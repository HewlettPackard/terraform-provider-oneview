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

func TestAccLabels_1(t *testing.T) {
	var label ov.AssignedLabel

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLabelsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLabels,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLabelsExists(
						"oneview_labels.test", &label),
					resource.TestCheckResourceAttr(
						"oneview_labels.test", "resource_uri", "/rest/labels/resources/resource_uri",
					),
				),
			},
			{
				ResourceName:      testAccLabels,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckLabelsExists(n string, label *ov.AssignedLabel) resource.TestCheckFunc {
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

		testLabel, err := config.ovClient.GetAssignedLabels(utils.Nstring(rs.Primary.ID))
		if err != nil {
			return err
		}
		if testLabel.ResourceUri.String() != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}
		*label = testLabel
		return nil
	}
}

func testAccCheckLabelsDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oneview_labels" {
			continue
		}

		testLabel, _ := config.ovClient.GetAssignedLabels(utils.Nstring(rs.Primary.ID))

		if testLabel.ResourceUri != "" {
			return fmt.Errorf("Labels still exists")
		}
	}

	return nil
}

var testAccLabels = `resource "oneview_labels" "test" {
  resource_uri = "/rest/labels/resources/resource_uri"
}`
