provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

data "oneview_scope" "scope" {
  name = "<scope>"
}

# Updates initial scope uris and name of the resource created from main.tf
resource "oneview_ethernet_network" "ethernetnetwork" {
  name                  = "<ethernet_rename>"
  ethernet_network_type = "<ethernet_network_type>"
  initial_scope_uris    = [data.oneview_scope.scope.uri]
  type                  = "<type>"
  vlan_id               = 200
}
/*
resource "oneview_ethernet_network" "ethernetnetwork_1" {
  name    = "Auto-Ethernet-1"
  ethernet_network_type = "Tagged"
  type    = "ethernet-networkV4"
  vlan_id = 101
}

resource "oneview_ethernet_network" "ethernetnetwork_2" {
  name    = "Auto-Ethernet-2"
  ethernet_network_type = "Tagged"
  type    = "ethernet-networkV4"
  vlan_id = 102
}
*/
