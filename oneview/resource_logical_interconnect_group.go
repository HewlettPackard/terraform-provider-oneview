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
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"reflect"
	"strconv"
)

func resourceLogicalInterconnectGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogicalInterconnectGroupCreate,
		Read:   resourceLogicalInterconnectGroupRead,
		Update: resourceLogicalInterconnectGroupUpdate,
		Delete: resourceLogicalInterconnectGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "logical-interconnect-groupV3",
			},
			"interconnect_bay_set": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"redundancy_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enclosure_indexes": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Set: func(a interface{}) int {
					return a.(int)
				},
			},
			"initial_scope_uris": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"interconnect_map_entry_template": {
				Required: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bay_number": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"interconnect_type_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"enclosure_index": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  1,
						},
					},
				},
			},
			"uplink_set": {
				Optional: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "Ethernet",
						},
						"ethernet_network_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"logical_port_config": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"desired_speed": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "Auto",
									},
									"port_num": {
										Type:     schema.TypeSet,
										Required: true,
										Elem:     &schema.Schema{Type: schema.TypeInt},
										Set: func(a interface{}) int {
											return a.(int)
										},
									},
									"bay_num": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"enclosure_num": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  1,
									},
									"primary_port": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
								},
							},
						},
						"mode": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "Auto",
						},
						"network_uris": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						"lacp_timer": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "Short",
						},
						"native_network_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"internal_network_uris": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"telemetry_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "telemetry-configuration",
						},
						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"sample_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  12,
						},
						"sample_interval": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  300,
						},
					},
				},
			},
			"snmp_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"v3_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "snmp-configuration",
						},
						"read_community": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "public",
						},
						"system_contact": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_access": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						"trap_destination": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"community_string": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"enet_trap_categories": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
									"fc_trap_categories": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
									"vcm_trap_categories": {
										Type:     schema.TypeSet,
										Optional: true,
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
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
								},
							},
						},
					},
				},
			},
			"sflow_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"enabled": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"state": {
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
							Optional: true,
						},
						"sflow_agents": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bay_number": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"enclosure_index": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ip_addr": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"ip_mode": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"status": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"subnet_mask": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"sflow_collectors": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"collector_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"collector_id": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ip_address": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"max_datagram_size": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"max_header_size": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"port": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"sflow_network": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"vlan_id": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"uri": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"sflow_ports": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bay_number": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"collector_id": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"enclosure_index": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"icm_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"port_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"interconnect_settings": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "EthernetInterconnectSettingsV3",
						},
						"fast_mac_cache_failover": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"interconnect_utilization_alert": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"network_loop_protection": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"pause_flood_protection": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"rich_tlv": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"mac_refresh_interval": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  5,
						},
					},
				},
			},
			"igmp_settings": {
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
						"category": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"consistency_checking": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "ExactMatch",
						},
						"created": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dependent_resource_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"etag": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"igmp_snooping": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"prevent_flooding": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"proxy_reporting": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"igmp_idle_timeout_interval": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"igmp_snooping_vlan_ids": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"modified": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"port_flap_settings": {
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
						},
						"port_flap_threshold_per_interval": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"no_of_samples_declare_failures": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"consistency_checking": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_flap_protection_mode": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"state": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"quality_of_service": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "qos-aggregated-configuration",
						},
						"category": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"consistency_checking": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"created": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"etag": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"modified": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"state": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"active_qos_config": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"category": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"config_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"created": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"description": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"downlink_classification_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"etag": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"modified": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"state": {
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
									"uplink_classification_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"uri": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"qos_traffic_classifiers": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"qos_classification_mapping": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dot1p_class_mapping": {
																Type:     schema.TypeList,
																Optional: true,
																Elem:     &schema.Schema{Type: schema.TypeInt},
															},
															"dscp_class_mapping": {
																Type:     schema.TypeList,
																Optional: true,
																Elem:     &schema.Schema{Type: schema.TypeString},
															},
														},
													},
												},
												"qos_traffic_class": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bandwidth_share": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"class_name": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"egress_dot1p_value": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"enabled": {
																Type:     schema.TypeBool,
																Optional: true,
															},
															"max_bandwidth": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"real_time": {
																Type:     schema.TypeBool,
																Optional: true,
															},
															"roce": {
																Type:     schema.TypeBool,
																Optional: true,
															},
															"dcbx_configuration": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"application_protocol": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"default_max_bandwidth": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"default_min_bandwidth": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"priority_code_point": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"priority_flow_control": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"dcbx_ets_port": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"bay_number": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																					"enclosure_index": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"icm_name": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																					"max_bandwidth": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																					"min_bandwidth": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																					"port_name": {
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
										},
									},
								},
							},
						},
					},
				},
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified": {
				Type:     schema.TypeString,
				Computed: true,
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
			"fabric_uri": {
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

func resourceLogicalInterconnectGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	lig := ov.LogicalInterconnectGroup{
		Name: d.Get("name").(string),
		Type: d.Get("type").(string),
	}

	if val, ok := d.GetOk("interconnect_bay_set"); ok {
		lig.InterconnectBaySet = val.(int)
	}

	if val, ok := d.GetOk("redundancy_type"); ok {
		lig.RedundancyType = val.(string)
	}

	if val, ok := d.GetOk("enclosure_indexes"); ok {
		rawEnclosureIndexes := val.(*schema.Set).List()
		enclosureIndexes := make([]int, 0)
		for _, raw := range rawEnclosureIndexes {
			enclosureIndex := raw.(int)
			enclosureIndexes = append(enclosureIndexes, enclosureIndex)
		}
		lig.EnclosureIndexes = enclosureIndexes
	}

	if val, ok := d.GetOk("initial_scope_uris"); ok {
		rawInitialScopeUris := val.(*schema.Set).List()
		initialScopeUris := make([]utils.Nstring, len(rawInitialScopeUris))
		for _, raw := range rawInitialScopeUris {
			initialScopeUri := utils.Nstring(raw.(string))
			initialScopeUris = append(initialScopeUris, initialScopeUri)
		}
		lig.InitialScopeUris = initialScopeUris
	}

	interconnectMapEntryTemplates := make([]ov.InterconnectMapEntryTemplate, 0)
	rawInterconnectMapEntryTemplates := d.Get("interconnect_map_entry_template").(*schema.Set).List()
	for _, raw := range rawInterconnectMapEntryTemplates {
		interconnectMapEntryTemplate := raw.(map[string]interface{})
		interconnectTypeName := interconnectMapEntryTemplate["interconnect_type_name"].(string)
		interconnectType, err := config.ovClient.GetInterconnectTypeByName(interconnectTypeName)
		if err != nil {
			return err
		}
		if interconnectType.URI == "" {
			return fmt.Errorf("Could not find Interconnect Type from name: %s", interconnectTypeName)
		}

		enclosureLocation := ov.LocationEntry{
			RelativeValue: interconnectMapEntryTemplate["enclosure_index"].(int),
			Type:          "Enclosure",
		}
		locationEntries := make([]ov.LocationEntry, 0)
		locationEntries = append(locationEntries, enclosureLocation)

		bayLocation := ov.LocationEntry{
			RelativeValue: interconnectMapEntryTemplate["bay_number"].(int),
			Type:          "Bay",
		}
		locationEntries = append(locationEntries, bayLocation)
		logicalLocation := ov.LogicalLocation{
			LocationEntries: locationEntries,
		}
		interconnectMapEntryTemplates = append(interconnectMapEntryTemplates, ov.InterconnectMapEntryTemplate{
			LogicalLocation:              logicalLocation,
			EnclosureIndex:               interconnectMapEntryTemplate["enclosure_index"].(int),
			PermittedInterconnectTypeUri: interconnectType.URI,
		})
	}
	interconnectMapTemplate := ov.InterconnectMapTemplate{
		InterconnectMapEntryTemplates: interconnectMapEntryTemplates,
	}
	lig.InterconnectMapTemplate = &interconnectMapTemplate

	uplinkSetCount := d.Get("uplink_set.#").(int)
	uplinkSets := make([]ov.UplinkSets, 0)
	for i := 0; i < uplinkSetCount; i++ {
		uplinkSetPrefix := fmt.Sprintf("uplink_set.%d", i)
		uplinkSet := ov.UplinkSets{}
		if val, ok := d.GetOk(uplinkSetPrefix + ".name"); ok {
			uplinkSet.Name = val.(string)
		}
		if val, ok := d.GetOk(uplinkSetPrefix + ".network_type"); ok {
			uplinkSet.NetworkType = val.(string)
		}
		if val, ok := d.GetOk(uplinkSetPrefix + ".ethernet_network_type"); ok {
			uplinkSet.EthernetNetworkType = val.(string)
		}
		if val, ok := d.GetOk(uplinkSetPrefix + ".mode"); ok {
			uplinkSet.Mode = val.(string)
		}
		if val, ok := d.GetOk(uplinkSetPrefix + ".lacp_timer"); ok {
			uplinkSet.LacpTimer = val.(string)
		}
		if val, ok := d.GetOk(uplinkSetPrefix + ".native_network_uri"); ok {
			uplinkSet.NativeNetworkUri = utils.NewNstring(val.(string))
		}

		logicalPortCount := d.Get(uplinkSetPrefix + ".logical_port_config.#").(int)
		logicalPorts := make([]ov.LogicalPortConfigInfo, 0)
		for i := 0; i < logicalPortCount; i++ {
			logicalPortPrefix := fmt.Sprintf(uplinkSetPrefix+".logical_port_config.%d", i)
			rawPortLocations := d.Get(logicalPortPrefix + ".port_num").(*schema.Set).List()
			for _, raw := range rawPortLocations {
				logicalPort := ov.LogicalPortConfigInfo{}

				if val, ok := d.GetOk(logicalPortPrefix + ".desired_speed"); ok {
					logicalPort.DesiredSpeed = val.(string)
				}

				locationEntries := make([]ov.LocationEntry, 0)
				enclosureLocation := ov.LocationEntry{
					RelativeValue: d.Get(logicalPortPrefix + ".enclosure_num").(int),
					Type:          "Enclosure",
				}
				locationEntries = append(locationEntries, enclosureLocation)

				bayLocation := ov.LocationEntry{
					RelativeValue: d.Get(logicalPortPrefix + ".bay_num").(int),
					Type:          "Bay",
				}
				locationEntries = append(locationEntries, bayLocation)

				portLocation := ov.LocationEntry{
					RelativeValue: raw.(int),
					Type:          "Port",
				}
				locationEntries = append(locationEntries, portLocation)

				logicalLocation := ov.LogicalLocation{
					LocationEntries: locationEntries,
				}

				logicalPort.LogicalLocation = logicalLocation
				if _, ok := d.GetOk(logicalPortPrefix + ".primary_port"); ok {
					if uplinkSet.PrimaryPort == nil {
						uplinkSet.PrimaryPort = &logicalLocation
					}
				}

				logicalPorts = append(logicalPorts, logicalPort)
			}

		}
		uplinkSet.LogicalPortConfigInfos = logicalPorts

		rawNetUris := d.Get(uplinkSetPrefix + ".network_uris").(*schema.Set).List()
		netUris := make([]utils.Nstring, 0)
		for _, raw := range rawNetUris {
			netUris = append(netUris, utils.NewNstring(raw.(string)))
		}
		uplinkSet.NetworkUris = netUris

		uplinkSets = append(uplinkSets, uplinkSet)
	}
	lig.UplinkSets = uplinkSets

	sflowConfigurationPrefix := fmt.Sprintf("sflow_configuration.0")
	sflowConfiguration := ov.SflowConfiguration{}

	sflowAgentCount := d.Get(sflowConfigurationPrefix + ".sflow_agents.#").(int)
	sflowAgents := make([]ov.SflowAgent, 0)
	for i := 0; i < sflowAgentCount; i++ {
		sflowAgentPrefix := fmt.Sprintf(sflowConfigurationPrefix+".sflow_agents.%d", i)
		sflowAgent := ov.SflowAgent{}

		if val, ok := d.GetOk(sflowAgentPrefix + ".bay_number"); ok {
			sflowAgent.BayNumber = val.(int)
		}

		if val, ok := d.GetOk(sflowAgentPrefix + ".enclosure_index"); ok {
			sflowAgent.EnclosureIndex = val.(int)
		}

		if val, ok := d.GetOk(sflowAgentPrefix + ".ip_addr"); ok {
			sflowAgent.IpAddr = val.(string)
		}

		if val, ok := d.GetOk(sflowAgentPrefix + ".ip_mode"); ok {
			sflowAgent.IpMode = val.(string)
		}

		if val, ok := d.GetOk(sflowAgentPrefix + ".status"); ok {
			sflowAgent.Status = val.(string)
		}

		if val, ok := d.GetOk(sflowAgentPrefix + ".subnet_mask"); ok {
			sflowAgent.SubnetMask = val.(string)
		}

		sflowAgents = append(sflowAgents, sflowAgent)
	}
	sflowConfiguration.SflowAgents = sflowAgents

	sflowCollectorCount := d.Get(sflowConfigurationPrefix + ".sflow_collectors.#").(int)
	sflowCollectors := make([]ov.SflowCollector, 0)
	for i := 0; i < sflowCollectorCount; i++ {
		sflowCollectorPrefix := fmt.Sprintf(sflowConfigurationPrefix+".sflow_collectors.%d", i)
		sflowCollector := ov.SflowCollector{}

		if val, ok := d.GetOk(sflowCollectorPrefix + ".collector_enabled"); ok {
			enabled := val.(bool)
			sflowCollector.CollectorEnabled = &enabled
		}

		if val, ok := d.GetOk(sflowCollectorPrefix + ".collector_id"); ok {
			sflowCollector.CollectorId = val.(int)
		}

		if val, ok := d.GetOk(sflowCollectorPrefix + ".ip_address"); ok {
			sflowCollector.IPAddress = val.(string)
		}

		if val, ok := d.GetOk(sflowCollectorPrefix + ".max_datagram_size"); ok {
			sflowCollector.MaxDatagramSize = val.(int)
		}

		if val, ok := d.GetOk(sflowCollectorPrefix + ".max_header_size"); ok {
			sflowCollector.MaxHeaderSize = val.(int)
		}

		if val, ok := d.GetOk(sflowCollectorPrefix + ".name"); ok {
			sflowCollector.Name = val.(string)
		}

		if val, ok := d.GetOk(sflowCollectorPrefix + ".port"); ok {
			sflowCollector.Port = val.(int)
		}

		sflowCollectors = append(sflowCollectors, sflowCollector)
	}
	sflowConfiguration.SflowCollectors = sflowCollectors

	sflowNetworkPrefix := fmt.Sprintf(sflowConfigurationPrefix + ".sflow_network.0")
	sflowNetwork := ov.SflowNetwork{}

	if val, ok := d.GetOk(sflowNetworkPrefix + ".vlan_id"); ok {
		sflowNetwork.VlanId = val.(int)
	}

	if val, ok := d.GetOk(sflowNetworkPrefix + ".uri"); ok {
		sflowNetwork.URI = utils.Nstring(val.(string))
	}

	if val, ok := d.GetOk(sflowNetworkPrefix + ".name"); ok {
		sflowNetwork.Name = val.(string)
	}

	sflowConfiguration.SflowNetwork = &sflowNetwork

	sflowPortCount := d.Get(sflowConfigurationPrefix + ".sflow_ports.#").(int)
	sflowPorts := make([]ov.SflowPort, 0)
	for i := 0; i < sflowPortCount; i++ {
		sflowPortPrefix := fmt.Sprintf(sflowConfigurationPrefix+".sflow_ports.%d", i)
		sflowPort := ov.SflowPort{}

		if val, ok := d.GetOk(sflowPortPrefix + ".bay_number"); ok {
			sflowPort.BayNumber = val.(int)
		}

		if val, ok := d.GetOk(sflowPortPrefix + ".enclosure_index"); ok {
			sflowPort.EnclosureIndex = val.(int)
		}

		if val, ok := d.GetOk(sflowPortPrefix + ".collector_id"); ok {
			sflowPort.CollectorId = val.(int)
		}

		if val, ok := d.GetOk(sflowPortPrefix + ".icm_name"); ok {
			sflowPort.IcmName = val.(string)
		}

		if val, ok := d.GetOk(sflowPortPrefix + ".port_num"); ok {
			sflowPort.PortName = val.(string)
		}

		sflowPorts = append(sflowPorts, sflowPort)
	}
	sflowConfiguration.SflowPorts = sflowPorts

	if val, ok := d.GetOk(sflowConfigurationPrefix + ".category"); ok {
		sflowConfiguration.Category = val.(string)
	}
	if val, ok := d.GetOk(sflowConfigurationPrefix + ".description"); ok {
		sflowConfiguration.Description = utils.NewNstring(val.(string))
	}
	if val, ok := d.GetOk(sflowConfigurationPrefix + ".enabled"); ok {
		enabled := val.(bool)
		sflowConfiguration.Enabled = &enabled
	}
	if val, ok := d.GetOk(sflowConfigurationPrefix + ".name"); ok {
		sflowConfiguration.Name = val.(string)
	}
	if val, ok := d.GetOk(sflowConfigurationPrefix + ".state"); ok {
		sflowConfiguration.State = val.(string)
	}
	if val, ok := d.GetOk(sflowConfigurationPrefix + ".status"); ok {
		sflowConfiguration.Status = val.(string)
	}
	if val, ok := d.GetOk(sflowConfigurationPrefix + ".uri"); ok {
		sflowConfiguration.URI = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk(sflowConfigurationPrefix + ".type"); ok {
		sflowConfiguration.Type = val.(string)
		lig.SflowConfiguration = &sflowConfiguration
	}

	rawInternalNetUris := d.Get("internal_network_uris").(*schema.Set).List()
	internalNetUris := make([]utils.Nstring, len(rawInternalNetUris))
	for i, raw := range rawInternalNetUris {
		internalNetUris[i] = utils.NewNstring(raw.(string))
	}
	lig.InternalNetworkUris = internalNetUris

	telemetryConfigPrefix := fmt.Sprintf("telemetry_configuration.0")
	telemetryConfiguration := ov.TelemetryConfiguration{}
	if val, ok := d.GetOk(telemetryConfigPrefix + ".sample_count"); ok {
		telemetryConfiguration.SampleCount = val.(int)
	}
	if val, ok := d.GetOk(telemetryConfigPrefix + ".sample_interval"); ok {
		telemetryConfiguration.SampleInterval = val.(int)
	}
	if val, ok := d.GetOk(telemetryConfigPrefix + ".enabled"); ok {
		enabled := val.(bool)
		telemetryConfiguration.EnableTelemetry = &enabled
	}
	if telemetryConfiguration != (ov.TelemetryConfiguration{}) {
		telemetryConfiguration.Type = d.Get(telemetryConfigPrefix + ".type").(string)
		lig.TelemetryConfiguration = &telemetryConfiguration
	}

	snmpConfigPrefix := fmt.Sprintf("snmp_configuration.0")
	snmpConfiguration := ov.SnmpConfiguration{}
	if val, ok := d.GetOk(snmpConfigPrefix + ".enabled"); ok {
		enabled := val.(bool)
		snmpConfiguration.Enabled = &enabled
	}
	if val, ok := d.GetOk(snmpConfigPrefix + ".v3_enabled"); ok {
		v3Enabled := val.(bool)
		snmpConfiguration.V3Enabled = &v3Enabled
	}
	if val, ok := d.GetOk(snmpConfigPrefix + ".read_community"); ok {
		snmpConfiguration.ReadCommunity = val.(string)
	}
	if val, ok := d.GetOk(snmpConfigPrefix + ".system_contact"); ok {
		snmpConfiguration.SystemContact = val.(string)
	}
	rawSnmpAccess := d.Get(snmpConfigPrefix + ".snmp_access").(*schema.Set).List()
	snmpAccess := make([]string, len(rawSnmpAccess))
	for i, raw := range rawSnmpAccess {
		snmpAccess[i] = raw.(string)
	}
	snmpConfiguration.SnmpAccess = snmpAccess

	trapDestinationCount := d.Get(snmpConfigPrefix + ".trap_destination.#").(int)
	trapDestinations := make([]ov.TrapDestination, 0, trapDestinationCount)
	for i := 0; i < trapDestinationCount; i++ {
		trapDestinationPrefix := fmt.Sprintf(snmpConfigPrefix+".trap_destination.%d", i)

		rawEnetTrapCategories := d.Get(trapDestinationPrefix + ".enet_trap_categories").(*schema.Set).List()
		enetTrapCategories := make([]string, len(rawEnetTrapCategories))
		for i, raw := range rawEnetTrapCategories {
			enetTrapCategories[i] = raw.(string)
		}

		rawFcTrapCategories := d.Get(trapDestinationPrefix + ".fc_trap_categories").(*schema.Set).List()
		fcTrapCategories := make([]string, len(rawFcTrapCategories))
		for i, raw := range rawFcTrapCategories {
			fcTrapCategories[i] = raw.(string)
		}

		rawVcmTrapCategories := d.Get(trapDestinationPrefix + ".vcm_trap_categories").(*schema.Set).List()
		vcmTrapCategories := make([]string, len(rawVcmTrapCategories))
		for i, raw := range rawVcmTrapCategories {
			vcmTrapCategories[i] = raw.(string)
		}

		rawTrapSeverities := d.Get(trapDestinationPrefix + ".trap_severities").(*schema.Set).List()
		trapSeverities := make([]string, len(rawTrapSeverities))
		for i, raw := range rawTrapSeverities {
			trapSeverities[i] = raw.(string)
		}

		trapDestination := ov.TrapDestination{
			TrapDestination:    d.Get(trapDestinationPrefix + ".trap_destination").(string),
			CommunityString:    d.Get(trapDestinationPrefix + ".community_string").(string),
			TrapFormat:         d.Get(trapDestinationPrefix + ".trap_format").(string),
			EnetTrapCategories: enetTrapCategories,
			FcTrapCategories:   fcTrapCategories,
			VcmTrapCategories:  vcmTrapCategories,
			TrapSeverities:     trapSeverities,
		}
		trapDestinations = append(trapDestinations, trapDestination)
	}
	if trapDestinationCount > 0 {
		snmpConfiguration.TrapDestinations = trapDestinations
	}

	if val, ok := d.GetOk(snmpConfigPrefix + ".type"); ok {
		snmpConfiguration.Type = val.(string)
		lig.SnmpConfiguration = &snmpConfiguration
	}

	interconnectSettingsPrefix := fmt.Sprintf("interconnect_settings")
	if val, ok := d.GetOk(interconnectSettingsPrefix + ".type"); ok {
		interconnectSettings := ov.EthernetSettings{}

		macFailoverEnabled := d.Get(interconnectSettingsPrefix + ".fast_mac_cache_failover").(bool)
		interconnectSettings.EnableFastMacCacheFailover = &macFailoverEnabled

		networkLoopProtectionEnabled := d.Get(interconnectSettingsPrefix + ".network_loop_protection").(bool)
		interconnectSettings.EnableNetworkLoopProtection = &networkLoopProtectionEnabled

		pauseFloodProtectionEnabled := d.Get(interconnectSettingsPrefix + ".pause_flood_protection").(bool)
		interconnectSettings.EnablePauseFloodProtection = &pauseFloodProtectionEnabled

		if val1, ok := d.GetOk(interconnectSettingsPrefix + ".rich_tlv"); ok {
			enabled := val1.(bool)
			interconnectSettings.EnableRichTLV = &enabled
		}

		if val1, ok := d.GetOk(interconnectSettingsPrefix + ".interconnect_utilization_alert"); ok {
			enabled := val1.(bool)
			interconnectSettings.EnableInterconnectUtilizationAlert = &enabled
		}

		if val1, ok := d.GetOk(interconnectSettingsPrefix + ".mac_refresh_interval"); ok {
			interconnectSettings.MacRefreshInterval = val1.(int)
		}

		interconnectSettings.Type = val.(string)
		lig.EthernetSettings = &interconnectSettings
	}

	rawigmpsetting := d.Get("igmp_settings").([]interface{})
	igmpSetting := ov.IgmpSettings{}
	for _, val := range rawigmpsetting {

		rawlval := val.(map[string]interface{})

		enableigmpsnooping := rawlval["igmp_snooping"].(bool)
		enablepreventflooding := rawlval["prevent_flooding"].(bool)
		enableproxyreporting := rawlval["proxy_reporting"].(bool)

		igmpSetting.Created = rawlval["created"].(string)
		igmpSetting.Category = utils.Nstring(rawlval["category"].(string))
		igmpSetting.Type = rawlval["type"].(string)
		igmpSetting.ConsistencyChecking = rawlval["consistency_checking"].(string)
		igmpSetting.Description = rawlval["description"].(string)
		igmpSetting.ETAG = utils.Nstring(rawlval["etag"].(string))
		igmpSetting.EnableIgmpSnooping = &enableigmpsnooping
		igmpSetting.EnablePreventFlooding = &enablepreventflooding
		igmpSetting.EnableProxyReporting = &enableproxyreporting
		igmpSetting.ID = rawlval["id"].(string)
		igmpSetting.IgmpIdleTimeoutInterval = rawlval["igmp_idle_timeout_interval"].(int)
		igmpSetting.IgmpSnoopingVlanIds = rawlval["igmp_snooping_vlan_ids"].(string)
		igmpSetting.Modified = rawlval["modified"].(string)
		igmpSetting.Name = rawlval["name"].(string)
		igmpSetting.State = rawlval["state"].(string)
		igmpSetting.Status = rawlval["status"].(string)
		igmpSetting.URI = utils.Nstring(rawlval["uri"].(string))
	}
	lig.IgmpSettings = &igmpSetting

	portFlapSettingsData := d.Get("port_flap_settings").([]interface{})
	for _, raw := range portFlapSettingsData {

		portFlapSettingRawData := raw.(map[string]interface{})

		portFlapSettingStructure := ov.PortFlapProtection{}
		portFlapSettingStructure.Type = portFlapSettingRawData["type"].(string)
		portFlapSettingStructure.URI = utils.Nstring(portFlapSettingRawData["uri"].(string))
		portFlapSettingStructure.Category = utils.Nstring(portFlapSettingRawData["category"].(string))
		portFlapSettingStructure.ETAG = portFlapSettingRawData["etag"].(string)
		portFlapSettingStructure.Created = portFlapSettingRawData["created"].(string)
		portFlapSettingStructure.Modified = portFlapSettingRawData["modified"].(string)
		portFlapSettingStructure.ID = portFlapSettingRawData["id"].(string)
		portFlapSettingStructure.Name = portFlapSettingRawData["name"].(string)
		portFlapSettingStructure.DetectionInterval = portFlapSettingRawData["detection_interval"].(int)
		portFlapSettingStructure.PortFlapThresholdPerInterval = portFlapSettingRawData["port_flap_threshold_per_interval"].(int)
		portFlapSettingStructure.NoOfSamplesDeclareFailures = portFlapSettingRawData["no_of_samples_declare_failures"].(int)
		portFlapSettingStructure.ConsistencyChecking = portFlapSettingRawData["consistency_checking"].(string)
		portFlapSettingStructure.PortFlapProtectionMode = portFlapSettingRawData["port_flap_protection_mode"].(string)
		portFlapSettingStructure.Description = utils.Nstring(portFlapSettingRawData["description"].(string))
		portFlapSettingStructure.State = portFlapSettingRawData["state"].(string)
		portFlapSettingStructure.Status = portFlapSettingRawData["status"].(string)
		lig.PortFlapProtection = &portFlapSettingStructure
	}

	if val, ok := d.GetOk("quality_of_service"); ok {
		rawQoss := val.([]interface{})
		ovQos := ov.QosConfiguration{}
		if len(rawQoss) != 0 {
			for _, rawQosConfig := range rawQoss {
				rawQos := rawQosConfig.(map[string]interface{})
				rawActiveQosConfigs := rawQos["active_qos_config"].([]interface{})
				ovActiveQosConfig := ov.ActiveQosConfig{}
				if len(rawActiveQosConfigs) != 0 {
					for _, rawActiveQosConfig := range rawActiveQosConfigs {
						activeQosConfig := rawActiveQosConfig.(map[string]interface{})
						rawQosClassifiers := activeQosConfig["qos_traffic_classifiers"].([]interface{})
						ovQosTrafficClassifier := make([]ov.QosTrafficClassifier, 0)
						if len(rawQosClassifiers) != 0 {
							for _, rawQosClassifier := range rawQosClassifiers {
								qosClassifier := rawQosClassifier.(map[string]interface{})
								ovQosClassificationMapping := ov.QosClassificationMapping{}
								rawQosClassificationMappings := qosClassifier["qos_classification_mapping"].([]interface{})
								if rawQosClassificationMappings != nil {
									for _, rawQosClassificationMapping := range rawQosClassificationMappings {
										qosClassificationMapping := rawQosClassificationMapping.(map[string]interface{})
										rawDot1pClassMappings := qosClassificationMapping["dot1p_class_mapping"].([]interface{})
										if qosClassificationMapping["dot1p_class_mapping"] != nil {
											dot1pClassMapping := make([]int, 0)
											for _, raw := range rawDot1pClassMappings {
												dot1pClassMapping = append(dot1pClassMapping, raw.(int))
											}
											ovQosClassificationMapping.Dot1pClassMapping = dot1pClassMapping
										}
										rawDscpClassMappings := qosClassificationMapping["dscp_class_mapping"].([]interface{})
										if qosClassificationMapping["dscp_class_mapping"] != nil {
											dscpClassMapping := make([]string, 0)
											for _, raw := range rawDscpClassMappings {
												dscpClassMapping = append(dscpClassMapping, raw.(string))
											}
											ovQosClassificationMapping.DscpClassMapping = dscpClassMapping
										}
									}
								}
								ovQosTrafficClass := ov.QosTrafficClass{}
								rawQosTrafficClasses := qosClassifier["qos_traffic_class"].([]interface{})
								if len(rawQosTrafficClasses) != 0 {
									for _, rawQosTrafficClass := range rawQosTrafficClasses {
										qosTrafficClass := rawQosTrafficClass.(map[string]interface{})
										rawDcbxConfigurations := qosTrafficClass["dcbx_configuration"].([]interface{})
										ovDcbxConfiguration := ov.DcbxConfigurations{}
										if len(rawDcbxConfigurations) != 0 {
											for _, rawDcbxConfiguration := range rawDcbxConfigurations {
												dcbxConfiguration := rawDcbxConfiguration.(map[string]interface{})
												rawDcbxEtsPorts := dcbxConfiguration["dcbx_ets_port"].([]interface{})
												ovDcbxEtsPorts := make([]ov.DcbxEtsPort, 0)
												if len(rawDcbxEtsPorts) != 0 {
													for _, rawDcbxPort := range rawDcbxEtsPorts {
														dcbxPort := rawDcbxPort.(map[string]interface{})
														ovDcbxEtsPorts = append(ovDcbxEtsPorts, ov.DcbxEtsPort{
															BayNumber:      dcbxPort["bay_number"].(string),
															EnclosureIndex: dcbxPort["enclosure_index"].(int),
															IcmName:        dcbxPort["icm_name"].(string),
															MaxBandwidth:   dcbxPort["max_bandwidth"].(string),
															MinBandwidth:   dcbxPort["min_bandwidth"].(string),
															PortName:       dcbxPort["port_name"].(string),
														})
													}
												}
												ovDcbxConfiguration = ov.DcbxConfigurations{
													ApplicationProtocol:     dcbxConfiguration["application_protocol"].(string),
													DefaultMaximumBandwidth: dcbxConfiguration["default_max_bandwidth"].(string),
													DefaultMinimumBandwidth: dcbxConfiguration["default_min_bandwidth"].(string),
													PriorityCodePoint:       dcbxConfiguration["priority_code_point"].(string),
													PriorityFlowControl:     dcbxConfiguration["priority_flow_control"].(string),
												}
												if len(ovDcbxEtsPorts) != 0 {
													ovDcbxConfiguration.DcbxEtsPorts = ovDcbxEtsPorts
												}
											}
										}
										ovQosTrafficClass = ov.QosTrafficClass{
											BandwidthShare:   qosTrafficClass["bandwidth_share"].(string),
											ClassName:        qosTrafficClass["class_name"].(string),
											EgressDot1pValue: GetIntPointer(qosTrafficClass["egress_dot1p_value"].(int)),
											Enabled:          GetBoolPointer(qosTrafficClass["enabled"].(bool)),
											MaxBandwidth:     qosTrafficClass["max_bandwidth"].(int),
											RealTime:         GetBoolPointer(qosTrafficClass["real_time"].(bool)),
											Roce:             GetBoolPointer(qosTrafficClass["roce"].(bool)),
										}
										if !(reflect.DeepEqual(ovDcbxConfiguration, ov.DcbxConfigurations{})) {
											ovQosTrafficClass.DcbxConfiguration = &ovDcbxConfiguration
										}
									}
								}
								if reflect.DeepEqual(ovQosClassificationMapping, ov.QosClassificationMapping{}) && ovQosTrafficClass == (ov.QosTrafficClass{}) {
									continue
								} else if reflect.DeepEqual(ovQosClassificationMapping, ov.QosClassificationMapping{}) {
									ovQosTrafficClassifier = append(ovQosTrafficClassifier, ov.QosTrafficClassifier{
										QosTrafficClass: &ovQosTrafficClass,
									})
								} else if ovQosTrafficClass == (ov.QosTrafficClass{}) {
									ovQosTrafficClassifier = append(ovQosTrafficClassifier, ov.QosTrafficClassifier{
										QosClassificationMapping: &ovQosClassificationMapping,
									})
								} else {
									ovQosTrafficClassifier = append(ovQosTrafficClassifier, ov.QosTrafficClassifier{
										QosClassificationMapping: &ovQosClassificationMapping,
										QosTrafficClass:          &ovQosTrafficClass,
									})
								}
							}
						}
						ovActiveQosConfig = ov.ActiveQosConfig{
							Category:                   utils.NewNstring(activeQosConfig["category"].(string)),
							ConfigType:                 activeQosConfig["config_type"].(string),
							Created:                    activeQosConfig["created"].(string),
							Description:                utils.NewNstring(activeQosConfig["description"].(string)),
							DownlinkClassificationType: activeQosConfig["downlink_classification_type"].(string),
							ETAG:                       activeQosConfig["etag"].(string),
							Modified:                   activeQosConfig["modified"].(string),
							Name:                       activeQosConfig["name"].(string),
							State:                      activeQosConfig["state"].(string),
							Status:                     activeQosConfig["status"].(string),
							Type:                       activeQosConfig["type"].(string),
							UplinkClassificationType:   activeQosConfig["uplink_classification_type"].(string),
							URI:                        utils.NewNstring(activeQosConfig["uri"].(string)),
							QosTrafficClassifiers:      ovQosTrafficClassifier,
						}
					}
				}
				ovQos = ov.QosConfiguration{
					Category:            rawQos["category"].(string),
					ConsistencyChecking: rawQos["consistency_checking"].(string),
					Created:             rawQos["created"].(string),
					Description:         utils.NewNstring(rawQos["description"].(string)),
					ETAG:                rawQos["etag"].(string),
					Modified:            rawQos["modified"].(string),
					Name:                rawQos["name"].(string),
					State:               rawQos["state"].(string),
					Status:              rawQos["status"].(string),
					Type:                rawQos["type"].(string),
					URI:                 utils.NewNstring(rawQos["uri"].(string)),
					ActiveQosConfig:     &ovActiveQosConfig,
				}
			}

		}
		lig.QosConfiguration = &ovQos
	}
	ligError := config.ovClient.CreateLogicalInterconnectGroup(lig)
	d.SetId(d.Get("name").(string))
	if ligError != nil {
		d.SetId("")
		return ligError
	}
	return resourceLogicalInterconnectGroupRead(d, meta)
}

func resourceLogicalInterconnectGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	logicalInterconnectGroup, err := config.ovClient.GetLogicalInterconnectGroupByName(d.Id())

	if err != nil || logicalInterconnectGroup.URI.IsNil() {
		d.SetId("")
		return nil
	}

	d.Set("name", logicalInterconnectGroup.Name)
	d.Set("type", logicalInterconnectGroup.Type)
	d.Set("created", logicalInterconnectGroup.Created)
	d.Set("modified", logicalInterconnectGroup.Modified)
	d.Set("uri", logicalInterconnectGroup.URI.String())
	d.Set("status", logicalInterconnectGroup.Status)
	d.Set("category", logicalInterconnectGroup.Category)
	d.Set("state", logicalInterconnectGroup.State)
	d.Set("fabric_uri", logicalInterconnectGroup.FabricUri.String())
	d.Set("etag", logicalInterconnectGroup.ETAG)
	d.Set("description", logicalInterconnectGroup.Description)
	d.Set("interconnect_settings.0.interconnect_utilization_alert", logicalInterconnectGroup.EthernetSettings.EnableInterconnectUtilizationAlert)
	d.Set("interconnect_bay_set", logicalInterconnectGroup.InterconnectBaySet)
	d.Set("redundancy_type", logicalInterconnectGroup.RedundancyType)

	if logicalInterconnectGroup.QosConfiguration != nil {
		qosTrafficClassifiers := make([]map[string]interface{}, 0, len(logicalInterconnectGroup.QosConfiguration.ActiveQosConfig.QosTrafficClassifiers))
		for _, qosTrafficClassifier := range logicalInterconnectGroup.QosConfiguration.ActiveQosConfig.QosTrafficClassifiers {
			qosClassificationMapping := make([]map[string]interface{}, 0, 1)
			if qosTrafficClassifier.QosClassificationMapping != nil {
				dot1pClassMappings := make([]interface{}, 0)
				if qosTrafficClassifier.QosClassificationMapping.Dot1pClassMapping != nil {
					for _, dot1pClassMapping := range qosTrafficClassifier.QosClassificationMapping.Dot1pClassMapping {
						dot1pClassMappings = append(dot1pClassMappings, dot1pClassMapping)
					}
				}
				dscpClassMappings := make([]interface{}, 0)
				if qosTrafficClassifier.QosClassificationMapping.DscpClassMapping != nil {
					for _, dscpClassMapping := range qosTrafficClassifier.QosClassificationMapping.DscpClassMapping {
						dscpClassMappings = append(dscpClassMappings, dscpClassMapping)
					}
				}
				qosClassificationMapping = append(qosClassificationMapping, map[string]interface{}{
					"dot1p_class_mapping": dot1pClassMappings,
					"dscp_class_mapping":  dscpClassMappings,
				})
			}
			dcbxConfigurations := make([]map[string]interface{}, 0, 1)
			if qosTrafficClassifier.QosTrafficClass.DcbxConfiguration != nil {
				dcbxEtsPorts := make([]map[string]interface{}, 0, len(qosTrafficClassifier.QosTrafficClass.DcbxConfiguration.DcbxEtsPorts))
				for _, dcbxEtsPort := range qosTrafficClassifier.QosTrafficClass.DcbxConfiguration.DcbxEtsPorts {
					dcbxEtsPorts = append(dcbxEtsPorts, map[string]interface{}{
						"bay_number":      dcbxEtsPort.BayNumber,
						"enclosure_index": dcbxEtsPort.EnclosureIndex,
						"icm_name":        dcbxEtsPort.IcmName,
						"max_bandwidth":   dcbxEtsPort.MaxBandwidth,
						"min_bandwidth":   dcbxEtsPort.MinBandwidth,
						"port_name":       dcbxEtsPort.PortName,
					})
				}
				dcbxConfigurations = append(dcbxConfigurations, map[string]interface{}{
					"application_protocol":  qosTrafficClassifier.QosTrafficClass.DcbxConfiguration.ApplicationProtocol,
					"default_max_bandwidth": qosTrafficClassifier.QosTrafficClass.DcbxConfiguration.DefaultMaximumBandwidth,
					"default_min_bandwidth": qosTrafficClassifier.QosTrafficClass.DcbxConfiguration.DefaultMinimumBandwidth,
					"priority_code_point":   qosTrafficClassifier.QosTrafficClass.DcbxConfiguration.PriorityCodePoint,
					"priority_flow_control": qosTrafficClassifier.QosTrafficClass.DcbxConfiguration.PriorityFlowControl,
					"dcbx_ets_port":         dcbxEtsPorts,
				})
			}
			qosTrafficClass := make([]map[string]interface{}, 0, 1)
			qosTrafficClass = append(qosTrafficClass, map[string]interface{}{
				"bandwidth_share":    qosTrafficClassifier.QosTrafficClass.BandwidthShare,
				"class_name":         qosTrafficClassifier.QosTrafficClass.ClassName,
				"egress_dot1p_value": *qosTrafficClassifier.QosTrafficClass.EgressDot1pValue,
				"enabled":            *qosTrafficClassifier.QosTrafficClass.Enabled,
				"max_bandwidth":      qosTrafficClassifier.QosTrafficClass.MaxBandwidth,
				"real_time":          *qosTrafficClassifier.QosTrafficClass.RealTime,
				"roce":               *qosTrafficClassifier.QosTrafficClass.Roce,
				"dcbx_configuration": dcbxConfigurations,
			})
			qosTrafficClassifiers = append(qosTrafficClassifiers, map[string]interface{}{
				"qos_traffic_class":          qosTrafficClass,
				"qos_classification_mapping": qosClassificationMapping,
			})
		}
		activeQosConfig := make([]map[string]interface{}, 0, 1)
		activeQosConfig = append(activeQosConfig, map[string]interface{}{
			"category":                     logicalInterconnectGroup.QosConfiguration.ActiveQosConfig.Category,
			"config_type":                  logicalInterconnectGroup.QosConfiguration.ActiveQosConfig.ConfigType,
			"created":                      logicalInterconnectGroup.QosConfiguration.ActiveQosConfig.Created,
			"description":                  logicalInterconnectGroup.QosConfiguration.ActiveQosConfig.Description,
			"downlink_classification_type": logicalInterconnectGroup.QosConfiguration.ActiveQosConfig.DownlinkClassificationType,
			"etag":                         logicalInterconnectGroup.QosConfiguration.ActiveQosConfig.ETAG,
			"modified":                     logicalInterconnectGroup.QosConfiguration.ActiveQosConfig.Modified,
			"name":                         logicalInterconnectGroup.QosConfiguration.ActiveQosConfig.Name,
			"state":                        logicalInterconnectGroup.QosConfiguration.ActiveQosConfig.State,
			"status":                       logicalInterconnectGroup.QosConfiguration.ActiveQosConfig.Status,
			"type":                         logicalInterconnectGroup.QosConfiguration.ActiveQosConfig.Type,
			"uplink_classification_type":   logicalInterconnectGroup.QosConfiguration.ActiveQosConfig.UplinkClassificationType,
			"uri":                          logicalInterconnectGroup.QosConfiguration.ActiveQosConfig.URI,
			"qos_traffic_classifiers":      qosTrafficClassifiers,
		})
		qosConfiguration := make([]map[string]interface{}, 0, 1)
		qosConfiguration = append(qosConfiguration, map[string]interface{}{
			"type":                 logicalInterconnectGroup.QosConfiguration.Type,
			"category":             logicalInterconnectGroup.QosConfiguration.Category,
			"consistency_checking": logicalInterconnectGroup.QosConfiguration.ConsistencyChecking,
			"created":              logicalInterconnectGroup.QosConfiguration.Created,
			"description":          logicalInterconnectGroup.QosConfiguration.Description,
			"etag":                 logicalInterconnectGroup.QosConfiguration.ETAG,
			"modified":             logicalInterconnectGroup.QosConfiguration.Modified,
			"name":                 logicalInterconnectGroup.QosConfiguration.Name,
			"state":                logicalInterconnectGroup.QosConfiguration.State,
			"status":               logicalInterconnectGroup.QosConfiguration.Status,
			"uri":                  logicalInterconnectGroup.QosConfiguration.URI,
			"active_qos_config":    activeQosConfig,
		})
		d.Set("quality_of_service", qosConfiguration)
	}

	enclosureIndexes := make([]interface{}, len(logicalInterconnectGroup.EnclosureIndexes))
	for i, enclosureIndexVal := range logicalInterconnectGroup.EnclosureIndexes {
		enclosureIndexes[i] = enclosureIndexVal
	}
	d.Set("enclosure_indexes", schema.NewSet(func(a interface{}) int { return a.(int) }, enclosureIndexes))

	initialScopeUris := make([]interface{}, len(logicalInterconnectGroup.InitialScopeUris))
	for i, initialScopeUriVal := range logicalInterconnectGroup.InitialScopeUris {
		initialScopeUris[i] = initialScopeUriVal
	}
	d.Set("initial_scope_uris", schema.NewSet(func(a interface{}) int { return a.(int) }, initialScopeUris))

	interconnectMapEntryTemplates := make([]map[string]interface{}, 0, len(logicalInterconnectGroup.InterconnectMapTemplate.InterconnectMapEntryTemplates))
	for _, interconnectMapEntryTemplate := range logicalInterconnectGroup.InterconnectMapTemplate.InterconnectMapEntryTemplates {
		interconnectType, err := config.ovClient.GetInterconnectTypeByUri(interconnectMapEntryTemplate.PermittedInterconnectTypeUri)
		if err != nil {
			return err
		}
		if interconnectType.Name == "" {
			return fmt.Errorf("Could not find interconnectType with URI %s", interconnectMapEntryTemplate.PermittedInterconnectTypeUri.String())
		}
		var bayNum int
		var enclosureIndex int
		if interconnectMapEntryTemplate.LogicalLocation.LocationEntries[0].Type == "Bay" {
			bayNum = interconnectMapEntryTemplate.LogicalLocation.LocationEntries[0].RelativeValue
			enclosureIndex = interconnectMapEntryTemplate.LogicalLocation.LocationEntries[1].RelativeValue
		} else {
			bayNum = interconnectMapEntryTemplate.LogicalLocation.LocationEntries[1].RelativeValue
			enclosureIndex = interconnectMapEntryTemplate.LogicalLocation.LocationEntries[0].RelativeValue
		}

		interconnectMapEntryTemplates = append(interconnectMapEntryTemplates, map[string]interface{}{
			"interconnect_type_name": interconnectType.Name,
			"bay_number":             bayNum,
			"enclosure_index":        enclosureIndex,
		})
	}

	d.Set("interconnect_map_entry_template", interconnectMapEntryTemplates)

	uplinkSets := make([]map[string]interface{}, 0, len(logicalInterconnectGroup.UplinkSets))
	for i, uplinkSet := range logicalInterconnectGroup.UplinkSets {

		primaryPortEnclosure := 0
		primaryPortBay := 0
		primaryPortPort := 0

		if uplinkSet.PrimaryPort != nil {
			for _, primaryPortLocation := range uplinkSet.PrimaryPort.LocationEntries {
				if primaryPortLocation.Type == "Bay" {
					primaryPortBay = primaryPortLocation.RelativeValue
				}
				if primaryPortLocation.Type == "Enclosure" {
					primaryPortEnclosure = primaryPortLocation.RelativeValue
				}
				if primaryPortLocation.Type == "Port" {
					primaryPortPort = primaryPortLocation.RelativeValue
				}
			}
		}

		logicalPortConfigs := make([]map[string]interface{}, 0, len(uplinkSet.LogicalPortConfigInfos))
		for _, logicalPortConfigInfo := range uplinkSet.LogicalPortConfigInfos {
			portEnclosure := 0
			portBay := 0
			portPort := 0
			primaryPort := false
			for _, portLocation := range logicalPortConfigInfo.LogicalLocation.LocationEntries {
				if portLocation.Type == "Bay" {
					portBay = portLocation.RelativeValue
				}
				if portLocation.Type == "Enclosure" {
					portEnclosure = portLocation.RelativeValue
				}
				if portLocation.Type == "Port" {
					portPort = portLocation.RelativeValue
				}
			}
			if primaryPortEnclosure == portEnclosure && primaryPortBay == portBay && primaryPortPort == portPort {
				primaryPort = true
			}

			portPorts := make([]interface{}, 0)
			portPorts = append(portPorts, portPort)

			included := false
			for j, portConfig := range logicalPortConfigs {
				if portConfig["bay_num"] == portBay && portConfig["enclosure_num"] == portEnclosure {
					included = true
					portSet := logicalPortConfigs[j]["port_num"].(*schema.Set)
					portSet.Add(portPort)
				}
			}

			if included == false {
				logicalPortConfigs = append(logicalPortConfigs, map[string]interface{}{
					"desired_speed": logicalPortConfigInfo.DesiredSpeed,
					"primary_port":  primaryPort,
					"port_num":      schema.NewSet(func(a interface{}) int { return a.(int) }, portPorts),
					"bay_num":       portBay,
					"enclosure_num": portEnclosure,
				})
			}
		}

		//Oneview returns an unordered list so order it to match the configuration file
		logicalPortCount := d.Get("uplink_set." + strconv.Itoa(i) + ".logical_port_config.#").(int)
		oneviewLogicalPortCount := len(logicalPortConfigs)
		for j := 0; j < logicalPortCount; j++ {
			currBay := d.Get("uplink_set." + strconv.Itoa(i) + ".logical_port_config." + strconv.Itoa(j) + ".bay_num").(int)
			for k := 0; k < oneviewLogicalPortCount; k++ {
				if currBay == logicalPortConfigs[k]["bay_num"] && j <= k {
					logicalPortConfigs[j], logicalPortConfigs[k] = logicalPortConfigs[k], logicalPortConfigs[j]
				}
			}
		}

		networkUris := make([]interface{}, len(uplinkSet.NetworkUris))
		for i, networkUri := range uplinkSet.NetworkUris {
			networkUris[i] = networkUri.String()
		}

		uplinkSets = append(uplinkSets, map[string]interface{}{
			"network_type":          uplinkSet.NetworkType,
			"ethernet_network_type": uplinkSet.EthernetNetworkType,
			"name":                  uplinkSet.Name,
			"mode":                  uplinkSet.Mode,
			"lacp_timer":            uplinkSet.LacpTimer,
			"native_network_uri":    uplinkSet.NativeNetworkUri,
			"logical_port_config":   logicalPortConfigs,
			"network_uris":          schema.NewSet(schema.HashString, networkUris),
		})
	}
	uplinkCount := d.Get("uplink_set.#").(int)
	oneviewUplinkCount := len(uplinkSets)
	for i := 0; i < uplinkCount; i++ {
		currUplinkName := d.Get("uplink_set." + strconv.Itoa(i) + ".name").(string)
		for j := 0; j < oneviewUplinkCount; j++ {
			if currUplinkName == uplinkSets[j]["name"] && i <= j {
				uplinkSets[i], uplinkSets[j] = uplinkSets[j], uplinkSets[i]
			}
		}
	}
	d.Set("uplink_set", uplinkSets)

	internalNetworkUris := make([]interface{}, len(logicalInterconnectGroup.InternalNetworkUris))
	for i, internalNetworkUri := range logicalInterconnectGroup.InternalNetworkUris {
		internalNetworkUris[i] = internalNetworkUri
	}
	d.Set("internal_network_uris", internalNetworkUris)

	if logicalInterconnectGroup.SflowConfiguration != nil {
		sflowAgents := make([]map[string]interface{}, 0, len(logicalInterconnectGroup.SflowConfiguration.SflowAgents))
		for i := 0; i < len(logicalInterconnectGroup.SflowConfiguration.SflowAgents); i++ {

			sflowAgents = append(sflowAgents, map[string]interface{}{
				"bay_number":      logicalInterconnectGroup.SflowConfiguration.SflowAgents[i].BayNumber,
				"enclosure_index": logicalInterconnectGroup.SflowConfiguration.SflowAgents[i].EnclosureIndex,
				"ip_addr":         logicalInterconnectGroup.SflowConfiguration.SflowAgents[i].IpAddr,
				"ip_mode":         logicalInterconnectGroup.SflowConfiguration.SflowAgents[i].IpMode,
				"subnet_mask":     logicalInterconnectGroup.SflowConfiguration.SflowAgents[i].SubnetMask,
				"status":          logicalInterconnectGroup.SflowConfiguration.SflowAgents[i].Status,
			})
		}

		sflowCollectors := make([]map[string]interface{}, 0, len(logicalInterconnectGroup.SflowConfiguration.SflowCollectors))
		for i := 0; i < len(logicalInterconnectGroup.SflowConfiguration.SflowCollectors); i++ {

			sflowCollectors = append(sflowCollectors, map[string]interface{}{
				"collector_enabled": logicalInterconnectGroup.SflowConfiguration.SflowCollectors[i].CollectorEnabled,
				"collector_id":      logicalInterconnectGroup.SflowConfiguration.SflowCollectors[i].CollectorId,
				"ip_address":        logicalInterconnectGroup.SflowConfiguration.SflowCollectors[i].IPAddress,
				"max_datagram_size": logicalInterconnectGroup.SflowConfiguration.SflowCollectors[i].MaxDatagramSize,
				"max_header_size":   logicalInterconnectGroup.SflowConfiguration.SflowCollectors[i].MaxHeaderSize,
				"name":              logicalInterconnectGroup.SflowConfiguration.SflowCollectors[i].Name,
				"port":              logicalInterconnectGroup.SflowConfiguration.SflowCollectors[i].Port,
			})
		}
		sflowNetwork := make([]map[string]interface{}, 0, 1)
		if logicalInterconnectGroup.SflowConfiguration.SflowNetwork != nil {
			sflowNetwork = append(sflowNetwork, map[string]interface{}{
				"name":    logicalInterconnectGroup.SflowConfiguration.SflowNetwork.Name,
				"vlan_id": logicalInterconnectGroup.SflowConfiguration.SflowNetwork.VlanId,
				"uri":     logicalInterconnectGroup.SflowConfiguration.SflowNetwork.URI.String(),
			})
		}

		sflowPorts := make([]map[string]interface{}, 0, len(logicalInterconnectGroup.SflowConfiguration.SflowPorts))
		for i := 0; i < len(logicalInterconnectGroup.SflowConfiguration.SflowPorts); i++ {
			sflowPorts = append(sflowPorts, map[string]interface{}{
				"bay_number":      logicalInterconnectGroup.SflowConfiguration.SflowPorts[i].BayNumber,
				"collector_id":    logicalInterconnectGroup.SflowConfiguration.SflowPorts[i].CollectorId,
				"enclosure_index": logicalInterconnectGroup.SflowConfiguration.SflowPorts[i].EnclosureIndex,
				"icm_name":        logicalInterconnectGroup.SflowConfiguration.SflowPorts[i].IcmName,
				"port_name":       logicalInterconnectGroup.SflowConfiguration.SflowPorts[i].PortName,
			})
		}

		sflowConfigurations := make([]map[string]interface{}, 0, 1)

		sflowConfigurations = append(sflowConfigurations, map[string]interface{}{
			"category":         logicalInterconnectGroup.SflowConfiguration.Category,
			"description":      logicalInterconnectGroup.SflowConfiguration.Description.String(),
			"enabled":          *logicalInterconnectGroup.SflowConfiguration.Enabled,
			"name":             logicalInterconnectGroup.SflowConfiguration.Name,
			"state":            logicalInterconnectGroup.SflowConfiguration.State,
			"status":           logicalInterconnectGroup.SflowConfiguration.Status,
			"type":             logicalInterconnectGroup.SflowConfiguration.Type,
			"uri":              logicalInterconnectGroup.SflowConfiguration.URI.String(),
			"sflow_agents":     sflowAgents,
			"sflow_collectors": sflowCollectors,
			"sflow_network":    sflowNetwork,
			"sflow_ports":      sflowPorts,
		})

		d.Set("sflow_configuration", sflowConfigurations)

	}
	telemetryConfigurations := make([]map[string]interface{}, 0, 1)
	telemetryConfigurations = append(telemetryConfigurations, map[string]interface{}{
		"enabled":         *logicalInterconnectGroup.TelemetryConfiguration.EnableTelemetry,
		"sample_count":    logicalInterconnectGroup.TelemetryConfiguration.SampleCount,
		"sample_interval": logicalInterconnectGroup.TelemetryConfiguration.SampleInterval,
		"type":            logicalInterconnectGroup.TelemetryConfiguration.Type,
	})
	d.Set("telemetry_configuration", telemetryConfigurations)

	trapDestinations := make([]map[string]interface{}, 0, 1)
	for _, trapDestination := range logicalInterconnectGroup.SnmpConfiguration.TrapDestinations {

		enetTrapCategories := make([]interface{}, len(trapDestination.EnetTrapCategories))
		for i, enetTrapCategory := range trapDestination.EnetTrapCategories {
			enetTrapCategories[i] = enetTrapCategory
		}

		fcTrapCategories := make([]interface{}, len(trapDestination.FcTrapCategories))
		for i, fcTrapCategory := range trapDestination.FcTrapCategories {
			fcTrapCategories[i] = fcTrapCategory
		}

		vcmTrapCategories := make([]interface{}, len(trapDestination.VcmTrapCategories))
		for i, vcmTrapCategory := range trapDestination.VcmTrapCategories {
			vcmTrapCategories[i] = vcmTrapCategory
		}

		trapSeverities := make([]interface{}, len(trapDestination.TrapSeverities))
		for i, trapSeverity := range trapDestination.TrapSeverities {
			trapSeverities[i] = trapSeverity
		}

		trapDestinations = append(trapDestinations, map[string]interface{}{
			"trap_destination":     trapDestination.TrapDestination,
			"community_string":     trapDestination.CommunityString,
			"trap_format":          trapDestination.TrapFormat,
			"enet_trap_categories": schema.NewSet(schema.HashString, enetTrapCategories),
			"fc_trap_categories":   schema.NewSet(schema.HashString, fcTrapCategories),
			"vcm_trap_categories":  schema.NewSet(schema.HashString, vcmTrapCategories),
			"trap_severities":      schema.NewSet(schema.HashString, trapSeverities),
		})
	}

	//Oneview returns an unordered list so order it to match the configuration file
	trapDestinationCount := d.Get("snmp_configuration.0.trap_destination.#").(int)
	oneviewTrapDestinationCount := len(trapDestinations)
	for i := 0; i < trapDestinationCount; i++ {
		currDest := d.Get("snmp_configuration.0.trap_destination." + strconv.Itoa(i) + ".trap_destination").(string)
		for j := 0; j < oneviewTrapDestinationCount; j++ {
			if currDest == trapDestinations[j]["trap_destination"] && i <= j {
				trapDestinations[i], trapDestinations[j] = trapDestinations[j], trapDestinations[i]
			}
		}
	}

	snmpAccess := make([]interface{}, len(logicalInterconnectGroup.SnmpConfiguration.SnmpAccess))
	for i, snmpAccessIP := range logicalInterconnectGroup.SnmpConfiguration.SnmpAccess {
		snmpAccess[i] = snmpAccessIP
	}

	snmpConfiguration := make([]map[string]interface{}, 0, 1)
	snmpConfiguration = append(snmpConfiguration, map[string]interface{}{
		"enabled":          *logicalInterconnectGroup.SnmpConfiguration.Enabled,
		"v3_enabled":       *logicalInterconnectGroup.SnmpConfiguration.V3Enabled,
		"read_community":   logicalInterconnectGroup.SnmpConfiguration.ReadCommunity,
		"snmp_access":      schema.NewSet(schema.HashString, snmpAccess),
		"system_contact":   logicalInterconnectGroup.SnmpConfiguration.SystemContact,
		"type":             logicalInterconnectGroup.SnmpConfiguration.Type,
		"trap_destination": trapDestinations,
	})
	d.Set("snmp_configuration", snmpConfiguration)

	interconnectSettings := make([]map[string]interface{}, 0, 1)
	interconnectSetting := map[string]interface{}{
		"type":                    logicalInterconnectGroup.EthernetSettings.Type,
		"fast_mac_cache_failover": *logicalInterconnectGroup.EthernetSettings.EnableFastMacCacheFailover,
		"network_loop_protection": *logicalInterconnectGroup.EthernetSettings.EnableNetworkLoopProtection,
		"pause_flood_protection":  *logicalInterconnectGroup.EthernetSettings.EnablePauseFloodProtection,
		"rich_tlv":                *logicalInterconnectGroup.EthernetSettings.EnableRichTLV,
		"mac_refresh_interval":    logicalInterconnectGroup.EthernetSettings.MacRefreshInterval,
	}

	interconnectSetting["interconnect_utilization_alert"] = *logicalInterconnectGroup.EthernetSettings.EnableInterconnectUtilizationAlert
	interconnectSettings = append(interconnectSettings, interconnectSetting)
	d.Set("interconnect_settings", interconnectSettings)

	if logicalInterconnectGroup.IgmpSettings != nil {
		igmpSettings := make([]map[string]interface{}, 0, 1)
		igmpSetting := map[string]interface{}{
			"category":                   logicalInterconnectGroup.IgmpSettings.Category,
			"consistency_checking":       logicalInterconnectGroup.IgmpSettings.ConsistencyChecking,
			"created":                    logicalInterconnectGroup.IgmpSettings.Created,
			"dependent_resource_uri":     logicalInterconnectGroup.IgmpSettings.DependentResourceUri,
			"description":                logicalInterconnectGroup.IgmpSettings.Description,
			"etag":                       logicalInterconnectGroup.IgmpSettings.ETAG,
			"igmp_snooping":              *logicalInterconnectGroup.IgmpSettings.EnableIgmpSnooping,
			"prevent_flooding":           *logicalInterconnectGroup.IgmpSettings.EnablePreventFlooding,
			"proxy_reporting":            *logicalInterconnectGroup.IgmpSettings.EnableProxyReporting,
			"id":                         logicalInterconnectGroup.IgmpSettings.ID,
			"igmp_idle_timeout_interval": logicalInterconnectGroup.IgmpSettings.IgmpIdleTimeoutInterval,
			"igmp_snooping_vlan_ids":     logicalInterconnectGroup.IgmpSettings.IgmpSnoopingVlanIds,
			"modified":                   logicalInterconnectGroup.IgmpSettings.Modified,
			"name":                       logicalInterconnectGroup.IgmpSettings.Name,
			"state":                      logicalInterconnectGroup.IgmpSettings.State,
			"status":                     logicalInterconnectGroup.IgmpSettings.Status,
			"type":                       logicalInterconnectGroup.IgmpSettings.Type,
			"uri":                        logicalInterconnectGroup.IgmpSettings.URI,
		}
		igmpSettings = append(igmpSettings, igmpSetting)
		d.Set("igmp_settings", igmpSettings)
	}

	if logicalInterconnectGroup.PortFlapProtection != nil {
		portFlapSettings := make([]map[string]interface{}, 0, 1)
		portFlapSetting := map[string]interface{}{
			"type":                             logicalInterconnectGroup.PortFlapProtection.Type,
			"uri":                              logicalInterconnectGroup.PortFlapProtection.URI,
			"category":                         logicalInterconnectGroup.PortFlapProtection.Category,
			"etag":                             logicalInterconnectGroup.PortFlapProtection.ETAG,
			"created":                          logicalInterconnectGroup.PortFlapProtection.Created,
			"modified":                         logicalInterconnectGroup.PortFlapProtection.Modified,
			"id":                               logicalInterconnectGroup.PortFlapProtection.ID,
			"name":                             logicalInterconnectGroup.PortFlapProtection.Name,
			"detection_interval":               logicalInterconnectGroup.PortFlapProtection.DetectionInterval,
			"port_flap_threshold_per_interval": logicalInterconnectGroup.PortFlapProtection.PortFlapThresholdPerInterval,
			"no_of_samples_declare_failures":   logicalInterconnectGroup.PortFlapProtection.NoOfSamplesDeclareFailures,
			"consistency_checking":             logicalInterconnectGroup.PortFlapProtection.ConsistencyChecking,
			"port_flap_protection_mode":        logicalInterconnectGroup.PortFlapProtection.PortFlapProtectionMode,
			"description":                      logicalInterconnectGroup.PortFlapProtection.Description,
			"state":                            logicalInterconnectGroup.PortFlapProtection.State,
			"status":                           logicalInterconnectGroup.PortFlapProtection.Status,
		}
		portFlapSettings = append(portFlapSettings, portFlapSetting)
		d.Set("port_flap_settings", portFlapSettings)
	}
	return nil
}

func resourceLogicalInterconnectGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteLogicalInterconnectGroup(d.Get("name").(string))
	if err != nil {
		return err
	}
	return nil
}

func GetIntPointer(value int) *int {
	return &value
}

func GetBoolPointer(value bool) *bool {
	return &value
}

func resourceLogicalInterconnectGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	lig := ov.LogicalInterconnectGroup{
		Name: d.Get("name").(string),
		Type: d.Get("type").(string),
		URI:  utils.NewNstring(d.Get("uri").(string)),
	}

	if val, ok := d.GetOk("interconnect_bay_set"); ok {
		lig.InterconnectBaySet = val.(int)
	}

	if val, ok := d.GetOk("redundancy_type"); ok {
		lig.RedundancyType = val.(string)
	}

	if val, ok := d.GetOk("enclosure_indexes"); ok {
		rawEnclosureIndexes := val.(*schema.Set).List()
		enclosureIndexes := make([]int, len(rawEnclosureIndexes))
		for i, raw := range rawEnclosureIndexes {
			enclosureIndexes[i] = raw.(int)
		}
		lig.EnclosureIndexes = enclosureIndexes
	}

	if val, ok := d.GetOk("initial_scope_uris"); ok {
		rawInitialScopeUris := val.(*schema.Set).List()
		initialScopeUris := make([]utils.Nstring, len(rawInitialScopeUris))
		for _, raw := range rawInitialScopeUris {
			initialScopeUri := utils.Nstring(raw.(string))
			initialScopeUris = append(initialScopeUris, initialScopeUri)
		}
		lig.InitialScopeUris = initialScopeUris
	}

	interconnectMapEntryTemplates := make([]ov.InterconnectMapEntryTemplate, 0)
	rawInterconnectMapEntryTemplates := d.Get("interconnect_map_entry_template").(*schema.Set).List()
	for _, raw := range rawInterconnectMapEntryTemplates {
		interconnectMapEntryTemplate := raw.(map[string]interface{})
		interconnectTypeName := interconnectMapEntryTemplate["interconnect_type_name"].(string)
		interconnectType, err := config.ovClient.GetInterconnectTypeByName(interconnectTypeName)
		if err != nil {
			return err
		}
		if interconnectType.URI == "" {
			return fmt.Errorf("Could not find Interconnect Type from name: %s", interconnectTypeName)
		}

		enclosureLocation := ov.LocationEntry{
			RelativeValue: interconnectMapEntryTemplate["enclosure_index"].(int),
			Type:          "Enclosure",
		}
		locationEntries := make([]ov.LocationEntry, 0)
		locationEntries = append(locationEntries, enclosureLocation)

		bayLocation := ov.LocationEntry{
			RelativeValue: interconnectMapEntryTemplate["bay_number"].(int),
			Type:          "Bay",
		}
		locationEntries = append(locationEntries, bayLocation)
		logicalLocation := ov.LogicalLocation{
			LocationEntries: locationEntries,
		}
		interconnectMapEntryTemplates = append(interconnectMapEntryTemplates, ov.InterconnectMapEntryTemplate{
			LogicalLocation:              logicalLocation,
			EnclosureIndex:               interconnectMapEntryTemplate["enclosure_index"].(int),
			PermittedInterconnectTypeUri: interconnectType.URI,
		})
	}

	interconnectMapTemplate := ov.InterconnectMapTemplate{
		InterconnectMapEntryTemplates: interconnectMapEntryTemplates,
	}
	lig.InterconnectMapTemplate = &interconnectMapTemplate

	uplinkSetCount := d.Get("uplink_set.#").(int)
	uplinkSets := make([]ov.UplinkSets, 0)
	for i := 0; i < uplinkSetCount; i++ {
		uplinkSetPrefix := fmt.Sprintf("uplink_set.%d", i)
		uplinkSet := ov.UplinkSets{}
		if val, ok := d.GetOk(uplinkSetPrefix + ".name"); ok {
			uplinkSet.Name = val.(string)
		}
		if val, ok := d.GetOk(uplinkSetPrefix + ".network_type"); ok {
			uplinkSet.NetworkType = val.(string)
		}
		if val, ok := d.GetOk(uplinkSetPrefix + ".ethernet_network_type"); ok {
			uplinkSet.EthernetNetworkType = val.(string)
		}
		if val, ok := d.GetOk(uplinkSetPrefix + ".mode"); ok {
			uplinkSet.Mode = val.(string)
		}
		if val, ok := d.GetOk(uplinkSetPrefix + ".lacp_timer"); ok {
			uplinkSet.LacpTimer = val.(string)
		}
		if val, ok := d.GetOk(uplinkSetPrefix + ".native_network_uri"); ok {
			uplinkSet.NativeNetworkUri = utils.NewNstring(val.(string))
		}

		logicalPortCount := d.Get(uplinkSetPrefix + ".logical_port_config.#").(int)
		logicalPorts := make([]ov.LogicalPortConfigInfo, 0)
		for i := 0; i < logicalPortCount; i++ {
			logicalPortPrefix := fmt.Sprintf(uplinkSetPrefix+".logical_port_config.%d", i)
			rawPortLocations := d.Get(logicalPortPrefix + ".port_num").(*schema.Set).List()
			for _, raw := range rawPortLocations {
				logicalPort := ov.LogicalPortConfigInfo{}

				if val, ok := d.GetOk(logicalPortPrefix + ".desired_speed"); ok {
					logicalPort.DesiredSpeed = val.(string)
				}

				locationEntries := make([]ov.LocationEntry, 0)
				enclosureLocation := ov.LocationEntry{
					RelativeValue: d.Get(logicalPortPrefix + ".enclosure_num").(int),
					Type:          "Enclosure",
				}
				locationEntries = append(locationEntries, enclosureLocation)

				bayLocation := ov.LocationEntry{
					RelativeValue: d.Get(logicalPortPrefix + ".bay_num").(int),
					Type:          "Bay",
				}
				locationEntries = append(locationEntries, bayLocation)

				portLocation := ov.LocationEntry{
					RelativeValue: raw.(int),
					Type:          "Port",
				}
				locationEntries = append(locationEntries, portLocation)

				logicalLocation := ov.LogicalLocation{
					LocationEntries: locationEntries,
				}

				logicalPort.LogicalLocation = logicalLocation
				if _, ok := d.GetOk(logicalPortPrefix + ".primary_port"); ok {
					if uplinkSet.PrimaryPort == nil {
						uplinkSet.PrimaryPort = &logicalLocation
					}
				}

				logicalPorts = append(logicalPorts, logicalPort)
			}

		}
		uplinkSet.LogicalPortConfigInfos = logicalPorts

		rawNetUris := d.Get(uplinkSetPrefix + ".network_uris").(*schema.Set).List()
		netUris := make([]utils.Nstring, 0)
		for _, raw := range rawNetUris {
			netUris = append(netUris, utils.NewNstring(raw.(string)))
		}
		uplinkSet.NetworkUris = netUris

		uplinkSets = append(uplinkSets, uplinkSet)
	}
	lig.UplinkSets = uplinkSets

	rawInternalNetUris := d.Get("internal_network_uris").(*schema.Set).List()
	internalNetUris := make([]utils.Nstring, len(rawInternalNetUris))
	for i, raw := range rawInternalNetUris {
		internalNetUris[i] = utils.NewNstring(raw.(string))
	}
	lig.InternalNetworkUris = internalNetUris

	telemetryConfigPrefix := fmt.Sprintf("telemetry_configuration.0")
	telemetryConfiguration := ov.TelemetryConfiguration{}
	if val, ok := d.GetOk(telemetryConfigPrefix + ".sample_count"); ok {
		telemetryConfiguration.SampleCount = val.(int)
	}
	if val, ok := d.GetOk(telemetryConfigPrefix + ".sample_interval"); ok {
		telemetryConfiguration.SampleInterval = val.(int)
	}
	if val, ok := d.GetOk(telemetryConfigPrefix + ".enabled"); ok {
		enabled := val.(bool)
		telemetryConfiguration.EnableTelemetry = &enabled
	}
	if telemetryConfiguration != (ov.TelemetryConfiguration{}) {
		telemetryConfiguration.Type = d.Get(telemetryConfigPrefix + ".type").(string)
		lig.TelemetryConfiguration = &telemetryConfiguration
	}

	sflowConfigurationPrefix := fmt.Sprintf("sflow_configuration.0")
	sflowConfiguration := ov.SflowConfiguration{}

	sflowAgentCount := d.Get(sflowConfigurationPrefix + ".sflow_agents.#").(int)
	sflowAgents := make([]ov.SflowAgent, 0)
	for i := 0; i < sflowAgentCount; i++ {
		sflowAgentPrefix := fmt.Sprintf(sflowConfigurationPrefix+".sflow_agents.%d", i)
		sflowAgent := ov.SflowAgent{}

		if val, ok := d.GetOk(sflowAgentPrefix + ".bay_number"); ok {
			sflowAgent.BayNumber = val.(int)
		}

		if val, ok := d.GetOk(sflowAgentPrefix + ".enclosure_index"); ok {
			sflowAgent.EnclosureIndex = val.(int)
		}

		if val, ok := d.GetOk(sflowAgentPrefix + ".ip_addr"); ok {
			sflowAgent.IpAddr = val.(string)
		}

		if val, ok := d.GetOk(sflowAgentPrefix + ".ip_mode"); ok {
			sflowAgent.IpMode = val.(string)
		}

		if val, ok := d.GetOk(sflowAgentPrefix + ".status"); ok {
			sflowAgent.Status = val.(string)
		}

		if val, ok := d.GetOk(sflowAgentPrefix + ".subnet_mask"); ok {
			sflowAgent.SubnetMask = val.(string)
		}

		sflowAgents = append(sflowAgents, sflowAgent)
	}
	sflowConfiguration.SflowAgents = sflowAgents

	sflowCollectorCount := d.Get(sflowConfigurationPrefix + ".sflow_collectors.#").(int)
	sflowCollectors := make([]ov.SflowCollector, 0)
	for i := 0; i < sflowCollectorCount; i++ {
		sflowCollectorPrefix := fmt.Sprintf(sflowConfigurationPrefix+".sflow_collectors.%d", i)
		sflowCollector := ov.SflowCollector{}

		if val, ok := d.GetOk(sflowCollectorPrefix + ".collector_enabled"); ok {
			enabled := val.(bool)
			sflowCollector.CollectorEnabled = &enabled
		}

		if val, ok := d.GetOk(sflowCollectorPrefix + ".collector_id"); ok {
			sflowCollector.CollectorId = val.(int)
		}

		if val, ok := d.GetOk(sflowCollectorPrefix + ".ip_address"); ok {
			sflowCollector.IPAddress = val.(string)
		}

		if val, ok := d.GetOk(sflowCollectorPrefix + ".max_datagram_size"); ok {
			sflowCollector.MaxDatagramSize = val.(int)
		}

		if val, ok := d.GetOk(sflowCollectorPrefix + ".max_header_size"); ok {
			sflowCollector.MaxHeaderSize = val.(int)
		}

		if val, ok := d.GetOk(sflowCollectorPrefix + ".name"); ok {
			sflowCollector.Name = val.(string)
		}

		if val, ok := d.GetOk(sflowCollectorPrefix + ".port"); ok {
			sflowCollector.Port = val.(int)
		}

		sflowCollectors = append(sflowCollectors, sflowCollector)
	}
	sflowConfiguration.SflowCollectors = sflowCollectors

	sflowNetworkPrefix := fmt.Sprintf(sflowConfigurationPrefix + ".sflow_network.0")
	sflowNetwork := ov.SflowNetwork{}

	if val, ok := d.GetOk(sflowNetworkPrefix + ".vlan_id"); ok {
		sflowNetwork.VlanId = val.(int)
	}

	if val, ok := d.GetOk(sflowNetworkPrefix + ".uri"); ok {
		sflowNetwork.URI = utils.Nstring(val.(string))
	}

	if val, ok := d.GetOk(sflowNetworkPrefix + ".name"); ok {
		sflowNetwork.Name = val.(string)
	}

	sflowConfiguration.SflowNetwork = &sflowNetwork

	sflowPortCount := d.Get(sflowConfigurationPrefix + ".sflow_ports.#").(int)
	sflowPorts := make([]ov.SflowPort, 0)
	for i := 0; i < sflowPortCount; i++ {
		sflowPortPrefix := fmt.Sprintf(sflowConfigurationPrefix+".sflow_ports.%d", i)
		sflowPort := ov.SflowPort{}

		if val, ok := d.GetOk(sflowPortPrefix + ".bay_number"); ok {
			sflowPort.BayNumber = val.(int)
		}

		if val, ok := d.GetOk(sflowPortPrefix + ".enclosure_index"); ok {
			sflowPort.EnclosureIndex = val.(int)
		}

		if val, ok := d.GetOk(sflowPortPrefix + ".collector_id"); ok {
			sflowPort.CollectorId = val.(int)
		}

		if val, ok := d.GetOk(sflowPortPrefix + ".icm_name"); ok {
			sflowPort.IcmName = val.(string)
		}

		if val, ok := d.GetOk(sflowPortPrefix + ".port_num"); ok {
			sflowPort.PortName = val.(string)
		}

		sflowPorts = append(sflowPorts, sflowPort)
	}
	sflowConfiguration.SflowPorts = sflowPorts

	if val, ok := d.GetOk(sflowConfigurationPrefix + ".category"); ok {
		sflowConfiguration.Category = val.(string)
	}
	if val, ok := d.GetOk(sflowConfigurationPrefix + ".description"); ok {
		sflowConfiguration.Description = utils.NewNstring(val.(string))
	}
	if val, ok := d.GetOk(sflowConfigurationPrefix + ".enabled"); ok {
		enabled := val.(bool)
		sflowConfiguration.Enabled = &enabled
	}
	if val, ok := d.GetOk(sflowConfigurationPrefix + ".name"); ok {
		sflowConfiguration.Name = val.(string)
	}
	if val, ok := d.GetOk(sflowConfigurationPrefix + ".state"); ok {
		sflowConfiguration.State = val.(string)
	}
	if val, ok := d.GetOk(sflowConfigurationPrefix + ".status"); ok {
		sflowConfiguration.Status = val.(string)
	}
	if val, ok := d.GetOk(sflowConfigurationPrefix + ".uri"); ok {
		sflowConfiguration.URI = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk(sflowConfigurationPrefix + ".type"); ok {
		sflowConfiguration.Type = val.(string)
		lig.SflowConfiguration = &sflowConfiguration
	}

	snmpConfigPrefix := fmt.Sprintf("snmp_configuration.0")
	snmpConfiguration := ov.SnmpConfiguration{}
	if val, ok := d.GetOk(snmpConfigPrefix + ".enabled"); ok {
		enabled := val.(bool)
		snmpConfiguration.Enabled = &enabled
	}
	if val, ok := d.GetOk(snmpConfigPrefix + ".v3_enabled"); ok {
		v3Enabled := val.(bool)
		snmpConfiguration.V3Enabled = &v3Enabled
	}
	if val, ok := d.GetOk(snmpConfigPrefix + ".read_community"); ok {
		snmpConfiguration.ReadCommunity = val.(string)
	}
	if val, ok := d.GetOk(snmpConfigPrefix + ".system_contact"); ok {
		snmpConfiguration.SystemContact = val.(string)
	}
	rawSnmpAccess := d.Get(snmpConfigPrefix + ".snmp_access").(*schema.Set).List()
	snmpAccess := make([]string, len(rawSnmpAccess))
	for i, raw := range rawSnmpAccess {
		snmpAccess[i] = raw.(string)
	}

	trapDestinationCount := d.Get(snmpConfigPrefix + ".trap_destination.#").(int)
	trapDestinations := make([]ov.TrapDestination, 0, trapDestinationCount)
	for i := 0; i < trapDestinationCount; i++ {
		trapDestinationPrefix := fmt.Sprintf(snmpConfigPrefix+".trap_destination.%d", i)

		rawEnetTrapCategories := d.Get(trapDestinationPrefix + ".enet_trap_categories").(*schema.Set).List()
		enetTrapCategories := make([]string, len(rawEnetTrapCategories))
		for i, raw := range rawEnetTrapCategories {
			enetTrapCategories[i] = raw.(string)
		}

		rawFcTrapCategories := d.Get(trapDestinationPrefix + ".fc_trap_categories").(*schema.Set).List()
		fcTrapCategories := make([]string, len(rawFcTrapCategories))
		for i, raw := range rawFcTrapCategories {
			fcTrapCategories[i] = raw.(string)
		}

		rawVcmTrapCategories := d.Get(trapDestinationPrefix + ".vcm_trap_categories").(*schema.Set).List()
		vcmTrapCategories := make([]string, len(rawVcmTrapCategories))
		for i, raw := range rawVcmTrapCategories {
			vcmTrapCategories[i] = raw.(string)
		}

		rawTrapSeverities := d.Get(trapDestinationPrefix + ".trap_severities").(*schema.Set).List()
		trapSeverities := make([]string, len(rawTrapSeverities))
		for i, raw := range rawTrapSeverities {
			trapSeverities[i] = raw.(string)
		}

		trapDestination := ov.TrapDestination{
			TrapDestination:    d.Get(trapDestinationPrefix + ".trap_destination").(string),
			CommunityString:    d.Get(trapDestinationPrefix + ".community_string").(string),
			TrapFormat:         d.Get(trapDestinationPrefix + ".trap_format").(string),
			EnetTrapCategories: enetTrapCategories,
			FcTrapCategories:   fcTrapCategories,
			VcmTrapCategories:  vcmTrapCategories,
			TrapSeverities:     trapSeverities,
		}
		trapDestinations = append(trapDestinations, trapDestination)
	}
	if trapDestinationCount > 0 {
		snmpConfiguration.TrapDestinations = trapDestinations
	}

	snmpConfiguration.SnmpAccess = snmpAccess
	if val, ok := d.GetOk(snmpConfigPrefix + ".type"); ok {
		snmpConfiguration.Type = val.(string)
		lig.SnmpConfiguration = &snmpConfiguration
	}

	ligCall, _ := config.ovClient.GetLogicalInterconnectGroupByName(d.Get("name").(string))

	interconnectSettingsPrefix := fmt.Sprintf("interconnect_settings.0")
	if val, ok := d.GetOk(interconnectSettingsPrefix + ".type"); ok {
		interconnectSettings := ov.EthernetSettings{}

		macFailoverEnabled := d.Get(interconnectSettingsPrefix + ".fast_mac_cache_failover").(bool)
		interconnectSettings.EnableFastMacCacheFailover = &macFailoverEnabled

		networkLoopProtectionEnabled := d.Get(interconnectSettingsPrefix + ".network_loop_protection").(bool)
		interconnectSettings.EnableNetworkLoopProtection = &networkLoopProtectionEnabled

		pauseFloodProtectionEnabled := d.Get(interconnectSettingsPrefix + ".pause_flood_protection").(bool)
		interconnectSettings.EnablePauseFloodProtection = &pauseFloodProtectionEnabled

		if val1, ok := d.GetOk(interconnectSettingsPrefix + ".rich_tlv"); ok {
			enabled := val1.(bool)
			interconnectSettings.EnableRichTLV = &enabled
		}

		if val1, ok := d.GetOk(interconnectSettingsPrefix + ".interconnect_utilization_alert"); ok {
			enabled := val1.(bool)
			interconnectSettings.EnableInterconnectUtilizationAlert = &enabled
		}

		if val1, ok := d.GetOk(interconnectSettingsPrefix + ".mac_refresh_interval"); ok {
			interconnectSettings.MacRefreshInterval = val1.(int)
		}
		interconnectSettings.DependentResourceUri = ligCall.EthernetSettings.DependentResourceUri
		interconnectSettings.Type = val.(string)
		lig.EthernetSettings = &interconnectSettings
	}

	rawigmpsetting := d.Get("igmp_settings").([]interface{})
	igmpSetting := ov.IgmpSettings{}
	for _, val := range rawigmpsetting {

		rawlval := val.(map[string]interface{})

		enableigmpsnooping := rawlval["igmp_snooping"].(bool)
		enablepreventflooding := rawlval["prevent_flooding"].(bool)
		enableproxyreporting := rawlval["proxy_reporting"].(bool)

		igmpSetting.Created = rawlval["created"].(string)
		igmpSetting.Category = utils.Nstring(rawlval["category"].(string))
		igmpSetting.Type = rawlval["type"].(string)
		igmpSetting.ConsistencyChecking = rawlval["consistency_checking"].(string)
		igmpSetting.DependentResourceUri = ligCall.IgmpSettings.DependentResourceUri
		igmpSetting.Description = rawlval["description"].(string)
		igmpSetting.ETAG = utils.Nstring(rawlval["etag"].(string))
		igmpSetting.EnableIgmpSnooping = &enableigmpsnooping
		igmpSetting.EnablePreventFlooding = &enablepreventflooding
		igmpSetting.EnableProxyReporting = &enableproxyreporting
		igmpSetting.ID = rawlval["id"].(string)
		igmpSetting.IgmpIdleTimeoutInterval = rawlval["igmp_idle_timeout_interval"].(int)
		igmpSetting.IgmpSnoopingVlanIds = rawlval["igmp_snooping_vlan_ids"].(string)
		igmpSetting.Modified = rawlval["modified"].(string)
		igmpSetting.Name = rawlval["name"].(string)
		igmpSetting.State = rawlval["state"].(string)
		igmpSetting.Status = rawlval["status"].(string)
		igmpSetting.URI = utils.Nstring(rawlval["uri"].(string))
	}
	lig.IgmpSettings = &igmpSetting

	rawPortFlapSetting := d.Get("port_flap_settings").([]interface{})
	PortFlapSetting := ov.PortFlapProtection{}
	for _, val := range rawPortFlapSetting {

		rawlval := val.(map[string]interface{})
		PortFlapSetting.Type = rawlval["type"].(string)
		PortFlapSetting.URI = utils.Nstring(rawlval["uri"].(string))
		PortFlapSetting.Category = utils.Nstring(rawlval["category"].(string))
		PortFlapSetting.ETAG = rawlval["etag"].(string)
		PortFlapSetting.Created = rawlval["created"].(string)
		PortFlapSetting.Modified = rawlval["modified"].(string)
		PortFlapSetting.ID = rawlval["id"].(string)
		PortFlapSetting.Name = rawlval["name"].(string)
		PortFlapSetting.DetectionInterval = rawlval["detection_interval"].(int)
		PortFlapSetting.PortFlapThresholdPerInterval = rawlval["port_flap_threshold_per_interval"].(int)
		PortFlapSetting.NoOfSamplesDeclareFailures = rawlval["no_of_samples_declare_failures"].(int)
		PortFlapSetting.ConsistencyChecking = rawlval["consistency_checking"].(string)
		PortFlapSetting.PortFlapProtectionMode = rawlval["port_flap_protection_mode"].(string)
		PortFlapSetting.Description = utils.Nstring(rawlval["description"].(string))
		PortFlapSetting.State = rawlval["state"].(string)
		PortFlapSetting.Status = rawlval["status"].(string)
	}
	lig.PortFlapProtection = &PortFlapSetting
	if val, ok := d.GetOk("quality_of_service"); ok {
		rawQoss := val.([]interface{})
		ovQos := ov.QosConfiguration{}
		if len(rawQoss) != 0 {
			for _, rawQosConfig := range rawQoss {
				rawQos := rawQosConfig.(map[string]interface{})
				rawActiveQosConfigs := rawQos["active_qos_config"].([]interface{})
				ovActiveQosConfig := ov.ActiveQosConfig{}
				if len(rawActiveQosConfigs) != 0 {
					for _, rawActiveQosConfig := range rawActiveQosConfigs {
						activeQosConfig := rawActiveQosConfig.(map[string]interface{})
						rawQosClassifiers := activeQosConfig["qos_traffic_classifiers"].([]interface{})
						ovQosTrafficClassifier := make([]ov.QosTrafficClassifier, 0)
						if len(rawQosClassifiers) != 0 {
							for _, rawQosClassifier := range rawQosClassifiers {
								qosClassifier := rawQosClassifier.(map[string]interface{})
								ovQosClassificationMapping := ov.QosClassificationMapping{}
								rawQosClassificationMappings := qosClassifier["qos_classification_mapping"].([]interface{})
								if rawQosClassificationMappings != nil {
									for _, rawQosClassificationMapping := range rawQosClassificationMappings {
										qosClassificationMapping := rawQosClassificationMapping.(map[string]interface{})
										rawDot1pClassMappings := qosClassificationMapping["dot1p_class_mapping"].([]interface{})
										if qosClassificationMapping["dot1p_class_mapping"] != nil {
											dot1pClassMapping := make([]int, 0)
											for _, raw := range rawDot1pClassMappings {
												dot1pClassMapping = append(dot1pClassMapping, raw.(int))
											}
											ovQosClassificationMapping.Dot1pClassMapping = dot1pClassMapping
										}
										rawDscpClassMappings := qosClassificationMapping["dscp_class_mapping"].([]interface{})
										if qosClassificationMapping["dscp_class_mapping"] != nil {
											dscpClassMapping := make([]string, 0)
											for _, raw := range rawDscpClassMappings {
												dscpClassMapping = append(dscpClassMapping, raw.(string))
											}
											ovQosClassificationMapping.DscpClassMapping = dscpClassMapping
										}
									}
								}
								ovQosTrafficClass := ov.QosTrafficClass{}
								rawQosTrafficClasses := qosClassifier["qos_traffic_class"].([]interface{})
								if len(rawQosTrafficClasses) != 0 {
									for _, rawQosTrafficClass := range rawQosTrafficClasses {
										qosTrafficClass := rawQosTrafficClass.(map[string]interface{})
										rawDcbxConfigurations := qosTrafficClass["dcbx_configuration"].([]interface{})
										ovDcbxConfiguration := ov.DcbxConfigurations{}
										if len(rawDcbxConfigurations) != 0 {
											for _, rawDcbxConfiguration := range rawDcbxConfigurations {
												dcbxConfiguration := rawDcbxConfiguration.(map[string]interface{})
												rawDcbxEtsPorts := dcbxConfiguration["dcbx_ets_port"].([]interface{})
												ovDcbxEtsPorts := make([]ov.DcbxEtsPort, 0)
												if len(rawDcbxEtsPorts) != 0 {
													for _, rawDcbxPort := range rawDcbxEtsPorts {
														dcbxPort := rawDcbxPort.(map[string]interface{})
														ovDcbxEtsPorts = append(ovDcbxEtsPorts, ov.DcbxEtsPort{
															BayNumber:      dcbxPort["bay_number"].(string),
															EnclosureIndex: dcbxPort["enclosure_index"].(int),
															IcmName:        dcbxPort["icm_name"].(string),
															MaxBandwidth:   dcbxPort["max_bandwidth"].(string),
															MinBandwidth:   dcbxPort["min_bandwidth"].(string),
															PortName:       dcbxPort["port_name"].(string),
														})
													}
												}
												ovDcbxConfiguration = ov.DcbxConfigurations{
													ApplicationProtocol:     dcbxConfiguration["application_protocol"].(string),
													DefaultMaximumBandwidth: dcbxConfiguration["default_max_bandwidth"].(string),
													DefaultMinimumBandwidth: dcbxConfiguration["default_min_bandwidth"].(string),
													PriorityCodePoint:       dcbxConfiguration["priority_code_point"].(string),
													PriorityFlowControl:     dcbxConfiguration["priority_flow_control"].(string),
												}
												if len(ovDcbxEtsPorts) != 0 {
													ovDcbxConfiguration.DcbxEtsPorts = ovDcbxEtsPorts
												}
											}
										}
										ovQosTrafficClass = ov.QosTrafficClass{
											BandwidthShare:   qosTrafficClass["bandwidth_share"].(string),
											ClassName:        qosTrafficClass["class_name"].(string),
											EgressDot1pValue: GetIntPointer(qosTrafficClass["egress_dot1p_value"].(int)),
											Enabled:          GetBoolPointer(qosTrafficClass["enabled"].(bool)),
											MaxBandwidth:     qosTrafficClass["max_bandwidth"].(int),
											RealTime:         GetBoolPointer(qosTrafficClass["real_time"].(bool)),
											Roce:             GetBoolPointer(qosTrafficClass["roce"].(bool)),
										}
										if !(reflect.DeepEqual(ovDcbxConfiguration, ov.DcbxConfigurations{})) {
											ovQosTrafficClass.DcbxConfiguration = &ovDcbxConfiguration
										}
									}
								}
								if reflect.DeepEqual(ovQosClassificationMapping, ov.QosClassificationMapping{}) && ovQosTrafficClass == (ov.QosTrafficClass{}) {
									continue
								} else if reflect.DeepEqual(ovQosClassificationMapping, ov.QosClassificationMapping{}) {
									ovQosTrafficClassifier = append(ovQosTrafficClassifier, ov.QosTrafficClassifier{
										QosTrafficClass: &ovQosTrafficClass,
									})
								} else if ovQosTrafficClass == (ov.QosTrafficClass{}) {
									ovQosTrafficClassifier = append(ovQosTrafficClassifier, ov.QosTrafficClassifier{
										QosClassificationMapping: &ovQosClassificationMapping,
									})
								} else {
									ovQosTrafficClassifier = append(ovQosTrafficClassifier, ov.QosTrafficClassifier{
										QosClassificationMapping: &ovQosClassificationMapping,
										QosTrafficClass:          &ovQosTrafficClass,
									})
								}
							}
						}
						ovActiveQosConfig = ov.ActiveQosConfig{
							Category:                   utils.NewNstring(activeQosConfig["category"].(string)),
							ConfigType:                 activeQosConfig["config_type"].(string),
							Created:                    activeQosConfig["created"].(string),
							Description:                utils.NewNstring(activeQosConfig["description"].(string)),
							DownlinkClassificationType: activeQosConfig["downlink_classification_type"].(string),
							ETAG:                       activeQosConfig["etag"].(string),
							Modified:                   activeQosConfig["modified"].(string),
							Name:                       activeQosConfig["name"].(string),
							State:                      activeQosConfig["state"].(string),
							Status:                     activeQosConfig["status"].(string),
							Type:                       activeQosConfig["type"].(string),
							UplinkClassificationType:   activeQosConfig["uplink_classification_type"].(string),
							URI:                        utils.NewNstring(activeQosConfig["uri"].(string)),
							QosTrafficClassifiers:      ovQosTrafficClassifier,
						}
					}
				}
				ovQos = ov.QosConfiguration{
					Category:            rawQos["category"].(string),
					ConsistencyChecking: rawQos["consistency_checking"].(string),
					Created:             rawQos["created"].(string),
					Description:         utils.NewNstring(rawQos["description"].(string)),
					ETAG:                rawQos["etag"].(string),
					Modified:            rawQos["modified"].(string),
					Name:                rawQos["name"].(string),
					State:               rawQos["state"].(string),
					Status:              rawQos["status"].(string),
					Type:                rawQos["type"].(string),
					URI:                 utils.NewNstring(rawQos["uri"].(string)),
					ActiveQosConfig:     &ovActiveQosConfig,
				}
			}

		}
		lig.QosConfiguration = &ovQos
	}
	err := config.ovClient.UpdateLogicalInterconnectGroup(lig)
	if err != nil {
		return err
	}
	d.SetId(d.Get("name").(string))
	return resourceLogicalInterconnectGroupRead(d, meta)
}
