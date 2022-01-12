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

func TestAccSNMPv3TrapDestinations_1(t *testing.T) {
	var SNMPv3TrapDestinations ov.SNMPv3Trap

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSNMPv3TrapDestinationsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSNMPv3TrapDestinations,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSNMPv3TrapDestinationsExists(
						"oneview_appliance_snmpv3_trap_destinations.test", &SNMPv3TrapDestinations),
					resource.TestCheckResourceAttr(
						"oneview_appliance_snmpv3_trap_destinations.test", "destination_address", "test_destination_address",
					),
					resource.TestCheckResourceAttr(
						"oneview_appliance_snmpv3_trap_destinations.test", "user_id", "test_user_id",
					),
					resource.TestCheckResourceAttr(
						"oneview_appliance_snmpv3_trap_destinations.test", "port", "162",
					),
				),
			},
			{
				ResourceName:      testAccSNMPv3TrapDestinations,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckSNMPv3TrapDestinationsExists(n string, SNMPv3TrapDestinations *ov.SNMPv3Trap) resource.TestCheckFunc {
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

		testSNMPv3TrapDestinations, err := config.ovClient.GetSNMPv3TrapDestinationsById(rs.Primary.ID)
		if err != nil {
			return err
		}
		if testSNMPv3TrapDestinations.ID != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*SNMPv3TrapDestinations = testSNMPv3TrapDestinations
		return nil
	}
}

func testAccCheckSNMPv3TrapDestinationsDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview_appliance_snmpv3_trap_destinations" {
			continue
		}

		testTrap, _ := config.ovClient.GetSNMPv3TrapDestinationsById(rs.Primary.ID)

		if testTrap.ID != "" {
			return fmt.Errorf("SNMPv3TrapDestinations still exists")
		}
	}

	return nil
}

var testAccSNMPv3TrapDestinations = `resource "oneview_appliance_snmpv3_trap_destinations" "test" {
    destination_address = "test_destination_address"
    port = 162
    user_id = "test_user_id"
}`
