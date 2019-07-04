---
layout: "oneview"
page_title: "Oneview: logical_interconnect"
sidebar_current: "docs-oneview-logical-interconnect"
description: |-
 Imports a Logical Interconnect from appliance.
---

# oneview\_logical\_interconnect

Import a logical interconnect

## Example Usage

```js
resource "oneview_logical_interconnect" "default" {
 # Empty body
}
```
## Terraform Command to import

	terraform import oneview_logical_interconnect.default <logical-interconnect-name>
	
## Argument Reference

The following arguments are supported: 

* `logical-interconnect-name` - (Required) A unique name for the resource as per the appliance.

- - -
description: |-
  Updates Logical Interconnect
---

# oneview\_logical\_interconnect

Update logical interconnect.

## Example Usage

```js
resource "oneview_logical_interconnect" "default" {
  uri = "rest/logical-interconnects/d0432852-28a7-4060-ba49-57ca973ef6c2"
  update_string = "update_type"
}
```

## Argument Reference

The following arguments are supported: 

* `update_type` - (Required) Type of update of Logical Interconnect.

	| NO |        Type of Update                          |   Update String               |
	|----|------------------------------------------------|-------------------------------|
	|  1 |`UpdateLogicalInterconnectConsistentStateById`  |'updateComplianceById'         |

- - -

* `uri`	- Unique Resource Identifier for a resource.
