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

resource "oneview_ethernet_network" "ethernetnetwork" {
  name               = "TestEthNetwork_terraform"
  type               = "ethernet-networkV4"
  vlan_id            = 100
  initial_scope_uris = [data.oneview_scope.scope_obj.uri]
}

# Updates the resource created above
# To update uncomment the below and ad the attributes  to be updated

resource "oneview_ethernet_network" "ethernetnetwork" {
  name               = "TestEthNetwork_terraform_Rename"
  type               = "ethernet-networkV4"
  vlan_id            = "102"
  initial_scope_uris = [data.oneview_scope.scope_obj.uri]
  depends_on         = [oneview_ethernet_network.ethernetnetwork]
}

#Testing data source
data "oneview_ethernet_network" "ethernetnetwork" {
  name       = "TestEthNetwork-updatenew"
  depends_on = [oneview_ethernet_network.ethernetnetwork]
}

#Importing Existing resource
resource "oneview_ethernet_network" "import_eth" {
}

