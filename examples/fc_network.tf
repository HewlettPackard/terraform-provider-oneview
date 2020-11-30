provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2200
  ov_ifmatch    = "*"
}

data "oneview_scope" "scope_obj" {
  name = "testing"
}

resource "oneview_fc_network" "FCNetwork" {
  name               = "TestFCNetwork"
  type               = "fc-networkV4"
  initial_scope_uris = [data.oneview_scope.scope_obj.uri]
}

resource "oneview_fc_network" "FCNetwork" {
  name = "TestFCNetwork_Renamed"
  type = "fc-networkV4"
}

# Testing data source
data "oneview_fc_network" "fc_network" {
  name = "SYN03_Frame1"
}

output "oneview_fc_network_value" {
  value = data.oneview_fc_network.fc_network.fabric_type
}

#Importing Existing resource
resource "oneview_fc_network" "import_fc" {
}

