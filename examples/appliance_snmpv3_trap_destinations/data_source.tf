provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Testing data source
data "oneview_appliance_snmpv3_trap_destinations" "snmptrap" {
  id_field = "67003649-af34-4a92-a46a-137855ddc8f7"
}

output "oneview_snmpv3_trap_destinations" {
  value = data.oneview_appliance_snmpv3_trap_destinations.snmptrap
}