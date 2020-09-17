provider "oneview" {
  ov_username   = <ov_username>
  ov_password   = <ov_password>
  ov_endpoint   = <ov_ip>
  ov_sslverify  = false
  ov_apiversion = <ov_apiversion>
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
	name = "TestNetworkSet_update"
	native_network_uri = ""
	type = "network-setV5"
	network_uris = ["${data.oneview_ethernet_network.eth.uri}",]
	initial_scope_uris = ["${data.oneview_scope.scope_obj.uri}"]
}

resource "oneview_network_set" "NetworkSet" {
	name = "TestNetworkSet_update"
	native_network_uri = ""
	type = "network-setV5"
	network_uris = ["${data.oneview_ethernet_network.eth.uri}","${data.oneview_ethernet_network.eth1.uri}"]
	initial_scope_uris = ["${data.oneview_scope.scope_obj.uri}"]
}

// Example for data source
data "oneview_network_set" "network_set" {
        name = "TestNetworkSet_update"
}

output "oneview_network_set_value" {
        value = "${data.oneview_network_set.network_set.uri}"
}

//Importing an existing resource from appliance
resource "oneview_network_set" "import_ns" {
}