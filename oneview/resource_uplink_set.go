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
				Required: true,
			},
			"network_uris": {
				Required: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"fc_network_uris": {
				Required: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"fcoe_network_uris": {
				Required: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"connection_mode": {
				Type:     schema.TypeString,
				Required: true,
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
				Optional: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lacptimer": {
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
			"port_config_infos": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"desired_speed": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"bay_number": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"enclosure_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_number": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"remote_chassis_id": {
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
		},
	}
}

func resourceUplinkSetCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	const Bay = "Bay"
	const Enclosure = "Enclosure"
	const Port = "Port"

	uplinkSet := ov.UplinkSet{
		Name:                           d.Get("name").(string),
		LogicalInterconnectURI:         utils.NewNstring(d.Get("logical_interconnect_uri").(string)),
		ConnectionMode:                 d.Get("connection_mode").(string),
		NetworkType:                    d.Get("network_type").(string),
		EthernetNetworkType:            d.Get("ethernet_network_type").(string),
		Type:                           d.Get("type").(string),
		ManualLoginRedistributionState: d.Get("manual_login_redistribution_state").(string),
	}

	networkUriList := d.Get("network_uris").(*schema.Set).List()
	networkUris := make([]utils.Nstring, 0)

	for _, raw := range networkUriList {
		networkUris = append(networkUris, utils.NewNstring(raw.(string)))
	}
	uplinkSet.NetworkURIs = networkUris

	fcNetworkUriList := d.Get("fc_network_uris").(*schema.Set).List()
	fcNetworkUris := make([]utils.Nstring, 0)

	for _, raw := range fcNetworkUriList {
		fcNetworkUris = append(fcNetworkUris, utils.NewNstring(raw.(string)))
	}
	uplinkSet.FcNetworkURIs = fcNetworkUris

	fcoeNetworkUriList := d.Get("fcoe_network_uris").(*schema.Set).List()
	fcoeNetworkUris := make([]utils.Nstring, 0)

	for _, raw := range fcoeNetworkUriList {
		fcoeNetworkUris = append(fcoeNetworkUris, utils.NewNstring(raw.(string)))
	}
	uplinkSet.FcoeNetworkURIs = fcoeNetworkUris

	// Getting list of port config info
	portConfigInfosList := d.Get("port_config_infos").(*schema.Set).List()
	portConfigInfos := make([]ov.PortConfigInfos, 0)

	for _, raw := range portConfigInfosList {
		portConfigInfo := raw.(map[string]interface{})

		desiredSpeed := portConfigInfo["desired_speed"].(string)
		portUri := portConfigInfo["port_uri"].(string)

		enclosureLocation := ov.LocationEntries{
			Value: portConfigInfo["enclosure_uri"].(string),
			Type:  Enclosure,
		}
		locationEntries := make([]ov.LocationEntries, 0)
		locationEntries = append(locationEntries, enclosureLocation)

		bayLocation := ov.LocationEntries{
			Value: portConfigInfo["bay_number"].(string),
			Type:  Bay,
		}
		locationEntries = append(locationEntries, bayLocation)

		portLocation := ov.LocationEntries{
			Value: portConfigInfo["port_number"].(string),
			Type:  Port,
		}
		locationEntries = append(locationEntries, portLocation)

		location := ov.Location{
			LocationEntries: locationEntries,
		}

		expectedNeighbor := ov.ExpectedNeighbor{
			RemoteChassisId: portConfigInfo["remote_chassis_id"].(string),
			RemotePortId:    portConfigInfo["remote_port_id"].(string),
		}

		portConfigInfos = append(portConfigInfos, ov.PortConfigInfos{
			DesiredSpeed:     desiredSpeed,
			PortUri:          portUri,
			ExpectedNeighbor: &expectedNeighbor,
			Location:         location,
		})
	}

	uplinkSet.PortConfigInfos = portConfigInfos

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
	d.Set("etag", uplinkSet.Etag)
	d.Set("reachability", uplinkSet.Reachability)
	d.Set("network_type", uplinkSet.NetworkType)
	d.Set("ethernet_network_type", uplinkSet.EthernetNetworkType)
	d.Set("port_config_infos", uplinkSet.PortConfigInfos)
	d.Set("connection_mode", uplinkSet.ConnectionMode)
	d.Set("lacptimer", uplinkSet.LacpTimer)
	d.Set("native_network_uri", uplinkSet.NativeNetworkUri)
	d.Set("fc_mode", uplinkSet.FcMode)

	return nil
}

func resourceUplinkSetUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	const Bay = "Bay"
	const Enclosure = "Enclosure"
	const Port = "Port"

	uplinkSet := ov.UplinkSet{
		Name:                           d.Get("name").(string),
		LogicalInterconnectURI:         utils.NewNstring(d.Get("logical_interconnect_uri").(string)),
		ConnectionMode:                 d.Get("connection_mode").(string),
		NetworkType:                    d.Get("network_type").(string),
		EthernetNetworkType:            d.Get("ethernet_network_type").(string),
		Type:                           d.Get("type").(string),
		URI:                            utils.NewNstring(d.Get("uri").(string)),
		ManualLoginRedistributionState: d.Get("manual_login_redistribution_state").(string),
	}

	networkUriList := d.Get("network_uris").(*schema.Set).List()
	networkUris := make([]utils.Nstring, 0)

	for _, raw := range networkUriList {
		networkUris = append(networkUris, utils.NewNstring(raw.(string)))
	}
	uplinkSet.NetworkURIs = networkUris

	fcNetworkUriList := d.Get("fc_network_uris").(*schema.Set).List()
	fcNetworkUris := make([]utils.Nstring, 0)

	for _, raw := range fcNetworkUriList {
		fcNetworkUris = append(fcNetworkUris, utils.NewNstring(raw.(string)))
	}
	uplinkSet.FcNetworkURIs = fcNetworkUris

	fcoeNetworkUriList := d.Get("fcoe_network_uris").(*schema.Set).List()
	fcoeNetworkUris := make([]utils.Nstring, 0)

	for _, raw := range fcoeNetworkUriList {
		fcoeNetworkUris = append(fcoeNetworkUris, utils.NewNstring(raw.(string)))
	}
	uplinkSet.FcoeNetworkURIs = fcoeNetworkUris

	// Getting list of port config info
	portConfigInfosList := d.Get("port_config_infos").(*schema.Set).List()
	portConfigInfos := make([]ov.PortConfigInfos, 0)

	for _, raw := range portConfigInfosList {
		portConfigInfo := raw.(map[string]interface{})

		desiredSpeed := portConfigInfo["desired_speed"].(string)
		portUri := portConfigInfo["port_uri"].(string)

		enclosureLocation := ov.LocationEntries{
			Value: portConfigInfo["enclosure_uri"].(string),
			Type:  Enclosure,
		}
		locationEntries := make([]ov.LocationEntries, 0)
		locationEntries = append(locationEntries, enclosureLocation)

		bayLocation := ov.LocationEntries{
			Value: portConfigInfo["bay_number"].(string),
			Type:  Bay,
		}
		locationEntries = append(locationEntries, bayLocation)

		portLocation := ov.LocationEntries{
			Value: portConfigInfo["port_number"].(string),
			Type:  Port,
		}
		locationEntries = append(locationEntries, portLocation)

		location := ov.Location{
			LocationEntries: locationEntries,
		}

		expectedNeighbor := ov.ExpectedNeighbor{
			RemoteChassisId: portConfigInfo["remote_chassis_id"].(string),
			RemotePortId:    portConfigInfo["remote_port_id"].(string),
		}

		portConfigInfos = append(portConfigInfos, ov.PortConfigInfos{
			DesiredSpeed:     desiredSpeed,
			PortUri:          portUri,
			ExpectedNeighbor: &expectedNeighbor,
			Location:         location,
		})
	}

	uplinkSet.PortConfigInfos = portConfigInfos

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
