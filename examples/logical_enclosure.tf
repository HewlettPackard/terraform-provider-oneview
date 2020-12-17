provider "oneview" {
  ov_username =   var.username
  ov_password =   var.password
  ov_endpoint =   var.endpoint
  ov_sslverify =  var.ssl_enabled
  ov_apiversion = 2200
  ov_ifmatch = "*"
}

data "oneview_enclosure_group" "enclosure_group" {
  name = "Auto-TestEG"
}

resource "oneview_logical_enclosure" "LogicalEnclosure" {
  name                = "TestLE"
  enclosure_uris      = ["/rest/enclosures/0000000000A66101", "/rest/enclosures/0000000000A66102", "/rest/enclosures/0000000000A66103"]
  enclosure_group_uri = data.oneview_enclosure_group.enclosure_group.uri
}

# Import an Exisiting Logical Enclosure
# terraform import oneview_logical_enclosure.LogicalEnclosure <LE_Name>
resource "oneview_logical_enclosure" "LogicalEnclosure" {
}

# Update by Group 
resource "oneview_logical_enclosure" "LogicalEnclosure" {
  name                = "TestLE"
  enclosure_uris      = ["/rest/enclosures/0000000000A66101", "/rest/enclosures/0000000000A66102", "/rest/enclosures/0000000000A66103"]
  enclosure_group_uri = data.oneview_enclosure_group.enclosure_group.uri
  update_type         = "updateByGroup"
}

# Datasource
data "oneview_logical_enclosure" "logical_enclosure" {
  name                = "TESTLE"
  enclosure_group_uri = data.oneview_enclosure_group.enclosure_group.uri
}

output "oneview_logical_enclosure_value" {
  value = data.oneview_logical_enclosure.logical_enclosure.name
}
