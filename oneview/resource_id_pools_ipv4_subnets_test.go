// (C) Copyright 2021 Hewlett Packard Enterprise Development LP
//
// Licensed under the Apache License, Version 2.0 (the "License");
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed
// under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package oneview

import (
	"fmt"
	"testing"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIPv4Subnets_1(t *testing.T) {
	var subnet ov.Ipv4Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSubnet,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(
						"oneview_id_pools_ipv4_subnets.test", &subnet),
					resource.TestCheckResourceAttr(
						"oneview_id_pools_ipv4_subnets.test", "network_id", "10.1.0.0",
					),
					resource.TestCheckResourceAttr(
						"oneview_id_pools_ipv4_subnets.test", "subnet_mask", "255.255.0.1",
					),
					resource.TestCheckResourceAttr(
						"oneview_id_pools_ipv4_subnets.test", "gateway", "10.1.0.1",
					),
				),
			},
			{
				ResourceName:      testAccSubnet,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckSubnetExists(n string, subnet *ov.Ipv4Subnet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found :%v", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config, err := testProviderConfig()
		if err != nil {
			return err
		}

		testSubnet, err := config.ovClient.GetIPv4SubnetbyId(rs.Primary.ID)
		if err != nil {
			return err
		}
		*subnet = testSubnet
		return nil
	}
}

func testAccCheckSubnetDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview_id_pools_ipv4_subnets" {
			continue
		}

		testNet, _ := config.ovClient.GetIPv4SubnetbyId(rs.Primary.ID)

		if testNet.NetworkId != "" {
			return fmt.Errorf("Subnet still exists")
		}
	}

	return nil
}

var testAccSubnet = `
  resource "oneview_id_pools_ipv4_subnets" "test" {
       network_id="10.1.0.0"
       subnet_mask="255.255.0.1"
       gateway="10.1.0.1"
  }`
