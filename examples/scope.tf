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

resource "oneview_scope" "scope_inst" {
	name = "TestScope"
	description = "Testing creation of scope"
	type = "ScopeV3"
	initial_scope_uris = ["${data.oneview_scope.scope.uri}"]
	added_resource_uris = ["/rest/fc-networks/ac24da0d-d993-4ac5-b16c-c42c2842a918"]
}

# Updates the resource created above 
# To update uncomment the below and ad the attributes  to be updated
/*resource "oneview_scope" "scope_inst" {
	name = "TestScope_Renamed"
	type = "ScopeV3"
	description = "Rename the Scope"
}*/
