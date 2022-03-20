provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Testing data source
data "oneview_ethernet_network" "ethernetnetworks" {
  name = "TestEthNetwork_terraform_Rename"
}

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

#Importing Existing resource
#resource "oneview_ethernet_network" "import_eth"{
#}
