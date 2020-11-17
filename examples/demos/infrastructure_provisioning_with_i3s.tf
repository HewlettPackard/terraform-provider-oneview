/*   Infrastructure Provisioning with Storage and Image Streamer 

        Be able to provision compute (with server settings), networking, and storage.
        Create a server profile template with the following options:
                Network connections
                Boot mode
                Boot settings
		Bios settings
		Firmware
		SAN storage 
		OS deployement settings
                Create a server profile from a server profile template and assign to hardware
*/

provider "oneview" {
  ov_username =   "${var.username}"
  ov_password =   "${var.password}"
  ov_endpoint =   "${var.endpoint}"
  ov_sslverify =  "${var.ssl_enabled}"
  i3s_endpoint =  "https://15.212.167.172" #"${var.i3s_endpoint}"
  ov_apiversion = 2000
  ov_ifmatch = "*"
}

# Creates Sever Profile Templates 
resource "oneview_server_profile_template" "ServerProfileTemplate" {
	name = "TestServerProfileTemplateTerraform"
	type = "ServerProfileTemplateV8"
	enclosure_group = "EG-P"
	server_hardware_type = "SY 660 Gen9 1"
	network = [{
		id = 1
		name = "Deployment Network A"
		function_type = "Ethernet"
		network_uri = "/rest/ethernet-networks/caaa21fc-1b38-4f20-9abb-221e411a5a34"
		port_id = "Mezz 6:1-a"
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
		network_uri = "/rest/ethernet-networks/caaa21fc-1b38-4f20-9abb-221e411a5a34"
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
		#compliance_control = "CheckedMinimum"
	}
	volume_attachments = [{
		id = 1
		lun_type = "Manual"
		lun = "10"
		boot_volume_priority = "NotBootable"
		volume = [{
                        template_uri = "/rest/storage-volume-templates/1a1c6a3c-587e-4930-8b0e-abf40124053e"
                        properties = [{
                                name = "vol_name"
                                size = 268435456
                                storage_pool = "/rest/storage-pools/9923DE4C-F571-4B64-8C3E-ABF40112FE60"
                        }]
		}]
		storage_paths = [{
			is_enabled = true
			connection_id = 3
			target_selector = "Auto"
			targets = []
			},
			{
			is_enabled = true
			connection_id = 4
			target_selector = "Auto"
			targets = []
			}]
	}]
	os_deployment_settings = {
		os_deployment_plan_name = "HPE - Foundation 1.0 - create empty OS Volume-2017-10-13"
		os_custom_attributes = [{
            name="VolumeSize"
            value="1"},
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

# Creates Server Profile with above defined Server Profile Template.
resource "oneview_server_profile" "SP" {
  name = "TestSpTerraform"
  hardware_name = "MXQ646057D, bay 4" #"SYN03_Frame1, bay 3"
  type = "ServerProfileV12"
  template = "test_withi3S" #"${oneview_server_profile_template.ServerProfileTemplate.name}"
    power_state = "off"
    os_deployment_settings = {
      os_custom_attributes = [{
        name="VolumeSize"
        value="1"}]
    }
 depends_on = ["oneview_server_profile_template.ServerProfileTemplate"]
}
