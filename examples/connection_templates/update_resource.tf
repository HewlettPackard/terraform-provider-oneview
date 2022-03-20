provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

resource "oneview_connection_templates" "connectionTemplates" {
  name = "renamed-connection-template"
  bandwidth {
      maximum_bandwidth = 8000
      typical_bandwidth = 2500
  }
}