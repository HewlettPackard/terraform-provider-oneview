provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2800
  ov_ifmatch    = "*"
}

#Importing Existing resource
resource "oneview_fc_network" "import_fc" {
}

