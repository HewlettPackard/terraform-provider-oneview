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

func dataSourceLogicalInterconnect() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLogicalInterconnectRead,

		Schema: map[string]*schema.Schema{
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enclosure_uris": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"ethernet_settings": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dependent_resource_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"interconnect_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
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
				},
			},
			"fusion_domain_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"interconnect_map": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interconnect_map_entries": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enclosure_index": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"interconnect_uri": {

										Type:     schema.TypeString,
										Computed: true,
									},
									"location": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"location_entries": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"value": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"logical_downlink_uri": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"permitted_interconnect_type_uri": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"interconnects": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"logical_interconnect_group_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"snmp_configuration": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"read_community": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"system_contact": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_access": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						"trap_destination": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"community_string": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"enet_trap_categories": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
									"fc_trap_categories": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
									"vcm_trap_categories": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
									"trap_destination": {
										Type:     schema.TypeString,
										Required: true,
									},
									"trap_format": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "SNMPv1",
									},
									"trap_severities": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
								},
							},
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"v3_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"stacking_health": {
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

func dataSourceLogicalInterconnectRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	id := d.Get("name").(string)
	logInt, err := config.ovClient.GetLogicalInterconnectById(id)

	if err != nil || logInt.URI.IsNil() {
		d.SetId("")
		return nil
	}

	d.SetId(id)

	d.Set("category", logInt.Category)
	d.Set("description", logInt.Description)
	d.Set("fusion_domain_uri", logInt.FusionDomainUri)
	d.Set("logical_interconnect_group_uri", logInt.LogicalInterconnectGroupUri)
	d.Set("name", logInt.Name)
	d.Set("stacking_health", logInt.StackingHealth)
	d.Set("type", logInt.Type)
	d.Set("uri", logInt.URI)
	d.Set("enclosure_uris", logInt.EnclosureUris)

	if logInt.EthernetSettings != nil {
		ethernetSettings := make([]map[string]interface{}, 0)

		ethernetSettings = append(ethernetSettings, map[string]interface{}{
			"category":               logInt.EthernetSettings.Category,
			"dependent_resource_uri": logInt.EthernetSettings.DependentResourceUri.String(),
			"description":            logInt.EthernetSettings.Description.String(),
			"id":                     logInt.EthernetSettings.ID,
			"interconnect_type":      logInt.EthernetSettings.InterconnectType,
			"name":                   logInt.EthernetSettings.Name,
			"type":                   logInt.EthernetSettings.Type,
			"uri":                    logInt.EthernetSettings.URI.String(),
		})
		d.Set("ethernet_settings", ethernetSettings)
	}

	if logInt.InterconnectMap != nil {
		interconnectMapEntries := make([]map[string]interface{}, 0, len(logInt.InterconnectMap.InterconnectMapEntries))
		for _, interconnectMapEntry := range logInt.InterconnectMap.InterconnectMapEntries {
			location := make([]map[string]interface{}, 0, 1)
			locationEntries := make([]map[string]interface{}, 0, len(interconnectMapEntry.Location.LocationEntries))
			for _, locationEntry := range interconnectMapEntry.Location.LocationEntries {
				locationEntries = append(locationEntries, map[string]interface{}{
					"type":  locationEntry.Type,
					"value": locationEntry.Value,
				})
			}
			location = append(location, map[string]interface{}{
				"location_entries": locationEntries,
			})
			interconnectMapEntries = append(interconnectMapEntries, map[string]interface{}{
				"location":                        location,
				"logical_downlink_uri":            interconnectMapEntry.LogicalDownlinkUri.String(),
				"permitted_interconnect_type_uri": interconnectMapEntry.PermittedInterconnectTypeUri.String(),
				"interconnect_uri":                interconnectMapEntry.InterconnectUri,
				"enclosure_index":                 interconnectMapEntry.EnclosureIndex,
			})
		}
		interconnectMap := make([]map[string]interface{}, 0, 1)
		interconnectMap = append(interconnectMap, map[string]interface{}{
			"interconnect_map_entries": interconnectMapEntries,
		})
		d.Set("interconnect_map", interconnectMap)
	}

	if logInt.Interconnects != nil {
		interconnects := make([]interface{}, len(logInt.Interconnects))
		for i, interconnect := range logInt.Interconnects {
			interconnects[i] = interconnect
		}

		d.Set("interconnects", interconnects)
	}

	if logInt.SnmpConfiguration != nil {
		trapDestinations := make([]map[string]interface{}, 0, len(logInt.SnmpConfiguration.TrapDestinations))
		for _, trapDestination := range logInt.SnmpConfiguration.TrapDestinations {
			trapDestinations = append(trapDestinations, map[string]interface{}{
				"trap_destination":     trapDestination.TrapDestination,
				"community_string":     trapDestination.CommunityString,
				"trap_format":          trapDestination.TrapFormat,
				"enet_trap_categories": trapDestination.EnetTrapCategories,
				"fc_trap_categories":   trapDestination.FcTrapCategories,
				"vcm_trap_categories":  trapDestination.VcmTrapCategories,
				"trap_severities":      trapDestination.TrapSeverities,
			})
		}

		snmpConfiguration := make([]map[string]interface{}, 0, 1)
		snmpConfiguration = append(snmpConfiguration, map[string]interface{}{
			"enabled":          *logInt.SnmpConfiguration.Enabled,
			"v3_enabled":       *logInt.SnmpConfiguration.V3Enabled,
			"read_community":   logInt.SnmpConfiguration.ReadCommunity,
			"snmp_access":      logInt.SnmpConfiguration.SnmpAccess,
			"system_contact":   logInt.SnmpConfiguration.SystemContact,
			"type":             logInt.SnmpConfiguration.Type,
			"trap_destination": trapDestinations,
		})
		d.Set("snmp_configuration", snmpConfiguration)
	}
	return nil
}
