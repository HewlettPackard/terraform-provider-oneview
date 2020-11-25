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

func TestAccVolume_1(t *testing.T) {
	var volume ov.StorageVolume

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUplinkSet,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVolumeExists(
						"oneview_volume.test", &volume),
					resource.TestCheckResourceAttr(
						"oneview_volume.test", "name", "terraform volume",
					),
				),
			},
			{
				ResourceName:      testAccVolume,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckVolumeExists(n string, volume *ov.StorageVolume) resource.TestCheckFunc {
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

		testVolume, err := config.ovClient.GetStorageVolumeByName(rs.Primary.ID)
		if err != nil {
			return err
		}
		if testVolume.Name != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*volume = testVolume
		return nil
	}
}

func testAccCheckVolumeDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview_volume" {
			continue
		}

		testVolume, _ := config.ovClient.GetStorageVolumeByName(rs.Primary.ID)

		if testVolume.Name != "" {
			return fmt.Errorf("Volume still exists")
		}
	}

	return nil
}

var testAccVolume = `resource "oneview_volume" "test" {
    count = 1
    name = "terraform volume"
    port_config_infos {}
    network_uris {}
    fc_network_uris {}
    fcoe_network_uris {}
  }`
