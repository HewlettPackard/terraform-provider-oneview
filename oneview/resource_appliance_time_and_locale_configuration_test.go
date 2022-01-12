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

func TestAccTimeandLocaleConfiguration_1(t *testing.T) {
	var TimeandLocaleConfiguration ov.ApplianceTimeandLocal

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTimeandLocaleConfigurationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTimeandLocaleConfiguration,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTimeandLocaleConfigurationExists(
						"oneview_appliance_time_and_locale.test", &TimeandLocaleConfiguration),
					resource.TestCheckResourceAttr(
						"oneview_appliance_time_and_locale.test", "locale", "test_locale",
					),
					resource.TestCheckResourceAttr(
						"oneview_appliance_time_and_locale.test", "timezone", "test_zone",
					),
					resource.TestCheckResourceAttr(
						"oneview_appliance_time_and_locale.test", "ntpserver", "[\"server\"]",
					),
				),
			},
			{
				ResourceName:      testAccTimeandLocaleConfiguration,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckTimeandLocaleConfigurationExists(n string, TimeandLocaleConfiguration *ov.ApplianceTimeandLocal) resource.TestCheckFunc {
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

		testTimeandLocaleConfiguration, err := config.ovClient.GetApplianceTimeandLocals("", "", "", "")
		if err != nil {
			return err
		}
		if testTimeandLocaleConfiguration.Locale != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*TimeandLocaleConfiguration = testTimeandLocaleConfiguration
		return nil
	}
}

func testAccCheckTimeandLocaleConfigurationDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview_connection_templates" {
			continue
		}

		testNet, _ := config.ovClient.GetApplianceTimeandLocals("", "", "", "")

		if testNet.Locale != "" {
			return fmt.Errorf("TimeandLocaleConfiguration still exists")
		}
	}

	return nil
}

var testAccTimeandLocaleConfiguration = `
  resource "oneview_appliance_time_and_locale" "test" {
	locale = "test_locale"
    timezone = "test_zone"
	ntpserver = ["server"]
  }`
