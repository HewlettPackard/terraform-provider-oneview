provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# To update typical and maximum bandwidth values of any network, you will need to fetch the respective connection template uri from a network as below: 
/*
data "oneview_ethernet_network" "network" {
  name = "test-network"
}
output "connection_template_uri" {
  value = data.oneview_ethernet_network.network.connection_template_uri
}
*/
# Then you will need to import the connection template and update typical and maximum bandwidth values. 
# To import an existing connection template, retain this configuration and run below command. 
# $ terraform import oneview_connection_templates.connectionTemplates <connection-template-uri>
# Then just update the update_resource.tf with your desired values and execute it.

resource "oneview_connection_templates" "connectionTemplates" {
}
