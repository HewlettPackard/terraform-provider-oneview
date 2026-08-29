---
layout: "oneview"
page_title: "Oneview: connection_templates"
sidebar_current: "docs-connection_templates"
description: |-
 Gets information about an existing connection_templates.
---

# oneview\_connection\_templates

Use this data source to access the attributes of a connection templates.

## Example Usage

```hcl
data "oneview_connection_templates" "test" {
 name = "defaultConnectionTemplates"
}

output "oneview_connection_templates_value" {
 value = data.oneview_connection_templates.test.uri
}
```

```hcl
data "oneview_connection_templates" "testingViaUri" {
 uri = "/rest/connection-templates/defaultConnectionTemplate" 
}

output "oneview_connection_templates_value" {
 value = data.oneview_connection_templates.testingViaUri.uri
}
```

## Argument Reference

* `name` - The name of the connection template.
* `uri` - The URI of the resource.

## Attributes Reference

* `bandwidth` - The bandwidth assigned to the connection-template.

* `maximumBandwidth` - The maximum bandwidth of a network connection, expressed in Mbps.

* `typicalBandwidth` - The typical bandwidth of a network connection, expressed in Mbps.

* `category` - Identifies the resource type.

* `description` - Brief description of the resource.

* `etag` - Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resource.

* `state` -  Current state of the resource.

* `status` - Overall health status of the resource.

* `uri` - The URI of the resource.

* `type` - Type of the resource.
