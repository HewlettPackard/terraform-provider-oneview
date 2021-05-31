---
layout: "oneview"
page_title: "Oneview: Appliance Time and Locale Configuration"
sidebar_current: "docs-appliance_time_and_locale"
description: |-
 Creates appliance_time_and_locale.
---

# oneview\_appliance\_time\_and\_locale

Configures the appliance time and locale settings.

## Example Usage

```js
resource "oneview_appliance_time_and_locale" "snmptrap" {
    locale = "en_US.UTF-8"
    date_time = "2014-09-11T12:10:33"
    timezone = "UTC"
    ntp_servers = ["16.110.135.123", "16.85.40.52"]
}
```

## Argument Reference

The following arguments are supported: 

* `date_time` - Date and time of the appliance.

* `locale` - Contains locale details.

* `ntpServers` - List of NTP servers.

* `timezone` - Time zone of the appliance is in UTC.

---

* `locale_display_name` - (Optional) Display name of the locale set on the appliance.

## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `category` - Identifies the resource type.

* `etag` - Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resource.

* `uri` - The URI of the resource.

* `type` - Type of the resource.