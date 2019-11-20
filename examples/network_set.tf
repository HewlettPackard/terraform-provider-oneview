provider "oneview" {
  ov_username   = <ov_username>
  ov_password   = <ov_password>
  ov_endpoint   = <ov_ip>
  ov_sslverify  = false
  ov_apiversion = <ov_apiversion>
}

resource "oneview_network_set" "NetworkSet" {
	name = "TestNetworkSet_update"
	native_network_uri = ""
	type = "network-setV4"
	network_uris = ["/rest/ethernet-networks/51c17169-b114-4aa2-9688-92f04ac86863"]
	initial_scope_uris = ["/rest/scopes/62a57803-da7d-4b66-aee0-6e5d8854264b"]
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
