---
layout: "oneview"
page_title: "Oneview: server_profile_template"
sidebar_current: "docs-server_profile_template"
description: |-
 Gets information about an existing server_profile_template.
---

# oneview\_server\_profile\_template

Use this data source to access the attributes of a Server Profile Template.

## Example Usage

```hcl
data "oneview_server_profile_template" "test" {
 name = "Test server profile template"
}

output "oneview_server_profile_template_value" {
 value = "${data.oneview_server_profile_template.test.uri}"
}
```

## Argument Reference

* `name` - (Required) A unique name for the resource.

* `affinity` - (Optional) This identifies the behavior of the server profiles created from this template when the server 
hardware is removed or replaced. This can be set to Bay or BayAndServer. 
This defaults to Bay.
  
* `network` - (Optional) Network connection to be configured for the server. Can be specified multiple times. 
Network configuration is specified below.
  
* `hide_unused_flex_nics` - (Optional) Hides flex nics that aren't in use.
  This defaults to true.

* `serial_number_type` - (Optional) Specifies the type of Serial Number and UUID to be programmed into the server ROM. 
The value can be 'Virtual' or 'Physical'. Changing this forces a new resource.
This defaults to 'Virtual'.
  
* `wwn_type` - (Optional) Specifies the type of WWN address to be programmed into the IO devices. The value can be 'Virtual' 
or 'Physical'. Changing this forces a new resource. 
This defaults to 'Virtual'.

* `mac_type` - (Optional) Specifies the type of MAC address to be programmed into the IO devices. The value can be 'Virtual'
or 'Physical'. Changing this forces a new resource.
This defaults to 'Virtual'.

* `uri` - The URI of the created resource.

* `eTag` - Entity tag/version ID of the resource.
