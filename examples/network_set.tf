provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2400
  ov_ifmatch    = "*"
}

data "oneview_ethernet_network" "eth" {
  name = "Prod_1104"
}

data "oneview_ethernet_network" "eth1" {
  name = "Prod_1103"
}

data "oneview_scope" "scope_obj" {
  name = "test"
}

resource "oneview_network_set" "NetworkSet" {
  name               = "TestNetworkSet_update"
  native_network_uri = ""
  type               = "network-setV5"
  network_uris       = [data.oneview_ethernet_network.eth.uri]
  initial_scope_uris = [data.oneview_scope.scope_obj.uri]
}
/*
resource "oneview_network_set" "NetworkSet" {
  name               = "TestNetworkSet_update"
  native_network_uri = ""
  type               = "network-setV5"
  network_uris       = [data.oneview_ethernet_network.eth.uri, data.oneview_ethernet_network.eth1.uri]
  initial_scope_uris = [data.oneview_scope.scope_obj.uri]
}

# Example for data source
data "oneview_network_set" "network_set" {
  name = "TestNetworkSet_update"
  depends_on = [oneview_network_set.NetworkSet]
}

output "oneview_network_set_value" {
  value = data.oneview_network_set.network_set.uri
  depends_on = [oneview_network_set.network_set]
}

# Importing an existing resource from appliance
resource "oneview_network_set" "import_ns" {
}
*/
