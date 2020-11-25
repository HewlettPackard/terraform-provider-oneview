// (C) Copyright 2019 Hewlett Packard Enterprise Development LP
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

func dataSourceEthernetNetwork() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEthernetNetworkRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"purpose": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_network": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"smart_link": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ethernet_network_type": {
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
			"fabric_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scopesuri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"initial_scope_uris": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
		},
	}
}

func dataSourceEthernetNetworkRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)
	eNet, err := config.ovClient.GetEthernetNetworkByName(name)
	if err != nil || eNet.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.Set("name", eNet.Name)
	d.Set("vlan_id", eNet.VlanId)
	d.Set("purpose", eNet.Purpose)
	d.Set("smart_link", eNet.SmartLink)
	d.Set("private_network", eNet.PrivateNetwork)
	d.Set("ethernet_network_type", eNet.EthernetNetworkType)
	d.Set("type", eNet.Type)
	d.Set("created", eNet.Created)
	d.Set("modified", eNet.Modified)
	d.Set("uri", eNet.URI.String())
	d.Set("connection_template_uri", eNet.ConnectionTemplateUri.String())
	d.Set("status", eNet.Status)
	d.Set("category", eNet.Category)
	d.Set("state", eNet.State)
	d.Set("fabric_uri", eNet.FabricUri.String())
	d.Set("etag", eNet.ETAG)
	d.Set("scopesuri", eNet.ScopesUri.String())
	d.Set("initial_scope_uris", eNet.InitialScopeUris)
	d.SetId(name)
	return nil
}
