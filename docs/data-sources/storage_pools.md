---
layout: "oneview"
page_title: "Oneview: storage_pools"
sidebar_current: "docs-oneview-storage-pools"
description: |-
  Gets details of a storage-pool.
---

# oneview\_storage\_pool

Use this data source to access the attributes of a storage pool.

## Example Usage

```hcl
data "oneview_storage_pool" "test" {
 name = "Test storage pool"
}

output "oneview_storage_pool_value" {
 value = "${data.oneview_storage_pool.test.uri}"
}
```

## Argument Reference

The following arguments are supported: 

* `name` - (Required)  Display name for the resource.

## Attributes Reference

* `category` - Resource category used for authorizations and resource type groupings.

* `description` - Brief description of the resource.

* `state` - Current state of the resource.

* `status` - Overall health status of the resource.

* `eTag` - Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resource.

* `uri` - The canonical URI of the resource.

* `storage_system_uri` - URI of the storage system in which this storage pool is present.

* `total_capacity` -  Total capacity of the storage pool in bytes.

* `free_capacity` - Free capacity available from the storage pool in bytes

* `is_managed` -  Indicates whether or not the pool is managed 

* `storage_pool_device_specific_attributes` -  Device specific properties for the storage pool.

* `type` - Uniquely identifies the type of the JSON object.
