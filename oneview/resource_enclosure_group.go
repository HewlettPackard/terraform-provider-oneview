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
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceEnclosureGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceEnclosureGroupCreate,
		Read:   resourceEnclosureGroupRead,
		Update: resourceEnclosureGroupUpdate,
		Delete: resourceEnclosureGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ambient_temperature_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_logical_interconnect_groups": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"category": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "enclosure-groups",
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enclosure_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"enclosure_type_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"initial_scope_uris": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"interconnect_bay_mapping_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"interconnect_bay_mappings": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enclosure_index": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"interconnect_bay": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"logical_interconnect_group_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ip_addressing_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_range_uris": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_mapping_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"port_mappings": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interconnect_bay": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mid_plane_port": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"power_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scopes_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"stacking_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func resourceEnclosureGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	enclosureGroup := ov.EnclosureGroup{
		Name: d.Get("name").(string),
	}
	if val, ok := d.GetOk("interconnect_bay_mappings"); ok {
		rawInterconnectBayMappings := val.(*schema.Set).List()
		interconnectBayMappings := make([]ov.InterconnectBayMap, 0)
		for _, raw := range rawInterconnectBayMappings {
			interconnectBayMappingItem := raw.(map[string]interface{})
			interconnectBayMappings = append(interconnectBayMappings, ov.InterconnectBayMap{
				EnclosureIndex:              interconnectBayMappingItem["enclosure_index"].(int),
				InterconnectBay:             interconnectBayMappingItem["interconnect_bay"].(int),
				LogicalInterconnectGroupUri: utils.NewNstring(interconnectBayMappingItem["logical_interconnect_group_uri"].(string))})
		}
		enclosureGroup.InterconnectBayMappings = interconnectBayMappings
	}
	if val, ok := d.GetOk("initial_scope_uris"); ok {
		rawinitialScopeUris := val.(*schema.Set).List()
		initialScopeUris := make([]utils.Nstring, len(rawinitialScopeUris))
		for i, rawData := range rawinitialScopeUris {
			scope, _ := config.ovClient.GetScopeByName(rawData.(string))
			initialScopeUris[i] = utils.Nstring(scope.URI)
		}
		enclosureGroup.InitialScopeUris = initialScopeUris
	}
	if val, ok := d.GetOk("ambient_temperature_mode"); ok {
		enclosureGroup.AmbientTemperatureMode = val.(string)
	}
	if val, ok := d.GetOk("ambient_temperature_mode"); ok {
		enclosureGroup.AmbientTemperatureMode = val.(string)
	}
	if val, ok := d.GetOk("enclosure_count"); ok {
		enclosureGroup.EnclosureCount = val.(int)
	}
	if val, ok := d.GetOk("power_mode"); ok {
		enclosureGroup.PowerMode = val.(string)
	}
	if val, ok := d.GetOk("ip_addressing_mode"); ok {
		enclosureGroup.IpAddressingMode = val.(string)
	}
	if val, ok := d.GetOk("ip_range_uris"); ok {
		rawIPRangeUris := val.(*schema.Set).List()
		ipRangeUris := make([]utils.Nstring, 0)
		for _, rawData := range rawIPRangeUris {
			ipRangeUris = append(ipRangeUris, utils.Nstring(rawData.(string)))
		}
		enclosureGroup.IpRangeUris = ipRangeUris
	}
	enclosureGroupError := config.ovClient.CreateEnclosureGroup(enclosureGroup)
	d.SetId(d.Get("name").(string))
	if enclosureGroupError != nil {
		d.SetId("")
		return enclosureGroupError
	}
	return resourceEnclosureGroupRead(d, meta)
}
func resourceEnclosureGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	enclosureGroup, err := config.ovClient.GetEnclosureGroupByName(d.Id())
	if err != nil || enclosureGroup.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.Set("ambient_temperature_mode", enclosureGroup.AmbientTemperatureMode)
	d.Set("associated_logical_interconnect_groups", enclosureGroup.AssociatedLogicalInterconnectGroups)
	d.Set("category", enclosureGroup.Category)
	d.Set("description", enclosureGroup.Description)
	d.Set("eTag", enclosureGroup.ETAG)
	d.Set("enclosure_count", enclosureGroup.EnclosureCount)
	d.Set("enclosure_type_uri", enclosureGroup.EnclosureTypeUri.String())
	d.Set("initial_scope_uris", enclosureGroup.InitialScopeUris)
	d.Set("interconnect_bay_mapping_count", enclosureGroup.InterconnectBayMappingCount)
	d.Set("interconnect_bay_mappings", enclosureGroup.InterconnectBayMappings)
	d.Set("ip_addressing_mode", enclosureGroup.IpAddressingMode)
	d.Set("ip_range_uris", enclosureGroup.IpRangeUris)
	d.Set("name", enclosureGroup.Name)
	d.Set("os_deployment_settings", enclosureGroup.OsDeploymentSettings)
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
func resourceEnclosureGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	enclosureGroup := ov.EnclosureGroup{
		URI:                         utils.NewNstring(d.Get("uri").(string)),
		InterconnectBayMappingCount: d.Get("interconnect_bay_mapping_count").(int),
		Type:                        d.Get("type").(string),
		StackingMode:                d.Get("stacking_mode").(string),
	}

	rawInterconnectBayMappings := d.Get("interconnect_bay_mappings").(*schema.Set).List()
	interconnectBayMappings := make([]ov.InterconnectBayMap, 0)
	for _, raw := range rawInterconnectBayMappings {
		interconnectBayMappingItem := raw.(map[string]interface{})

		interconnectBayMappings = append(interconnectBayMappings, ov.InterconnectBayMap{
			EnclosureIndex:              interconnectBayMappingItem["enclosure_index"].(int),
			InterconnectBay:             interconnectBayMappingItem["interconnect_bay"].(int),
			LogicalInterconnectGroupUri: utils.NewNstring(interconnectBayMappingItem["logical_interconnect_group_uri"].(string))})
	}
	enclosureGroup.InterconnectBayMappings = interconnectBayMappings

	// Optional Parameters
	if val, ok := d.GetOk("ambient_temperature_mode"); ok {
		enclosureGroup.AmbientTemperatureMode = val.(string)
	}

	if val, ok := d.GetOk("associated_logical_interconnect_groups"); ok {
		rawData := val.(*schema.Set).List()
		associatedLIG := make([]string, 0)
		for _, rawDataItem := range rawData {
			associatedLIG = append(associatedLIG, rawDataItem.(string))
		}
		enclosureGroup.AssociatedLogicalInterconnectGroups = associatedLIG
	}

	if val, ok := d.GetOk("category"); ok {
		enclosureGroup.Category = val.(string)
	}

	if val, ok := d.GetOk("description"); ok {
		enclosureGroup.Description = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk("etag"); ok {
		enclosureGroup.ETAG = val.(string)
	}

	if val, ok := d.GetOk("enclosure_count"); ok {
		enclosureGroup.EnclosureCount = val.(int)
	}

	if val, ok := d.GetOk("enclosure_type_uri"); ok {
		enclosureGroup.EnclosureTypeUri = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk("ip_addressing_mode"); ok {
		enclosureGroup.IpAddressingMode = val.(string)
	}

	if val, ok := d.GetOk("initial_scope_uris"); ok {
		rawinitialScopeUris := val.(*schema.Set).List()
		initialScopeUris := make([]utils.Nstring, 0)
		for i, rawData := range rawinitialScopeUris {
			scope, _ := config.ovClient.GetScopeByName(rawData.(string))
			initialScopeUris[i] = utils.Nstring(scope.URI)
		}
		enclosureGroup.InitialScopeUris = initialScopeUris
	}

	if val, ok := d.GetOk("name"); ok {
		enclosureGroup.Name = val.(string)
	}

	if val, ok := d.GetOk("port_mapping_count"); ok {
		enclosureGroup.PortMappingCount = val.(int)
	}

	if val, ok := d.GetOk("port_mappings"); ok {
		rawPortMappings := val.(*schema.Set).List()
		portMappings := make([]ov.PortMap, 0)
		for _, raw := range rawPortMappings {
			portMappingItem := raw.(map[string]interface{})

			portMappings = append(portMappings, ov.PortMap{
				InterconnectBay: portMappingItem["interconnect_bay"].(int),
				MidplanePort:    portMappingItem["mid_plane_port"].(int)})
		}
		enclosureGroup.PortMappings = portMappings
	}

	if val, ok := d.GetOk("power_mode"); ok {
		enclosureGroup.PowerMode = val.(string)
	}

	if val, ok := d.GetOk("scopes_uri"); ok {
		enclosureGroup.ScopesUri = utils.NewNstring(val.(string))
	}

	err := config.ovClient.UpdateEnclosureGroup(enclosureGroup)
	if err != nil {
		return err
	}
	d.SetId(d.Get("name").(string))

	return resourceEnclosureGroupRead(d, meta)
}

func resourceEnclosureGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteEnclosureGroup(d.Get("name").(string))
	if err != nil {
		return err
	}
	return nil
}
