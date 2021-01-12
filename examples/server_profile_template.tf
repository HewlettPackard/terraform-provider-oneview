provider "oneview" {
  ov_username =   var.username
  ov_password =   var.password
  ov_endpoint =   var.endpoint
  ov_sslverify =  var.ssl_enabled
  ov_apiversion = 2400
  ov_ifmatch = "*"
}


data "oneview_scope" "scope" {
  name = "test"
}

#	Create 
resource "oneview_server_profile_template" "ServerProfileTemplate" {
  name                 = "TestServerProfileTemplate"
  type                 = "ServerProfileTemplateV8"
  enclosure_group      = "test_EG"
  server_hardware_type = "SY 480 Gen9 1"
  initial_scope_uris   = [data.oneview_scope.scope.uri]
}
 

# 	Update 
resource "oneview_server_profile_template" "ServerProfileTemplate" {
  name                 = "TestServerProfileTemplate_Renamed"
  type                 = "ServerProfileTemplateV8"
  enclosure_group      = "test_EG"
  server_hardware_type = "SY 480 Gen9 1"
  initial_scope_uris   = [data.oneview_scope.scope.uri]
}


# 	Datasource 
data "oneview_server_profile_template" "server_profile_template" {
  name = "TestServerProfileTemplate"
  depends_on = [oneview_server_profile_template.ServerProfileTemplate]
}

output "oneiew_server_hardware_type_value" {
  value = data.oneview_server_profile_template.server_profile_template.uri
} 
