---
layout: "oneview"
page_title: "Oneview: storage-volume"
sidebar_current: "docs-oneview-storage-volume"
description: |-
  Creates a storage-volume.
---

# oneview\_storage\_volume

Creates a storage volume.

## Example Usage

```js
resource "oneview_volume" "default" {
  name = "default-storage-volume"
  properties = {
    "name" = "default-storage-volume"
    "storage_pool"= "storage-pool-uri",
    "size"= size,
    "provisioning_type"= "provisioning-type",
  }
  template_uri= "storage-template-uri",
  is_permanent= true,
}
```

## Argument Reference

The following arguments are supported: 

* `name` -(Required) A unique name for the resource.

---

* `storage_pool` -(Required) The uri of storage pool.

* `size` -(Required) Provisioning size for the volume.

* `provisioning_type` - (Required) Type of provisioning for the volume.

* `template_uri` -(Required) The uri of storage pool.

---
description: |-
 Imports a Storage Volume from appliance.
---

Import a storage volume.

## Example Usage

```js
resource "oneview_volume" "default" {
 # Empty body
}
```
## Terraform Command to import

	terraform import oneview_volume.default <storage-volume-name>
	
## Argument Reference

The following arguments are supported: 

* `storage-volume-name` - (Required) A unique name for the resource as per the appliance.
