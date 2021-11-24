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
	"errors"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTimeAndLocale() *schema.Resource {
	return &schema.Resource{
		Read:   resourceTimeAndLocaleRead,
		Create: resourceTimeAndLocaleCreate,
		Update: resourceTimeAndLocaleUpdate,
		Delete: resourceTimeAndLocaleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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
				Required: true,
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
				Required: true,
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
				Required: true,
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

func resourceTimeAndLocaleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	Locale := ov.ApplianceTimeandLocal{
		Locale:            d.Get("locale").(string),
		DateTime:          utils.NewNstring(d.Get("date_time").(string)),
		Timezone:          utils.NewNstring(d.Get("timezone").(string)),
		LocaleDisplayName: utils.NewNstring(d.Get("locale_displayname").(string)),
	}
	if val, ok := d.GetOk("ntp_servers"); ok {
		rawNtpServers := val.(*schema.Set).List()
		NtpServers := make([]utils.Nstring, len(rawNtpServers))
		for i, raw := range rawNtpServers {
			NtpServers[i] = utils.Nstring(raw.(string))
		}
		Locale.NtpServers = NtpServers
	}

	err := config.ovClient.CreateApplianceTimeandLocal(Locale)
	if err != nil {
		return err
	}
	d.SetId(d.Get("locale").(string))
	return resourceTimeAndLocaleRead(d, meta)
}

func resourceTimeAndLocaleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	timeLocale, err := config.ovClient.GetApplianceTimeandLocals("", "", "", "")
	if err != nil {
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

func resourceTimeAndLocaleDelete(d *schema.ResourceData, meta interface{}) error {
	err := errors.New("this resource do not support delete request")
	return err
}

func resourceTimeAndLocaleUpdate(d *schema.ResourceData, meta interface{}) error {
	err := errors.New("this resource do not support update request")
	return err
}
