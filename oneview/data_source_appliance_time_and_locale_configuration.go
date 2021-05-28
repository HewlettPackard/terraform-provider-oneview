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
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
				Set:      schema.HashString,
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
	id := d.Get("id").(string)
	snmpTrap, err := config.ovClient.GetSNMPv3TrapDestinationsById(id)
	if err != nil || snmpTrap.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.Set("type", snmpTrap.Type)
	d.Set("created", snmpTrap.Created)
	d.Set("modified", snmpTrap.Modified)
	d.Set("uri", snmpTrap.URI.String())
	d.Set("destination_address", snmpTrap.DestinationAddress)
	d.Set("category", snmpTrap.Category)
	d.Set("user_uri", snmpTrap.UserURI)
	d.Set("etag", snmpTrap.ETAG)
	d.Set("user_id", snmpTrap.UserID)
	d.Set("id", snmpTrap.ID)
	d.Set("port", snmpTrap.Port)
	d.SetId(id)
	return nil
}
