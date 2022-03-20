provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

data "oneview_server_hardware" "server_hardware" {
  name = "0000A66102, bay 4"
}

output "oneview_server_hardware_value" {
  value = data.oneview_server_hardware.server_hardware.uri
}

