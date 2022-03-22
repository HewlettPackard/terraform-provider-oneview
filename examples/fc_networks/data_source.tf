provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Testing data source
data "oneview_fc_network" "fc_network" {
  name = "TestFCNetwork_Renamed"
}

output "oneview_fc_network_value" {
  value = data.oneview_fc_network.fc_network.fabric_type
}

/*
#Importing Existing resource
resource "oneview_fc_network" "import_fc"{
}
*/
