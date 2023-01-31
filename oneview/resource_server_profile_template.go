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
	"log"
	"reflect"
	"strconv"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServerProfileTemplate() *schema.Resource {

	return &schema.Resource{
		Create: resourceServerProfileTemplateCreate,
		Read:   resourceServerProfileTemplateRead,
		Update: resourceServerProfileTemplateUpdate,
		Delete: resourceServerProfileTemplateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"affinity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"bios_option": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compliance_control": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"manage_bios": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"overridden_settings": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
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
			"boot": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compliance_control": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
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
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compliance_control": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"manage_mode": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"mode": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"pxe_boot_policy": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"secure_boot": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
				Computed: true,
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compliance_control": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"connections": {
							Optional: true,
							Computed: true,
							Type:     schema.TypeSet,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"boot": {
										Optional: true,
										Computed: true,
										Type:     schema.TypeList,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"boot_vlan_id": {
													Type:     schema.TypeString,
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
												"boot_target": {
													Type:     schema.TypeSet,
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
												"iscsi": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
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
												"priority": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"targets": {
													Type:     schema.TypeSet,
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
										Computed: true,
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
									"isolated_trunk": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"lag_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
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
									},
									"requested_mbps": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"requested_vfs": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"mac_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"manage_connections": {
							Type:     schema.TypeBool,
							Required: true,
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
				Optional: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enclosure_group": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enclosure_group_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"firmware": {
				Optional: true,
				Computed: true,
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compliance_control": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"firmware_activation_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"firmware_baseline_uri": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"firmware_install_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"force_install_firmware": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"manage_firmware": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"hide_unused_flex_nics": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"initial_scope_uris": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"iscsi_initiator_name_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_storage": {
				Optional: true,
				Computed: true,
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compliance_control": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "CheckedMinimum",
						},
						"controller": {
							Optional: true,
							Computed: true,
							Type:     schema.TypeList,
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
									"initialize": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"import_configuration": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"logical_drives": {
										Optional: true,
										Computed: true,
										Type:     schema.TypeList,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
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
												"name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"num_physical_drives": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"num_spare_drives": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"raid_level": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"sas_logical_jbod_id": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"mode": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"predictive_spare_rebuild": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"sas_logical_jbod": {
							Optional: true,
							Type:     schema.TypeSet,
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
								},
							},
						},
					},
				},
			},
			"mac_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_processor": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compliance_control": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"manage_mp": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"reapply_state": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mp_settings": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"administrator_account": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"delete_administrator_account": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"password": {
													Type:      schema.TypeString,
													Optional:  true,
													Sensitive: true,
												},
											},
										},
									},
									"directory": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"directory_authentication": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"directory_generic_ldap": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"directory_server_address": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"directory_server_port": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"directory_server_certificate": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"directory_user_context": {
													Type:     schema.TypeSet,
													Optional: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
													Set:      schema.HashString,
												},
												"ilo_distinguished_name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"password": {
													Type:      schema.TypeString,
													Optional:  true,
													Sensitive: true,
												},
												"kerberos_authentication": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"kerberos_realm": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"kerberos_kdc_server_address": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"kerberos_kdc_server_port": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"kerberos_key_tab": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"ilo_host_name": {
										MaxItems: 1,
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hostname": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"key_manager": {
										MaxItems: 1,
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"primary_server_address": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"primary_server_port": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"secondary_server_address": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"secondary_server_port": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"redundancy_required": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"group_name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"certificate_name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"login_name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"password": {
													Type:      schema.TypeString,
													Optional:  true,
													Sensitive: true,
												},
											},
										},
									},
									"directory_groups": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"group_dn": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"group_sid": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"user_config_priv": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"remote_console_priv": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"virtual_media_priv": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"virtual_power_and_reset_priv": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"ilo_config_priv": {
													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},
									"local_accounts": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"user_name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"display_name": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"password": {
													Type:      schema.TypeString,
													Optional:  true,
													Sensitive: true,
												},
												"user_config_priv": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"remote_console_priv": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"virtual_media_priv": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"virtual_power_and_reset_priv": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"ilo_config_priv": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"login_priv": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"host_bios_config_priv": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"host_nic_config_priv": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"host_storage_config_priv": {
													Type:     schema.TypeBool,
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
			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"refresh_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"san_storage": {
				Optional: true,
				Computed: true,
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compliance_control": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
			"volume_attachments": {
				Optional: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"associated_template_attachment_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"storage_paths": {
							Optional: true,
							Computed: true,
							Type:     schema.TypeList,
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
													Type:     schema.TypeInt,
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
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"initial_scope_uris": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Set: schema.HashString,
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
												"provisioning_type": {
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

			"scopes_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial_number_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_hardware_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_hardware_type_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_profile_description": {
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

			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"wwn_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceServerProfileTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	serverProfileTemplate := ov.ServerProfile{
		Name:                     d.Get("name").(string),
		Type:                     d.Get("type").(string),
		Affinity:                 d.Get("affinity").(string),
		SerialNumberType:         d.Get("serial_number_type").(string),
		WWNType:                  d.Get("wwn_type").(string),
		MACType:                  d.Get("mac_type").(string),
		HideUnusedFlexNics:       d.Get("hide_unused_flex_nics").(bool),
		Description:              d.Get("description").(string),
		ServerProfileDescription: d.Get("server_profile_description").(string),
	}

	if d.Get("enclosure_group") != "" {
		enclosureGroup, err := config.ovClient.GetEnclosureGroupByName(d.Get("enclosure_group").(string))
		if err != nil {
			return err
		}
		serverProfileTemplate.EnclosureGroupURI = enclosureGroup.URI
	}

	serverHardwareType, err := config.ovClient.GetServerHardwareTypeByName(d.Get("server_hardware_type").(string))
	if err != nil {
		return err
	}
	serverProfileTemplate.ServerHardwareTypeURI = serverHardwareType.URI

	if val, ok := d.GetOk("management_processor"); ok {
		mps := val.([]interface{})
		ovManagementProcessor := ov.ManagementProcessors{}
		for _, rawMp := range mps {
			// extracting management processor
			mp := rawMp.(map[string]interface{})
			// extracting MpSettings
			ovMpSettings := ov.MpSettings{}
			mpSettings := mp["mp_settings"].([]interface{})
			for _, mpSettingg := range mpSettings {
				mpSetting := mpSettingg.(map[string]interface{})
				// extracting administrator account
				ovAdminAcc := ov.AdministratorAccount{}
				rawAdminAcc := mpSetting["administrator_account"].([]interface{})
				for _, adminAccs := range rawAdminAcc {
					adminAcc := adminAccs.(map[string]interface{})
					ovAdminAcc = ov.AdministratorAccount{
						DeleteAdministratorAccount: GetBoolPointer(adminAcc["delete_administrator_account"].(bool)),
						Password:                   adminAcc["password"].(string),
					}
				}
				// extracting directory
				rawDirectory := mpSetting["directory"].([]interface{})
				ovDirectory := ov.Directory{}
				for _, directoryy := range rawDirectory {
					directory := directoryy.(map[string]interface{})

					directoryUserContexts := make([]string, 0)
					for _, raw := range directory["directory_user_context"].(*schema.Set).List() {
						directoryUserContexts = append(directoryUserContexts, raw.(string))
					}

					ovDirectory = ov.Directory{
						DirectoryAuthentication:    directory["directory_authentication"].(string),
						DirectoryGenericLDAP:       GetBoolPointer(directory["directory_generic_ldap"].(bool)),
						DirectoryServerAddress:     directory["directory_server_address"].(string),
						DirectoryServerPort:        directory["directory_server_port"].(int),
						DirectoryServerCertificate: directory["directory_server_certificate"].(string),
						DirectoryUserContext:       directoryUserContexts,
						IloObjectDistinguishedName: directory["ilo_distinguished_name"].(string),
						Password:                   directory["password"].(string),
						KerberosAuthentication:     GetBoolPointer(directory["kerberos_authentication"].(bool)),
						KerberosRealm:              directory["kerberos_realm"].(string),
						KerberosKDCServerAddress:   directory["kerberos_kdc_server_address"].(string),
						KerberosKDCServerPort:      directory["kerberos_kdc_server_port"].(int),
						KerberosKeytab:             directory["kerberos_key_tab"].(string),
					}
				}

				// extracting hostname
				rawHostName := mpSetting["ilo_host_name"].([]interface{})
				ovHostName := ov.IloHostName{}
				for _, hostMap := range rawHostName {
					host := hostMap.(map[string]interface{})
					ovHostName = ov.IloHostName{
						HostName: host["hostname"].(string),
					}
				}

				// extracting key manager
				rawKeyManager := mpSetting["key_manager"].([]interface{})
				ovKeyManager := ov.KeyManager{}
				for _, keyManagerr := range rawKeyManager {
					keyManager := keyManagerr.(map[string]interface{})
					ovKeyManager = ov.KeyManager{
						PrimaryServerAddress:   keyManager["primary_server_address"].(string),
						PrimaryServerPort:      keyManager["primary_server_port"].(int),
						SecondaryServerAddress: keyManager["secondary_server_address"].(string),
						SecondaryServerPort:    keyManager["secondary_server_port"].(int),
						RedundancyRequired:     GetBoolPointer(keyManager["redundancy_required"].(bool)),
						GroupName:              keyManager["group_name"].(string),
						CertificateName:        keyManager["certificate_name"].(string),
						LoginName:              keyManager["login_name"].(string),
						Password:               keyManager["password"].(string),
					}
				}

				// extracting directory groups
				rawDirectoryGroups := mpSetting["directory_groups"].(*schema.Set).List()
				ovDirectoryGroups := make([]ov.DirectoryGroups, 0)
				for _, directoryGroupp := range rawDirectoryGroups {
					directoryGroup := directoryGroupp.(map[string]interface{})
					ovDirectoryGroups = append(ovDirectoryGroups, ov.DirectoryGroups{
						GroupDN:                  directoryGroup["group_dn"].(string),
						GroupSID:                 directoryGroup["group_sid"].(string),
						UserConfigPriv:           GetBoolPointer(directoryGroup["user_config_priv"].(bool)),
						RemoteConsolePriv:        GetBoolPointer(directoryGroup["remote_console_priv"].(bool)),
						VirtualMediaPriv:         GetBoolPointer(directoryGroup["virtual_media_priv"].(bool)),
						VirtualPowerAndResetPriv: GetBoolPointer(directoryGroup["virtual_power_and_reset_priv"].(bool)),
						ILOConfigPriv:            GetBoolPointer(directoryGroup["ilo_config_priv"].(bool)),
					})
				}

				// extracting local accounts
				rawLocalAccounts := mpSetting["local_accounts"].(*schema.Set).List()
				ovLocalAccounts := make([]ov.LocalAccounts, 0)
				for _, localAccounts := range rawLocalAccounts {
					localAccount := localAccounts.(map[string]interface{})
					ovLocalAccounts = append(ovLocalAccounts, ov.LocalAccounts{
						UserName:                 localAccount["user_name"].(string),
						DisplayName:              localAccount["display_name"].(string),
						Password:                 localAccount["password"].(string),
						UserConfigPriv:           GetBoolPointer(localAccount["user_config_priv"].(bool)),
						RemoteConsolePriv:        GetBoolPointer(localAccount["remote_console_priv"].(bool)),
						VirtualMediaPriv:         GetBoolPointer(localAccount["virtual_media_priv"].(bool)),
						VirtualPowerAndResetPriv: GetBoolPointer(localAccount["virtual_power_and_reset_priv"].(bool)),
						ILOConfigPriv:            GetBoolPointer(localAccount["ilo_config_priv"].(bool)),
						LoginPriv:                GetBoolPointer(localAccount["login_priv"].(bool)),
						HostBIOSConfigPriv:       GetBoolPointer(localAccount["host_bios_config_priv"].(bool)),
						HostNICConfigPriv:        GetBoolPointer(localAccount["host_nic_config_priv"].(bool)),
						HostStorageConfigPriv:    GetBoolPointer(localAccount["host_storage_config_priv"].(bool)),
					})
				}

				// setting MpSettings
				ovMpSettings = ov.MpSettings{
					AdministratorAccount: ovAdminAcc,
					LocalAccounts:        ovLocalAccounts,
					Directory:            ovDirectory,
					DirectoryGroups:      ovDirectoryGroups,
					KeyManager:           ovKeyManager,
					IloHostName:          ovHostName,
				}
			}
			// setting ManagementProcessor
			ovManagementProcessor = ov.ManagementProcessors{
				ManageMp:  mp["manage_mp"].(bool),
				MpSetting: ovMpSettings,
			}
		}
		serverProfileTemplate.ManagementProcessors = ovManagementProcessor
	}

	if val, ok := d.GetOk("connection_settings"); ok {
		connections := val.([]interface{})
		for _, rawConSettings := range connections {
			rawConSetting := rawConSettings.(map[string]interface{})
			rawNetwork := rawConSetting["connections"].(*schema.Set).List()
			networks := make([]ov.Connection, 0)
			for i, rawNet := range rawNetwork {
				rawNetworkItem := rawNet.(map[string]interface{})
				bootOptions := ov.BootOption{}
				if rawNetworkItem["boot"] != nil {
					rawBoots := rawNetworkItem["boot"].([]interface{})
					for _, rawBoot := range rawBoots {
						bootItem := rawBoot.(map[string]interface{})
						bootTargets := []ov.BootTarget{}
						rawBootTargets := bootItem["boot_target"].(*schema.Set).List()
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
									Chaplevel:            rawIscsiItem["chap_level"].(string),
									FirstBootTargetIp:    rawIscsiItem["first_boot_target_ip"].(string),
									FirstBootTargetPort:  rawIscsiItem["first_boot_target_port"].(string),
									SecondBootTargetIp:   rawIscsiItem["second_boot_target_ip"].(string),
									SecondBootTargetPort: rawIscsiItem["second_boot_target_port"].(string),
								}
							}
						}
						val, err := strconv.Atoi(bootItem["boot_vlan_id"].(string))
						if err != nil {
							return fmt.Errorf("invalid boot_vlan_id: %s", err)
						}
						bootOptions = ov.BootOption{
							Priority:         bootItem["priority"].(string),
							EthernetBootType: bootItem["ethernet_boot_type"].(string),
							BootVolumeSource: bootItem["boot_volume_source"].(string),
							Iscsi:            &iscsi,
							Targets:          bootTargets,
							BootVlanId:       val,
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
							IpAddressSource: rawIpv4Item["ip_address_source"].(string),
						}
					}
				}

				networkv200 := ov.Connectionv200{
					RequestedVFs: rawNetworkItem["requested_vfs"].(string),
				}

				networks = append(networks, ov.Connection{
					ID:             rawNetworkItem["id"].(int),
					Name:           rawNetworkItem["name"].(string),
					IsolatedTrunk:  rawNetworkItem["isolated_trunk"].(bool),
					LagName:        rawNetworkItem["lag_name"].(string),
					FunctionType:   rawNetworkItem["function_type"].(string),
					NetworkURI:     utils.NewNstring(rawNetworkItem["network_uri"].(string)),
					PortID:         rawNetworkItem["port_id"].(string),
					Connectionv200: networkv200,
					RequestedMbps:  rawNetworkItem["requested_mbps"].(string),
				})

				if !(reflect.DeepEqual(bootOptions, ov.BootOption{})) {
					networks[i].Boot = &bootOptions

				}
				if ipv4 != (ov.Ipv4Option{}) {
					networks[i].Ipv4 = &ipv4
				}

			}
			serverProfileTemplate.ConnectionSettings = ov.ConnectionSettings{
				Connections:       networks,
				ComplianceControl: rawConSetting["compliance_control"].(string),
				ManageConnections: rawConSetting["manage_connections"].(bool),
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
			serverProfileTemplate.Boot = ov.BootManagement{
				ManageBoot: rawBoot["manage_boot"].(bool),
				Order:      bootOrder,
			}
			if _, ok := d.GetOk("boot_mode"); ok {
				rawBootMode := d.Get("boot_mode").([]interface{})[0].(map[string]interface{})
				manageMode := rawBootMode["manage_mode"].(bool)
				serverProfileTemplate.BootMode = ov.BootModeOption{
					ManageMode:    &manageMode,
					Mode:          rawBootMode["mode"].(string),
					PXEBootPolicy: utils.Nstring(rawBootMode["pxe_boot_policy"].(string)),
					SecureBoot:    rawBootMode["secure_boot"].(string),
				}
			}
		}
	}

	if val, ok := d.GetOk("bios_option"); ok {
		rawBiosOption := val.([]interface{})
		biosOption := ov.BiosOption{}
		for _, raw := range rawBiosOption {
			rawBiosItem := raw.(map[string]interface{})

			overriddenSettings := make([]ov.BiosSettings, 0)
			rawOverriddenSetting := rawBiosItem["overridden_settings"].(*schema.Set).List()

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
		serverProfileTemplate.Bios = &biosOption
	}

	if val, ok := d.GetOk("initial_scope_uris"); ok {
		initialScopeUrisOrder := val.(*schema.Set).List()
		initialScopeUris := make([]utils.Nstring, len(initialScopeUrisOrder))
		for i, raw := range initialScopeUrisOrder {
			initialScopeUris[i] = utils.Nstring(raw.(string))
		}
		serverProfileTemplate.InitialScopeUris = initialScopeUris
	}

	if _, ok := d.GetOk("firmware"); ok {
		rawFirmware := d.Get("firmware").([]interface{})
		firmware := ov.FirmwareOption{}
		for _, raw := range rawFirmware {
			firmwareItem := raw.(map[string]interface{})
			firmware = ov.FirmwareOption{
				ComplianceControl:      firmwareItem["compliance_control"].(string),
				ForceInstallFirmware:   firmwareItem["force_install_firmware"].(bool),
				FirmwareBaselineUri:    utils.NewNstring(firmwareItem["firmware_baseline_uri"].(string)),
				ManageFirmware:         firmwareItem["manage_firmware"].(bool),
				FirmwareInstallType:    firmwareItem["firmware_install_type"].(string),
				FirmwareActivationType: firmwareItem["firmware_activation_type"].(string),
			}
		}
		serverProfileTemplate.Firmware = firmware
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
					ld := ov.LogicalDriveV3{
						Bootable:          &boot,
						RaidLevel:         logicalDrivesItem["raid_level"].(string),
						Accelerator:       logicalDrivesItem["accelerator"].(string),
						DriveTechnology:   logicalDrivesItem["drive_technology"].(string),
						Name:              logicalDrivesItem["name"].(string),
						NumPhysicalDrives: logicalDrivesItem["num_physical_drives"].(int),
					}
					if val := logicalDrivesItem["sas_logical_jbod_id"].(string); val != "" {
						val, _ := strconv.Atoi(logicalDrivesItem["sas_logical_jbod_id"].(string))
						ld.SasLogicalJBODId = val
					}
					if val := logicalDrivesItem["num_spare_drives"].(string); val != "" {
						val, _ := strconv.Atoi(logicalDrivesItem["num_spare_drives"].(string))
						ld.NumSpareDrives = val
					}

					logicalDrives = append(logicalDrives, ld)

				}
				init, _ := controllerData["initialize"].(bool)
				localStorageEmbeddedController = append(localStorageEmbeddedController, ov.LocalStorageEmbeddedController{
					DeviceSlot:             controllerData["device_slot"].(string),
					DriveWriteCache:        controllerData["drive_write_cache"].(string),
					Initialize:             &init,
					Mode:                   controllerData["mode"].(string),
					PredictiveSpareRebuild: controllerData["predictive_spare_rebuild"].(string),
					LogicalDrives:          logicalDrives,
				})
			}
			// Gets Local Storage Sas Jbods Body
			rawLocalStorageSasJbod := localStorageItem["sas_logical_jbod"].(*schema.Set).List()
			logicalJbod := make([]ov.LogicalJbod, 0)
			for _, raw3 := range rawLocalStorageSasJbod {
				sasLogicalJbodData := raw3.(map[string]interface{})
				logicalJbod = append(logicalJbod, ov.LogicalJbod{
					Description:       sasLogicalJbodData["description"].(string),
					DeviceSlot:        sasLogicalJbodData["drive_slot"].(string),
					DriveMaxSizeGB:    sasLogicalJbodData["drive_max_size_gb"].(int),
					DriveMinSizeGB:    sasLogicalJbodData["drive_min_size_gb"].(int),
					DriveTechnology:   sasLogicalJbodData["drive_technology"].(string),
					EraseData:         sasLogicalJbodData["erase_data"].(bool),
					ID:                sasLogicalJbodData["id"].(int),
					Name:              sasLogicalJbodData["name"].(string),
					NumPhysicalDrives: sasLogicalJbodData["num_physical_drive"].(int),
					Persistent:        sasLogicalJbodData["persistent"].(bool),
				})
			}
			localStorage = ov.LocalStorageOptions{
				ComplianceControl: localStorageItem["compliance_control"].(string),
				Controllers:       localStorageEmbeddedController,
				SasLogicalJBODs:   logicalJbod,
			}
		}
		serverProfileTemplate.LocalStorage = localStorage
	}

	// get SAN storage data if provided
	if _, ok := d.GetOk("san_storage"); ok {
		rawSanStorage := d.Get("san_storage").([]interface{})
		sanStorage := ov.SanStorageOptions{}
		for _, raw := range rawSanStorage {
			sanStorageItem := raw.(map[string]interface{})
			sansystemcred := make([]ov.SanSystemCredential, 0)
			if sanStorageItem["san_system_credentials"] != "" {
				rawsansystemCredentials := sanStorageItem["san_system_credentials"].([]interface{})

				for _, raw3 := range rawsansystemCredentials {
					sansystemcreddata := raw3.(map[string]interface{})
					sansystemcred = append(sansystemcred, ov.SanSystemCredential{
						ChapLevel:        sansystemcreddata["chap_level"].(string),
						ChapName:         sansystemcreddata["chap_name"].(string),
						ChapSecret:       sansystemcreddata["chap_secret"].(string),
						ChapSource:       sansystemcreddata["chap_source"].(string),
						MutualChapName:   sansystemcreddata["mutual_chap_name"].(string),
						MutualChapSecret: sansystemcreddata["mutual_chap_secret"].(string),
						StorageSystemUri: utils.NewNstring(sansystemcreddata["storage_system_uri"].(string)),
					})
				}
			}

			sanStorage = ov.SanStorageOptions{
				ComplianceControl:    sanStorageItem["compliance_control"].(string),
				HostOSType:           sanStorageItem["host_os_type"].(string),
				ManageSanStorage:     sanStorageItem["manage_san_storage"].(bool),
				SanSystemCredentials: sansystemcred,
			}

		}

		serverProfileTemplate.SanStorage = sanStorage
	}

	// Get volume attachment data for san storage
	if _, ok := d.GetOk("volume_attachments"); ok {

		rawVolumeAttachments := d.Get("volume_attachments").([]interface{})
		volumeAttachments := make([]ov.VolumeAttachment, 0)

		for i, rawVolumeAttachment := range rawVolumeAttachments {
			volumeAttachmentItem := rawVolumeAttachment.(map[string]interface{})
			volumes := ov.Volume{}

			if volumeAttachmentItem["volume"] != nil {
				rawVolume := volumeAttachmentItem["volume"].([]interface{})
				for _, rawVol := range rawVolume {
					volumeItem := rawVol.(map[string]interface{})
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
					rawInitialScopeUris := volumeItem["initial_scope_uris"].(*schema.Set).List()
					initialScopeUris := make([]utils.Nstring, len(rawInitialScopeUris))
					for i, raw := range rawInitialScopeUris {
						initialScopeUris[i] = utils.Nstring(raw.(string))
					}
					volumes.InitialScopeUris = initialScopeUris
					volumes = ov.Volume{
						TemplateUri: utils.NewNstring(volumeItem["template_uri"].(string)),
					}
					if !reflect.DeepEqual(properties, ov.PropertiesSP{}) {
						volumes.Properties = &properties
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
				IsPermanent:                    GetBoolPointer(volumeAttachmentItem["is_permanent"].(bool)),
				LUN:                            volumeAttachmentItem["lun"].(string),
				LUNType:                        volumeAttachmentItem["lun_type"].(string),
				VolumeURI:                      utils.NewNstring(volumeAttachmentItem["volume_uri"].(string)),
				VolumeStorageSystemURI:         utils.NewNstring(volumeAttachmentItem["volume_storage_system_uri"].(string)),
				AssociatedTemplateAttachmentId: volumeAttachmentItem["associated_template_attachment_id"].(string),
				State:                          volumeAttachmentItem["state"].(string),
				Status:                         volumeAttachmentItem["status"].(string),
				StoragePaths:                   storagePaths,
				BootVolumePriority:             volumeAttachmentItem["boot_volume_priority"].(string),
			})
			if !reflect.DeepEqual(volumes, ov.Volume{}) {
				volumeAttachments[i].Volume = &volumes
			}

		}
		serverProfileTemplate.SanStorage.VolumeAttachments = volumeAttachments
	}
	sptError := config.ovClient.CreateProfileTemplate(serverProfileTemplate)
	d.SetId(d.Get("name").(string))
	if sptError != nil {
		d.SetId("")
		return sptError
	}
	return resourceServerProfileTemplateRead(d, meta)
}

