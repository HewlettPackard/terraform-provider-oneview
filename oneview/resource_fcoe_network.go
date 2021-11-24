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
			},
			"connection_template_uri": {
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

	// updates connection_template
	if rawVal, ok := d.GetOk("bandwidth"); ok {
		bandwidthVal := rawVal.([]interface{})
		for _, bandwidth := range bandwidthVal {
			rawBandwidth := bandwidth.(map[string]interface{})
			// get fcoe network by name
			fcNet, er := config.ovClient.GetFCoENetworkByName(d.Get("name").(string))
			if er != nil {
				log.Printf("unable to get fcoe network for connection_template_uri: %s", er)
			}
			// get connection template by uri
			conTemp, er := config.ovClient.GetConnectionTemplateByURI(fcNet.ConnectionTemplateUri)
			if er != nil {
				log.Printf("unable to get connection template by uri: %s", er)
			}
			if fcNet.ConnectionTemplateUri.String() != "" {
				// filter URI
				id := path.Base(fcNet.ConnectionTemplateUri.String())
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
	d.Set("connection_template_uri", fcoeNet.ConnectionTemplateUri.String())
	d.Set("status", fcoeNet.Status)
	d.Set("category", fcoeNet.Category)
	d.Set("state", fcoeNet.State)
	d.Set("fabricuri", fcoeNet.FabricUri.String())
	d.Set("etag", fcoeNet.ETAG)
	d.Set("managedsanuri", fcoeNet.ManagedSanUri)
	d.Set("scopesuri", fcoeNet.ScopesUri.String())

	// reads bandwidth from connection template
	conTemp, err := config.ovClient.GetConnectionTemplateByURI(fcoeNet.ConnectionTemplateUri)
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

	// reads scopes from fcoe network
	scopes, err := config.ovClient.GetScopeFromResource(fcoeNet.URI.String())
	if err != nil {
		log.Printf("unable to fetch scopes: %s", err)
	} else {
		d.Set("initial_scope_uris", scopes.ScopeUris)
	}
	return nil
}

func resourceFCoENetworkUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	newFCoENet := ov.FCoENetwork{
		ETAG:                  d.Get("etag").(string),
		URI:                   utils.NewNstring(d.Get("uri").(string)),
		VlanId:                d.Get("vlanid").(int),
		Name:                  d.Get("name").(string),
		ConnectionTemplateUri: utils.NewNstring(d.Get("connection_template_uri").(string)),
		Type:                  d.Get("type").(string),
	}
	if d.HasChange("vlanid") {
		return fmt.Errorf("vlan Id can not be changed")
	}
	if d.HasChange("initial_scope_uris") {
		return fmt.Errorf("Initial scope uri can not be updated")
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

	if d.HasChange("initial_scope_uris") {
		// updates scopes on fcoe network
		val := d.Get("initial_scope_uris").(*schema.Set).List()
		err := UpdateScopeUris(meta, val, newFCoENet.URI.String())
		if err != nil {
			return err
		}
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
