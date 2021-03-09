provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2600
  ov_ifmatch    = "*"
}

data "oneview_scope" "scope" {
  name = "Auto-Scope"
}

# Creation of Server Profile without template

resource "oneview_server_profile" "SP" {
  name               = "TestSP2"
  hardware_name      = "0000A66102, bay 5"
  type               = "ServerProfileV12"
  enclosure_group    = "Auto-EG"
  initial_scope_uris = [data.oneview_scope.scope.uri]
}

