provider "oneview" {
  ov_username =   "${var.username}"
  ov_password =   "${var.password}"
  ov_endpoint =   "${var.endpoint}"
  ov_sslverify =  "${var.ssl_enabled}"
  ov_apiversion = 2200
  ov_ifmatch = "*"
}

variable "LE_name" {
 type = "string"
 description = "Logical Enclosure Name"
}

variable "EG_name" {
 type = "string"
 description = "Enclosure Group Name"
}

# Fetching Enclosure Group
data "oneview_enclosure_group" "enclosure_group" {
        name = "${var.EG_name}"
}

/*
# Creating Logical Enclosure
resource "oneview_logical_enclosure" "LogicalEnclosure" {
	name = "${var.LE_name}"
	enclosure_uris = ["/rest/enclosures/0000000000A66101","/rest/enclosures/0000000000A66102","/rest/enclosures/0000000000A66103"]
	enclosure_group_uri = "${data.oneview_enclosure_group.enclosure_group.uri}"
}
*/

/*
# Import an Exisiting Logical Enclosure
# terraform import oneview_logical_enclosure.LogicalEnclosure <LE_Name>
resource "oneview_logical_enclosure" "LogicalEnclosure" {
}
*/

/*
# Update by Group 
resource "oneview_logical_enclosure" "LogicalEnclosure" {
  name = "${var.LE_name}"
  enclosure_uris = ["/rest/enclosures/0000000000A66101","/rest/enclosures/0000000000A66102","/rest/enclosures/0000000000A66103"]
  enclosure_group_uri = "${data.oneview_enclosure_group.enclosure_group.uri}"
  update_type = "updateByGroup"
}
*/

/*
# Datasource
data "oneview_logical_enclosure" "logical_enclosure" {
	name = "${var.LE_name}"
	enclosure_group_uri = "${data.oneview_enclosure_group.enclosure_group.uri}"
}

output "oneview_logical_enclosure_value" {
        value = "${data.oneview_logical_enclosure.logical_enclosure.name}"
}
*/
