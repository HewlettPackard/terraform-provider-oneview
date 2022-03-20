provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

data "oneview_enclosure" "enclosure" {
  name = "0000A66101"
}

resource "oneview_enclosure" "import_enc" {
  op                  = "replace"
  path                = "/name"
  value               = "0000A66101 Renamed"
  name                = "0000A66101"
  enclosure_group_uri = data.oneview_enclosure.enclosure.enclosure_group_uri
}

