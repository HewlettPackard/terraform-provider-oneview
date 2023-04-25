provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Extracting Storage System
data "oneview_storage_system" "ss_inst" {
  name = "<storage_system_name2>"
}

output "oneview_storage_system_value" {
  value = data.oneview_storage_system.ss_inst.uri
}

# Testing import of existing resource
/*
resource "oneview_storage_system" "ss_import"{
}
*/
