provider "oneview" {
	ov_username = "administrator"
	ov_password = "madhav123"
	ov_endpoint = "https://10.170.16.44"
	ov_sslverify = false
	ov_apiversion = 800
	ov_ifmatch = "*"
}

resource "oneview_fc_network" "FCNetwork" {
	name = "TestFCNetwork"
	type = "fc-networkV4"
	initial_scope_uris = ["/rest/scopes/8aa405fb-bc62-42e5-9ca6-4a544e7ffdec", "/rest/scopes/bf7df2e6-ebbb-4c6b-bc7a-be8cdac930bf"]
}

/*resource "oneview_fc_network" "FCNetwork" {
	name = "TestFCNetwork_Renamed"
	type = "fc-networkV4"
}*/
