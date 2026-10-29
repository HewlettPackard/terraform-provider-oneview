---
layout: "oneview"
page_title: "Oneview: Firmware Drivers"
sidebar_current: "docs-firmware_drivers"
description: |-
 Gets information about an existing firmware_drivers.
---

# oneview\_firmware\_drivers

Use this data source to access the attributes of a firmware baseline.

## Example Usage

```hcl
data "oneview_firmware_drivers" "test" {
 name= "HPE Synergy Service Pack"
 version="SY-2021.11.01"
}

output "oneview_firmware_drivers_value" {
 value = "${data.oneview_firmware_drivers.test}"
}
```

## Argument Reference

* `name` - (Required) Display name for the resource. for e.g. "HPE Synergy Service Pack"
* `version` - (Required) The version of the firmware bundle. For e.g. "SY-2021.11.01"

## Attributes Reference

* `baseline_short_name` - The short name is the 'SPP ' string plus the version of the firmware bundle.

* `bundle_size` - The size of the firmware baseline resource in bytes.

* `bundle_type` - The typical bandwidth of a network connection, expressed in Mbps.

* `category` - Identifies the resource type.

* `description` - Brief description of the resource.

* `etag` - Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resource.

* `esxi_os_driver_meta_data` -  Relative Path to retrieve the zipFiles present in vmw directory of SPP.

* `fw_components` - The list of the components in the firmware baseline resource.
  
  * `component_version` - The version of this component.

  * `file_name` - The actual file name of this component.

  * `name` - The name of this component.

  * `sw_key_name_list` - A list of the software key name for this component.

* `hotfixes` - List of hotfixes the current baseline contains. This array is empty for ServicePack and hotfixes. Only Custom bundles will have this array elements.

  * `hotfix_name` - The name of the hotfix.

  * `release_date` - The release date of the hotfix.

  * `resource_id` - Unique ID of the hotfix.

* `hpsum_version` - The SUM version of the firmware baseline resource.
  
* `iso_file_name` - The full iso file name which includes the '.iso' extension.

* `last_task_uri` - The most recent task tracker uri; for internal use only.

* `locations` - List of locations where the baseline exists. Contains ServicePack mirrors information.

* `mirrorlist` -  The list of ServicePack mirrors on the corresponding repositories present on the appliance.

* `parent_bundle` - Details of the firmware baseline from which the current custom baseline is created. This is empty for ServicePack and hotfixes. Only Custom bundles will have this value.

  * `parent_bundle_name` - The name of parent firmware bundle.

  * `release_date` - The release date of the parent firmware bundle.

  * `version` - The version of the parent firmware bundle.

* `release_date` - The release date of the firmware baseline resource in GMT time.

* `resource_id` - The resource id of the firmware baseline resource.

* `resource_state` - The resource state of the firmware baseline resource.

* `scopes_uri` - The URI for the firmware bundle resource scope assignments.

* `signature_file_name` -  The signature file name for the Gen10 and above hotfixes.

* `signature_file_required` - The value of this attribute will be true only for Gen10 and above hotfixes.

* `state` -  Current state of the resource.

* `status` - Overall health status of the resource.

* `supported_languages` - The languages supported by this firmware baseline resource.

* `supported_os_list` - Type of the resource.
  
* `sw_packages_full_path` -  Full path to retrieve the software packages in this firmware baseline resource; for internal use only.

* `uri` - The URI of the resource.

* `type` - Type of the resource.

* `uuid` - The unique id of the firmware baseline resource.

* `version` - The version of the firmware bundle

* `xml_key_name` - The baseline xml file name for internal use.