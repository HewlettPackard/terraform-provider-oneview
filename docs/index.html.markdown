---
layout: "oneview"
page_title: "Provider: OneView"
sidebar_current: "docs-oneview-index"
description: |-
  The Oneview provider is used to interact with your on premise OneView system. The provider needs to be configured with the proper credentials before it can be used. 
---

#Oneview Provider 

 The Oneview provider is used to interact with [OneView](https://www.hpe.com/us/en/integrated-systems/software.html). 
 The provider needs to be configured with the proper credentials before it can be used. 

##Example Usage
```js
//Configure the Oneview Provider
provider "oneview" {
  username = "username"
  password = "password123"
  endpoint = oneview_url.com
  sslverify = true
  apiversion = 200
}

//Create a new ethernet network
resource "oneview_ethernet_network" {
  // ...
}
```
## Authentication

The Oneview provider supports static credentials and environment variables.

##Configuration Reference

The following keys can be used to configure the provider.

* `username` - (Optional) This is the OneView username. 
  It must be provided or sourced from ONEVIEW_OV_USER environment variable.

* `password` - (Optional) This is the OneView password. 
  It must be provided or sourced from ONEVIEW_OV_PASSWORD environment variable.
  
* `endpoint` - (Optional) This is the OneView URL.
  It must be provided or sourced from ONEVIEW_OV_ENDPOINT environment variable.

* `sslverify` - (Optional) This is a boolean value for whether ssl is enabled. 
  It must be provided or sourced from ONEVIEW_OV_SSLVERIFY environment variable.

* `apiversion` - (Optional) This specifies what API version to use.
  It must be provided or sourced from ONEVIEW_OV_API_VERSION environment variable.

* `domain` - (Optional) This is the domain to use for the oneview system.
  It can be provided or sourced from ONEVIEW_OV_DOMAIN environment variable.
