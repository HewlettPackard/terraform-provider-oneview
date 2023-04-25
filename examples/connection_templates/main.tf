provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

data "oneview_ethernet_network" "network" {
  name = "Auto_Ethernet_1"
}
output "connection_template_uri" {
  value = data.oneview_ethernet_network.network.connection_template_uri
}


#resource "oneview_connection_templates" "connectionTemplates" {
#}
