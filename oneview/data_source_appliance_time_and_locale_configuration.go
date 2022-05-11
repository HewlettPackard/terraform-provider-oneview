// (C) Copyright 2021 Hewlett Packard Enterprise Development LP
//
// Licensed under the Apache License, Version 2.0 (the "License");
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed
// under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package oneview

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTimeAndLocale() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTimeAndLocaleRead,

		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"locale": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"locale_displayname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ntp_servers": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
			"polling_interval": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"timezone": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceTimeAndLocaleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	timeLocale, err := config.ovClient.GetApplianceTimeandLocals("", "", "", "")
	if err != nil {
		d.SetId("")
		return err
	} else if timeLocale.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.Set("type", timeLocale.Type)
	d.Set("created", timeLocale.Created)
	d.Set("modified", timeLocale.Modified)
	d.Set("uri", timeLocale.URI.String())
	d.Set("date_time", timeLocale.DateTime)
	d.Set("category", timeLocale.Category)
	d.Set("locale", timeLocale.Locale)
	d.Set("etag", timeLocale.ETAG)
	d.Set("locale_displayname", timeLocale.LocaleDisplayName)
	d.Set("polling_interval", timeLocale.PollingInterval)
	d.Set("timezone", timeLocale.Timezone)
	d.Set("ntp_servers", timeLocale.NtpServers)
	d.SetId(timeLocale.Locale)
	return nil
}
