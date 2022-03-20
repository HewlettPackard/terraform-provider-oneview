provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Testing data source
# Gets connection template data from name
data "oneview_connection_templates" "connectionTemplates" {
  name = "defaultConnectionTemplate"
}

output "oneview_connection_template_value" {
  value = data.oneview_connection_templates.connectionTemplates
}

# Gets connection template data from uri
data "oneview_connection_templates" "connectionTemplatesURI" {
  uri = data.oneview_connection_templates.connectionTemplates.uri
}

output "oneview_connection_template_valueURI" {
  value = data.oneview_connection_templates.connectionTemplatesURI.name
}
