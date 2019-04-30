provider "oneview" {
        ov_username = "<ov-username>"
        ov_password = "<ov-password>"
        ov_endpoint = "<ov-endpoint>"
        ov_sslverify = false
        ov_apiversion = 800
        ov_ifmatch = "*"
}

data "oneview_interconnect_type" "interconnect_type" {
	name = "Synergy 10Gb Pass-Thru Module"
}

output "oneview_interconnect_type_value" {
	value = "${data.oneview_interconnect_type.interconnect_type.type}"
//	value = "Null"
}
