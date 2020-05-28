---
layout: "oneview"
page_title: "Oneview: hypervisor_cluster_profile"
sidebar_current: "docs-oneview-hypervisor-cluster profile"
description: |-
  Creates a hypervisor cluster profile.
---

# oneview\_hypervisor\_cluster\_profile

Creates an hypervisor cluster profile.

## Example Usage

```js
resource "oneview_hypervisor_cluster_profile" "default" {
  "type"="HypervisorClusterProfileV3",
    "name"="Cluster7",
    "description"="",
    "hypervisor_type"="Vmware",
    "hypervisor_manager_uri"="/rest/hypervisor-managers/063055b5-4703-4b0c-8aea-60b23c2de157",
    "path"="DC2",
    "hypervisor_cluster_settings"={  
                                  "type"="Vmware",
                                  "drs_enabled"=true,
                                  "ha_enabled"=false,
                                  "multi_nic_v_motion"=false,
                                  "virtual_switch_type"="Standard"
                               },
    "hypervisor_host_profile_template"={  
        "server_profile_template_uri"="/rest/server-profile-templates/278cadfb-2e86-4a05-8932-972553518259",
        "deployment_plan"={  
        "deployment_plan_uri"="",
        "server_password"="",
        }
        "host_prefix"="Test-Cluster-host"
     },
}
```

## Argument Reference

The following arguments are supported: 

* `name` - (Required) A unique name for the resource.
* `type` - (Required) Indicates the display name of the Hypervisor cluster profile.


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `add_host_requests` - Request specifications for adding host to cluster.

* `category` - category Identifies the resource type 

* `compliance_state` -  Cluster Profile Compliance State.

* `created` - Date and time when the resource was created

* `description` -  Description of the resource.

* `e_tag` - ETAG or version ID of this resource.

* `hypervisor_cluster_settings` - User preferences for the hypervisor cluster profile.

* `hypervisor_cluster_uri` - URI of the hypervisor cluster resource.

* `hypervisor_host_profile_template` - Template for the hypervisor host profile.

* `hypervisor_host_profile_uris` - URIs of the host profiles

* `hypervisor_manager_uri` - URI of the hypervisor manager in which the cluster will be created.

* `hypervisor_type` - Type of the hypervisor in the cluster.

* `ip_pools` -  IP Pool configuration. This is used to configure IP and related attributes on the hosts.

* `mgmt_ip_settings_override` - Common management IP settings like subnet, gateway used in the case of user provided management IP.

* `modified` - Timestamp when this resource was last modified.

* `name` -  Name of the resource.

* `path` -  Name of the resource.

* `refresh_state` -  Indicates if the resource is currently refreshing.

* `scopes_uri` - The URI for the resource scope assignments

* `shared_storage_volumes` - set of storage volumes, which will be considered as cluster volumes and source type for these volumes will be cluster profile.

* `state` -  Current state of the resource. Valid values include New, Configuring, Error, Active, Removed, Refreshing and NonCompliant.

* `state_reason` - Indicates the reason the resource in its current state.

* `status` - Current status of this resource. Valid values include Unknown, OK, Disabled, Warning, Critical.