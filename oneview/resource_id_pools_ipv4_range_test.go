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

func TestAccIPv4Range_1(t *testing.T) {
	var ranges ov.Ipv4Range

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRangeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRange,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRangeExists(
						"oneview_id_pools_ipv4_range.test", &ranges),
					resource.TestCheckResourceAttr(
						"oneview_id_pools_ipv4_range.test", "name", "Range",
					),
					resource.TestCheckResourceAttr(
						"oneview_id_pools_ipv4_range.test", "subnet_uri", "/resst/fake",
					),
				),
			},
			{
				ResourceName:      testAccRange,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckRangeExists(n string, ranges *ov.Ipv4Range) resource.TestCheckFunc {
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

		testRange, err := config.ovClient.GetIPv4RangebyId("", rs.Primary.ID)
		if err != nil {
			return err
		}
		*ranges = testRange
		return nil
	}
}

func testAccCheckRangeDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview_id_pools_ipv4_range" {
			continue
		}

		testNet, _ := config.ovClient.GetIPv4RangebyId("", rs.Primary.ID)

		if testNet.URI != "" {
			return fmt.Errorf("Range still exists")
		}
	}

	return nil
}

var testAccRange = `
  resource "oneview_id_pools_ipv4_range" "test" {
       name="IpRange"
       subnet_uri="/rest/fake"
  }`
