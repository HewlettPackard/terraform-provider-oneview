provider "oneview" {
	ov_username = "<ov_username>"
	ov_password = "<ov_password>"
	ov_endpoint = "<ov_endpoint>"
	ov_sslverify = false
	ov_apiversion = <ov_apiversion>
	ov_ifmatch = "*"
}

data "oneview_enclosure_group" "enclosure_group" {
        name = "EnclosureGroupDemo"
}

resource "oneview_logical_enclosure" "LogicalEnclosure" {
	name = "TESTLE"
	enclosure_uris = ["/rest/enclosures/0000000000A66101","/rest/enclosures/0000000000A66102","/rest/enclosures/0000000000A66103"]
	enclosure_group_uri = "${data.oneview_enclosure_group.enclosure_group.uri}"
}

/* Update by Group

resource "oneview_logical_enclosure" "LogicalEnclosure" {
  name = "TESTLE"
  enclosure_uris = ["/rest/enclosures/0000000000A66101","/rest/enclosures/0000000000A66102","/rest/enclosures/0000000000A66103"]
  enclosure_group_uri = "${data.oneview_enclosure_group.enclosure_group.uri}"
  update_type = "updateByGroup"
}
*/

# Datasource

/*
data "oneview_logical_enclosure" "logical_enclosure" {
    name = "TESTLE"
	enclosure_group_uri = "${data.oneview_enclosure_group.enclosure_group.uri}"
}

output "oneview_logical_enclosure_value" {
        value = "${data.oneview_logical_enclosure.logical_enclosure.name}"
}
*/
