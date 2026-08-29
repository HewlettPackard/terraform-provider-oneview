---
layout: "oneview"
page_title: "Oneview: scope"
sidebar_current: "docs-oneview-scope"
description: |-
  Creates a scope.
---

# oneview\_scope

Creates a scope.

## Example Usage

```js
resource "oneview_scope" "default" {
  name = "test-scope"
  type = "ScopeV3"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A unique name for the resource.

* `type` - (Required) Identifies the Scope type.

- - -

* `Description` - (Optional) A description of the Scope.

* `added_resource_uris` - (Optional)  List of resource URIs to be added when creating or editing a scope. It is only meaningful at resource creation time and while updating. It is relevant during retrieval as well.

* `removed_resource_uris` - (Optional)  List of resource URIs to be removed. It is only meaningful while updating.

* `initial_resource_uris` - (Optional) A list of URIs of the scopes to which the resource is assigned. It is only meaningful at resource creation time and while updating. It is relevant during retrieval as well. 

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `uri` - The URI of the created resource.

* `eTag` - Entity tag/version ID of the resource.
