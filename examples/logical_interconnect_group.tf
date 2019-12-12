provider "oneview" {
        ov_username = "<ov-username>"
        ov_password = "<ov-password>"
        ov_endpoint = "<ov-endpoint>"
        ov_sslverify = false
        ov_apiversion = <ov-apiversion>
        ov_ifmatch = "*"
}

# Creates a logical interconnect group or updates if already existing
resource "oneview_logical_interconnect_group" "LIG" {
  name = "TestLIG"
  type = "logical-interconnect-groupV7"
  interconnect_bay_set = 1
  redundancy_type = "NonRedundantASide"
  enclosure_indexes = [-1]

  interconnect_map_entry_template = [
    {
      bay_number             = 1
      interconnect_type_name = "Virtual Connect SE 16Gb FC Module for Synergy"
      enclosure_index = -1
    },
  ]
}

/* Test for data source

data "oneview_logical_interconnect_group" "logical_interconnect_group" {
        name = "TestLIG"
}

output "lig_value" {
        value = "${data.oneview_logical_interconnect_group.logical_interconnect_group.redundancy_type}"
}*/
