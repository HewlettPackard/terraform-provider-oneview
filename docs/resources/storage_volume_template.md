---
layout: "oneview"
page_title: "Oneview: storage_volume_templates"
sidebar_current: "docs-oneview-storage-volume-templates"
description: |-
  Creates a storage volume template.
---

# oneview\_storage\_volume\_template

Creates a storage Volume Templates.

## Example Usage

```js
resource "oneview_storage_volume_template" "default" {
  name = "SampleSVT"
  tp_name = []
  tp_description = []
  tp_storage_pool = []
  tp_is_shareable = []
  tp_provisioning_type = []
  root_template_uri = "uri"
}
```

## Argument Reference

The following arguments are supported: 

* `name` - (Required) Name of the User Template.

* `root_template_uri` - (Required) Uri of the root template that the user template is associated with.

---

* `tp_description` -  The template specific parameters.

* `tp_is_shareable` - The template specific parameters.

* `tp_name` - The template specific parameters.

* `tp_is_shareable` - The template specific parameters.

* `tp_provisioning_type` - The template specific parameters.

