provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}


# Fetching Network Resource Uri
data "oneview_ethernet_network" "ethernetnetworks" {
  name = "<network_name>"
}

# Updates created labels after running main.tf 
resource "oneview_label" "anyName" {
        resource_uri = data.oneview_ethernet_network.ethernetnetworks.uri
        labels {
                name = "<label_rename_1>"
        }
        labels{
                name = "<label_rename_2>"
        }
      	labels {
           name = "<label_name_3>" 
	}	
}
