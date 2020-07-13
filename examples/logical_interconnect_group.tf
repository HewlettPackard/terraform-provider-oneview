provider "oneview" {
        ov_username = "<ov_username>"
        ov_password = "<ov_password>"
        ov_endpoint = "<ov_endpoint>"
        ov_sslverify = false
        ov_apiversion = <ov_apiversion>
        ov_ifmatch = "*"
}


/* 
# Creates a logical interconnect group or updates if already existing
resource "oneview_logical_interconnect_group" "LIG" {
  name = "TestLIG5"
  type = "logical-interconnect-groupV8"
  interconnect_bay_set = 1
  redundancy_type = "NonRedundantASide"
  enclosure_indexes = [-1]

  igmp_settings = [
   {
      consistency_checking = "ExactMatch",
      igmp_idle_timeout_interval =  260,
      igmp_snooping =  true,
      prevent_flooding = true,
      proxy_reporting =  true,
   },
  ]

  interconnect_map_entry_template = [
    {
      bay_number             = 1
      interconnect_type_name = "Virtual Connect SE 16Gb FC Module for Synergy"
      enclosure_index = -1
    },
  ]
}
*/

/*  
# Test for data source  
data "oneview_logical_interconnect_group" "logical_interconnect_group" {
        name = "TestLIG4"
}

output "lig_value" {
        value = "${data.oneview_logical_interconnect_group.logical_interconnect_group.redundancy_type}"
}
*/

/*
# Importing an existing resource from appliance
resource "oneview_logical_interconnect_group" "import_lig"{
}
*/
