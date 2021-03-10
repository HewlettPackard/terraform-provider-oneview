provider "oneview" {
  ov_username =   "${var.username}"
  ov_password =   "${var.password}"
  ov_endpoint =   "${var.endpoint}"
  ov_sslverify =  "${var.ssl_enabled}"
  ov_apiversion = 2600
  ov_ifmatch = "*"
}





# Creates server profile template
resource "oneview_server_profile_template" "ServerProfileTemplate" {
        name = "TestServerProfileTemplate"
        type = "ServerProfileTemplateV8"
//        enclosure_group = "Auto-EG"
	enclosure_group = "EG"
        server_hardware_type = "SY 480 Gen9 2"
//        initial_scope_uris = ["${data.oneview_scope.scope.uri}"]
        bios_option = {
      		manage_bios = "true"	
		overridden_settings = [
		{
			id =  "TimeFormat"
			value = "Utc"
		}]
	}
	boot_order = ["HardDisk"]
	boot_mode = {
		manage_mode = "true"
		mode = "UEFIOptimized"
		pxe_boot_policy = "Auto"
	}
	local_storage = {
		controller = [{
			device_slot =  "Embedded"
			drive_write_cache = "Unmanaged"
			initialize = "true"
			mode = "RAID"
			predictive_spare_rebuild = "Unmanaged"
			logical_drives = [{
				accelerator = "Unmanaged"
				bootable = "true"
				drive_technology = "SasHdd"
				name = "TestLd"
				num_physical_drives = 2
				raid_level = "RAID1"
			}]
		}]
	}
	connection_settings {
		manage_connections = true
		connections {
			id            = 1
			name          = "Management"
			function_type = "Ethernet"
			network_uri   = "/rest/ethernet-networks/42d5c77c-970f-46dc-b5b6-9835fc4cadc2"
			port_id       = "Mezz 3:1-a"
			boot {
				priority           = "Primary"
				ethernet_boot_type = "PXE"
			}
		}
	}
}
