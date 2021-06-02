provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 3000
  ov_ifmatch    = "*"
}

# Data source for oneview version

data "oneview_version" "ver"  {

}
output "oneview_version_value" {
value = data.oneview_version.ver.current_version
}