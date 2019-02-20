provider "oneview" {
	ov_username = "<ov_username>"
	ov_password = "<ov_password>"
	ov_endpoint = "<ov_endpoint>"
	ov_sslverify = false
	ov_apiversion = <ov_apiversion>
	ov_ifmatch = "*"
}

resource "oneview_logical_interconnect" "LogicalInterconnectImport" {}

/* Run

terraform import oneview_logical_interconnect.LogicalInterconnectImport "EN1, interconnect 3"

in terminal

*/