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

output "scope_value" {
        value = "${data.oneview_scope.scope.uri}"
}

resource "oneview_scope" "scope_inst" {
	name = "TestScope"
	description = "Create Scope"
	type = "ScopeV3"
	initial_scope_uris = ["/rest/scopes/8aa405fb-bc62-42e5-9ca6-4a544e7ffdec"]
	added_resource_uris = ["/rest/fc-networks/9886a1ac-accb-4089-a33f-349dd449982a"]
}

# Updates the resource created above 
# To update uncomment the below and ad the attributes  to be updated
/*resource "oneview_scope" "scope_inst" {
	name = "TestScope_Renamed"
	type = "ScopeV3"
	description = "Rename the Scope"
}*/
