provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}
data  "oneview_appliance_snmpv3_user" "snmpv3user" {
  user_name="user"    
}
output "oneview_appliance_snmpv3_user_uri" {
  value = data.oneview_appliance_snmpv3_user.snmpv3user.uri
}