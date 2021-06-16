provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

data "oneview_ipv4_range" "ipranges" {
  id = "b1b869f8-3d5a-4d4a-b0a2-fb6634f045d6"
}

output "ipranges_value" {
  value = data.oneview_ipv4_range.ipranges
}