provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Testing Data Source
data "oneview_scope" "scope" {
  name = "Auto-Scope"
}

#Importing Existing resource
#resource "oneview_scope" "import_scope"{
#}
