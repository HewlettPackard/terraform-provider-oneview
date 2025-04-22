---
layout: "oneview"
page_title: "Oneview: verison"
sidebar_current: "docs-oneview-version"
description: |-
   Gets information about the range of supported API version .
---

# oneview\_version

Use this data source to access the attributes of version

## Example Usage

```js
data "oneview_version" "ver"  {

}
```

## Argument Reference

No argument is needed.
- - -


## Attributes Reference

The following computed attributes are exported:

* `current_version` - The latest supported API version.

* `minimum_version` - The minimum supported API version

* `id` - The appliance ip address.
