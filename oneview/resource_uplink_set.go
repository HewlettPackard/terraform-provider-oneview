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
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUplinkSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceUplinkSetCreate,
		Read:   resourceUplinkSetRead,
		Update: resourceUplinkSetUpdate,
		Delete: resourceUplinkSetDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"logical_interconnect_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_uris": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"fc_network_uris": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"fcoe_network_uris": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"connection_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ethernet_network_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "uplink-setV4",
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
			"eTag": {
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
			"lcaptimer": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"manual_login_redistribution_state": {
				Type:     schema.TypeString,
				Required: true,
			},
			"native_network_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reachability": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fc_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expected_neighbour": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"remote_chasis_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"remote_port_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"port_config_infos": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"desired_speed": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"location": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"location_entries": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"type": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceUplinkSetCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	uplinkSet := ov.UplinkSet{
		Name: d.Get("name").(string),
		LogicalInterconnectURI: utils.NewNstring(d.Get("logical_interconnect_uri").(string)),
		ConnectionMode:         d.Get("connection_mode").(string),
		NetworkType:            d.Get("network_type").(string),
		EthernetNetworkType:    d.Get("ethernet_network_type").(string),
		Type:                   d.Get("type").(string),
		ManualLoginRedistributionState: d.Get("manual_login_redistribution_state").(string),
	}

	if val, ok := d.GetOk("network_uris"); ok {
		rawNetworkUris := val.(*schema.Set).List()
		NetworkUris := make([]utils.Nstring, len(rawNetworkUris))
		for i, raw := range rawNetworkUris {
			NetworkUris[i] = utils.Nstring(raw.(string))
		}
		uplinkSet.NetworkURIs = NetworkUris
	}

	if val, ok := d.GetOk("fc_network_uris"); ok {
		rawFcNetworkUris := val.(*schema.Set).List()
		FcNetworkUris := make([]utils.Nstring, len(rawFcNetworkUris))
		for i, raw := range rawFcNetworkUris {
			FcNetworkUris[i] = utils.Nstring(raw.(string))
		}
		uplinkSet.FcNetworkURIs = FcNetworkUris
	}

	if val, ok := d.GetOk("fcoe_network_uris"); ok {
		rawFcoeNetworkUris := val.(*schema.Set).List()
		FcoeNetworkUris := make([]utils.Nstring, len(rawFcoeNetworkUris))
		for i, raw := range rawFcoeNetworkUris {
			FcoeNetworkUris[i] = utils.Nstring(raw.(string))
		}
		uplinkSet.FcoeNetworkURIs = FcoeNetworkUris
	}

	portConfigInfosCount := d.Get("port_config_infos.#").(int)
	portConfigInfosAll := make([]ov.PortConfigInfos, 0)

	for i := 0; i < portConfigInfosCount; i++ {
		portConfigInfosPrefix := fmt.Sprintf("port_config_infos.%d", i)
		location := ov.Location{}

		locationPrefix := fmt.Sprintf(portConfigInfosPrefix + ".location.0")
		locationEntriesCount := d.Get(locationPrefix + ".location_entries.#").(int)
		locationEntriesAll := make([]ov.LocationEntries, 0)

		for i := 0; i < locationEntriesCount; i++ {
			locationEntriesPrefix := fmt.Sprintf(locationPrefix+".locationEntries.%d", i)
			locationEntries := ov.LocationEntries{
				Value: d.Get(locationEntriesPrefix + ".value").(string),
				Type: d.Get(locationEntriesPrefix + ".type").(string),
			}
			locationEntriesAll = append(locationEntriesAll, locationEntries)
		}

		if locationEntriesCount > 0 {
			location.LocationEntries = locationEntriesAll
		}
		portConfigInfos := ov.PortConfigInfos{}
		portConfigInfos.Location = location

		if val, ok := d.GetOk(portConfigInfosPrefix + ".desired_speed"); ok {
			portConfigInfos.DesiredSpeed = val.(string)
		}

		portConfigInfosAll = append(portConfigInfosAll, portConfigInfos)

	}
	uplinkSet.PortConfigInfos = portConfigInfosAll

	uplinkSetError := config.ovClient.CreateUplinkSet(uplinkSet)
	d.SetId(d.Get("name").(string))
	if uplinkSetError != nil {
		d.SetId("")
		return uplinkSetError
	}
	return resourceUplinkSetRead(d, meta)
}

func resourceUplinkSetRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	uplinkSet, err := config.ovClient.GetUplinkSetByName(d.Get("name").(string))
	if err != nil || uplinkSet.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.Set("name", uplinkSet.Name)
	d.Set("logical_interconnect_uri", uplinkSet.LogicalInterconnectURI)
	d.Set("network_uris", uplinkSet.NetworkURIs)
	d.Set("manual_login_redistribution_state", uplinkSet.ManualLoginRedistributionState)
	d.Set("description", uplinkSet.Description)
	d.Set("type", uplinkSet.Type)
	d.Set("uri", uplinkSet.URI.String())
	d.Set("fcoe_network_uris", uplinkSet.FcoeNetworkURIs)
	d.Set("status", uplinkSet.Status)
	d.Set("category", uplinkSet.Category)
	d.Set("state", uplinkSet.State)
	d.Set("fc_network_uris", uplinkSet.FcNetworkURIs)
	d.Set("created", uplinkSet.Created)
	d.Set("modified", uplinkSet.Modified)
	d.Set("eTag", uplinkSet.Etag)
	d.Set("reachability", uplinkSet.Reachability)
	d.Set("expected_neighbor", uplinkSet.ExpectedNeighbor)
	d.Set("network_type", uplinkSet.NetworkType)
	d.Set("ethernet_network_type", uplinkSet.EthernetNetworkType)
	d.Set("port_config_infos", uplinkSet.PortConfigInfos)
	return nil
}

func resourceUplinkSetUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	uplinkSet := ov.UplinkSet{
		Etag:        d.Get("eTag").(string),
		URI:         utils.NewNstring(d.Get("uri").(string)),
		Name:        d.Get("name").(string),
		Type:        d.Get("type").(string),
		Description: utils.NewNstring(d.Get("description").(string)),
	}

	err := config.ovClient.UpdateUplinkSet(uplinkSet)
	if err != nil {
		return err
	}
	d.SetId(d.Get("name").(string))

	return resourceUplinkSetRead(d, meta)
}

func resourceUplinkSetDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteUplinkSet(d.Get("name").(string))
	if err != nil {
		return err
	}
	return nil
}
