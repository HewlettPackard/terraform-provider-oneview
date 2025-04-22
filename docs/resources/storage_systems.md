---
layout: "oneview"
page_title: "Oneview: storage_systems"
sidebar_current: "docs-oneview-storage-systems"
description: |-
  Creates a storage-system.
---

# oneview\_storage\_syatem

Creates a storage system.

## Example Usage

```js
resource "oneview_storage_system" "default" {
  hostname = "ip_address_or_hostname"
  family = "StoreServ_or_StoreVirtual"
  username = "username"
  password = "password"
}
```

## Argument Reference

The following arguments are supported: 

* `hostname` - (Required) IP address or hostname of the storage system.

* `username` - (Required)  User name for the storage system.

* `password` - (Required)  Password for the specified user.

* `family` - (Required)  Either StoreVirtual or StoreServ to indicate the type of storage system being managed.

---

* `name` - (Optional)  Used to indicate the name of the StoreVirtual cluster to manage.

