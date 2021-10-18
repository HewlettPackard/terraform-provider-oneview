---
layout: "oneview"
page_title: "Oneview: logical_interconnect_group"
sidebar_current: "docs-oneview-logical-interconnect-group"
description: |-
  Creates a logical interconnect group.
---

# oneview\_logical\_interconnect\_group

Creates a logical interconnect group.

Note: Instead of providing list of port nums, from terraform-oneview-provider v6.3 we need to provide single port_num for each logical_port_config block.
## Example Usage

```js
resource "oneview_logical_interconnect_group" "default" {
  name = "test-logical-interconnect-group"
  
  internal_network_uris = ["${oneview_ethernet_network.default.0.uri}"]
  
  interconnect_settings {
    fast_mac_cache_failover = false
    igmp_timeout_interval = 250
  }
  
  interconnect_map_entry_template {
    interconnect_type_name = "HP VC FlexFabric-20/40 F8 Module"
    bay_number = 1
  }
  
  snmp_configuration {
    read_community = "Group 1"
    system_contact = "admin"
    snmp_access = ["192.168.1.101"]
    
    trap_destination {
      trap_destination = "127.0.0.1"
      enet_trap_categories = ["Port Thresholds", "Other"]
      trap_severities = ["Info"]
    }
  }
  
  uplink_set {
    name = "uplink-default"
    network_uris = ["${oneview_ethernet_network.test.1.uri}"]
    logical_port_config {
      bay_num = 4
      port_num = 20
    }
    logical_port_config {
      bay_num = 4
      port_num = 21
    }
  }
}
```

## Argument Reference

The following arguments are supported: 

* `name` - (Required) A unique name for the resource.

- - -

* `internal_network_uris` - (Optional) A list of internal network URIs

* `interconnect_settings` - (Optional) A interconnect settings block. If not supplied a default will be used.
  Interconnect Settings documented below.

* `telemetry_configuration` - (Optional) The controls for collection of interconnect statistics. 
  If not supplied a default will be used. Telemetry Configuration documented below.
  
* `snmp_configuration` - (Optional) The SNMP configuration for the logical interconnect group. 
  If not supplied a default will be used. Snmp Configuration is documented below.

* `interconnect_bay_set` - (Optional) This option is required on Synergy hardware, but not needed for C7000.
  The Interconnect bay set associated with the logical interconnect group. Supported interconnect bay set is
  1 (bays 1&4), 2 (bays 2&5), or 3 (bays 3&6).

* `redundancy_type` - (Optional) This option is required on Synergy hardware, but not needed for C7000.
  The type of enclosure redundancy. Supported enclosure redundancy types are HighlyAvailable,
  NonRedundantASide, NonRedundantBSide, and Redundant.

* `enclosure_indexes` - (Optional) This option is required on Synergy hardware, but not needed for C7000.
  The list of enclosure indices that are specified by this logical interconnect group. The value [-1] indicates
  that this is a single enclosure logical interconnect group for Virtual Connect SE FC Modules. The value [1]
  indicates that this is a single enclosure logical interconnect group for other supported interconnects. If
  you are building a logical interconnect group for use with a three enclosures interconnect link topology, the
  value needs to be [1,2,3].

* `interconnect_map_entry_template` - (Optional) Interconnect map associated with the logical interconnect group.
  This can be specified multiple times. Interconnect Map Entry Template is documented below. 

* `uplink_set` - (Optional) List of uplink sets in the logical interconnect group.
  This can be specified multiple times. Uplink Set is documented below. 

Interconnect Settings support the following:

* `fast_mac_cache_failover` - (Optional) This will cause Ethernet packets to be tranmitted on newly-active links.
  Defaults to true.

* `igmp_snooping` - (Optional) Allows modules to monitor the IGMP IP multicast membership activities.
  Defaults to false.
  
* `network_loop_protection` - (Optional) Enables or disables network loop protection.
  Defaults to true.
  
* `pause_flood_protection` - (Optional) Enables Pause Flood Control protection.
  Defaults to true. 
  
* `igmp_timeout_interval` - (Optional) IGMP snooping idle time out intervals in seconds.
  Defaults to 260
  
* `mac_refresh_interval` - (Optional)  MAC Cache Fail Over refresh intervals in seconds.
  Defaults to 5

Telemetry Configuration supports the following:

* `enabled` - (Optional) Enables telemetry. Defaults to true.

* `sample_count` - (Optional) Telemetry sample count. Defaults to 12.

* `sample_interval` - (Optional) Telemetry sample interval in seconds. Defaults to 300.

Snmp Configuration supports the following: 

* `enabled` - (Optional) Enables SNMP v1 and v2. Defaults to true.

* `v3_enabled` - (Optional) Enables SNMP v3.  Defaults to false.  Must be set to true
  for a Virtual Connect SE 40Gb F8 Module.

* `read_community` - (Optional) Authentication string for read-only access.
  Defaults to public.

* `system_contact` - (Optional) Person to notify when system problems occur.

* `snmp_access` - (Optional) The access list allowed for GET operations.

* `trap_destination` - (Optional) The list of configured trap destinations.
  This can be specified multiple times. Trap Destination options specified below.

Trap Destination supports the following:

* `trap_destination` - (Required) The trap destination IP address or host name.

* `trap_format` - (Optional) The trap format (SNMP version) for this trap destination.
  Defaults to SNMPv1.

* `community_string` - (Optional)  Authentication string for the trap destination.
  Defaults to public. 

* `enet_trap_categories` - (Optional)  Filter the traps for this trap destination by the list of configured Ethernet traps.

* `fc_trap_categories` - (Optional)  Filter the traps for this trap destination by the list of configured Fibre Channel traps.

* `vcm_trap_categories` - (Optional) Filter the traps for this trap destination by the list of configured VCM traps.

* `trap_severities` - (Optional) Filter the traps for this trap destination by the list of configured severities

Interconnect Map Entry Template supports the following:

* `interconnect_type_name` - (Required) The interconnect type name for the bay.

* `bay_number` - (Required) The bay number to use. 

* `enclosure_index` - (Optional) The enclosure to use. Defaults to 1.

Uplink Set supports the following:

* `name` - (Required) A unique name for the resource.

* `network_uris` - (Optional) The set of network URIs. NOTE: for Ethernet Uplink Set Groups, 
  all Ethernet Networks must have unique VLAN IDs.

* `native_network_uri` - (Optional) The Ethernet native network URI.

* `network_type` - (Optional) The type of network. Defaults to Ethernet.

* `mode` - (Optional) The Ethernet uplink failover mode. Defaults to Auto.

* `lacp_timer` - (Optional) The LACP timer. Defaults to Long.

* `logical_port_config` - (Optional) The detailed configuration properties for the uplink ports.
  Logical Port Config is documented below.

Logical Port Config supports the followings:

* `bay_num` - (Required) The bay number to use. 

* `port_num` - (Required) The single port number to use.

* `enclosure_num` - (Optional) The enclosure to use. Defaults to 1.

* `primary_port` - (Optional) The Ethernet primary failover port. Defaults to false.

* `desired_speed` - (Optional) The port speed you prefer it to use. Defaults to Auto.

* `desired_fec_mode` - (Optional)  The desire FEC mode of logical port.

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource shall be assigned.
It is meaningful at resource creation time, during resource update, and it is included on resource retrieval as well.


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `uri` - The URI of the created resource.

* `eTag` - Entity tag/version ID of the resource.
