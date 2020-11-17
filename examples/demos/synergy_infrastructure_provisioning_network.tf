provider "oneview" {
  ov_username =   "${var.username}"
  ov_password =   "${var.password}"
  ov_endpoint =   "${var.endpoint}"
  ov_sslverify =  "${var.ssl_enabled}"
  ov_apiversion = 2200
  ov_ifmatch = "*"
}

/*	Below example creates ethernet network to be added as internal network 
	and fetches another ethernet network to be added to logical interconnect group resource:-
*/

# CREATION OF ETHERNET NETWORK 
resource "oneview_ethernet_network" "ethernet_network" {
	name = "<network_name>"
	type = "ethernet-networkV4"
	vlan_id = 171
}

# GET THE ETHERNET NETWORK TO GET THE URI TO ASSIGN TO UPLINKSET 
data "oneview_ethernet_network" "eth_net" {
        name = "<network_name>"
}

# ADDING NEWTWORK TO LOGICAL INTERCONNECT GROUP USING 'internalNetworkUris' 
resource "oneview_logical_interconnect_group" "logical_interconnect_group" {
	type = "logical-interconnect-groupV8"
	name = "<LIG_name>"
	interconnect_bay_set = 3
	enclosure_indexes = [1, 2, 3]
	redundancy_type = "HighlyAvailable"
	internal_network_uris = ["${oneview_ethernet_network.ethernet_network.uri}"]
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


# PERFORMING UPDATE FROM GROUP ON LOGICAL INTERCONNECT TO BRING BACK IT TO CONSISTENT STATE  

/*
   To import existing LIG and LI resources from appliance into the teriraform instances follow the below steps.
	1. Uncomment the below code snippet, which contains empty resource declarations and comment the rest of the script from that point.
	2. Save the file and execute the following commands to import.
		terraform import oneview_logical_interconnect_group.logical_interconnect_group <resource_name>
		terraform import oneview_logical_interconnect.logical_interconnect <resource_id>
	Example:
		terraform import oneview_logical_interconnect_group.logical_inte
rconnect_group TestLIG
		terraform import oneview_logical_interconnect.li d4468f89-4442-4324-9c01-624c7382db2d
	3. Once the resources are imported, comment the below empty resource declarations and uncomment the rest of the script and execute the script for the rest of the operations to be performed.
*/

/*
resource "oneview_logical_interconnect_group" "logical_interconnect_group" {
}


resource "oneview_logical_interconnect" "logical_interconnect"{
}
*/

/* 
resource "oneview_logical_interconnect" "logical_interconnect" {
	update_type = "updateComplianceById"
	depends_on = ["oneview_logical_interconnect_group.logical_interconnect_group"]
}
*/
