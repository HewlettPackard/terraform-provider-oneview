---
layout: "oneview"
page_title: "Oneview: server_profile"
sidebar_current: "docs-oneview-server-profile"
description: |-
  Create/Update a Server Profile.
---

# oneview\_server\_profile
Creates or updates a Server Profile.
## Example Usage

Create server profile with local storages.
```hcl
resource "oneview_server_profile" "default" {
  name = "test-server-profile"
  hardware_name = "00AT092 bay 2"
  type = "ServerProfileV10"
  local_storage {
    controller {
      device_slot       = "Embedded"
      drive_write_cache = "Unmanaged"
      initialize        = false
      import_configuration = false
      mode                     = "RAID"
      predictive_spare_rebuild = "Unmanaged"
      logical_drives {
        name                = "LD-1"
	drive_number	    = 1
	num_physical_drives = 2
        bootable            = false
        drive_technology    = "SasHdd"
        raid_level          = "RAID1"
        accelerator         = "Unmanaged"
      }
      logical_drives {
        name                = "LD-2"
	drive_number	    = 2
	num_physical_drives = 2
        bootable            = false
        drive_technology    = "SasHdd"
        raid_level          = "RAID1"
        accelerator         = "Unmanaged"
      }
    }
  }
}
```
Update: Removing LD-1 Logical Drive from the server profile.

Note: If you want to delete a block from the server profile resource, you will need to keep it as a empty block in your configuration.
```hcl
resource "oneview_server_profile" "default" {
  name = "test-server-profile"
  hardware_name = "00AT092 bay 2"
  type = "ServerProfileV10"
  local_storage {
    controller {
      device_slot       = "Embedded"
      drive_write_cache = "Unmanaged"
      initialize        = false
      import_configuration = false
      mode                     = "RAID"
      predictive_spare_rebuild = "Unmanaged"
      
      // Keeping logical_drive as an empty block in its position to delete it.
      logical_drives {
      }
      
      logical_drives {
        name                = "LD-2"
	drive_number	    = 2
	num_physical_drives = 2
        bootable            = false
        drive_technology    = "SasHdd"
        raid_level          = "RAID1"
        accelerator         = "Unmanaged"
      }
    }
  }
}
```
Create server profile using server profile template.
```hcl
resource "oneview_server_profile" "default" {
  name = "test-server-profile"
  template = "${oneview_server_profile_template.test.name}"
}
```
Update request for server profile
```hcl
resource "oneview_server_profile" "default" {
       operation_type {
          name = "reapply_SAN_storage"        
          
        }
        name = "TestSP"
        type = "ServerProfileV10"
        server_hardware_type = "SY 480 Gen9 3"
        enclosure_group = "SYN03_EC"
        hardware_name = "SYN03_Frame3, bay 1"
}
```

Update request for renaming  logical drive name
```hcl
resource "oneview_server_profile" "default" {
      operation_type {
          name = "rename_logical_drive"
          logical_drive_value{
            device_slot_name="Embedded"
            current_logical_drive_name="ld2"
            new_logical_drive_name="ld3"
          }        
          
        } 
        name = "TestSP"
        type = "ServerProfileV10"
        server_hardware_type = "SY 480 Gen9 3"
        enclosure_group = "SYN03_EC"
        hardware_name = "SYN03_Frame3, bay 1"
}
```
Following update operation can be performed by giving the update string for given type of update.
| NO          | Type of Update | Update String |
 | ----------- | -------------- | ------------- |
 | 1           | `Update the server profile from the server profile template.`       | 'update_from_template'         |
  | 2           | `Stage a server profile to update from its server profile template the next time the server is powered off via OneView.` | 'update_from_template_stage'         |
 | 3           | `Cancel the staging on an update from template for a server profile that was previously staged.	`        | 'cancel_update_from_template' |
 | 4           | `Refresh the server profile`        | 'refresh' |
 | 5           | `Reapply the server profile's firmware`        | 'reapply_firmware' |
  | 6           | `Reapply the server profile's connections`        | 'reapply_connection' |
 | 7           | `Reapply the server profile's adapters and boot settings`        | 'reapply_adpater_boot_settings' |
  | 8           | `Rename a server profile's logical drive`        | 'rename_logical_drive' |
 | 9           | `Reapply the server profile's local storage`        | 'reapply_local_storage' |
 | 10           | `Reapply the server profile's SAN storage`        | 'reapply_SAN_storage' |
 | 11           | `Reapply the server profile's bios settings`        | 'reapply_bios_settings' |
 | 12           | `Regenerate CHAP Secrets`        | 'regenerate_CHAP_secrets' |
 | 13           | `Reapply the server iLO settings`        | 'reapply_ilo_settings' |

## Argument Reference
- - -
The following arguments are supported: 

* `name` - (Required) A unique name for the resource.

* `template` - (Optional) The name of the template you will use for the server profile. 

* `Operation_type` - (Optional) patch operation can be performed by giving the update string for given type of update. 

- - -

* `public_connection` - (Optional) The name of the network that is going out to the public.

* `hardware_name` - (Optional) The name of the Server Hardware the server will be provisioned on.
  If this isn't used, a server hardware will be picked based on compatibility with the server profile template and any hw_filter(s) (see below).

* `hw_filter` - (Optional) List of filters to apply to the search for HW. See the OneView API docs pertaining to common filter query params, but the basic format is `[not] {attribute} {operator} '{value}'`. For example, `hw_filter = ["memoryMb >= 4096", "processorCoreCount = 4", "processorSpeedMhz >= 2400", "processorType regex '^Intel.*'"]`

