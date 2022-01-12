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

func TestAccSshAccess_1(t *testing.T) {
	var sshAccess ov.ApplianceSshAccess

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccsshAccess,
				Check: resource.ComposeTestCheckFunc(
					testAccChecksshAccessExists(
						"oneview_appliance_ssh_access.test", &sshAccess),
					resource.TestCheckResourceAttr(
						"oneview_appliance_ssh_access.test", "allow_ssh_access", "true",
					),
				),
			},
			{
				Config: testAccsshAccessUpdated,
				Check: resource.ComposeTestCheckFunc(
					testAccChecksshAccessExists(
						"oneview_appliance_ssh_access.test", &sshAccess),
					resource.TestCheckResourceAttr(
						"oneview_appliance_ssh_access.test", "allow_ssh_access", "false",
					),
				),
			},
			{
				ResourceName:      testAccsshAccess,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccChecksshAccessExists(n string, sshAccess *ov.ApplianceSshAccess) resource.TestCheckFunc {
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

		testsshAccess, err := config.ovClient.GetSshAccess()
		if err != nil {
			return err
		}
		if testsshAccess.Type != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*sshAccess = testsshAccess
		return nil
	}
}

func testAccsshAccessDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview_appliance_ssh_access" {
			continue
		}

		testNet, _ := config.ovClient.GetSshAccess()

		if testNet.Type != "" {
			return fmt.Errorf("sshAccess still exists")
		}
	}

	return nil
}

var testAccsshAccess = `
  resource "oneview_appliance_ssh_access" "test" {
    allow_ssh_access = true
  }`

var testAccsshAccessUpdated = `
  resource "oneview_connection_templates" "test" {
    allow_ssh_access = false
  }`
