provider "oneview" {
  ov_username   = "administrator"
  ov_password   = "admin123"
  ov_endpoint   = "https://10.1.20.1"
  ov_sslverify  = false
  ov_apiversion = 2000
  ov_ifmatch    = "*"
}

// 1 - IMPORTING RESOURCE FIRST

# resource "oneview_network_set" "network_set" { 
# }

# resource "oneview_logical_interconnect_group" "logical_interconnect_group" {
# }
# resource "oneview_ethernet_network" "network_Prod_60" {
  
# }


# // 2 - ADDING THE NEW NETWORK 

# # // CREATION OF ETHERNET NETWORK
resource "oneview_ethernet_network" "network_Prod_60" {
  name    = "Prod_60"
  type    = "ethernet-networkV4"
  vlan_id = 60
}



# /* GETTING ETHERNET NETWORK URIS FROM THE NETWORK SET */
data "oneview_network_set" "network_set" {
  name = "Prod"
}

output "network_set_Prod_uris" {
  value = "${data.oneview_network_set.network_set.network_uris}"
}

data "oneview_ethernet_network" "network_Prod_60" {
  name = "Prod_60"
}
output "oneview_ethernet_network" {
  value = "${data.oneview_ethernet_network.network_Prod_60.uri}"
}
resource "oneview_network_set" "network_set" {
  name = "Prod"
  native_network_uri = ""
  network_set_type = "Regular"
  type = "network-setV5"
  network_uris = ["${data.oneview_network_set.network_set.network_uris}", "${data.oneview_ethernet_network.network_Prod_60.uri}"]
  depends_on = ["oneview_ethernet_network.network_Prod_60"]
}

# /* GETTING NETWORK URIs FOR THE LIG DECLARATION*/
data "oneview_ethernet_network" "network_Mgmt" {
  name = "Mgmt"
}

data "oneview_fc_network" "network_SAN_A_FC" {
  name = "SAN A FC"
}

data "oneview_fc_network" "network_SAN_B_FC" {
  name = "SAN B FC"
}

output "network_SAN_B_FC_uri" {
  value = "${data.oneview_fc_network.network_SAN_B_FC.uri}"
}

# /* GETTING THE UPLINK SET URIs FOR THE LIG DECLARATION*/
# data "oneview_uplink_set" "uplink_set_Mgmt" {
#   name = "Mgmt"
# }

# output "oneview_uplink_set_value" {
#   value = "${data.oneview_uplink_set.uplink_set_Mgmt.port_config_infos}"
# }

# data "oneview_ethernet_network" "network_Prod_60" {
#   name = "Prod_60"
# }

