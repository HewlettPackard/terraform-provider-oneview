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

func dataSourceServerProfileTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServerProfileTemplateRead,

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
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"affinity": {
				Type:     schema.TypeString,
				Computed: true,
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
					},
				},
			},
			"bios_option": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
										Type:     schema.TypeString,
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
							//							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
						"sas_logical_jbod": {
							Optional: true,
							Type:     schema.TypeList,
							//							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"description": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"drive_slot": {
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
			"hide_unused_flex_nics": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_hardware_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enclosure_group": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial_number_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"wwn_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mac_type": {
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
		d.SetId("")
		return nil
	}

	d.SetId(name)
	d.Set("name", spt.Name)
	d.Set("type", spt.Type)
	d.Set("affinity", spt.Affinity)
	d.Set("uri", spt.URI.String())

	enclosureGroup, err := config.ovClient.GetEnclosureGroupByUri(spt.EnclosureGroupURI)
	if err != nil {
		return err
	}
	d.Set("enclosure_group", enclosureGroup.Name)

	serverHardwareType, err := config.ovClient.GetServerHardwareTypeByUri(spt.ServerHardwareTypeURI)
	if err != nil {
		return err
	}

	d.Set("server_hardware_type", serverHardwareType.Name)

	d.Set("etag", spt.ETAG)
	d.Set("serial_number_type", spt.SerialNumberType)
	d.Set("wwn_type", spt.WWNType)
	d.Set("mac_type", spt.MACType)
	d.Set("hide_unused_flex_nics", spt.HideUnusedFlexNics)

	if len(spt.ConnectionSettings.Connections) != 0 {
		// Get connections
		connections := make([]map[string]interface{}, 0, len(spt.ConnectionSettings.Connections))
		for i := 0; i < len(spt.ConnectionSettings.Connections); i++ {
			// Gets Boot for Connection
			iscsi := make([]map[string]interface{}, 0, 1)
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

			// Gets Boot Settings
			connectionBoot := make([]map[string]interface{}, 0, 1)
			if spt.ConnectionSettings.Connections[i].Boot != nil {
				connectionBoot = append(connectionBoot, map[string]interface{}{
					"priority":           spt.ConnectionSettings.Connections[i].Boot.Priority,
					"boot_vlan_id":       spt.ConnectionSettings.Connections[i].Boot.BootOptionV3.BootVlanId,
					"ethernet_boot_type": spt.ConnectionSettings.Connections[i].Boot.EthernetBootType,
					"boot_volume_source": spt.ConnectionSettings.Connections[i].Boot.BootVolumeSource,
					"iscsi":              iscsi,
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
				"function_type":  spt.ConnectionSettings.Connections[i].FunctionType,
				"network_uri":    spt.ConnectionSettings.Connections[i].NetworkURI,
				"port_id":        spt.ConnectionSettings.Connections[i].PortID,
				"requested_mbps": spt.ConnectionSettings.Connections[i].RequestedMbps,
				"id":             spt.ConnectionSettings.Connections[i].ID,
				"isolated_trunk": spt.ConnectionSettings.Connections[i].IsolatedTrunk,
				"lag_name":       spt.ConnectionSettings.Connections[i].LagName,
				"mac_type":       spt.ConnectionSettings.Connections[i].MacType,
				"managed":        spt.ConnectionSettings.Connections[i].Managed,
				"network_name":   spt.ConnectionSettings.Connections[i].NetworkName,
				"boot":           connectionBoot,
				"ipv4":           connectionIpv4,
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
		"manage_boot": spt.Boot.ManageBoot,
		"boot_order":  bootOrder,
	})
	d.Set("boot", boot)

	overriddenSettings := make([]interface{}, 0, len(spt.Bios.OverriddenSettings))
	for _, overriddenSetting := range spt.Bios.OverriddenSettings {
		overriddenSettings = append(overriddenSettings, map[string]interface{}{
			"id":    overriddenSetting.ID,
			"value": overriddenSetting.Value,
		})
	}
	if spt.Bios != nil {
		biosOptions := make([]map[string]interface{}, 0, 1)
		biosOptions = append(biosOptions, map[string]interface{}{
			"manage_bios":         spt.Bios.ManageBios,
			"overridden_settings": overriddenSettings,
		})

		d.Set("bios_option", biosOptions)
	}

	if len(spt.LocalStorage.Controllers) != 0 {
		// Gets Storage Controller Body
		controllers := make([]map[string]interface{}, 0, len(spt.LocalStorage.Controllers))
		for i := 0; i < len(spt.LocalStorage.Controllers); i++ {
			logicalDrives := make([]map[string]interface{}, 0, len(spt.LocalStorage.Controllers[i].LogicalDrives))
			for j := 0; j < len(spt.LocalStorage.Controllers[i].LogicalDrives); j++ {
				logicalDrives = append(logicalDrives, map[string]interface{}{
					"bootable":            *spt.LocalStorage.Controllers[i].LogicalDrives[j].Bootable,
					"accelerator":         spt.LocalStorage.Controllers[i].LogicalDrives[j].Accelerator,
					"drive_technology":    spt.LocalStorage.Controllers[i].LogicalDrives[j].DriveTechnology,
					"name":                spt.LocalStorage.Controllers[i].LogicalDrives[j].Name,
					"num_physical_drives": spt.LocalStorage.Controllers[i].LogicalDrives[j].NumPhysicalDrives,
					"num_spare_drives":    spt.LocalStorage.Controllers[i].LogicalDrives[j].NumSpareDrives,
					"sas_logical_jbod_id": spt.LocalStorage.Controllers[i].LogicalDrives[j].SasLogicalJBODId,
					"raid_level":          spt.LocalStorage.Controllers[i].LogicalDrives[j].RaidLevel,
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
			"manage_local_storage": spt.LocalStorage.ManageLocalStorage,
			"initialize":           spt.LocalStorage.Initialize,
			"controller":           controllers,
			"sas_logical_jbod":     sasLogDrives,
		})
		d.Set("local_storage", localStorage)
	}
	return nil
}
