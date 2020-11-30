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
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceUplinkSet() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUplinkSetRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"logical_interconnect_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_uris": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"fc_network_uris": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"fcoe_network_uris": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"connection_mode": {
				Type:     schema.TypeString,
				Computed: true,
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
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lacptimer": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"manual_login_redistribution_state": {
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
			"port_config_infos": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"desired_speed": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"port_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bay_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enclosure_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"port_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"remote_chassis_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"remote_port_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceUplinkSetRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	uplinkSet, err := config.ovClient.GetUplinkSetByName(d.Get("name").(string))
	if err != nil || uplinkSet.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.SetId(d.Get("name").(string))
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
	d.Set("network_type", uplinkSet.NetworkType)
	d.Set("ethernet_network_type", uplinkSet.EthernetNetworkType)
	d.Set("port_config_infos", uplinkSet.PortConfigInfos)
	d.Set("connection_mode", uplinkSet.ConnectionMode)
	d.Set("lacptimer", uplinkSet.LacpTimer)
	d.Set("native_network_uri", uplinkSet.NativeNetworkUri)
	d.Set("fc_mode", uplinkSet.FcMode)

	return nil
}
