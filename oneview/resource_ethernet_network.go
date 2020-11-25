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

func resourceEthernetNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceEthernetNetworkCreate,
		Read:   resourceEthernetNetworkRead,
		Update: resourceEthernetNetworkUpdate,
		Delete: resourceEthernetNetworkDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"purpose": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "General",
			},
			"private_network": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"smart_link": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ethernet_network_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Tagged",
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ethernet-networkV3",
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
			"fabric_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
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

func resourceEthernetNetworkCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	eNet := ov.EthernetNetwork{
		Name:                d.Get("name").(string),
		VlanId:              d.Get("vlan_id").(int),
		Purpose:             d.Get("purpose").(string),
		SmartLink:           d.Get("smart_link").(bool),
		PrivateNetwork:      d.Get("private_network").(bool),
		EthernetNetworkType: d.Get("ethernet_network_type").(string),
		Type:                d.Get("type").(string),
		Description:         utils.NewNstring(d.Get("description").(string)),
	}
	if val, ok := d.GetOk("initial_scope_uris"); ok {
		rawInitialScopeUris := val.(*schema.Set).List()
		initialScopeUris := make([]utils.Nstring, len(rawInitialScopeUris))
		for i, raw := range rawInitialScopeUris {
			initialScopeUris[i] = utils.Nstring(raw.(string))
		}
		eNet.InitialScopeUris = initialScopeUris
	}

	eNetError := config.ovClient.CreateEthernetNetwork(eNet)
	d.SetId(d.Get("name").(string))
	if eNetError != nil {
		d.SetId("")
		return eNetError
	}
	return resourceEthernetNetworkRead(d, meta)
}

func resourceEthernetNetworkRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	eNet, err := config.ovClient.GetEthernetNetworkByName(d.Id())
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

	return nil
}

func resourceEthernetNetworkUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	newENet := ov.EthernetNetwork{
		ETAG:                  d.Get("etag").(string),
		URI:                   utils.NewNstring(d.Get("uri").(string)),
		VlanId:                d.Get("vlan_id").(int),
		Purpose:               d.Get("purpose").(string),
		Name:                  d.Get("name").(string),
		PrivateNetwork:        d.Get("private_network").(bool),
		SmartLink:             d.Get("smart_link").(bool),
		ConnectionTemplateUri: utils.NewNstring(d.Get("connection_template_uri").(string)),
		Type:                  d.Get("type").(string),
	}

	err := config.ovClient.UpdateEthernetNetwork(newENet)
	if err != nil {
		return err
	}
	d.SetId(d.Get("name").(string))

	return resourceEthernetNetworkRead(d, meta)
}

func resourceEthernetNetworkDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteEthernetNetwork(d.Get("name").(string))
	if err != nil {
		return err
	}
	return nil
}
