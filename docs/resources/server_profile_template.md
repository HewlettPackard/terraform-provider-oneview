---
layout: "oneview"
page_title: "Oneview: server_profile_template"
sidebar_current: "docs-oneview-server-profile-template"
description: |-
  Creates a server profile template.
---

# oneview\_server\_profile\_template

Creates a server profile template.

## Example Usage

```hcl
resource "oneview_server_profile_template" "ServerProfileTemplateWithLocalStorage" {
  name                 = "TestServerProfileTemplate_with_local_storage"
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
```

## Argument Reference

The following arguments are supported: 

* `name` - (Required) A unique name for the resource.

* `enclosure_group` - (Required) Identifies the enclosure group name for which the Server Profile Template was designed. 
The enclosure group is determined when the profile template is created and cannot be modified. 

* `server_hardware_type` - (Required) Identifies the server hardware type name for which the Server Profile Template was 
designed. The server hardware type is determined when the profile template is created and cannot be modified.

- - -

* `affinity` - (Optional) This identifies the behavior of the server profiles created from this template when the server 
hardware is removed or replaced. This can be set to Bay or BayAndServer. 
This defaults to Bay.
  
* `bios` - (Optional) Server BIOS settings.
Bios configuration is specified below.

* `boot_mode` - (Optional) Boot mode settings to be configured on the server. For Gen 7 and Gen 8 servers these values should not be specified.
Boot mode configuration is specified below.

* `boot_order`- (Optional) Defines the order in which boot will be attempted on the available devices. Different hardware take different boot orders. Refer to the api documentation for your specific boot order options.

* `firmware` - (Optional) Firmware attributes required to configure firmware.
Firmware configuration is specified below.

* `hide_unused_flex_nics` - (Optional) Hides flex nics that aren't in use.
This defaults to true.

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource shall be assigned.
It is meaningful at resource creation time, during resource update, and it is included on resource retrieval as well.

* `local_storage` - (Optional) Local storage settings to be configured on the server.

* `logical_drives` - (Optional) List of logical drives associated with the controller.

* `mac_type` - (Optional) Specifies the type of MAC address to be programmed into the IO devices. The value can be 'Virtual' or 'Physical'. Changing this forces a new resource.
This defaults to 'Virtual'.

* `manage_connections` - (Optional)  Identifies whether connections should be managed by server profile template. If this is false then the compliance check between server profile template and server profile is disabled, this allows a server profile created from a profile template to define any connectivity.

