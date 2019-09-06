---
layout: "oneview"
page_title: "Oneview: storage-volume"
sidebar_current: "docs-storage-volume"
description: |-
 Gets information about an existing storage_volume.
---

# oneview\_storage_volume

Use this data source to access the attributes of a Storage Volume.

## Example Usage

```hcl
data "oneview_volume" "test" {
 name = "Test volume"
}

output "oneview_volume_value" {
 value = "${data.oneview_volume.test.uri}"
}
```

## Argument Reference

* `name` - (Required) Name of logical enclosure

## Attributes Reference

* `category` - Used to identify the kind of resource.

* `created` - Date and time when the resource is created.

* `device_specific_attributes` - Attributes specific for the volume.

* `type` - Type of the resource.

* `description` - The description of the enclosure.

* `storage_pool_uri` - Uri for the storage pool

* `scopes_uri` - The URI for the resource scope assignment.

* `uri` - The canonical URI of the resource.

