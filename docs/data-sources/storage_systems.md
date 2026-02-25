---
layout: "oneview"
page_title: "Oneview: storage_systems"
sidebar_current: "docs-oneview-storage-systems"
description: |-
  Creates a storage-system.
---

# oneview\_storage\_syatem

Use this data source to access the attributes of a storage system.

## Example Usage

```hcl
data "oneview_storage_system" "test" {
 name = "Test storage system"
}

output "oneview_storage_system_value" {
 value = "${data.oneview_storage_system.test.uri}"
}
```

## Argument Reference

The following arguments are supported: 

* `name` - (Required)  Display name for the resource.

## Attributes Reference

* `hostname` - IP address or hostname of the storage system. 

* `credentials` - Credentials used to authenticate with the storage system.

* `username` - User name for the storage system.

* `password` - Password for the specified user.

* `family` - Family of the device (e.g. StoreServ or StoreVirtual)

* `category` - Resource category used for authorizations and resource type groupings.

* `description` - Brief description of the resource.

* `state` - Current state of the resource.

* `status` - Overall health status of the resource.

* `eTag` - Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resource.

* `uri` - The canonical URI of the resource.

* `storage_pools_uri` - URI to get a list of pools related to the storage system.

* `total_capacity` -  Total capacity of the storage system in bytes.

* `ports` - A list of target ports for the storage system.

* `storage_system_device_specific_attributes` - Properties specific to the storage system type for a storage system..

* `managed_pool` - Managed pools.