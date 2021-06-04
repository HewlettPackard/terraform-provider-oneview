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
  name = "Auto-Ethernet-1"
}

# Create Labels for the Ethernet Network
resource "oneview_label" "anyName" {
	resource_uri = data.oneview_ethernet_network.ethernetnetworks.uri
        labels {
          	name = "NewLabel"
        }
	labels{
		name = "NewLabele2UpdatedwithoutUri"
	}
}
/* Importing Existing resource
resource "oneview_label" "anyName" {
}
*/
