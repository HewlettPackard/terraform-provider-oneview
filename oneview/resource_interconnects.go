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

func resourceInterconnects() *schema.Resource {
	return &schema.Resource{
		Read: resourceInterconnectsRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"base_wwn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"category": {
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
			"device_reset_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"eTag": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"edge_virtual_bridging_available": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_cut_through": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_fast_mac_cache_failover": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_igmp_snooping": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_network_loop_protection": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_pause_flood_protection": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_rich_tlv": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_storm_control": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_tagged_lldp": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enclosure_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enclosure_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enclosure_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"firmware_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"icm_licenses": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"license": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"consumed_count": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"license_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"required_count": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"state": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"igmp_idle_timeout_interval": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"igmp_snooping_vlan_ids": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"initial_scope_uris": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      HashString,
			},
			"interconnect_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"interconnect_location": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"location_entries": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"value": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"interconnect_mac": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"interconnect_type_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_address_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip_address_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"lldp_ip_address_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lldp_ipv4_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lldp_ipv6_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"logical_interconnect_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_bandwidth": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mgmt_interface": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"migration_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"model": {
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
			"network_loop_protection_interval": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"part_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ports": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"associated_uplink_set_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"available": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"bay_number": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"capability": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      HashString,
						},
						"category": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"config_port_types": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      HashString,
						},
						"connector_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"created": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"dcbx_info": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dcbx_ap_reason": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"dcbx_pfc_reason": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"dcbx_pg_reason": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"dcbx_status": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"eTag": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"fc_port_properties": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fcf_mac": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"logins": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"logins_count": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"neighbor_interconnect_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"op_online": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"op_online_reason": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"principle_interconnect_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"principle_interconnect_name_list": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      HashString,
									},
									"trunk_master": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"wwnn": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"wwpn": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"interconnect_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"lag_id": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"lag_states": {
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
						"neighbor": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"link_label": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"link_uri": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"remote_chassis_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"remote_chassis_id_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"remote_mgmt_address": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"remote_mgmt_address_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"remote_port_description": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"remote_port_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"remote_port_id_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"remote_system_capabilities": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"remote_system_description": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"remote_system_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"remote_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"operational_speed": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"paired_port_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_health_status": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_monitor_config_info": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_running_capability_type": {
							Type:     schema.Type,
							Optional: true,
						},
						"port_split_mode": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_status": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_status_reason": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_type_extended": {
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
						"subports": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port_number": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"port_status": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"port_status_reason": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"vendor_specific_port_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlans": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"power_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"qos_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
									"eTag": {
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
																Type:     schema.TypeSet,
																Optional: true,
																Elem:     &schema.Schema{Type: schema.TypeInt},
																Set: func(a interface{}) int {
																	return a.(int)
																},
															},
															"dscp_class_mapping": {
																Type:     schema.TypeSet,
																Optional: true,
																Elem:     &schema.Schema{Type: schema.TypeString},
																Set:      HashString,
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
																Type:     schema.TypeString,
																Optional: true,
															},
															"real_time": {
																Type:     schema.TypeBool,
																Optional: true,
															},
														},
													},
												},
											},
										},
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
										Type:     schema.Type,
										Optional: true,
									},
									"uri": {
										Type:     schema.Type,
										Optional: true,
									},
								},
							},
						},
						"category": {
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
						"eTag": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"inactive_fcoe_qos_config": {
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
									"eTag": {
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
																Type:     schema.TypeSet,
																Optional: true,
																Elem:     &schema.Schema{Type: schema.TypeInt},
																Set: func(a interface{}) int {
																	return a.(int)
																},
															},
															"dscp_class_mapping": {
																Type:     schema.TypeSet,
																Optional: true,
																Elem:     &schema.Schema{Type: schema.TypeString},
																Set:      HashString,
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
																Type:     schema.TypeString,
																Optional: true,
															},
															"real_time": {
																Type:     schema.TypeBool,
																Optional: true,
															},
														},
													},
												},
											},
										},
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
										Type:     schema.Type,
										Optional: true,
									},
									"uri": {
										Type:     schema.Type,
										Optional: true,
									},
								},
							},
						},
						"inactive_non_fcoe_qos_config": {
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
									"eTag": {
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
																Type:     schema.TypeSet,
																Optional: true,
																Elem:     &schema.Schema{Type: schema.TypeInt},
																Set: func(a interface{}) int {
																	return a.(int)
																},
															},
															"dscp_class_mapping": {
																Type:     schema.TypeSet,
																Optional: true,
																Elem:     &schema.Schema{Type: schema.TypeString},
																Set:      HashString,
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
																Type:     schema.TypeString,
																Optional: true,
															},
															"real_time": {
																Type:     schema.TypeBool,
																Optional: true,
															},
														},
													},
												},
											},
										},
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
										Type:     schema.Type,
										Optional: true,
									},
									"uri": {
										Type:     schema.Type,
										Optional: true,
									},
								},
							},
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
						"uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"remote_support": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"remote_support_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"support_data_collection_state": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"support_data_collection_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"support_data_collections_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"support_settings": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"destination": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"support_current_state": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"support_state": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"roles": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      HashString,
			},
			"scopes_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"serial_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"snmp_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
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
						"eTag": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
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
						"read_community": {
							Type:     schema.TypeString,
							Optional: true,
						},
						// TODO
						"snmp_access": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      HashString,
						},
					/*	"snmp_users": {
							Type:     schema.Type,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.Type},
							Set:      "",
						},*/
						"state": {
							Type:     schema.Type,
							Optional: true,
						},
						"status": {
							Type:     schema.Type,
							Optional: true,
						},
						"system_contact": {
							Type:     schema.TypeString,
							Optional: true,
						},
						// TODO
						/*"trap_destinations": {
							Type:     schema.Type,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.Type},
							Set:      "",
						},*/
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"v3_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"spare_part_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stacking_domain_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"stacking_domain_role": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stacking_member_id": {
				Type:     schema.TypeInt,
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
			"storm_control_polling_interval": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"storm_control_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sub_port_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"uid_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"unsupported_capabilities": {
				Type:     schema.Type,
				Optional: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceInterconnectsRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	eNet, err := config.ovClient.GetEthernetNetworkByName(d.Id())
	if err != nil || eNet.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.Set("name", eNet.Name)
	d.Set("vlan_id", eNet.VlanId)
	d.Set("purpose", eNet.Purpose)
	d.Set("smart_link", eNet.SmartLink)
	d.Set("private_network", eNet.PrivateNetwork)
	d.Set("ethernet_network_type", eNet.EthernetNetworkType)
	d.Set("type", eNet.Type)
	d.Set("created", eNet.Created)
	d.Set("modified", eNet.Modified)
	d.Set("uri", eNet.URI.String())
	d.Set("connection_template_uri", eNet.ConnectionTemplateUri.String())
	d.Set("status", eNet.Status)
	d.Set("category", eNet.Category)
	d.Set("state", eNet.State)
	d.Set("fabric_uri", eNet.FabricUri.String())
	d.Set("eTag", eNet.ETAG)
	return nil
}
