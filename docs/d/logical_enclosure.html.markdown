---
layout: "oneview"
page_title: "Oneview: logical_enclosure"
sidebar_current: "docs-logical_enclosure"
description: |-
 Gets information about an existing logical_enclosure.
---

# oneview\_logical_enclosure

Use this data source to access the attributes of a Logical Enclosure.

## Example Usage

```hcl
data "oneview_logical_enclosure" "test" {
 name = "Test logical enclosure"
}

output "oneview_logical_enclosure_value" {
 value = "${data.oneview_logical_enclosure.test.uri}"
}
```

## Argument Reference

* `name` - (Required) Name of logical enclosure

## Attributes Reference

* `ambient_temperature_mode` - The environment in which the logical enclosure should be optimized to operate.If not specified default value Standard is assumed

* `category` - Used to identify the kind of resource.

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource is assigned.

* `created` - Date and time when the resource is created.

* `delete_failed` - Indicator that the deletion of a logical enclosure failed.

* `type` - Type of the resource.

* `deployment_mode` - Specifies the OS deployment network configuration.

* `deployment_network_uri` - Specify the URI of the network used for OS deployment.

* `manage_os_deployment` - Indicate whether the OS deployment is enabed.

* `deployment_cluster_uri` - URI of the deployment cluster.

* `deployment_manager_settings` - The settings of the deployment manager for this Logical Enclosure.

* `description` - The description of the enclosure.

* `enclosure_group_uri` - URI of the enclosure group associated with the logical enclosure.

* `enclosure_uris` - A set of enclosure URIs associated with the logical enclosure.

* `firmwire_base_line_uri` - The URI of the firmware baseline to apply to the enclosure.

* `firmware_update_on` - Option that speicifies the component type within the enclosure which has to be updated.

* `force_install_firmware` - Force installation of firmware even if same or newer version is installed.

* `logical_interconnect_update_mode` - User can request a disruptive or non-disruptive mode of update for the logical interconnects.

* `update_firmware_on_unmanaged_interconnect` - User can indicate whether or not to update unmanaged interconnects within the logical enclosure.

* `validate_if_li_firmware_update_is_non_disruptive` -  User can validate whether the logical interconnect firmware update will cause disruption to the traffic that the interconnects are carrying.

* `ip_addressing_mode` - Manage IPv4 address allocation for interconnects and  device bay management processors.

* `dns_servers` - A list of DNS servers for the IPv4Range.

* `domain` - The domain of the IPv4Range.

* `gateway` - The gateway of the IPv4Range.

* `ip_range_uri` - The URI of the IPv4Range resource .

* `subnet_mask` - The submnet mask of the IPv4Range.

* `power_mode` - Power mode of the logical enclosure.

* `scaling_state` - Current resource state of the logical enclosure.

* `scopes_uri` - The URI for the resource scope assignment.

* `uri` - The canonical URI of the resource.

