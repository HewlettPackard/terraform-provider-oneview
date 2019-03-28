provider "oneview" {
	ov_username = "administrator"
	ov_password = "madhav123"
	ov_endpoint = "https://10.170.16.44"
	ov_sslverify = false
	ov_apiversion = 800
	ov_ifmatch = "*"
}

data "oneview_scope" "scope_obj" {
        name = "Scope_Sample"
}


resource "oneview_ethernet_network" "ethernetnetwork" {
	name = "TestEthNetwork_terraform"
	type = "ethernet-networkV4"
	vlan_id = 100
	initial_scope_uris = ["${data.oneview_scope.scope_obj.uri}"] 
}

# Updates the resource created above
# To update uncomment the below and ad the attributes  to be updated

/*resource "oneview_ethernet_network" "EthernetNetwork" {
	name = "TestEthNetwork-updatenew"
	type = "ethernet-networkV4"
	vlan_id = "10"
	initial_scope_uris = ["${data.oneview_scope.scope_obj.uri}"] 
	
}*/
