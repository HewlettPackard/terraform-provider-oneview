provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

data  "oneview_appliance_snmpv3_user" "snmpv3user" {
  user_name = "<user_name_terra>"
}

# Creates SNMPv3 Trap Destination
resource "oneview_appliance_snmpv3_trap_destinations" "snmptrap" {
    destination_address = "1.1.1.1"
    port = 162
    user_id = data.oneview_appliance_snmpv3_user.snmpv3user.user_id
}