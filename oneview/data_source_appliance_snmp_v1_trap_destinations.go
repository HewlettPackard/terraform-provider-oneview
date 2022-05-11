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

func dataSourceSNMPv1TrapDestination() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSNMPv1TrapDestinationRead,

		Schema: map[string]*schema.Schema{
			"community_string": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"destination": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"destination_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSNMPv1TrapDestinationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	id := d.Get("destination_id").(string)
	snmpTrap, err := config.ovClient.GetSNMPv1TrapDestinationsById(id)
	if err != nil {
		d.SetId("")
		return err
	} else if snmpTrap.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.Set("community_string", snmpTrap.CommunityString)
	d.Set("destination", snmpTrap.Destination)
	d.Set("port", snmpTrap.Port)
	d.Set("uri", snmpTrap.URI.String())
	d.SetId(id)
	return nil
}
