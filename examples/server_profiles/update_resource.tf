provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Creation of Server Profile without template
resource "oneview_server_profile" "SP" {
  name            = "Test-SP"
  hardware_name   = "0000A66102, bay 3"
  type            = "ServerProfileV12"
  enclosure_group = "EG"
  update_type     = "put"
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
      initialize               = false
      import_configuration     = false
      mode                     = "RAID"
      predictive_spare_rebuild = "Unmanaged"

      /* To remove first logical drive we are kepping logical_drives as an empty block.
   Note: Do not remove the block, you will need to keep it as empty in order to delete it from the Oneview. */

      logical_drives {
      }

      logical_drives {
        accelerator         = "Unmanaged"
        bootable            = false
        drive_technology    = "SasHdd"
        drive_number        = 2
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
      mac_type       = "Virtual"
      function_type  = "Ethernet"
      network_uri    = "/rest/ethernet-networks/1aa324ee-5e86-4428-a1cd-97ed26020d9d"
      port_id        = "Mezz 3:1-a"
      requested_mbps = "2500"
      boot {
        boot_vlan_id       = 0
        priority           = "Primary"
        ethernet_boot_type = "PXE"
      }
    }

    /* To remove second connection we are kepping connection as an empty block.
   Note: Do not omit the block, you will need to keep it as empty in order to delete it from the Oneview. */

    connections {
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
