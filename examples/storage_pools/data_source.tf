provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Reads Storage Pool resource
data "oneview_storage_pool" "storage_pool" {
  name = "<storage_pool_terraform>"
}

output "oneview_storage_pool_value" {
  value = data.oneview_storage_pool.storage_pool.uri
}

