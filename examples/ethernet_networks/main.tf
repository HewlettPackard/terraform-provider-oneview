provider "oneview" {
        ov_username =   "${var.username}"
        ov_password =   "${var.password}"
        ov_endpoint =   "${var.endpoint}"
        ov_sslverify =  "${var.ssl_enabled}"
        ov_apiversion = 2800
        ov_ifmatch = "*"

}

# Creates Ethernet Network Resource
resource "oneview_ethernet_network" "ethernetnetwork" {
	name = "TestEthNetwork_terraform"
	type = "ethernet-networkV4"
	vlan_id = 100
}
