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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccServerCertificate_2(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccServerCertificateData,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("oneview_server_certificate.test", "type", "CertificateInfoV2"),
					resource.TestCheckResourceAttr("oneview_server_certificate.test", "alias_name", "TestServerCertificate"),
				),
			},
		},
	})
}

var testAccServerCertificateData = `
  data "oneview_server_certificate" "test" {
    type = "CertificateInfoV2"		  
    alias_name = "TestServerCertificate"
  }`
