provider "oneview" {
        ov_username = "<ov_username>"
        ov_password = "<ov_password"
        ov_endpoint = "<ov_endpoint>"
        ov_sslverify = false
        ov_apiversion = <ov_apiversion>
        ov_ifmatch = "*"
}

data "oneview_scope" "scope" {
        name = "test"
}
/*
resource "oneview_server_profile_template" "ServerProfileTemplate" {
	name = "TestServerProfileTemplate"
	type = "ServerProfileTemplateV8"
	enclosure_group = "enclosureGp"
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
/* 	Datasource */
data "oneview_server_profile_template" "server_profile_template" {
	name = "TestServerProfileTemplate"
}

output "oneiew_server_hardware_type_value" {
	value = "${data.oneview_server_profile_template.server_profile_template.uri}"
}