// flattens management processor
func flattenMp(d *schema.ResourceData) map[string]interface{} {
	if val, ok := d.GetOk("management_processor"); ok {
		valn := val.([]interface{})
		if len(valn) != 0 {
			vall := valn[0].(map[string]interface{})
			if len(vall) != 0 {
				valmp := vall["mp_settings"].([]interface{})
				if len(valmp) != 0 {
					return valmp[0].(map[string]interface{})
				}
			}
		}

	}
	return nil
}

func resourceServerProfileTemplateRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	spt, err := config.ovClient.GetProfileTemplateByName(d.Id())
	if err != nil || spt.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.Set("affinity", spt.Affinity)
	d.Set("category", spt.Category)
	d.Set("created", spt.Created)
	d.Set("description", spt.Description)
	d.Set("server_profile_description", spt.ServerProfileDescription)

	if spt.EnclosureGroupURI != "" {
		enclosureGroup, err := config.ovClient.GetEnclosureGroupByUri(spt.EnclosureGroupURI)
		if err != nil {
			return err
		}
		d.Set("enclosure_group", enclosureGroup.Name)
		d.Set("enclosure_group_uri", spt.EnclosureGroupURI)
	}
	d.Set("etag", spt.ETAG)
	d.Set("hide_unused_flex_nics", spt.HideUnusedFlexNics)
	d.Set("iscsi_initiator_name", spt.IscsiInitiatorName)
	d.Set("iscsi_initiator_name_type", spt.IscsiInitiatorNameType)
	d.Set("mac_type", spt.MACType)
	d.Set("modified", spt.Modified)
	d.Set("name", spt.Name)
	d.Set("refresh_state", spt.RefreshState)
	d.Set("scopes_uri", spt.ScopesUri)
	d.Set("serial_number_type", spt.SerialNumberType)

	serverHardwareType, err := config.ovClient.GetServerHardwareTypeByUri(spt.ServerHardwareTypeURI)
	if err != nil {
		return err
	}
	d.Set("server_hardware_type", serverHardwareType.Name)
	d.Set("server_hardware_type_uri", spt.ServerHardwareTypeURI.String())
	d.Set("state", spt.State)
	d.Set("status", spt.Status)
	d.Set("type", spt.Type)
	d.Set("uri", spt.URI.String())
	d.Set("wwn_type", spt.WWNType)

	// reads scope from SPT resource
	scopes, err := config.ovClient.GetScopeFromResource(spt.URI.String())
	if err != nil {
		log.Printf("unable to fetch scopes: %s", err)
	} else {
		d.Set("initial_scope_uris", scopes.ScopeUris)
	}

	bootOrder := make([]interface{}, 0)
	if len(spt.Boot.Order) != 0 {
		for _, currBoot := range spt.Boot.Order {
			bootOrder = append(bootOrder, currBoot)
		}
	}
	boot := make([]map[string]interface{}, 0, 1)
	boot = append(boot, map[string]interface{}{
		"manage_boot":        spt.Boot.ManageBoot,
		"compliance_control": spt.Boot.ComplianceControl,
		"boot_order":         bootOrder,
	})
	d.Set("boot", boot)
	emptyBootMode := ov.BootModeOption{}
	if spt.BootMode != emptyBootMode {
		bootMode := make([]map[string]interface{}, 0, 1)
		bootMode = append(bootMode, map[string]interface{}{
			"compliance_control": spt.BootMode.ComplianceControl,
			"manage_mode":        spt.BootMode.ManageMode,
			"mode":               spt.BootMode.Mode,
			"pxe_boot_policy":    spt.BootMode.PXEBootPolicy,
			"secure_boot":        spt.BootMode.SecureBoot,
		})
		d.Set("boot_mode", bootMode)
	}

	if spt.Bios != nil {
		biosOptions := make([]map[string]interface{}, 0, 1)
		overriddenSettings := make([]interface{}, 0)
		if len(spt.Bios.OverriddenSettings) > 0 {
			for _, overriddenSetting := range spt.Bios.OverriddenSettings {
				overriddenSettings = append(overriddenSettings, map[string]interface{}{
					"id":    overriddenSetting.ID,
					"value": overriddenSetting.Value,
				})
			}
		}
		biosOptions = append(biosOptions, map[string]interface{}{
			"manage_bios":         spt.Bios.ManageBios,
			"overridden_settings": overriddenSettings,
		})

		d.Set("bios_option", biosOptions)
	}

	emptyFirmware := ov.FirmwareOption{}
	if spt.Firmware != emptyFirmware {
		firmware := make([]map[string]interface{}, 0, 1)
		firmware = append(firmware, map[string]interface{}{
			"compliance_control":       spt.Firmware.ComplianceControl,
			"firmware_baseline_uri":    spt.Firmware.FirmwareBaselineUri,
			"force_install_firmware":   spt.Firmware.ForceInstallFirmware,
			"manage_firmware":          spt.Firmware.ManageFirmware,
			"firmware_install_type":    spt.Firmware.FirmwareInstallType,
			"firmware_activation_type": spt.Firmware.FirmwareActivationType,
		})
		d.Set("firmware", firmware)
	}

	emptyManagementProcessor := ov.IntManagementProcessor{}
	if !reflect.DeepEqual(spt.ManagementProcessor, emptyManagementProcessor) {
		mpSettings := make([]map[string]interface{}, 0)
		if len(spt.ManagementProcessor.MpSettings) != 0 {
			// initializing schema variables...
			adminAcc := make([]map[string]interface{}, 0)
			directory := make([]map[string]interface{}, 0)
			keyManager := make([]map[string]interface{}, 0)
			directoryGroups := make([]map[string]interface{}, 0)
			localAccounts := make([]map[string]interface{}, 0)
			hostName := make([]map[string]interface{}, 0)

			for _, val := range spt.ManagementProcessor.MpSettings {

				if val.SettingType == "AdministratorAccount" {
					// initializing 0th location...
					adminAc := map[string]interface{}{}
					// adding attributes if they exists...

					if daa, ok := val.Args["deleteAdministratorAccount"]; ok {
						adminAc["delete_administrator_account"] = daa
					}
					if pass, ok := val.Args["password"]; ok {
						if pass != nil {
							adminAc["password"] = pass
						}
					}
					// extracts MpSettings to re-set it
					valmpp := flattenMp(d)
					if valmpp != nil {
						vals := valmpp["administrator_account"].([]interface{})
						for _, x := range vals {
							xx := x.(map[string]interface{})
							adminAc["password"] = xx["password"]
						}
					}
					adminAcc = append(adminAcc, adminAc)

				}

				if val.SettingType == "Directory" {
					// initializing 0th location...
					directoryy := map[string]interface{}{}

					// adding attributes if they exists...
					if dgl, ok := val.Args["directoryAuthentication"]; ok {
						directoryy["directory_authentication"] = dgl
					}
					if dgl, ok := val.Args["directoryGenericLDAP"]; ok {
						directoryy["directory_generic_ldap"] = dgl
					}
					if dsa, ok := val.Args["directoryServerAddress"]; ok {
						directoryy["directory_server_address"] = dsa
					}
					if dsp, ok := val.Args["directoryServerPort"]; ok {
						directoryy["directory_server_port"] = dsp
					}
					if dsc, ok := val.Args["directoryServerCertificate"]; ok {
						directoryy["directory_server_certificate"] = dsc
					}
					if iodn, ok := val.Args["iloObjectDistinguishedName"]; ok {
						directoryy["ilo_distinguished_name"] = iodn
					}
					if p, ok := val.Args["password"]; ok {
						if p != nil {
							directoryy["password"] = p
						}
					}
					if ka, ok := val.Args["kerberosAuthentication"]; ok {
						directoryy["kerberos_authentication"] = ka
					}
					if kr, ok := val.Args["kerberosRealm"]; ok {
						directoryy["kerberos_realm"] = kr
					}
					if kksa, ok := val.Args["kerberosKDCServerAddress"]; ok {
						directoryy["kerberos_kdc_server_address"] = kksa
					}
					if kksp, ok := val.Args["kerberosKDCServerPort"]; ok {
						directoryy["kerberos_kdc_server_port"] = kksp
					}
					if kkt, ok := val.Args["kerberosKeytab"]; ok {
						directoryy["kerberos_key_tab"] = kkt
					}
					if duc, ok := val.Args["directoryUserContext"]; ok {
						ducSet := []string{}
						switch reflect.TypeOf(duc).Kind() {
						case reflect.Slice:
							s := reflect.ValueOf(duc)
							for i := 0; i < s.Len(); i++ {
								ducSet = append(ducSet, s.Index(i).Interface().(string))
							}
							directoryy["directory_user_context"] = ducSet
						}
					}
					// extracts MpSettings to re-set it
					valmpp := flattenMp(d)
					if valmpp != nil {
						vals := valmpp["directory"].([]interface{})
						for _, x := range vals {
							xx := x.(map[string]interface{})
							directoryy["password"] = xx["password"]
						}
					}
					directory = append(directory, directoryy)

				}
				if val.SettingType == "DirectoryGroups" {
					if dga, ok := val.Args["directoryGroupAccounts"]; ok {
						// dectectng type of interface value
						switch reflect.TypeOf(dga).Kind() {
						case reflect.Slice:
							s := reflect.ValueOf(dga)
							for i := 0; i < s.Len(); i++ {
								// since slice elements will be maps converting interface to map
								w := s.Index(i).Interface().(map[string]interface{})
								dg := map[string]interface{}{}
								if gd, ok := w["groupDN"]; ok {
									dg["group_dn"] = gd
								}
								if gs, ok := w["groupSID"]; ok {
									dg["group_sid"] = gs
								}
								if ucp, ok := w["userConfigPriv"]; ok {
									dg["user_config_priv"] = ucp
								}
								if rcp, ok := w["remoteConsolePriv"]; ok {
									dg["remote_console_priv"] = rcp
								}
								if vmp, ok := w["virtualMediaPriv"]; ok {
									dg["virtual_media_priv"] = vmp
								}
								if vparp, ok := w["virtualPowerAndResetPriv"]; ok {
									dg["virtual_power_and_reset_priv"] = vparp
								}
								if icp, ok := w["iLOConfigPriv"]; ok {
									dg["ilo_config_priv"] = icp
								}
								// adding directory groups
								directoryGroups = append(directoryGroups, dg)
							}
						}
					}
				}

				if val.SettingType == "LocalAccounts" {
					if las, ok := val.Args["localAccounts"]; ok {
						// dectectng type of interface value
						switch reflect.TypeOf(las).Kind() {
						case reflect.Slice:
							s := reflect.ValueOf(las)
							for i := 0; i < s.Len(); i++ {
								// since slice elements will be maps converting interface to map
								w := s.Index(i).Interface().(map[string]interface{})
								la := map[string]interface{}{}
								if un, ok := w["userName"]; ok {
									la["user_name"] = un
								}
								if dn, ok := w["displayName"]; ok {
									la["display_name"] = dn
								}
								if p, ok := w["password"]; ok {
									if p != nil {
										la["password"] = p
									}
								}
								if ucp, ok := w["userConfigPriv"]; ok {
									la["user_config_priv"] = ucp
								}
								if rcp, ok := w["remoteConsolePriv"]; ok {
									la["remote_console_priv"] = rcp
								}
								if vmp, ok := w["virtualMediaPriv"]; ok {
									la["virtual_media_priv"] = vmp
								}
								if vpar, ok := w["virtualPowerAndResetPriv"]; ok {
									la["virtual_power_and_reset_priv"] = vpar
								}
								if icp, ok := w["iLOConfigPriv"]; ok {
									la["ilo_config_priv"] = icp
								}
								if lp, ok := w["loginPriv"]; ok {
									la["login_priv"] = lp
								}
								if hbcp, ok := w["hostBIOSConfigPriv"]; ok {
									la["host_bios_config_priv"] = hbcp
								}
								if hncp, ok := w["hostNICConfigPriv"]; ok {
									la["host_nic_config_priv"] = hncp
								}
								if hscp, ok := w["hostStorageConfigPriv"]; ok {
									la["host_storage_config_priv"] = hscp
								}
								// adding local account
								localAccounts = append(localAccounts, la)
							}

							//extracts MpSettings to re-set it

							valmpp := flattenMp(d)
							if valmpp != nil {
								vals := valmpp["local_accounts"].(*schema.Set).List()
								for i, x := range vals {
									xx := x.(map[string]interface{})
									localAccounts[i]["password"] = xx["password"]
								}
							}
						}
					}
				}

				if val.SettingType == "Hostname" {
					hostname := map[string]interface{}{}
					if host, ok := val.Args["hostName"]; ok {
						hostname["ilo_host_name"] = host
					}
					hostName = append(hostName, hostname)
				}

				if val.SettingType == "KeyManager" {
					// initializing 0th location...
					keyManagerr := map[string]interface{}{}
					// extratcing values if exists...
					if psa, ok := val.Args["primaryServerAddress"]; ok {
						keyManagerr["primary_server_address"] = psa
					}
					if psp, ok := val.Args["primaryServerPort"]; ok {
						keyManagerr["primary_server_port"] = psp
					}
					if ssa, ok := val.Args["secondaryServerAddress"]; ok {
						keyManagerr["secondary_server_address"] = ssa
					}
					if ssp, ok := val.Args["secondaryServerPort"]; ok {
						keyManagerr["secondary_server_port"] = ssp
					}
					if rr, ok := val.Args["redundancyRequired"]; ok {
						keyManagerr["redundancy_required"] = rr
					}
					if gn, ok := val.Args["groupName"]; ok {
						keyManagerr["group_name"] = gn
					}
					if cn, ok := val.Args["certificateName"]; ok {
						keyManagerr["certificate_name"] = cn
					}
					if ln, ok := val.Args["loginName"]; ok {
						keyManagerr["login_name"] = ln
					}
					if p, ok := val.Args["password"]; ok {
						keyManagerr["password"] = p
					}
					// extracts MpSettings to re-set it
					valmpp := flattenMp(d)
					if valmpp != nil {
						vals := valmpp["key_manager"].([]interface{})
						for _, x := range vals {
							xx := x.(map[string]interface{})
							keyManagerr["password"] = xx["password"]
						}
					}
					keyManager = append(keyManager, keyManagerr)

				}

			}
			// setting MpSettings
			mpSettings = append(mpSettings, map[string]interface{}{
				"administrator_account": adminAcc,
				"directory":             directory,
				"key_manager":           keyManager,
				"directory_groups":      directoryGroups,
				"local_accounts":        localAccounts,
				"ilo_host_name":         hostName,
			})
		}

		mp := make([]map[string]interface{}, 0)
		mp = append(mp, map[string]interface{}{
			"compliance_control": spt.ManagementProcessor.ComplianceControl,
			"manage_mp":          spt.ManagementProcessor.ManageMp,
			"mp_settings":        mpSettings,
		})
		d.Set("management_processor", mp)
	}

	if len(spt.ConnectionSettings.Connections) != 0 {
		// Get connections
		connections := make([]map[string]interface{}, 0, len(spt.ConnectionSettings.Connections))
		for _, connection := range spt.ConnectionSettings.Connections {
			// Gets Boot for Connection
			iscsi := make([]map[string]interface{}, 0, 1)
			bootTargets := make([]map[string]interface{}, 0, 0)
			// Gets Boot Settings
			connectionBoot := make([]map[string]interface{}, 0, 1)
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
					})
				}

				connectionBoot = append(connectionBoot, map[string]interface{}{
					"priority":           connection.Boot.Priority,
					"ethernet_boot_type": connection.Boot.EthernetBootType,
					"boot_volume_source": connection.Boot.BootVolumeSource,
					"iscsi":              iscsi,
					"boot_target":        bootTargets,
				})
				if connection.Boot.BootVlanId != 0 {
					connectionBoot[0]["boot_vlan_id"] = strconv.Itoa(connection.Boot.BootVlanId)
				}
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
				"boot":           connectionBoot,
				"function_type":  connection.FunctionType,
				"id":             connection.ID,
				"ipv4":           connectionIpv4,
				"isolated_trunk": connection.IsolatedTrunk,
				"lag_name":       connection.LagName,
				"mac_type":       connection.MacType,
				"name":           connection.Name,
				"network_name":   connection.NetworkName,
				"network_uri":    connection.NetworkURI,
				"port_id":        connection.PortID,
				"requested_mbps": connection.RequestedMbps,
				"requested_vfs":  connection.RequestedVFs,
			})
		}

		// flatten connection settings to overwrite port_id if equals to "Auto"
		if getVal, ok := d.GetOk("connection_settings"); ok {
			conSetVal := getVal.([]interface{})
			for _, rawConSet := range conSetVal {
				conSet := rawConSet.(map[string]interface{})
				consVal := conSet["connections"].(*schema.Set).List()
				// iterating through connections from state
				for i, rawConVal := range consVal {
					con := rawConVal.(map[string]interface{})
					// iterating through connections from refresh
					for _, conVal := range connections {
						if conVal["id"] == con["id"] {
							// overrides port_id
							if con["port_id"] == "Auto" {
								connections[i]["port_id"] = "Auto"
							}
						}
					}
				}
			}
		}

		// Connection Settings
		connectionSettings := make([]map[string]interface{}, 0, 1)
		connectionSettings = append(connectionSettings, map[string]interface{}{
			"manage_connections": spt.ConnectionSettings.ManageConnections,
			"compliance_control": spt.ConnectionSettings.ComplianceControl,
			"connections":        connections,
		})
		d.Set("connection_settings", connectionSettings)
	}

	if len(spt.LocalStorage.Controllers) != 0 {
		// Gets Storage Controller Body
		controllers := make([]map[string]interface{}, 0, len(spt.LocalStorage.Controllers))
		for i := 0; i < len(spt.LocalStorage.Controllers); i++ {
			logicalDrives := make([]map[string]interface{}, 0, len(spt.LocalStorage.Controllers[i].LogicalDrives))
			for j := 0; j < len(spt.LocalStorage.Controllers[i].LogicalDrives); j++ {
				ld := map[string]interface{}{
					"bootable":            *spt.LocalStorage.Controllers[i].LogicalDrives[j].Bootable,
					"accelerator":         spt.LocalStorage.Controllers[i].LogicalDrives[j].Accelerator,
					"drive_technology":    spt.LocalStorage.Controllers[i].LogicalDrives[j].DriveTechnology,
					"name":                spt.LocalStorage.Controllers[i].LogicalDrives[j].Name,
					"num_physical_drives": spt.LocalStorage.Controllers[i].LogicalDrives[j].NumPhysicalDrives,
					"raid_level":          spt.LocalStorage.Controllers[i].LogicalDrives[j].RaidLevel,
				}
				if val := spt.LocalStorage.Controllers[i].LogicalDrives[j].SasLogicalJBODId; val != 0 {
					ld["sas_logical_jbod_id"] = strconv.Itoa(val)
				}

				if val := spt.LocalStorage.Controllers[i].LogicalDrives[j].NumSpareDrives; val != 0 {
					ld["num_spare_drives"] = strconv.Itoa(val)
				}
				logicalDrives = append(logicalDrives, ld)
			}
			controllers = append(controllers, map[string]interface{}{
				"device_slot":              spt.LocalStorage.Controllers[i].DeviceSlot,
				"initialize":               *spt.LocalStorage.Controllers[i].Initialize,
				"import_configuration":     spt.LocalStorage.Controllers[i].ImportConfiguration,
				"drive_write_cache":        spt.LocalStorage.Controllers[i].DriveWriteCache,
				"mode":                     spt.LocalStorage.Controllers[i].Mode,
				"predictive_spare_rebuild": spt.LocalStorage.Controllers[i].PredictiveSpareRebuild,
				"logical_drives":           logicalDrives,
			})
		}
		// Gets Sas Logical Jbod Controller Body
		sasLogDrives := make([]map[string]interface{}, 0, len(spt.LocalStorage.SasLogicalJBODs))
		for i := 0; i < len(spt.LocalStorage.SasLogicalJBODs); i++ {
			sasLogDrives = append(sasLogDrives, map[string]interface{}{
				"description":        spt.LocalStorage.SasLogicalJBODs[i].Description,
				"device_slot":        spt.LocalStorage.SasLogicalJBODs[i].DeviceSlot,
				"drive_max_size_gb":  spt.LocalStorage.SasLogicalJBODs[i].DriveMaxSizeGB,
				"drive_min_size_sb":  spt.LocalStorage.SasLogicalJBODs[i].DriveMinSizeGB,
				"drive_technology":   spt.LocalStorage.SasLogicalJBODs[i].DriveTechnology,
				"erase_data":         spt.LocalStorage.SasLogicalJBODs[i].EraseData,
				"id":                 spt.LocalStorage.SasLogicalJBODs[i].ID,
				"name":               spt.LocalStorage.SasLogicalJBODs[i].Name,
				"num_physical_drive": spt.LocalStorage.SasLogicalJBODs[i].NumPhysicalDrives,
				"persistent":         spt.LocalStorage.SasLogicalJBODs[i].Persistent,
			})
		}
		// Gets Local Storage Body
		localStorage := make([]map[string]interface{}, 0, 1)
		localStorage = append(localStorage, map[string]interface{}{
			"compliance_control": spt.LocalStorage.ComplianceControl,
			"controller":         controllers,
			"sas_logical_jbod":   sasLogDrives,
		})
		d.Set("local_storage", localStorage)
	}

	sanSystemCredentials := make([]interface{}, 0)
	if len(spt.SanStorage.SanSystemCredentials) != 0 {
		for i := 0; i < len(spt.SanStorage.SanSystemCredentials); i++ {
			sanSystemCredentials = append(sanSystemCredentials, map[string]interface{}{
				"chap_level":         spt.SanStorage.SanSystemCredentials[i].ChapLevel,
				"storage_system_uri": spt.SanStorage.SanSystemCredentials[i].StorageSystemUri.String(),
			})
		}
	}

	SanStorageOptions := make([]map[string]interface{}, 0, 1)
	SanStorageOptions = append(SanStorageOptions, map[string]interface{}{
		"compliance_control":     spt.SanStorage.ComplianceControl,
		"host_os_type":           spt.SanStorage.HostOSType,
		"manage_san_storage":     spt.SanStorage.ManageSanStorage,
		"san_system_credentials": sanSystemCredentials,
	})
	d.Set("san_storage", SanStorageOptions)
	volumeAttachments := make([]interface{}, 0)
	if len(spt.SanStorage.VolumeAttachments) != 0 {
		for i := 0; i < len(spt.SanStorage.VolumeAttachments); i++ {
			storagePaths := make([]interface{}, 0)
			if len(spt.SanStorage.VolumeAttachments[i].StoragePaths) != 0 {
				for j := 0; j < len(spt.SanStorage.VolumeAttachments[i].StoragePaths); j++ {
					targets := make([]interface{}, 0)
					if len(spt.SanStorage.VolumeAttachments[i].StoragePaths[j].Targets) != 0 {
						for k := 0; k < len(spt.SanStorage.VolumeAttachments[i].StoragePaths[j].Targets); k++ {
							targets = append(targets, map[string]interface{}{
								"ip_address": spt.SanStorage.VolumeAttachments[i].StoragePaths[j].Targets[k].IpAddress,
								"name":       spt.SanStorage.VolumeAttachments[i].StoragePaths[j].Targets[k].Name,
								"tcp_port":   spt.SanStorage.VolumeAttachments[i].StoragePaths[j].Targets[k].TcpPort,
							})
						}

					}
					storagePaths = append(storagePaths, map[string]interface{}{
						"connection_id":   spt.SanStorage.VolumeAttachments[i].StoragePaths[j].ConnectionID,
						"is_enabled":      spt.SanStorage.VolumeAttachments[i].StoragePaths[j].IsEnabled,
						"network_uri":     spt.SanStorage.VolumeAttachments[i].StoragePaths[j].NetworkUri.String(),
						"status":          spt.SanStorage.VolumeAttachments[i].StoragePaths[j].Status,
						"target_selector": spt.SanStorage.VolumeAttachments[i].StoragePaths[j].TargetSelector,
						"targets":         targets,
					})
				}

			}
			volumes := make([]interface{}, 0)
			if spt.SanStorage.VolumeAttachments[i].Volume != nil {

				properties := make([]interface{}, 0)
				if spt.SanStorage.VolumeAttachments[i].Volume.Properties != nil {

					properties = append(properties, map[string]interface{}{
						"data_protection_level":            spt.SanStorage.VolumeAttachments[i].Volume.Properties.DataProtectionLevel,
						"data_transfer_limit":              spt.SanStorage.VolumeAttachments[i].Volume.Properties.DataTransferLimit,
						"description":                      spt.SanStorage.VolumeAttachments[i].Volume.Properties.Description,
						"folder":                           spt.SanStorage.VolumeAttachments[i].Volume.Properties.Folder,
						"iops_limit":                       spt.SanStorage.VolumeAttachments[i].Volume.Properties.IopsLimit,
						"is_deduplicated":                  spt.SanStorage.VolumeAttachments[i].Volume.Properties.IsDeduplicated,
						"is_encrypted":                     spt.SanStorage.VolumeAttachments[i].Volume.Properties.IsEncrypted,
						"is_pinned":                        spt.SanStorage.VolumeAttachments[i].Volume.Properties.IsPinned,
						"is_shareable":                     spt.SanStorage.VolumeAttachments[i].Volume.Properties.IsShareable,
						"name":                             spt.SanStorage.VolumeAttachments[i].Volume.Properties.Name,
						"performance_policy":               spt.SanStorage.VolumeAttachments[i].Volume.Properties.PerformancePolicy,
						"provisioning_type":                spt.SanStorage.VolumeAttachments[i].Volume.Properties.ProvisioningType,
						"size":                             spt.SanStorage.VolumeAttachments[i].Volume.Properties.Size,
						"volume_set":                       spt.SanStorage.VolumeAttachments[i].Volume.Properties.VolumeSet,
						"is_data_reduction_enabled":        spt.SanStorage.VolumeAttachments[i].Volume.Properties.IsDataReductionEnabled,
						"is_adaptive_optimization_enabled": spt.SanStorage.VolumeAttachments[i].Volume.Properties.IsAdaptiveOptimizationEnabled,
						"is_compressed":                    spt.SanStorage.VolumeAttachments[i].Volume.Properties.IsCompressed,
						"snapshot_pool":                    spt.SanStorage.VolumeAttachments[i].Volume.Properties.SnapshotPool,
						"storage_pool":                     spt.SanStorage.VolumeAttachments[i].Volume.Properties.StoragePool,
						"template_version":                 spt.SanStorage.VolumeAttachments[i].Volume.Properties.TemplateVersion,
					})

				}
				volumes = append(volumes, map[string]interface{}{
					"initial_scope_uris": spt.SanStorage.VolumeAttachments[i].Volume.InitialScopeUris,
					"template_uri":       spt.SanStorage.VolumeAttachments[i].Volume.TemplateUri.String(),
					"properties":         properties,
				})

			}

			volumeAttachments = append(volumeAttachments, map[string]interface{}{

				"associated_template_attachment_id": spt.SanStorage.VolumeAttachments[i].AssociatedTemplateAttachmentId,
				"boot_volume_priority":              spt.SanStorage.VolumeAttachments[i].BootVolumePriority,
				"id":                                spt.SanStorage.VolumeAttachments[i].ID,
				"is_permanent":                      spt.SanStorage.VolumeAttachments[i].IsPermanent,
				"lun":                               spt.SanStorage.VolumeAttachments[i].LUN,
				"lun_type":                          spt.SanStorage.VolumeAttachments[i].LUNType,
				"state":                             spt.SanStorage.VolumeAttachments[i].State,
				"status":                            spt.SanStorage.VolumeAttachments[i].Status,
				"storage_paths":                     storagePaths,
				"volume_storage_system_uri":         spt.SanStorage.VolumeAttachments[i].VolumeStorageSystemURI,
				"volume_uri":                        spt.SanStorage.VolumeAttachments[i].VolumeURI,
				"volume":                            volumes,
			})
		}

	}

	d.Set("volume_attachments", volumeAttachments)

	return nil
}

func resourceServerProfileTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	serverProfileTemplate, err := config.ovClient.GetProfileTemplateByName(d.Id())
	if err != nil {
		return err
	}

	if d.HasChange("initial_scope_uris") {
		// updates scopes for SPT
		val := d.Get("initial_scope_uris").(*schema.Set).List()
		err := UpdateScopeUris(meta, val, serverProfileTemplate.URI.String())
		if err != nil {
			return err
		}
	}

	if d.HasChange("name") {
		val := d.Get("name")
		serverProfileTemplate.Name = val.(string)
	}

	if d.HasChange("affinity") {
		val := d.Get("affinity")
		serverProfileTemplate.Affinity = val.(string)
	}
	if d.HasChange("serial_number_type") {
		val := d.Get("serial_number_type")
		serverProfileTemplate.SerialNumberType = val.(string)
	}

	if d.HasChange("wwn_type") {
		val := d.Get("wwn_type")
		serverProfileTemplate.WWNType = val.(string)
	}
	if d.HasChange("mac_type") {
		val := d.Get("mac_type")
		serverProfileTemplate.MACType = val.(string)
	}
	if d.HasChange("hide_unused_flex_nics") {
		val := d.Get("hide_unused_flex_nics")
		serverProfileTemplate.HideUnusedFlexNics = val.(bool)
	}

	if d.HasChange("description") {
		val := d.Get("description")
		serverProfileTemplate.Description = val.(string)
	}
	if d.HasChange("server_profile_description") {
		val := d.Get("server_profile_description")
		serverProfileTemplate.ServerProfileDescription = val.(string)
	}

	if d.Get("enclosure_group") != "" {
		val := d.Get("enclosure_group")
		enclosureGroup, err := config.ovClient.GetEnclosureGroupByName(val.(string))
		if err != nil {
			return err
		}
		serverProfileTemplate.EnclosureGroupURI = enclosureGroup.URI
	}

	valsht := d.Get("server_hardware_type")
	serverHardwareType, err := config.ovClient.GetServerHardwareTypeByName(valsht.(string))
	if err != nil {
		return err
	}
	serverProfileTemplate.ServerHardwareTypeURI = serverHardwareType.URI

	if d.HasChange("management_processor") {
		val := d.Get("management_processor")
		mps := val.([]interface{})
		ovManagementProcessor := ov.ManagementProcessors{}
		for _, rawMp := range mps {
			// extracting management processor
			mp := rawMp.(map[string]interface{})
			// extracting MpSettings
			ovMpSettings := ov.MpSettings{}
			mpSettings := mp["mp_settings"].([]interface{})
			for _, mpSettingg := range mpSettings {
				mpSetting := mpSettingg.(map[string]interface{})
				// extracting administrator account
				rawAdminAcc := mpSetting["administrator_account"].([]interface{})
				ovAdminAcc := ov.AdministratorAccount{}
				for _, adminAccs := range rawAdminAcc {
					adminAcc := adminAccs.(map[string]interface{})
					ovAdminAcc = ov.AdministratorAccount{
						DeleteAdministratorAccount: GetBoolPointer(adminAcc["delete_administrator_account"].(bool)),
						Password:                   adminAcc["password"].(string),
					}
				}
				// extracting directory
				rawDirectory := mpSetting["directory"].([]interface{})
				ovDirectory := ov.Directory{}
				for _, directoryy := range rawDirectory {
					directory := directoryy.(map[string]interface{})

					directoryUserContexts := make([]string, 0)
					for _, raw := range directory["directory_user_context"].(*schema.Set).List() {
						directoryUserContexts = append(directoryUserContexts, raw.(string))
					}

					ovDirectory = ov.Directory{
						DirectoryAuthentication:    directory["directory_authentication"].(string),
						DirectoryGenericLDAP:       GetBoolPointer(directory["directory_generic_ldap"].(bool)),
						DirectoryServerAddress:     directory["directory_server_address"].(string),
						DirectoryServerPort:        directory["directory_server_port"].(int),
						DirectoryServerCertificate: directory["directory_server_certificate"].(string),
						DirectoryUserContext:       directoryUserContexts,
						IloObjectDistinguishedName: directory["ilo_distinguished_name"].(string),
						Password:                   directory["password"].(string),
						KerberosAuthentication:     GetBoolPointer(directory["kerberos_authentication"].(bool)),
						KerberosRealm:              directory["kerberos_realm"].(string),
						KerberosKDCServerAddress:   directory["kerberos_kdc_server_address"].(string),
						KerberosKDCServerPort:      directory["kerberos_kdc_server_port"].(int),
						KerberosKeytab:             directory["kerberos_key_tab"].(string),
					}
				}

				// extracting hostname
				rawHostName := mpSetting["ilo_host_name"].([]interface{})
				ovHostName := ov.IloHostName{}
				for _, hostMap := range rawHostName {
					host := hostMap.(map[string]interface{})
					ovHostName = ov.IloHostName{
						HostName: host["hostname"].(string),
					}
				}

				// extracting key manager
				rawKeyManager := mpSetting["key_manager"].([]interface{})
				ovKeyManager := ov.KeyManager{}
				for _, keyManagerr := range rawKeyManager {
					keyManager := keyManagerr.(map[string]interface{})
					ovKeyManager = ov.KeyManager{
						PrimaryServerAddress:   keyManager["primary_server_address"].(string),
						PrimaryServerPort:      keyManager["primary_server_port"].(int),
						SecondaryServerAddress: keyManager["secondary_server_address"].(string),
						SecondaryServerPort:    keyManager["secondary_server_port"].(int),
						RedundancyRequired:     GetBoolPointer(keyManager["redundancy_required"].(bool)),
						GroupName:              keyManager["group_name"].(string),
						CertificateName:        keyManager["certificate_name"].(string),
						LoginName:              keyManager["login_name"].(string),
						Password:               keyManager["password"].(string),
					}
				}

				// extracting directory groups
				rawDirectoryGroups := mpSetting["directory_groups"].(*schema.Set).List()
				ovDirectoryGroups := make([]ov.DirectoryGroups, 0)
				for _, directoryGroupp := range rawDirectoryGroups {
					directoryGroup := directoryGroupp.(map[string]interface{})
					ovDirectoryGroups = append(ovDirectoryGroups, ov.DirectoryGroups{
						GroupDN:                  directoryGroup["group_dn"].(string),
						GroupSID:                 directoryGroup["group_sid"].(string),
						UserConfigPriv:           GetBoolPointer(directoryGroup["user_config_priv"].(bool)),
						RemoteConsolePriv:        GetBoolPointer(directoryGroup["remote_console_priv"].(bool)),
						VirtualMediaPriv:         GetBoolPointer(directoryGroup["virtual_media_priv"].(bool)),
						VirtualPowerAndResetPriv: GetBoolPointer(directoryGroup["virtual_power_and_reset_priv"].(bool)),
						ILOConfigPriv:            GetBoolPointer(directoryGroup["ilo_config_priv"].(bool)),
					})
				}

				// extracting local accounts
				rawLocalAccounts := mpSetting["local_accounts"].(*schema.Set).List()
				ovLocalAccounts := make([]ov.LocalAccounts, 0)
				for _, localAccounts := range rawLocalAccounts {
					localAccount := localAccounts.(map[string]interface{})
					ovLocalAccounts = append(ovLocalAccounts, ov.LocalAccounts{
						UserName:                 localAccount["user_name"].(string),
						DisplayName:              localAccount["display_name"].(string),
						Password:                 localAccount["password"].(string),
						UserConfigPriv:           GetBoolPointer(localAccount["user_config_priv"].(bool)),
						RemoteConsolePriv:        GetBoolPointer(localAccount["remote_console_priv"].(bool)),
						VirtualMediaPriv:         GetBoolPointer(localAccount["virtual_media_priv"].(bool)),
						VirtualPowerAndResetPriv: GetBoolPointer(localAccount["virtual_power_and_reset_priv"].(bool)),
						ILOConfigPriv:            GetBoolPointer(localAccount["ilo_config_priv"].(bool)),
						LoginPriv:                GetBoolPointer(localAccount["login_priv"].(bool)),
						HostBIOSConfigPriv:       GetBoolPointer(localAccount["host_bios_config_priv"].(bool)),
						HostNICConfigPriv:        GetBoolPointer(localAccount["host_nic_config_priv"].(bool)),
						HostStorageConfigPriv:    GetBoolPointer(localAccount["host_storage_config_priv"].(bool)),
					})
				}

				ovMpSettings = ov.MpSettings{
					AdministratorAccount: ovAdminAcc,
					LocalAccounts:        ovLocalAccounts,
					Directory:            ovDirectory,
					DirectoryGroups:      ovDirectoryGroups,
					KeyManager:           ovKeyManager,
					IloHostName:          ovHostName,
				}
			}
			ovManagementProcessor = ov.ManagementProcessors{
				ManageMp:  mp["manage_mp"].(bool),
				MpSetting: ovMpSettings,
			}

		}
		serverProfileTemplate.ManagementProcessors = ovManagementProcessor
	}

	if d.HasChange("connection_settings") {
		val := d.Get("connection_settings")

		connections := val.([]interface{})
		for _, rawConSettings := range connections {
			rawConSetting := rawConSettings.(map[string]interface{})
			rawNetwork := rawConSetting["connections"].(*schema.Set).List()
			networks := make([]ov.Connection, 0)
			for _, rawNet := range rawNetwork {
				rawNetworkItem := rawNet.(map[string]interface{})
				bootOptions := ov.BootOption{}
				if rawNetworkItem["boot"] != nil {
					rawBoots := rawNetworkItem["boot"].([]interface{})
					for _, rawBoot := range rawBoots {
						bootItem := rawBoot.(map[string]interface{})

						bootTargets := []ov.BootTarget{}
						rawBootTargets := bootItem["boot_target"].(*schema.Set).List()
						if rawBootTargets != nil {
							for _, rawBootTarget := range rawBootTargets {
								bootTarget := rawBootTarget.(map[string]interface{})
								bootTargett := ov.BootTarget{
									LUN:       bootTarget["lun"].(string),
									ArrayWWPN: bootTarget["array_wwpn"].(string),
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
									Chaplevel:            rawIscsiItem["chap_level"].(string),
									FirstBootTargetIp:    rawIscsiItem["first_boot_target_ip"].(string),
									FirstBootTargetPort:  rawIscsiItem["first_boot_target_port"].(string),
									SecondBootTargetIp:   rawIscsiItem["second_boot_target_ip"].(string),
									SecondBootTargetPort: rawIscsiItem["second_boot_target_port"].(string),
								}
							}
						}
						val, err := strconv.Atoi(bootItem["boot_vlan_id"].(string))
						if err != nil {
							return fmt.Errorf("invalid boot_vlan_id: %s", err)
						}
						bootOptions = ov.BootOption{
							Priority:         bootItem["priority"].(string),
							BootVolumeSource: bootItem["boot_volume_source"].(string),
							EthernetBootType: bootItem["ethernet_boot_type"].(string),
							Iscsi:            &iscsi,
							Targets:          bootTargets,
							BootVlanId:       val,
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
							IpAddressSource: rawIpv4Item["ip_address_source"].(string),
						}
					}
				}

				network := ov.Connection{
					ID:            rawNetworkItem["id"].(int),
					Name:          rawNetworkItem["name"].(string),
					FunctionType:  rawNetworkItem["function_type"].(string),
					NetworkURI:    utils.NewNstring(rawNetworkItem["network_uri"].(string)),
					PortID:        rawNetworkItem["port_id"].(string),
					RequestedMbps: rawNetworkItem["requested_mbps"].(string),
					IsolatedTrunk: rawNetworkItem["isolated_trunk"].(bool),
					LagName:       rawNetworkItem["lag_name"].(string),
				}

				networkv200 := ov.Connectionv200{
					RequestedVFs: rawNetworkItem["requested_vfs"].(string),
				}
				network.Connectionv200 = networkv200
				if ipv4 != (ov.Ipv4Option{}) {
					network.Ipv4 = &ipv4
				}
				if !(reflect.DeepEqual(bootOptions, ov.BootOption{})) {
					network.Boot = &bootOptions

				}
				networks = append(networks, network)

			}

			serverProfileTemplate.ConnectionSettings = ov.ConnectionSettings{
				Connections:       networks,
				ComplianceControl: rawConSetting["compliance_control"].(string),
				ManageConnections: rawConSetting["manage_connections"].(bool),
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
			serverProfileTemplate.Boot = ov.BootManagement{
				ManageBoot: rawBoot["manage_boot"].(bool),
				Order:      bootOrder,
			}

			rawBootMode := d.Get("boot_mode").([]interface{})[0].(map[string]interface{})
			manageMode := rawBootMode["manage_mode"].(bool)
			serverProfileTemplate.BootMode = ov.BootModeOption{
				ManageMode:    &manageMode,
				Mode:          rawBootMode["mode"].(string),
				PXEBootPolicy: utils.Nstring(rawBootMode["pxe_boot_policy"].(string)),
			}

		}
	}

	if d.HasChange("bios_option") {
		val := d.Get("bios_option")
		rawBiosOption := val.([]interface{})
		biosOption := ov.BiosOption{}
		for _, raw := range rawBiosOption {
			rawBiosItem := raw.(map[string]interface{})
			rawOverriddenSetting := rawBiosItem["overridden_settings"].(*schema.Set).List()
			overriddenSettings := make([]ov.BiosSettings, 0)

			for _, raw2 := range rawOverriddenSetting {
				rawOverriddenSettingItem := raw2.(map[string]interface{})

				overriddenSetting := ov.BiosSettings{
					ID:    rawOverriddenSettingItem["id"].(string),
					Value: rawOverriddenSettingItem["value"].(string),
				}
				overriddenSettings = append(overriddenSettings, overriddenSetting)

			}
			biosOption.OverriddenSettings = overriddenSettings
			manageBios := rawBiosItem["manage_bios"].(bool)
			biosOption.ManageBios = &manageBios
		}
		serverProfileTemplate.Bios = &biosOption
	}

	if d.HasChange("initial_scope_uris") {
		return fmt.Errorf("Initial scope uri can not be updated")
	}

	// Get firmware details
	if d.HasChange("firmware") {
		rawFirmware := d.Get("firmware").([]interface{})
		firmware := ov.FirmwareOption{}
		for _, raw := range rawFirmware {
			firmwareItem := raw.(map[string]interface{})
			firmware = ov.FirmwareOption{
				ComplianceControl:      firmwareItem["compliance_control"].(string),
				ForceInstallFirmware:   firmwareItem["force_install_firmware"].(bool),
				FirmwareBaselineUri:    utils.NewNstring(firmwareItem["firmware_baseline_uri"].(string)),
				ManageFirmware:         firmwareItem["manage_firmware"].(bool),
				FirmwareInstallType:    firmwareItem["firmware_install_type"].(string),
				FirmwareActivationType: firmwareItem["firmware_activation_type"].(string),
			}
		}
		serverProfileTemplate.Firmware = firmware
	}

	// Get local storage data if provided
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
						DriveTechnology:   logicalDrivesItem["drive_technology"].(string),
						Name:              logicalDrivesItem["name"].(string),
						NumPhysicalDrives: logicalDrivesItem["num_physical_drives"].(int),
					}

					if logicalDrivesItem["sas_logical_jbod_id"].(string) != "" {
						val, err := strconv.Atoi(logicalDrivesItem["sas_logical_jbod_id"].(string))
						if err != nil {
							return fmt.Errorf("invalid sas_logical_jbod_id: %s", err)
						}
						l := len(logicalDrives) - 1
						logicalDrives[l].SasLogicalJBODId = val
					}
					if val := logicalDrivesItem["num_spare_drives"].(string); val != "" {
						val, _ := strconv.Atoi(logicalDrivesItem["num_spare_drives"].(string))
						logicalDrive.NumSpareDrives = val
					}

					logicalDrives = append(logicalDrives, logicalDrive)

				}
				init, _ := controllerData["initialize"].(bool)
				localStorageEmbeddedController := ov.LocalStorageEmbeddedController{
					DeviceSlot:             controllerData["device_slot"].(string),
					DriveWriteCache:        controllerData["drive_write_cache"].(string),
					Initialize:             &init,
					ImportConfiguration:    controllerData["import_configuration"].(bool),
					Mode:                   controllerData["mode"].(string),
					PredictiveSpareRebuild: controllerData["predictive_spare_rebuild"].(string),
					LogicalDrives:          logicalDrives,
				}
				localStorageEmbeddedControllers = append(localStorageEmbeddedControllers, localStorageEmbeddedController)
			}
			rawLocalStorageSasJbod := localStorageItem["sas_logical_jbod"].(*schema.Set).List()
			logicalJbods := make([]ov.LogicalJbod, 0)
			for _, raw3 := range rawLocalStorageSasJbod {
				sasLogicalJbodData := raw3.(map[string]interface{})
				logicalJbod := ov.LogicalJbod{
					Description:       sasLogicalJbodData["description"].(string),
					DeviceSlot:        sasLogicalJbodData["drive_slot"].(string),
					DriveMaxSizeGB:    sasLogicalJbodData["drive_max_size_gb"].(int),
					DriveMinSizeGB:    sasLogicalJbodData["drive_min_size_gb"].(int),
					DriveTechnology:   sasLogicalJbodData["drive_technology"].(string),
					EraseData:         sasLogicalJbodData["erase_data"].(bool),
					ID:                sasLogicalJbodData["id"].(int),
					Name:              sasLogicalJbodData["name"].(string),
					NumPhysicalDrives: sasLogicalJbodData["num_physical_drive"].(int),
					Persistent:        sasLogicalJbodData["persistent"].(bool),
				}

				logicalJbods = append(logicalJbods, logicalJbod)

			}
			localStorage = ov.LocalStorageOptions{
				ComplianceControl: localStorageItem["compliance_control"].(string),
				Controllers:       localStorageEmbeddedControllers,
				SasLogicalJBODs:   logicalJbods,
			}
		}
		serverProfileTemplate.LocalStorage = localStorage
	}

	// get SAN storage data if provided
	if d.HasChange("san_storage") {
		rawSanStorage := d.Get("san_storage").([]interface{})
		sanStorage := ov.SanStorageOptions{}
		for _, raw := range rawSanStorage {
			sanStorageItem := raw.(map[string]interface{})
			sanStorage = ov.SanStorageOptions{
				HostOSType:       sanStorageItem["host_os_type"].(string),
				ManageSanStorage: sanStorageItem["manage_san_storage"].(bool),
			}
		}
		serverProfileTemplate.SanStorage = sanStorage
	}

	// Get volume attachment data for san storage
	if d.HasChange("volume_attachments") {
		rawVolumeAttachments := d.Get("volume_attachments").([]interface{})
		volumeAttachments := make([]ov.VolumeAttachment, 0)
		for _, rawVolumeAttachment := range rawVolumeAttachments {
			volumeAttachmentItem := rawVolumeAttachment.(map[string]interface{})
			volumes := ov.Volume{}
			if volumeAttachmentItem["volume"] != nil {
				rawVolume := volumeAttachmentItem["volume"].([]interface{})
				for _, rawVol := range rawVolume {
					volumeItem := rawVol.(map[string]interface{})
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
					rawInitialScopeUris := volumeItem["initial_scope_uris"].(*schema.Set).List()
					initialScopeUris := make([]utils.Nstring, len(rawInitialScopeUris))
					for i, raw := range rawInitialScopeUris {
						initialScopeUris[i] = utils.Nstring(raw.(string))
					}
					volumes.InitialScopeUris = initialScopeUris

					volumes = ov.Volume{
						TemplateUri: utils.NewNstring(volumeItem["template_uri"].(string)),
					}
					if !reflect.DeepEqual(properties, ov.PropertiesSP{}) {
						volumes.Properties = &properties
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
						ConnectionID:   storagePathItem["connection_id"].(int),
						IsEnabled:      storagePathItem["is_enabled"].(bool),
						NetworkUri:     utils.NewNstring(storagePathItem["network_uri"].(string)),
						Status:         storagePathItem["status"].(string),
						Targets:        targets,
						TargetSelector: storagePathItem["target_selector"].(string),
					}
					storagePaths = append(storagePaths, storagePath)
				}
			}
			volumeAttachment := ov.VolumeAttachment{
				ID:                             volumeAttachmentItem["id"].(int),
				IsPermanent:                    GetBoolPointer(volumeAttachmentItem["is_permanent"].(bool)),
				LUN:                            volumeAttachmentItem["lun"].(string),
				LUNType:                        volumeAttachmentItem["lun_type"].(string),
				VolumeURI:                      utils.NewNstring(volumeAttachmentItem["volume_uri"].(string)),
				VolumeStorageSystemURI:         utils.NewNstring(volumeAttachmentItem["volume_storage_system_uri"].(string)),
				AssociatedTemplateAttachmentId: volumeAttachmentItem["associated_template_attachment_id"].(string),
				State:                          volumeAttachmentItem["state"].(string),
				Status:                         volumeAttachmentItem["status"].(string),
				StoragePaths:                   storagePaths,
				BootVolumePriority:             volumeAttachmentItem["boot_volume_priority"].(string),
			}
			if !reflect.DeepEqual(volumes, ov.Volume{}) {
				volumeAttachment.Volume = &volumes
			}
			volumeAttachments = append(volumeAttachments, volumeAttachment)

		}
		serverProfileTemplate.SanStorage.VolumeAttachments = volumeAttachments
	}

	errC := config.ovClient.UpdateProfileTemplate(serverProfileTemplate)
	if errC != nil {
		return errC
	}
	d.SetId(d.Get("name").(string))

	return resourceServerProfileTemplateRead(d, meta)
}

func resourceServerProfileTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteProfileTemplate(d.Get("name").(string))
	if err != nil {
		return err
	}

	return nil
}
