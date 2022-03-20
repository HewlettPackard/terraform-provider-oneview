provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Example for data source
data "oneview_network_set" "network_set" {
  name = "TestNetworkSet_update"
}

output "oneview_network_set_value" {
  value = data.oneview_network_set.network_set.uri
}

# Importing an existing resource from appliance
#resource "oneview_network_set" "import_ns" {
#}
