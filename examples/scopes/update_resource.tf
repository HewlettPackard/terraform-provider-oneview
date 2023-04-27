provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

resource "oneview_scope" "scope" {
  name        = "test-scope"
  description = "Testing creation of scope"
  type        = "<type>"
}

# Updates the resource created from main.tf
resource "oneview_scope" "scope_inst" {
  type               = "<type>"
  name               = "<scope_rename>"
  description        = "Rename the Scope"
  initial_scope_uris = [oneview_scope.scope.uri]
  depends_on         = [oneview_scope.scope]
}

