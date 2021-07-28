provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2800
  ov_ifmatch    = "*"
}

# Testing data source
data "oneview_ethernet_network" "ethernetnetworks1" {
  name = "TestNetwork_1"
}

data "oneview_ethernet_network" "ethernetnetworks2" {
  name = "iscsi_nw"
}

data "oneview_scope" "scope" {
  name = "Auto-Scope"
}

# Create Server Profile with local storage
resource "oneview_server_profile" "SPWithLocalStorage" {
  name                 = "TestSP_with_local_storage"
  hardware_name        = "0000A66101, bay 5"
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
  connection_settings {
    manage_connections = true
    connections {
      id             = 1
      name           = "Management-01"
      isolated_trunk = false
      managed        = true
      function_type  = "Ethernet"
      network_uri    = data.oneview_ethernet_network.ethernetnetworks1.uri
      port_id        = "Auto"
      requested_mbps = "2500"
      boot {
        priority           = "Primary"
        ethernet_boot_type = "PXE"
      }
    }
  }
}

# Create Server Profile with OS deployment settings
resource "oneview_server_profile" "SPWithOSDS" {
  name                 = "TestSP_with_osds"
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
  os_deployment_settings {
    force_os_deployment = false
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
      network_uri    = data.oneview_ethernet_network.ethernetnetworks2.uri
      port_id        = "Mezz 3:1-a"
      requested_mbps = "2500"
      boot {
        boot_volume_source = "UserDefined"
        priority           = "Primary"
        ethernet_boot_type = "iSCSI"
        iscsi {
          initiator_name_source = "ProfileInitiatorName"
        }
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
      network_uri    = data.oneview_ethernet_network.ethernetnetworks2.uri
      port_id        = "Mezz 3:2-a"
      requested_mbps = "2500"
      boot {
        boot_volume_source = "UserDefined"
        priority           = "Secondary"
        ethernet_boot_type = "iSCSI"
        iscsi {
          initiator_name_source = "ProfileInitiatorName"
        }
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
      network_uri    = data.oneview_ethernet_network.ethernetnetworks1.uri
      port_id        = "Auto"
      requested_mbps = "2500"
    }
  }
}

# Creation of Server Profile with logical storage and connections
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

# Creation of Server Profile on DL server
resource "oneview_server_profile" "SPWithDL" {
  name                 = "TestSP_DL_Server"
  hardware_name        = "<ilo_ip>"
  server_hardware_type  = "DL360 Gen10 1"
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
  boot_mode {
    manage_mode     = true
    mode            = "UEFIOptimized"
    pxe_boot_policy = "Auto"
  }
}

