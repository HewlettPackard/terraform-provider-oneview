---
layout: "oneview"
page_title: "Oneview: SNMPv3 User"
sidebar_current: "docs-snmpv3_user"
description: |-
 Gets information about an existing snmpv3_user.
---

# oneview\_appliance\_snmpv3\_user

Use this data source to access the attributes of a snmpv3_user.

## Example Usage

```hcl
data "oneview_appliance_snmpv3_user" "test" {
 id_field = "67003649-af34-4a92-a46a-137855ddc8f7"
}

output "oneview_snmpv3_user" {
 value = "${data.oneview_snmpv3_user.test}"
}
```

## Argument Reference

* `user_name` - (Required) Thes smp v3 user username. The USM User Name

## Attributes Reference

* `authentication_passphrase` - The passphrase used for authentication.

* `authentication_protocol` - The protocol used for authentication.

* `category` - The category is used to help identify the kind of resource.

* `created` - Date and time when the resource was created.

* `etag` - Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resource.
  
* `modified` - Date and time when the resource was last modified

* `user_id` - The USM User Id
  
* `privacy_passphrase` - The passphrase used to ensure privacy. Supported length for passphrase is 8 - 32 characters.
  
* `privacy_protocol` - An indication of whether messages sent to be protected from disclosure and if so, the type of privacy protocol which is used
  
* `security_level` - The Level of Security that determines if the message needs to be protected from disclosure and if the message needs to be authenticated.

* `type` - Type of the resource.

* `uri` - The URI of the resource.

* `user_name` - The USM User Name

