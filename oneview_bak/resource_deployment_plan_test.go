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

	"github.com/HewlettPackard/oneview-golang/i3s"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccDeploymentPlan_1(t *testing.T) {
	var deploymentPlan i3s.DeploymentPlan

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDeploymentPlanDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDeploymentPlan,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeploymentPlanExists(
						"oneview-deployment_plan.test", &deploymentPlan),
					resource.TestCheckResourceAttr(
						"oneview-deployment_plan.test", "name", "terraform deployment plan",
					),
				),
			},
			{
				ResourceName:      testAccDeploymentPlan,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckDeploymentPlanExists(n string, deploymentPlan *i3s.DeploymentPlan) resource.TestCheckFunc {
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

		testDeploymentPlan, err := config.i3sClient.GetDeploymentPlanByName(rs.Primary.ID)
		if err != nil {
			return err
		}
		if testDeploymentPlan.Name != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*deploymentPlan = testDeploymentPlan
		return nil
	}
}

func testAccCheckDeploymentPlanDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview-deployment_plan" {
			continue
		}

		testLig, _ := config.i3sClient.GetDeploymentPlanByName(rs.Primary.ID)

		if testLig.Name != "" {
			return fmt.Errorf("DeploymenttPlan still exists")
		}
	}

	return nil
}

var testAccDeploymentPlan = `resource "oneview-deployment_plan" "test" {
    count = 1
    name = "terraform deployment plan"
    oe_build_plan_uri ="rest/build-plans/S8T45F"
    type = "OEDeploymentPlanV5"
  }`
