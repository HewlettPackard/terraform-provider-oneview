provider "oneview" {
        ov_username = "<ov-username>"
        ov_password = "<ov-password>"
        ov_endpoint = "<ov-endpoint>"
        ov_sslverify = false
        ov_apiversion = <ov_apiversion>
        ov_ifmatch = "*"
}

data "oneview_interconnect_type" "interconnect_type" {
	name = "Virtual Connect SE 40Gb F8 Module for Synergy"
}

output "oneview_interconnect_type_value" {
	value = "${data.oneview_interconnect_type.interconnect_type.type}"
}
