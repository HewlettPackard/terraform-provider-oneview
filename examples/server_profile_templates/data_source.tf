provider "oneview" {
  ov_username =   "${var.username}"
  ov_password =   "${var.password}"
  ov_endpoint =   "${var.endpoint}"
  ov_sslverify =  "${var.ssl_enabled}"
  ov_apiversion = 2600
  ov_ifmatch = "*"
}

# Datasource for server profile template
data "oneview_server_profile_template" "server_profile_template" {
	name = "TestServerProfileTemplateRenamed"
}

output "oneiew_server_hardware_type_value" {
	value = "${data.oneview_server_profile_template.server_profile_template.uri}"
}

