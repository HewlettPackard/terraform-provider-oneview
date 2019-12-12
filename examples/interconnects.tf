provider "oneview" {
	ov_username = "<username>"
	ov_password = "<password>"
	ov_endpoint = "<endpoint>"
	ov_sslverify = false
	ov_apiversion = <ov_apiversion>
}

data "oneview_interconnect" "interconnect" {
	name = "EN1, interconnect 3"
}

output "oneiew_interconnect_value" {
	value = "${data.oneview_interconnect.interconnect.uri}"
}
