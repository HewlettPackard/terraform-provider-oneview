provider "oneview" {
  ov_username =   "${var.username}"
  ov_password =   "${var.password}"
  ov_endpoint =   "${var.endpoint}"
  ov_sslverify =  "${var.ssl_enabled}"
  ov_apiversion = 2600
  ov_ifmatch = "*"
}

data "oneview_scope" "scope" {
        name = "Auto-Scope"
}

# Creates server profile template
resource "oneview_server_profile_template" "ServerProfileTemplate" {
        name = "TestServerProfileTemplate"
        type = "ServerProfileTemplateV8"
        enclosure_group = "Auto-EG"
        server_hardware_type = "SY 480 Gen9 1"
        initial_scope_uris = ["${data.oneview_scope.scope.uri}"]
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
}

data "oneview_scope" "scope" {
        name = "Auto-Scope"
}

# Updates server profile template from main.tf
resource "oneview_server_profile_template" "ServerProfileTemplate" {
        name = "TestServerProfileTemplateRenamed"
        type = "ServerProfileTemplateV8"
        enclosure_group = "Auto-EG"
        server_hardware_type = "SY 480 Gen9 1"
        initial_scope_uris = ["${data.oneview_scope.scope.uri}"]
}

# Datasource for server profile template
data "oneview_server_profile_template" "server_profile_template" {
	name = "TestServerProfileTemplateRenamed"
}

output "oneiew_server_hardware_type_value" {
	value = "${data.oneview_server_profile_template.server_profile_template.uri}"
}

