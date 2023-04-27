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
  hardware_name         = "<server_hardware_name>"
  type                  = "ServerProfileV12"
  server_hardware_type  = "<server_hardware_type_terraform>"
  enclosure_group       = "<enclosure_group_name>"
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

/*
# Patch request - Server profile Refresh
resource "oneview_server_profile" "SP" {
        update_type = "patch"
        options = [
        {
          op = "replace"
          path = "/refreshState"
          value = "RefreshPending"
        }
        ]
        name = "TestSP_Renamed"
        type = "ServerProfileV12"        
        enclosure_group = "EG-Synergy-Local"
        hardware_name = "Synergy-Encl-2, bay 8"
}
*/
