provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Testing data source
data "oneview_firmware_drivers" "drivers" {
  id = "HPE Synergy Service Pack,SY-2021.11.01"
}

output "firmware_drivers_value" {
  value = data.oneview_firmware_drivers.drivers
}
