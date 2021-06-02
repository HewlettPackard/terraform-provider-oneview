---
layout: "oneview"
page_title: "Oneview: server_profile"
sidebar_current: "docs-oneview-server-profile"
description: |-
	Gets information about an existing server_profile.
---

# oneview\_server\_profile

Use this data source to access the attributes of a Server Profile.

## Example Usage

```js
data "oneview_server_profile" "sp" {
        name = "TestAll"
}

output "oneview_server_profile_value" {
        value = "${data.oneview_server_profile.sp.uri}"
}
```

## Argument Reference

The following arguments are supported: 

* `name` - (Required) A unique name for the resource.

- - -

* `public_connection` -  The name of the network that is going out to the public.

* `hardware_name` -  The name of the Server Hardware the server will be provisioned on.
  If this isn't used, a server hardware will be picked based on compatibility with the server profile template and any hw_filter(s) (see below).

* `type` -  The server profile version to be provisioned. Defaults to ServerProfileV5.
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

* `san_storage` - The profile san storage configuration.

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
