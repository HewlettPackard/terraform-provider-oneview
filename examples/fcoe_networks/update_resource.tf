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

# Updates the resource created above
# To update uncomment the below and add the attributes to be updated
resource "oneview_fcoe_network" "FCoENetwork" {
  name   = "<network_rename>"
  type   = "<type>"
  vlanid = 202
  initial_scope_uris = [data.oneview_scope.scope_obj.uri]
}

