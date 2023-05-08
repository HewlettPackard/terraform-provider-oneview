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
  name = "Gen10 Service Pack for ProLiant"
  version="2022.09.01.00"
}

output "firmware_drivers_value" {
  value = data.oneview_firmware_drivers.drivers
}
