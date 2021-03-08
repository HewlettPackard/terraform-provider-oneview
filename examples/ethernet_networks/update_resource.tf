provider "oneview" {
        ov_username =   "${var.username}"
        ov_password =   "${var.password}"
        ov_endpoint =   "${var.endpoint}"
        ov_sslverify =  "${var.ssl_enabled}"
        ov_apiversion = 2600
        ov_ifmatch = "*"

}

# Updates the resource created from main.tf
resource "oneview_ethernet_network" "ethernetnetwork" {
	name = "TestEthNetwork_terraform_Rename"
	type = "ethernet-networkV4"
	vlan_id = 100
}
