provider "oneview" {
        ov_username = "<ov-username>"
        ov_password = "<ov-password>"
        ov_endpoint = "<ov-endpoint>"
        ov_sslverify = false
        ov_apiversion = <ov-apiversion>
        ov_ifmatch = "*"
}

data "oneview_scope" "scope" {
	name = "test"
}

data "oneview_logical_interconnect_group" "logical_interconnect_group" {
        name = "test"
}

resource "oneview_enclosure_group" "eg_inst" {
	name = "TestEnclosureGroup"
	description = "Testing creation of Enclosure Group"
	ip_addressing_mode = "External"
	enclosure_count = 3
	initial_scope_uris = ["scope_1", "scope_2"]
	interconnect_bay_mappings = [
	{
		interconnect_bay = 3
		logical_interconnect_group_uri = "${data.oneview_logical_interconnect_group.logical_interconnect_group.uri}"
	},
	{
		interconnect_bay = 6
		logical_interconnect_group_uri = "${data.oneview_logical_interconnect_group.logical_interconnect_group.uri}"
	}
	]
}

# Updates the resource created above
# To update uncomment the below lines and add the values to the attributes mentioned
/* resource "oneview_enclosure_group" "eg_inst" {
	name = "TestEnclosureGroup_Renamed"
	enclosure_count = 1
	ip_addressing_mode = "External"
	power_mode = "RedundantPowerFeed"
	interconnect_bay_mappings = [
        {
                interconnect_bay = 3
                logical_interconnect_group_uri =  "${data.oneview_logical_interconnect_group.logical_interconnect_group.uri}"
        },
        {
                interconnect_bay = 6
                logical_interconnect_group_uri =  "${data.oneview_logical_interconnect_group.logical_interconnect_group.uri}"
        }
        ]
	interconnect_bay_mapping_count = 2
	type = "EnclosureGroupV8"
	stacking_mode = "Enclosure"
}

# Test for data source 

data "oneview_enclosure_group" "enclosure_group" {
        name = "TestEnclosureGroup_Renamed"
}

output "oneview_enclosure_group_value" {
        value = "${data.oneview_enclosure_group.enclosure_group.uri}"
}
*/
