provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2400
  ov_ifmatch    = "*"
}

data "oneview_interconnect" "interconnect" {
  name = "0000A66101, interconnect 3"
}

output "oneiew_interconnect_value" {
  value = data.oneview_interconnect.interconnect.uri
}

