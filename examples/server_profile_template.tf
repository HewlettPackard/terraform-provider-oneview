provider "oneview" {
  ov_username =   "${var.username}"
  ov_password =   "${var.password}"
  ov_endpoint =   "${var.endpoint}"
  ov_sslverify =  "${var.ssl_enabled}"
  ov_apiversion = 2400
  ov_ifmatch = "*"
}

data "oneview_scope" "scope" {
        name = "testing"
}
/*
resource "oneview_server_profile_template" "ServerProfileTemplate" {
        name = "TestServerProfileTemplate"
        type = "ServerProfileTemplateV8"
        enclosure_group = "EG-Synergy-Local"
        server_hardware_type = "SY 480 Gen9 1"
        initial_scope_uris = ["${data.oneview_scope.scope.uri}"]
	boot_order = ["HardDisk"]

	boot_mode = {
		manage_mode = true
		mode = "UEFIOptimized"
		pxe_boot_policy = "Auto"
	}
	local_storage = {
		controller = [{
			device_slot =  "Embedded",
			drive_write_cache =  "Unmanaged",
			initialize = true,
			logical_drives = [{
				accelerator = "Unmanaged",
				bootable = true,
				drive_technology =  "SasHdd",
				name =  "TestLd",
				num_physical_drives =  2,
				raid_level =  "RAID1",
			}]
		}]
	}
}
*/
/* 	Update 
resource "oneview_server_profile_template" "ServerProfileTemplate" {
	name = "TestServerProfileTemplate_Renamed"
	type = "ServerProfileTemplateV8"
	enclosure_group = "enclosureGp"
	server_hardware_type = "SY 480 Gen9 1"
	initial_scope_uris = ["${data.oneview_scope.scope.uri}"]
}
*/
/* 	Datasource 
data "oneview_server_profile_template" "server_profile_template" {
	name = "TestServerProfileTemplate"
}

output "oneiew_server_hardware_type_value" {
	value = "${data.oneview_server_profile_template.server_profile_template.uri}"
}
*/
