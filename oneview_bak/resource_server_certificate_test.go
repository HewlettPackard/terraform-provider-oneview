// (C) Copyright 2020 Hewlett Packard Enterprise Development LP
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

func TestAccServerCertificate_1(t *testing.T) {
	var serverCertificate ov.ServerCertificate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckServerCertificateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServerCertificate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerCertificateExists(
						"oneview_server_certificate.test", &serverCertificate),
					resource.TestCheckResourceAttr(
						"oneview_server_certificate.test", "alias_name", "Terraform Server Certificate1",
					),
				),
			},
			{
				ResourceName:      testAccServerCertificate,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckServerCertificateExists(n string, serverCertificate *ov.ServerCertificate) resource.TestCheckFunc {
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

		testServerCertificate, err := config.ovClient.GetServerCertificateByName(rs.Primary.ID)
		if err != nil {
			return err
		}
		if testServerCertificate.Name != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*serverCertificate = testServerCertificate
		return nil
	}
}

func testAccCheckServerCertificateDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview_server_certificate" {
			continue
		}

		testServerCertificate, _ := config.ovClient.GetServerCertificateByName(rs.Primary.ID)

		if testServerCertificate.Name != "" {
			return fmt.Errorf("ServerCertificate still exists")
		}
	}

	return nil
}

var testAccServerCertificate = `resource "oneview_server_certificate" "test" {
  count = 1
  alias_name = "Terraform Server Certificate1"
}`
