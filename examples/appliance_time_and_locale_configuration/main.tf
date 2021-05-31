provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 3000
  ov_ifmatch    = "*"
}

#Create Appliance Time and Locale
resource "oneview_appliance_time_and_locale" "timelocale" {
    locale = "en_US.UTF-8"
    date_time = "2014-09-11T12:10:33"
    timezone = "UTC"
    ntp_servers = ["16.110.135.123", "16.85.40.52"]
}