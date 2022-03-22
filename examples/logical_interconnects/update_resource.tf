provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Updates logical Interconnect Port Flap Settings
resource "oneview_logical_interconnect" "logical_interconnect" {
  update_type = "updatePortFlapSettings"
  port_flap_settings {
    port_flap_protection_mode        = "Detect"
    port_flap_threshold_per_interval = 2
    detection_interval               = 20
    no_of_samples_declare_failures   = 2
    name                             = "PortFlapSettingsUpdated"
    consistency_checking             = "ExactMatch"
  }
}

# Returns logical interconnects to a consistent state. The current logical interconnect state is compared to the associated logical interconnect group.
/*
resource "oneview_logical_interconnect" "logical_interconnect" {
	update_type = "updateComplianceById"
}
*/
