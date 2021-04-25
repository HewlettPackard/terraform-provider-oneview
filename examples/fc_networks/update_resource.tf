provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2800
  ov_ifmatch    = "*"
}

# Updates the created resource with local name FCNetwork
resource "oneview_fc_network" "FCNetwork" {
  name = "TestFCNetwork_Renamed"
  type = "fc-networkV4"
}

