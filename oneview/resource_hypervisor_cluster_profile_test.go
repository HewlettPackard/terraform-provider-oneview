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

func TestAccHypervisorClusterProfile_1(t *testing.T) {
	var hypervisorClusterProfile ov.HypervisorClusterProfile

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHypervisorClusterProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHypervisorClusterProfile,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHypervisorClusterProfileExists(
						"oneview_hypervisor_manager.test", &hypervisorClusterProfile),
					resource.TestCheckResourceAttr(
						"oneview_hypervisor_manager.test", "name", "Terraform Hypervisor Cluster Profile1",
					),
				),
			},
			{
				ResourceName:      testAccHypervisorClusterProfile,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckHypervisorClusterProfileExists(n string, hypervisorClusterProfile *ov.HypervisorClusterProfile) resource.TestCheckFunc {
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

		testHypervisorClusterProfile, err := config.ovClient.GetHypervisorClusterProfileByName(rs.Primary.ID)
		if err != nil {
			return err
		}
		if testHypervisorClusterProfile.Name != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*hypervisorClusterProfile = testHypervisorClusterProfile
		return nil
	}
}

func testAccCheckHypervisorClusterProfileDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview_hypervisor_manager" {
			continue
		}

		testHypervisorClusterProfile, _ := config.ovClient.GetHypervisorClusterProfileByName(rs.Primary.ID)

		if testHypervisorClusterProfile.Name != "" {
			return fmt.Errorf("HypervisorClusterProfile still exists")
		}
	}

	return nil
}

var testAccHypervisorClusterProfile = `resource "oneview_hypervisor_manager" "test" {
  count = 1
  name = "Terraform Hypervisor Cluster Profile1"
}`
