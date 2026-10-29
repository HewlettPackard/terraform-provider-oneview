---
layout: "oneview"
page_title: "Oneview: appliance_ssh_access"
sidebar_current: "docs-appliance_ssh_access"
description: |-
 Updates appliance_ssh_access.
---

# oneview\_appliance\_ssh\_access

Configures the appliance ssh access.

## Example Usage

```js
resource "oneview_appliance_ssh_access" "sshaccess" {
   allow_ssh_access = false
}
```

## Argument Reference

The following arguments are supported: 

* `allow_ssh_access` - Indicates whether the appliance accepts SSH connections, including access to the maintenance console.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `category` - Identifies the resource type.

* `etag` - Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resource.

* `type` - Uniquely identifies the type of the JSON object.

* `uri` - The URI of the resource.
