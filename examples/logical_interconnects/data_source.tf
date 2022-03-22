provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Get Logical Interconnect to terraform
data "oneview_logical_interconnect" "logical_interconnect" {
  name = "Auto-LE-Auto-LIG"
}

output "oneview_logical_interconnect_value" {
  value = data.oneview_logical_interconnect.logical_interconnect.uri
}

