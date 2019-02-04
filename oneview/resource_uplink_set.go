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
				Type:     schema.TypeBool,
				Optional: true,
			},
			"network_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ethernet_network_type": {
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
			"manual_logi_redistribution_state": {
				Type:     schema.TypeString,
				Computed: true,
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
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"desired_speed": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"location": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"primary_port_location": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"location": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"location_entries": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"location": {
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
		LogicalInterconnectURI:         d.Get("logical_interconnect_uri"),
		ConnectionMode:                 g.Get("connection_mode"),
		NetworkType:                    d.Get("network_type"),
		EthernetNetworkType:            d.Get("ethernet_network_type"),
		Type:                           d.Get("type"),
		ManualLoginRedistributionState: d.Get("manual_login_redistribution_state"),
	}

	if val, ok := d.GetOk("network_uris"); ok {
		rawNetworkUris := val.(*schema.Set).List()
		NetworkUris := make([]utils.Nstring, len(rawNetworkUris))
		for i, raw := range rawNetworkUris {
			NetworkUris[i] = utils.Nstring(raw.(string))
		}
		uplinkSet.NetworkUris = NetworkUris	
	}

	if val, ok := d.GetOk("fc_network_uris"); ok {
		rawFcNetworkUris := val.(*schema.Set).List()
		FcNetworkUris := make([]utils.Nstring, len(rawFcNetworkUris))
		for i, raw := range rawFcNetworkUris {
			FcNetworkUris[i] = utils.Nstring(raw.(string))
		}
		uplinkSet.FcNetworkUris = FcNetworkUris
	}

	if val, ok := d.GetOk("fcoe_network_uris"); ok {
		rawFcoeNetworkUris := val.(*schema.Set).List()
		FcoeNetworkUris := make([]utils.Nstring, len(rawFcoeNetworkUris))
		for i, raw := range rawFcoeNetworkUris {
			FcoeNetworkUris[i] = utils.Nstring(raw.(string))
		}
		uplinkSet.FcoeNetworkUris = FcoeNetworkUris
	}

	portLocationEntriesPrefix := fmt.Sprintf("location_entries.0")
	portLocationEntries := ov.LocationEntries{}
	if val, ok := d.GetOk(portLocationEntries + ".value"); ok {
		portLocationEntries.DesiredSpeed = val.(string)
	}
	if val, ok := d.GetOk(portLocationEntries + ".type"); ok {
		portLocationEntries.location = &portLocation
	}
	
	portLocationPrefix := fmt.Sprintf("location.0")
	portLocation := ov.Location{}
	if val, ok := d.GetOk(portLocation + ".location_entries"); ok {
		portLocation =  &portLocationEntries
	}
	
	portConfigInfosPrefix := fmt.Sprintf("port_config_infos.0")
	portConfigInfos := ov.PortConfigInfos{}
	if val, ok := d.GetOk(portConfigInfosPrefix + ".desired_speed"); ok {
		portConfigInfos.DesiredSpeed = val.(string)
	}
	if val, ok := d.GetOk(portConfigInfosPrefix + ".location"); ok {
		portConfigInfos.location = &portLocation
	}
	if portConfigInfos != (ov.PortConfigInfos{}) {
		uplink_set.portConfigInfos = &portConfigInfos
	}

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
        d.Set("logical_interconnect_uri", uplinkSet.LogicalInterconnectUri)
        d.Set("network_uris", uplinkSet.NetworkURIs.String())
        d.Set("manual_login_redistribution_state", uplinkSet.ManualLoginRedistributionState)
        d.Set("description", uplinkSet.Description)
        d.Set("type", uplinkSet.Type)
        d.Set("uri", uplinkSet.URI.String())
        d.Set("fcoe_network_uris", uplinkSet.FcoeNetworkURIs.String())
        d.Set("status", uplinkSet.Status)
        d.Set("category", uplinkSet.Category)
        d.Set("state", uplinkSet.State)
        d.Set("fc_network_uris", uplinkSet.FcNetworkURIs.String())
        d.Set("created", uplinkSet.Created)
        d.Set("modified", uplinkSet.Modified)
        d.Set("eTag", uplinkSet.ETAG)
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
		ETAG:                    d.Get("eTag").(string),
		URI:                     utils.NewNstring(d.Get("uri").(string)),
		Name:                    d.Get("name").(string),
		FabricType:              d.Get("fabric_type").(string),
		LinkStabilityTime:       d.Get("link_stability_time").(int),
		AutoLoginRedistribution: d.Get("auto_login_redistribution").(bool),
		Type: d.Get("type").(string),
		ConnectionTemplateUri: utils.NewNstring(d.Get("connection_template_uri").(string)),
		Description:           d.Get("description").(string),
	}

	err := config.ovClient.UpdateFcNetwork(uplinkSet)
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