* `connection_settings` - Connection downlinks associated with server profile. 
  *  `reapply_state` - Current reapply state of the connection downlinks associated with the server profile.
  * `connection` - A list of connections associated with server profile.
    * `allocated_mbps` -  The transmit throughput (mbps) currently allocated to this connection. 
    * `name` - A string used to identify the respective connection.
    * `allocated_vfs` - The number of virtual functions allocated to this connection.
    * `function_type` - Type of function required for the connection. functionType cannot be modified after the connection is created.
    * `network_uri` - The name of the network or network set to be connected.
    * `port_id` - Identifies the port/FlexNIC on an adapter used for this connection.
    * `requested_mbps` - The transmit throughput (mbps) that should be allocated to this connection.
    * `requested_vfs` - The SR-IOV virtual functions that should be allocated to this connection.
    *  `state` - The state of a connection.
    *  `status` - The status of a connection.
    *  `wwnn` - The node WWN address that is currently programmed on the FlexNic. 
    *  `wwpn` - The port WWN address that is currently programmed on the FlexNIC. 
    *  `wwpn_type` -  Specifies the type of WWN address to be porgrammed on the FlexNIC. 
    *  `id` - A unique identifier for this connection. 
    *  `interconnect_port` - The interconnect port associated with the connection.
    *  `interconnect_uri` - The interconnectUri associated with the connection.
    *  `isolated_trunk` - When selected, for each PVLAN domain, primary VLAN ID tags will translated to the isolated VLAN ID tags for traffic egressing to the downlink ports. 
    *  `lag_name` - The link aggregation group name for a server profile connection. 
    *  `mac` - The MAC address that is currently programmed on the FlexNic. 
    *  `mac_type` - Specifies the type of MAC address to be programmed into the IO Devices. 
    *  `managed` - Indicates whether the connection is capable of Virtual Connect functionality and managed by OneView. 
    *  `maximum_mbps` - Maximum transmit throughput (mbps) allowed on this connection.
    *  `network_name` - The name of the network or network set to be connected. 
    *  `private_vlan_port_type` - Private Vlan port type.This is a read only field
    *  `boot` - Indicates that the server will attempt to boot from this connection.
    	* `priority` - Indicates the boot priority for this connection. 
    	* `boot_vlan_id` - The virtual LAN ID of the boot connection
    	* `ethernet_boot_type` - Indicates the boot protocol for a connection with Ethernet functionType. 
    	* `boot_volume_source` - Indicates boot volume source for the connection.
    	* `targets` - Defines the boot targets that the server will attempt to boot. 
    	  * `array_wwpn` - The wwpn of the target device that provides access to the Boot Volume.
    	  * `lun` - The LUN of the Boot Volume presented by the target device.
    	* `iscsi` - This object contains the iSCSI parameters of the connection when functionType is iSCSI, or when functionType is Ethernet and ethernetBootType is iSCSI.
    	  * `chap_level` - The iSCSI Challenge Handshake Authentication Protocol (CHAP) method. 
    	  * `chap_name` - The iSCSI CHAP name. 
    	  * `chap_secret` - The iSCSI CHAP secret. 
    	  * `initiator_name` - The unique identifier of the iSCSI initiator in iQN, EUI or NAA format.
    	  * `initiator_name_source` - Indicates how the iSCSI initiator name initiatorName was created. 
    	  * `mutual_chap_name` - The iSCSI Mutual Challenge Handshake Authentication Protocol (Mutual-CHAP) name. 
    	  * `mutual_chap_secret` - The iSCSI Mutual-CHAP secret.
    	  * `boot_target_lun` - The LUN number of the iSCSI target volume.
    	  * `boot_target_name` - The unique identifier of the iSCSI target volume in iQN, EUI or NAA format.
    	  * `first_boot_target_ip` - The IP address of the iSCSI target volume that is used first to attempt to connect. 
    	  * `first_boot_target_port` - The port number to be used for the iSCSI target volume designated by firstBootTargetIp. 
    	  * `second_boot_target_ip` - The IP address of the iSCSI target volume that is used (if given) if the connection designated by firstBootTargetIp fails
    	  * `second_boot_target_port` - The port number to be used for the iSCSI target volume designated by secondBootTargetIp.
    	* `ipv4` - The IP information for a connection. This is only used for iSCSI connections. 
    	  * `gateway` - The gateway for the iSCSI initiator. 
    	  * `ip_address` - The IPv4 address of the iSCSI initiator. When creating a connection, if ipAddressSource is DHCP, then this must be omitted.
    	  * `subnet_mask` - The subnet mask of the iSCSI initiator. 
    	  * `ip_address_source` - Specifies how the IPv4 parameters are to be supplied.

* `management_processor` - Server management processor settings.
	* `compliance_control` -  Defines the compliance type of template's management processor settings with the corresponding profile's Management Processor settings. Valid values are "Checked" and "Unchecked".
	* `manage_mp` -  Indicates whether the management processor settings are configured using the server profile. Value can be 'true' or 'false'
	* `mp_settings` - The management processor settings to be modified. Below are the attributes supported.
		* `administrator_account` - Below attributes are supported for addministrator account.
			* `delete_administrator_account` and `password`
		* `directory` - Below attributes are support for directory.
			* `directory_authentication`, `directory_generic_ldap`, `directory_server_address`, `directory_server_port`, `directory_server_certificate`, `directory_user_context`, `ilo_distinguished_name`, `password`, `kerberos_authentication`, `kerberos_realm`, `kerberos_kdc_server_address`, `kerberos_kdc_server_port`, and `kerberos_key_tab`.
		* `key_manager` - Below attributes are supported for key manager. 
			* `primary_server_address`, `primary_server_port`, `secondary_server_address`, `secondary_server_port`, `redundancy_required`, `group_name`, `certificate_name`, `login_name`, and `password`.
		* `directory_groups`  - Below attributes are supported for directory groups. 
			* `group_dn`, `group_sid`, `user_config_priv`, `remote_console_priv`, `virtual_media_priv`, `virtual_power_and_reset_priv`, and `ilo_config_priv`.
		* `local_accounts` - Below attributes are supported for local accounts.
			* `user_name`, `display_name`, `password`, `user_config_priv`, `remote_console_priv`, `virtual_media_priv`, `virtual_power_and_reset_priv`, and `ilo_config_priv`. 
			*  `login_priv`, `host_bios_config_priv`, `host_nic_config_priv`, `host_storage_config_priv` attributes are supported for Gen 10 onwards.

* `os_deployment_settings` - (Optional) OS Deployment settings applicable when deployment is invoked through a server profile.
OS Deploymennt Stting configuration is specified below.

* `san_storage` - (Optional)The profile SAN storage configuration.
San Storage configuration is specified below.

