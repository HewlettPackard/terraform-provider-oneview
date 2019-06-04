provider "oneview" {
  ov_username   = <ov_username>
  ov_password   = <ov_password>
  ov_endpoint   = <ov_ip>
  ov_sslverify  = false
  ov_apiversion = <ov_apiversion>
}

resource "oneview_uplink_set" "UplinkSet" {
  name                     = "TestUplinkSet0100"
  type                     = "uplink-setV4"
  logical_interconnect_uri = "/rest/logical-interconnects/34abde89-d9a8-4f72-aa16-2c19bb548b11"
  network_uris             = ["/rest/ethernet-networks/e022d0a9-4fc5-415c-9aeb-5e7047ae657a"]
  fc_network_uris          = []
  fcoe_network_uris        = []
  port_config_infos 	   = []
  manual_login_redistribution_state = "NotSupported"
  connection_mode                   = "Auto"
  network_type                      = "Ethernet"
  ethernet_network_type             = "Tagged"
}
