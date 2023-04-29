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

# Updates the created resource with local name FCNetwork and scopes
resource "oneview_fc_network" "FCNetwork" {
  name               =  "<fcn_rename>"
  fabric_type        =  "<fabric_type>"
  type               =  "<type>"
  initial_scope_uris = [data.oneview_scope.scope_obj.uri]
}
