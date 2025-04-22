---
layout: "oneview"
page_title: "Oneview: rack_manager"
sidebar_current: "docs-rack_manager"
description: |-
 Gets information about an existing rack_manager.
---

# oneview\_rack\_manager

Use this data source to access the attributes of a rack manager.

## Example Usage

```hcl
data "oneview_rack_manager" "test" {
 name = "Test rack manager"
}

output "oneview_rack_manager_value" {
 value = "${data.oneview_rack_manager.test.uri}"
}
```

## Argument Reference

* `name` - (Required) The name of the rack manager.

## Attributes Reference

* `category` - Identifies the resource type.

* `description` - Brief description of the resource.

* `etag` - Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resource.

* `licensing_intent` -  The license type is automatically applied when you add a rack manager. Monitored rack managers use a HPE OneView Standard license.

* `id` - This is the id of the rack manager.

* `location` -  Location of the rack manager.

* `model` - The rack manager model name.

* `part_number` - This is the part number of the rack manager.

* `refresh_state` - This is the refresh state of the rack manager.

* `remote_support_uri` -  Remote support URI.

* `scopes_uri` - The URI for the resource scope assignments.

* `serial_uumber` - Serial number of the rack manager.

* `status` - Overall health status of the resource.

* `support_data_collection_state` - Current remote support data collection state of the rack manager.

* `support_data_collection_type"` - Current remote support data collection type of the rack manager.

* `support_data_collections_uri` - Uri To dataCollection.

* `support_state` -  Current remote support state of the rack manager.

* `uri` - The URI of the resource.

* `type` - Type of the resource.

* `initial_scope_uris` - (Optional) A list of URIs of the scopes to which the resource is assigned.

