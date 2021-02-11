provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2400
  ov_ifmatch    = "*"
}

data "oneview_scope" "scope" {
  name = "testing"
}
/*
# Creation of Server Profile without template
resource "oneview_server_profile" "SP" {
  name               = "TestSP2"
  hardware_name      = "0000A66101, bay 3"
  type               = "ServerProfileV12"
  enclosure_group    = "EG"
  bios_option {
    manage_bios = "true"
    overridden_settings {
      id    = "TimeFormat"
      value = "Utc"
    }
  }
  boot_order = ["HardDisk"]
  boot_mode {
    manage_mode     = "true"
    mode            = "UEFIOptimized"
    pxe_boot_policy = "Auto"
  }
  local_storage {
    controller {
      device_slot              = "Embedded"
      drive_write_cache        = "Unmanaged"
      initialize               = "true"
      mode                     = "RAID"
      predictive_spare_rebuild = "Unmanaged"
      logical_drives {
        accelerator         = "Unmanaged"
        bootable            = "true"
        drive_technology    = "SasHdd"
        name                = "TestLd"
        num_physical_drives = 2
        raid_level          = "RAID1"
      }
    }
  }
  initial_scope_uris = [data.oneview_scope.scope.uri]
}
*/

# Updating Server profile
resource "oneview_server_profile" "SP" {
  name                 = "TestSP_Renamed"
  hardware_name        = "0000A66101, bay 3"
  type                 = "ServerProfileV12"
  enclosure_group      = "EG"
  server_hardware_type = "SY 660 Gen9 1"
  bios_option {
    manage_bios = "true"
    overridden_settings {
      id    = "TimeFormat"
      value = "Utc"
    }
  } 
  local_storage {
    controller {
      device_slot              = "Embedded"
      drive_write_cache        = "Unmanaged"
      initialize               = "true"
      mode                     = "RAID"
      predictive_spare_rebuild = "Unmanaged"
      logical_drives {
        accelerator         = "Unmanaged"
        bootable            = "false"
        drive_technology    = "SasHdd"
        name                = "TestLdUpdate"
        num_physical_drives = 2
        raid_level          = "RAID1"
      }
    }
  }
  initial_scope_uris   = [data.oneview_scope.scope.uri]
  update_type          = "put"
}

/*
# Patch request - Server profile Refresh
resource "oneview_server_profile" "SP" {
  update_type = "patch"
  options {
    op    = "replace"
    path  = "/refreshState"
    value = "RefreshPending"
  }
  name                 = "TestSP_Renamed"
  type                 = "ServerProfileV12"
  server_hardware_type = "SY 660 Gen9 2"
  enclosure_group      = "test_EG"
  hardware_name        = "0000A66101, bay 3"
}

#Data source for server profile

data "oneview_server_profile" "sp" {
  name = "TestSP_Renamed"
}

output "oneview_server_profile_value" {
  value = data.oneview_server_profile.sp.uri
}

# To import an existing server profile to terraform, use the below code and run the following command:

# terraform import <resource>.<instance_name> <resource_name>
# Eg: terraform import oneview_server_profile.serverProfile Test
resource "oneview_server_profile" "serverProfile" {
}
*/
