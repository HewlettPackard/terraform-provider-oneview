provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2400
  ov_ifmatch    = "*"
}

data "oneview_ethernet_network" "ethernetnetwork" {
  name = "Prod_1103"
}

resource "oneview_uplink_set" "UplinkSet" {
  name                              = "TestUplinkSet0100"
  type                              = "uplink-setV7"
  logical_interconnect_uri          = "/rest/logical-interconnects/05c45bdf-a2eb-4461-95b9-3e971b9e105e"
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
  name = "TestUplinkSet0100"
  depends_on = [oneview_uplink_set.UplinkSet]
}

output "oneview_uplink_set_value" {
  value = data.oneview_uplink_set.uplink_set.uri
  depends_on = [ oneview_uplink_set.uplink_set]
}

/*
#Importing Existing resource
resource "oneview_uplink_set" "import_us" {
}
*/
