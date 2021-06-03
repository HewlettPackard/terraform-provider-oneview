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
	"strings"

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
							Optional: true,
						},
						"reapply_state": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"overridden_settings": {
							Optional: true,
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
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manage_connections": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"reapply_state": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"compliance_control": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"connections": {
							Optional: true,
							Type:     schema.TypeList,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"allocated_mbps": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"allocated_vfs": {
										Type:     schema.TypeInt,
										Optional: true,
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
										Default:  "Lom 1:1-a",
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
										Optional: true,
									},
									"status": {
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
										Optional: true,
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
										Optional: true,
									},
									"maximum_mbps": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"network_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"private_vlan_port_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"request_": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"boot": {
										Optional: true,
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
												"iscsi": {
													Type:     schema.TypeList,
													Optional: true,
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
															},
															"initiator_name_source": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"mutual_chap_name": {
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
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"force_install_firmware": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"firmware_baseline_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"consistency_state": {
							Type:     schema.TypeString,
							Optional: true,
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
							Optional: true,
						},
						"manage_firmware": {
							Type:     schema.TypeBool,
							Optional: true,
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
							Optional: true,
						},
						"sas_logical_jbod": {
							Optional: true,
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
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"san_storage": {
				Optional: true,
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
						"server_hardware_type_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"server_hardware_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"serial_number": {
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
			// schema for ov.SanStorage.VolumeAttachments
			"volume_attachments": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"associated_template_attachment_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"lun": {
							Type:     schema.TypeString,
							Required: true,
						},
						"lun_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"boot_volume_priority": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"volume_storage_system_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"volume_uri": {
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
						"volume": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"initial_scope_uris": {
										Type:     schema.TypeString,
										Optional: true,
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
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"volume_set": {
													Type:     schema.TypeString,
													Optional: true,
												},
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
												"is_shareable": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"is_encrypted": {
													Type:     schema.TypeBool,
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
												"is_pinned": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"is_deduplicated": {
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
												"size": {
													Type:     schema.TypeInt,
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
						"storage_paths": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"network_uri": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"target_selector": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"is_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"connection_id": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"targets": {
										Type:     schema.TypeSet,
										Optional: true,
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
													Type:     schema.TypeInt,
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
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Optional: true,
			},
			"os_deployment_settings": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"os_custom_attributes": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
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
			"associated_server": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"category": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"created": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enclosure_bay": {
				Type:     schema.TypeString,
				Optional: true,
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
				Type:     schema.TypeList,
				Optional: true,
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
		if !strings.EqualFold(serverHardware.PowerState, "off") {
			return errors.New("Server Hardware must be powered off to assign to the server profile")
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

	if val, ok := d.GetOk("associated_server"); ok {
		serverProfile.AssociatedServer = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk("category"); ok {
		serverProfile.Category = val.(string)
	}

	if val, ok := d.GetOk("created"); ok {
		serverProfile.Created = val.(string)
	}

	if val, ok := d.GetOk("description"); ok {
		serverProfile.Description = val.(string)
	}

	if val, ok := d.GetOk("etag"); ok {
		serverProfile.ETAG = val.(string)
	}

	if val, ok := d.GetOk("enclosure_bay"); ok {
		serverProfile.EnclosureBay = val.(int)
	}

	if val, ok := d.GetOk("in_progress"); ok {
		serverProfile.InProgress = val.(bool)
	}

	if val, ok := d.GetOk("iscsi_initiator_name"); ok {
		serverProfile.IscsiInitiatorName = val.(string)
	}

	if val, ok := d.GetOk("iscsi_initiator_name_type"); ok {
		serverProfile.IscsiInitiatorNameType = val.(string)
	}

	if val, ok := d.GetOk("modified"); ok {
		serverProfile.Modified = val.(string)
	}

	if val, ok := d.GetOk("profile_uuid"); ok {
		serverProfile.ProfileUUID = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk("refresh_state"); ok {
		serverProfile.RefreshState = val.(string)
	}

	if val, ok := d.GetOk("scopes_uri"); ok {
		serverProfile.ScopesUri = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk("server_hardware_reapply_state"); ok {
		serverProfile.ServerHardwareReapplyState = val.(string)
	}

	if val, ok := d.GetOk("server_hardware_type_uri"); ok {
		serverProfile.ServerHardwareTypeURI = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk("service_manager"); ok {
		serverProfile.ServiceManager = val.(string)
	}

	if val, ok := d.GetOk("state"); ok {
		serverProfile.State = val.(string)
	}

	if val, ok := d.GetOk("status"); ok {
		serverProfile.Status = val.(string)
	}

	if val, ok := d.GetOk("task_uri"); ok {
		serverProfile.TaskURI = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk("template_compliance"); ok {
		serverProfile.TemplateCompliance = val.(string)
	}

	if val, ok := d.GetOk("uuid"); ok {
		serverProfile.UUID = utils.NewNstring(val.(string))
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
									Chaplevel:            rawIscsiItem["chap_level"].(string),
									FirstBootTargetIp:    rawIscsiItem["first_boot_target_ip"].(string),
									FirstBootTargetPort:  rawIscsiItem["first_boot_target_ip"].(string),
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
			rawBootMode := d.Get("boot_mode").(*schema.Set).List()[0].(map[string]interface{})
			manageMode := rawBootMode["manage_mode"].(bool)
			serverProfile.BootMode = ov.BootModeOption{
				ManageMode:    &manageMode,
				Mode:          rawBootMode["mode"].(string),
				PXEBootPolicy: utils.Nstring(rawBootMode["pxe_boot_policy"].(string)),
				SecureBoot:    rawBootMode["secure_boot"].(string),
			}
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
				ConsistencyState:         firmwareItem["consistency_state"].(string),
				FirmwareActivationType:   firmwareItem["firmware_activation_type"].(string),
				FirmwareScheduleDateTime: firmwareItem["firmware_schedule_date_time"].(string),
				ReapplyState:             firmwareItem["reapply_state"].(string),
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
						DriveNumber:       logicalDrivesItem["drive_number"].(int),
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
					Status:            sasLogicalJbodData["status"].(string),
					SasLogicalJBODUri: utils.NewNstring(sasLogicalJbodData["sas_logical_jbod_uri"].(string)),
				})
			}
			localStorage = ov.LocalStorageOptions{
				ManageLocalStorage: localStorageItem["manage_local_storage"].(bool),
				Initialize:         localStorageItem["initialize"].(bool),
				Controllers:        localStorageEmbeddedController,
				ReapplyState:       localStorageItem["reapply_state"].(string),
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
				HostOSType:            sanStorageItem["host_os_type"].(string),
				ManageSanStorage:      sanStorageItem["manage_san_storage"].(bool),
				ServerHardwareTypeURI: utils.NewNstring(sanStorageItem["server_hardware_type_uri"].(string)),
				ServerHardwareURI:     utils.NewNstring(sanStorageItem["server_hardware_uri"].(string)),
				SerialNumber:          sanStorageItem["serial_number"].(string),
				Type:                  sanStorageItem["type"].(string),
				URI:                   utils.NewNstring(sanStorageItem["uri"].(string)),
			}
		}
		serverProfile.SanStorage = sanStorage
	}
	if _, ok := d.GetOk("volume_attachments"); ok {
		rawVolumeAttachments := d.Get("volume_attachments").(*schema.Set).List()
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
					volumes = ov.Volume{
						IsPermanent:      &tempIsPermanent,
						Properties:       &properties,
						InitialScopeUris: utils.NewNstring(volumeItem["initial_scope_uris"].(string)),
						TemplateUri:      utils.NewNstring(volumeItem["template_uri"].(string)),
					}
				}
			}
			// get volumeAttachemts.storagepaths
			storagePaths := make([]ov.StoragePath, 0)
			if volumeAttachmentItem["storage_paths"] != nil {
				rawStoragePaths := volumeAttachmentItem["storage_paths"].(*schema.Set).List()

				for _, rawStoragePath := range rawStoragePaths {
					storagePathItem := rawStoragePath.(map[string]interface{})

					// get volumeAttachemts.storagepaths.targets
					targets := make([]ov.Target, 0)
					if storagePathItem["targets"] != nil {
						rawStorageTargets := storagePathItem["targets"].(*schema.Set).List()
						for _, rawStorageTarget := range rawStorageTargets {
							storageTargetItem := rawStorageTarget.(map[string]interface{})
							targets = append(targets, ov.Target{
								IpAddress: storageTargetItem["ip_address"].(string),
								Name:      storageTargetItem["name"].(string),
								TcpPort:   storageTargetItem["tcp_port"].(int),
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
		for _, raw := range rawOsDeploySetting {
			osDeploySettingItem := raw.(map[string]interface{})

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
			// If Name already imported from SPT, overwrite its value from SP
			for _, temp1 := range osCustomAttributes {
				for j, temp2 := range serverProfile.OSDeploymentSettings.OSCustomAttributes {
					if temp1.Name == temp2.Name {
						serverProfile.OSDeploymentSettings.OSCustomAttributes[j].Value = temp1.Value
					}
				}
			}

		}
	}
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
			// Gets Boot Settings
			if connection.Boot != nil {
				if connection.Boot.Iscsi != nil {
					iscsi = append(iscsi, map[string]interface{}{
						"chap_level":              connection.Boot.Iscsi.Chaplevel,
						"initiator_name_source":   connection.Boot.Iscsi.InitiatorNameSource,
						"first_boot_target_ip":    connection.Boot.Iscsi.FirstBootTargetIp,
						"first_boot_target_port":  connection.Boot.Iscsi.FirstBootTargetPort,
						"second_boot_target_ip":   connection.Boot.Iscsi.SecondBootTargetIp,
						"second_boot_target_port": connection.Boot.Iscsi.SecondBootTargetPort,
					})
				}
				connectionBoot = append(connectionBoot, map[string]interface{}{
					"priority":           connection.Boot.Priority,
					"boot_vlan_id":       connection.Boot.BootOptionV3.BootVlanId,
					"ethernet_boot_type": connection.Boot.EthernetBootType,
					"boot_volume_source": connection.Boot.BootVolumeSource,
					"iscsi":              iscsi,
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
				"function_type":  connection.FunctionType,
				"network_uri":    connection.NetworkURI,
				"port_id":        connection.PortID,
				"requested_mbps": connection.RequestedMbps,
				"id":             connection.ID,
				"name":           connection.Name,
				"isolated_trunk": connection.IsolatedTrunk,
				"lag_name":       connection.LagName,
				"mac_type":       connection.MacType,
				"managed":        connection.Managed,
				"network_name":   connection.NetworkName,
				"boot":           connectionBoot,
				"ipv4":           connectionIpv4,
			})
		}
		// Connection Settings
		connectionSettings := make([]map[string]interface{}, 0, 1)
		connectionSettings = append(connectionSettings, map[string]interface{}{
			"manage_connections": serverProfile.ConnectionSettings.ManageConnections,
			"compliance_control": serverProfile.ConnectionSettings.ComplianceControl,
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

	overriddenSettings := make([]interface{}, 0, len(serverProfile.Bios.OverriddenSettings))
	for _, overriddenSetting := range serverProfile.Bios.OverriddenSettings {
		overriddenSettings = append(overriddenSettings, map[string]interface{}{
			"id":    overriddenSetting.ID,
			"value": overriddenSetting.Value,
		})
	}
	if serverProfile.Bios != nil {
		biosOptions := make([]map[string]interface{}, 0, 1)
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
	return nil
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
		serverProfile := ov.ServerProfile{
			Type: d.Get("type").(string),
			Name: d.Get("name").(string),
			URI:  utils.NewNstring(d.Get("uri").(string)),
		}

		var serverHardware ov.ServerHardware
		if val, ok := d.GetOk("hardware_name"); ok {
			var err error
			serverHardware, err = config.ovClient.GetServerHardwareByName(val.(string))
			if err != nil {
				return err
			}
			if !strings.EqualFold(serverHardware.PowerState, "off") {
				return fmt.Errorf("Server Hardware must be powered off to assign to server profile")
			}
			serverProfile.ServerHardwareURI = serverHardware.URI
		}

		if val, ok := d.GetOk("template"); ok {
			serverProfileTemplate, err := config.ovClient.GetProfileTemplateByName(val.(string))
			if err != nil || serverProfileTemplate.URI.IsNil() {
				return err
			}
			serverProfile.ServerProfileTemplateURI = serverProfileTemplate.URI
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

		if val, ok := d.GetOk("associated_server"); ok {
			serverProfile.AssociatedServer = utils.NewNstring(val.(string))
		}

		if val, ok := d.GetOk("category"); ok {
			serverProfile.Category = val.(string)
		}

		if val, ok := d.GetOk("created"); ok {
			serverProfile.Created = val.(string)
		}

		if val, ok := d.GetOk("description"); ok {
			serverProfile.Description = val.(string)
		}

		if val, ok := d.GetOk("etag"); ok {
			serverProfile.ETAG = val.(string)
		}

		if val, ok := d.GetOk("enclosure_bay"); ok {
			serverProfile.EnclosureBay = val.(int)
		}

		if val, ok := d.GetOk("in_progress"); ok {
			serverProfile.InProgress = val.(bool)
		}

		if val, ok := d.GetOk("iscsi_initiator_name"); ok {
			serverProfile.IscsiInitiatorName = val.(string)
		}

		if val, ok := d.GetOk("iscsi_initiator_name_type"); ok {
			serverProfile.IscsiInitiatorNameType = val.(string)
		}

		if val, ok := d.GetOk("modified"); ok {
			serverProfile.Modified = val.(string)
		}

		if val, ok := d.GetOk("profile_uuid"); ok {
			serverProfile.ProfileUUID = utils.NewNstring(val.(string))
		}

		if val, ok := d.GetOk("refresh_state"); ok {
			serverProfile.RefreshState = val.(string)
		}

		if val, ok := d.GetOk("scopes_uri"); ok {
			serverProfile.ScopesUri = utils.NewNstring(val.(string))
		}

		if val, ok := d.GetOk("server_hardware_reapply_state"); ok {
			serverProfile.ServerHardwareReapplyState = val.(string)
		}

		if val, ok := d.GetOk("server_hardware_type_uri"); ok {
			serverProfile.ServerHardwareTypeURI = utils.NewNstring(val.(string))
		}

		if val, ok := d.GetOk("service_manager"); ok {
			serverProfile.ServiceManager = val.(string)
		}

		if val, ok := d.GetOk("state"); ok {
			serverProfile.State = val.(string)
		}

		if val, ok := d.GetOk("status"); ok {
			serverProfile.Status = val.(string)
		}

		if val, ok := d.GetOk("task_uri"); ok {
			serverProfile.TaskURI = utils.NewNstring(val.(string))
		}

		if val, ok := d.GetOk("template_compliance"); ok {
			serverProfile.TemplateCompliance = val.(string)
		}

		if val, ok := d.GetOk("uuid"); ok {
			serverProfile.UUID = utils.NewNstring(val.(string))
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

							iscsi := ov.BootIscsi{}
							if bootItem["iscsi"] != nil {
								rawIscsis := bootItem["iscsi"].([]interface{})
								for _, rawIscsi := range rawIscsis {
									rawIscsiItem := rawIscsi.(map[string]interface{})
									iscsi = ov.BootIscsi{
										Chaplevel:            rawIscsiItem["chap_level"].(string),
										InitiatorName:        rawIscsiItem["initiator_name"].(string),
										InitiatorNameSource:  rawIscsiItem["initiator_name_source"].(string),
										FirstBootTargetIp:    rawIscsiItem["first_boot_target_ip"].(string),
										FirstBootTargetPort:  rawIscsiItem["first_boot_target_ip"].(string),
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
								SubnetMask:      rawIpv4Item["subne_mask"].(string),
								IpAddress:       rawIpv4Item["ip_address"].(string),
								IpAddressSource: rawIpv4Item["ip_address_source"].(string),
							}
						}
					}
					connectionV200 := ov.Connectionv200{
						RequestedVFs: rawNetworkItem["requested_vfs"].(string),
					}

					networks = append(networks, ov.Connection{
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
				rawBootMode := d.Get("boot_mode").(*schema.Set).List()[0].(map[string]interface{})
				manageMode := rawBootMode["manage_mode"].(bool)
				serverProfile.BootMode = ov.BootModeOption{
					ManageMode:    &manageMode,
					Mode:          rawBootMode["mode"].(string),
					PXEBootPolicy: utils.Nstring(rawBootMode["pxe_boot_policy"].(string)),
				}
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

		// Get firmware details
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
					ConsistencyState:         firmwareItem["consistency_state"].(string),
					FirmwareActivationType:   firmwareItem["firmware_activation_type"].(string),
					FirmwareScheduleDateTime: firmwareItem["firmware_schedule_date_time"].(string),
					ReapplyState:             firmwareItem["reapply_state"].(string),
				}
			}
			serverProfile.Firmware = firmware
		}
		if _, ok := d.GetOk("local_storage"); ok {
			rawLocalStorage := d.Get("local_storage").([]interface{})
			localStorage := ov.LocalStorageOptions{}
			for _, raw := range rawLocalStorage {
				localStorageItem := raw.(map[string]interface{})
				rawLocalStorageController := localStorageItem["controller"].([]interface{})
				localStorageEmbeddedController := make([]ov.LocalStorageEmbeddedController, 0)
				for _, raw2 := range rawLocalStorageController {
					controllerData := raw2.(map[string]interface{})
					rawLogicalDrives := controllerData["logical_drives"].([]interface{})
					logicalDrives := make([]ov.LogicalDriveV3, 0)
					for _, rawLogicalDrive := range rawLogicalDrives {
						logicalDrivesItem := rawLogicalDrive.(map[string]interface{})
						boot := logicalDrivesItem["bootable"].(bool)
						logicalDrives = append(logicalDrives, ov.LogicalDriveV3{
							Bootable:          &boot,
							RaidLevel:         logicalDrivesItem["raid_level"].(string),
							Accelerator:       logicalDrivesItem["accelerator"].(string),
							DriveNumber:       logicalDrivesItem["drive_number"].(int),
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
						Status:            sasLogicalJbodData["status"].(string),
						SasLogicalJBODUri: utils.NewNstring(sasLogicalJbodData["sas_logical_jbod_uri"].(string)),
					})
				}
				localStorage = ov.LocalStorageOptions{
					ManageLocalStorage: localStorageItem["manage_local_storage"].(bool),
					Initialize:         localStorageItem["initialize"].(bool),
					Controllers:        localStorageEmbeddedController,
					ReapplyState:       localStorageItem["reapply_state"].(string),
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
					HostOSType:            sanStorageItem["host_os_type"].(string),
					ManageSanStorage:      sanStorageItem["manage_san_storage"].(bool),
					ServerHardwareTypeURI: utils.NewNstring(sanStorageItem["server_hardware_type_uri"].(string)),
					ServerHardwareURI:     utils.NewNstring(sanStorageItem["server_hardware_uri"].(string)),
					SerialNumber:          sanStorageItem["serial_number"].(string),
					Type:                  sanStorageItem["type"].(string),
					URI:                   utils.NewNstring(sanStorageItem["uri"].(string)),
				}
			}
			serverProfile.SanStorage = sanStorage
		}
		if _, ok := d.GetOk("volume_attachments"); ok {
			rawVolumeAttachments := d.Get("volume_attachments").(*schema.Set).List()
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
						volumes = ov.Volume{
							IsPermanent:      &tempIsPermanent,
							Properties:       &properties,
							InitialScopeUris: utils.NewNstring(volumeItem["initial_scope_uris"].(string)),
							TemplateUri:      utils.NewNstring(volumeItem["template_uri"].(string)),
						}
					}
				}
				// get volumeAttachemts.storagepaths
				storagePaths := make([]ov.StoragePath, 0)
				if volumeAttachmentItem["storage_paths"] != nil {
					rawStoragePaths := volumeAttachmentItem["storage_paths"].(*schema.Set).List()

					for _, rawStoragePath := range rawStoragePaths {
						storagePathItem := rawStoragePath.(map[string]interface{})

						// get volumeAttachemts.storagepaths.targets
						targets := make([]ov.Target, 0)
						if storagePathItem["targets"] != nil {
							rawStorageTargets := storagePathItem["targets"].(*schema.Set).List()
							for _, rawStorageTarget := range rawStorageTargets {
								storageTargetItem := rawStorageTarget.(map[string]interface{})
								targets = append(targets, ov.Target{
									IpAddress: storageTargetItem["ip_address"].(string),
									Name:      storageTargetItem["name"].(string),
									TcpPort:   storageTargetItem["tcp_port"].(int),
								})
							}
						}

						storagePaths = append(storagePaths, ov.StoragePath{
							IsEnabled:      storagePathItem["is_enabled"].(bool),
							Status:         storagePathItem["status"].(string),
							ConnectionID:   storagePathItem["connection_id"].(int),
							NetworkUri:     utils.NewNstring(storagePathItem["network_uri"].(string)),
							TargetSelector: storagePathItem["target_selector"].(string),
							Targets:        targets,
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
			for _, raw := range rawOsDeploySetting {
				osDeploySettingItem := raw.(map[string]interface{})

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
				// If Name already imported from SPT, overwrite its value from SP
				for _, temp1 := range osCustomAttributes {
					for j, temp2 := range serverProfile.OSDeploymentSettings.OSCustomAttributes {
						if temp1.Name == temp2.Name {
							serverProfile.OSDeploymentSettings.OSCustomAttributes[j].Value = temp1.Value
						}
					}
				}

			}
		}

		err := config.ovClient.UpdateServerProfile(serverProfile)
		if err != nil {
			return err
		}
		d.SetId(d.Get("name").(string))

	}

	return resourceServerProfileRead(d, meta)

}

func resourceServerProfileDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteProfile(d.Get("name").(string))
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
