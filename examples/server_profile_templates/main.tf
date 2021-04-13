provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2600
  ov_ifmatch    = "*"
}

data "oneview_scope" "scope" {
  name = "Auto-Scope"
}

# Creates server profile template
resource "oneview_server_profile_template" "ServerProfileTemplate" {
  name                 = "TestServerProfileTemplater"
  type                 = "ServerProfileTemplateV8"
  enclosure_group      = "Auto-EG"
  server_hardware_type = "SY 480 Gen9 1"
  initial_scope_uris   = [data.oneview_scope.scope.uri]
  bios_option {
    manage_bios = true
    overridden_settings {
      id    = "TimeFormat"
      value = "Utc"
    }
  }
  boot_order = ["HardDisk"]
  boot_mode {
    manage_mode     = true
    mode            = "UEFIOptimized"
    pxe_boot_policy = "Auto"
  }
  local_storage {
    controller {
      device_slot       = "Embedded"
      drive_write_cache = "Unmanaged"
      initialize        = true
      import_configuration = false 
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
/*  connection_settings {
	manage_connections = true
	connections {
		id            = 1
		name          = "Management"
		function_type = "Ethernet"
		network_uri   = "/rest/ethernet-networks/9ada4ca7-5832-4378-b972-a44d7f2ce652"
		port_id       = "Mezz 3:1-a"
		boot {
			priority           = "Primary"
			ethernet_boot_type = "PXE"
		}
	}
  }*/
}
