---
layout: "oneview"
page_title: "Oneview: enclosure"
sidebar_current: "docs-enclosure"
description: |-
 Gets information about an existing enclosure.
---

# oneview\_enclosure

Use this data source to access the attributes of a enclosure.

## Example Usage

```hcl
data "oneview_enclosure" "test" {
 name = "Test enclsoure"
}

output "oneview_enclosure_value" {
 value = "${data.oneview_enclosure.test.uri}"
}
```

## Argument Reference

* `name` - (Required) The name of the enclsoure.

## Attributes Reference

* `active_oa_preferred_ip` - Preferred IP address for the enclosure's active OA.

* `asset_tag` -  Asset tag for the enclosure.

* `category` - Identifies the resource type.

* `description` - Brief description of the resource.

* `device_bay_count` - Number of device bays in the enclosure.

* `device_bays` -  List of device bays in the enclosure.

* `enclosure_group_uri` - URI for this enclosure's enclosure-group.

* `enclosure_type` - The type of the enclosure, eg, 'C7000' or 'SY12000' or 'SDX'.

* `etag` - Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resource.

* `fw_baseline_name` - The name of the current firmware baseline.

* `fw_baseline_uri` - The firmware baseline associated with this enclosure.

* `interconnect_bay_count` - Number of interconnect bays in the enclosure.

* `interconnect_bays` - List of interconnect bays in the enclosure.

* `is_fw_managed` - Flag indicating whether the firmware is managed.

* `licensing_intent` - The licensing policy for all the servers in the enclosure.

* `part_number` - Part number of the enclosure.

* `rack_name` - Name of the rack in which the enclosure resides.

* `refresh_state` -  Indicates if the resource is currently refreshing.

* `scopesUri` - The URI for the resource scope assignments.

* `serial_number` - Serial number of the enclosure.

* `standby_oa_preferred_ip` - Preferred IP address for the enclosure's standby OA.

* `state` - Current resource state of the enclosure.

* `state_reason` - Indicates the reason the resource in its current state.

* `status` - Overall health status of the enclosure.

* `type` - Type of the resource.

* `uri` - The URI of the resource.

* `uuid` -  UUID of the enclosure.

* `vcm_domain_id` - Domain ID of the enclosure.

* `vcm_domain_name` - Domain name of the enclosure.

* `vcm_mode` - Flag which indicates whether the enclosure is managed by an appliance.

* `vcm_url` -  URL of the enclosure's management appliance.