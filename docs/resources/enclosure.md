---
layout: "oneview"
page_title: "Oneview: enclosure"
sidebar_current: "docs-oneview-enclosure"
description: |-
  Creates an enclosure.
---

# oneview\_enclosure

Creates an enclosure.

## Example Usage

```js
resource "oneview_enclosure" "default" {
  host_name = "Hostname_or_IP"
  licensing_intent = "Oneview"
  password = "password"
  user_name = "username"
}
```

## Argument Reference

The following arguments are supported: 

* `host_name` - (Required) Hostname identifies the primary or standby OA of the enclosure to be added. 

* `licensing_intent` - (Required)  The licensing policy for all the servers in the enclosure.

* `password` - (Required) Password for the specified user name.

* `user_name` - (Required) OA administrator user name (e.g. Administrator).

---

* `enclosure_group_uri` - (Optional) URI of an existing enclosure group to which the enclosure should be added. 

* `enclosure_uri` - (Optional) The URI for an enclosure that was previously added. When re-adding an enclosure, either enclosure URI or hostname can be used to identify the enclosure. 

* `firmware_baseline_uri` - (Optional) The URI of the firmware baseline to apply to the enclosure.

* `force` - (Optional) Use this optional flag with caution, because force-adding an enclosure makes it unmanagable by any other system managing it and removes any existing configuration. 

* `force_install_firmware` - (Optional) Force installation of firmware even if same or newer version is installed. 

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource shall be initially assigned. 

* `scopesUri` - (Optional) A list of URIs of the scopes to which the resource is assigned. 

* `state` - (Optional) Specifies the state of the enclosure to be added. Set to "monitored" to inventory and monitor the hardware.

* `update_firmware_on` - (Optional) Specifies whether the firmware baseline should be applied to the logical interconnect in addition to the enclosure.
