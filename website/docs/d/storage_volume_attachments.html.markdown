---
layout: "oneview"
page_title: "Oneview: storage_volume_attachments"
sidebar_current: "docs-oneview-storage-volume-attachments"
description: |-
  Creates a storage-volume attachment.
---

# oneview\_storage\_volume\_attachment

Use this data source to access the attributes of a storage volume attachment.

## Example Usage

```hcl
data "oneview_storage_volume_attachment" "test" {
 name = "Test storage volume attachment"
}

output "oneview_storage_volume_attachment_value" {
 value = "${data.oneview_storage_volume_attachment.test.uri}"
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

* `storage_system_uri` - URI to of the associated storage system.

* `storage_volume_uri` - URI of the Storage Volume the Storage Volume Attachment is connected to.

* `host` - Contains information used to descibe the host using this Storage Volume Attachment 

* `paths` - Contains information used to describe the data path(s) connecting the initiator to the Storage System target(s) 
