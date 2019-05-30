provider "oneview" {
        ov_username = "<ov-username>"
        ov_password = "<ov-password>"
        ov_endpoint = "<ov-endpoint>"
        ov_sslverify = false
        ov_apiversion = <ov-apiversion>
        ov_ifmatch = "*"
}

data "oneview_server_hardware" "server_hardware" {
	name = "0000A66101, bay 3"
}

output "oneview_server_hardware_value" {
	value = "${data.oneview_server_hardware.server_hardware.uri}"
}
