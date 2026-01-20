---
layout: "oneview"
page_title: "Oneview: server_certificate"
sidebar_current: "docs-oneview-server-certificate"
description: |-
  Creates a server certificate.
---

# oneview\_server\_certificate

Creates an server certificate.

## Example Usage

```js
resource "oneview_server_certificate" "default" {
   certificate_details = [{
     base64_data="<base64_date>"
     alias_name = "test-server-certificate"
     type="CertificateDetailsV2"
}]
}
```

## Argument Reference

The following arguments are supported: 

* `alias_name` - A unique name for the resource.
* `type` - (Required) Indicates the type of the Server certificate.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

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
