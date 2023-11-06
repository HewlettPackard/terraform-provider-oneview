provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
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
  name               = "TestSP_with_local_storage"
  hardware_name      = "0000A66101, bay 5"
  type               = "ServerProfileV12"
  enclosure_group    = "Auto-EG"

  force_flags = ["ignoreServerHealth"] // supported: ignoreSANWarnings, ignoreServerHealth, ignoreLSWarnings, all; default: none

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
        name                = "TestLd"
        num_physical_drives = 2
        raid_level          = "RAID1"
      }
    }
  }
  connection_settings {
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
#Create a Server Profile using a Server profile template

# data "oneview_server_hardware" "sh" {
#   name = "172.18.31.4"
# }

# data "oneview_server_profile_template" "spt" {
#   name = "spt"
# }


# resource "oneview_server_profile" "sp_from_spt" {
#   name                 = "sp_from_spt"
#   hardware_name        = data.oneview_server_hardware.sh.name
#   server_hardware_type = data.oneview_server_profile_template.spt.server_hardware_type
#   template             = data.oneview_server_profile_template.spt.name
#   scopes_uri           = data.oneview_scope.scope.uri
# }

/*
# Create Server Profile with OS deployment settings
resource "oneview_server_profile" "SPWithOSDS" {
  name               = "TestSP_with_osds"
  hardware_name      = "0000A66102, bay 5"
  type               = "ServerProfileV12"
  enclosure_group    = "Auto-EG"
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
    force_os_deployment     = false
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
        ip_address_source = "SubnetPool"
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
        ip_address_source = "SubnetPool"
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
      initialize               = false
      import_configuration     = false
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
      id             = 1
      name           = "Management-01"
      isolated_trunk = false
      mac_type       = "Virtual"
      function_type  = "Ethernet"
      network_uri    = "/rest/ethernet-networks/1aa324ee-5e86-4428-a1cd-97ed26020d9d"
      port_id        = "Mezz 3:1-a"
      requested_mbps = "2500"
      boot {
        boot_vlan_id       = 0
        priority           = "Primary"
        ethernet_boot_type = "PXE"
      }
    }
    connections {
      id             = 2
      name           = "Management-02-Updatejhgj"
      isolated_trunk = false
      function_type  = "iSCSI"
      network_uri    = "/rest/ethernet-networks/8b875382-04d0-4027-9263-fdb95b1c15ef"
      port_id        = "Mezz 3:1-b"
      requested_mbps = "2500"
      mac_type       = "Virtual"
      boot {
        priority = "NotBootable"
      }
    }
  }
}

# Creation of Server Profile on DL server
# Enclosure group and affinity are not supported for DL server
resource "oneview_server_profile" "SPWithDL" {
  name                 = "TestSP_DL_Server"
  hardware_name        = "<ilo_ip>"  
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
}

# Creation of Server Profile with connections, local storage and management processor
resource "oneview_server_profile" "SP" {
  name                 = "SPi101"
  hardware_name        = "0000A66101, bay 1"
  type                 = "ServerProfileV12"
  enclosure_group      = "EG" 
  update_type          = "put"

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
      initialize               = false
      import_configuration     = false
      mode                     = "Mixed"
      predictive_spare_rebuild = "Unmanaged"

      logical_drives {
        accelerator         = "Unmanaged"
        bootable            = false
        drive_technology    = "SasHdd"
        name                = "TestLd"
        num_physical_drives = 2
        raid_level          = "RAID1"
      }

      logical_drives {
        accelerator         = "Unmanaged"
        bootable            = false
        drive_technology    = "SasHdd"
        name                = "TestLd3"
        num_physical_drives = 2
        raid_level          = "RAID1"
      }
    }
  }

  connection_settings {

    connections {
      id             = 1
      name           = "Management-01"
      isolated_trunk = false
      mac_type       = "Virtual"
      function_type  = "Ethernet"
      network_uri    = data.oneview_ethernet_network.ethernetnetwork1.uri // "/rest/ethernet-networks/1aa324ee-5e86-4428-a1cd-97ed26020d9d"
      port_id        = "Auto"
      requested_mbps = "2500"
      boot {
        boot_vlan_id       = 0
        priority           = "Primary"
        ethernet_boot_type = "PXE"
      }
    }

    connections {
      id             = 2
      name           = "Management-02-Update"
      isolated_trunk = false
      function_type  = "iSCSI"
      network_uri    = data.oneview_ethernet_network.ethernetnetwork2.uri //"/rest/ethernet-networks/8b875382-04d0-4027-9263-fdb95b1c15ef"
      port_id        = "Auto"
      requested_mbps = "2500"
      mac_type       = "Virtual"

      boot {
        priority = "NotBootable"
      }
    }
  }

  management_processor {
    manage_mp = true
    mp_settings {

      local_accounts {
        user_name                    = "test_UserNamei-Update"
        display_name                 = "test_DisplayName"
        password                     = "test_password"
        user_config_priv             = true
        remote_console_priv          = true
        virtual_media_priv           = true
        virtual_power_and_reset_priv = true
        ilo_config_priv              = true
      }

      directory {
        directory_authentication     = "defaultSchema"
        directory_generic_ldap       = false
        directory_server_address     = "ldap.example.com"
        directory_server_port        = 636
        directory_server_certificate = "-----BEGIN CERTIFICATE-----\nMIIBozCCAQwCCQCWGqL41Y6YKTANBgkqhkiG9w0BAQUFADAWMRQwEgYDVQQDEwtD\nb21tb24gTmFtZTAeFw0xNzA3MTQxOTQzMjZaFw0xODA3MTQxOTQzMjZaMBYxFDAS\nBgNVBAMTC0NvbW1vbiBOYW1lMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCf\nCNTrU4AZF044Rtu8jiGR6Ce1u9K6GJE+60VCau2y4A2z5B5kKA2XnyP+2JpLRRA8\n8PjEyVJuL1fJomGF74L305j6ucetXZGcEy26XNyKFOtsBeoHtjkISYNTMxikjvC1\nXHctTYds0D6Q6u7igkN9ew8ngn61LInFqb6dLm+CmQIDAQABMA0GCSqGSIb3DQEB\nBQUAA4GBAFVOQ8zXFNHdXVa045onbkx8pgM2zK5VQ69YFtlAymFDWaS7a5M+96JW\n2c3001GDGZcW6fGqW+PEyu3COImRlEhEHaZKs511I7RfckMzZ3s7wPrQrC8WQLqI\ntiZtCWfUX7tto7YDdmfol7bHiaSbrLUv4H/B7iS9FGemA+nrghCK\n-----END CERTIFICATE-----"
        directory_user_context       = ["OU=US,OU=Users,OU=Accounts,dc=Subdomain,dc=example,dc=com", "ou=People,o=example.com"]
        ilo_distinguished_name       = "service"
        password                     = "test_password"
        kerberos_authentication      = false
      }

      directory_groups {
        group_dn                     = "ilos.example.com,ou=Groups,o=example.com"
        group_sid                    = "S-1-5-12"
        user_config_priv             = false
        remote_console_priv          = true
        virtual_media_priv           = true
        virtual_power_and_reset_priv = true
        ilo_config_priv              = true
        ##below attrbutes are meant for gen 10 and above
        login_priv                   = true
        host_bios_config_priv        = true
        host_nic_config_priv         = true
        host_storage_config_priv     = true

      }

      key_manager {
        primary_server_address   = "192.0.2.91"
        primary_server_port      = 9000
        secondary_server_address = "192.0.2.92"
        secondary_server_port    = 9000
        redundancy_required      = true
        group_name               = "GRP"
        certificate_name         = "Local CA"
        login_name               = "deployment"
        password                 = "test_password"
      }

      administrator_account {
        delete_administrator_account = false
        password                     = "test_password"
      }
    }
  }
}
*/
