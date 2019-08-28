provider "oneview" {
	ov_username = "<ov-username>"
	ov_password = "<ov-password>"
	ov_endpoint = "<ov-endpoint>"
	ov_sslverify = false
	ov_apiversion = <ov-apiversion>
	ov_ifmatch = "*"
}

data "oneview_scope" "scope" {
	name = "Scope_Sample"
}

resource "oneview_enclosure_group" "eg_inst" {
	name = "TestEnclosureGroup"
	description = "Testing creation of Enclosure Group"
	ip_addressing_mode = "External"
	enclosure_count = 3
	initial_scope_uris = ["${data.oneview_scope.scope.uri}"]
	interconnect_bay_mappings = [
	{
		interconnect_bay = 1
		logical_interconnect_group_uri = "/rest/sas-logical-interconnect-groups/a14f144f-359d-4ce2-b2c1-446382a30ad0"
	},
	{
		interconnect_bay = 4
		logical_interconnect_group_uri = "/rest/sas-logical-interconnect-groups/a14f144f-359d-4ce2-b2c1-446382a30ad0"
	}
	]
}

# Updates the resource created above
# To update uncomment the below lines and add the values to the attributes mentioned
/*resource "oneview_enclosure_group" "eg_inst" {
	name = "TestEnclosureGroup_Renamed"
	enclosure_count = 3 
	ip_addressing_mode = "External"
	power_mode = "RedundantPowerFeed"
	interconnect_bay_mappings = [
        {
                interconnect_bay = 1
                logical_interconnect_group_uri = "/rest/sas-logical-interconnect-groups/a14f144f-359d-4ce2-b2c1-446382a30ad0"
        },
        {
                interconnect_bay = 4
                logical_interconnect_group_uri = "/rest/sas-logical-interconnect-groups/a14f144f-359d-4ce2-b2c1-446382a30ad0"
        }
        ]
	interconnect_bay_mapping_count = 2
	type = "EnclosureGroupV7"
	stacking_mode = "Enclosure"
}*/

/* Test for data source

data "oneview_enclosure_group" "enclosure_group" {
        name = "EnclosureGroupDemo"
}

output "oneview_enclosure_group_value" {
        value = "${data.oneview_enclosure_group.enclosure_group.uri}"
}*/
