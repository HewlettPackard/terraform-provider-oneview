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

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceServerProfileTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServerProfileTemplateRead,

		Schema: map[string]*schema.Schema{
			"affinity": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"boot": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compliance_control": {
							Type:     schema.TypeString,
							Optional: true,
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
			"enclosure_group_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"boot_mode": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compliance_control": {
							Type:     schema.TypeString,
							Optional: true,
						},
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
					},
				},
			},
			"bios_option": {
				Type:     schema.TypeSet,
				Optional: true,
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
				MaxItems: 1,
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
									},
									"lag_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"managed": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Required: true,
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
								},
							},
						},
					},
				},
			},
			"firmware": {
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compliance_control": {
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
						"force_install_firmware": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"manage_firmware": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"hide_unused_flex_nics": {
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
			"iscsi_initiator_name_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"local_storage": {
				Optional: true,
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compliance_control": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"controller": {
							Optional: true,
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
						"manage_local_storage": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"initialize": {
							Type:     schema.TypeBool,
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
								},
							},
						},
					},
				},
			},
			"mac_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_processor": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compliance_control": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"manage_mp": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"reapply_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mp_settings": {
							Computed: true,
							Type:     schema.TypeList,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"administrator_account": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"delete_administrator_account": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"password": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"directory": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"directory_authentication": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"directory_generic_ldap": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"directory_server_address": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"directory_server_port": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"directory_server_certificate": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"directory_user_context": {
													Type:     schema.TypeSet,
													Computed: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
													Set:      schema.HashString,
												},
												"ilo_distinguished_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"password": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"kerberos_authentication": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"kerberos_realm": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"kerberos_kdc_server_address": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"kerberos_kdc_server_port": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"kerberos_key_tab": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"key_manager": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"primary_server_address": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"primary_server_port": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"secondary_server_address": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"secondary_server_port": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"redundancy_required": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"group_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"certificate_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"login_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"password": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"directory_groups": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"group_dn": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"group_sid": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"user_config_priv": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"remote_console_priv": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"virtual_media_priv": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"virtual_power_and_reset_priv": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"ilo_config_priv": {
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"local_accounts": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"user_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"password": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"user_config_priv": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"remote_console_priv": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"virtual_media_priv": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"virtual_power_and_reset_priv": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"ilo_config_priv": {
													Type:     schema.TypeBool,
													Computed: true,
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
										Computed: true,
									},
									"chap_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"chap_secret": {
										Type:     schema.TypeString,
										Optional: true,
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
										Optional: true,
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
							Optional: true,
						},
						"lun_type": {
							Type:     schema.TypeString,
							Required: true,
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
						"volume": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"initial_scope_uris": {
										Computed: true,
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
				Computed: true,
			},
			"server_hardware_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_hardware_type_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_profile_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enclosure_group": {
				Type:     schema.TypeString,
				Computed: true,
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
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"wwn_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceServerProfileTemplateRead(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*Config)
	name := d.Get("name").(string)

	spt, err := config.ovClient.GetProfileTemplateByName(name)
	if err != nil || spt.URI.IsNil() {
		return fmt.Errorf("error fetching server profile template: %w", err)
	}

	d.Set("affinity", spt.Affinity)
	d.Set("category", spt.Category)
	d.Set("created", spt.Created)
	d.Set("description", spt.Description)
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

	// reads scope from SPT resource
	scopes, err := config.ovClient.GetScopeFromResource(spt.URI.String())
	if err != nil {
		log.Printf("unable to fetch scopes: %s", err)
	} else {
		d.Set("initial_scope_uris", scopes.ScopeUris)
	}

	d.Set("iscsi_initiator_name_type", spt.IscsiInitiatorNameType)
	d.Set("mac_type", spt.MACType)
	d.Set("modified", spt.Modified)
	d.SetId(name)
	d.Set("name", spt.Name)
	d.Set("refresh_state", spt.RefreshState)
	d.Set("scopes_uri", spt.ScopesUri)
	d.Set("serial_number_type", spt.SerialNumberType)
	serverHardwareType, err := config.ovClient.GetServerHardwareTypeByUri(spt.ServerHardwareTypeURI)
	if err != nil {
		return err
	}
	d.Set("server_hardware_type", serverHardwareType.Name)
	d.Set("server_hardware_type_uri", spt.ServerHardwareTypeURI)
	d.Set("server_profile_description", spt.ServerProfileDescription)
	d.Set("state", spt.State)
	d.Set("status", spt.Status)
	d.Set("type", spt.Type)
	d.Set("uri", spt.URI.String())

	d.Set("wwn_type", spt.WWNType)

	if len(spt.ConnectionSettings.Connections) != 0 {
		// Get connections
		connections := make([]map[string]interface{}, 0, len(spt.ConnectionSettings.Connections))
		for i := 0; i < len(spt.ConnectionSettings.Connections); i++ {
			// Gets Boot for Connection
			iscsi := make([]map[string]interface{}, 0, 1)
			targets := make([]map[string]interface{}, 0)
			if spt.ConnectionSettings.Connections[i].Boot != nil {
				if spt.ConnectionSettings.Connections[i].Boot.Iscsi != nil {
					iscsi = append(iscsi, map[string]interface{}{
						"chap_level":              spt.ConnectionSettings.Connections[i].Boot.Iscsi.Chaplevel,
						"initiator_name_source":   spt.ConnectionSettings.Connections[i].Boot.Iscsi.InitiatorNameSource,
						"first_boot_target_ip":    spt.ConnectionSettings.Connections[i].Boot.Iscsi.FirstBootTargetIp,
						"first_boot_target_port":  spt.ConnectionSettings.Connections[i].Boot.Iscsi.FirstBootTargetPort,
						"second_boot_target_ip":   spt.ConnectionSettings.Connections[i].Boot.Iscsi.SecondBootTargetIp,
						"second_boot_target_port": spt.ConnectionSettings.Connections[i].Boot.Iscsi.SecondBootTargetPort,
					})
				}

				// Get Boot targets list

				if len(spt.ConnectionSettings.Connections[i].Boot.Targets) != 0 {
					for j := 0; j < len(spt.ConnectionSettings.Connections[i].Boot.Targets); j++ {
						targets = append(targets, map[string]interface{}{
							"array_wwpn": spt.ConnectionSettings.Connections[i].Boot.Targets[j].ArrayWWPN,
							"lun":        spt.ConnectionSettings.Connections[i].Boot.Targets[j].LUN,
						})
					}
				}
			}

			// Gets Boot Settings
			connectionBoot := make([]map[string]interface{}, 0, 1)
			if spt.ConnectionSettings.Connections[i].Boot != nil {
				connectionBoot = append(connectionBoot, map[string]interface{}{
					"priority":           spt.ConnectionSettings.Connections[i].Boot.Priority,
					"boot_vlan_id":       spt.ConnectionSettings.Connections[i].Boot.BootVlanId,
					"ethernet_boot_type": spt.ConnectionSettings.Connections[i].Boot.EthernetBootType,
					"boot_volume_source": spt.ConnectionSettings.Connections[i].Boot.BootVolumeSource,
					"iscsi":              iscsi,
					"targets":            targets,
				})
			}
			// Get IPV4 Settings for Connection
			connectionIpv4 := make([]map[string]interface{}, 0, 1)
			if spt.ConnectionSettings.Connections[i].Ipv4 != nil {
				connectionIpv4 = append(connectionIpv4, map[string]interface{}{
					"gateway":           spt.ConnectionSettings.Connections[i].Ipv4.Gateway,
					"ip_address":        spt.ConnectionSettings.Connections[i].Ipv4.IpAddress,
					"subnet_mask":       spt.ConnectionSettings.Connections[i].Ipv4.SubnetMask,
					"ip_address_source": spt.ConnectionSettings.Connections[i].Ipv4.IpAddressSource,
				})
			}

			// Gets Connection Body
			connections = append(connections, map[string]interface{}{
				"boot":           connectionBoot,
				"function_type":  spt.ConnectionSettings.Connections[i].FunctionType,
				"id":             spt.ConnectionSettings.Connections[i].ID,
				"ipv4":           connectionIpv4,
				"isolated_trunk": spt.ConnectionSettings.Connections[i].IsolatedTrunk,
				"lag_name":       spt.ConnectionSettings.Connections[i].LagName,
				"managed":        spt.ConnectionSettings.Connections[i].Managed,
				"name":           spt.ConnectionSettings.Connections[i].Name,
				"network_name":   spt.ConnectionSettings.Connections[i].NetworkName,
				"network_uri":    spt.ConnectionSettings.Connections[i].NetworkURI,
				"port_id":        spt.ConnectionSettings.Connections[i].PortID,
				"requested_mbps": spt.ConnectionSettings.Connections[i].RequestedMbps,
			})
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

	bootOrder := make([]interface{}, 0)
	if len(spt.Boot.Order) != 0 {
		for _, currBoot := range spt.Boot.Order {
			bootOrder = append(bootOrder, currBoot)
		}
	}
	boot := make([]map[string]interface{}, 0, 1)
	boot = append(boot, map[string]interface{}{
		"compliance_control": spt.Boot.ComplianceControl,
		"manage_boot":        spt.Boot.ManageBoot,
		"boot_order":         bootOrder,
	})
	d.Set("boot", boot)

	bootModeOptions := make([]map[string]interface{}, 0, 1)
	bootModeOptions = append(bootModeOptions, map[string]interface{}{
		"manage_mode":     spt.BootMode.ManageMode,
		"mode":            spt.BootMode.Mode,
		"pxe_boot_policy": spt.BootMode.PXEBootPolicy.String(),
		"secure_boot":     spt.BootMode.SecureBoot,
	})

	d.Set("boot_mode", bootModeOptions)
	if spt.Bios != nil {
		overriddenSettings := make([]interface{}, 0, len(spt.Bios.OverriddenSettings))
		for _, overriddenSetting := range spt.Bios.OverriddenSettings {
			overriddenSettings = append(overriddenSettings, map[string]interface{}{
				"id":    overriddenSetting.ID,
				"value": overriddenSetting.Value,
			})
		}

		biosOptions := make([]map[string]interface{}, 0, 1)
		biosOptions = append(biosOptions, map[string]interface{}{
			"compliance_control":  spt.Bios.ComplianceControl,
			"manage_bios":         spt.Bios.ManageBios,
			"overridden_settings": overriddenSettings,
		})

		d.Set("bios_option", biosOptions)
	}

	firmwareOption := make([]map[string]interface{}, 0, 1)
	firmwareOption = append(firmwareOption, map[string]interface{}{
		"compliance_control":       spt.Firmware.ComplianceControl,
		"firmware_activation_type": spt.Firmware.FirmwareActivationType,
		"firmware_baseline_uri":    spt.Firmware.FirmwareBaselineUri.String(),
		"firmware_install_type":    spt.Firmware.FirmwareInstallType,
		"force_install_firmware":   spt.Firmware.ForceInstallFirmware,
		"manage_firmware":          spt.Firmware.ManageFirmware,
	})
	d.Set("firmware", firmwareOption)

	// Management Processor
	emptyManagementProcessor := ov.IntManagementProcessor{}
	if !reflect.DeepEqual(spt.ManagementProcessor, emptyManagementProcessor) {
		mpSettings := make([]map[string]interface{}, 0)
		if len(spt.ManagementProcessor.MpSettings) != 0 {
			// initializing schema variables...
			adminAcc := make([]map[string]interface{}, 1)
			directory := make([]map[string]interface{}, 1)
			keyManager := make([]map[string]interface{}, 1)
			directoryGroups := make([]map[string]interface{}, 0)
			localAccounts := make([]map[string]interface{}, 0)

			for _, val := range spt.ManagementProcessor.MpSettings {

				if val.SettingType == "AdministratorAccount" {
					// initializing 0th location...
					adminAcc[0] = map[string]interface{}{}
					// adding attributes if they exists...

					if daa, ok := val.Args["deleteAdministratorAccount"]; ok {
						adminAcc[0]["delete_administrator_account"] = daa
					}
					if pass, ok := val.Args["password"]; ok {
						if pass != nil {
							adminAcc[0]["password"] = pass
						}
					}
				}

				if val.SettingType == "Directory" {
					// initializing 0th location...
					directory[0] = map[string]interface{}{}

					// adding attributes if they exists...
					if dgl, ok := val.Args["directoryGenericLDAP"]; ok {
						directory[0]["directory_generic_ldap"] = dgl
					}
					if dsa, ok := val.Args["directoryServerAddress"]; ok {
						directory[0]["directory_server_address"] = dsa
					}
					if dsp, ok := val.Args["directoryServerPort"]; ok {
						directory[0]["directory_server_port"] = dsp
					}
					if dsc, ok := val.Args["directoryServerCertificate"]; ok {
						directory[0]["directory_server_certificate"] = dsc
					}
					if iodn, ok := val.Args["iloObjectDistinguishedName"]; ok {
						directory[0]["ilo_distinguished_name"] = iodn
					}
					if p, ok := val.Args["password"]; ok {
						if p != nil {
							directory[0]["password"] = p
						}
					}
					if ka, ok := val.Args["kerberosAuthentication"]; ok {
						directory[0]["kerberos_authentication"] = ka
					}
					if kr, ok := val.Args["kerberosRealm"]; ok {
						directory[0]["kerberos_realm"] = kr
					}
					if kksa, ok := val.Args["kerberosKDCServerAddress"]; ok {
						directory[0]["kerberos_kdc_server_address"] = kksa
					}
					if kksp, ok := val.Args["kerberosKDCServerPort"]; ok {
						directory[0]["kerberos_kdc_server_port"] = kksp
					}
					if kkt, ok := val.Args["kerberosKeytab"]; ok {
						directory[0]["kerberos_key_tab"] = kkt
					}
					if duc, ok := val.Args["directoryUserContext"]; ok {
						ducSet := []string{}
						switch reflect.TypeOf(duc).Kind() {
						case reflect.Slice:
							s := reflect.ValueOf(duc)
							for i := 0; i < s.Len(); i++ {
								ducSet = append(ducSet, s.Index(i).Interface().(string))
							}
							directory[0]["directory_user_context"] = ducSet
						}
					}
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
								// adding local account
								localAccounts = append(localAccounts, la)
							}
						}
					}
				}

				if val.SettingType == "KeyManager" {
					// initializing 0th location...
					keyManager[0] = map[string]interface{}{}
					// extratcing values if exists...
					if psa, ok := val.Args["primaryServerAddress"]; ok {
						keyManager[0]["primary_server_address"] = psa
					}
					if psp, ok := val.Args["primaryServerPort"]; ok {
						keyManager[0]["primary_server_port"] = psp
					}
					if ssa, ok := val.Args["secondaryServerAddress"]; ok {
						keyManager[0]["secondary_server_address"] = ssa
					}
					if ssp, ok := val.Args["secondaryServerPort"]; ok {
						keyManager[0]["secondary_server_port"] = ssp
					}
					if rr, ok := val.Args["redundancyRequired"]; ok {
						keyManager[0]["redundancy_required"] = rr
					}
					if gn, ok := val.Args["groupName"]; ok {
						keyManager[0]["group_name"] = gn
					}
					if cn, ok := val.Args["certificateName"]; ok {
						keyManager[0]["certificate_name"] = cn
					}
					if ln, ok := val.Args["loginName"]; ok {
						keyManager[0]["login_name"] = ln
					}
					if p, ok := val.Args["password"]; ok {
						keyManager[0]["password"] = p
					}
				}

			}
			// setting MpSettings
			mpSettings = append(mpSettings, map[string]interface{}{
				"administrator_account": adminAcc,
				"directory":             directory,
				"key_manager":           keyManager,
				"directory_groups":      directoryGroups,
				"local_accounts":        localAccounts,
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
	sanSystemCredentials := make([]interface{}, 0)
	if len(spt.SanStorage.SanSystemCredentials) != 0 {
		for i := 0; i < len(spt.SanStorage.SanSystemCredentials); i++ {
			sanSystemCredentials = append(sanSystemCredentials, map[string]interface{}{
				"chap_level":         spt.SanStorage.SanSystemCredentials[i].ChapLevel,
				"chap_name":          spt.SanStorage.SanSystemCredentials[i].ChapName,
				"chap_secret":        spt.SanStorage.SanSystemCredentials[i].ChapSecret,
				"chap_source":        spt.SanStorage.SanSystemCredentials[i].ChapSource,
				"mutual_chap_name":   spt.SanStorage.SanSystemCredentials[i].MutualChapName,
				"mutual_chap_secret": spt.SanStorage.SanSystemCredentials[i].MutualChapSecret,
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
					"is_permanent":       spt.SanStorage.VolumeAttachments[i].Volume.IsPermanent,
					"template_uri":       spt.SanStorage.VolumeAttachments[i].Volume.TemplateUri.String(),
					"properties":         properties,
				})

			}

			volumeAttachments = append(volumeAttachments, map[string]interface{}{
				"associated_template_attachment_id": spt.SanStorage.VolumeAttachments[i].AssociatedTemplateAttachmentId,
				"boot_volume_priority":              spt.SanStorage.VolumeAttachments[i].BootVolumePriority,
				"id":                                spt.SanStorage.VolumeAttachments[i].ID,
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

	if len(spt.LocalStorage.Controllers) != 0 {
		// Gets Storage Controller Body
		controllers := make([]map[string]interface{}, 0, len(spt.LocalStorage.Controllers))
		for i := 0; i < len(spt.LocalStorage.Controllers); i++ {
			logicalDrives := make([]map[string]interface{}, 0, len(spt.LocalStorage.Controllers[i].LogicalDrives))
			for j := 0; j < len(spt.LocalStorage.Controllers[i].LogicalDrives); j++ {
				logicalDrives = append(logicalDrives, map[string]interface{}{
					"accelerator":         spt.LocalStorage.Controllers[i].LogicalDrives[j].Accelerator,
					"bootable":            *spt.LocalStorage.Controllers[i].LogicalDrives[j].Bootable,
					"drive_technology":    spt.LocalStorage.Controllers[i].LogicalDrives[j].DriveTechnology,
					"name":                spt.LocalStorage.Controllers[i].LogicalDrives[j].Name,
					"num_physical_drives": spt.LocalStorage.Controllers[i].LogicalDrives[j].NumPhysicalDrives,
					"num_spare_drives":    spt.LocalStorage.Controllers[i].LogicalDrives[j].NumSpareDrives,
					"raid_level":          spt.LocalStorage.Controllers[i].LogicalDrives[j].RaidLevel,
					"sas_logical_jbod_id": spt.LocalStorage.Controllers[i].LogicalDrives[j].SasLogicalJBODId,
				})
			}
			controllers = append(controllers, map[string]interface{}{
				"device_slot":              spt.LocalStorage.Controllers[i].DeviceSlot,
				"initialize":               *spt.LocalStorage.Controllers[i].Initialize,
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
				"drive_min_size_gb":  spt.LocalStorage.SasLogicalJBODs[i].DriveMinSizeGB,
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
	return nil
}
