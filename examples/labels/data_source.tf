provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Fetching Network Resource URI
data "oneview_ethernet_network" "ethernetnetworks" {
  name = "Auto-Ethernet-1"
}

# Testing data source
data "oneview_label" "lablesDataSource" {
  resource_uri = data.oneview_ethernet_network.uri
}

output "oneview_label_value" {
  value = data.oneview_label.lablesDataSource.labels
}
