provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

data "oneview_scope" "scope" {
  name = "<scope>"
}

data "oneview_ethernet_network""ethernetnetworks1"{
  name = "<network_1>"
}

# Updates server profile template from main.tf
resource "oneview_server_profile_template" "ServerProfileTemplateWithLocalStorage" {
  name                 = "<spt_rename>"
  type                 = "<type_terraform>"
  enclosure_group      = "<enclosure_group_name_terraform>"
  server_hardware_type = "<server_hardware_type_name_terraform>"
  initial_scope_uris   = [data.oneview_scope.scope.uri]
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
        name                = "TestLD-01"
        num_physical_drives = 2
        raid_level          = "RAID1"
      }
    }
    /* To remove second connection we are kepping connection as an empty block.
   Note: Do not omit the block, you will need to keep it as empty in order to delete it from the Oneview. */

  }
}
