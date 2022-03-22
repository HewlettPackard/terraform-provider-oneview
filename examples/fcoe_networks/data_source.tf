provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Testing data source
data "oneview_fcoe_network" "fcoe_network_obj" {
  name = "TestFCoENetwork_Terraform_Renamed"
}

output "oneview_fcoe_network" {
  value = data.oneview_fcoe_network.fcoe_network_obj.vlan_id
}

# Importing Existing resource
#resource "oneview_fcoe_network" "import_fcoe"{
#}
