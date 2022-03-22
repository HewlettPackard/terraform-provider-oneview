provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

data "oneview_scope" "scope" {
  name = "Auto-Scope"
}

# Creates Ethernet Network Resource with required bandwidth
resource "oneview_ethernet_network" "ethernetnetwork" {
  name                  = "TestEthNetwork_terraform"
  ethernet_network_type = "Tagged"
  type                  = "ethernet-networkV4"
  vlan_id               = 100
  initial_scope_uris    = [data.oneview_scope.scope.uri]
  bandwidth {
    maximum_bandwidth = 10000
    typical_bandwidth = 1500
  }
}
/*
# Creates 2 Ethernet Network Resource
resource "oneview_ethernet_network" "ethernetnetwork_1" {
  name                  = "Auto-Ethernet-1"
  ethernet_network_type = "Tagged"
  type                  = "ethernet-networkV4"
  vlan_id               = 101
}

resource "oneview_ethernet_network" "ethernetnetwork_2" {
  name                  = "Auto-Ethernet-2"
  ethernet_network_type = "Tagged"
  type                  = "ethernet-networkV4"
  vlan_id               = 102
}

# Creating tunnel ethernet network.vland id is not needed.
# resource "oneview_ethernet_network" "ethernetnetwork_2" {
#   name                  = "Auto-Ethernet-2"
#   ethernet_network_type = "Tunnel"
#   type                  = "ethernet-networkV4"

# }
#Creating untagged ethernet network . vland id is not needed.
# resource "oneview_ethernet_network" "ethernetnetwork_2" {
#   name                  = "Auto-Ethernet-2"
#   ethernet_network_type = "Untagged"
#   type                  = "ethernet-networkV4"

# }
*/
