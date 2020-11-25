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

func TestAccEnlosureGroup_1(t *testing.T) {
	var enclosureGroup ov.EnclosureGroup

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckEnclosureGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccEnclosureGroup,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEnclosureGroupExists(
						"oneview-enclosure_group.test", &enclosureGroup),
					resource.TestCheckResourceAttr(
						"oneview-enclosure_group.test", "name", "terraform enclosure group",
					),
				),
			},
			{
				ResourceName:      testAccEnclosureGroup,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckEnclosureGroupExists(n string, enclosureGroup *ov.EnclosureGroup) resource.TestCheckFunc {
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

		testEnclosureGroup, err := config.ovClient.GetEnclosureGroupByName(rs.Primary.ID)
		if err != nil {
			return err
		}
		if testEnclosureGroup.Name != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*enclosureGroup = testEnclosureGroup
		return nil
	}
}

func testAccCheckEnclosureGroupDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview-enclosure_group" {
			continue
		}

		testEg, _ := config.ovClient.GetEnclosureGroupByName(rs.Primary.ID)

		if testEg.Name != "" {
			return fmt.Errorf("Enclsoure still exists")
		}
	}

	return nil
}

var testAccEnclosureGroup = `resource "oneview-enclosure_group" "test" {
    count = 1
    name = "terraform enclosure group"	
  }`
