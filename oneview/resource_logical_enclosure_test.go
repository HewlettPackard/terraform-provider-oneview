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

func TestAccLogicalEnclosure_1(t *testing.T) {
	var logicalEnclosure ov.LogicalEnclosure

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLogicalEnclosureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLogicalEnclosure,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLogicalEnclosureExists(
						"oneview_logical_enclosure.test", &logicalEnclosure),
					resource.TestCheckResourceAttr(
						"oneview_logical_enclosure.test", "name", "Terraform le 1",
					),
				),
			},
			{
				Config: testAccLogicalEnclosureUpdated,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLogicalEnclosureExists(
						"oneview_logical_enclosure.test", &logicalEnclosure),
					resource.TestCheckResourceAttr(
						"oneview_logical_enclosure.test", "name", "Terraform le 2",
					),
				),
			},
			{
				ResourceName:      testAccLogicalEnclosure,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckLogicalEnclosureExists(n string, logicalEnclosure *ov.LogicalEnclosure) resource.TestCheckFunc {
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

		testLogicalEnclosure, err := config.ovClient.GetLogicalEnclosureByName(rs.Primary.ID)
		if err != nil {
			return err
		}
		if testLogicalEnclosure.Name != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*logicalEnclosure = testLogicalEnclosure
		return nil
	}
}

func testAccCheckLogicalEnclosureDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview_logical_enclosure" {
			continue
		}

		testNet, _ := config.ovClient.GetLogicalEnclosureByName(rs.Primary.ID)

		if testNet.Name != "" {
			return fmt.Errorf("NetworkSet still exists")
		}
	}

	return nil
}

var testAccLogicalEnclosure = `
  resource "oneview_logical_enclosure" "test" {
    name = "Terraform le 1"
  }`

var testAccLogicalEnclosureUpdated = `
  resource "oneview_logical_enclosure" "test" {
    name = "Terraform le 2"
  }`
