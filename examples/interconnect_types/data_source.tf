provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

data "oneview_interconnect_type" "interconnect_type" {
  name = "Virtual Connect SE 40Gb F8 Module for Synergy"
}

output "oneview_interconnect_type_value" {
  value = data.oneview_interconnect_type.interconnect_type.type
}

