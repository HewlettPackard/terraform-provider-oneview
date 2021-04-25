provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2800
  ov_ifmatch    = "*"
}

# Updates the resource created above
# To update uncomment the below and add the attributes to be updated
resource "oneview_fcoe_network" "FCoENetwork" {
  name   = "TestFCoENetwork_Terraform_Renamed"
  type   = "fcoe-networkV4"
  vlanid = 202
}

