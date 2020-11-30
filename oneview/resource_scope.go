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
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceScope() *schema.Resource {
	return &schema.Resource{
		Create: resourceScopeCreate,
		Read:   resourceScopeRead,
		Update: resourceScopeUpdate,
		Delete: resourceScopeDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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
				Required: true,
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

func resourceScopeCreate(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*Config)

	scope := ov.Scope{
		Type:        d.Get("type").(string),
		Name:        d.Get("name").(string),
		Description: utils.NewNstring(d.Get("description").(string)),
	}

	if val, ok := d.GetOk("initial_scope_uris"); ok {
		rawinitialScopeUris := val.(*schema.Set).List()
		initialScopeUris := make([]utils.Nstring, len(rawinitialScopeUris))
		for i, rawData := range rawinitialScopeUris {
			initialScopeUris[i] = utils.Nstring(rawData.(string))
		}
		scope.InitialScopeUris = initialScopeUris
	}

	if val, ok := d.GetOk("added_resource_uris"); ok {
		rawAddedResourceUris := val.(*schema.Set).List()
		addedScopeUris := make([]utils.Nstring, len(rawAddedResourceUris))
		for i, rawData := range rawAddedResourceUris {
			addedScopeUris[i] = utils.Nstring(rawData.(string))
		}
		scope.AddedResourceUris = addedScopeUris
	}

	if val, ok := d.GetOk("removed_resource_uris"); ok {
		rawRemovedScopeUris := val.(*schema.Set).List()
		removedScopeUris := make([]utils.Nstring, len(rawRemovedScopeUris))
		for i, rawData := range rawRemovedScopeUris {
			removedScopeUris[i] = utils.Nstring(rawData.(string))
		}
		scope.RemovedResourceUris = removedScopeUris
	}

	err := config.ovClient.CreateScope(scope)
	d.SetId(d.Get("name").(string))
	if err != nil {
		d.SetId("")
		return err
	}

	return resourceScopeRead(d, meta)
}

func resourceScopeRead(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*Config)

	scope, err := config.ovClient.GetScopeByName(d.Id())
	if err != nil || scope.URI.IsNil() {
		d.SetId("")
		return nil
	}
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

func resourceScopeUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	scope := ov.Scope{
		URI:         utils.NewNstring(d.Get("uri").(string)),
		Type:        d.Get("type").(string),
		Name:        d.Get("name").(string),
		Description: utils.NewNstring(d.Get("description").(string)),
	}

	if val, ok := d.GetOk("added_resource_uris"); ok {
		rawAddedResourceUris := val.(*schema.Set).List()
		addedResourceUris := make([]utils.Nstring, len(rawAddedResourceUris))
		for i, rawData := range rawAddedResourceUris {
			addedResourceUris[i] = utils.Nstring(rawData.(string))
		}
		scope.AddedResourceUris = addedResourceUris
	}

	if val, ok := d.GetOk("removed_resource_uris"); ok {
		rawRemovedResourceUris := val.(*schema.Set).List()
		removedResourceUris := make([]utils.Nstring, len(rawRemovedResourceUris))
		for i, rawData := range rawRemovedResourceUris {
			removedResourceUris[i] = utils.Nstring(rawData.(string))
		}
		scope.RemovedResourceUris = removedResourceUris
	}

	err := config.ovClient.UpdateScope(scope)
	if err != nil {
		d.SetId("")
		return err
	}

	d.SetId(d.Get("name").(string))
	return resourceScopeRead(d, meta)
}

func resourceScopeDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteScope(d.Get("name").(string))
	if err != nil {
		return err
	}
	return nil
}
