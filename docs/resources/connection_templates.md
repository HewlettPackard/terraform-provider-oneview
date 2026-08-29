---
layout: "oneview"
page_title: "Oneview: connection_templates"
sidebar_current: "docs-oneview-connection-templates"
description: |-
 Imports a connection template from appliance.
---

# oneview\_connection\_templates

Import a connection template

## Example Usage

```js
resource "oneview_connection_templates" "default" {
 # Empty body
}
```
## Terraform Command to import

	terraform import oneview_connection_templates.default <connection-template-uri>

## Argument Reference

The following arguments are supported: 

* `uri` - (Required) A unique uri for the resource as per the appliance.

- - -
description: |-
  Updates Connection Template
---

# oneview\_connection\_templates

Update Connection Template.

## Example Usage

```js
resource "oneview_connection_templates" "default" {
  name = "renamed-connection-template"
  bandwidth {
      maximum_bandwidth = 8000
      typical_bandwidth = 2500
  }
}
```

## Argument Reference

The following arguments are supported: 

* `uri` - (Required) The uri for the connection template.

* `maximumbandwidth` - (Required) The maximum bandwidth of a network connection, expressed in Mbps.

* `typicalBandwidth` - (Required) The typical bandwidth of a network connection, expressed in Mbps.
