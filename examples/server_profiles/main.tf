provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2800
  ov_ifmatch    = "*"
}

# Creation of Server Profile without template
resource "oneview_server_profile" "SP" {
  name                 = "Test-SP"
  hardware_name        = "0000A66102, bay 3"
  type                 = "ServerProfileV12"
  enclosure_group      = "EG"
  server_hardware_type  = "SY 660 Gen9 1" 
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
      device_slot       = "Embedded"
      drive_write_cache = "Unmanaged"
      initialize        = false
      import_configuration = false
      mode                     = "RAID"
      predictive_spare_rebuild = "Unmanaged"
      logical_drives {
        accelerator         = "Unmanaged"
        bootable            = false
        drive_technology    = "SasHdd"
        name                = "TestLd1"
        num_physical_drives = 2
        raid_level          = "RAID1"
      }
      logical_drives {
        accelerator         = "Unmanaged"
        bootable            = false
        drive_technology    = "SasHdd"
        name                = "TestLd2"
        num_physical_drives = 2
        raid_level          = "RAID1"
      }
    }
  }
  connection_settings {
    connections {
      id            = 1
      name          = "Management-01"
      isolated_trunk = false
      mac_type          = "Virtual"
      function_type = "Ethernet"
      network_uri   = "/rest/ethernet-networks/1aa324ee-5e86-4428-a1cd-97ed26020d9d"
      port_id       = "Mezz 3:1-a"
      requested_mbps = "2500"
      boot {
        boot_vlan_id       = 0
        priority           = "Primary"
        ethernet_boot_type = "PXE"
      }
    }
    connections {
      id            = 2
      name          = "Management-02-Updatejhgj"
      isolated_trunk = false  
      function_type = "iSCSI"
      network_uri   = "/rest/ethernet-networks/8b875382-04d0-4027-9263-fdb95b1c15ef"
      port_id       = "Mezz 3:1-b"
      requested_mbps = "2500"
      mac_type       = "Virtual"
      boot {
	priority = "NotBootable"
      }
    }                                                                                 
  }
}

