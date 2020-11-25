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

func dataSourceInterconnectType() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceInterconnectTypeRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"interconnect_bay_set": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"interconnect_capabilities": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_bandwidth_in_gbps": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"capabilities": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed: true,
						},
					},
				},
			},
			"downlink_port_capability": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"max_bandwidth_in_gbps": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"total_sub_port": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"port_info": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"downlink_capable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"port_capabilities": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"port_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"port_number": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"paired_port_name": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						"uplink_capable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"created": {
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
		},
	}
}

func dataSourceInterconnectTypeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	interconnectType, err := config.ovClient.GetInterconnectTypeByName(d.Get("name").(string))
	if err != nil || interconnectType.URI.IsNil() {
		d.SetId("")

		return nil
	}
	d.SetId(d.Get("name").(string))
	d.Set("name", interconnectType.Name)
	d.Set("type", interconnectType.Type)
	d.Set("created", interconnectType.Created)
	d.Set("modified", interconnectType.Modified)
	d.Set("uri", interconnectType.URI.String())
	d.Set("status", interconnectType.Status)
	d.Set("category", interconnectType.Category)
	d.Set("state", interconnectType.State)
	d.Set("port_info", interconnectType.PortInfos)
	d.Set("etag", interconnectType.ETAG)
	d.Set("description", interconnectType.Description)

	DownlinkPortCapability := make([]map[string]interface{}, 0, 1)
	DownlinkPortCapability = append(DownlinkPortCapability, map[string]interface{}{
		"created":               interconnectType.DownlinkPortCapability.Created,
		"max_bandwidth_in_gbps": interconnectType.DownlinkPortCapability.MaxBandwidthInGbps,
		"uri":                   interconnectType.DownlinkPortCapability.URI,
	})
	d.Set("downlink_port_capability", DownlinkPortCapability)

	InterconnectCapabilities := make([]map[string]interface{}, 0, 1)

	capabilitiesMap := make([]interface{}, len(interconnectType.InterconnectCapabilities.Capabilities))
	for i, capabilities := range interconnectType.InterconnectCapabilities.Capabilities {
		capabilitiesMap[i] = capabilities
	}

	InterconnectCapabilities = append(InterconnectCapabilities, map[string]interface{}{

		"capabilities":          schema.NewSet(schema.HashString, capabilitiesMap),
		"max_bandwidth_in_gbps": interconnectType.InterconnectCapabilities.MaxBandwidthInGbps,
	})
	d.Set("interconnect_capabilities", InterconnectCapabilities)

	return nil
}
