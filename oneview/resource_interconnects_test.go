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

func TestAccInterconnects_1(t *testing.T) {
	var interconnect ov.Interconnect

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccInterconnect,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInterconnectExists(
						"oneview_interconnect.test", &interconnect),
					resource.TestCheckResourceAttr(
						"oneview_interconnect.test", "name", "Terraform Interconnect 1",
					),
				),
			},
			{
				ResourceName:      testAccInterconnect,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckInterconnectExists(n string, interconnect *ov.Interconnect) resource.TestCheckFunc {
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

		testInterconnect, err := config.ovClient.GetInterconnectByName(rs.Primary.ID)
		if err != nil {
			return err
		}
		if testInterconnect.Name != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*interconnect = testInterconnect
		return nil
	}
}

var testAccInterconnect = `resource "oneview_interconnect" "test" {
  count = 1
  name = "Terraform Interconnect 1"
}`
