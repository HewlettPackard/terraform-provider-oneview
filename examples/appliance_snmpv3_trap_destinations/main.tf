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

# Creates SNMPv3 Trap Destination
resource "oneview_appliance_snmpv3_trap_destinations" "snmptrap" {
    destination_address = "<destination_address>"
    port = 162
    user_id = data.oneview_appliance_snmpv3_user.snmptrap.user_id
}