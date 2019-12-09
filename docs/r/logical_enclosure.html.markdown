---
layout: "oneview"
page_title: "Oneview: logical_enclosure"
sidebar_current: "docs-oneview-logical-enclosure"
description: |-
  Creates a logical-enclosure.
---

# oneview\_logical\_enclosure

Creates a logical enclosure.

## Example Usage

```js
resource "oneview_logical_enclosure" "default" {
  name = "default-logical-enclosure"
  enclosure_uris = ["${oneview_enclosure1.uri}", 
                                 "${oneview_enclosure2.uri}"]
  enclosure_group_uri = "${oneview_enclosure_group.uri}"
}
```

## Argument Reference

The following arguments are supported: 

* `name` -(Optional) A unique name for the resource.

---

* `enclosure_uris` -(Optional) The set of uris associated with the enclosure.

* `enclosure_group_uri` - (Optional) The uri of the enclosure group. 

