provider "oneview" {
        ov_username =   "${var.username}"
        ov_password =   "${var.password}"
        ov_endpoint =   "${var.endpoint}"
        ov_sslverify =  "${var.ssl_enabled}"
        ov_apiversion = 2400
        ov_ifmatch = "*"

}

data "oneview_scope" "scope_obj" {
  name = "testing"
}

data "oneview_scope" "scope_obj1" {
  name = "testing_1"
}

resource "oneview_fcoe_network" "FCoENetwork" {
  name               = "TestFCoENetwork_Terraform"
  type               = "fcoe-networkV4"
  vlanId             = 202
  initial_scope_uris = ["${data.oneview_scope.scope_obj.uri}", "${data.oneview_scope.scope_obj1.uri}"]
}

# Updates the resource created above
# To update uncomment the below and add the attributes to be updated
/*
resource "oneview_fcoe_network" "FCoENetwork" {
  name   = "TestFCoENetwork_Terraform_Renamed"
  type   = "fcoe-networkV4"
  vlanId = 202
}

# Testing data source
data "oneview_fcoe_network" "fcoe_network_obj" {
        name = "TestFCoENetwork_Terraform_Renamed"
}

output "oneview_fcoe_network" {
        value = "${data.oneview_fcoe_network.fcoe_network_obj.vlan_id}"
}


# Importing Existing resource
resource "oneview_fcoe_network" "import_fcoe"{
}
*/

