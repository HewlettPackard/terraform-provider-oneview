provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

data "oneview_scope" "scope_obj" {
  name = "Auto-Scope"
}

# Updates the created resource with local name FCNetwork and scopes
resource "oneview_fc_network" "FCNetwork" {
  name               = "TestFCNetwork_Renamed"
  fabric_type        = "FabricAttach"
  type               = "fc-networkV4"
  initial_scope_uris = [data.oneview_scope.scope_obj.uri]
}

