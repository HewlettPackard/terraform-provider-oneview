provider "oneview" {
  ov_username =   "${var.username}"
  ov_password =   "${var.password}"
  ov_endpoint =   "${var.endpoint}"
  ov_sslverify =  "${var.ssl_enabled}"
  ov_apiversion = 2800
  ov_ifmatch = "*"
}

data "oneview_scope" "scope" {
        name = "Auto-Scope"
}

# Updates server profile template from main.tf
resource "oneview_server_profile_template" "ServerProfileTemplate" {
        name = "TestServerProfileTemplateRenamed"
        type = "ServerProfileTemplateV8"
        enclosure_group = "Auto-EG"
        server_hardware_type = "SY 480 Gen9 2"
        initial_scope_uris = ["${data.oneview_scope.scope.uri}"]
}
