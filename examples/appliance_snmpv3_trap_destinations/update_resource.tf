provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

data "oneview_appliance_snmpv3_user" "snmptrap" {
    user_name = "user"
}


# Update SNMPv3 Trap Destination - changing port 162 to 190
resource "oneview_appliance_snmpv3_trap_destinations" "snmptrap" {
    destination_address = "1.1.1.1"
    port = 190
    user_id = data.oneview_appliance_snmpv3_user.snmptrap.user_id
}
