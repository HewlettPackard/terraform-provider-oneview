provider "oneview" {
  ov_username =   "${var.username}"
  ov_password =   "${var.password}"
  ov_endpoint =   "${var.endpoint}"
  ov_sslverify =  "${var.ssl_enabled}"
  ov_apiversion = 2400
  ov_ifmatch = "*"
}

 
# Creates a logical interconnect group or updates if already existing
/*
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
  port_flap_settings = [
    {
      type = "portFlapProtection"
      port_flap_protection_mode = "Detect"
      port_flap_threshold_per_interval =  2
      detection_interval = 20
      no_of_samples_declare_failures = 3
      name = "PortFlapSettings"
      consistency_checking = "ExactMatch"
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
        name = "TestLIG5"
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
