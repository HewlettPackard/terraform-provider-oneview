
provider "oneview" {
	ov_username = <ov-username>
	ov_password = <ov-password>
	ov_endpoint = <ov-endpoint>
	ov_sslverify = false
	ov_apiversion = <ov-apiversion>
}

/*
	While executing the following script follow the below steps
	1. Use target option to update/delete the specific resource. Otherwise there is chance of deleting the Network and NetworkSet first and then updating the LIG and LI.
	2. To update the LIG and LI follow the below commands
		terraform apply -target=oneview_logical_interconnect_group.logical_interconnect_group -target=oneview_logical_interconnect.logical_interconnect
	3. Once the LIG and LI are updated, delete the NetworkSet and Network in order using the below commands.
		terraform destroy -target=oneview_network_set.network_et
		terraform destroy -target=oneview_ethernet_network.ethernet_network	
*/

/* GET THE ETHERNET NETWORK TO GET THE URI TO ASSIGN TO UPLINKSET*/
data "oneview_ethernet_network" "eth_net" {
        name = "TestEth"
}

/* REMOVING THE NETWORK FROM LIG */
resource "oneview_logical_interconnect_group" "logical_interconnect_group" {
	type = "logical-interconnect-groupV5"
	name = "LIG"
	interconnect_bay_set = 3
	enclosure_indexes = [1, 2, 3]
	redundancy_type = "HighlyAvailable"
	interconnect_map_entry_template = [{
	  enclosure_index = 1
	  bay_number = 3
	  interconnect_type_name = "Virtual Connect SE 40Gb F8 Module for Synergy"
	 },
	 {
	  enclosure_index = 2
	  bay_number = 6
	  interconnect_type_name = "Virtual Connect SE 40Gb F8 Module for Synergy"
	 },
	 {
	  enclosure_index = 2
	  bay_number = 3
	  interconnect_type_name = "Synergy 20Gb Interconnect Link Module"
	 },
	 {
	  enclosure_index = 3
	  bay_number = 3
	  interconnect_type_name = "Synergy 20Gb Interconnect Link Module"
	 },
	 {
	  enclosure_index = 1
	  bay_number = 6
	  interconnect_type_name = "Synergy 20Gb Interconnect Link Module"
	 },
	 {
	  enclosure_index = 3
	  bay_number = 6
	  interconnect_type_name = "Synergy 20Gb Interconnect Link Module"
	 }]

	uplink_set = [{
		network_type = "Ethernet"
		network_uris = ["${data.oneview_ethernet_network.eth_net.uri}"]
		name = "UPlinkSet1"
		logical_port_config = [{
			 port_num = [61]
			 bay_num = 3
			 enclosure_num = 1
		}]
	}]
}

/* PERFORMING UPDATE FROM GROUP ON LOGICAL INTERCONNECT TO BRING BACK IT TO CONSISTENT STATE */
resource "oneview_logical_interconnect" "logical_interconnect" {
	update_type = "updateComplianceById"
	depends_on = ["oneview_logical_interconnect_group.logical_interconnect_group"]
}
