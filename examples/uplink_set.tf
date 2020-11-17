provider "oneview" {
  ov_username =   "${var.username}"
  ov_password =   "${var.password}"
  ov_endpoint =   "${var.endpoint}"
  ov_sslverify =  "${var.ssl_enabled}"
  ov_apiversion = 2200
  ov_ifmatch = "*"
}

# Fetching Logical Interconnect
data "oneview_logical_interconnect" "logical_interconnect" {
        name = "6c9d7d01-c176-43c8-b043-6fc0a65f4f9b"
}

# Fetching Network
data "oneview_ethernet_network" "ethernetnetwork" {
  name = "TestNetwork_1"
}

# Creating Upllink Set
resource "oneview_uplink_set" "UplinkSet" {
  name                     = "TestUplinkSet0100"
  type                     = "uplink-setV7"
  logical_interconnect_uri = "${data.oneview_logical_interconnect.logical_interconnect.uri}"
  network_uris             = ["${data.oneview_ethernet_network.ethernetnetwork.uri}",]
  fc_network_uris          = []
  fcoe_network_uris        = []
  port_config_infos 	   = []
  manual_login_redistribution_state = "NotSupported"
  connection_mode                   = "Auto"
  network_type                      = "Ethernet"
  ethernet_network_type             = "Tagged"
}

/*
# Example for data source
data "oneview_uplink_set" "uplink_set" {
        name = "up1"
}

output "oneview_uplink_set_value" {
        value = "${data.oneview_uplink_set.uplink_set.uri}"
}


#Importing Existing resource
resource "oneview_uplink_set" "import_us"{
}
*/
