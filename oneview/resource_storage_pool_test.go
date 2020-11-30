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

func TestAccStoragePool_1(t *testing.T) {
	var storagePool ov.StoragePool

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccStoragePool,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStoragePoolExists(
						"oneview_storage_pool.test", &storagePool),
					resource.TestCheckResourceAttr(
						"oneview_storage_pool.test", "name", "Test",
					),
				),
			},
			{
				Config: testAccStoragePoolUpdated,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStoragePoolExists(
						"oneview_storage_pool.test", &storagePool),
					resource.TestCheckResourceAttr(
						"oneview_storage_pool.test", "name", "Test2",
					),
				),
			},
			{
				ResourceName:      testAccStoragePool,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckStoragePoolExists(n string, storagePool *ov.StoragePool) resource.TestCheckFunc {
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

		testStoragePool, err := config.ovClient.GetStoragePoolByName(rs.Primary.ID)
		if err != nil {
			return err
		}
		if testStoragePool.Name != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*storagePool = testStoragePool
		return nil
	}
}

var testAccStoragePool = `
  resource "oneview_storage_pool" "test" {
    name = "Test"
  }`

var testAccStoragePoolUpdated = `
  resource "oneview_storage_pool" "test" {
    name = "Test2"
  }`
