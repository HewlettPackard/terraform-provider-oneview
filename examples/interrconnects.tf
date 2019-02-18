provider "oneview" {
	ov_username = "<username>"
	ov_password = "<password>"
	ov_endpoint = "<endpoint>"
	ov_sslverify = false
	ov_apiversion = 800
}

resource "oneview_interconnect" "InterconnectImport" {}

/* Run Example

terraform import oneview_interconnect.InterconnectImport "EN1, interconnect 3"

*/
