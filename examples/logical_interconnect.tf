provider "oneview" {
	ov_username = "<ov_username>"
	ov_password = "<ov_password>"
	ov_endpoint = "<ov_endpoint>"
	ov_sslverify = false
	ov_apiversion = <ov_apiversion>
	ov_ifmatch = "*"
}

data "oneview_logical_interconnect" "logical_interconnect" {
	name = "d4468f89-4442-4324-9c01-624c7382db2d"
}

output "oneview_logical_interconnect_value" {
	value = "${data.oneview_logical_interconnect.logical_interconnect.uri}"
}
