provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Update SNMPv3 Trap Destination - changing port 162 to 190
data "oneview_appliance_snmpv3_user" "snmptrap" {
    user_name = "<user_name_terra>"
}