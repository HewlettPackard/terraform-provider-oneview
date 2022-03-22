provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}


data "oneview_scope" "scope_obj" {
  name = "Auto-Scope"
}

// Adds Rack server to the appliance
// Licensing can be OneView or OneviewNoiLO for Managed
resource "oneview_server_hardware" "sh" {
  configuration_state = "Managed"
  hostname            = "<serverIp>"
  username            = "dcs"
  password            = "dcs"
  licensing_intent    = "OneView"
  initial_scope_uris  = [data.oneview_scope.scope_obj.uri]
}

// To import server hardware you can use server hardware name. 
// For below example run terraform import oneview_server_hardware.sh <name-of-the-server>
// resource "oneview_server_hardware" "sh" {
// }
