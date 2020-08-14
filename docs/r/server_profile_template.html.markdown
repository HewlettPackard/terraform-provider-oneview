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

```js
resource "oneview_server_profile_template" "default" {
  name = "test-server-profile-template"
  enclosure_group = "my_enclosure_group"
  server_hardware_type = "BL460c Gen9 1"
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

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource shall be initially assigned. It is only meaningful at resource creation time and is not included on resource retrieval.

* `local_storage` - (Optional) Local storage settings to be configured on the server.

* `logical_drives` - (Optional) List of logical drives associated with the controller.

* `mac_type` - (Optional) Specifies the type of MAC address to be programmed into the IO devices. The value can be 'Virtual' or 'Physical'. Changing this forces a new resource.
This defaults to 'Virtual'.

* `manage_connections` - (Optional)  Identifies whether connections should be managed by server profile template. If this is false then the compliance check between server profile template and server profile is disabled, this allows a server profile created from a profile template to define any connectivity.

* `network` - (Optional) Network connection to be configured for the server. Can be specified multiple times. 
Network configuration is specified below.

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

Network supports the following:

* `name` - (Required) A unique name for the resource.

* `function_type` - (Required) Type of function required for the connection. Values can be 'Ethernet' or 'FibreChannel'
Changing this forces a new resoure.

* `network_uri` - (Required) Identifies the network or network set to be connected. 

* `port_id` - (Optional) Identifies the port (FlexNIC) used for this connection. Defaults to "Lom 1:1-a".

* `requested_mbps` - (Optional) The transmit throughput (mbps) that should be allocated to this connection.
Defaults to `2500`

* `boot` - (Optional) Boot setting required for the server profile template.

* `id` - (Optional) A unique identifier for this connection.

* `ipv4` - (Optional)  The IP information for a connection. This is only used for iSCSI connections. It must be omitted for other connection types.

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
