---
layout: "oneview"
page_title: "Oneview: appliance_ssh_access"
sidebar_current: "docs-appliance_ssh_access"
description: |-
 Gets information about an existing appliance_ssh_access.
---

# oneview\_appliance\_ssh\_access

Use this data source to access the attributes of appliance ssh access.

## Example Usage

```hcl
data "oneview_appliance_ssh_access" "test" {
}

output "oneview_appliance_ssh_access_value" {
 value = "${data.oneview_appliance_ssh_access.test.uri}"
}
```

## Argument Reference

No argument is required to retrive the data source of existing appliance ssh access.

## Attributes Reference

* `allowSshAccess` -  Indicates whether the appliance accepts SSH connections, including access to the maintenance console.

* `category` - Identifies the resource type.

* `created` - Date and time when the resource was created.

* `etag` - Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resource.

* `modified` - Date and time when the resource was last modified.

* `type` - Uniquely identifies the type of the JSON object.

* `uri` - The URI of the resource.
