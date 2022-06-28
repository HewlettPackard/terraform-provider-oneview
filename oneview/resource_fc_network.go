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
				Required: true,
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
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "fc-networkV2",
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

func resourceFCNetworkCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	fcNet := ov.FCNetwork{
		Name:                    d.Get("name").(string),
		FabricType:              d.Get("fabric_type").(string),
		LinkStabilityTime:       d.Get("link_stability_time").(int),
		ManagedSanURI:           utils.NewNstring(d.Get("managed_san_uri").(string)),
		AutoLoginRedistribution: d.Get("auto_login_redistribution").(bool),
		Type:                    d.Get("type").(string),
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
	// updates connection_template
	if rawVal, ok := d.GetOk("bandwidth"); ok {
		bandwidthVal := rawVal.([]interface{})
		for _, bandwidth := range bandwidthVal {
			rawBandwidth := bandwidth.(map[string]interface{})
			// get fc network by name
			fcNet, er := config.ovClient.GetFCNetworkByName(d.Get("name").(string))
			if er != nil {
				log.Printf("unable to get fc network for connection_template_uri: %s", er)
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
	d.Set("scopesuri", fcNet.ScopesUri.String())

	// reads bandwidth from connection template
	conTemp, err := config.ovClient.GetConnectionTemplateByURI(fcNet.ConnectionTemplateUri)
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

	// reads scopes from fc network
	scopes, err := config.ovClient.GetScopeFromResource(fcNet.URI.String())
	if err != nil {
		log.Printf("unable to fetch scopes: %s", err)
	} else {
		d.Set("initial_scope_uris", scopes.ScopeUris)
	}

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
	}

	if d.HasChange("fabric_type") {
		return fmt.Errorf("fabric type can not be changed")
	}

	if d.HasChange("initial_scope_uris") {
		// updates scopes on fc network
		val := d.Get("initial_scope_uris").(*schema.Set).List()
		err := UpdateScopeUris(meta, val, fcNet.URI.String())
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
