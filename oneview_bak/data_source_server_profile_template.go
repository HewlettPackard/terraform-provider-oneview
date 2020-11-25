// (C) Copyright 2019 Hewlett Packard Enterprise Development LP
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
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceServerProfileTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServerProfileTemplateRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"boot_order": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"affinity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network": {
				Optional: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"function_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"port_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"requested_mbps": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"hide_unused_flex_nics": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_hardware_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enclosure_group": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial_number_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"wwn_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mac_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceServerProfileTemplateRead(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*Config)
	name := d.Get("name").(string)

	spt, err := config.ovClient.GetProfileTemplateByName(name)
	if err != nil || spt.URI.IsNil() {
		d.SetId("")
		return nil
	}

	d.SetId(name)
	d.Set("name", spt.Name)
	d.Set("type", spt.Type)
	d.Set("affinity", spt.Affinity)
	d.Set("uri", spt.URI.String())

	enclosureGroup, err := config.ovClient.GetEnclosureGroupByUri(spt.EnclosureGroupURI)
	if err != nil {
		return err
	}
	d.Set("enclosure_group", enclosureGroup.Name)

	serverHardwareType, err := config.ovClient.GetServerHardwareTypeByUri(spt.ServerHardwareTypeURI)
	if err != nil {
		return err
	}

	d.Set("server_hardware_type", serverHardwareType.Name)

	d.Set("etag", spt.ETAG)
	d.Set("serial_number_type", spt.SerialNumberType)
	d.Set("wwn_type", spt.WWNType)
	d.Set("mac_type", spt.MACType)
	d.Set("hide_unused_flex_nics", spt.HideUnusedFlexNics)

	var connections []ov.Connection
	if len(spt.ConnectionSettings.Connections) != 0 {
		connections = spt.ConnectionSettings.Connections
	}
	if len(connections) != 0 {
		networks := make([]map[string]interface{}, 0, len(connections))
		for _, rawNet := range connections {
			networks = append(networks, map[string]interface{}{
				"name":           rawNet.Name,
				"function_type":  rawNet.FunctionType,
				"network_uri":    rawNet.NetworkURI.String(),
				"port_id":        rawNet.PortID,
				"requested_mbps": rawNet.RequestedMbps,
				"id":             rawNet.ID,
			})
		}
		d.Set("network", networks)
	}

	if spt.Boot.ManageBoot {
		bootOrder := make([]interface{}, len(spt.Boot.Order))
		for i, currBoot := range spt.Boot.Order {
			bootOrder[i] = currBoot
		}
		d.Set("boot_order", bootOrder)
	}

	return nil
}
