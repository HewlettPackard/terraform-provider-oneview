provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

data "oneview_scope" "scope_obj" {
  name = "<scope>"
}

# Creates FCOE Network Resource with required bandwidth
resource "oneview_fcoe_network" "FCoENetwork" {
  name               = "<network_name_terra>"
  type               = "<type>"   
  vlanid             = 202
  initial_scope_uris = [data.oneview_scope.scope_obj.uri]
  bandwidth {
    maximum_bandwidth = 10000
    typical_bandwidth = 1500
  }
}
