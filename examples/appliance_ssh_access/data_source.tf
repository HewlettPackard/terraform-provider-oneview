provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Testing data source
data "oneview_appliance_ssh_access" "sshAccess" {
}

output "oneview_appliance_ssh_access_value" {
  value = data.oneview_appliance_ssh_access.sshAccess
}