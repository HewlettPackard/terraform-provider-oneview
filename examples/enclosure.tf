provider "oneview" {
	ov_username = "<ov_username>"
	ov_password = "<ov_password>"
	ov_endpoint = "<ov_endpoint>"
	ov_sslverify = false
	ov_apiversion = <ov_apiversion>
	ov_ifmatch = "*"
}

data "oneview_enclosure_group" "enclosure_group" {
        name = "TestEnclosureGroup_Renamed"
}

data "oneview_scope" "scope" {
        name = "test"
}

resource "oneview_enclosure" "enclosure_inst" {
	enclosure_group_uri = "${data.oneview_enclosure_group.enclosure_group.uri}"
	host_name = "<enclosure_hostname>"
	user_name = "<enclosure_username>"
	password = "<enclosure_password>"
	licensing_intent = "Oneview"
	initial_scope_uris = ["scope_name1", "scope_name2"]
	name = "Encl2"
}

#Importing Existing resource to update
#run the below command to import the resource:
#terraform import oneview_enclosure.enclosure_inst <your_enclosure_name>

resource "oneview_enclosure" "import_enc"{
}

# Update Enclosure
resource "oneview_enclosure" "import_enc" {
	op = "replace"
	path = "/name"
	value = "Enclosure_Renamed"
	name = "Enclosure_Renamed"
}

# Testing data source
data "oneview_enclosure" "enclosure" {
        name = "SYN03_Frame1"
}

output "oneview_enclosure_value" {
        value = "${data.oneview_enclosure.enclosure.uuid}"
}
