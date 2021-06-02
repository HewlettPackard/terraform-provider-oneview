provider "oneview" {
	ov_username   = var.username
	ov_password   = var.password
	ov_endpoint   = var.endpoint
	ov_sslverify  = var.ssl_enabled
	ov_apiversion = 3000
	ov_ifmatch    = "*"
  }
  
# Creating Custom Firmware Service Pack
resource "oneview_firmware_drivers" "drivers" {
  baseline_uri = "/rest/firmware-drivers/Synergy_Custom_SPP_2021_02_01_Z7550-97110"
  hotfix_uris = ["/rest/firmware-drivers/cp042959"]
  custom_baseline_name = "Terraform_SPP"
  initial_scope_uris = ["/rest/scopes/2f4cbc4f-a1a2-4073-bdf4-829041b3398b", "/rest/scopes/bda31962-f4e5-4969-b036-bc0c63b0ccd3"]
}