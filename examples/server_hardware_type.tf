provider "oneview" {
	ov_username = "<ov_username>"
	ov_password = "<ov_password>"
	ov_endpoint = "<ov_endpoint>"
	ov_sslverify = false
	ov_apiversion = <ov_apiversion>
	ov_ifmatch = "*"
}

data "oneview_server_hardware_type" "server_hardware_type" {
	name = "SY 480 Gen9 2"
}

output "oneiew_server_hardware_type_value" {
	value = "${data.oneview_server_hardware_type.server_hardware_type.uri}"
}
