provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2600
  ov_ifmatch    = "*"
}

data "oneview_scope" "scope" {
  name = "testing"
}

resource "oneview_server_profile_template" "ServerProfileTemplate" {
  name                 = "TestServerProfileTemplate101"
  type                 = "ServerProfileTemplateV8"
  enclosure_group      = "EG"
  server_hardware_type = "SY 480 Gen9 1"
  initial_scope_uris   = [data.oneview_scope.scope.uri]
  boot_order           = ["HardDisk"]
  boot_mode {
    manage_mode     = true
    mode            = "UEFIOptimized"
    pxe_boot_policy = "Auto"
  }
  bios_option {
    manage_bios = true
    overridden_settings {
      id    = "TimeFormat"
      value = "Utc"
    }
  }
  local_storage {
    controller {
      device_slot              = "Embedded"
      drive_write_cache        = "Unmanaged"
      initialize               = true
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
    manage_connections = true
    connections {
      id            = 1
      name          = "Deployment Network A"
      function_type = "Ethernet"
      network_uri   = "/rest/ethernet-networks/728af64d-c9aa-4287-a331-ba4c6654dd15"
      port_id       = "Mezz 3:1-a"
      boot {
        priority           = "Primary"
        ethernet_boot_type = "PXE"
      }
    }
  }
}

resource "oneview_server_profile_template" "ServerProfileTemplate" {
  name                 = "TestServerProfileTemplate102"
  type                 = "ServerProfileTemplateV8"
  enclosure_group      = "EG"
  server_hardware_type = "SY 480 Gen9 1"
  initial_scope_uris   = [data.oneview_scope.scope.uri]
  local_storage {
    controller {
      device_slot              = "Embedded"
      drive_write_cache        = "Unmanaged"
      initialize               = true
      mode                     = "RAID"
      predictive_spare_rebuild = "Unmanaged"
      logical_drives {
        accelerator         = "Unmanaged"
        bootable            = false
        drive_technology    = "SasHdd"
        name                = "TestLd"
        num_physical_drives = 2
        raid_level          = "RAID1"
      }
    }
  }
  connection_settings {
    manage_connections = true
    connections {
      id            = 2
      name          = "Deployment Network B"
      function_type = "Ethernet"
      network_uri   = "/rest/ethernet-networks/6afb79a2-44be-4055-aa56-1c632149d386"
      port_id       = "Mezz 3:2-a"
    }
  }
}

/* 	Update */
resource "oneview_server_profile_template" "ServerProfileTemplate" {
  name                 = "TestServerProfileTemplate_Renamed"
  type                 = "ServerProfileTemplateV8"
  enclosure_group      = "enclosureGp"
  server_hardware_type = "SY 480 Gen9 1"
  initial_scope_uris   = [data.oneview_scope.scope.uri]
}

/* 	Datasource */
data "oneview_server_profile_template" "server_profile_template" {
  name = "TestServerProfileTemplate"
}

output "oneiew_server_hardware_type_value" {
  value = data.oneview_server_profile_template.server_profile_template.uri
}

