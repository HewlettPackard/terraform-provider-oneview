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

* `initial_scope_uris`  - A list of URIs of the scopes to which the resource shall be initially assigned.
 
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



Network Resource Supports following:
* `boot` - Indicates that the server will attempt to boot from this connection.
  
* `name` - (Required) A unique name for the resource.

* `function_type` - (Required) Type of function required for the connection. Values can be 'Ethernet' or 'FibreChannel' Changing this forces a new resoure.


* `isolocated_trunk` - (Optional) When selected, for each PVLAN domain, primary VLAN ID tags will translated to the isolated VLAN ID tags for traffic egressing to the downlink ports.
  
* `lag_name` - (Optional) The link aggregation group name for a server profile connection.
  
* `network_uri` - (Required) Identifies the network or network set to be connected. 

* `port_id` - (Optional) Identifies the port (FlexNIC) used for this connection. Defaults to "Lom 1:1-a".

* `requested_mbps` - (Optional) The transmit throughput (mbps) that should be allocated to this connection.Defaults to `2500`

* `id` - (Optional) A unique identifier for this connection.


* `ipv4` - (Optional) The IP information for a connection. This is only used for iSCSI connections. It must be omitted for other connection types.
  


