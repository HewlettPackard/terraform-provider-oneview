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
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceServerProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerProfileCreate,
		Read:   resourceServerProfileRead,
		Update: resourceServerProfileUpdate,
		Delete: resourceServerProfileDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"boot": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manage_boot": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"boot_order": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
					},
				},
			},
			"boot_mode": {
				Optional: true,
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manage_mode": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"mode": {
							Type:     schema.TypeString,
							Required: true,
						},
						"pxe_boot_policy": {
							Type:     schema.TypeString,
							Required: true,
						},
						"secure_boot": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"bios_option": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manage_bios": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"consistency_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"reapply_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"overridden_settings": {
							Optional: true,
							Computed: true,
							Type:     schema.TypeList,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
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
			"connection_settings": {
				Optional: true,
				Computed: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manage_connections": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"reapply_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compliance_control": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"connections": {
							Optional: true,
							Computed: true,
							Type:     schema.TypeList,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"allocated_mbps": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"allocated_vfs": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"function_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"network_uri": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"port_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"requested_mbps": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"requested_vfs": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"wwnn": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"wwpn": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"wwpn_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"interconnect_port": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"interconnect_uri": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"isolated_trunk": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"lag_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"mac": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"mac_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"managed": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"maximum_mbps": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"network_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"private_vlan_port_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"boot": {
										Optional: true,
										Computed: true,
										Type:     schema.TypeList,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"priority": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"boot_vlan_id": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"ethernet_boot_type": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"boot_volume_source": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"boot_target": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"array_wwpn": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"lun": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"iscsi": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"chap_level": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"chap_name": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"chap_secret": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"initiator_name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"initiator_name_source": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"mutual_chap_name": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"mutual_chap_secret": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"boot_target_lun": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"boot_target_name": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"first_boot_target_ip": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"first_boot_target_port": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"second_boot_target_ip": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"second_boot_target_port": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"ipv4": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"gateway": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"ip_address": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"subnet_mask": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"ip_address_source": {
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
			"server_hardware_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enclosure_group": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"affinity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hide_unused_flex_nics": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"initial_scope_uris": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"serial_number_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"wwn_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"firmware": {
				Optional: true,
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"force_install_firmware": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"firmware_baseline_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"consistency_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"firmware_activation_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"firmware_schedule_date_time": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"reapply_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"manage_firmware": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"firmware_install_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"local_storage": {
				Optional: true,
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manage_local_storage": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"initialize": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"controller": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeList,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"import_configuration": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"device_slot": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"initialize": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"drive_write_cache": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"mode": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"predictive_spare_rebuild": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"logical_drives": {
										Optional: true,
										Computed: true,
										Type:     schema.TypeList,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"bootable": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"accelerator": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"drive_number": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"drive_technology": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"num_physical_drives": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"num_spare_drives": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"sas_logical_jbod_id": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"raid_level": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"reapply_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sas_logical_jbod": {
							Optional: true,
							Computed: true,
							Type:     schema.TypeList,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"description": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"device_slot": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"drive_max_size_gb": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"drive_min_size_gb": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"drive_technology": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"erase_data": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"id": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"num_physical_drive": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"persistent": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"sas_logical_jbod_uri": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"san_storage": {
				Optional: true,
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_os_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"manage_san_storage": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"san_system_credentials": {
							Optional: true,
							Type:     schema.TypeList,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"chap_level": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"chap_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"chap_secret": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"chap_source": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"mutual_chap_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"mutual_chap_secret": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"storage_system_uri": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			// schema for ov.SanStorage.VolumeAttachments
			"volume_attachments": {
				Optional: true,
				Computed: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"associated_template_attachment_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"boot_volume_priority": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"is_permanent": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"lun": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"lun_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"volume_storage_system_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"volume_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"storage_paths": {
							Optional: true,
							Computed: true,
							Type:     schema.TypeList,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"connection_id": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"is_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"network_uri": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_selector": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"targets": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip_address": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"tcp_port": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"volume": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"initial_scope_uris": {
										Optional: true,
										Type:     schema.TypeSet,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Set: schema.HashString,
									},
									"is_permanent": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"template_uri": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"properties": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"data_protection_level": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"data_transfer_limit": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"description": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"folder": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"iops_limit": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"is_deduplicated": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"is_encrypted": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"is_pinned": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"is_shareable": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"performance_policy": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"provisioning_type": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"size": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"volume_set": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"is_data_reduction_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"is_adaptive_optimization_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"is_compressed": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"snapshot_pool": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"storage_pool": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"template_version": {
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
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"template": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hw_filter": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"hardware_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_connection": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ilo_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hardware_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_mac": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_slot_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"power_state": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: func(v interface{}, k string) (warning []string, errors []error) {
					val := v.(string)
					if val != "on" && val != "off" {
						errors = append(errors, fmt.Errorf("%q must be 'on' or 'off'", k))
					}
					return
				},
			},
			"options": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"op": {
							Required: true,
							Type:     schema.TypeString,
						},
						"path": {
							Required: true,
							Type:     schema.TypeString,
						},
						"value": {
							Required: true,
							Type:     schema.TypeString,
						},
					},
				},
			},
			"update_type": {
				Type:     schema.TypeString,
				Default:  "put",
				Optional: true,
			},
			"os_deployment_settings": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deploy_method": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"deployment_mac": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"deployment_port_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"force_os_deployment": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
						},
						"os_custom_attributes": {
							Computed: true,
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"constraints": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"os_deployment_plan_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"os_deployment_plan_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"os_volume": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"reapply_state": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"associated_server": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enclosure_bay": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enclosure_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"in_progress": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"iscsi_initiator_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"iscsi_initiator_name_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_processor": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manage_mp": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"mp_settings": {
							Optional: true,
							Type:     schema.TypeList,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"args": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"setting_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"reapply_state": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"modified": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"refresh_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scopes_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_hardware_reapply_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_hardware_type_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_manager": {
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
			"task_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"template_compliance": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceServerProfileCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	serverProfile := ov.ServerProfile{}

	if val, ok := d.GetOk("template"); ok {

		serverProfileByTemplate, err := config.ovClient.GetProfileTemplateByName(val.(string))
		if err != nil || serverProfileByTemplate.URI.IsNil() {
			return err
		}

		serverProfile = serverProfileByTemplate
		serverProfile.ServerProfileTemplateURI = serverProfileByTemplate.URI
		serverProfile.ConnectionSettings = ov.ConnectionSettings{
			Connections: serverProfile.ConnectionSettings.Connections,
		}
	}

	serverProfile.Type = d.Get("type").(string)
	serverProfile.Name = d.Get("name").(string)

	var serverHardware ov.ServerHardware
	if val, ok := d.GetOk("hardware_name"); ok {
		var err error
		serverHardware, err = config.ovClient.GetServerHardwareByName(val.(string))
		if err != nil {
			return err
		}
		serverProfile.ServerHardwareURI = serverHardware.URI
	}

	if val, ok := d.GetOk("affinity"); ok {
		serverProfile.Affinity = val.(string)
	}

	if val, ok := d.GetOk("serial_number_type"); ok {
		serverProfile.SerialNumberType = val.(string)
	}

	if val, ok := d.GetOk("wwn_type"); ok {
		serverProfile.WWNType = val.(string)
	}

	if val, ok := d.GetOk("mac_type"); ok {
		serverProfile.MACType = val.(string)
	}

	if val, ok := d.GetOk("hide_unused_flex_nics"); ok {
		serverProfile.HideUnusedFlexNics = val.(bool)
	}

	if val, ok := d.GetOk("description"); ok {
		serverProfile.Description = val.(string)
	}

	if val, ok := d.GetOk("enclosure_bay"); ok {
		serverProfile.EnclosureBay = val.(int)
	}

	if val, ok := d.GetOk("enclosure_uri"); ok {
		serverProfile.EnclosureURI = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk("iscsi_initiator_name"); ok {
		serverProfile.IscsiInitiatorName = val.(string)
	}

	if val, ok := d.GetOk("iscsi_initiator_name_type"); ok {
		serverProfile.IscsiInitiatorNameType = val.(string)
	}

	if val, ok := d.GetOk("scopes_uri"); ok {
		serverProfile.ScopesUri = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk("server_hardware_type_uri"); ok {
		serverProfile.ServerHardwareTypeURI = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk("service_manager"); ok {
		serverProfile.ServiceManager = val.(string)
	}

	if val, ok := d.GetOk("template_compliance"); ok {
		serverProfile.TemplateCompliance = val.(string)
	}

	if val, ok := d.GetOk("enclosure_group"); ok {
		enclosureGroup, err := config.ovClient.GetEnclosureGroupByName(val.(string))
		if err != nil {
			return err
		}
		serverProfile.EnclosureGroupURI = enclosureGroup.URI
	}

	if val, ok := d.GetOk("server_hardware_type"); ok {
		serverHardwareType, err := config.ovClient.GetServerHardwareTypeByName(val.(string))
		if err != nil {
			return err
		}
		serverProfile.ServerHardwareTypeURI = serverHardwareType.URI
	}

	if val, ok := d.GetOk("connection_settings"); ok {
		connections := val.([]interface{})
		for _, rawConSettings := range connections {
			rawConSetting := rawConSettings.(map[string]interface{})
			rawNetwork := rawConSetting["connections"].([]interface{})
			networks := make([]ov.Connection, 0)
			for i := 0; i < len(rawNetwork); i++ {
				rawNetworkItem := rawNetwork[i].(map[string]interface{})
				bootOptions := ov.BootOption{}
				if rawNetworkItem["boot"] != nil {
					rawBoots := rawNetworkItem["boot"].([]interface{})
					for _, rawBoot := range rawBoots {
						bootItem := rawBoot.(map[string]interface{})
						bootTargets := []ov.BootTarget{}
						rawBootTargets := bootItem["boot_target"].([]interface{})
						if rawBootTargets != nil {
							for _, rawBootTarget := range rawBootTargets {
								bootTarget := rawBootTarget.(map[string]interface{})
								bootTargets = append(bootTargets, ov.BootTarget{
									LUN:       bootTarget["lun"].(string),
									ArrayWWPN: bootTarget["array_wwpn"].(string),
								})
							}
						}
						iscsi := ov.BootIscsi{}
						if bootItem["iscsi"] != nil {
							rawIscsis := bootItem["iscsi"].([]interface{})
							for _, rawIscsi := range rawIscsis {
								rawIscsiItem := rawIscsi.(map[string]interface{})
								iscsi = ov.BootIscsi{
									BootTargetLun:        rawIscsiItem["boot_target_lun"].(string),
									BootTargetName:       rawIscsiItem["boot_target_name"].(string),
									ChapName:             rawIscsiItem["chap_name"].(string),
									ChapSecret:           rawIscsiItem["chap_secret"].(string),
									InitiatorName:        rawIscsiItem["initiator_name"].(string),
									InitiatorNameSource:  rawIscsiItem["initiator_name_source"].(string),
									MutualChapName:       rawIscsiItem["mutual_chap_name"].(string),
									MutualChapSecret:     rawIscsiItem["mutual_chap_secret"].(string),
									Chaplevel:            rawIscsiItem["chap_level"].(string),
									FirstBootTargetIp:    rawIscsiItem["first_boot_target_ip"].(string),
									FirstBootTargetPort:  rawIscsiItem["first_boot_target_port"].(string),
									SecondBootTargetIp:   rawIscsiItem["second_boot_target_ip"].(string),
									SecondBootTargetPort: rawIscsiItem["second_boot_target_port"].(string),
								}
							}
						}

						bootOptionV3 := ov.BootOptionV3{
							BootVlanId: bootItem["boot_vlan_id"].(int),
						}
						bootOptions = ov.BootOption{
							Priority:         bootItem["priority"].(string),
							BootOptionV3:     bootOptionV3,
							EthernetBootType: bootItem["ethernet_boot_type"].(string),
							BootVolumeSource: bootItem["boot_volume_source"].(string),
							Iscsi:            &iscsi,
							Targets:          bootTargets,
						}
					}
				}

				ipv4 := ov.Ipv4Option{}
				if rawNetworkItem["ipv4"] != nil {
					rawIpv4s := rawNetworkItem["ipv4"].([]interface{})
					for _, rawIpv4 := range rawIpv4s {
						rawIpv4Item := rawIpv4.(map[string]interface{})
						ipv4 = ov.Ipv4Option{
							Gateway:         rawIpv4Item["gateway"].(string),
							IpAddressSource: rawIpv4Item["ip_address_source"].(string),
						}
					}
				}
				connectionV200 := ov.Connectionv200{
					RequestedVFs: rawNetworkItem["requested_vfs"].(string),
				}

				networks = append(networks, ov.Connection{
					ID:               rawNetworkItem["id"].(int),
					Name:             rawNetworkItem["name"].(string),
					FunctionType:     rawNetworkItem["function_type"].(string),
					InterconnectPort: rawNetworkItem["interconnect_port"].(int),
					IsolatedTrunk:    rawNetworkItem["isolated_trunk"].(bool),
					LagName:          rawNetworkItem["lag_name"].(string),
					MAC:              utils.NewNstring(rawNetworkItem["mac"].(string)),
					MacType:          rawNetworkItem["mac_type"].(string),
					WWPN:             utils.NewNstring(rawNetworkItem["wwpn"].(string)),
					WWPNType:         rawNetworkItem["wwpn_type"].(string),
					NetworkURI:       utils.NewNstring(rawNetworkItem["network_uri"].(string)),
					PortID:           rawNetworkItem["port_id"].(string),
					Connectionv200:   connectionV200,
					RequestedMbps:    rawNetworkItem["requested_mbps"].(string),
					Ipv4:             &ipv4,
				})
				if len(rawNetworkItem["boot"].([]interface{})) != 0 {
					networks[i].Boot = &bootOptions
				}

			}

			serverProfile.ConnectionSettings = ov.ConnectionSettings{
				Connections: networks,
			}
		}
	}

	if val, ok := d.GetOk("boot"); ok {
		boot := val.([]interface{})
		for _, rawBoots := range boot {
			rawBoot := rawBoots.(map[string]interface{})
			rawBootOrder := rawBoot["boot_order"].(*schema.Set).List()
			bootOrder := make([]string, len(rawBootOrder))
			for i, raw := range rawBootOrder {
				bootOrder[i] = raw.(string)
			}
			serverProfile.Boot = ov.BootManagement{
				ManageBoot: rawBoot["manage_boot"].(bool),
				Order:      bootOrder,
			}
		}
	}
	if _, ok := d.GetOk("boot_mode"); ok {
		rawBootMode := d.Get("boot_mode").(*schema.Set).List()[0].(map[string]interface{})
		manageMode := rawBootMode["manage_mode"].(bool)
		serverProfile.BootMode = ov.BootModeOption{
			ManageMode:    &manageMode,
			Mode:          rawBootMode["mode"].(string),
			PXEBootPolicy: utils.Nstring(rawBootMode["pxe_boot_policy"].(string)),
			SecureBoot:    rawBootMode["secure_boot"].(string),
		}
	}

	if val, ok := d.GetOk("bios_option"); ok {
		rawBiosOption := val.([]interface{})
		biosOption := ov.BiosOption{}
		for _, raw := range rawBiosOption {
			rawBiosItem := raw.(map[string]interface{})

			overriddenSettings := make([]ov.BiosSettings, 0)
			rawOverriddenSetting := rawBiosItem["overridden_settings"].([]interface{})

			for _, raw2 := range rawOverriddenSetting {
				rawOverriddenSettingItem := raw2.(map[string]interface{})
				overriddenSettings = append(overriddenSettings, ov.BiosSettings{
					ID:    rawOverriddenSettingItem["id"].(string),
					Value: rawOverriddenSettingItem["value"].(string),
				})
			}
			manageBios := rawBiosItem["manage_bios"].(bool)
			biosOption = ov.BiosOption{
				ManageBios:         &manageBios,
				OverriddenSettings: overriddenSettings,
			}
		}

		serverProfile.Bios = &biosOption

	}

	if val, ok := d.GetOk("initial_scope_uris"); ok {
		initialScopeUrisOrder := val.(*schema.Set).List()
		initialScopeUris := make([]utils.Nstring, len(initialScopeUrisOrder))
		for i, raw := range initialScopeUrisOrder {
			initialScopeUris[i] = utils.Nstring(raw.(string))
		}
		serverProfile.InitialScopeUris = initialScopeUris
	}

	// Get firmware details if provided
	if _, ok := d.GetOk("firmware"); ok {
		rawFirmware := d.Get("firmware").(*schema.Set).List()
		firmware := ov.FirmwareOption{}
		for _, raw := range rawFirmware {
			firmwareItem := raw.(map[string]interface{})
			firmware = ov.FirmwareOption{
				ForceInstallFirmware:     firmwareItem["force_install_firmware"].(bool),
				FirmwareBaselineUri:      utils.NewNstring(firmwareItem["firmware_baseline_uri"].(string)),
				ManageFirmware:           firmwareItem["manage_firmware"].(bool),
				FirmwareInstallType:      firmwareItem["firmware_install_type"].(string),
				FirmwareActivationType:   firmwareItem["firmware_activation_type"].(string),
				FirmwareScheduleDateTime: firmwareItem["firmware_schedule_date_time"].(string),
			}
		}
		serverProfile.Firmware = firmware
	}
	// Get local storage data if provided
	if _, ok := d.GetOk("local_storage"); ok {
		rawLocalStorage := d.Get("local_storage").([]interface{})
		localStorage := ov.LocalStorageOptions{}
		for _, raw := range rawLocalStorage {
			localStorageItem := raw.(map[string]interface{})
			// Gets Local Storage Controller body
			rawLocalStorageController := localStorageItem["controller"].([]interface{})
			localStorageEmbeddedController := make([]ov.LocalStorageEmbeddedController, 0)
			for _, raw2 := range rawLocalStorageController {
				controllerData := raw2.(map[string]interface{})
				// Gets Local Storage Controller's Logical Drives
				rawLogicalDrives := controllerData["logical_drives"].([]interface{})
				logicalDrives := make([]ov.LogicalDriveV3, 0)
				for _, rawLogicalDrive := range rawLogicalDrives {
					logicalDrivesItem := rawLogicalDrive.(map[string]interface{})
					boot := logicalDrivesItem["bootable"].(bool)
					logicalDrives = append(logicalDrives, ov.LogicalDriveV3{
						Bootable:          &boot,
						RaidLevel:         logicalDrivesItem["raid_level"].(string),
						Accelerator:       logicalDrivesItem["accelerator"].(string),
						DriveTechnology:   logicalDrivesItem["drive_technology"].(string),
						Name:              logicalDrivesItem["name"].(string),
						NumPhysicalDrives: logicalDrivesItem["num_physical_drives"].(int),
						NumSpareDrives:    logicalDrivesItem["num_spare_drives"].(int),
						SasLogicalJBODId:  logicalDrivesItem["sas_logical_jbod_id"].(int),
					})
				}
				init := controllerData["initialize"].(bool)
				localStorageEmbeddedController = append(localStorageEmbeddedController, ov.LocalStorageEmbeddedController{
					DeviceSlot:             controllerData["device_slot"].(string),
					DriveWriteCache:        controllerData["drive_write_cache"].(string),
					Initialize:             &init,
					ImportConfiguration:    controllerData["import_configuration"].(bool),
					Mode:                   controllerData["mode"].(string),
					PredictiveSpareRebuild: controllerData["predictive_spare_rebuild"].(string),
					LogicalDrives:          logicalDrives,
				})
			}

			// Gets Local Storage Sas Jbods Body
			rawLocalStorageSasJbod := localStorageItem["sas_logical_jbod"].([]interface{})
			logicalJbod := make([]ov.LogicalJbod, 0)
			for _, raw3 := range rawLocalStorageSasJbod {
				sasLogicalJbodData := raw3.(map[string]interface{})
				logicalJbod = append(logicalJbod, ov.LogicalJbod{
					Description:       sasLogicalJbodData["description"].(string),
					DeviceSlot:        sasLogicalJbodData["device_slot"].(string),
					DriveMaxSizeGB:    sasLogicalJbodData["drive_max_size_gb"].(int),
					DriveMinSizeGB:    sasLogicalJbodData["drive_min_size_gb"].(int),
					DriveTechnology:   sasLogicalJbodData["drive_technology"].(string),
					EraseData:         sasLogicalJbodData["erase_data"].(bool),
					ID:                sasLogicalJbodData["id"].(int),
					Name:              sasLogicalJbodData["name"].(string),
					NumPhysicalDrives: sasLogicalJbodData["num_physical_drive"].(int),
					Persistent:        sasLogicalJbodData["persistent"].(bool),
					SasLogicalJBODUri: utils.NewNstring(sasLogicalJbodData["sas_logical_jbod_uri"].(string)),
				})
			}
			localStorage = ov.LocalStorageOptions{
				ManageLocalStorage: localStorageItem["manage_local_storage"].(bool),
				Initialize:         localStorageItem["initialize"].(bool),
				Controllers:        localStorageEmbeddedController,
				SasLogicalJBODs:    logicalJbod,
			}
		}
		serverProfile.LocalStorage = localStorage
	}

	// get SAN storage data if provided
	if _, ok := d.GetOk("san_storage"); ok {
		rawSanStorage := d.Get("san_storage").(*schema.Set).List()
		sanStorage := ov.SanStorageOptions{}
		for _, raw := range rawSanStorage {
			sanStorageItem := raw.(map[string]interface{})
			sanStorage = ov.SanStorageOptions{
				HostOSType:       sanStorageItem["host_os_type"].(string),
				ManageSanStorage: sanStorageItem["manage_san_storage"].(bool),
			}
		}
		serverProfile.SanStorage = sanStorage
	}
	if _, ok := d.GetOk("volume_attachments"); ok {
		rawVolumeAttachments := d.Get("volume_attachments").([]interface{})
		volumeAttachments := make([]ov.VolumeAttachment, 0)
		for _, rawVolumeAttachment := range rawVolumeAttachments {
			volumeAttachmentItem := rawVolumeAttachment.(map[string]interface{})
			volumes := ov.Volume{}
			if volumeAttachmentItem["volume"] != nil {
				rawVolume := volumeAttachmentItem["volume"].([]interface{})
				for _, rawVol := range rawVolume {
					volumeItem := rawVol.(map[string]interface{})
					tempIsPermanent := volumeItem["is_permanent"].(bool)
					properties := ov.PropertiesSP{}
					if volumeItem["properties"] != nil {
						rawVolumeProperties := volumeItem["properties"].([]interface{})
						for _, rawVolProp := range rawVolumeProperties {
							propertyItem := rawVolProp.(map[string]interface{})
							tempIsShareable := propertyItem["is_shareable"].(bool)
							properties = ov.PropertiesSP{
								DataProtectionLevel:           propertyItem["data_protection_level"].(string),
								DataTransferLimit:             propertyItem["data_transfer_limit"].(int),
								Description:                   propertyItem["description"].(string),
								Folder:                        propertyItem["folder"].(string),
								IopsLimit:                     propertyItem["iops_limit"].(int),
								IsShareable:                   &tempIsShareable,
								IsEncrypted:                   propertyItem["is_encrypted"].(bool),
								IsDeduplicated:                propertyItem["is_deduplicated"].(bool),
								IsPinned:                      propertyItem["is_pinned"].(bool),
								IsCompressed:                  propertyItem["is_compressed"].(bool),
								IsAdaptiveOptimizationEnabled: propertyItem["is_adaptive_optimization_enabled"].(bool),
								IsDataReductionEnabled:        propertyItem["is_data_reduction_enabled"].(bool),
								Name:                          propertyItem["name"].(string),
								PerformancePolicy:             propertyItem["performance_policy"].(string),
								ProvisioningType:              propertyItem["provisioning_type"].(string),
								Size:                          propertyItem["size"].(int),
								SnapshotPool:                  utils.NewNstring(propertyItem["snapshot_pool"].(string)),
								StoragePool:                   utils.NewNstring(propertyItem["storage_pool"].(string)),
								TemplateVersion:               propertyItem["template_version"].(string),
								VolumeSet:                     utils.NewNstring(propertyItem["volume_set"].(string)),
							}
						}
					}

					if val, ok := d.GetOk(volumeItem["initial_scope_uris"].(string)); ok {
						rawInitialScopeUris := val.(*schema.Set).List()
						initialScopeUris := make([]utils.Nstring, len(rawInitialScopeUris))
						for i, raw := range rawInitialScopeUris {
							initialScopeUris[i] = utils.Nstring(raw.(string))
						}
						volumes.InitialScopeUris = initialScopeUris
					}
					volumes = ov.Volume{
						IsPermanent: &tempIsPermanent,
						Properties:  &properties,

						TemplateUri: utils.NewNstring(volumeItem["template_uri"].(string)),
					}
				}
			}
			// get volumeAttachemts.storagepaths
			storagePaths := make([]ov.StoragePath, 0)
			if volumeAttachmentItem["storage_paths"] != nil {
				rawStoragePaths := volumeAttachmentItem["storage_paths"].([]interface{}) //.(*schema.Set).List()

				for _, rawStoragePath := range rawStoragePaths {
					storagePathItem := rawStoragePath.(map[string]interface{})

					// get volumeAttachemts.storagepaths.targets
					targets := make([]ov.Target, 0)
					if storagePathItem["targets"] != nil {
						rawStorageTargets := storagePathItem["targets"].([]interface{}) //.(*schema.Set).List()
						for _, rawStorageTarget := range rawStorageTargets {
							storageTargetItem := rawStorageTarget.(map[string]interface{})
							targets = append(targets, ov.Target{
								IpAddress: storageTargetItem["ip_address"].(string),
								Name:      storageTargetItem["name"].(string),
								TcpPort:   storageTargetItem["tcp_port"].(string),
							})
						}
					}

					storagePaths = append(storagePaths, ov.StoragePath{
						ConnectionID:   storagePathItem["connection_id"].(int),
						IsEnabled:      storagePathItem["is_enabled"].(bool),
						NetworkUri:     utils.NewNstring(storagePathItem["network_uri"].(string)),
						Status:         storagePathItem["status"].(string),
						Targets:        targets,
						TargetSelector: storagePathItem["target_selector"].(string),
					})
				}
			}
			volumeAttachments = append(volumeAttachments, ov.VolumeAttachment{
				ID:                             volumeAttachmentItem["id"].(int),
				LUN:                            volumeAttachmentItem["lun"].(string),
				LUNType:                        volumeAttachmentItem["lun_type"].(string),
				VolumeURI:                      utils.NewNstring(volumeAttachmentItem["volume_uri"].(string)),
				VolumeStorageSystemURI:         utils.NewNstring(volumeAttachmentItem["volume_storage_system_uri"].(string)),
				AssociatedTemplateAttachmentId: volumeAttachmentItem["associated_template_attachment_id"].(string),
				State:                          volumeAttachmentItem["state"].(string),
				Status:                         volumeAttachmentItem["status"].(string),
				StoragePaths:                   storagePaths,
				BootVolumePriority:             volumeAttachmentItem["boot_volume_priority"].(string),
				Volume:                         &volumes,
			})
		}
		serverProfile.SanStorage.VolumeAttachments = volumeAttachments
	}

	if val, ok := d.GetOk("os_deployment_settings"); ok {
		rawOsDeploySetting := val.(*schema.Set).List()
		osDeploySetting := ov.OSDeploymentSettings{}
		for _, raw := range rawOsDeploySetting {
			osDeploySettingItem := raw.(map[string]interface{})
			osdp := ""
			if osDeploySettingItem["os_deployment_plan_uri"] != "" {
				osdp = osDeploySettingItem["os_deployment_plan_uri"].(string)
			} else if osDeploySettingItem["os_deployment_plan_name"] != "" {
				osDeploymentPlan, err := config.ovClient.GetOSDeploymentPlanByName(osDeploySettingItem["os_deployment_plan_name"].(string))
				if err != nil {
					return err
				}
				if osDeploymentPlan.URI == "" {
					return fmt.Errorf("Could not find deployment plan by name: %s", osDeploySettingItem["os_deployment_plan_name"].(string))
				}
				osdp = osDeploymentPlan.URI.String()
			}

			osCustomAttributes := make([]ov.OSCustomAttribute, 0)
			if osDeploySettingItem["os_custom_attributes"] != nil {
				rawOsDeploySettings := osDeploySettingItem["os_custom_attributes"].(*schema.Set).List()
				for _, rawDeploySetting := range rawOsDeploySettings {
					rawOsDeploySetting := rawDeploySetting.(map[string]interface{})
					osCustomAttributes = append(osCustomAttributes, ov.OSCustomAttribute{
						Name:  rawOsDeploySetting["name"].(string),
						Value: rawOsDeploySetting["value"].(string),
					})
				}
			}

			osDeploySetting = ov.OSDeploymentSettings{
				ForceOsDeployment:   osDeploySettingItem["force_os_deployment"].(bool),
				DeployMethod:        osDeploySettingItem["deploy_method"].(string),
				DeploymentMac:       osDeploySettingItem["deployment_mac"].(string),
				DeploymentPortId:    osDeploySettingItem["deployment_port_id"].(string),
				OSDeploymentPlanUri: utils.NewNstring(osdp),
				OSCustomAttributes:  osCustomAttributes,
			}
		}
		serverProfile.OSDeploymentSettings = osDeploySetting
	}
	//Cleaning up SP  by removing spt related fields
	cleanupSp(&serverProfile)

	err := config.ovClient.SubmitNewProfile(serverProfile)
	d.SetId(d.Get("name").(string))

	if err != nil {
		d.SetId("")
		return err
	} else if d.Get("power_state").(string) == "on" {
		if err := serverHardware.PowerOn(); err != nil {
			return err
		}
	}
	return resourceServerProfileRead(d, meta)
}

func cleanupSp(sp *ov.ServerProfile) {
	sp.Bios.ComplianceControl = ""
	sp.Boot.ComplianceControl = ""
	sp.BootMode.ComplianceControl = ""
	sp.ConnectionSettings.ComplianceControl = ""
	sp.Firmware.ComplianceControl = ""
	sp.LocalStorage.ComplianceControl = ""
	sp.ManagementProcessor.ComplianceControl = ""
	sp.OSDeploymentSettings.ComplianceControl = ""
	sp.SanStorage.ComplianceControl = ""
	sp.ServerProfileDescription = ""
}

func resourceServerProfileRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	serverProfile, err := config.ovClient.GetProfileByName(d.Id())

	if err != nil || serverProfile.URI.IsNil() {
		d.SetId("")
		return nil
	}

	// when server hardware is assigned
	if serverProfile.ServerHardwareURI != "" {
		serverHardware, err := config.ovClient.GetServerHardwareByUri(serverProfile.ServerHardwareURI)
		if err != nil {
			return err
		}
		d.Set("enclosure_bay", serverProfile.EnclosureBay)
		d.Set("enclosure_uri", serverProfile.EnclosureURI)
		d.Set("hardware_uri", serverHardware.URI.String())
		d.Set("hardware_name", serverHardware.Name)
		d.Set("ilo_ip", serverHardware.GetIloIPAddress())
	}
	d.Set("serial_number", serverProfile.SerialNumber.String())

	if val, ok := d.GetOk("public_connection"); ok {
		publicConnection, err := serverProfile.GetConnectionByName(val.(string))
		if err != nil {
			return err
		}
		d.Set("public_mac", publicConnection.MAC)
		d.Set("public_slot_id", publicConnection.ID)
	}

	d.Set("name", serverProfile.Name)
	d.Set("type", serverProfile.Type)
	d.Set("uri", serverProfile.URI.String())

	enclosureGroup, err := config.ovClient.GetEnclosureGroupByUri(serverProfile.EnclosureGroupURI)
	if err != nil {
		return err
	}
	d.Set("enclosure_group", enclosureGroup.Name)

	serverHardwareType, err := config.ovClient.GetServerHardwareTypeByUri(serverProfile.ServerHardwareTypeURI)
	if err != nil {
		return err
	}
	d.Set("server_hardware_type", serverHardwareType.Name)
	d.Set("affinity", serverProfile.Affinity)
	d.Set("serial_number_type", serverProfile.SerialNumberType)
	d.Set("wwn_type", serverProfile.WWNType)
	d.Set("mac_type", serverProfile.MACType)
	d.Set("hide_unused_flex_nics", serverProfile.HideUnusedFlexNics)
	d.Set("associated_server", serverProfile.AssociatedServer.String())
	d.Set("category", serverProfile.Category)
	d.Set("created", serverProfile.Created)
	d.Set("description", serverProfile.Description)
	d.Set("etag", serverProfile.ETAG)
	d.Set("in_progress", serverProfile.InProgress)
	d.Set("initial_scope_uris", serverProfile.InitialScopeUris)
	d.Set("iscsi_initiator_name", serverProfile.IscsiInitiatorName)
	d.Set("iscsi_initiator_name_type", serverProfile.IscsiInitiatorNameType)
	d.Set("modified", serverProfile.Modified)
	d.Set("profile_uuid", serverProfile.ProfileUUID.String())
	d.Set("refresh_state", serverProfile.RefreshState)
	d.Set("scopes_uri", serverProfile.ScopesUri)
	d.Set("server_hardware_reapply_state", serverProfile.ServerHardwareReapplyState)
	d.Set("server_hardware_type_uri", serverProfile.ServerHardwareTypeURI.String())
	d.Set("service_manager", serverProfile.ServiceManager)
	d.Set("state", serverProfile.State)
	d.Set("status", serverProfile.Status)
	d.Set("task_uri", serverProfile.TaskURI.String())
	d.Set("template", serverProfile.ServerProfileTemplateURI.String())
	d.Set("template_compliance", serverProfile.TemplateCompliance)
	d.Set("uuid", serverProfile.UUID.String())

	if len(serverProfile.ConnectionSettings.Connections) != 0 {
		// Get connections
		connections := make([]map[string]interface{}, 0, len(serverProfile.ConnectionSettings.Connections))
		for _, connection := range serverProfile.ConnectionSettings.Connections {
			iscsi := make([]map[string]interface{}, 0, 1)
			connectionBoot := make([]map[string]interface{}, 0, 1)
			bootTargets := make([]map[string]interface{}, 0, len(connection.Boot.Targets))
			// Gets Boot Settings
			if connection.Boot != nil {

				if connection.Boot.Targets != nil {
					for _, bootTarget := range connection.Boot.Targets {
						bootTargets = append(bootTargets, map[string]interface{}{
							"lun":        bootTarget.LUN,
							"array_wwpn": bootTarget.ArrayWWPN,
						})
					}
				}

				if connection.Boot.Iscsi != nil {
					iscsi = append(iscsi, map[string]interface{}{
						"chap_level":              connection.Boot.Iscsi.Chaplevel,
						"initiator_name_source":   connection.Boot.Iscsi.InitiatorNameSource,
						"first_boot_target_ip":    connection.Boot.Iscsi.FirstBootTargetIp,
						"first_boot_target_port":  connection.Boot.Iscsi.FirstBootTargetPort,
						"second_boot_target_ip":   connection.Boot.Iscsi.SecondBootTargetIp,
						"second_boot_target_port": connection.Boot.Iscsi.SecondBootTargetPort,
						"mutual_chap_name":        connection.Boot.Iscsi.MutualChapName,
						"mutual_chap_secret":      connection.Boot.Iscsi.MutualChapSecret,
						"boot_target_name":        connection.Boot.Iscsi.BootTargetName,
						"boot_target_lun":         connection.Boot.Iscsi.BootTargetLun,
						"chap_name":               connection.Boot.Iscsi.ChapName,
						"chap_secret":             connection.Boot.Iscsi.ChapSecret,
						"initiator_name":          connection.Boot.Iscsi.InitiatorName,
					})
				}
				connectionBoot = append(connectionBoot, map[string]interface{}{
					"priority":           connection.Boot.Priority,
					"boot_vlan_id":       connection.Boot.BootOptionV3.BootVlanId,
					"ethernet_boot_type": connection.Boot.EthernetBootType,
					"boot_volume_source": connection.Boot.BootVolumeSource,
					"iscsi":              iscsi,
					"boot_target":        bootTargets,
				})
			}
			// Get IPV4 Settings for Connection
			connectionIpv4 := make([]map[string]interface{}, 0, 1)
			if connection.Ipv4 != nil {
				connectionIpv4 = append(connectionIpv4, map[string]interface{}{
					"gateway":           connection.Ipv4.Gateway,
					"ip_address":        connection.Ipv4.IpAddress,
					"subnet_mask":       connection.Ipv4.SubnetMask,
					"ip_address_source": connection.Ipv4.IpAddressSource,
				})
			}
			// Gets Connection Body
			connections = append(connections, map[string]interface{}{
				"function_type":          connection.FunctionType,
				"allocated_mbps":         connection.AllocatedMbps,
				"allocated_vfs":          connection.Connectionv200.AllocatedVFs,
				"interconnect_uri":       connection.InterconnectURI.String(),
				"maximum_mbps":           connection.MaximumMbps,
				"state":                  connection.State,
				"status":                 connection.Status,
				"private_vlan_port_type": connection.PrivateVlanPortType,
				"network_uri":            connection.NetworkURI,
				"port_id":                connection.PortID,
				"requested_mbps":         connection.RequestedMbps,
				"id":                     connection.ID,
				"name":                   connection.Name,
				"isolated_trunk":         connection.IsolatedTrunk,
				"lag_name":               connection.LagName,
				"mac_type":               connection.MacType,
				"managed":                connection.Managed,
				"network_name":           connection.NetworkName,
				"boot":                   connectionBoot,
				"ipv4":                   connectionIpv4,
			})
		}
		// Connection Settings
		connectionSettings := make([]map[string]interface{}, 0, 1)
		connectionSettings = append(connectionSettings, map[string]interface{}{
			"manage_connections": serverProfile.ConnectionSettings.ManageConnections,
			"compliance_control": serverProfile.ConnectionSettings.ComplianceControl,
			"reapply_state":      serverProfile.ConnectionSettings.ReapplyState,
			"connections":        connections,
		})
		d.Set("connection_settings", connectionSettings)
	}

	bootOrder := make([]interface{}, 0)
	if len(serverProfile.Boot.Order) != 0 {
		for _, currBoot := range serverProfile.Boot.Order {
			bootOrder = append(bootOrder, currBoot)
		}
	}
	boot := make([]map[string]interface{}, 0, 1)
	boot = append(boot, map[string]interface{}{
		"manage_boot": serverProfile.Boot.ManageBoot,
		"boot_order":  bootOrder,
	})
	d.Set("boot", boot)

	emptyBootMode := ov.BootModeOption{}
	if serverProfile.BootMode != emptyBootMode {
		bootMode := make([]map[string]interface{}, 0, 1)
		bootMode = append(bootMode, map[string]interface{}{
			"manage_mode":     serverProfile.BootMode.ManageMode,
			"mode":            serverProfile.BootMode.Mode,
			"pxe_boot_policy": serverProfile.BootMode.PXEBootPolicy,
			"secure_boot":     serverProfile.BootMode.SecureBoot,
		})
		d.Set("boot_mode", bootMode)
	}

	if serverProfile.Bios != nil {
		biosOptions := make([]map[string]interface{}, 0, 1)
		overriddenSettings := make([]interface{}, 0)
		if len(serverProfile.Bios.OverriddenSettings) > 0 {
			for _, overriddenSetting := range serverProfile.Bios.OverriddenSettings {
				overriddenSettings = append(overriddenSettings, map[string]interface{}{
					"id":    overriddenSetting.ID,
					"value": overriddenSetting.Value,
				})
			}
		}
		biosOptions = append(biosOptions, map[string]interface{}{
			"manage_bios":         serverProfile.Bios.ManageBios,
			"reapply_state":       serverProfile.Bios.ReapplyState,
			"consistency_state":   serverProfile.Bios.ConsistencyState,
			"overridden_settings": overriddenSettings,
		})

		d.Set("bios_option", biosOptions)
	}

	if len(serverProfile.LocalStorage.Controllers) != 0 {
		// Gets Local Storage Body
		localStorage := make([]map[string]interface{}, 0, 1)
		// Gets Storage Controller Body
		controllers := make([]map[string]interface{}, 0, len(serverProfile.LocalStorage.Controllers))
		for i := 0; i < len(serverProfile.LocalStorage.Controllers); i++ {
			logicalDrives := make([]map[string]interface{}, 0, len(serverProfile.LocalStorage.Controllers[i].LogicalDrives))
			for j := 0; j < len(serverProfile.LocalStorage.Controllers[i].LogicalDrives); j++ {
				logicalDrives = append(logicalDrives, map[string]interface{}{
					"bootable":            serverProfile.LocalStorage.Controllers[i].LogicalDrives[j].Bootable,
					"accelerator":         serverProfile.LocalStorage.Controllers[i].LogicalDrives[j].Accelerator,
					"drive_number":        serverProfile.LocalStorage.Controllers[i].LogicalDrives[j].DriveNumber,
					"drive_technology":    serverProfile.LocalStorage.Controllers[i].LogicalDrives[j].DriveTechnology,
					"name":                serverProfile.LocalStorage.Controllers[i].LogicalDrives[j].Name,
					"num_physical_drives": serverProfile.LocalStorage.Controllers[i].LogicalDrives[j].NumPhysicalDrives,
					"num_spare_drives":    serverProfile.LocalStorage.Controllers[i].LogicalDrives[j].NumSpareDrives,
					"sas_logical_jbod_id": serverProfile.LocalStorage.Controllers[i].LogicalDrives[j].SasLogicalJBODId,
					"raid_level":          serverProfile.LocalStorage.Controllers[i].LogicalDrives[j].RaidLevel,
				})
			}
			controllers = append(controllers, map[string]interface{}{
				"device_slot":              serverProfile.LocalStorage.Controllers[i].DeviceSlot,
				"initialize":               serverProfile.LocalStorage.Controllers[i].Initialize,
				"import_configuration":     serverProfile.LocalStorage.Controllers[i].ImportConfiguration,
				"drive_write_cache":        serverProfile.LocalStorage.Controllers[i].DriveWriteCache,
				"mode":                     serverProfile.LocalStorage.Controllers[i].Mode,
				"predictive_spare_rebuild": serverProfile.LocalStorage.Controllers[i].PredictiveSpareRebuild,
				"logical_drives":           logicalDrives,
			})
		}
		// Gets Sas Logical Jbod Controller Body
		sasLogDrives := make([]map[string]interface{}, 0, len(serverProfile.LocalStorage.SasLogicalJBODs))
		for i := 0; i < len(serverProfile.LocalStorage.SasLogicalJBODs); i++ {
			sasLogDrives = append(sasLogDrives, map[string]interface{}{
				"description":          serverProfile.LocalStorage.SasLogicalJBODs[i].Description,
				"device_slot":          serverProfile.LocalStorage.SasLogicalJBODs[i].DeviceSlot,
				"drive_max_size_gb":    serverProfile.LocalStorage.SasLogicalJBODs[i].DriveMaxSizeGB,
				"drive_min_size_sb":    serverProfile.LocalStorage.SasLogicalJBODs[i].DriveMinSizeGB,
				"drive_technology":     serverProfile.LocalStorage.SasLogicalJBODs[i].DriveTechnology,
				"erase_data":           serverProfile.LocalStorage.SasLogicalJBODs[i].EraseData,
				"id":                   serverProfile.LocalStorage.SasLogicalJBODs[i].ID,
				"name":                 serverProfile.LocalStorage.SasLogicalJBODs[i].Name,
				"num_physical_drive":   serverProfile.LocalStorage.SasLogicalJBODs[i].NumPhysicalDrives,
				"persistent":           serverProfile.LocalStorage.SasLogicalJBODs[i].Persistent,
				"sas_logical_jbod_uri": serverProfile.LocalStorage.SasLogicalJBODs[i].SasLogicalJBODUri.String(),
				"status":               serverProfile.LocalStorage.SasLogicalJBODs[i].Status,
			})
		}
		localStorage = append(localStorage, map[string]interface{}{
			"manage_local_storage": serverProfile.LocalStorage.ManageLocalStorage,
			"initialize":           serverProfile.LocalStorage.Initialize,
			"controller":           controllers,
			"reapply_state":        serverProfile.LocalStorage.ReapplyState,
			"sas_logical_jbod":     sasLogDrives,
		})
		d.Set("local_storage", localStorage)

	}

	if serverProfile.OSDeploymentSettings.OSDeploymentPlanUri != "" {
		osCustomAttributes := make([]map[string]interface{}, 0, len(serverProfile.OSDeploymentSettings.OSCustomAttributes))
		for i := 0; i < len(serverProfile.OSDeploymentSettings.OSCustomAttributes); i++ {
			osCustomAttributes = append(osCustomAttributes, map[string]interface{}{
				"name":        serverProfile.OSDeploymentSettings.OSCustomAttributes[i].Name,
				"type":        serverProfile.OSDeploymentSettings.OSCustomAttributes[i].Type,
				"value":       serverProfile.OSDeploymentSettings.OSCustomAttributes[i].Value,
				"constraints": serverProfile.OSDeploymentSettings.OSCustomAttributes[i].Constraints,
			})
		}

		osdp, err := config.ovClient.GetOSDeploymentPlan(serverProfile.OSDeploymentSettings.OSDeploymentPlanUri)
		if err != nil {
			return err
		}
		osDeploymentPlanName := osdp.Name

		osDeploymentSettingslist := make([]map[string]interface{}, 0, 1)
		osDeploymentSettingslist = append(osDeploymentSettingslist, map[string]interface{}{
			"deploy_method":           serverProfile.OSDeploymentSettings.DeployMethod,
			"deployment_mac":          serverProfile.OSDeploymentSettings.DeploymentMac,
			"deployment_port_id":      serverProfile.OSDeploymentSettings.DeploymentPortId,
			"force_os_deployment":     serverProfile.OSDeploymentSettings.ForceOsDeployment,
			"os_custom_attributes":    osCustomAttributes,
			"os_deployment_plan_name": osDeploymentPlanName,
			"os_deployment_plan_uri":  serverProfile.OSDeploymentSettings.OSDeploymentPlanUri.String(),
			"os_volume":               serverProfile.OSDeploymentSettings.OSVolumeUri.String(),
			"reapply_state":           serverProfile.OSDeploymentSettings.ReapplyState,
		})
		d.Set("os_deployment_settings", osDeploymentSettingslist)
	}

	sanSystemCredentials := make([]interface{}, 0)
	if len(serverProfile.SanStorage.SanSystemCredentials) != 0 {
		for i := 0; i < len(serverProfile.SanStorage.SanSystemCredentials); i++ {
			sanSystemCredentials = append(sanSystemCredentials, map[string]interface{}{
				"chap_level":         serverProfile.SanStorage.SanSystemCredentials[i].ChapLevel,
				"storage_system_uri": serverProfile.SanStorage.SanSystemCredentials[i].StorageSystemUri.String(),
			})
		}
	}
	SanStorageOptions := make([]map[string]interface{}, 0, 1)
	SanStorageOptions = append(SanStorageOptions, map[string]interface{}{
		"compliance_control":     serverProfile.SanStorage.ComplianceControl,
		"host_os_type":           serverProfile.SanStorage.HostOSType,
		"manage_san_storage":     serverProfile.SanStorage.ManageSanStorage,
		"san_system_credentials": sanSystemCredentials,
	})
	d.Set("san_storage", SanStorageOptions)

	volumeAttachments := make([]interface{}, 0)
	if len(serverProfile.SanStorage.VolumeAttachments) != 0 {
		for i := 0; i < len(serverProfile.SanStorage.VolumeAttachments); i++ {
			storagePaths := make([]interface{}, 0)
			if len(serverProfile.SanStorage.VolumeAttachments[i].StoragePaths) != 0 {
				for j := 0; j < len(serverProfile.SanStorage.VolumeAttachments[i].StoragePaths); j++ {
					targets := make([]interface{}, 0)
					if len(serverProfile.SanStorage.VolumeAttachments[i].StoragePaths[j].Targets) != 0 {
						for k := 0; k < len(serverProfile.SanStorage.VolumeAttachments[i].StoragePaths[j].Targets); k++ {
							targets = append(targets, map[string]interface{}{
								"ip_address": serverProfile.SanStorage.VolumeAttachments[i].StoragePaths[j].Targets[k].IpAddress,
								"name":       serverProfile.SanStorage.VolumeAttachments[i].StoragePaths[j].Targets[k].Name,
								"tcp_port":   serverProfile.SanStorage.VolumeAttachments[i].StoragePaths[j].Targets[k].TcpPort,
							})
						}

					}
					storagePaths = append(storagePaths, map[string]interface{}{
						"connection_id":   serverProfile.SanStorage.VolumeAttachments[i].StoragePaths[j].ConnectionID,
						"is_enabled":      serverProfile.SanStorage.VolumeAttachments[i].StoragePaths[j].IsEnabled,
						"network_uri":     serverProfile.SanStorage.VolumeAttachments[i].StoragePaths[j].NetworkUri.String(),
						"status":          serverProfile.SanStorage.VolumeAttachments[i].StoragePaths[j].Status,
						"target_selector": serverProfile.SanStorage.VolumeAttachments[i].StoragePaths[j].TargetSelector,
						"targets":         targets,
					})
				}

			}
			volumes := make([]interface{}, 0)
			if serverProfile.SanStorage.VolumeAttachments[i].Volume != nil {

				properties := make([]interface{}, 0)
				if serverProfile.SanStorage.VolumeAttachments[i].Volume.Properties != nil {

					properties = append(properties, map[string]interface{}{
						"data_protection_level":            serverProfile.SanStorage.VolumeAttachments[i].Volume.Properties.DataProtectionLevel,
						"data_transfer_limit":              serverProfile.SanStorage.VolumeAttachments[i].Volume.Properties.DataTransferLimit,
						"description":                      serverProfile.SanStorage.VolumeAttachments[i].Volume.Properties.Description,
						"folder":                           serverProfile.SanStorage.VolumeAttachments[i].Volume.Properties.Folder,
						"iops_limit":                       serverProfile.SanStorage.VolumeAttachments[i].Volume.Properties.IopsLimit,
						"is_deduplicated":                  serverProfile.SanStorage.VolumeAttachments[i].Volume.Properties.IsDeduplicated,
						"is_encrypted":                     serverProfile.SanStorage.VolumeAttachments[i].Volume.Properties.IsEncrypted,
						"is_pinned":                        serverProfile.SanStorage.VolumeAttachments[i].Volume.Properties.IsPinned,
						"is_shareable":                     serverProfile.SanStorage.VolumeAttachments[i].Volume.Properties.IsShareable,
						"name":                             serverProfile.SanStorage.VolumeAttachments[i].Volume.Properties.Name,
						"performance_policy":               serverProfile.SanStorage.VolumeAttachments[i].Volume.Properties.PerformancePolicy,
						"provisioning_type":                serverProfile.SanStorage.VolumeAttachments[i].Volume.Properties.ProvisioningType,
						"size":                             serverProfile.SanStorage.VolumeAttachments[i].Volume.Properties.Size,
						"volume_set":                       serverProfile.SanStorage.VolumeAttachments[i].Volume.Properties.VolumeSet,
						"is_data_reduction_enabled":        serverProfile.SanStorage.VolumeAttachments[i].Volume.Properties.IsDataReductionEnabled,
						"is_adaptive_optimization_enabled": serverProfile.SanStorage.VolumeAttachments[i].Volume.Properties.IsAdaptiveOptimizationEnabled,
						"is_compressed":                    serverProfile.SanStorage.VolumeAttachments[i].Volume.Properties.IsCompressed,
						"snapshot_pool":                    serverProfile.SanStorage.VolumeAttachments[i].Volume.Properties.SnapshotPool,
						"storage_pool":                     serverProfile.SanStorage.VolumeAttachments[i].Volume.Properties.StoragePool,
						"template_version":                 serverProfile.SanStorage.VolumeAttachments[i].Volume.Properties.TemplateVersion,
					})

				}
				volumes = append(volumes, map[string]interface{}{
					"initial_scope_uris": serverProfile.SanStorage.VolumeAttachments[i].Volume.InitialScopeUris,
					"is_permanent":       serverProfile.SanStorage.VolumeAttachments[i].Volume.IsPermanent,
					"template_uri":       serverProfile.SanStorage.VolumeAttachments[i].Volume.TemplateUri.String(),
					"properties":         properties,
				})

			}

			volumeAttachments = append(volumeAttachments, map[string]interface{}{
				"associated_template_attachment_id": serverProfile.SanStorage.VolumeAttachments[i].AssociatedTemplateAttachmentId,
				"boot_volume_priority":              serverProfile.SanStorage.VolumeAttachments[i].BootVolumePriority,
				"id":                                serverProfile.SanStorage.VolumeAttachments[i].ID,
				"is_permanent":                      serverProfile.SanStorage.VolumeAttachments[i].IsPermanent,
				"lun":                               serverProfile.SanStorage.VolumeAttachments[i].LUN,
				"lun_type":                          serverProfile.SanStorage.VolumeAttachments[i].LUNType,
				"state":                             serverProfile.SanStorage.VolumeAttachments[i].State,
				"status":                            serverProfile.SanStorage.VolumeAttachments[i].Status,
				"storage_paths":                     storagePaths,
				"volume_storage_system_uri":         serverProfile.SanStorage.VolumeAttachments[i].VolumeStorageSystemURI,
				"volume_uri":                        serverProfile.SanStorage.VolumeAttachments[i].VolumeURI,
				"volume":                            volumes,
			})
		}

	}
	d.Set("volume_attachments", volumeAttachments)
	return nil
}

// IsZeroOfUnderlyingType returns true if value is null
func IsZeroOfUnderlyingType(x interface{}) bool {
	if reflect.ValueOf(x).Kind() == reflect.Ptr {
		return true
	}
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

// IsStructNil return true if struct is null
func IsStructNil(x interface{}) bool {
	v := reflect.ValueOf(x)
	for j := 0; j < v.NumField(); j++ {
		if !IsZeroOfUnderlyingType(v.Field(j).Interface()) {
			return true
		}
	}
	return false
}

func resourceServerProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	updateType := d.Get("update_type").(string)

	if updateType == "patch" {
		serverProfile := ov.ServerProfile{
			Name: d.Get("name").(string),
			Type: d.Get("type").(string),
			URI:  utils.NewNstring(d.Get("uri").(string)),
		}

		rawOptions := d.Get("options").(*schema.Set).List()
		options := make([]ov.Options, len(rawOptions))
		for i, rawData := range rawOptions {
			option := rawData.(map[string]interface{})
			options[i] = ov.Options{
				Op:    option["op"].(string),
				Path:  option["path"].(string),
				Value: option["value"].(string)}
		}

		error := config.ovClient.PatchServerProfile(serverProfile, options)
		d.SetId(d.Get("name").(string))
		if error != nil {
			d.SetId("")
			return error
		}
	}

	if updateType == "put" {
		serverProfile, err := config.ovClient.GetProfileByName(d.Get("name").(string))

		var serverHardware ov.ServerHardware
		if d.HasChange("hardware_name") {
			val := d.Get("hardware_name")
			var err error
			serverHardware, err = config.ovClient.GetServerHardwareByName(val.(string))
			if err != nil {
				return err
			}
			serverProfile.ServerHardwareURI = serverHardware.URI
		}

		if d.HasChange("template") {
			val := d.Get("template")
			if val != "null" {
				serverProfileTemplate, err := config.ovClient.GetProfileTemplateByName(val.(string))
				if err != nil || serverProfileTemplate.URI.IsNil() {
					return err
				}
				serverProfile.ServerProfileTemplateURI = serverProfileTemplate.URI
			}
		}

		if d.HasChange("affinity") {
			val := d.Get("affinity")
			serverProfile.Affinity = val.(string)
		}

		if d.HasChange("serial_number_type") {
			val := d.Get("serial_number_type")
			serverProfile.SerialNumberType = val.(string)
		}

		if d.HasChange("wwn_type") {
			val := d.Get("wwn_type")
			serverProfile.WWNType = val.(string)
		}

		if d.HasChange("mac_type") {
			val := d.Get("mac_type")
			serverProfile.MACType = val.(string)
		}

		if d.HasChange("hide_unused_flex_nics") {
			val := d.Get("hide_unused_flex_nics")
			serverProfile.HideUnusedFlexNics = val.(bool)
		}

		if d.HasChange("description") {
			val := d.Get("description")
			serverProfile.Description = val.(string)
		}

		if d.HasChange("enclosure_bay") {
			val := d.Get("enclosure_bay")
			serverProfile.EnclosureBay = val.(int)
		}

		if d.HasChange("enclosure_uri") {
			val := d.Get("enclosure_uri")
			serverProfile.EnclosureURI = utils.NewNstring(val.(string))
		}

		if d.HasChange("iscsi_initiator_name") {
			val := d.Get("iscsi_initiator_name")
			serverProfile.IscsiInitiatorName = val.(string)
		}

		if d.HasChange("iscsi_initiator_name_type") {
			val := d.Get("iscsi_initiator_name_type")
			serverProfile.IscsiInitiatorNameType = val.(string)
		}

		if d.HasChange("profile_uuid") {
			val := d.Get("profile_uuid")
			serverProfile.ProfileUUID = utils.NewNstring(val.(string))
		}

		if d.HasChange("scopes_uri") {
			val := d.Get("scopes_uri")
			serverProfile.ScopesUri = utils.NewNstring(val.(string))
		}

		if d.HasChange("server_hardware_type_uri") {
			val := d.Get("server_hardware_type_uri")
			serverProfile.ServerHardwareTypeURI = utils.NewNstring(val.(string))
		}

		if d.HasChange("service_manager") {
			val := d.Get("service_manager")
			serverProfile.ServiceManager = val.(string)
		}

		if d.HasChange("template_compliance") {
			val := d.Get("template_compliance")
			serverProfile.TemplateCompliance = val.(string)
		}

		if d.HasChange("uuid") {
			val := d.Get("uuid")
			serverProfile.UUID = utils.NewNstring(val.(string))
		}

		if d.HasChange("enclosure_group") {
			val := d.Get("enclosure_group")
			enclosureGroup, err := config.ovClient.GetEnclosureGroupByName(val.(string))
			if err != nil {
				return err
			}
			serverProfile.EnclosureGroupURI = enclosureGroup.URI
		}

		if d.HasChange("server_hardware_type") {
			val := d.Get("server_hardware_type")
			serverHardwareType, err := config.ovClient.GetServerHardwareTypeByName(val.(string))
			if err != nil {
				return err
			}
			serverProfile.ServerHardwareTypeURI = serverHardwareType.URI
		}

		if d.HasChange("connection_settings") {
			val := d.Get("connection_settings")
			connections := val.([]interface{})
			for _, rawConSettings := range connections {
				rawConSetting := rawConSettings.(map[string]interface{})
				rawNetwork := rawConSetting["connections"].([]interface{})
				networks := make([]ov.Connection, 0)
				for i := 0; i < len(rawNetwork); i++ {
					rawNetworkItem := rawNetwork[i].(map[string]interface{})
					bootOptions := ov.BootOption{}
					if rawNetworkItem["boot"] != nil {
						rawBoots := rawNetworkItem["boot"].([]interface{})
						for _, rawBoot := range rawBoots {
							bootItem := rawBoot.(map[string]interface{})

							bootTargets := []ov.BootTarget{}
							rawBootTargets := bootItem["boot_target"].([]interface{})
							if rawBootTargets != nil {
								for _, rawBootTarget := range rawBootTargets {
									bootTarget := rawBootTarget.(map[string]interface{})
									bootTargett := ov.BootTarget{
										LUN:       bootTarget["lun"].(string),
										ArrayWWPN: bootTarget["array_wwpn"].(string),
									}
									// checks if all elements are empty
									// skips if empty
									if !IsStructNil(bootTargett) {
										continue
									}
									bootTargets = append(bootTargets, bootTargett)
								}
							}
							iscsi := ov.BootIscsi{}
							if bootItem["iscsi"] != nil {
								rawIscsis := bootItem["iscsi"].([]interface{})
								for _, rawIscsi := range rawIscsis {
									rawIscsiItem := rawIscsi.(map[string]interface{})
									iscsi = ov.BootIscsi{
										BootTargetLun:        rawIscsiItem["boot_target_lun"].(string),
										BootTargetName:       rawIscsiItem["boot_target_name"].(string),
										ChapName:             rawIscsiItem["chap_name"].(string),
										ChapSecret:           rawIscsiItem["chap_secret"].(string),
										InitiatorName:        rawIscsiItem["initiator_name"].(string),
										InitiatorNameSource:  rawIscsiItem["initiator_name_source"].(string),
										MutualChapName:       rawIscsiItem["mutual_chap_name"].(string),
										MutualChapSecret:     rawIscsiItem["mutual_chap_secret"].(string),
										Chaplevel:            rawIscsiItem["chap_level"].(string),
										FirstBootTargetIp:    rawIscsiItem["first_boot_target_ip"].(string),
										FirstBootTargetPort:  rawIscsiItem["first_boot_target_port"].(string),
										SecondBootTargetIp:   rawIscsiItem["second_boot_target_ip"].(string),
										SecondBootTargetPort: rawIscsiItem["second_boot_target_port"].(string),
									}
								}
							}
							bootOptionV3 := ov.BootOptionV3{
								BootVlanId: bootItem["boot_vlan_id"].(int),
							}
							bootOptions = ov.BootOption{
								Priority:         bootItem["priority"].(string),
								BootOptionV3:     bootOptionV3,
								EthernetBootType: bootItem["ethernet_boot_type"].(string),
								BootVolumeSource: bootItem["boot_volume_source"].(string),
								Iscsi:            &iscsi,
								Targets:          bootTargets,
							}
						}
					}
					ipv4 := ov.Ipv4Option{}
					if rawNetworkItem["ipv4"] != nil {
						rawIpv4s := rawNetworkItem["ipv4"].([]interface{})
						for _, rawIpv4 := range rawIpv4s {
							rawIpv4Item := rawIpv4.(map[string]interface{})
							ipv4 = ov.Ipv4Option{
								Gateway:         rawIpv4Item["gateway"].(string),
								SubnetMask:      rawIpv4Item["subnet_mask"].(string),
								IpAddress:       rawIpv4Item["ip_address"].(string),
								IpAddressSource: rawIpv4Item["ip_address_source"].(string),
							}
						}
					}
					connectionV200 := ov.Connectionv200{
						RequestedVFs: rawNetworkItem["requested_vfs"].(string),
					}

					network := ov.Connection{
						ID:               rawNetworkItem["id"].(int),
						IsolatedTrunk:    rawNetworkItem["isolated_trunk"].(bool),
						LagName:          rawNetworkItem["lag_name"].(string),
						MAC:              utils.NewNstring(rawNetworkItem["mac"].(string)),
						MacType:          rawNetworkItem["mac_type"].(string),
						WWPN:             utils.NewNstring(rawNetworkItem["wwpn"].(string)),
						WWPNType:         rawNetworkItem["wwpn_type"].(string),
						Connectionv200:   connectionV200,
						InterconnectPort: rawNetworkItem["interconnect_port"].(int),
						Name:             rawNetworkItem["name"].(string),
						FunctionType:     rawNetworkItem["function_type"].(string),
						NetworkURI:       utils.NewNstring(rawNetworkItem["network_uri"].(string)),
						PortID:           rawNetworkItem["port_id"].(string),
						RequestedMbps:    rawNetworkItem["requested_mbps"].(string),
						Ipv4:             &ipv4,
					}
					// checks if all elements are empty
					// skips connection if empty
					if !IsStructNil(network) {
						continue
					}
					networks = append(networks, network)
					if len(rawNetworkItem["boot"].([]interface{})) != 0 {
						networks[len(networks)-1].Boot = &bootOptions
					}

				}
				serverProfile.ConnectionSettings = ov.ConnectionSettings{
					Connections: networks,
				}
			}

		}

		if d.HasChange("boot") {
			val := d.Get("boot")
			boot := val.([]interface{})
			for _, rawBoots := range boot {
				rawBoot := rawBoots.(map[string]interface{})
				rawBootOrder := rawBoot["boot_order"].(*schema.Set).List()
				bootOrder := make([]string, len(rawBootOrder))
				for i, raw := range rawBootOrder {
					bootOrder[i] = raw.(string)
				}
				serverProfile.Boot = ov.BootManagement{
					ManageBoot: rawBoot["manage_boot"].(bool),
					Order:      bootOrder,
				}
			}
		}
		if d.HasChange("boot_mode") {
			rawBootMode := d.Get("boot_mode").(*schema.Set).List()[0].(map[string]interface{})
			manageMode := rawBootMode["manage_mode"].(bool)

			serverProfile.BootMode = ov.BootModeOption{
				ManageMode:    &manageMode,
				Mode:          rawBootMode["mode"].(string),
				PXEBootPolicy: utils.Nstring(rawBootMode["pxe_boot_policy"].(string)),
			}

		}

		if d.HasChange("bios_option") {
			val := d.Get("bios_option")
			rawBiosOption := val.([]interface{})
			biosOption := ov.BiosOption{}
			for _, raw := range rawBiosOption {
				rawBiosItem := raw.(map[string]interface{})
				if _, ok := d.GetOk("overridden_settings"); ok {
					overriddenSettings := make([]ov.BiosSettings, 0)
					rawOverriddenSetting := rawBiosItem["overridden_settings"].([]interface{})
					for _, raw2 := range rawOverriddenSetting {
						rawOverriddenSettingItem := raw2.(map[string]interface{})
						if d.HasChanges(rawOverriddenSettingItem["id"].(string), rawOverriddenSettingItem["value"].(string)) {
							overriddenSetting := ov.BiosSettings{
								ID:    rawOverriddenSettingItem["id"].(string),
								Value: rawOverriddenSettingItem["value"].(string),
							}
							// checks if all elements are empty
							// skips if empty
							if !IsStructNil(overriddenSetting) {
								continue
							}
							overriddenSettings = append(overriddenSettings, overriddenSetting)
						}
					}
					biosOption = ov.BiosOption{
						OverriddenSettings: overriddenSettings,
					}
				}
				manageBios := rawBiosItem["manage_bios"].(bool)
				biosOption = ov.BiosOption{
					ManageBios: &manageBios,
				}
			}
			serverProfile.Bios = &biosOption
		}

		if d.HasChange("initial_scope_uris") {
			val := d.Get("initial_scope_uris")
			initialScopeUrisOrder := val.(*schema.Set).List()
			initialScopeUris := make([]utils.Nstring, len(initialScopeUrisOrder))
			for i, raw := range initialScopeUrisOrder {
				initialScopeUris[i] = utils.Nstring(raw.(string))
			}
			serverProfile.InitialScopeUris = initialScopeUris
		}

		// Get firmware details
		if d.HasChange("firmware") {
			rawFirmware := d.Get("firmware").(*schema.Set).List()
			firmware := ov.FirmwareOption{}
			for _, raw := range rawFirmware {
				firmwareItem := raw.(map[string]interface{})
				firmware = ov.FirmwareOption{
					ForceInstallFirmware:     firmwareItem["force_install_firmware"].(bool),
					FirmwareBaselineUri:      utils.NewNstring(firmwareItem["firmware_baseline_uri"].(string)),
					ManageFirmware:           firmwareItem["manage_firmware"].(bool),
					FirmwareInstallType:      firmwareItem["firmware_install_type"].(string),
					FirmwareActivationType:   firmwareItem["firmware_activation_type"].(string),
					FirmwareScheduleDateTime: firmwareItem["firmware_schedule_date_time"].(string),
				}
			}
			serverProfile.Firmware = firmware
		}
		if d.HasChange("local_storage") {
			rawLocalStorage := d.Get("local_storage").([]interface{})
			localStorage := ov.LocalStorageOptions{}
			for _, raw := range rawLocalStorage {
				localStorageItem := raw.(map[string]interface{})
				rawLocalStorageController := localStorageItem["controller"].([]interface{})
				localStorageEmbeddedControllers := make([]ov.LocalStorageEmbeddedController, 0)
				for _, raw2 := range rawLocalStorageController {
					controllerData := raw2.(map[string]interface{})
					rawLogicalDrives := controllerData["logical_drives"].([]interface{})
					logicalDrives := make([]ov.LogicalDriveV3, 0)
					for _, rawLogicalDrive := range rawLogicalDrives {
						logicalDrivesItem := rawLogicalDrive.(map[string]interface{})
						boot := logicalDrivesItem["bootable"].(bool)
						logicalDrive := ov.LogicalDriveV3{
							Bootable:          &boot,
							RaidLevel:         logicalDrivesItem["raid_level"].(string),
							Accelerator:       logicalDrivesItem["accelerator"].(string),
							DriveNumber:       logicalDrivesItem["drive_number"].(int),
							DriveTechnology:   logicalDrivesItem["drive_technology"].(string),
							Name:              logicalDrivesItem["name"].(string),
							NumPhysicalDrives: logicalDrivesItem["num_physical_drives"].(int),
							NumSpareDrives:    logicalDrivesItem["num_spare_drives"].(int),
							SasLogicalJBODId:  logicalDrivesItem["sas_logical_jbod_id"].(int),
						}
						// checks if all elements are empty
						// skip if empty
						if !IsStructNil(logicalDrive) {
							continue
						}
						logicalDrives = append(logicalDrives, logicalDrive)
					}
					init := controllerData["initialize"].(bool)
					// Same implementtion can be acquired for controllers as well, to check all the data in the struct.

					localStorageEmbeddedController := ov.LocalStorageEmbeddedController{
						DeviceSlot:             controllerData["device_slot"].(string),
						DriveWriteCache:        controllerData["drive_write_cache"].(string),
						Initialize:             &init,
						ImportConfiguration:    controllerData["import_configuration"].(bool),
						Mode:                   controllerData["mode"].(string),
						PredictiveSpareRebuild: controllerData["predictive_spare_rebuild"].(string),
						LogicalDrives:          logicalDrives,
					}
					// checks if all elements are empty
					// skips connection if empty
					if !IsStructNil(localStorageEmbeddedController) {
						continue
					}
					localStorageEmbeddedControllers = append(localStorageEmbeddedControllers, localStorageEmbeddedController)

				}
				rawLocalStorageSasJbod := localStorageItem["sas_logical_jbod"].([]interface{})
				logicalJbods := make([]ov.LogicalJbod, 0)
				for _, raw3 := range rawLocalStorageSasJbod {
					sasLogicalJbodData := raw3.(map[string]interface{})
					logicalJbod := ov.LogicalJbod{
						Description:       sasLogicalJbodData["description"].(string),
						DeviceSlot:        sasLogicalJbodData["device_slot"].(string),
						DriveMaxSizeGB:    sasLogicalJbodData["drive_max_size_gb"].(int),
						DriveMinSizeGB:    sasLogicalJbodData["drive_min_size_gb"].(int),
						DriveTechnology:   sasLogicalJbodData["drive_technology"].(string),
						EraseData:         sasLogicalJbodData["erase_data"].(bool),
						ID:                sasLogicalJbodData["id"].(int),
						Name:              sasLogicalJbodData["name"].(string),
						NumPhysicalDrives: sasLogicalJbodData["num_physical_drive"].(int),
						Persistent:        sasLogicalJbodData["persistent"].(bool),
						SasLogicalJBODUri: utils.NewNstring(sasLogicalJbodData["sas_logical_jbod_uri"].(string)),
					}
					// checks if all elements are empty
					// skip if empty
					if !IsStructNil(logicalJbod) {
						continue
					}
					logicalJbods = append(logicalJbods, logicalJbod)
				}
				localStorage = ov.LocalStorageOptions{
					ManageLocalStorage: localStorageItem["manage_local_storage"].(bool),
					Initialize:         localStorageItem["initialize"].(bool),
					Controllers:        localStorageEmbeddedControllers,
					SasLogicalJBODs:    logicalJbods,
				}
			}
			serverProfile.LocalStorage = localStorage
		}

		// get SAN storage data if provided
		if d.HasChange("san_storage") {
			rawSanStorage := d.Get("san_storage").(*schema.Set).List()
			sanStorage := ov.SanStorageOptions{}
			for _, raw := range rawSanStorage {
				sanStorageItem := raw.(map[string]interface{})
				sanStorage = ov.SanStorageOptions{
					HostOSType:       sanStorageItem["host_os_type"].(string),
					ManageSanStorage: sanStorageItem["manage_san_storage"].(bool),
				}
			}
			serverProfile.SanStorage = sanStorage
		}

		if d.HasChange("volume_attachments") {
			rawVolumeAttachments := d.Get("volume_attachments").([]interface{})
			volumeAttachments := make([]ov.VolumeAttachment, 0)
			for _, rawVolumeAttachment := range rawVolumeAttachments {
				volumeAttachmentItem := rawVolumeAttachment.(map[string]interface{})
				volumes := ov.Volume{}
				if volumeAttachmentItem["volume"] != nil {
					rawVolume := volumeAttachmentItem["volume"].(*schema.Set).List()
					for _, rawVol := range rawVolume {
						volumeItem := rawVol.(map[string]interface{})
						tempIsPermanent := volumeItem["is_permanent"].(bool)
						properties := ov.PropertiesSP{}
						if volumeItem["properties"] != nil {
							rawVolumeProperties := volumeItem["properties"].(*schema.Set).List()
							for _, rawVolProp := range rawVolumeProperties {
								propertyItem := rawVolProp.(map[string]interface{})
								tempIsShareable := propertyItem["is_shareable"].(bool)
								properties = ov.PropertiesSP{
									DataProtectionLevel:           propertyItem["data_protection_level"].(string),
									DataTransferLimit:             propertyItem["data_transfer_limit"].(int),
									Description:                   propertyItem["description"].(string),
									Folder:                        propertyItem["folder"].(string),
									IopsLimit:                     propertyItem["iops_limit"].(int),
									IsShareable:                   &tempIsShareable,
									IsEncrypted:                   propertyItem["is_encrypted"].(bool),
									IsDeduplicated:                propertyItem["is_deduplicated"].(bool),
									IsPinned:                      propertyItem["is_pinned"].(bool),
									IsCompressed:                  propertyItem["is_compressed"].(bool),
									IsAdaptiveOptimizationEnabled: propertyItem["is_adaptive_optimization_enabled"].(bool),
									IsDataReductionEnabled:        propertyItem["is_data_reduction_enabled"].(bool),
									Name:                          propertyItem["name"].(string),
									PerformancePolicy:             propertyItem["performance_policy"].(string),
									ProvisioningType:              propertyItem["provisioning_type"].(string),
									Size:                          propertyItem["size"].(int),
									SnapshotPool:                  utils.NewNstring(propertyItem["snapshot_pool"].(string)),
									StoragePool:                   utils.NewNstring(propertyItem["storage_pool"].(string)),
									TemplateVersion:               propertyItem["template_version"].(string),
									VolumeSet:                     utils.NewNstring(propertyItem["volume_set"].(string)),
								}
							}
						}

						if val, ok := d.GetOk(volumeItem["initial_scope_uris"].(string)); ok {
							rawInitialScopeUris := val.(*schema.Set).List()
							initialScopeUris := make([]utils.Nstring, len(rawInitialScopeUris))
							for i, raw := range rawInitialScopeUris {
								initialScopeUris[i] = utils.Nstring(raw.(string))
							}
							volumes.InitialScopeUris = initialScopeUris
						}
						volumes = ov.Volume{
							IsPermanent: &tempIsPermanent,
							Properties:  &properties,
							TemplateUri: utils.NewNstring(volumeItem["template_uri"].(string)),
						}
					}
				}
				// get volumeAttachemts.storagepaths
				storagePaths := make([]ov.StoragePath, 0)
				if volumeAttachmentItem["storage_paths"] != nil {
					rawStoragePaths := volumeAttachmentItem["storage_paths"].([]interface{})

					for _, rawStoragePath := range rawStoragePaths {
						storagePathItem := rawStoragePath.(map[string]interface{})

						// get volumeAttachemts.storagepaths.targets
						targets := make([]ov.Target, 0)
						if storagePathItem["targets"] != nil {
							rawStorageTargets := storagePathItem["targets"].([]interface{})
							for _, rawStorageTarget := range rawStorageTargets {
								storageTargetItem := rawStorageTarget.(map[string]interface{})
								targets = append(targets, ov.Target{
									IpAddress: storageTargetItem["ip_address"].(string),
									Name:      storageTargetItem["name"].(string),
									TcpPort:   storageTargetItem["tcp_port"].(string),
								})
							}
						}

						storagePath := ov.StoragePath{
							IsEnabled:      storagePathItem["is_enabled"].(bool),
							Status:         storagePathItem["status"].(string),
							ConnectionID:   storagePathItem["connection_id"].(int),
							NetworkUri:     utils.NewNstring(storagePathItem["network_uri"].(string)),
							TargetSelector: storagePathItem["target_selector"].(string),
							Targets:        targets,
						}

						// checks if all elements are empty
						// skips if empty
						if !IsStructNil(storagePath) {
							continue
						}
						storagePaths = append(storagePaths, storagePath)

					}
				}
				volumeAttachment := ov.VolumeAttachment{
					ID:                             volumeAttachmentItem["id"].(int),
					LUN:                            volumeAttachmentItem["lun"].(string),
					LUNType:                        volumeAttachmentItem["lun_type"].(string),
					VolumeURI:                      utils.NewNstring(volumeAttachmentItem["volume_uri"].(string)),
					VolumeStorageSystemURI:         utils.NewNstring(volumeAttachmentItem["volume_storage_system_uri"].(string)),
					AssociatedTemplateAttachmentId: volumeAttachmentItem["associated_template_attachment_id"].(string),
					StoragePaths:                   storagePaths,
					BootVolumePriority:             volumeAttachmentItem["boot_volume_priority"].(string),
					Volume:                         &volumes,
				}
				// checks if all elements are empty
				// skips if empty
				if !IsStructNil(volumeAttachment) {
					continue
				}
				volumeAttachments = append(volumeAttachments, volumeAttachment)

			}
			serverProfile.SanStorage.VolumeAttachments = volumeAttachments
		}

		if d.HasChange("os_deployment_settings") {
			val := d.Get("os_deployment_settings")
			rawOsDeploySetting := val.(*schema.Set).List()
			osDeploySetting := ov.OSDeploymentSettings{}
			for _, raw := range rawOsDeploySetting {
				osDeploySettingItem := raw.(map[string]interface{})
				osdp := ""
				if osDeploySettingItem["os_deployment_plan_uri"] != "" {
					osdp = osDeploySettingItem["os_deployment_plan_uri"].(string)
				} else if osDeploySettingItem["os_deployment_plan_name"] != "" {
					osDeploymentPlan, err := config.ovClient.GetOSDeploymentPlanByName(osDeploySettingItem["os_deployment_plan_name"].(string))
					if err != nil {
						return err
					}
					if osDeploymentPlan.URI == "" {
						return fmt.Errorf("Could not find deployment plan by name: %s", osDeploySettingItem["os_deployment_plan_name"].(string))
					}
					osdp = osDeploymentPlan.URI.String()
				}

				osCustomAttributes := make([]ov.OSCustomAttribute, 0)
				if osDeploySettingItem["os_custom_attributes"] != nil {
					rawOsDeploySettings := osDeploySettingItem["os_custom_attributes"].(*schema.Set).List()
					for _, rawDeploySetting := range rawOsDeploySettings {
						rawOsDeploySetting := rawDeploySetting.(map[string]interface{})
						osCustomAttributes = append(osCustomAttributes, ov.OSCustomAttribute{
							Name:  rawOsDeploySetting["name"].(string),
							Value: rawOsDeploySetting["value"].(string),
						})
					}
				}

				osDeploySetting = ov.OSDeploymentSettings{
					ForceOsDeployment:   osDeploySettingItem["force_os_deployment"].(bool),
					DeployMethod:        osDeploySettingItem["deploy_method"].(string),
					DeploymentMac:       osDeploySettingItem["deployment_mac"].(string),
					DeploymentPortId:    osDeploySettingItem["deployment_port_id"].(string),
					OSDeploymentPlanUri: utils.NewNstring(osdp),
					OSCustomAttributes:  osCustomAttributes,
				}
			}
			serverProfile.OSDeploymentSettings = osDeploySetting
		}

		err = config.ovClient.UpdateServerProfile(serverProfile)

		if err != nil {
			return err
		}
		d.SetId(d.Get("name").(string))

	}

	return resourceServerProfileRead(d, meta)

}

func resourceServerProfileDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteProfile(d.Id())
	if err != nil {
		return err
	}
	return nil
}

func getServerHardware(config *Config, serverProfileTemplate ov.ServerProfile, filters []string) (hw ov.ServerHardware, err error) {
	ovMutexKV.Lock(serverProfileTemplate.EnclosureGroupURI.String())
	defer ovMutexKV.Unlock(serverProfileTemplate.EnclosureGroupURI.String())

	var (
		hwlist ov.ServerHardwareList
		f      = []string{"serverHardwareTypeUri='" + serverProfileTemplate.ServerHardwareTypeURI.String() + "'",
			"serverGroupUri='" + serverProfileTemplate.EnclosureGroupURI.String() + "'",
			"state='NoProfileApplied'"}
	)

	f = append(f, filters...)

	if hwlist, err = config.ovClient.GetServerHardwareList(f, "name:desc", "", "", ""); err != nil {
		if _, ok := err.(*json.SyntaxError); ok && len(filters) > 0 {
			return hw, fmt.Errorf("%s. It's likely your hw_filter(s) are incorrectly formatted", err)
		}
		return hw, err
	}
	for _, h := range hwlist.Members {
		if _, reserved := serverHardwareURIs[h.URI.String()]; !reserved {
			serverHardwareURIs[h.URI.String()] = true // Mark as reserved
			h.Client = config.ovClient                // The SDK GetServerHardwareList method doesn't set the
			// client, so we need to do it here. See https://github.com/HewlettPackard/oneview-golang/issues/103
			return h, nil
		}
	}

	return hw, errors.New("No blades that are compatible with the template are available!")
}
