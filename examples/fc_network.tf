provider "oneview" {
	ov_username = "<ov_username>"
	ov_password = "<ov_password"
	ov_endpoint = "<ov_endpoint>"
	ov_sslverify = false
	ov_apiversion = <ov_apiversion>
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
