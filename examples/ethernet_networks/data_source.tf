provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2800
  ov_ifmatch    = "*"
}

# Testing data source
data "oneview_ethernet_network" "ethernetnetworks" {
get_type = "GetAll"
}

output "oneview_ethernet_network_value" {
  value = data.oneview_ethernet_network.ethernetnetworks
}

data "oneview_ethernet_network" "ethernetnetworkbyname" {
get_type = "GetById"
name = "Test"
}

output "oneview_ethernet_network_by_name_value" {
  value = data.oneview_ethernet_network.ethernetnetworkbyname
}

resource "oneview_ethernet_network" "ethernetnetwork_1" {
  name    = "Auto-Ethernet-1"
  type    = "ethernet-networkV4"
  vlan_id = 101
}

resource "oneview_ethernet_network" "ethernetnetwork_2" {
  name    = "Auto-Ethernet-2"
  type    = "ethernet-networkV4"
  vlan_id = 102
}

#Importing Existing resource
#resource "oneview_ethernet_network" "import_eth"{
#}
