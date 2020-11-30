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

func TestAccStorageSystem_1(t *testing.T) {
	var storageSystem ov.StorageSystem

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStorageSystemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccStorageSystem,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageSystemExists(
						"oneview_storage_system.test", &storageSystem),
					resource.TestCheckResourceAttr(
						"oneview_storage_system.test", "name", "Terraform le 1",
					),
				),
			},
			{
				Config: testAccStorageSystemUpdated,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageSystemExists(
						"oneview_storage_system.test", &storageSystem),
					resource.TestCheckResourceAttr(
						"oneview_storage_system.test", "name", "Terraform le 2",
					),
				),
			},
			{
				ResourceName:      testAccStorageSystem,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckStorageSystemExists(n string, storageSystem *ov.StorageSystem) resource.TestCheckFunc {
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

		testStorageSystem, err := config.ovClient.GetStorageSystemByName(rs.Primary.ID)
		if err != nil {
			return err
		}
		if testStorageSystem.Name != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*storageSystem = testStorageSystem
		return nil
	}
}

func testAccCheckStorageSystemDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview_storage_system" {
			continue
		}

		testNet, _ := config.ovClient.GetStorageSystemByName(rs.Primary.ID)

		if testNet.Name != "" {
			return fmt.Errorf("NetworkSet still exists")
		}
	}

	return nil
}

var testAccStorageSystem = `
  resource "oneview_storage_system" "test" {
    name = "Terraform le 1"
  }`

var testAccStorageSystemUpdated = `
  resource "oneview_storage_system" "test" {
    name = "Terraform le 2"
  }`
