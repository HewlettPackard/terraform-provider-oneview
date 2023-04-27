provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

data "oneview_ethernet_network" "eth1" {
  name = "<ethernet1_name>"
}

data "oneview_ethernet_network" "eth2" {
  name = "<ethernet2_name>"
}

data "oneview_scope" "scope_obj" {
  name = "<scope>"
}

resource "oneview_network_set" "NetworkSet" {
  name               = "<network_name>"
  native_network_uri = ""
  type               = "<type>"
  network_uris       = [data.oneview_ethernet_network.eth1.uri, data.oneview_ethernet_network.eth2.uri]
  initial_scope_uris = [data.oneview_scope.scope_obj.uri]
}

