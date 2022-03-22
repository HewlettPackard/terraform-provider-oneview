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
    destination_id = "81"
    community_string = "Test5"
    destination = "192.0.16.4"
    port = 172
}
