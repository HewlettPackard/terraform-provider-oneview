provider "oneview" {
	ov_username = "administrator"
	ov_password = "madhav123"
	ov_endpoint = "https://10.170.16.44"
	ov_sslverify = false
	ov_apiversion = 800
	ov_ifmatch = "*"
}

resource "oneview_uplink_set" "UplinkSet" {
	name = "TestUplinkSet"
	type = "uplink-setV4"
	logical_interconnect_uri = "/rest/logical-interconnect-groups/0d3f98d6-0c10-4565-b9db-3907d5534e1d"
	network_uris = []
	fc_network_uris = []
	fcoe_network_uris = []
	port_config_infos = {}
	manual_login_redistribution_state = "NotSupported"
	connection_mode = "Auto"
	network_type = "Ethernet"
	ethernet_network_type = "Tagged"
}
