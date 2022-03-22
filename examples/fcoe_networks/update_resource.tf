provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

data "oneview_scope" "scope_obj" {
  name = "Auto-Scope-Test"
}

# Updates the resource created above
# To update uncomment the below and add the attributes to be updated
resource "oneview_fcoe_network" "FCoENetwork" {
  name   = "TestFCoENetwork_Terraform_Renamed"
  type   = "fcoe-networkV4"
  vlanid = 202
  initial_scope_uris = [data.oneview_scope.scope_obj.uri]
}

