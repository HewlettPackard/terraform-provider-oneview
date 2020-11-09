provider "oneview" {
  ov_username =   "${var.username}"
  ov_password =   "${var.password}"
  ov_endpoint =   "${var.endpoint}"
  ov_sslverify =  "${var.ssl_enabled}"
  ov_apiversion = 2200
  ov_ifmatch = "*"
}

data "oneview_interconnect" "interconnect" {
	name = "Synergy-Encl-1, interconnect 6"
}

output "oneiew_interconnect_value" {
	value = "${data.oneview_interconnect.interconnect.uri}"
}
