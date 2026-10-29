---
layout: "oneview"
page_title: "Oneview: logical_enclosure"
sidebar_current: "docs-oneview-logical-enclosure"
description: |-
  Creates a logical-enclosure.
---

# oneview\_logical\_enclosure

Creates a logical enclosure with firmware.

## Example Usage

```hcl
resource "oneview_logical_enclosure" "default" {
  name = "default-logical-enclosure"
  enclosure_uris = [data.oneview_enclosure1.uri,
                                 data.oneview_enclosure2.uri]
  enclosure_group_uri = data.oneview_enclosure_group.uri
  firmware {
	firmware_baseline_uri = "/rest/firmware-drivers/Synergy_Custom_SPP_2021_02_01_Z7550-97110"
	force_install_firmware = true
  }
}
```

Import a logical enclosure.

## Example Usage

```js
resource "oneview_logical_enclosure" "default" {
 # Empty body
}
```
## Terraform Command to import

        terraform import oneview_logical_enclosure.default <logical-enclosure-name>


Update logical enclosure.

## Example Usage

```js
resource "oneview_logical_enclosure" "default" {
        name = "default-logical-enclosure"
        enclosure_uris = [data.oneview_enclosure1.uri,
                                 data.oneview_enclosure2.uri]
        enclosure_group_uri = data.oneview_enclosure_group.uri
        update_type = "update-type"
	firmware {
		firmware_baseline_uri = "/rest/firmware-drivers/SPPGen9Snap3_2015_0221_71"
         	firmware_update_on = "EnclosureOnly"
         	force_install_firmware = false
         	validate_if_li_firmware_update_is_non_disruptive = false
         	logical_interconnect_update_mode = "Parallel"
         	update_firmware_on_unmanaged_interconnect = false
	}
}
```

## Argument Reference

The following arguments are supported: 

* `name` -(Required) A unique name for the resource.

* `enclosure_uris` -(Required) The set of uris associated with the enclosure.

* `enclosure_group_uri` - (Required) The uri of the enclosure group. 

---

* `firmware_baseline_uri` - The URI of the firmware baseline to apply to the enclosure.

* `firmware_update_on` - Option that speicifies the component type within the enclosure which has to be updated.

* `force_install_firmware` - Force installation of firmware even if same or newer version is installed.

* `logical_interconnect_update_mode` - User can request a disruptive or non-disruptive mode of update for the logical interconnects.

* `update_firmware_on_unmanaged_interconnect` - User can indicate whether or not to update unmanaged interconnects within the logical enclosure.

* `validate_if_li_firmware_update_is_non_disruptive` -  User can validate whether the logical interconnect firmware update will cause disruption to the traffic that the interconnects are carrying. 

* `update_type` - Type of update of Logical Enclosure.
* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource shall be assigned.
It is meaningful at resource creation time, during resource update, and it is included on resource retrieval as well.

* `update_type` - (Required) Type of update of Logical Enclosure.

 | NO  | Type of Update                    | Update String   |
 | --- | --------------------------------- | --------------- |
 | 1   | `UpdateLogicalEnclosure`          | 'update'        |
 | 2   | `UpdateFromGroupLogicalEnclosure` | 'updateByGroup' |

