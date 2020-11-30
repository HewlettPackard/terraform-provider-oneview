// (C) Copyright 2016 Hewlett Packard Enterprise Development LP
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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccUplinkSet_1(t *testing.T) {
	var uplinkSet ov.UplinkSet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckUplinkSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUplinkSet,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUplinkSetExists(
						"oneview-uplink_set.test", &uplinkSet),
					resource.TestCheckResourceAttr(
						"oneview-uplink_set.test", "name", "terraform uplink set",
					),
				),
			},
			{
				ResourceName:      testAccUplinkSet,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckUplinkSetExists(n string, uplinkSet *ov.UplinkSet) resource.TestCheckFunc {
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

		testUplinkSet, err := config.ovClient.GetUplinkSetByName(rs.Primary.ID)
		if err != nil {
			return err
		}
		if testUplinkSet.Name != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*uplinkSet = testUplinkSet
		return nil
	}
}

func testAccCheckUplinkSetDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview-uplink_set" {
			continue
		}

		testLig, _ := config.ovClient.GetUplinkSetByName(rs.Primary.ID)

		if testLig.Name != "" {
			return fmt.Errorf("UplinkSet still exists")
		}
	}

	return nil
}

var testAccUplinkSet = `resource "oneview-uplink_set" "test" {
    count = 1
    name = "terraform uplink set"
    port_config_infos {}
    network_uris {}
    fc_network_uris {}
    fcoe_network_uris {}
  }`