* `serial_number_type` - (Optional) Specifies the type of Serial Number and UUID to be programmed into the server ROM. 
The value can be 'Virtual' or 'Physical'. Changing this forces a new resource.
This defaults to 'Virtual'.
  
* `type` - (Optional) Uniquely identifies the type of the JSON object.
This Defaults to ServerProfileTemplateV1

* `volume_attachments` - (Optional) The list of storage volume attachments.
Volume Attachment configuration is specified below.

* `wwn_type` - (Optional) Specifies the type of WWN address to be programmed into the IO devices. The value can be 'Virtual' 
or 'Physical'. Changing this forces a new resource. 
This defaults to 'Virtual'.

Bios Options Supports the following:

* `manage_bios` - (Required) Indicates whether the BIOS settings should be configured on the server profiles created from the template. Value can be 'true' or 'false'.

* `overridden_settings` - (Optional) The BIOS settings to be modified. All omitted BIOS settings are reset to their factory default values.

Boot Mode Supports the following:

* `manage_mode` - (Required) Indicates whether the boot mode should be configured on server profiles created from the template. Value can be 'true' or 'false'. The value defaults to 'false' when unspecified.

* `mode` - (Required) The environment used for server boot operation. Possible values are "UEFI", "UEFIOptimized" or "BIOS". This field is required only when "manageMode" is set to 'true'.

* `pxe_boot_policy` - (Required) Defines the filtering or priority of the PXE boot options for each enabled NIC port. This field is required only when the "mode" is set to "UEFI" or "UEFIOptimized". 

Firmware Supports the following:

* `force_install_firmware` - (Optional) Force installation of firmware even if same or newer version is installed. Downgrading the firmware can result in the installation of unsupported firmware which can cause the hardware to cease operating. Value can be 'true' or 'false'.

* `firmware_baseline_uri` - (Optional) Identifies the firmware baseline to be applied to the server hardware.

* `firmware_activation_type` - (Optional) Specifies when the applied Service Pack for ProLiant (SPP) will be activated.

* `manage_firmware` - (Optional) Indicates that the server firmware should be configured on the server profiles created from the template. Value can be 'true' or 'false'.

* `firmware_install_type` - (Optional) Force installation of firmware even if same or newer version is installed. Downgrading the firmware can result in the installation of unsupported firmware which can cause the hardware to cease operating. Value can be 'true' or 'false'.

Os Deployment Settings support the following:

* `os_deployment_plan_name` - (Optional) Identifies the OS deployment plan.

* `os_volume_uri` - (Optional) Identifies the OS deployment plan. Use GET /rest/os-deployment-plans to retrieve the list of available OS deployment plans.

* `os_custom_attributes` - (Optional) Identifies the custom attributes to be configured with the OS deployment plan.

San Storage Supports the following:

* `host_os_type` - (Optional) The operating system type of the host. To retrieve the list of supported host OS types, issue a REST Get request using the /rest/storage-systems/host-types API.

* `manage_san_storage` - (Optional) Identifies whether SAN storage should be configured on server profiles created from the template.

* `server_hardware_type_uri` - (Optional) Identifies the server hardware type for which the Server Profile Template was designed. The serverHardwareTypeUri is determined when the profile template is created. Use GET /rest/server-hardware-types to retrieve the list of server hardware types.

Volume Attachments Support the following:

* `associated_template_attachment_id` - (Optional) A "key" value uniquely identifying the definition of a volume attachment in a template. 

* `id` - (Optional) The ID of the storage volume attachment.

* `lun` - (Optional) Logical Unit Number. Typically a value from 0 to 255. A blank value is not allowed when lunType is "Manual". If lunType is "Auto", lun can be left blank. 

* `lun_type` - (Optional) "Manual" or "Auto". If Manual, a lun value is required. If Auto, the lun value is returned as a read only value.

* `boot_volume_priority` - (Optional) Identifies whether the volume will be used as a boot volume and with what priority. This attribute can only be set on private volumes.

* `volume_storage_system_uri` - (Optional) The URI of the storage system associated with this volume attachment. This field is read-only when the storage pool URI in the volume properties is specified.

* `volume_uri` - (Optional) The URI of the storage volume associated with this volume attachment. This attribute is required when specifying an existing volume to attach, but is not used when defining a new volume to be created.

* `storage_paths` - (Optional) A list of host-to-target path associations.

* `volume` - (Optional) Contains properties describing a volume to be created. This attribute is required when defining a new volume to be created, but is not used when an existing volume is attached.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `uri` - The URI of the created resource.

* `eTag` - Entity tag/version ID of the resource.
