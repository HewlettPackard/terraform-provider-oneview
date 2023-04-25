provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Fetching Newtork Resource URI
data "oneview_ethernet_network" "ethernetnetworks" {
  name = "<network_name>"
}

# Create Labels for the Ethernet Network
resource "oneview_label" "anyName" {
	resource_uri = data.oneview_ethernet_network.ethernetnetworks.uri
        labels {
          	name = "<label_name_1>"
        }
	labels{
		name = "<label_name_2>"
	}
}
/* Importing Existing resource
resource "oneview_label" "anyName" {
}
*/
