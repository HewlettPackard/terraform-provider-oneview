---
layout: "oneview"
page_title: "Provider: OneView"
sidebar_current: "docs-oneview-index"
description: |-
  The Oneview provider is used to interact with your on premise OneView system. The provider needs to be configured with the proper credentials before it can be used.
---

# Oneview Provider

 The Oneview provider is used to interact with [OneView](https://www.hpe.com/us/en/integrated-systems/software.html).
 The provider needs to be configured with the proper credentials before it can be used.

## Example Usage
```js
//Configure the Oneview Provider
provider "oneview" {
  ov_username = "username"
  ov_password = "password123"
  ov_endpoint = oneview_url.com
  ov_sslverify = true
  ov_apiversion = 2400
  ov_domain = "Local"
  ov_ifmatch = "*"
}

//Create a new ethernet network
resource "oneview_ethernet_network" {
  // ...
}
```
## Authentication

The Oneview provider supports static credentials and environment variables.

## Configuration Reference

The following keys can be used to configure the provider.

* `ov_username` - (Optional) This is the OneView username.
  It must be provided or sourced from ONEVIEW_OV_USER environment variable.

* `ov_password` - (Optional) This is the OneView password.
  It must be provided or sourced from ONEVIEW_OV_PASSWORD environment variable.

* `ov_endpoint` - (Optional) This is the OneView URL.
  It must be provided or sourced from ONEVIEW_OV_ENDPOINT environment variable.

* `ov_sslverify` - (Optional) This is a boolean value for whether ssl is enabled.
  It must be provided or sourced from ONEVIEW_OV_SSLVERIFY environment variable.

* `ov_apiversion` - (Optional) This specifies what API version to use.
  It must be provided or sourced from ONEVIEW_OV_API_VERSION environment variable.

* `ov_domain` - (Optional) This is the domain to use for the oneview system.
  It can be provided or sourced from ONEVIEW_OV_DOMAIN environment variable.

* `ov_ifmatch` - (Optional) This is the if-Match(ETag) request header to use for the oneview system.
    It can be provided or sourced from ONEVIEW_OV_IF_MATCH environment variable.