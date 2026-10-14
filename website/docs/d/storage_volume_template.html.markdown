---
layout: "oneview"
page_title: "Oneview: storage_volume_templates"
sidebar_current: "docs-oneview-storage-systems"
description: |-
  Gets information about an existing Storage volume template.
---

# oneview\_storage\_syatem

Use this data source to access the attributes of a storage system.

## Example Usage

```hcl
data "oneview_storage_volume_template" "test" {
 name = "Test storage Volume Template"
}

output "oneview_storage_volume_template_value" {
 value = "${data.oneview_storage_volume_template.test.root_template_uri}"
}
```

## Argument Reference

The following arguments are supported: 

* `name` - (Required)  Display name for the resource.

## Attributes Reference

* `category` - Resource category used for authorizations and resource type groupings.

* `compatible_storage_systems_uri` - Uri to return a list of storage systems that supports this storage template.

* `description` -  Brief description of the resource.

* `eTag` - Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resource.

* `family` - Storage family that the storage volume template is associated with.

* `is_root` - Describes if the storage volume template is a user template or a root template.

* `name` - Display name for the resource.

* `tp_name`, `tp_storage_pool`, `tp_size`, `tp_provisioning_type`, `tp_snapshot_pool`, `tp_data_transfer_limit`, `tp_is_deduplicated`,
	`tp_is_encrypted`, `tp_is_pinned`, `tp_iops_limit`, `tp_folder`, `tp_template_version`, `tp_performance_policy`, 
	`tp_volume_set`, `tp_description`, `tp_is_adaptive_optimization_enabled`, `tp_is_compressed`, `tp_data_protection_level`,
	`tp_is_shareable` - The properties that define the template. The set of properties match those that are returned from the root template for the storage system 
						associated with the target storage pool.

* `root_template_uri` -  Uri of the root template that the storage volume template is associated with.

* `scopes_uri` -  URI of list of scopes for this storage template.

* `state` - Current state of the resource.

* `status` - Overall health status of the resource.

* `storage_pool_uri` - Uri of the storage pool that the storage volume template is associated with.

* `type` -  Uniquely identifies the type of the JSON object.

* `uri` - The canonical URI of the resource.

* `uuid` - Unique identifier for the storage volume template.

* `version` -  Version of the storage volume template.