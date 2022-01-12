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
	"fmt"
	"log"
	"path"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
				Optional: true,
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
				Required: true,
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
			"subnet_uri": {
				Optional: true,
				Computed: true,
				Type:     schema.TypeString,
			},
			"initial_scope_uris": {
				Optional: true,
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"bandwidth": {
				Optional: true,
				Computed: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"maximum_bandwidth": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"typical_bandwidth": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
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
		SubnetUri:           utils.NewNstring(d.Get("subnet_uri").(string)),
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
	if eNetError != nil {
		d.SetId("")
		return eNetError
	}

	// updates connection_template
	if rawVal, ok := d.GetOk("bandwidth"); ok {
		bandwidthVal := rawVal.([]interface{})
		for _, bandwidth := range bandwidthVal {
			rawBandwidth := bandwidth.(map[string]interface{})
			// get ethernet network by name
			eNet, er := config.ovClient.GetEthernetNetworkByName(d.Get("name").(string))
			if er != nil {
				log.Printf("unable to get ethernet network for connection_template_uri: %s", er)
			}
			// get connection template by uri
			conTemp, er := config.ovClient.GetConnectionTemplateByURI(eNet.ConnectionTemplateUri)
			if er != nil {
				log.Printf("unable to get connection template by uri: %s", er)
			}
			if eNet.ConnectionTemplateUri.String() != "" {
				// filter URI
				id := path.Base(eNet.ConnectionTemplateUri.String())
				// update the con_temp with required bandwidth
				BandwidthOptions := ov.BandwidthType{
					MaximumBandwidth: rawBandwidth["maximum_bandwidth"].(int),
					TypicalBandwidth: rawBandwidth["typical_bandwidth"].(int),
				}
				conTemp.Bandwidth = BandwidthOptions
				_, er = config.ovClient.UpdateConnectionTemplate(id, conTemp)
				if er != nil {
					log.Printf("unable to update the connection template: %s", er)
				}
			}
		}

	}
	d.SetId(d.Get("name").(string))
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
	d.Set("subnet_uri", eNet.SubnetUri.String())

	// reads bandwidth from connection template
	conTemp, err := config.ovClient.GetConnectionTemplateByURI(eNet.ConnectionTemplateUri)
	if err != nil {
		log.Printf("unable to fetch connection template: %s", err)
	} else {
		bandwidth := make([]interface{}, 0)
		bw := map[string]interface{}{}
		bw["typical_bandwidth"] = conTemp.Bandwidth.TypicalBandwidth
		bw["maximum_bandwidth"] = conTemp.Bandwidth.MaximumBandwidth
		bandwidth = append(bandwidth, bw)
		d.Set("bandwidth", bandwidth)
	}
	// reads scopes from ethernet network
	scopes, err := config.ovClient.GetScopeFromResource(eNet.URI.String())
	if err != nil {
		log.Printf("unable to fetch scopes: %s", err)
	} else {
		d.Set("initial_scope_uris", scopes.ScopeUris)
	}
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

	if d.HasChange("vlan_id") {
		return fmt.Errorf("vlan Id can not be changed")
	}

	if d.HasChange("initial_scope_uris") {
		// updates scopes on ethernet network
		val := d.Get("initial_scope_uris").(*schema.Set).List()
		err := UpdateScopeUris(meta, val, newENet.URI.String())
		if err != nil {
			return err
		}
	}

	if d.HasChange("bandwidth") {
		rawVal := d.Get("bandwidth")
		bandwidthVal := rawVal.([]interface{})
		for _, bandwidth := range bandwidthVal {
			rawBandwidth := bandwidth.(map[string]interface{})
			conTempURI := utils.NewNstring(d.Get("connection_template_uri").(string))
			// get connection template by uri
			conTemp, err := config.ovClient.GetConnectionTemplateByURI(conTempURI)
			if err != nil {
				return fmt.Errorf("unable to retrieve connection template: %s", err)
			}
			// filter URI
			id := path.Base(conTempURI.String())
			// update the con_temp with required bandwidth
			BandwidthOptions := ov.BandwidthType{
				MaximumBandwidth: rawBandwidth["maximum_bandwidth"].(int),
				TypicalBandwidth: rawBandwidth["typical_bandwidth"].(int),
			}
			conTemp.Bandwidth = BandwidthOptions
			conTemp, err = config.ovClient.UpdateConnectionTemplate(id, conTemp)
			if err != nil {
				return fmt.Errorf("unable to update bandwidth: %s", err)
			}
		}
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
