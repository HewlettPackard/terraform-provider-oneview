---
layout: "oneview"
page_title: "Oneview: server_profile_template"
sidebar_current: "docs-server_profile_template"
description: |-
 Gets information about an existing server_profile_template.
---

# oneview\_server\_profile\_template

Use this data source to access the attributes of a Server Profile Template.

## Example Usage

```hcl
data "oneview_server_profile_template" "test" {
 name = "Test server profile template"
}

output "oneview_server_profile_template_value" {
 value = "${data.oneview_server_profile_template.test.uri}"
}
```

## Argument Reference

* `name` - (Required) A unique name for the resource.

* `affinity` - (Optional) This identifies the behavior of the server profiles created from this template when the server 
hardware is removed or replaced. This can be set to Bay or BayAndServer. 
This defaults to Bay.
* `bios` - (Optional) Server BIOS settings.
  
* `boot`- (Optional) Defines the order in which boot will be attempted on the available devices. Different hardware 
take different boot orders. Refer to the api documentation for your specific boot order options.
* `boot_mode `- (Optional)  Boot mode settings to be configured on the server.
* `category` - (Optional) Identifies the resource category. This field should always be set to 'server-profile-templates'.
* `connectionSettings` - (Optional) The profile connections configuration.

Connection settings configuration is specified below.

  
* `enclosure_group` - (Required) Identifies the enclosure group name for which the Server Profile Template was designed. 
The enclosure group is determined when the profile template is created and cannot be modified.

* `eTag` - Entity tag/version ID of the resource.

* `firmware` - Defines and enables firmware baseline management.

*  `hide_unused_flex_nics` - (Optional) Hides flex nics that aren't in use.
  This defaults to true.

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource is assigned.
 
* `iscsi_initiator_name_type` - When set to UserDefined, the value of iscsiInitatorName will need to be provided in the server profile. 
 
* `local_storage` - Local storage settings to be configured on the server.
 
* `mac_type` - (Optional) Specifies the type of MAC address to be programmed into the IO devices. The value can be 'Virtual'
  
* `management_processor` - (Optional) Server management processor settings.
   
* `os_deployment_settings` - OS deployment settings applicable when deployment is invoked through a server profile
  
* `refresh_state` - Current refresh State of this Server Profile Template.
  
* `san_storage` - The profile SAN storage configuration.
  
* `scopes_uri`-   The URI for the resource scope assignments.

* `serial_number_type` - (Optional) Specifies the type of Serial Number and UUID to be programmed into the server ROM. The value can be 'Virtual' or 'Physical'. Changing this forces a new resource.
This defaults to 'Virtual'.
 
*  `server_hardware_type` - (Required) Identifies the server hardware type name for which the Server Profile Template was designed. The server hardware type is determined when the profile template is created and cannot be modified.

* `server_hardware_type_uri` Identifies the server hardware type for which the Server Profile Template was designed

* `status` - Overall health status of this Server Profile Template.
 
* `uri` - The URI of the created resource.
 
* `wwn_type` - (Optional) Specifies the type of WWN address to be programmed into the IO devices. The value can be 'Virtual' or 'Physical'. Changing this forces a new resource. This defaults to 'Virtual'or 'Physical'. Changing this forces a new resource.This defaults to 'Virtual'.

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
			* `user_name`, `display_name`, `password`, `user_config_priv`, `remote_console_priv`, `virtual_media_priv`, `virtual_power_and_reset_priv`, and `ilo_config_priv`
