provider "oneview" {
	ov_username = "<ov_username>"
	ov_password = "<ov_password"
	ov_endpoint = "<ov_endpoint>"
	ov_sslverify = false
	ov_apiversion = <ov_apiversion>
	ov_ifmatch = "*"
}

data "oneview_scope" "scope_obj" {
        name = "test"
}

resource "oneview_fc_network" "FCNetwork" {
	name = "TestFCNetwork"
	type = "fc-networkV4"
	initial_scope_uris = ["${data.oneview_scope.scope_obj.uri}"]
}

/*resource "oneview_fc_network" "FCNetwork" {
	name = "TestFCNetwork_Renamed"
	type = "fc-networkV4"
}*/

/* Testing data source
data "oneview_fc_network" "fc_network" {
        name = "SYN03_Frame1"
}

output "oneview_fc_network_value" {
        value = "${data.oneview_fc_network.fc_network.fabric_type}"
}
*/

/*
#Importing Existing resource
resource "oneview_fc_network" "import_fc"{
}
*/
