provider "oneview" {
  ov_username   = "administrator"
  ov_password   = "madhav123"
  ov_endpoint   = "https://10.170.16.41"
  ov_sslverify  = false
  ov_apiversion = 600
  ov_ifmatch    = "*"
}

resource "oneview_uplink_set" "UplinkSet" {
  name                     = "TestUplinkSet"
  type                     = "uplink-setV4"
  logical_interconnect_uri = "/rest/logical-interconnects/fe57cf01-91c1-4a1d-b057-dff3e881d2b9"
  network_uris             = ["/rest/ethernet-networks/7771c389-5919-4037-8ff6-00aa593dbc4d"]
  fc_network_uris          = [""]
  fcoe_network_uris        = [""]
  port_config_infos 	   = [{
				    "desired_speed" = "Auto"
				    "location" = {
				      "location_entries" = {
								"value"="1"
								"type"="Bay"
								}
    					}
  			     }]
  manual_login_redistribution_state = "NotSupported"
  connection_mode                   = "Auto"
  network_type                      = "Ethernet"
  ethernet_network_type             = "Tagged"
}
