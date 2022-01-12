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

func TestAccTask_1(t *testing.T) {
	var task ov.Task

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTaskDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTask,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTaskExists(
						"oneview_task.test", &task),
					resource.TestCheckResourceAttr(
						"oneview_task.test", "task_id", "tgh4564-yfj7893",
					),
				),
			},
			{
				ResourceName:      testAccTask,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckTaskExists(n string, task *ov.Task) resource.TestCheckFunc {
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

		testTask, err := config.ovClient.GetTasksById("", "", "", "", rs.Primary.ID)
		if err != nil {
			return err
		}
		*task = testTask
		return nil
	}
}

func testAccCheckTaskDestroy(s *terraform.State) error {

	return nil
}

var testAccTask = `resource "oneview_task" "test" {
  task_id = "tgh4564-yfj7893"
}`
