// (C) Copyright 2016 Hewlett Packard Enterprise Development LP
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

func dataSourceNetworkSet() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkSetRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_uris": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"native_network_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_template_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scopes_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_set_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceNetworkSetRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	netSet, err := config.ovClient.GetNetworkSetByName(d.Get("name").(string))
	if err != nil || netSet.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.SetId(d.Get("name").(string))
	d.Set("name", netSet.Name)
	d.Set("type", netSet.Type)
	d.Set("created", netSet.Created)
	d.Set("description", netSet.Description)
	d.Set("etag", netSet.ETAG)
	d.Set("modified", netSet.Modified)
	d.Set("native_network_uri", netSet.NativeNetworkUri)
	d.Set("uri", netSet.URI.String())
	d.Set("connection_template_uri", netSet.ConnectionTemplateUri.String())
	d.Set("status", netSet.Status)
	d.Set("category", netSet.Category)
	d.Set("state", netSet.State)

	networkUris := make([]interface{}, len(netSet.NetworkUris))
	for i := 0; i < len(netSet.NetworkUris); i++ {
		networkUris[i] = netSet.NetworkUris[i].String()
	}

	rawNetUris := d.Get("network_uris").(*schema.Set).List()
	for i, currNetworkUri := range rawNetUris {
		for j := 0; j < len(networkUris); j++ {
			if currNetworkUri.(string) == networkUris[j] && i <= len(networkUris)-1 {
				networkUris[i], networkUris[j] = networkUris[j], networkUris[i]
			}
		}
	}
	d.Set("network_uris", networkUris)
	d.Set("scopes_uri", netSet.ScopesUri)
	d.Set("network_set_type", netSet.NetworkSetType)

	return nil

}
