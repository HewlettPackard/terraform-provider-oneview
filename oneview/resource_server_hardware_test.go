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
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccServerHardware_1(t *testing.T) {
	var serverHardware ov.ServerHardware

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckServerHardwareDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServerHardware,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerHardwareExists(
						"oneview_server_hardware.test", &serverHardware),
					resource.TestCheckResourceAttr(
						"oneview_server_hardware.test", "host_name", "172.1.1.1",
					),
				),
			},
			{
				ResourceName:      testAccServerHardware,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckServerHardwareExists(n string, serverHardware *ov.ServerHardware) resource.TestCheckFunc {
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

		testServerHardware, err := config.ovClient.GetServerHardwareByName(rs.Primary.ID)
		if err != nil {
			return err
		}
		if testServerHardware.Name != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*serverHardware = testServerHardware
		return nil
	}
}

func testAccCheckServerHardwareDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview_server_hardware" {
			continue
		}

		testSH, _ := config.ovClient.GetServerHardwareByUri(utils.Nstring(rs.Primary.ID))

		if testSH.URI != "" {
			return fmt.Errorf("Server Hardware still exists")
		}
	}

	return nil
}

var testAccServerHardware = `
  resource "oneview_server_hardware" "sh" {
	host_name   = "172.1.1.1"
  	username = "dcs"
  	password = "dcs"
  	licensing_intent = "OneView"
  }`
