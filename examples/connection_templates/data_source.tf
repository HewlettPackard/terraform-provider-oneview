provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 3000
  ov_ifmatch    = "*"
}

# Testing data source
data "oneview_connection_templates" "connectionTemplates" {
  name = "defaultConnectionTemplate"
}

output "oneview_connection_template_value" {
  value = data.oneview_connection_templates.connectionTemplates
}