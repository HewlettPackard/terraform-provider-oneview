---
layout: "oneview"
page_title: "Oneview: uplink_set"
sidebar_current: "docs-uplink_set"
description: |-
 Gets information about an existing uplink_set.
---

# oneview\_uplink_set

Use this data source to access the attributes of a Uplink Set.

## Example Usage

```hcl
data "oneview_uplink_set" "test" {
 name = "Test"
}

output "oneview_uplink_set_value" {
 value = "${data.oneview_uplink_set.test.uri}"
}
```

## Argument Reference

* `name` - (Required) Name of the uplink set whose details need to be retrieved.

## Attributes Reference

* `"logical_interconnect_uri"` - URI of the LI the uplink set is configured for.

* `"network_uris"` - Networks associated with teh uplink set.

* `"fc_network_uris"` - FC Networks associated with teh uplink set.

* `"fcoe_network_uris"` - FCOE Networks associated with teh uplink set.

* `type` - Type of the resource.

* `connection_mode`

* `uri` - URI of the uplink set

* `connection_mode` - Connection mode of the uplink set

* `port_config_infos` - Port configuration information of the uolink set

* `network_type` - Network type of the uplink set

* `ethernet_network_type`

* `lacptimer` - The LACP timer value

* `"manual_login_redistribution_state"` - The current state of Manual Login Redistribution

* `native_network_uri` -  The network that is designated as the native network. All external untagged traffic which ingresses the uplink set will be placed on this network. 

* `reachability` -  The reachability of the logical uplink. Possible values are redundantly reachable, reachable, unreachable or unknown. 
