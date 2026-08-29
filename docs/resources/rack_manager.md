---
layout: "oneview"
page_title: "Oneview: rack_manager"
sidebar_current: "docs-oneview-rack-manager"
description: |-
  Adds a Rack Manager.
---

# oneview\_rack\_manager

Adds a rack manager. Returns the rack manager URI on successful addition.

## Example Usage

```js
resource "oneview_rack_manager" "RM" {
   hostname               = "1.1.1.1"
   username               = "user"
   password               = "password"
}
```

## Argument Reference

The following arguments are supported: 
* `force` - The default is false. Specify true to force the addition and take ownership away from another manager.
* `hostname` - (Required) Hostname identifies the RMC of the rack manager to be added. It can be specified as either a hostname or IPv4 address.

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource shall be assigned.
It is meaningful at resource creation time, during resource update, and it is included on resource retrieval as well.
* `password` - (Required) Password for the specified user name.
* `username` - (Required) RMC administrator user name (e.g. Administrator).


- - -


