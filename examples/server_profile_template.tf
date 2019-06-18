provider "oneview" {
	ov_username = "<ov_username>"
	ov_password = "<ov_password"
	ov_endpoint = "<ov_endpoint>"
	ov_sslverify = false
	ov_apiversion = <ov_apiversion>
	ov_ifmatch = "*"
}

resource "oneview_server_profile_template" "ServerProfileTemplate" {
	name = "TestServerProfileTemplate"
	type = "ServerProfileTemplateV5"
	enclosure_group = "SYN03_EC"
	server_hardware_type = "SY 480 Gen9 2"
	initial_scope_uris = ["/rest/scopes/74877630-9a22-4061-9db4-d12b6c4cfee0"]
}

/* 	Update
resource "oneview_server_profile_template" "ServerProfileTemplate" {
	name = "TestServerProfileTemplate_Renamed"
	type = "ServerProfileTemplateV5"
	enclosure_group = "SYN03_EC"
	server_hardware_type = "SY 480 Gen9 3"
	initial_scope_uris = ["/rest/scopes/74877630-9a22-4061-9db4-d12b6c4cfee0"]
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