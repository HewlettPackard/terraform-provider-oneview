/*   Infrastructure Provisioning without Storage and Image Stramear

	Be able to provision compute (with server settings), networking, and storage.
	Create a server profile template with the following options:
		Network connections
		Boot mode
		Boot settings
		Create a server profile from a server profile template and assign to hardware
		Power on server
*/

provider "oneview" {
	ov_username = "<ov_username>"
	ov_password = "<ov_password>"
	ov_endpoint = "<ov_endpoint>"
	ov_sslverify = false
	ov_apiversion = <ov_apiversion>
	ov_ifmatch = "*"
}

//Create Networks
resource "oneview_ethernet_network" "ethernetnetwork" {
  name = "UCENet"
  type = "ethernet-networkV4"
  vlan_id = 100
}

resource "oneview_fc_network" "fc_network" {
  name = "UCFCNET"
  type = "fc-networkV4"
}

resource "oneview_logical_interconnect_group" "logical_interconnect_group" {
  name = "UCLIG"
  type = "logical-interconnect-groupV5"
  interconnect_bay_set = 3
  redundancy_type = "HighlyAvailable"
  enclosure_indexes = [1, 2, 3]

  interconnect_map_entry_template = [
    {
      bay_number             = 3
      interconnect_type_name = "Virtual Connect SE 40Gb F8 Module for Synergy"
      enclosure_index = 1
    },
    {
      bay_number             = 6
      interconnect_type_name = "Synergy 20Gb Interconnect Link Module"
      enclosure_index = 1
    },
    {
      bay_number             = 3
      interconnect_type_name = "Synergy 20Gb Interconnect Link Module"
      enclosure_index = 2
    },
    {
      bay_number             = 6
      interconnect_type_name = "Virtual Connect SE 40Gb F8 Module for Synergy"
      enclosure_index = 2
    },
    {
      bay_number             = 3
      interconnect_type_name = "Synergy 20Gb Interconnect Link Module"
      enclosure_index = 3
    },
    {
      bay_number             = 6
      interconnect_type_name = "Synergy 20Gb Interconnect Link Module"
      enclosure_index = 3
    },

  ]
}

// Create enclosure group
resource "oneview_enclosure_group" "enclosure_group" {
  name = "UCEG"
  ip_addressing_mode = "DHCP"
  enclosure_count = 3
  type = "EnclosureGroupV7"
  interconnect_bay_mappings = [
    {
      interconnect_bay = 3
      logical_interconnect_group_uri = "${oneview_logical_interconnect_group.logical_interconnect_group.uri}"
    },
    {
      interconnect_bay = 6
      logical_interconnect_group_uri = "${oneview_logical_interconnect_group.logical_interconnect_group.uri}"
    },
  ]
}

resource "oneview_logical_enclosure" "logical_enclosure" {
  name = "UCLE"
  enclosure_uris = ["/rest/enclosures/0000000000A66101","/rest/enclosures/0000000000A66102","/rest/enclosures/0000000000A66103"]
  enclosure_group_uri = "${oneview_enclosure_group.enclosure_group.uri}"
}

// Get Server Hardware Type
data "oneview_server_hardware_type" "server_hardware_type" {
	name = "SY 480 Gen9 1"
}

// Get Server Hardware 
data "oneview_server_hardware" "server_hardware" {
  name = "0000A66101, bay 7"
}

// Create Server Profile Template with Network and BOOT
resource "oneview_server_profile_template" "server_profile_template" {
	name = "UCSPT"
	type = "ServerProfileTemplateV5"
	enclosure_group = "${oneview_enclosure_group.enclosure_group.name}"
	server_hardware_type = "${data.oneview_server_hardware_type.server_hardware_type.name}"
  network = [{
    name = "UcEthNetConn"
    function_type = "FibreChannel"
    port_id = "Mezz 2:1"
    network_uri = "${oneview_fc_network.fc_network.uri}"
    requested_mbps = ""
  }]
  manage_connections = true
  boot_order = ["HardDisk"]
  boot_mode = {
    manage_mode = true
    mode = "UEFI"
    pxe_boot_policy = "Auto"
  }
  depends_on = ["oneview_logical_enclosure.logical_enclosure"]
}

// Create Server Profile with Network and BOOT
resource "oneview_server_profile" "server_profile" {
  name = "UCSP"
  template = "${oneview_server_profile_template.server_profile_template.name}"
  hardware_name = "${data.oneview_server_hardware.server_hardware.name}"
  type = "ServerProfileV9"
  depends_on = ["oneview_logical_enclosure.logical_enclosure"]
}
