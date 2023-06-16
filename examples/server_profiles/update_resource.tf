provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}


data "oneview_ethernet_network" "ethernetnetworks1" {
  name = "Server_profile_network"
}

data "oneview_scope" "scope" {
  name = "Auto-Scope"
}
# Creation of Server Profile without template
resource "oneview_server_profile" "SPWithLocalStorage" {
  name                  = "TestSP_with_local_storage_renamed"
  hardware_name         = "<server_hardware_name_terra>"
  type                  = "ServerProfileV12"
  enclosure_group       = "<enclosure_group_name>"
  update_type           = "put"
  initial_scope_uris    = [data.oneview_scope.scope.uri]
  bios_option {
    manage_bios = true
    overridden_settings {
      id    = "TimeFormat"
      value = "Utc"
    }
  }
  boot {
    manage_boot = true
    boot_order  = ["HardDisk"]
  }
  boot_mode {
    manage_mode     = true
    mode            = "UEFIOptimized"
    pxe_boot_policy = "Auto"
  }
  local_storage {
    controller {
      device_slot              = "Embedded"
      drive_write_cache        = "Unmanaged"
      initialize               = true
      import_configuration     = false
      mode                     = "RAID"
      predictive_spare_rebuild = "Unmanaged"
      logical_drives {
        accelerator         = "Unmanaged"
        bootable            = true
        drive_technology    = "SasHdd"
        name                = "TestLd"
        num_physical_drives = 2
        raid_level          = "RAID1"
      }
    }
  }
  connection_settings {
    connections {
      id             = 1
      name           = "Management-01"
      isolated_trunk = false
      # managed        = true
      function_type  = "Ethernet"
      network_uri    = data.oneview_ethernet_network.ethernetnetworks1.uri
      port_id        = "Auto"
      requested_mbps = "2500"
      boot {
        priority           = "Primary"
        ethernet_boot_type = "PXE"
      }
    }
  }
}


# Patch request to reapply  the server profile's SAN storage,  local storage and rename the logical drive
/*
resource "oneview_server_profile" "SP" {        
 
  name = "SP"
  type = "ServerProfileV12"        
  enclosure_group = "EG"
  hardware_name = "0000A66101, bay 3"

 operation_type {
          name = "reapply_SAN_storage"        
          
        }
  operation_type {
          name = "reapply_local_storage"        
          
        }  
  operation_type {
          name = "rename_logical_drive"
          logical_drive_value{
            device_slot_name="Embedded"
            current_logical_drive_name="ld2"
            new_logical_drive_name="ld3"
          }        
          
        }         
}
*/




