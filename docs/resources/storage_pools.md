---
layout: "oneview"
page_title: "Oneview: storage_pools"
sidebar_current: "docs-oneview-storage-pools"
description: |-
  Updates a storage-pool.
---

# oneview\_storage\_pool

Updates an existing storage pool.

## Example Usage

```js
resource "oneview_storage_pool" "test" {
 name = "Test storage pool"
 is_manged = false
}
```

## Argument Reference

The following arguments are supported: 

* `name` - (Required)  Display name for the resource.

* `description` - Brief description of the resource.

* `uri` - The canonical URI of the resource.

* `total_capacity` -  Total capacity of the storage pool in bytes.

* `free_capacity` - Free capacity available from the storage pool in bytes

* `is_managed` -  Indicates whether or not the pool is managed 

* `storage_pool_device_specific_attributes` -  Device specific properties for the storage pool.
