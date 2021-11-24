package oneview

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccVersion(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVersionkData,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("oneview_version.test", "current_version", "120"),
					resource.TestCheckResourceAttr("oneview_version.test", "minimum_version", "3000"),
				),
			},
		},
	})
}

var testAccVersionkData = `
  data "oneview_version" "test" {    
  }`
