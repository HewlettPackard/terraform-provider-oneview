provider "oneview" {
	ov_username   = var.username
	ov_password   = var.password
	ov_endpoint   = var.endpoint
	ov_sslverify  = var.ssl_enabled
	ov_apiversion = var.api_version
	ov_ifmatch    = "*"
  }

data "oneview_firmware_drivers" "spp" {
  name = "Service Pack for Synergy"
  version="SY-2021.02.01"
}

data "oneview_firmware_drivers" "hotfix" {
  name = "HPE ProLiant Gen10 and Gen10 Plus Smart Array Controller Driver for VMware vSphere 6.5 (Driver Component)."
  version="2021.04.01"
}

data "oneview_scope" "scope" {
  name = "ScopeTest"
}

# Creating Custom Firmware Service Pack
resource "oneview_firmware_drivers" "drivers" {
  baseline_uri = data.oneview_firmware_drivers.spp.uri
  hotfix_uris = [data.oneview_firmware_drivers.hotfix.uri]
  custom_baseline_name = "Terraform_SPP"
  initial_scope_uris = [data.oneview_scope.scope.uri]
}