provider "oneview" {
	ov_username = "<ov_username>"
	ov_password = "<ov_password>"
	ov_endpoint = "<ov_endpoint>"
	ov_sslverify = false
	ov_apiversion = <ov_apiversion>
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

/* Update by Group

resource "oneview_logical_enclosure" "LogicalEnclosure" {
	name = "TESTLE"
	enclosure_uris = ["/rest/enclosures/0000000000A66101","/rest/enclosures/0000000000A66102","/rest/enclosures/0000000000A66103"]
	enclosure_group_uri = "/rest/enclosure-groups/4f7b2cf4-fdb5-4065-8886-54a9905659f9"
	update_type = "updateByGroup"
*/
