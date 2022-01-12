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

func TestAccConnectionTemplate_1(t *testing.T) {
	var connectionTemplate ov.ConnectionTemplate

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccconnectionTemplate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckconnectionTemplateExists(
						"oneview_connection_templates.test", &connectionTemplate),
					resource.TestCheckResourceAttr(
						"oneview_connection_templates.test", "name", "Terraform ct 1",
					),
				),
			},
			{
				Config: testAccconnectionTemplateUpdated,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckconnectionTemplateExists(
						"oneview_connection_templates.test", &connectionTemplate),
					resource.TestCheckResourceAttr(
						"oneview_connection_templates.test", "name", "Terraform ct 2",
					),
				),
			},
			{
				ResourceName:      testAccconnectionTemplate,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckconnectionTemplateExists(n string, connectionTemplate *ov.ConnectionTemplate) resource.TestCheckFunc {
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

		testconnectionTemplate, err := config.ovClient.GetConnectionTemplateByName(rs.Primary.ID)
		if err != nil {
			return err
		}
		if testconnectionTemplate.Name != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*connectionTemplate = testconnectionTemplate
		return nil
	}
}

func testAccCheckconnectionTemplateDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview_connection_templates" {
			continue
		}

		testNet, _ := config.ovClient.GetConnectionTemplateByName(rs.Primary.ID)

		if testNet.Name != "" {
			return fmt.Errorf("connectionTemplate still exists")
		}
	}

	return nil
}

var testAccconnectionTemplate = `
  resource "oneview_connection_templates" "test" {
    name = "Terraform ct 1"
  }`

var testAccconnectionTemplateUpdated = `
  resource "oneview_connection_templates" "test" {
    name = "Terraform ct 2"
  }`
