// (C) Copyright 2022 Hewlett Packard Enterprise Development LP
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
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRackManager() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRackManagerRead,
		Schema: map[string]*schema.Schema{
			"category": {
				Type: schema.TypeString,

				Optional: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"licensing_intent": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"model": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"part_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"refresh_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_support_uri": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scopes_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"initial_scope_uris": {
				Optional: true,
				ForceNew: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"support_data_collection_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"support_data_collection_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"support_data_collections_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"support_state": {
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
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceRackManagerRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	id := d.Get("name").(string)
	rm, err := config.ovClient.GetRackManagerByName(id)
	if err != nil || rm.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.SetId(rm.Id)
	d.Set("category", rm.Category)
	d.Set("created", rm.Created)
	d.Set("etag", rm.ETAG)

	d.Set("licensing_intent", rm.LicensingIntent)
	d.Set("location", rm.Location)
	d.Set("model", rm.Model)
	d.Set("modified", rm.Modified)
	d.Set("name", rm.Name)
	d.Set("part_number", rm.PartNumber)
	d.Set("refresh_state", rm.RefreshState)
	d.Set("remote_support_uri", rm.RemoteSupportUri)

	d.Set("scopes_uri", rm.ScopesUri)
	d.Set("serial_number", rm.SerialNumber)
	d.Set("state", rm.State)
	d.Set("status", rm.Status)
	d.Set("support_data_collection_state", rm.SupportDataCollectionState)
	d.Set("support_data_collection_type", rm.SupportDataCollectionType)
	d.Set("support_data_dollections_uri", rm.SupportDataCollectionsUri)
	d.Set("type", rm.Type)
	d.Set("uri", rm.URI.String())

	// reads scopes from rack manager
	scopes, err := config.ovClient.GetScopeFromResource(rm.URI.String())
	if err != nil {
		log.Printf("unable to fetch scopes: %s", err)
	} else {
		d.Set("initial_scope_uris", scopes.ScopeUris)
	}

	return nil
}
