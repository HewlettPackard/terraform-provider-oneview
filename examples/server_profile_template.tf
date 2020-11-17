provider "oneview" {
  ov_username =   "${var.username}"
  ov_password =   "${var.password}"
  ov_endpoint =   "${var.endpoint}"
  ov_sslverify =  "${var.ssl_enabled}"
  ov_apiversion = 2200
  ov_ifmatch = "*"
}

data "oneview_scope" "scope" {
        name = "testing"
}
/*
resource "oneview_server_profile_template" "ServerProfileTemplate" {
        name = "TestServerProfileTemplate"
        type = "ServerProfileTemplateV8"
        enclosure_group = "EG-Synergy-Local"
        server_hardware_type = "SY 480 Gen9 1"
        initial_scope_uris = ["${data.oneview_scope.scope.uri}"]
}
*/
/* 	Update 
resource "oneview_server_profile_template" "ServerProfileTemplate" {
	name = "TestServerProfileTemplate_Renamed"
	type = "ServerProfileTemplateV8"
	enclosure_group = "enclosureGp"
	server_hardware_type = "SY 480 Gen9 1"
	initial_scope_uris = ["${data.oneview_scope.scope.uri}"]
}
*/
/* 	Datasource 
data "oneview_server_profile_template" "server_profile_template" {
	name = "TestServerProfileTemplate"
}

output "oneiew_server_hardware_type_value" {
	value = "${data.oneview_server_profile_template.server_profile_template.uri}"
}
*/
