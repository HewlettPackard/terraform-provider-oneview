provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Testing data source
data "oneview_rack_manager" "dsrm"{
  name="5UF7201000"
}

output "dsrmop" {
  value = data.oneview_rack_manager.dsrm.uri
}
/*
#Importing Existing resource
resource "oneview_rack_manager" "import_rm"{
}
*/
