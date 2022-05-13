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

func dataSourceSNMPv3TrapDestination() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSNMPv3TrapDestinationRead,

		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"trap_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"engine_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"destination_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id_field": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"etag": {
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

func dataSourceSNMPv3TrapDestinationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	id := d.Get("id_field").(string)
	snmpTrap, err := config.ovClient.GetSNMPv3TrapDestinationsById(id)
	if err != nil {
		d.SetId("")
		return err
	} else if snmpTrap.URI.IsNil() {
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
	d.Set("id_field", snmpTrap.ID)
	d.Set("port", snmpTrap.Port)
	d.Set("engine_id", snmpTrap.EngineID)
	d.Set("trap_type", snmpTrap.TrapType)
	d.SetId(id)
	return nil
}
