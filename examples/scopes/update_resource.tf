provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

resource "oneview_scope" "scope" {
  name        = "Auto-Scope"
  description = "Testing creation of scope"
  type        = "ScopeV3"
}

# Updates the resource created from main.tf
resource "oneview_scope" "scope_inst" {
  type               = "ScopeV3"
  name               = "TestScope_Renamed"
  description        = "Rename the Scope"
  initial_scope_uris = [oneview_scope.scope.uri]
  depends_on         = [oneview_scope.scope]
}

