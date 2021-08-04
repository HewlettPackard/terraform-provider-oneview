provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 3000
  ov_ifmatch    = "*"
}

data "oneview_ethernet_network" "ethernetnetworks1" {
  name = "TestNetwork_1"
}

data "oneview_ethernet_network" "ethernetnetworks2" {
  name = "iscsi_nw"
}

data "oneview_scope" "scope" {
  name = "Auto-Scope"
}

# Creating Server Profile Template with Management Settings
resource "oneview_server_profile_template" "SPTwithMgmtSettings" {
  name                 = "TestServerProfileTemplate_with_local_storage"
  type                 = "ServerProfileTemplateV8"
  enclosure_group      = "Auto-EG"
  server_hardware_type = "SY 480 Gen9 1"
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

  management_processor{
    manage_mp	=	true
    mp_settings {
      local_accounts {
        user_name = "test_UserNamei-Update"
        display_name = "test_DisplayName"
        password = "test_password"
        user_config_priv = true
        remote_console_priv = true
        virtual_media_priv = true
        virtual_power_and_reset_priv = true
        ilo_config_priv = true
      }
      directory {
	directory_authentication = "defaultSchema"
	directory_generic_ldap = false
	directory_server_address = "ldap.example.com"
	directory_server_port	= 636
	directory_server_certificate = "-----BEGIN CERTIFICATE-----\nMIIBozCCAQwCCQCWGqL41Y6YKTANBgkqhkiG9w0BAQUFADAWMRQwEgYDVQQDEwtD\nb21tb24gTmFtZTAeFw0xNzA3MTQxOTQzMjZaFw0xODA3MTQxOTQzMjZaMBYxFDAS\nBgNVBAMTC0NvbW1vbiBOYW1lMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCf\nCNTrU4AZF044Rtu8jiGR6Ce1u9K6GJE+60VCau2y4A2z5B5kKA2XnyP+2JpLRRA8\n8PjEyVJuL1fJomGF74L305j6ucetXZGcEy26XNyKFOtsBeoHtjkISYNTMxikjvC1\nXHctTYds0D6Q6u7igkN9ew8ngn61LInFqb6dLm+CmQIDAQABMA0GCSqGSIb3DQEB\nBQUAA4GBAFVOQ8zXFNHdXVa045onbkx8pgM2zK5VQ69YFtlAymFDWaS7a5M+96JW\n2c3001GDGZcW6fGqW+PEyu3COImRlEhEHaZKs511I7RfckMzZ3s7wPrQrC8WQLqI\ntiZtCWfUX7tto7YDdmfol7bHiaSbrLUv4H/B7iS9FGemA+nrghCK\n-----END CERTIFICATE-----"
	directory_user_context = ["OU=US,OU=Users,OU=Accounts,dc=Subdomain,dc=example,dc=com", "ou=People,o=example.com" ]
	ilo_distinguished_name = "service"
	password = "test_password"
	kerberos_authentication = false
      }
      directory_groups {
	group_dn = "ilos.example.com,ou=Groups,o=example.com"
	group_sid = "S-1-5-12"
	user_config_priv = false
	remote_console_priv = true
	virtual_media_priv = true
	virtual_power_and_reset_priv = true
	ilo_config_priv = true
      }
      key_manager {
	primary_server_address = "192.0.2.91"
	primary_server_port = 9000
	secondary_server_address = "192.0.2.92"
	secondary_server_port = 9000
	redundancy_required = true
	group_name = "GRP"
	certificate_name = "Local CA"
	login_name = "deployment"
	password = "test_password"
      }
      administrator_account {
	delete_administrator_account = false
	password = "test_password"
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
  connection_settings {
    manage_connections = true
    compliance_control = "CheckedMinimum"
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
      network_uri    = data.oneview_ethernet_network.ethernetnetworks1.uri
      port_id        = "Auto"
      requested_mbps = "2500"
    }
  }
}

*/
# Creates server profile template with san storage
/*
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
      network_uri    = "/rest/fc-networks/37aae264-8fd5-4624-960d-10173bde5752"
    }
    connections {
      id             = 2
      name           = "Management-01"
      isolated_trunk = false     
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
