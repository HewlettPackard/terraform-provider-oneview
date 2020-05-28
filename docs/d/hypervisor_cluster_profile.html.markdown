---
layout: "oneview"
page_title: "Oneview: hypervisor_cluster_profile"
sidebar_current: "docs-hypervisor_cluster_profile"
description: |-
 Gets information about an existing hypervisor_cluster_profile.
---

# oneview\_hypervisor\_cluster\_profile

Use this data source to access the attributes of a hypervisor cluster profile.

## Example Usage

```hcl
data "oneview_hypervisor_cluster_profile" "test" {
 uri = "/rest/hypervisor-cluster-profiles/12343-as45677-56778af"
}

output "oneview_hypervisor_cluster_profile_value" {
 value = "${data.oneview_hypervisor_cluster_profile.test.type}"
}
```

## Argument Reference

* `uri` - (Required) The name of the hypervisor cluster profile uri.

## Attributes Reference

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

* `type` - Uniquely identifies the type of the JSON object