# // ADDING NEW NETWORK URI TO LOGICAL INTERCONNECT GROUP UPLINK SET URIS
resource "oneview_logical_interconnect_group" "logical_interconnect_group" {
  type                  = "logical-interconnect-groupV8"
  interconnect_bay_set  = 3
  enclosure_indexes     = [1, 2, 3]
  redundancy_type       = "HighlyAvailable"
  name                  = "test"
  internal_network_uris = []
  igmp_settings         = []
  sflow_configuration   = []

  interconnect_map_entry_template = [{
    enclosure_index        = 1
    bay_number             = 3
    interconnect_type_name = "Virtual Connect SE 40Gb F8 Module for Synergy"
  },
    {
      enclosure_index        = 2
      bay_number             = 6
      interconnect_type_name = "Virtual Connect SE 40Gb F8 Module for Synergy"
    },
    {
      enclosure_index        = 3
      bay_number             = 3
      interconnect_type_name = "Synergy 20Gb Interconnect Link Module"
    },
    {
      enclosure_index        = 3
      bay_number             = 6
      interconnect_type_name = "Synergy 20Gb Interconnect Link Module"
    },
    {
      enclosure_index        = 1
      bay_number             = 6
      interconnect_type_name = "Synergy 20Gb Interconnect Link Module"
    },
    {
      enclosure_index        = 2
      bay_number             = 3
      interconnect_type_name = "Synergy 20Gb Interconnect Link Module"
    },
  ]
  
  uplink_set = [
    {
      name         = "Mgmt"
      # network_type = "Ethernet"
      # lacp_timer   = "Short"
      # mode         = "Auto"
      network_uris = ["${data.oneview_ethernet_network.network_Mgmt.uri}"]
    },
    {
      name         = "US-ESX-vMotion"
      
    },
    
    {
      name         = "US-SAN-A-FC"
      # network_type = "Ethernet"
      # lacp_timer   = "Short"
      # mode         = "Auto"
      # network_uris = ["${data.oneview_ethernet_network.network_Mgmt.uri}"]

      # logical_port_config = [
      # 	{
      # 	 port_num = [63]
      # 	 bay_num = 6
      # 	 enclosure_num = 2
      # 	 primary_port = "false"
      # 	 desired_speed = "Auto"
      # 	 },
      # 	 {
      # 	 port_num = [63]
      # 	 bay_num = 3
      # 	 enclosure_num = 1
      # 	 primary_port = "false"
      # 	 desired_speed = "Auto"
      # 	 }
      # ]
    },
    {
      name         = "US-Image Streamer"
      # network_type = "Ethernet"
      # lacp_timer   = "Short"
      # mode         = "Auto"
      # //network_uris = ["${data.oneview_ethernet_network.network_Mgmt.uri}"]
    },
    {
      name         = "US-Prod"
      //network_type = "FibreChannel"
      network_uris = ["${data.oneview_network_set.network_set.network_uris}"]

      # logical_port_config = [
      # 	{
      # 	 port_num = [67]
      # 	 bay_num = 3
      # 	 enclosure_num = 1
      # 	 primary_port = "false"
      # 	 desired_speed = "Auto"
      #      }
      # ]
    },
     {
      name         = "US-SAN-B-FC"
    #   //network_type = "Ethernet"
    #   //lacp_timer   = "Short"
    #   mode         = "Auto"
    #  // network_uris = ["${data.oneview_ethernet_network.network_Mgmt.uri}"]
    },
    {
      name         = "US-ESX-Mgmt"
      # network_type = "Ethernet"
      # lacp_timer   = "Short"
      # mode         = "Auto"
      # //network_uris = ["${data.oneview_ethernet_network.network_Mgmt.uri}"]
    },   
    # {
    #   name         = "US-Prod"
    #   network_type = "Ethernet"
    #   lacp_timer   = "Short"
    #   mode         = "Auto"
    #   network_uris = ["${data.oneview_network_set.network_set.network_uris}"]

    #   # logical_port_config = [
    #   # 	{
    #   # 	 port_num = [65]
    #   # 	 bay_num = 3
    #   # 	 enclosure_num = 1
    #   # 	 primary_port = "false"
    #   # 	 desired_speed = "Auto"
    #   # 	 },
    #   # 	 {
    #   # 	 port_num = [65]
    #   # 	 bay_num = 6
    #   # 	 enclosure_num = 2
    #   # 	 primary_port = "false"
    #   # 	 desired_speed = "Auto"
    #   # 	}
    #   # ]
    # },
    # {
    #   name         = "US-SAN-B-FC"
    #   network_type = "FibreChannel"
    #   network_uris = ["${data.oneview_fc_network.network_SAN_B_FC.uri}"]

    #   # logical_port_config = [
    #   # 	{
    #   # 	 port_num = [67]
    #   # 	 bay_num = 6
    #   # 	 enclosure_num = 2
    #   # 	 primary_port = "false"
    #   # 	 desired_speed = "Auto"
    #   # 	 }
    #   # ]
    # },
   
    
  ]

  depends_on = ["oneview_ethernet_network.network_Prod_60"]
}

# /* GETTING THE LI URI*/
# data "oneview_logical_interconnect" "logical_interconnect_read" {
#         name = "LE-Synergy-Local-LIG-FlexFabric"
# }


# // PERFORMING UPDATE FROM GROUP ON LOGICAL INTERCONNECT TO BRING BACK IT TO CONSISTENT STATE 
# resource "oneview_logical_interconnect" "logical_interconnect" {
# 	update_type = "updateComplianceById"
# 	//uri = "${data.oneview_logical_interconnect.logical_interconnect_read.uri}"
# 	depends_on = ["oneview_logical_interconnect_group.logical_interconnect_group"]
# }


# # // ADDING THE NEW NETWORK TO THE NETWORK SET  
# resource "oneview_network_set" "network_set" { 
#   name = "Prod"
#   native_network_uri = ""
#   network_set_type = "Regular"
#   type = "network-setV5"
#   network_uris = ["${data.oneview_network_set.network_set_Prod.network_uris}","${oneview_ethernet_network.network_Prod_60.uri}"]
#   depends_on = ["oneview_logical_interconnect.logical_interconnect"]
#  }

