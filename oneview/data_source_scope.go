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

func dataSourceScope() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceScopeRead,

		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ext_attributes": {
				Optional: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"appliance_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"old_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scopes_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"initial_scope_uris": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"added_resource_uris": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"removed_resource_uris": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
		},
	}
}

func dataSourceScopeRead(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*Config)
	name := d.Get("name").(string)

	scope, err := config.ovClient.GetScopeByName(name)
	if err != nil || scope.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.SetId(name)
	d.Set("name", scope.Name)
	d.Set("description", scope.Description)
	d.Set("state", scope.State)
	d.Set("status", scope.Status)
	d.Set("type", scope.Type)
	d.Set("uri", scope.URI.String())
	d.Set("appliance_id", scope.ApplianceId)
	d.Set("category", scope.Category)
	d.Set("etag", scope.Etag)
	d.Set("old_uri", scope.OldUri.String())
	d.Set("scopes_uri", scope.ScopesUri.String())
	d.Set("initial_scope_uris", scope.InitialScopeUris)
	d.Set("added_resource_uris", scope.AddedResourceUris)
	d.Set("removed_resource_uris", scope.RemovedResourceUris)

	return nil
}
