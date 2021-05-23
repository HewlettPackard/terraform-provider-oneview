provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 3000
  ov_ifmatch    = "*"
}

resource "oneview_connection_templates" "connectionTemplates" {
  bandwidth {
      maximum_bandwidth = 8000
      typical_bandwidth = 2500
  }
}