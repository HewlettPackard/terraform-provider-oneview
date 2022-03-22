provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Updates the imported resource from main.tf
resource "oneview_storage_pool" "storage_pool" {
  name       = "CPG-SSD-AO"
  is_managed = true
}

