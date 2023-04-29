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
  name = "<ethernet_rename>"
}

resource "oneview_ethernet_network" "ethernetnetwork_1" {
  name    = "<data_source_terraform1>"
  ethernet_network_type = "<ethernet_network_type>"
  type    = "<type>"
  vlan_id = 200
}

resource "oneview_ethernet_network" "ethernetnetwork_2" {
  name    = "<data_source_terraform2>"
  ethernet_network_type = "<ethernet_network_type>"
  type    = "<type>"
  vlan_id = 2001
}

#Importing Existing resource
#resource "oneview_ethernet_network" "import_eth"{
#}
