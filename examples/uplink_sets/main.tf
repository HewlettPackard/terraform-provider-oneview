provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Fetching Logical Interconnect
data "oneview_logical_interconnect" "logical_interconnect" {
  name = "Auto-LE-Auto-LIG"
}

# Fetching Network
data "oneview_ethernet_network" "ethernetnetwork" {
  name = "Auto-Ethernet-1"
}

# Creating Uplink Set
resource "oneview_uplink_set" "UplinkSet" {
  name                              = "Auto-UplinkSet"
  type                              = "uplink-setV7"
  logical_interconnect_uri          = data.oneview_logical_interconnect.logical_interconnect.uri
  network_uris                      = [data.oneview_ethernet_network.ethernetnetwork.uri]
  fc_network_uris                   = []
  fcoe_network_uris                 = []
  manual_login_redistribution_state = "NotSupported"
  connection_mode                   = "Auto"
  network_type                      = "Ethernet"
  ethernet_network_type             = "Tagged"
}

# Example for data source
data "oneview_uplink_set" "uplink_set" {
  name = "test" //"Auto-UplinkSet"
}

output "oneview_uplink_set_value" {
  value = data.oneview_uplink_set.uplink_set.uri
}

#Importing Existing resource
#resource "oneview_uplink_set" "import_us"{
#}
