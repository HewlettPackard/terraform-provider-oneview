provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 3600
  ov_ifmatch    = "*"
}

#Existing Appliance Time and Locale Configurations
data "oneview_appliance_time_and_locale" "timelocale" {
}

output "locale_value" {
  value = data.oneview_appliance_time_and_locale.timelocale.locale
}