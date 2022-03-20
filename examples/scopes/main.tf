provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

data "oneview_ethernet_network" "ethernet_network" {
  name = "Auto-Ethernet-2"
}

data "oneview_scope" "initial_scope_uri" {
  name = "Auto-Scope"
}

resource "oneview_scope" "scope_inst" {
  name                = "TestScope"
  description         = "Testing creation of scope"
  type                = "ScopeV3"
  initial_scope_uris  = [data.oneview_scope.initial_scope_uri.uri]
  added_resource_uris = [data.oneview_ethernet_network.ethernet_network.uri]
}

