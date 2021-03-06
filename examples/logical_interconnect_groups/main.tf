provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2800
  ov_ifmatch    = "*"
}

# Create Logical Interconnect Group
resource "oneview_logical_interconnect_group" "logical_interconnect_group" {
  type                 = "logical-interconnect-groupV8"
  name                 = "Auto-LIG"
  interconnect_bay_set = 3
  enclosure_indexes    = [1, 2, 3]
  redundancy_type      = "HighlyAvailable"
  interconnect_map_entry_template {
    enclosure_index        = 1
    bay_number             = 3
    interconnect_type_name = "Virtual Connect SE 40Gb F8 Module for Synergy"
  }
  interconnect_map_entry_template {
    enclosure_index        = 2
    bay_number             = 6
    interconnect_type_name = "Virtual Connect SE 40Gb F8 Module for Synergy"
  }
  interconnect_map_entry_template {
    enclosure_index        = 2
    bay_number             = 3
    interconnect_type_name = "Synergy 20Gb Interconnect Link Module"
  }
  interconnect_map_entry_template {
    enclosure_index        = 3
    bay_number             = 3
    interconnect_type_name = "Synergy 20Gb Interconnect Link Module"
  }
  interconnect_map_entry_template {
    enclosure_index        = 1
    bay_number             = 6
    interconnect_type_name = "Synergy 20Gb Interconnect Link Module"
  }
  interconnect_map_entry_template {
    enclosure_index        = 3
    bay_number             = 6
    interconnect_type_name = "Synergy 20Gb Interconnect Link Module"
  }
  igmp_settings {
    consistency_checking       = "ExactMatch"
    igmp_idle_timeout_interval = 260
    igmp_snooping              = true
    prevent_flooding           = true
    proxy_reporting            = true
  }
  port_flap_settings {
    type                             = "portFlapProtection"
    port_flap_protection_mode        = "Detect"
    port_flap_threshold_per_interval = 2
    detection_interval               = 20
    no_of_samples_declare_failures   = 3
    name                             = "PortFlapSettings"
    consistency_checking             = "ExactMatch"
  }
}

