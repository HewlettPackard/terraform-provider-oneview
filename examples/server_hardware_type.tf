provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2400
  ov_ifmatch    = "*"
}

data "oneview_server_hardware_type" "server_hardware_type" {
  name = "SY 480 Gen9 2"
}

output "oneiew_server_hardware_type_value" {
  value = data.oneview_server_hardware_type.server_hardware_type.uri
}

