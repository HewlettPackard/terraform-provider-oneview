provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Updates port to 172 and Community String to "Test5"
resource "oneview_appliance_snmp_v1_trap_destinations" "snmp_v1" {
    destination_id = "<destination_id>"
    community_string = "<community_string_rename>"
    destination = "<destination_address>"
    port = 172
}
