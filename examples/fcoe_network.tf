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
data "oneview_scope" "scope_obj1" {
        name = "test1"
}

# resource "oneview_fcoe_network" "FCoENetwork" {
# 	name = "TestFCoENetwork_Terraform"
# 	type = "fcoe-networkV4"
#   vlanId = 202
#   initial_scope_uris = ["${data.oneview_scope.scope_obj.uri}", "${data.oneview_scope.scope_obj1.uri}"]
# }

# Updates the resource created above
# To update uncomment the below and add the attributes to be updated
/*resource "oneview_fcoe_network" "FCoENetwork" {
	name = "TestFCoENetwork_Terraform_Renamed"
	type = "fcoe-networkV4"
   vlanId = 202
}
*/
/* Testing data source
data "oneview_fcoe_network" "fcoe_network_obj" {
        name = "TestFCoENetwork_Terraform_Renamed"
}

output "oneview_fcoe_network" {
        value = "${data.oneview_fcoe_network.fcoe_network_obj.vlan_id}"
}
*/


#Importing Existing resource
/*resource "oneview_fcoe_network" "import_fcoe"{
}
*/