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

func dataSourceFCoENetwork() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFCoENetworkRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"connection_template_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"managed_san_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
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
			"fabric_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
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
			"scopesuri": {
				Computed: true,
				Type:     schema.TypeString,
			},
		},
	}
}

func dataSourceFCoENetworkRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	id := d.Get("name").(string)

	fcoeNet, err := config.ovClient.GetFCoENetworkByName(id)
	if err != nil || fcoeNet.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.SetId(id)
	d.Set("name", fcoeNet.Name)
	d.Set("fabric_uri", fcoeNet.FabricUri.String())
	d.Set("vlan_id", fcoeNet.VlanId)
	d.Set("description", fcoeNet.Description.String())
	d.Set("type", fcoeNet.Type)
	d.Set("uri", fcoeNet.URI.String())
	d.Set("connection_template_uri", fcoeNet.ConnectionTemplateUri.String())
	d.Set("managed_san_uri", fcoeNet.ManagedSanUri.String())
	d.Set("status", fcoeNet.Status)
	d.Set("category", fcoeNet.Category)
	d.Set("state", fcoeNet.State)
	d.Set("created", fcoeNet.Created)
	d.Set("modified", fcoeNet.Modified)
	d.Set("etag", fcoeNet.ETAG)
	d.Set("scopesuri", fcoeNet.ScopesUri.String())
	return nil
}
