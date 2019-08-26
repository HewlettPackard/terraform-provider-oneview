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

* `name` -(Required) A unique name for the resource.

---

* `enclosure_uris` -(Required) The set of uris associated with the enclosure.

* `enclosure_group_uri` - (Required) The uri of the enclosure group. 

---
description: |-
 Imports a Logical Enclosure from appliance.
---

Import a logical enclosure.

## Example Usage

```js
resource "oneview_logical_enclosure" "default" {
 # Empty body
}
```
## Terraform Command to import

	terraform import oneview_logical_enclosure.default <logical-enclosure-name>
	
## Argument Reference

The following arguments are supported: 

* `logical-enclosure-name` - (Required) A unique name for the resource as per the appliance.

- - -
description: |-
  Updates Logical Enclosure
---

# oneview\_logical\_enclosure

Update logical enclosure.

## Example Usage

```js
resource "oneview_logical_enclosure" "default" {
  	name = "default-logical-enclosure"
  	enclosure_uris = ["${oneview_enclosure1.uri}", "${oneview_enclosure2.uri}"]
  	enclosure_group_uri = "${oneview_enclosure_group.uri}"
	update_type = "update-type"
}
```

## Argument Reference

The following arguments are supported: 

* `name` -(Required) A unique name for the resource.

---

* `enclosure_uris` -(Required) The set of uris associated with the enclosure.

* `enclosure_group_uri` - (Required) The uri of the enclosure group. 

* `update_type` - (Required) Type of update of Logical Enclosure.

	| NO |        Type of Update                          |   Update String               |
	|----|------------------------------------------------|-------------------------------|
	|  1 |`UpdateLogicalEnclosure`			              |'update'                       |
	|  2 |`UpdateFromGroupLogicalEnclosure`			      |'updateByGroup'                |

