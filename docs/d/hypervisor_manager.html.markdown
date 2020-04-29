---
layout: "oneview"
page_title: "Oneview: hypervisor_manager"
sidebar_current: "docs-hypervisor_manager"
description: |-
 Gets information about an existing hypervisor_manager.
---

# oneview\_hypervisor\_manager

Use this data source to access the attributes of a hypervisor manager.

## Example Usage

```
data "oneview_hypervisor_manager" "test" {
 name = "Test hypervisor manager"
}

output "oneview_hypervisor_manager_value" {
 value = "${data.oneview_hypervisor_manager.test.uri}"
}
```

## Argument Reference

* `name` -  Host name or IP of this hypervisor manager.

## Attributes Reference

* `available_dvs_versions` -   Available distributed switch versions on this hypervisor manager.

* `category` -  Resource category used for authorization which is "hypervisor-managers".

* `description` - Description for this hypervisor manager.

* `display_name` -  Indicates the display name of the Hypervisor manager.

* `e_tag` -  Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resource.

* `hypervisor_type` -  Hypervisor type which can be Vmware or HyperV.

* `password` -   Password that can be used to authenticate with this hypervisor manager.

* `port` -   Port number of the hypervisor manager service.Default set to 443

* `preferences` - Default preferences to be used for cluster profiles with this hypervisor manager

* `refresh_state` -  Indicates if the resource is currently refreshing.

* `resource_paths` -   Represents inventory paths in the hypervisor manager where cluster can be added

* `state` - Current state of the resource. Valid values include Connected, Disconnected, Configuring and Error.

* `state_reason` -  Indicates the reason the resource in its current state.

* `status` - Current status of this resource. Valid values include Unknown, OK, Disabled, Warning, Critical.

* `type` -   Uniquely identifies the type of the JSON object

* `uri` - The URI of the resource.

* `username` - (Required) The username to  authenticate the hypervisor manager.
