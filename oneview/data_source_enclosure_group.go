// (C) Copyright 2018 Hewlett Packard Enterprise Development LP
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

func dataSourceEnclosureGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEnclosureGroupRead,

		Schema: map[string]*schema.Schema{
			"ambient_temperature_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_logical_interconnect_groups": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enclosure_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"enclosure_type_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"initial_scope_uris": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"interconnect_bay_mapping_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"interconnect_bay_mappings": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enclosure_index": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"interconnect_bay": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"logical_interconnect_group_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ip_addressing_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_range_uris": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port_mapping_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"port_mappings": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interconnect_bay": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"mid_plane_port": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"power_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scopes_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"stacking_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceEnclosureGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	enclosureGroup, err := config.ovClient.GetEnclosureGroupByName(d.Get("name").(string))
	if err != nil || enclosureGroup.URI.IsNil() {
		d.SetId("")
		return nil
	}

	d.SetId(d.Get("name").(string))
	d.Set("ambient_temperature_mode", enclosureGroup.AmbientTemperatureMode)
	d.Set("associated_logical_interconnect_groups", enclosureGroup.AssociatedLogicalInterconnectGroups)
	d.Set("category", enclosureGroup.Category)
	d.Set("description", enclosureGroup.Description)
	d.Set("etag", enclosureGroup.ETAG)
	d.Set("enclosure_count", enclosureGroup.EnclosureCount)
	d.Set("enclosure_type_uri", enclosureGroup.EnclosureTypeUri.String())
	d.Set("initial_scope_uris", enclosureGroup.InitialScopeUris)
	d.Set("interconnect_bay_mapping_count", enclosureGroup.InterconnectBayMappingCount)
	d.Set("interconnect_bay_mappings", enclosureGroup.InterconnectBayMappings)
	d.Set("ip_addressing_mode", enclosureGroup.IpAddressingMode)
	d.Set("ip_range_uris", enclosureGroup.IpRangeUris)
	d.Set("name", enclosureGroup.Name)
	d.Set("port_mapping_count", enclosureGroup.PortMappingCount)
	d.Set("port_mappings", enclosureGroup.PortMappings)
	d.Set("power_mode", enclosureGroup.PowerMode)
	d.Set("scopes_uri", enclosureGroup.ScopesUri.String())
	d.Set("stacking_mode", enclosureGroup.StackingMode)
	d.Set("status", enclosureGroup.Status)
	d.Set("type", enclosureGroup.Type)
	d.Set("uri", enclosureGroup.URI.String())
	return nil
}
