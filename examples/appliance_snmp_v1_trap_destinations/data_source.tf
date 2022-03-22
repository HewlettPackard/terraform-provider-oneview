provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

data "oneview_appliance_snmp_v1_trap_destinations" "snmp_v1" {
        destination_id = "4"
}

output "oneview_snmp_v1_trap_destinations_value" {
        value = data.oneview_appliance_snmp_v1_trap_destinations.snmp_v1
}