* `power_state` - (Optional) Power state to enforce; `"on"` or `"off"`

* `type` - (Optional) The server profile version to be provisioned. Defaults to ServerProfileV5.
  Use ServerProfileV6 to use Image Streamer.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `serial_number` - A 10-byte value that is exposed to the Operating System as the server hardware's
  Serial Number. The value can be a virtual serial number, user defined serial number or physical serial
  number read from the server's ROM. It cannot be modified after the profile is created.

* `public_mac` - The MAC address of the NIC your public network is attached to.
  Need to specify public_connection to access this value. 
  
* `public_slot_id` - The slot id of the NIC your public network is attached to.
  Need to specify public_connection to access this value. 
  
* `ilo_ip` - The ILO ip address that is managing the server.

* `hardware_uri` - The URI of the hardware the server is provisioned on.

* `boot_order` - Defines the order in which boot will be attempted on the available devices.

* `boot_mode` -  Boot mode settings to be configured on the server.

* `bios_option` - Server BIOS settings.

* `server_hardware_type` - Identifies the server hardware type for which the Server Profile was designed. 

* `enclosure_group` -  Identifies the enclosure group for which the Server Profile was designed.

* `affinity` - This identifies the behavior of the server profile when the server hardware is removed or replaced. 

* `hide_unused_flex_nics` - This setting controls the enumeration of physical functions that do not correspond to connections in a profile.

* `serial_number_type` -  Specifies the type of Serial Number and UUID to be programmed into the server ROM.

* `wwn_type` -  Specifies the type of WWN address to be programmed into the IO devices. 

* `mac_type` - Specifies the type of MAC address to be programmed into the IO Devices.

* `firmware` - Defines and enables firmware baseline management.

* `local_storage` -  Local storage settings to be configured on the server.

* `logical_drives` - The list of logical drives associated with the controller. 

* `san_storage` - The profile san storage configuration. The san storage attributes list is available in the oneview Rest API reference guide.

* `volume_attachments` - The list of storage volume attachments.

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
* `firmware` - Defines and enables firmware baseline management.
	* `consistency_state` - Consistency state of the firmware component.
	* `force_install_firmware` - Force installation of firmware even if same or newer version is installed. Value can be 'true' or 'false'.
	* `firmware_baseline_uri` - Identifies the firmware baseline to be applied to the server hardware.
	* `firmware_activation_type` - Specifies when the applied Service Pack for ProLiant (SPP) will be activated.
	* `manage_firmware` - Indicates that the server firmware should be configured on the server profiles created from the template. Value can be 'true' or 'false'.
	* `firmware_install_type` -  Specifies the way a firmware baseline is installed. This field is used if the 'manageFirmware' field is true. Options are FirmwareAndOSDrivers, FirmwareOnly and FirmwareOnlyOfflineMode
	* `firmwareScheduleDateTime` -  Identifies the date and time the Service Pack for Proliant (SPP) will be activated.
	* `reapply_state` - Current reapply state of the firmware component.
* `management_processor` - Server management processor settings.
	* `compliance_control` -  Defines the compliance type of template's management processor settings with the corresponding profile's Management Processor settings. Valid values are "Checked" and "Unchecked".
	* `manage_mp` -  Indicates whether the management processor settings are configured using the server profile. Value can be 'true' or 'false'
	* `reapply_state` - Current reapply state of the mpSettings component.
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
			* `login_priv`, `host_bios_config_priv`, `host_nic_config_priv`, `host_storage_config_priv` attributes are supported for Gen 10 onwards.

* `associated_server` - The serial number of the server hardware that the server profile is currently applied to or was most recently assigned to. This value is cleared if a different server profile is assigned to the server hardware.

* `category` - Identifies the resource category. This field should always be set to 'server-profiles'.

* `created` - The time that the Server Profile was created.

* `description` - The description of this Server Profile.

* `enclosure_bay` - Identifies the enclosure device bay number that the Server Profile is currently assgined to, if applicable.

* `etag` - Entity tag/version ID of the resource.

* `in_progress` - Indicates whether the task identified by taskUri is currently executing.

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource shall be assigned.
It is meaningful at resource creation time, during resource update, and it is included on resource retrieval as well.

* `iscsi_initiator_name` - When iscsiInitatorNameType is set to UserDefined, this field specifies the default iSCSI initiator identifier used by connections that have the InitiatorNameSource set to ProfileInitiatorName. This field is ignored when iscsiInitatorNameType is set to AutoGenerated.

* `iscsi_initiator_name_type` - When set to UserDefined, the value of iscsiInitatorName is used as provided. When set to AutoGenerated, the default iscsiInitatorName will be constructed to be unique to the server profile.

* `modified` - The time that the Server Profile was most recently modified.

* `profile_uuid` - The automatically generated 36-byte Universally Unique ID of the server profile.

* `refresh_state` - Current refresh State of this Server Profile.

* `scopes_uri` - The URI for the resource scope assignments.

* `server_hardware_reapply_state` - Current reapply state of the server that is associated with this server profile. This includes the virtual serial number, UUID, boot settings and server hardware adapter configuration.

* `server_hardware_type_uri` - The UTI of the server hardware type for which the Server Profile was designed.

* `service_manager` - Name of a service manager that is designated owner of the profile.

* `state` - Current State of this Server Profile.

* `status` - Overall health status of this Server Profile.

* `task_uri` - URI of the task currently executing or most recently executed on this server profile.

* `template_compliance` - The compliance state of the server profile with the server profile template.

* `uuid` - The 36-byte value that is exposed to the Operating System as the server hardware's Universally Unique ID.
