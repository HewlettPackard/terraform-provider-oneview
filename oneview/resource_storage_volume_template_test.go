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

func TestAccStorageVolumeTemplate_1(t *testing.T) {
	var storageVolumeTemplate ov.StorageVolumeTemplate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStorageVolumeTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccStorageVolumeTemplate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageVolumeTemplateExists(
						"oneview_storage_volume_template.test", &storageVolumeTemplate),
					resource.TestCheckResourceAttr(
						"oneview_storage_volume_template.test", "name", "DemoStorageVolumeTemplate",
					),
				),
			},
			{
				ResourceName:      testAccStorageVolumeTemplate,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckStorageVolumeTemplateExists(n string, storageVolumeTemplate *ov.StorageVolumeTemplate) resource.TestCheckFunc {
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

		testStorageVolumeTemplate, err := config.ovClient.GetStorageVolumeTemplateByName(rs.Primary.ID)
		if err != nil {
			return err
		}
		if testStorageVolumeTemplate.Name != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*storageVolumeTemplate = testStorageVolumeTemplate
		return nil
	}
}

func testAccCheckStorageVolumeTemplateDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview_storage_volume_template" {
			continue
		}

		testSVT, _ := config.ovClient.GetStorageVolumeTemplateByName(rs.Primary.ID)

		if testSVT.Name != "" {
			return fmt.Errorf("StorageVolumeTemplate still exists")
		}
	}

	return nil
}

var testAccStorageVolumeTemplate = `resource "oneview_storage_volume_template" "test" {
  count = 1
  name = "DemoStorageVolumeTemplate"
  root_template_uri = "/rest/storage-volume-templates/96196d4c-3cac-4d6b-ab6b-a93c0143ac75"
  description = "Test SVT"
}`
