provider "oneview" {
	ov_username = <ov-username>
	ov_password = <ov-password>
	ov_endpoint = <ov-endpoint>
	i3s_endpoint = <i3s-endpoint>
	ov_sslverify = false
	ov_apiversion = <ov-api-version>
	ov_ifmatch = "*"
}

resource "oneview_server_profile_template" "ServerProfileTemplate" {
	name = "TestServerProfileTemplateTerraform"
	type = "ServerProfileTemplateV6"
	enclosure_group = "SYN03_EC"
	server_hardware_type = "DL380p Gen8 1 (new name)"
	network = [{
		id = 1
		name = "Deployment Network A"
		function_type = "Ethernet"
		network_uri = "/rest/ethernet-networks/29af1597-7f2e-45d8-aaed-ee1be6c42ae2"
		port_id = "Mezz 3:1-a"
		boot = {
			priority = "Primary"
			ethernet_boot_type = "iSCSI"
			boot_volume_source = "UserDefined"
			iscsi = {
				initiator_name_source="ProfileInitiatorName"
				chap_level="None"
				}
			}
		ipv4 = {
			ip_address_source = "SubnetPool"
			}
		},
		{
		id = 2
		name = "Deployment Network B"
		function_type = "Ethernet"
		network_uri = "/rest/ethernet-networks/29af1597-7f2e-45d8-aaed-ee1be6c42ae2"
		port_id = "Mezz 3:2-a"
		boot = {
			priority = "Secondary"
			ethernet_boot_type = "iSCSI"
			boot_volume_source = "UserDefined"
			iscsi = { 
				initiator_name_source="ProfileInitiatorName"
				chap_level="None"
				}
			}
		ipv4 = {
			ip_address_source = "Subnetpool"
			}
		},
		{
		id = 3
		name = "fc01"
		function_type = "FibreChannel"
		port_id = "Mezz 3:1"
		network_uri = "/rest/fc-networks/429006d8-24e2-4c52-8e08-58a1ea1cb985"
		boot = {
			priority = "NotBootable"
			}
		},
		{
		id = 4
		name = "fc02"
		function_type = "FibreChannel"
		network_uri = "/rest/fc-networks/7884fa5e-1b5a-4f56-b52c-459884bccaea"
		port_id = "Mezz 3:2"
		boot = {
			priority = "NotBootable"
			}
		},
		{
		id = 5
		name = "mgmt1"
		port_id = "Auto"
		function_type = "Ethernet"
		requested_mbps = 5000
		network_uri = "/rest/ethernet-networks/6de2920a-8ad4-4cd8-865c-1907d3b4682e"
		boot = {
			priority = "NotBootable"
			}
		},
		{
		id = 6
		name = "mgmt2"
		function_type = "Ethernet"
		port_id = "Auto"
		network_uri = "/rest/ethernet-networks/3ebf86fb-89fd-4cf8-b369-441690555cea"
		requested_mbps = 10000
		boot = {
			priority = "NotBootable"
			}
		}]
	manage_connections = true
	boot_order = ["HardDisk"]

	boot_mode = {
		manage_mode = true
		mode = "UEFIOptimized"
		pxe_boot_policy = "Auto"
	}
	firmware = {
		force_install_firmware = false
		firmware_baseline_uri = "/rest/firmware-drivers/SPP_2018_11_20190205_for_HPE_Synergy_Z7550-96592"
		manage_firmware = true
		firmware_install_type = "FirmwareAndOSDrivers"
	}
	bios_option = {
		manage_bios = true
		overridden_settings = [
			{ id = "UsbControl"
			value = "UsbEnabled"},
			{ id = "PowerRegulator"
			value = "StaticHighPerf"},
			{ id = "CollabPowerControl"
			value = "Disabled"},
			{ id = "EnergyPerfBias"
			value = "MaxPerf"},
			{ id = "MinProcIdlePkgState"
			value = "NoState" },
			{ id = "NumaGroupSizeOpt"
			value = "Clustered" },
			{ id = "MinProcIdlePower"
			value = "NoCStates"
			}]
	}
	san_storage = {
		host_os_type = "Windows 2012 / WS2012 R2"
		manage_san_storage = true
		//compliance_control = "CheckedMinimum"
	}
	volume_attachments = [{
		id = 1
		lun_type = "Manual"
		lun = "10"
		volume_uri = "/rest/storage-volumes/B9981C13-EED1-4F21-B95F-A93D00D23E3F"
		boot_volume_priority = "NotBootable"
		storage_paths = [{
			is_enabled = true
			connection_id = 3
			target_selector = "Auto"
			storage_targets = []
			},
			{
			is_enabled = true
			connection_id = 4
			target_selector = "Auto"
			storage_targets = []
			}]
	}]
	
	os_deployment_settings = {
		os_deployment_plan_name = "RHEL"
		os_custom_attributes = [{
			 
            name="DiskName"
            value="/dev/sda"},
         { 
            name="DomainName"
            value="eco.demo.local"
          },
         { 
            name="FirstPartitionSize"
            value="10"
          
         },
	 {
	   name="HostName"
	  value="rhel7601"},
         { 
            name="LogicalVolumeGroupName"
            value="new_vol_group"},
         { 
            name="LogicalVolumeName"
            value="new_vol"},
         { 
            name="LogicalVolumeSize"
            value="15"},
         { 
            name="ManagementNIC1.constraint"
            value="dhcp"
         },
         { 
            name="ManagementNIC1.connectionid"
            value="5"
         },
         { 
            name="ManagementNIC1.networkuri"
            value="/rest/ethernet-networks/6de2920a-8ad4-4cd8-865c-1907d3b4682e"
         },
         { 
            name="ManagementNIC1.vlanid"
            value="0"
         },
         { 
            name="ManagementNIC2.connectionid"
            value="none"
         },
         { 
            name="ManagementNIC2.mac"
            value="none"
         },
         { 
            name="ManagementNIC3.connectionid"
            value="none"
         },
         { 
            name="ManagementNIC3.mac"
            value="none"
         },
         { 
            name="ManagementNIC4.connectionid"
            value="none"
         },
         { 
            name="ManagementNIC4.mac"
            value="none"
         },
         { 
            name="NewRootPassword"
            value="admin"
         },
         { 
            name="NewUser"
            value="admin"},
         { 
            name="NewUserPassword"
            value="admin"
         },
         { 
            name="SSH"
            value="Enabled" },
         { 
            name="SecondPartitionSize"
            value="10"},
         { 
            name="TotalMgmtNICs"
            value="1" 	        }]

	}
}


resource "oneview_server_profile" "SP" {
	name = "TestSpTerraform"
	hardware_name = "SYN03_Frame1, bay 3"
	type = "ServerProfileV10"
	template = "${oneview_server_profile_template.ServerProfileTemplate.name}"
	power_state = "on"
	os_deployment_settings = {
		os_custom_attributes = [{
			name="HostName"
			value="rheltest"
		}]
	}
	depends_on = ["oneview_server_profile_template.ServerProfileTemplate"]
}
