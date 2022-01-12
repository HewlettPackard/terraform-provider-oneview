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

func TestAccSNMPv3User_1(t *testing.T) {
	var SNMPv3User ov.SNMPv3User

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSNMPv3UserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSNMPv3User,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSNMPv3UserExists(
						"oneview_appliance_snmpv3_user.test", &SNMPv3User),
					resource.TestCheckResourceAttr(
						"oneview_appliance_snmpv3_user.test", "destination_address", "test_destination_address",
					),
					resource.TestCheckResourceAttr(
						"oneview_appliance_snmpv3_user.test", "user_id", "test_user_id",
					),
					resource.TestCheckResourceAttr(
						"oneview_appliance_snmpv3_user.test", "port", "162",
					),
				),
			},
			{
				ResourceName:      testAccSNMPv3User,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckSNMPv3UserExists(n string, SNMPv3User *ov.SNMPv3User) resource.TestCheckFunc {
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

		testSNMPv3User, err := config.ovClient.GetSNMPv3UserByUserName(rs.Primary.ID)
		if err != nil {
			return err
		}
		if testSNMPv3User.Id != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*SNMPv3User = testSNMPv3User
		return nil
	}
}

func testAccCheckSNMPv3UserDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview_appliance_snmpv3_user" {
			continue
		}

		testUser, _ := config.ovClient.GetSNMPv3UserByUserName(rs.Primary.ID)

		if testUser.Id != "" {
			return fmt.Errorf("SNMPv3User still exists")
		}
	}

	return nil
}

var testAccSNMPv3User = `resource "oneview_appliance_snmpv3_user" "test" {
  user_name                 = "user"
  security_level            = "Authentication and privacy"
  authentication_protocol   = "SHA1"
  authentication_passphrase = "authPass"
  privacy_protocol          = "AES-128"
  privacy_passphrase        = "1234567812345678"
}`
