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

func TestAccFirmwareDrivers_1(t *testing.T) {
	var FirmwareDrivers ov.FirmwareDrivers

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFirmwareDriversDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFirmwareDrivers,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFirmwareDriversExists(
						"oneview_firmware_drivers.test", &FirmwareDrivers),
					resource.TestCheckResourceAttr(
						"oneview_firmware_drivers.test", "custom_baseline_name", "Terraform_SPP",
					),
					resource.TestCheckResourceAttr(
						"oneview_firmware_drivers.test", "baseline_uri", "abc-abc-abc",
					),
				),
			},
			{
				ResourceName:      testAccFirmwareDrivers,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckFirmwareDriversExists(n string, FirmwareDrivers *ov.FirmwareDrivers) resource.TestCheckFunc {
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

		testFirmwareDrivers, err := config.ovClient.GetFirmwareBaselineById(rs.Primary.ID)
		if err != nil {
			return err
		}
		if testFirmwareDrivers.Name != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*FirmwareDrivers = testFirmwareDrivers
		return nil
	}
}

func testAccCheckFirmwareDriversDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview_firmware_drivers" {
			continue
		}

		testFirmware, _ := config.ovClient.GetFirmwareBaselineById(rs.Primary.ID)

		if testFirmware.Name != "" {
			return fmt.Errorf("FirmwareDrivers still exists")
		}
	}

	return nil
}

var testAccFirmwareDrivers = `resource "oneview_firmware_drivers" "test" {
    custom_baseline_name = "Terraform_SPP"
	baseline_uri = "abc-abc-abc"
}`
