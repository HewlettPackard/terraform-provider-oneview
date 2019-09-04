provider "oneview" {
	ov_username = "gireesan"
	ov_password = "P@ssw0rd!"
	ov_endpoint = "https://10.50.4.100"
	i3s_endpoint = "https://10.50.4.106"	
	ov_sslverify = false
	ov_apiversion = 1000
	ov_ifmatch = "*"
}

resource "oneview_server_profile_template" "ServerProfileTemplate" {
	name = "TestServerProfileTemplate"
	type = "ServerProfileTemplateV6"
	enclosure_group = "SYN03_EC"
	server_hardware_type = "SY 480 Gen9 2"
	network = [{
		name = "Deployment Network A"
		function_type = "Ethernet"
		network_uri = "/rest/ethernet-networks/29af1597-7f2e-45d8-aaed-ee1be6c42ae2"
		port_id = "Mezz 3:1-a"
		},
		{
		name = "Deployment Network B"
		function_type = "Ethernet"
		network_uri = "/rest/ethernet-networks/29af1597-7f2e-45d8-aaed-ee1be6c42ae2"
		port_id = "Mezz 3:2-a"
		},
		{
		name = ""
		function_type = "FibreChannel"
		port_id = "Mezz 3:1"
		network_uri = "/rest/fc-networks/429006d8-24e2-4c52-8e08-58a1ea1cb985"
		},
		{
		name = ""
		function_type = "FibreChannel"
		network_uri = "/rest/fc-networks/7884fa5e-1b5a-4f56-b52c-459884bccaea"
		port_id = "Mezz 3:2"
		}]
	manage_connections = true
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
	san_storage = {
		host_os_type = "Windows 2012 / WS2012 R2"
		manage_san_storage = true
	}
	volume_attachments = [{
		lun_type = "Auto"
		lun = ""
		volume_uri = "/rest/storage-volumes/B9981C13-EED1-4F21-B95F-A93D00D23E3F"
		storage_paths = [{
			is_enabled = true
			connection_id = 3
			storage_target_type = "Auto"
			storage_targets = []
		},
		{
		is_enabled = true
		connection_id = 4
		storage_target_type = "Auto"
		storage_targets = []
		}]
	}]
	os_deployment_settings = {
		os_deployment_plan_name = "HPE - Foundation 1.0 - create empty OS Volume-2017-10-13"
	}
}
