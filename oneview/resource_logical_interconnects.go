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
	"path"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLogicalInterconnect() *schema.Resource {
	return &schema.Resource{
		Read:   resourceLogicalInterconnectRead,
		Update: resourceLogicalInterconnectUpdate,
		Delete: resourceLogicalInterconnectDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"category": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enclosure_uris": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"ethernet_settings": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dependent_resource_uri": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interconnect_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fusion_domain_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interconnect_map": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interconnect_map_entries": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enclosure_index": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"interconnect_uri": {

										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"location": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"location_entries": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"value": {
																Type:     schema.TypeString,
																Optional: true,
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
										Optional: true,
										Computed: true,
									},
									"permitted_interconnect_type_uri": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"port_flap_settings": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"category": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"etag": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"created": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"modified": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"detection_interval": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"port_flap_threshold_per_interval": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"no_of_samples_declare_failures": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"consistency_checking": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port_flap_protection_mode": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"interconnects": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"logical_interconnect_group_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmp_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"read_community": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"system_contact": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"snmp_access": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						"trap_destination": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"community_string": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"enet_trap_categories": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
									"fc_trap_categories": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
									"vcm_trap_categories": {
										Type:     schema.TypeSet,
										Optional: true,
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
										Optional: true,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
								},
							},
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"v3_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"stacking_health": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogicalInterconnectRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	id := d.Id()

	logInt := ov.LogicalInterconnect{}
	allLi, err := config.ovClient.GetLogicalInterconnects("", "", "")
	for _, li := range allLi.Members {
		_, uuid := path.Split(string((li.URI)))
		if id == uuid {
			logInt = li
		}
	}
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
	if logInt.PortFlapProtection != nil {
		portFlapSettings := make([]map[string]interface{}, 0, 1)
		portFlapSettings = append(portFlapSettings, map[string]interface{}{
			"type":                             logInt.PortFlapProtection.Type,
			"uri":                              logInt.PortFlapProtection.URI,
			"category":                         logInt.PortFlapProtection.Category,
			"etag":                             logInt.PortFlapProtection.ETAG,
			"created":                          logInt.PortFlapProtection.Created,
			"modified":                         logInt.PortFlapProtection.Modified,
			"id":                               logInt.PortFlapProtection.ID,
			"name":                             logInt.PortFlapProtection.Name,
			"detection_interval":               logInt.PortFlapProtection.DetectionInterval,
			"port_flap_threshold_per_interval": logInt.PortFlapProtection.PortFlapThresholdPerInterval,
			"no_of_samples_declare_failures":   logInt.PortFlapProtection.NoOfSamplesDeclareFailures,
			"consistency_checking":             logInt.PortFlapProtection.ConsistencyChecking,
			"port_flap_protection_mode":        logInt.PortFlapProtection.PortFlapProtectionMode,
			"description":                      logInt.PortFlapProtection.Description,
			"state":                            logInt.PortFlapProtection.State,
			"status":                           logInt.PortFlapProtection.Status,
		})
		d.Set("port_flap_settings", portFlapSettings)
	}
	return nil
}

func resourceLogicalInterconnectUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	updateType := d.Get("update_type").(string)

	if updateType == "updateComplianceById" {
		id := d.Id()
		err := config.ovClient.UpdateLogicalInterconnectConsistentStateById(id)
		if err != nil {
			return err
		}
		d.SetId(id)
	}

	if updateType == "updatePortFlapSettings" {
		id := d.Id()
		rawPortFlapSetting := d.Get("port_flap_settings").(*schema.Set).List()
		portFlapSettings := ov.PortFlapProtection{}
		for _, val := range rawPortFlapSetting {
			rawlval := val.(map[string]interface{})
			portFlapSettings.Type = rawlval["type"].(string)
			portFlapSettings.URI = utils.Nstring(rawlval["uri"].(string))
			portFlapSettings.Created = rawlval["created"].(string)
			portFlapSettings.Category = utils.Nstring(rawlval["category"].(string))
			portFlapSettings.ETAG = rawlval["etag"].(string)
			portFlapSettings.Modified = rawlval["modified"].(string)
			portFlapSettings.ID = rawlval["id"].(string)
			portFlapSettings.Name = rawlval["name"].(string)
			portFlapSettings.DetectionInterval = rawlval["detection_interval"].(int)
			portFlapSettings.PortFlapThresholdPerInterval = rawlval["port_flap_threshold_per_interval"].(int)
			portFlapSettings.NoOfSamplesDeclareFailures = rawlval["no_of_samples_declare_failures"].(int)
			portFlapSettings.ConsistencyChecking = rawlval["consistency_checking"].(string)
			portFlapSettings.PortFlapProtectionMode = rawlval["port_flap_protection_mode"].(string)
			portFlapSettings.Description = utils.Nstring(rawlval["description"].(string))
			portFlapSettings.State = rawlval["state"].(string)
			portFlapSettings.Status = rawlval["status"].(string)
		}
		err := config.ovClient.UpdateLogicalInterconnectPortFlapSettings(portFlapSettings, id)
		if err != nil {
			return err
		}
		d.SetId(id)
	}
	return resourceLogicalInterconnectRead(d, meta)
}
func resourceLogicalInterconnectDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
