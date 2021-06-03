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
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceServerProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServerProfileRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"affinity": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Bay",
			},
			"associated_server": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bios_option": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"consistency_state": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"manage_bios": {
							Type:     schema.TypeBool,
							Required: true,
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
						"reapply_state": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"boot": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"boot_order": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						"manage_boot": {
							Type:     schema.TypeBool,
							Optional: true,
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
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_settings": {
				Optional: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
									"boot": {
										Optional: true,
										Type:     schema.TypeList,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"boot_vlan_id": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"boot_volume_source": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"ethernet_boot_type": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"iscsi": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"boot_target_lun": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"boot_target_name": {
																Type:     schema.TypeString,
																Optional: true,
															},
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
															"first_boot_target_ip": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"first_boot_target_port": {
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
												"priority": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"targets": {
													Type:     schema.TypeList,
													Optional: true,
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
											},
										},
									},
									"function_type": {
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
												"ip_address_source": {
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
									"network_uri": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"port_id": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "Lom 1:1-a",
									},
									"private_vlan_port_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"requested_mbps": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "2500",
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
								},
							},
						},
						"manage_connections": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"reapply_state": {
							Type:     schema.TypeString,
							Optional: true,
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
			"enclosure_bay": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"enclosure_group": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"enclosure_group_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enclosure_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"firmware": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"consistency_state": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"firmware_activation_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"firmware_baseline_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"firmware_install_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"firmware_schedule_date_time": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"force_install_firmware": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"manage_firmware": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"reapply_state": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"hardware_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hardware_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hide_unused_flex_nics": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"ilo_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"in_progress": {
				Type:     schema.TypeBool,
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
			"iscsi_initiator_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"iscsi_initiator_name_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"local_storage": {
				Optional: true,
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"controller": {
							Optional: true,
							Type:     schema.TypeList,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"device_slot": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"drive_write_cache": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"import_configuration": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"initialize": {
										Type:     schema.TypeBool,
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
												"name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"accelerator": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"bootable": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"drive_technology": {
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
												"raid_level": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"sas_logical_jbod_id": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"initialize": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"manage_local_storage": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"sas_logical_jbod": {
							Optional: true,
							Type:     schema.TypeList,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
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
									"num_physical_drive": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"persistent": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"mac_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Virtual",
				ForceNew: true,
			},
			"management_processor": {
				Type:     schema.TypeSet,
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
						},
					},
				},
			},
			"profile_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_connection": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_mac": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_slot_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"refresh_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"san_storage": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compliance_control": {
							Type:     schema.TypeString,
							Optional: true,
						},
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
			"scopes_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial_number_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Virtual",
				ForceNew: true,
			},
			"server_hardware_reapply_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_hardware_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"server_hardware_type_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_manager": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"task_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"template": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"template_compliance": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ServerProfileV9",
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Computed: true,
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
						"boot_volume_priority": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"id": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"is_permanent": {
							Type:     schema.TypeBool,
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
						"state": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": {
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

						"storage_paths": {
							Optional: true,
							Type:     schema.TypeSet,
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
										Optional: true,
									},
									"target_selector": {
										Type:     schema.TypeString,
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

			"wwn_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Virtual",
				ForceNew: true,
			},
		},
	}
}

func dataSourceServerProfileRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	serverProfile, err := config.ovClient.GetProfileByName(d.Get("name").(string))
	if err != nil || serverProfile.URI.IsNil() {
		d.SetId("")
		return nil
	}

	d.SetId(d.Get("name").(string))

	d.Set("name", serverProfile.Name)
	d.Set("affinity", serverProfile.Affinity)
	d.Set("associated_server", serverProfile.AssociatedServer.String())

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

	// Boot Mode
	bootMode := make([]map[string]interface{}, 0, 1)
	bootMode = append(bootMode, map[string]interface{}{
		"manage_mode":     serverProfile.BootMode.ManageMode,
		"mode":            serverProfile.BootMode.Mode,
		"pxe_boot_policy": serverProfile.BootMode.PXEBootPolicy,
		"secure_boot":     serverProfile.BootMode.SecureBoot,
	})
	d.Set("boot_mode", bootMode)
	d.Set("category", serverProfile.Category)

	if len(serverProfile.ConnectionSettings.Connections) != 0 {
		// Get connections
		connections := make([]map[string]interface{}, 0, len(serverProfile.ConnectionSettings.Connections))
		for _, connection := range serverProfile.ConnectionSettings.Connections {
			// Gets Boot for Connection
			iscsi := make([]map[string]interface{}, 0, 1)
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
			// Get Boot targets list
			targets := make([]map[string]interface{}, 0)
			if len(connection.Boot.Targets) != 0 {
				targets := make([]map[string]interface{}, 0, len(connection.Boot.Targets))
				for j := 0; j < len(connection.Boot.Targets); j++ {
					targets = append(targets, map[string]interface{}{
						"array_wwpn": connection.Boot.Targets[j].ArrayWWPN,
						"lun":        connection.Boot.Targets[j].LUN,
					})
				}
			}
			// Gets Boot Settings
			connectionBoot := make([]map[string]interface{}, 0, 1)
			if connection.Boot != nil {
				connectionBoot = append(connectionBoot, map[string]interface{}{
					"priority":           connection.Boot.Priority,
					"boot_vlan_id":       connection.Boot.BootOptionV3.BootVlanId,
					"ethernet_boot_type": connection.Boot.EthernetBootType,
					"boot_volume_source": connection.Boot.BootVolumeSource,
					"iscsi":              iscsi,
					"targets":            targets,
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
				"requested_vfs":  connection.RequestedVFs,
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

	d.Set("created", serverProfile.Created)
	d.Set("description", serverProfile.Description)
	d.Set("enclosure_group_uri", serverProfile.EnclosureGroupURI.String())

	enclosureGroup, err := config.ovClient.GetEnclosureGroupByUri(serverProfile.EnclosureGroupURI)
	if err != nil {
		return err
	}
	d.Set("enclosure_group", enclosureGroup.Name)
	d.Set("etag", serverProfile.ETAG)

	// Firmware
	firmware := make([]map[string]interface{}, 0, 1)
	firmware = append(firmware, map[string]interface{}{
		"consistency_state":           serverProfile.Firmware.ConsistencyState,
		"firmware_activation_type":    serverProfile.Firmware.FirmwareActivationType,
		"firmware_baseline_uri":       serverProfile.Firmware.FirmwareBaselineUri,
		"firmware_install_type":       serverProfile.Firmware.FirmwareInstallType,
		"firmware_schedule_date_time": serverProfile.Firmware.FirmwareScheduleDateTime,
		"force_install_firmware":      serverProfile.Firmware.ForceInstallFirmware,
		"manage_firmware":             serverProfile.Firmware.ManageFirmware,
		"reapply_state":               serverProfile.Firmware.ReapplyState,
	})
	d.Set("firmware", firmware)
	d.Set("hide_unused_flex_nics", serverProfile.HideUnusedFlexNics)
	d.Set("in_progress", serverProfile.InProgress)
	d.Set("initial_scope_uris", serverProfile.InitialScopeUris)
	d.Set("iscsi_initiator_name", serverProfile.IscsiInitiatorName)
	d.Set("iscsi_initiator_name_type", serverProfile.IscsiInitiatorNameType)

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
			"logical_drive":            logicalDrives,
		})
	}
	// Gets Sas Logical Jbod Controller Body
	sasLogDrives := make([]map[string]interface{}, 0, len(serverProfile.LocalStorage.SasLogicalJBODs))
	for i := 0; i < len(serverProfile.LocalStorage.SasLogicalJBODs); i++ {
		sasLogDrives = append(sasLogDrives, map[string]interface{}{
			"description":        serverProfile.LocalStorage.SasLogicalJBODs[i].Description,
			"device_slot":        serverProfile.LocalStorage.SasLogicalJBODs[i].DeviceSlot,
			"drive_max_size_gb":  serverProfile.LocalStorage.SasLogicalJBODs[i].DriveMaxSizeGB,
			"drive_min_size_sb":  serverProfile.LocalStorage.SasLogicalJBODs[i].DriveMinSizeGB,
			"drive_technology":   serverProfile.LocalStorage.SasLogicalJBODs[i].DriveTechnology,
			"erase_data":         serverProfile.LocalStorage.SasLogicalJBODs[i].EraseData,
			"id":                 serverProfile.LocalStorage.SasLogicalJBODs[i].ID,
			"name":               serverProfile.LocalStorage.SasLogicalJBODs[i].Name,
			"num_physical_drive": serverProfile.LocalStorage.SasLogicalJBODs[i].NumPhysicalDrives,
			"persistent":         serverProfile.LocalStorage.SasLogicalJBODs[i].Persistent,
		})
	}
	localStorage = append(localStorage, map[string]interface{}{
		"manage_local_storage": serverProfile.LocalStorage.ManageLocalStorage,
		"initialize":           serverProfile.LocalStorage.Initialize,
		"controller":           controllers,
		"sas_logical_jbod":     sasLogDrives,
	})
	d.Set("local_storage", localStorage)
	d.Set("mac_type", serverProfile.MACType)

	// Management Processor
	mpSettings := make([]interface{}, 0, len(serverProfile.ManagementProcessor.MpSettings))
	for _, mpSetting := range serverProfile.ManagementProcessor.MpSettings {
		mpSettings = append(mpSettings, map[string]interface{}{
			"args":         mpSetting.Args,
			"setting_type": mpSetting.SettingType,
		})
	}
	if serverProfile.ManagementProcessor != nil {
		managementProcessor := make([]map[string]interface{}, 0, 1)
		managementProcessor = append(managementProcessor, map[string]interface{}{
			"manage_mp":     serverProfile.ManagementProcessor.ManageMp,
			"reapply_state": serverProfile.ManagementProcessor.ReapplyState,
			"mp_settings":   mpSettings,
		})

		d.Set("management_processor", managementProcessor)
	}
	d.Set("profile_uuid", serverProfile.ProfileUUID.String())

	if val, ok := d.GetOk("public_connection"); ok {
		publicConnection, err := serverProfile.GetConnectionByName(val.(string))
		if err != nil {
			return err
		}
		d.Set("public_mac", publicConnection.MAC)
		d.Set("public_slot_id", publicConnection.ID)
	}

	d.Set("refresh_state", serverProfile.RefreshState)
	d.Set("scopes_uri", serverProfile.ScopesUri)
	d.Set("serial_number", serverProfile.SerialNumber.String())
	d.Set("serial_number_type", serverProfile.SerialNumberType)
	d.Set("server_hardware_reapply_state", serverProfile.ServerHardwareReapplyState)
	d.Set("server_hardware_type_uri", serverProfile.ServerHardwareTypeURI.String())

	serverHardwareType, err := config.ovClient.GetServerHardwareTypeByUri(serverProfile.ServerHardwareTypeURI)
	if err != nil {
		return err
	}
	d.Set("server_hardware_type", serverHardwareType.Name)

	// when server hardware is assigned
	if serverProfile.ServerHardwareURI != "" {
		serverHardware, err := config.ovClient.GetServerHardwareByUri(serverProfile.ServerHardwareURI)
		if err != nil {
			return err
		}
		d.Set("enclosure_bay", serverProfile.EnclosureBay)
		d.Set("enclosure_uri", serverProfile.EnclosureURI.String())
		d.Set("hardware_uri", serverHardware.URI.String())
		d.Set("hardware_name", serverHardware.Name)
		d.Set("ilo_ip", serverHardware.GetIloIPAddress())
	}
	sanSystemCredentials := make([]interface{}, 0)
	if len(serverProfile.SanStorage.SanSystemCredentials) != 0 {
		sanSystemCredentials := make([]map[string]interface{}, 0, len(serverProfile.SanStorage.SanSystemCredentials))
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
			})
		}

	}
	d.Set("volume_attachments", volumeAttachments)

	d.Set("service_manager", serverProfile.ServiceManager)
	d.Set("state", serverProfile.State)
	d.Set("status", serverProfile.Status)
	d.Set("task_uri", serverProfile.TaskURI.String())
	d.Set("template", serverProfile.ServerProfileTemplateURI.String())
	d.Set("template_compliance", serverProfile.TemplateCompliance)
	d.Set("type", serverProfile.Type)
	d.Set("uri", serverProfile.URI.String())
	d.Set("uuid", serverProfile.UUID.String())
	d.Set("wwn_type", serverProfile.WWNType)
	return nil
}
