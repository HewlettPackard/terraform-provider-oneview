provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

data "oneview_ethernet_network" "ethernet_network" {
  name = "<network_name>"
}

data "oneview_scope" "initial_scope_uri" {
  name = "<initial_scope>"
}

resource "oneview_scope" "scope_inst" {
  name                = "<scope_name>"
  description         = "Testing creation of scope"
  type                = "<type>"
  initial_scope_uris  = [data.oneview_scope.initial_scope_uri.uri]
  added_resource_uris = [data.oneview_ethernet_network.ethernet_network.uri]
}

