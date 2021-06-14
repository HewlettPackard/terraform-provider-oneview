provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2800
  ov_ifmatch    = "*"
}

# Testing data source
data "oneview_ethernet_network" "ethernetnetworks" {
  name = "TestNetwork_1"
}

data "oneview_scope" "scope" {
  name = "Auto-Scope"
}

# Creation of Server Profile without template
resource "oneview_server_profile" "SP" {
  name                 = "TestSP"
  hardware_name        = "0000A66102, bay 5"
  type                 = "ServerProfileV12"
  enclosure_group      = "Auto-EG"
  initial_scope_uris = [data.oneview_scope.scope.uri]
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
  os_deployment_settings {
    compliance_control      = "Checked"
    os_deployment_plan_name = "Simple Deployment-HPE-Support 1.0"
    os_custom_attributes {
      name  = "ExampleOption"
      value = "Option1"
    }
    os_custom_attributes {
      name  = "ExampleNumber"
      value = "50"
    }
    os_custom_attributes {
      name  = "ExampleString"
      value = "Hello World!"
    }
  }
  connection_settings {
    manage_connections = true
    connections {
      id             = 1
      name           = "Deployment Network A"
      isolated_trunk = false
      managed        = true
      function_type  = "Ethernet"
      network_uri    = "/rest/ethernet-networks/8bd7a0e0-75e8-40c7-9f52-d08b007e5270"
      port_id        = "Mezz 3:1-a"
      requested_mbps = "2500"
      boot {
        boot_volume_source = "UserDefined"
        priority           = "Primary"
        ethernet_boot_type = "iSCSI"
      }
      ipv4 {
        ip_address_source  = "SubnetPool"
      }
    }
    connections {
      id             = 2
      name           = "Deployment Network B"
      isolated_trunk = false
      managed        = true
      function_type  = "Ethernet"
      network_uri    = "/rest/ethernet-networks/8bd7a0e0-75e8-40c7-9f52-d08b007e5270"
      port_id        = "Mezz 3:2-a"
      requested_mbps = "2500"
      boot {
        boot_volume_source = "UserDefined"
        priority           = "Secondary"
        ethernet_boot_type = "iSCSI"
      }
      ipv4 {
        ip_address_source  = "SubnetPool"
      }
    }
    connections {
      id             = 3
      name           = "Management-01"
      isolated_trunk = false
      managed        = true
      function_type  = "Ethernet"
      network_uri    = data.oneview_ethernet_network.ethernetnetworks.uri
      port_id        = "Auto"
      requested_mbps = "2500"
#      boot {
#        priority           = "Primary"
#        ethernet_boot_type = "PXE"
#      }
    }
  }
}
