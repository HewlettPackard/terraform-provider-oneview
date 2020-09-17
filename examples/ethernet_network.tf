provider "oneview" {
        ov_username = "<ov-username>"
        ov_password = "<ov-password>"
        ov_endpoint = "<ov-endpoint>"
        ov_sslverify = false
        ov_apiversion = <ov-apiversion>
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
	vlan_id = "100"
	initial_scope_uris = ["${data.oneview_scope.scope_obj.uri}"] 
	
}*/

/*
#Testing data source
data "oneview_ethernet_network" "ethernetnetwork" {
  name = "TestEthNetwork-updatenew"
}

/*
#Importing Existing resource
resource "oneview_ethernet_network" "import_eth"{
}
*/
