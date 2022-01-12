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

func TestAccApplianceSNMPv1TrapDestinations_1(t *testing.T) {
	var ApplianceSNMPv1TrapDestinations ov.SNMPv1Trap

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckApplianceSNMPv1TrapDestinationsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplianceSNMPv1TrapDestinations,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplianceSNMPv1TrapDestinationsExists(
						"oneview_appliance_snmp_v1_trap_destinations.test", &ApplianceSNMPv1TrapDestinations),
					resource.TestCheckResourceAttr(
						"oneview_appliance_snmp_v1_trap_destinations.test", "destination_address", "1.1.1.1",
					),
					resource.TestCheckResourceAttr(
						"oneview_appliance_snmp_v1_trap_destinations.test", "destination_id", "4",
					),
					resource.TestCheckResourceAttr(
						"oneview_appliance_snmp_v1_trap_destinations.test", "community_string", "Test1",
					),
					resource.TestCheckResourceAttr(
						"oneview_appliance_snmp_v1_trap_destinations.test", "port", "162",
					),
				),
			},
			{
				ResourceName:      testAccApplianceSNMPv1TrapDestinations,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckApplianceSNMPv1TrapDestinationsExists(n string, SNMPv1TrapDestinations *ov.SNMPv1Trap) resource.TestCheckFunc {
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

		testSNMPv1TrapDestinations, err := config.ovClient.GetSNMPv1TrapDestinationsById(rs.Primary.ID)
		if err != nil {
			return err
		}
		*SNMPv1TrapDestinations = testSNMPv1TrapDestinations
		return nil
	}
}

func testAccCheckApplianceSNMPv1TrapDestinationsDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview_appliance_snmp_v1_trap_destinations" {
			continue
		}

		testTrap, _ := config.ovClient.GetSNMPv1TrapDestinationsById(rs.Primary.ID)

		if testTrap.Destination != "" {
			return fmt.Errorf("ApplianceSNMPv1TrapDestinations still exists")
		}
	}

	return nil
}

var testAccApplianceSNMPv1TrapDestinations = `resource "oneview_appliance_snmp_v1_trap_destinations" "test" {
    destination_address = "1.1.1.1"
    port = 162
    community_string = "Test1"
    destination_id = "4"
}`
