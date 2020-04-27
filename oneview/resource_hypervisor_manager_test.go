// (C) Copyright 2016 Hewlett Packard Enterprise Development LP
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
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccHypervisorManager_1(t *testing.T) {
	var fcNetwork ov.HypervisorManager

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHypervisorManagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHypervisorManager,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHypervisorManagerExists(
						"oneview_hypervisor_manager.test", &fcNetwork),
					resource.TestCheckResourceAttr(
						"oneview_hypervisor_manager.test", "name", "Terraform FC Network 1",
					),
				),
			},
			{
				ResourceName:      testAccHypervisorManager,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckHypervisorManagerExists(n string, fcNetwork *ov.HypervisorManager) resource.TestCheckFunc {
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

		testHypervisorManager, err := config.ovClient.GetHypervisorManagerByName(rs.Primary.ID)
		if err != nil {
			return err
		}
		if testHypervisorManager.Name != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*fcNetwork = testHypervisorManager
		return nil
	}
}

func testAccCheckHypervisorManagerDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview_hypervisor_manager" {
			continue
		}

		testFCNet, _ := config.ovClient.GetHypervisorManagerByName(rs.Primary.ID)

		if testFCNet.Name != "" {
			return fmt.Errorf("HypervisorManager still exists")
		}
	}

	return nil
}

var testAccHypervisorManager = `resource "oneview_hypervisor_manager" "test" {
  count = 1
  name = "Terraform FC Network 1"
}`
