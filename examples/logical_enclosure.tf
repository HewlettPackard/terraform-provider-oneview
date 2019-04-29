provider "oneview" {
	ov_username = "Administrator"
	ov_password = "madhav123"
	ov_endpoint = "https://10.170.16.44"
	ov_sslverify = false
	ov_apiversion = 600
	ov_ifmatch = "*"
}
/*data "oneview_logical_enclosure" "logical_enclosure" {
        name = "TESTLE"
	enclosure_group_uri = "/rest/enclosure-groups/4f7b2cf4-fdb5-4065-8886-54a9905659f9"
}

output "oneview_logical_enclosure_value" {
        value = "${data.oneview_logical_enclosure.logical_enclosure.name}"
}*/

resource "oneview_logical_enclosure" "LogicalEnclosure" {
	name = "TESTLE"
	enclosure_uris = ["/rest/enclosures/0000000000A66101","/rest/enclosures/0000000000A66102","/rest/enclosures/0000000000A66103"]
	enclosure_group_uri = "/rest/enclosure-groups/4f7b2cf4-fdb5-4065-8886-54a9905659f9"

}

