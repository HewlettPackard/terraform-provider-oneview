---
layout: "oneview"
page_title: "Oneview: hypervisor_manager"
sidebar_current: "docs-oneview-hypervisor-manager"
description: |-
  Creates a hypervisor manager.
---

# oneview\_hypervisor\_manager

Creates an hypervisor manager.

## Example Usage

```js
resource "oneview_hypervisor_manager" "default" {
  name = "test-hypervisor-manager"
}
```

## Argument Reference

The following arguments are supported: 

* `username` - (Required) An  username to authenticate the hypervisor manager.
* `type` - (Required) Uniquely identifies the type of the JSON object.


- - -
* `name` - (optional) Host name or IP of this hypervisor manager. 
* `password` - (Optional) Password that can be used to authenticate with this hypervisor manager. 
* `port` - (Optional) Port number of the hypervisor manager service.Default set to 443. 
* `display_name` - (Optional) A unique name for the resource.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `available_dvs_versions` -   Available distributed switch versions on this hypervisor manager.

* `category` -  Resource category used for authorization which is "hypervisor-managers".

* `description` - Description for this hypervisor manager.

* `display_name` -  Indicates the display name of the Hypervisor manager.

* `e_tag` -  Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resource.

* `hypervisor_type` -  Hypervisor type which can be Vmware or HyperV.

* `name` -  Host name or IP of this hypervisor manager.

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

* `username` -   User name that can be used to authenticate with this hypervisor manager.
