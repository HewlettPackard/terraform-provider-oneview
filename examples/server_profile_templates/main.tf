provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 3000
  ov_ifmatch    = "*"
}

data "oneview_ethernet_network" "ethernetNetwork" {
  name = "mgmt_untagged"
}

data "oneview_ethernet_network" "ethernetnetworks2" {
  name = "iscsi_nw"
}

data "oneview_fc_network" "fcNetwork" {
  name = "FC_FA"
}

data "oneview_scope" "scope" {
  name = "test"
}

# Creates server profile template with connections
resource "oneview_server_profile_template" "ServerProfileTemplateWithConnections" {
  name                 = "TestServerProfileTemplate_with_connections"
  type                 = "ServerProfileTemplateV8"
  enclosure_group      = "EG"
  server_hardware_type = "SY 480 Gen9 1"
  initial_scope_uris   = [data.oneview_scope.scope.uri]
  bios_option {
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
  }

  connection_settings {
    compliance_control = "CheckedMinimum"
    manage_connections = true
    connections {
      id             = 134
      name           = "Management-01"
      managed        = true
      function_type  = "Ethernet"
      network_uri    = data.oneview_ethernet_network.
        .uri
      isolated_trunk = false 
      port_id        = "Auto" 
      requested_mbps = "2500"
      boot {
        priority           = "Primary"
        ethernet_boot_type = "PXE"
      }
    }
    connections {
      id             = 135
      name           = "Management-02"
      managed        = true
      function_type  = "FibreChannel"
      network_uri    = data.oneview_fc_network.fcNetwork.uri
      isolated_trunk = false  
      requested_mbps = "2500"
      port_id        = "Auto"
      requested_vfs  = ""
      boot {
        priority           = "Primary"
	      boot_volume_source = "UserDefined"
      }
    }
  }
}
  
/*
# Creates server profile template with local storage
resource "oneview_server_profile_template" "ServerProfileTemplateWithLocalStorage" {
  name                 = "TestServerProfileTemplate_with_local_storage"
  type                 = "ServerProfileTemplateV8"
  enclosure_group      = "EG"
  server_hardware_type = "SY 480 Gen9 1"
  initial_scope_uris   = [data.oneview_scope.scope.uri]
  bios_option {
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
  }
  local_storage {
    controller {
      device_slot       = "Embedded"
      drive_write_cache = "Unmanaged"
      initialize        = true
      import_configuration = false 
      mode                     = "Mixed"
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
  }
}
*/
  
/*
# Creates server profile template with OS deployment settings
resource "oneview_server_profile_template" "ServerProfileTemplateWithOSDS" {
  name                 = "TestServerProfileTemplate_with_osds"
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
  boot {
    manage_boot         = true
    boot_order          = ["HardDisk"]
  }
  boot_mode {
    manage_mode     = true
    mode            = "UEFIOptimized"
    pxe_boot_policy = "Auto"
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
    compliance_control = "CheckedMinimum"
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
      network_uri    = data.oneview_ethernet_network.etherneNetwork.uri
      port_id        = "Auto"
      requested_mbps = "2500"
    }
  }
}
*/
  
/*
# Creates server profile template with san storage
resource "oneview_server_profile_template" "ServerProfileTemplateWithSanStorage" {
  name                 = "TestServerProfileTemplate_with_local_storage_san"
  type                 = "ServerProfileTemplateV8"
  enclosure_group      = "EG"
  server_hardware_type = "SY 480 Gen9 2"
  initial_scope_uris   = [data.oneview_scope.scope.uri]
  bios_option {
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
    secure_boot = "Unmanaged"
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
        name                = "TestLD-01"
        num_physical_drives = 2
        raid_level          = "RAID1"
      }
    }
  }
  connection_settings {
    manage_connections = true
    compliance_control = "CheckedMinimum"
    connections {
      id             = 1
      name           = "connection1"
      function_type  = "FibreChannel"
      port_id        = "Mezz 3:1-b"
      requested_mbps = 2500
      network_uri    = data.oneview_fc_network.fcNetwork.uri
    }
    connections {
      id             = 2
      name           = "Management-01"
      isolated_trunk = false     
      function_type  = "Ethernet"
      network_uri    = data.oneview_ethernet_network.ethernetNetwork.uri
      port_id        = "Auto"
      requested_mbps = "2500"
      boot {
        priority           = "Primary"
        ethernet_boot_type = "PXE"
      }
      
    }
    
  }
    san_storage {
    host_os_type       = "Windows 2012 / WS2012 R2"
    manage_san_storage = true
  }
  volume_attachments {
    boot_volume_priority = "NotBootable"
    id                   = 1    
    lun_type                  = "Auto"
    volume_storage_system_uri = "/rest/storage-systems/TXQ1000307"
    volume_uri                = "/rest/storage-volumes/9E1C08D5-8EDF-4600-A73E-AD3A00B1ACB7"
    storage_paths {
      connection_id   = 1
      network_uri     = "/rest/fc-networks/37aae264-8fd5-4624-960d-10173bde5752"
      target_selector = "Auto"
      is_enabled      = true
    }
  }
}
*/
  
# Creating Server Profile Template with DL Server
# Enclosure group and affinity are not supported for DL server
/*
resource "oneview_server_profile_template" "ServerProfileTemplateWithDLServer" {
  name                 = "TestSPT_DL_Server"
  server_hardware_type = "DL560 Gen10 1"
  boot_mode {
    manage_mode     = true
    mode            = "UEFIOptimized"
    pxe_boot_policy = "Auto"
  }
}
*/