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
  name               = "TestSP2"
  hardware_name      = "0000A66101, bay 3"
  server_hardware_type  = "SY 660 Gen9 2"
  type               = "ServerProfileV12"
  enclosure_group      = "EG"
  update_type          = "put"
  bios_option {
    consistency_state = "Consistent"
    reapply_state     = "NotApplying"
    manage_bios = true
    overridden_settings {
      id    = "TimeFormat"
      value = "Utc"
    }
  }
  boot {
    manage_boot		= true
    boot_order		= ["HardDisk"]
  }
  boot_mode {
    manage_mode     = true
    mode            = "UEFIOptimized"
    pxe_boot_policy = "Auto"
  secure_boot	    = "Unmanaged"
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
        bootable            = true
        drive_technology	    = "SasHdd"
	name                = "TestLD-01"
        num_physical_drives = 2
        raid_level          = "RAID1"
      }
    }
  }
  connection_settings {
    manage_connections = false 
    connections {
      id            = 1
      name          = "FC-01"
      isolated_trunk = false
      managed        = false 
      mac_type          = "Physical"
      function_type = "FibreChannel"
      network_uri   = "/rest/fc-networks/ab5e7b7e-2038-4899-9863-a524c452b549"
      port_id       = "Mezz 2:1"
    }
    connections {
      id            = 2
      allocated_mbps = 0
      allocated_vfs  = 0
      name          = "Eth-01"
      isolated_trunk = false
      managed        = true 
      mac_type          = "Virtual"
      maximum_mbps   = 0
      function_type = "Ethernet"
      network_uri   = "/rest/ethernet-networks/8e90d988-d5bd-4082-a17b-945aa2e8e4c0"
      port_id       = "Mezz 3:1-a"
      requested_mbps = 2500
      boot {
         boot_vlan_id       = 0
         priority           = "Primary"
        ethernet_boot_type = "PXE"
      }
    }
  }
}
