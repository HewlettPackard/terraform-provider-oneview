provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Testing data source

data "oneview_storage_volume_template" "d_svt" {
  name = "RenameDemoStorageTemplate"
}

output "oneview_svt_value" {
  value = data.oneview_storage_volume_template.d_svt.root_template_uri
}

# Importing an existing resource from the appliance.
#resource "oneview_storage_volume_template" "svt" {
#}