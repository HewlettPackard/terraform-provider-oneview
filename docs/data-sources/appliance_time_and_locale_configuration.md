---
layout: "oneview"
page_title: "Oneview: Appliance Time and Locale Configuration"
sidebar_current: "docs-appliance_time_and_locale"
description: |-
 Gets information about an existing appliance_time_and_locale.
---

# oneview\_appliance\_time\_and\_locale

Use this data source to access the attributes of a appliance time and locale configuration.

## Example Usage

```hcl
data "oneview_appliance_time_and_locale" "test" {
}

output "locale_value" {
 value = "${data.oneview_appliance_time_and_locale.test.locale}"
}
```

## Argument Reference

No argument is required to retrive the data source of existing appliance time and locale configuration 

## Attributes Reference

* `date_time` - Date and time of the appliance.

* `locale` - Contains locale details.

* `locale_display_name` - Display name of the locale set on the appliance.

* `ntpServers` - List of NTP servers.

* `etag` - Entity tag/version ID of the resource, the same value that is returned in the ETag header on a GET of the resource.

* `pollingInterval` -  Polling interval of NTP client.

* `timezone` - Time zone of the appliance is in UTC.

* `type` - Type of the resource.