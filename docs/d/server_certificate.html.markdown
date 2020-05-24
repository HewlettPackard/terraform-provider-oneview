---
layout: "oneview"
page_title: "Oneview: server_certificate"
sidebar_current: "docs-server_certificate"
description: |-
 Gets information about an existing server_certificate.
---

# oneview\_server\_certificate

Use this data source to access the attributes of a server certificate.

## Example Usage

```hcl
data "oneview_server_certificate" "test" {
alias_name = "Test server certificate"
}

output "oneview_server_certificate_value" {
 value = "${data.oneview_server_certificate.test.uri}"
}
```

## Argument Reference
Note:  Either alias_name or remote_ip  needs to be provided.
* `remote_ip` - Remote ip is used to fetch the server certificate of a remote host. 
* `alias_name` -  Retrieves the device or server certificate, already trusted in the appliance, with the specified aliasName
## Attributes Reference

* `alternative_name` -  An optional entry that contains additional names that apply to the owner of the certificate, possibly including additional email addresses, DNS names, IP addresses, or other identifiers.

* `base64_data` -  Encrypted content of the SSL certificate.

* `basic_constraints` - Basic constraints to control the usage of the certificate.

* `category` -  Identifies the resource type

* `common_name` - Common name of the remote appliance certificate.

* `contact_person` - Name of the contact person present in the remote appliance certificate.

* `country` -  Country code present in the remote appliance certificate.

* `created` - Date and time when the resource was created

* `description` - Brief description of the resource

* `dn_qualifier` - DN qualifier present in the remote appliance certificate.

* `e_tag` -  Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resource.

* `email` - Email address present in the remote appliance certificate in the format name@domain.

* `enhanced_key_usage` - Collection of object identifiers that are allowed to use the key associated with the remote appliance certificate.

* `expires_in_days` -  Number of days until the expiration of the certificate. Internet CA certificates are often no longer usable after they expire because after the certificate expires, the issuing authority is no longer required to maintain information about the status of the certificate.

* `given_name` -  The given name present in the remote appliance certificate.

* `initials` - Initials present in the remote appliance certificate.

* `issuer` -  Issuer details of the remote appliance certificate

