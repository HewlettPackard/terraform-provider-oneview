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
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFCNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceFCNetworkCreate,
		Read:   resourceFCNetworkRead,
		Update: resourceFCNetworkUpdate,
		Delete: resourceFCNetworkDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fabric_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "FabricAttach",
			},
			"link_stability_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"auto_login_redistribution": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
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
				Optional: true,
				Default:  "fc-networkV2",
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
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
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified": {
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

func resourceFCNetworkCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	fcNet := ov.FCNetwork{
		Name:                    d.Get("name").(string),
		FabricType:              d.Get("fabric_type").(string),
		LinkStabilityTime:       d.Get("link_stability_time").(int),
		ManagedSanURI:           utils.NewNstring(d.Get("managed_san_uri").(string)),
		AutoLoginRedistribution: d.Get("auto_login_redistribution").(bool),
		Type:                    d.Get("type").(string),
		Description:             utils.NewNstring(d.Get("description").(string)),
	}

	if val, ok := d.GetOk("initial_scope_uris"); ok {
		rawInitialScopeUris := val.(*schema.Set).List()
		initialScopeUris := make([]utils.Nstring, len(rawInitialScopeUris))
		for i, raw := range rawInitialScopeUris {
			initialScopeUris[i] = utils.Nstring(raw.(string))
		}
		fcNet.InitialScopeUris = initialScopeUris
	}
	fcNetError := config.ovClient.CreateFCNetwork(fcNet)
	d.SetId(d.Get("name").(string))
	if fcNetError != nil {
		d.SetId("")
		return fcNetError
	}
	return resourceFCNetworkRead(d, meta)
}

func resourceFCNetworkRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	fcNet, err := config.ovClient.GetFCNetworkByName(d.Id())
	if err != nil || fcNet.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.Set("name", fcNet.Name)
	d.Set("fabric_type", fcNet.FabricType)
	d.Set("link_stability_time", fcNet.LinkStabilityTime)
	d.Set("auto_login_redistribution", fcNet.AutoLoginRedistribution)
	d.Set("description", fcNet.Description.String())
	d.Set("type", fcNet.Type)
	d.Set("uri", fcNet.URI.String())
	d.Set("connection_template_uri", fcNet.ConnectionTemplateUri.String())
	d.Set("managed_san_uri", fcNet.ManagedSanURI.String())
	d.Set("status", fcNet.Status)
	d.Set("category", fcNet.Category)
	d.Set("state", fcNet.State)
	d.Set("fabric_uri", fcNet.FabricUri.String())
	d.Set("created", fcNet.Created)
	d.Set("modified", fcNet.Modified)
	d.Set("etag", fcNet.ETAG)
	d.Set("scopeuri", fcNet.ScopesUri.String())
	d.Set("initial_scope_uris", fcNet.InitialScopeUris)
	return nil
}

func resourceFCNetworkUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	fcNet := ov.FCNetwork{
		ETAG:                    d.Get("etag").(string),
		URI:                     utils.NewNstring(d.Get("uri").(string)),
		Name:                    d.Get("name").(string),
		FabricType:              d.Get("fabric_type").(string),
		LinkStabilityTime:       d.Get("link_stability_time").(int),
		AutoLoginRedistribution: d.Get("auto_login_redistribution").(bool),
		Type:                    d.Get("type").(string),
		ConnectionTemplateUri:   utils.NewNstring(d.Get("connection_template_uri").(string)),
		Description:             utils.NewNstring(d.Get("description").(string)),
	}

	err := config.ovClient.UpdateFcNetwork(fcNet)
	if err != nil {
		return err
	}
	d.SetId(d.Get("name").(string))

	return resourceFCNetworkRead(d, meta)
}

func resourceFCNetworkDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteFCNetwork(d.Get("name").(string))
	if err != nil {
		return err
	}
	return nil
}
