/*   Infrastructure Provisioning with Networks and Storage

	Be able to provision compute (with server settings), networking, and storage.
	Create a server profile template with the following options:
		Network connections
		Boot mode
		Boot settings
		Create a server profile from a server profile template and assign to hardware
		Power on server
*/

provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2200
  ov_ifmatch    = "*"
}

# Create Networks
resource "oneview_ethernet_network" "ethernetnetwork" {
  name    = "<network_name>"
  type    = "ethernet-networkV4"
  vlan_id = 100
}

resource "oneview_fc_network" "fc_network" {
  name = "<network_name>"
  type = "fc-networkV4"
}

# Create Logical Interconnect Group
resource "oneview_logical_interconnect_group" "logical_interconnect_group" {
  name                 = "<LIG_name>"
  type                 = "logical-interconnect-groupV8"
  interconnect_bay_set = 3
  redundancy_type      = "HighlyAvailable"
  enclosure_indexes    = [1, 2, 3]

  interconnect_map_entry_template {
    bay_number             = 3
    interconnect_type_name = "Virtual Connect SE 40Gb F8 Module for Synergy"
    enclosure_index        = 1
  }
  interconnect_map_entry_template {
    bay_number             = 6
    interconnect_type_name = "Synergy 20Gb Interconnect Link Module"
    enclosure_index        = 1
  }
  interconnect_map_entry_template {
    bay_number             = 3
    interconnect_type_name = "Synergy 20Gb Interconnect Link Module"
    enclosure_index        = 2
  }
  interconnect_map_entry_template {
    bay_number             = 6
    interconnect_type_name = "Virtual Connect SE 40Gb F8 Module for Synergy"
    enclosure_index        = 2
  }
  interconnect_map_entry_template {
    bay_number             = 3
    interconnect_type_name = "Synergy 20Gb Interconnect Link Module"
    enclosure_index        = 3
  }
  interconnect_map_entry_template {
    bay_number             = 6
    interconnect_type_name = "Synergy 20Gb Interconnect Link Module"
    enclosure_index        = 3
  }
  uplink_set {
    network_type          = "Ethernet"
    ethernet_network_type = "Tagged"
    name                  = "test"
    logical_port_config {
      desired_speed = "Auto"
      port_num      = [66]
      bay_num       = 3
      enclosure_num = 1
      primary_port  = false
    }
    mode         = "Auto"
    network_uris = [oneview_ethernet_network.ethernetnetwork.uri]
  }
}

# Create Enclosure Group
resource "oneview_enclosure_group" "enclosure_group" {
  name               = "<EG_name>"
  ip_addressing_mode = "DHCP"
  enclosure_count    = 3
  type               = "EnclosureGroupV7"
  interconnect_bay_mappings {
    interconnect_bay               = 3
    logical_interconnect_group_uri = oneview_logical_interconnect_group.logical_interconnect_group.uri
  }
  interconnect_bay_mappings {
    interconnect_bay               = 6
    logical_interconnect_group_uri = oneview_logical_interconnect_group.logical_interconnect_group.uri
  }
  depends_on = [oneview_logical_interconnect_group.logical_interconnect_group]
}

# Create Logical Enclosure
resource "oneview_logical_enclosure" "logical_enclosure" {
  name                = "<LE_name>"
  enclosure_uris      = ["/rest/enclosures/0000000000A66101", "/rest/enclosures/0000000000A66102", "/rest/enclosures/0000000000A66103"]
  enclosure_group_uri = oneview_enclosure_group.enclosure_group.uri
  depends_on          = [oneview_enclosure_group.enclosure_group]
}

# Get Server Hardware Type
data "oneview_server_hardware_type" "server_hardware_type" {
  name = "SY 480 Gen9 1"
}

# Get Server Hardware 
data "oneview_server_hardware" "server_hardware" {
  name = "0000A66101, bay 7"
}

# Create Server Profile Template with Network and BOOT
resource "oneview_server_profile_template" "server_profile_template" {
  name                 = "<SPT_name>"
  type                 = "ServerProfileTemplateV8"
  enclosure_group      = oneview_enclosure_group.enclosure_group.name
  server_hardware_type = data.oneview_server_hardware_type.server_hardware_type.name
  network {
    name           = "UcFCNetConn"
    function_type  = "FibreChannel"
    port_id        = "Mezz 2:1"
    network_uri    = oneview_fc_network.fc_network.uri
    requested_mbps = ""
  }
  network {
    name           = "UcEthNetConn"
    function_type  = "Ethernet"
    port_id        = "Mezz 3:1-a"
    network_uri    = oneview_fc_network.fc_network.uri
    requested_mbps = ""
  }
  manage_connections = true
  boot_order         = ["HardDisk"]
  boot_mode {
    manage_mode     = true
    mode            = "UEFI"
    pxe_boot_policy = "Auto"
  }
  depends_on = [oneview_logical_enclosure.logical_enclosure]
}

# Create Server Profile with Network and BOOT
resource "oneview_server_profile" "server_profile" {
  name          = "<SP_name>"
  template      = oneview_server_profile_template.server_profile_template.name
  hardware_name = data.oneview_server_hardware.server_hardware.name
  type          = "ServerProfileV12"
  depends_on    = [oneview_logical_enclosure.logical_enclosure]
}

