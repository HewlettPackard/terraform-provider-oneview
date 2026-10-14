---
layout: "oneview"
page_title: "Oneview: SNMPv3 User"
sidebar_current: "docs-snmpv3_user"
description: |-
  Creates a snmpv3_user.
---

# oneview\_snmpv3\_user

Creates a SNMPv3 Trap Destinations.

## Example Usage

```js
resource "oneview_appliance_snmpv3_user" "snmpuser" {
  user_name                 = "user"
  security_level            = "Authentication and privacy"
  authentication_protocol   = "SHA1"
  authentication_passphrase = "authPass"
  privacy_protocol          = "AES-128"
  privacy_passphrase        ="12345600"
}
```

## Argument Reference

The following arguments are supported: 

* `user_name` - The USM User Name. (Required)

* `security_level` - The Level of Security that determines if the message needs to be protected from disclosure and if the message needs to be authenticated.

* `authentication_protocol` - The protocol used for authentication.

* `authentication_passphrase` - The passphrase used for authentication.

* `privacy_protocol` - An indication of whether messages sent to be protected from disclosure and if so, the type of privacy protocol which is used.

* `privacy_passphrase` - The passphrase used to ensure privacy. Supported length for passphrase is 8 - 32 characters.


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

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