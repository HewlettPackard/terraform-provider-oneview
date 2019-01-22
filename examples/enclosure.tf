provider "oneview" {
	ov_username = "administrator"
	ov_password = "madhav123"
	ov_endpoint = "https://10.170.16.44"
	ov_sslverify = false
	ov_apiversion = 800
	ov_ifmatch = "*"
}

resource "oneview_enclosure" "enclosure_inst" {
	enclosure_group_uri = "/rest/enclosure-groups/22b775e2-57dd-46d1-beb1-b53b14d51ce4"
	host_name = "172.18.1.13"
	user_name = "dcs"
	password = "dcs"
	licensing_intent = "Oneview"
	initial_scope_uris = ["/rest/scopes/bf7df2e6-ebbb-4c6b-bc7a-be8cdac930bf"]
}
/*resource "oneview_enclosure" "enclosure_inst" {
	op = "replace"
	path = "/name"
	value = "Enclosure_Renamed"
}*/
