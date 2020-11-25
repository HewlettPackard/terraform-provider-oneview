// (C) Copyright 2020 Hewlett Packard Enterprise Development LP
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

func resourceFCoENetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceFCoENetworkCreate,
		Read:   resourceFCoENetworkRead,
		Update: resourceFCoENetworkUpdate,
		Delete: resourceFCoENetworkDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vlanid": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"connectiontemplateuri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "fcoe-network",
			},
			"managedsanuri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fabricuri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
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
			"scopesuri": {
				Optional: true,
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
		},
	}
}

func resourceFCoENetworkCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	fcoeNet := ov.FCoENetwork{
		Name:   d.Get("name").(string),
		VlanId: d.Get("vlanid").(int),
		Type:   d.Get("type").(string),
	}
	if val, ok := d.GetOk("initial_scope_uris"); ok {
		rawInitialScopeUris := val.(*schema.Set).List()
		initialScopeUris := make([]utils.Nstring, 0)

		for _, rawData := range rawInitialScopeUris {
			initialScopeUris = append(initialScopeUris, utils.Nstring(rawData.(string)))
		}
		fcoeNet.InitialScopeUris = initialScopeUris
	}
	fcoeNetError := config.ovClient.CreateFCoENetwork(fcoeNet)
	d.SetId(d.Get("name").(string))
	if fcoeNetError != nil {
		d.SetId("")
		return fcoeNetError
	}
	return resourceFCoENetworkRead(d, meta)
}

func resourceFCoENetworkRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	fcoeNet, fcoeNetError := config.ovClient.GetFCoENetworkByName(d.Id())
	if fcoeNetError != nil || fcoeNet.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.Set("vlanid", fcoeNet.VlanId)
	d.Set("created", fcoeNet.Created)
	d.Set("modified", fcoeNet.Modified)
	d.Set("uri", fcoeNet.URI.String())
	d.Set("connectiontemplateuri", fcoeNet.ConnectionTemplateUri.String())
	d.Set("status", fcoeNet.Status)
	d.Set("category", fcoeNet.Category)
	d.Set("state", fcoeNet.State)
	d.Set("fabricuri", fcoeNet.FabricUri.String())
	d.Set("etag", fcoeNet.ETAG)
	d.Set("managedsanuri", fcoeNet.ManagedSanUri)
	d.Set("description", fcoeNet.Description)
	d.Set("scopesuri", fcoeNet.ScopesUri.String())
	d.Set("initial_scope_uris", fcoeNet.InitialScopeUris)
	return nil
}

func resourceFCoENetworkUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	newFCoENet := ov.FCoENetwork{
		ETAG:                  d.Get("etag").(string),
		URI:                   utils.NewNstring(d.Get("uri").(string)),
		VlanId:                d.Get("vlanid").(int),
		Name:                  d.Get("name").(string),
		ConnectionTemplateUri: utils.NewNstring(d.Get("connectiontemplateuri").(string)),
		Type:                  d.Get("type").(string),
	}

	err := config.ovClient.UpdateFCoENetwork(newFCoENet)
	if err != nil {
		return err
	}
	d.SetId(d.Get("name").(string))

	return resourceFCoENetworkRead(d, meta)
}

func resourceFCoENetworkDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteFCoENetwork(d.Get("name").(string))
	if err != nil {
		return err
	}
	return nil
}
