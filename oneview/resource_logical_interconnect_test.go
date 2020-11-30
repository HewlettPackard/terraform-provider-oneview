// (C) Copyright 2019 Hewlett Packard Enterprise Development LP
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
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

func TestAccLogicalInterconnect(t *testing.T) {
	var logicalInterconnect ov.LogicalInterconnect

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAcclogicalInterconnect,
				Check: resource.ComposeTestCheckFunc(
					testAccChecklogicalInterconnectExists(
						"oneview_logical_interconnect.test", &logicalInterconnect),
					resource.TestCheckResourceAttr(
						"oneview_logical_interconnect.test", "uri", "Terraform le 1",
					),
				),
			},
			{
				Config: testAcclogicalInterconnectUpdated,
				Check: resource.ComposeTestCheckFunc(
					testAccChecklogicalInterconnectExists(
						"oneview_logical_interconnect.test", &logicalInterconnect),
					resource.TestCheckResourceAttr(
						"oneview_logical_interconnect.test", "uri", "Terraform le 2",
					),
				),
			},
			{
				ResourceName:      testAcclogicalInterconnect,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccChecklogicalInterconnectExists(n string, logicalInterconnect *ov.LogicalInterconnect) resource.TestCheckFunc {
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

		testlogicalInterconnect, err := config.ovClient.GetLogicalInterconnectById(rs.Primary.ID)
		if err != nil {
			return err
		}
		if testlogicalInterconnect.Name != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*logicalInterconnect = testlogicalInterconnect
		return nil
	}
}

func testAccChecklogicalInterconnectDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview_logical_interconnect" {
			continue
		}

		testNet, _ := config.ovClient.GetLogicalInterconnectById(rs.Primary.ID)

		if testNet.Name != "" {
			return fmt.Errorf("LogicalInterconnect still exists")
		}
	}

	return nil
}

var testAcclogicalInterconnect = `
  resource "oneview_logical_interconnect" "test" {
    name = "Terraform le 1"
  }`

var testAcclogicalInterconnectUpdated = `
  resource "oneview_logical_interconnect" "test" {
    name = "Terraform le 2"
  }`
