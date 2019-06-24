provider "oneview" {
	ov_username = "<ov_username>"
	ov_password = "<ov_password>"
	ov_endpoint = "<ov_endpoint>"
	ov_sslverify = false
	ov_apiversion = <ov_apiversion>
	ov_ifmatch = "*"
}

resource "oneview_enclosure" "enclosure_inst" {
	enclosure_group_uri = "/rest/enclosure-groups/22b775e2-57dd-46d1-beb1-b53b14d51ce4"
	host_name = "<enclosure_hostname>"
	user_name = "<enclosure_username>"
	password = "<enclosure_password>"
	licensing_intent = "Oneview"
	initial_scope_uris = ["/rest/scopes/bf7df2e6-ebbb-4c6b-bc7a-be8cdac930bf"]
}

/*resource "oneview_enclosure" "enclosure_inst" {
	op = "replace"
	path = "/name"
	value = "Enclosure_Renamed"
}*/

/* Testing data source
data "oneview_enclosure" "enclosure" {
        name = "SYN03_Frame1"
}

output "oneview_enclosure_value" {
        value = "${data.oneview_enclosure.enclosure.uuid}"
}

*/
