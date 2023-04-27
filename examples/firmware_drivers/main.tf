provider "oneview" {
	ov_username   = var.username
	ov_password   = var.password
	ov_endpoint   = var.endpoint
	ov_sslverify  = var.ssl_enabled
	ov_apiversion = var.api_version
	ov_ifmatch    = "*"
  }

data "oneview_firmware_drivers" "spp" {
  id = "Synergy_Custom_SPP_2021.02.01_Z7550-97110"
}

data "oneview_firmware_drivers" "hotfix" {
  id = "HPE ProLiant Gen10 and Gen10 Plus Smart Array Controller Driver for VMware vSphere 6.5 (Driver Component)."
}

data "oneview_scope" "scope" {
  name = "Auto-Scope"
}

# Creating Custom Firmware Service Pack
resource "oneview_firmware_drivers" "drivers" {
  baseline_uri = "data.oneview_firmware_drivers.spp.uri"
  hotfix_uris = [data.oneview_firmware_drivers.hotfix.uri]
  custom_baseline_name = "Terraform_SPP"
  initial_scope_uris = [data.oneview_scope.scope.uri]
}
