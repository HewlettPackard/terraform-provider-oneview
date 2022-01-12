provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 3600
  ov_ifmatch    = "*"
}

data "oneview_scope" "scope_obj" {
  name = "Auto-Scope"
}

# Creates FCOE Network Resource with required bandwidth
resource "oneview_fcoe_network" "FCoENetwork" {
  name               = "TestFCoENetwork_Terraform"
  type               = "fcoe-networkV4"
  vlanid             = 202
  initial_scope_uris = [data.oneview_scope.scope_obj.uri]
  bandwidth {
    maximum_bandwidth = 10000
    typical_bandwidth = 1500
  }
}
