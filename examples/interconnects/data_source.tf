provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2600
  ov_ifmatch    = "*"
}

data "oneview_interconnect" "interconnect" {
  name = "0000A66102, interconnect 2"
}

output "oneiew_interconnect_value" {
  value = data.oneview_interconnect.interconnect.uri
}

