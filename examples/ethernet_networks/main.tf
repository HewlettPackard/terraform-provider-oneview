provider "oneview" {
        ov_username =   "${var.username}"
        ov_password =   "${var.password}"
        ov_endpoint =   "${var.endpoint}"
        ov_sslverify =  "${var.ssl_enabled}"
        ov_apiversion = 2600
        ov_ifmatch = "*"

}

# Creates 3 Ethernet Network Resource
resource "oneview_ethernet_network" "ethernetnetwork" {
	name = "TestEthNetwork_terraform"
	type = "ethernet-networkV4"
	vlan_id = 100
}

resource "oneview_ethernet_network" "ethernetnetwork_1" {
	name = "Auto-Ethernet-1"
	type = "ethernet-networkV4"
	vlan_id = 101
}

resource "oneview_ethernet_network" "ethernetnetwork_2" {
	name = "Auto-Ethernet-2"
	type = "ethernet-networkV4"
	vlan_id = 102
}
