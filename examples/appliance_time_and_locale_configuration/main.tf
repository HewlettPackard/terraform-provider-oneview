provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Create Appliance Time and Locale
resource "oneview_appliance_time_and_locale" "timelocale" {
    locale = "en_US.UTF-8"
    timezone = "UTC"
    ntp_servers = ["16.110.135.123", "16.85.40.52"]
}