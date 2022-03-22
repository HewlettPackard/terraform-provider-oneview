provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Datasource for Volume
data "oneview_volume" "volume" {
  name = "testvol"
}

output "oneview_volume_value" {
  value = data.oneview_volume.volume.uri
}

