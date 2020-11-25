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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strings"
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
			"boot_order": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
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
							Type:     schema.TypeSet,
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
			"network": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"function_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"network_uri": {
							Type:     schema.TypeString,
							Required: true,
						},
						"port_id": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "Lom 1:1-a",
						},
						"requested_mbps": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "2500",
						},
						"id": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"boot": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"priority": {
										Type:     schema.TypeString,
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
										Type:     schema.TypeSet,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"chap_level": {
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
												"initiator_name_source": {
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
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gateway": {
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
			"server_hardware_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"enclosure_group": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"affinity": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Bay",
			},
			"hide_unused_flex_nics": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"manage_connections": {
				Type:     schema.TypeBool,
				Optional: true,
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
				Default:  "Virtual",
				ForceNew: true,
			},
			"wwn_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Virtual",
				ForceNew: true,
			},
			"mac_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Virtual",
				ForceNew: true,
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
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manage_local_storage": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"initialize": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"logical_drives": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bootable": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"raid_level": {
							Type:     schema.TypeString,
							Optional: true,
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
									"templateuri": {
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
				Default:  "ServerProfileV9",
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

	if _, ok := d.GetOk("network"); ok {
		rawNetwork := d.Get("network").(*schema.Set).List()
		networks := make([]ov.Connection, 0)
		for _, rawNet := range rawNetwork {
			rawNetworkItem := rawNet.(map[string]interface{})

			bootOptions := ov.BootOption{}
			if rawNetworkItem["boot"] != nil {
				rawBoots := rawNetworkItem["boot"].(*schema.Set).List()
				for _, rawBoot := range rawBoots {
					bootItem := rawBoot.(map[string]interface{})

					iscsi := ov.BootIscsi{}
					if bootItem["iscsi"] != nil {
						rawIscsis := bootItem["iscsi"].(*schema.Set).List()
						for _, rawIscsi := range rawIscsis {
							rawIscsiItem := rawIscsi.(map[string]interface{})
							iscsi = ov.BootIscsi{
								Chaplevel:            rawIscsiItem["chap_level"].(string),
								FirstBootTargetIp:    rawIscsiItem["first_boot_target_ip"].(string),
								FirstBootTargetPort:  rawIscsiItem["first_boot_target_ip"].(string),
								SecondBootTargetIp:   rawIscsiItem["second_boot_target_ip"].(string),
								SecondBootTargetPort: rawIscsiItem["second_boot_target_port"].(string),
							}
						}
					}

					bootOptions = ov.BootOption{
						Priority:         bootItem["priority"].(string),
						EthernetBootType: bootItem["ethernet_boot_type"].(string),
						BootVolumeSource: bootItem["boot_volume_source"].(string),
						Iscsi:            &iscsi,
					}
				}
			}

			ipv4 := ov.Ipv4Option{}
			if rawNetworkItem["ipv4"] != nil {
				rawIpv4s := rawNetworkItem["ipv4"].(*schema.Set).List()
				for _, rawIpv4 := range rawIpv4s {
					rawIpv4Item := rawIpv4.(map[string]interface{})
					ipv4 = ov.Ipv4Option{
						Gateway:         rawIpv4Item["gateway"].(string),
						IpAddressSource: rawIpv4Item["ip_address_source"].(string),
					}
				}
			}

			networks = append(networks, ov.Connection{
				ID:            rawNetworkItem["id"].(int),
				Name:          rawNetworkItem["name"].(string),
				FunctionType:  rawNetworkItem["function_type"].(string),
				NetworkURI:    utils.NewNstring(rawNetworkItem["network_uri"].(string)),
				PortID:        rawNetworkItem["port_id"].(string),
				RequestedMbps: rawNetworkItem["requested_mbps"].(string),
				Ipv4:          &ipv4,
				Boot:          &bootOptions,
			})
		}

		if val, ok := d.GetOk("manage_connections"); ok {
			serverProfile.ConnectionSettings.ManageConnections = val.(bool)
			serverProfile.ConnectionSettings.Connections = networks
		}
	}

	if val, ok := d.GetOk("boot_order"); ok {
		rawBootOrder := val.(*schema.Set).List()
		bootOrder := make([]string, len(rawBootOrder))
		for i, raw := range rawBootOrder {
			bootOrder[i] = raw.(string)
		}
		serverProfile.Boot.ManageBoot = true
		serverProfile.Boot.Order = bootOrder
		rawBootMode := d.Get("boot_mode").(*schema.Set).List()[0].(map[string]interface{})
		serverProfile.BootMode = ov.BootModeOption{
			ManageMode:    rawBootMode["manage_mode"].(bool),
			Mode:          rawBootMode["mode"].(string),
			PXEBootPolicy: utils.Nstring(rawBootMode["pxe_boot_policy"].(string)),
		}

	}

	if val, ok := d.GetOk("bios_option"); ok {
		rawBiosOption := val.(*schema.Set).List()
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
			biosOption = ov.BiosOption{
				ManageBios:         rawBiosItem["manage_bios"].(bool),
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
		rawLocalStorage := d.Get("local_storage").(*schema.Set).List()
		localStorage := ov.LocalStorageOptions{}
		for _, raw := range rawLocalStorage {
			localStorageItem := raw.(map[string]interface{})
			localStorage = ov.LocalStorageOptions{
				ManageLocalStorage: localStorageItem["manage_local_storage"].(bool),
				Initialize:         localStorageItem["initialize"].(bool),
			}
		}
		serverProfile.LocalStorage = localStorage
	}

	if _, ok := d.GetOk("logical_drives"); ok {
		rawLogicalDrives := d.Get("logical_drives").(*schema.Set).List()
		logicalDrives := make([]ov.LogicalDrive, 0)
		for _, rawLogicalDrive := range rawLogicalDrives {
			logicalDrivesItem := rawLogicalDrive.(map[string]interface{})
			logicalDrives = append(logicalDrives, ov.LogicalDrive{
				Bootable:  logicalDrivesItem["bootable"].(bool),
				RaidLevel: logicalDrivesItem["raid_level"].(string),
			})
		}
		serverProfile.LocalStorage.LogicalDrives = logicalDrives
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

	serverHardware, err := config.ovClient.GetServerHardwareByUri(serverProfile.ServerHardwareURI)
	if err != nil {
		return err
	}

	d.Set("hardware_uri", serverHardware.URI.String())
	d.Set("ilo_ip", serverHardware.GetIloIPAddress())
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

	var connections []ov.Connection
	if len(serverProfile.ConnectionSettings.Connections) != 0 {
		connections = serverProfile.ConnectionSettings.Connections
	}

	if len(connections) != 0 {
		networks := make([]map[string]interface{}, 0, len(connections))
		for _, rawNet := range connections {
			networks = append(networks, map[string]interface{}{
				"name":           rawNet.Name,
				"function_type":  rawNet.FunctionType,
				"network_uri":    rawNet.NetworkURI.String(),
				"port_id":        rawNet.PortID,
				"requested_mbps": rawNet.RequestedMbps,
			})
		}
		d.Set("network", networks)
	}

	if serverProfile.Boot.ManageBoot {
		bootOrder := make([]interface{}, 0)
		for _, currBoot := range serverProfile.Boot.Order {
			rawBootOrder := d.Get("boot_order").(*schema.Set).List()
			for _, raw := range rawBootOrder {
				if raw == currBoot {
					bootOrder = append(bootOrder, currBoot)
				}
			}
		}
		d.Set("boot_order", bootOrder)
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

		if _, ok := d.GetOk("network"); ok {
			rawNetwork := d.Get("network").(*schema.Set).List()
			networks := make([]ov.Connection, 0)
			for _, rawNet := range rawNetwork {
				rawNetworkItem := rawNet.(map[string]interface{})

				bootOptions := ov.BootOption{}
				if rawNetworkItem["boot"] != nil {
					rawBoots := rawNetworkItem["boot"].(*schema.Set).List()
					for _, rawBoot := range rawBoots {
						bootItem := rawBoot.(map[string]interface{})

						iscsi := ov.BootIscsi{}
						if bootItem["iscsi"] != nil {
							rawIscsis := bootItem["iscsi"].(*schema.Set).List()
							for _, rawIscsi := range rawIscsis {
								rawIscsiItem := rawIscsi.(map[string]interface{})
								iscsi = ov.BootIscsi{
									Chaplevel:            rawIscsiItem["chap_level"].(string),
									FirstBootTargetIp:    rawIscsiItem["first_boot_target_ip"].(string),
									FirstBootTargetPort:  rawIscsiItem["first_boot_target_ip"].(string),
									SecondBootTargetIp:   rawIscsiItem["second_boot_target_ip"].(string),
									SecondBootTargetPort: rawIscsiItem["second_boot_target_port"].(string),
								}
							}
						}

						bootOptions = ov.BootOption{
							Priority:         bootItem["priority"].(string),
							EthernetBootType: bootItem["ethernet_boot_type"].(string),
							BootVolumeSource: bootItem["boot_volume_source"].(string),
							Iscsi:            &iscsi,
						}
					}
				}

				ipv4 := ov.Ipv4Option{}
				if rawNetworkItem["ipv4"] != nil {
					rawIpv4s := rawNetworkItem["ipv4"].(*schema.Set).List()
					for _, rawIpv4 := range rawIpv4s {
						rawIpv4Item := rawIpv4.(map[string]interface{})
						ipv4 = ov.Ipv4Option{
							Gateway:         rawIpv4Item["gateway"].(string),
							IpAddressSource: rawIpv4Item["ip_address_source"].(string),
						}
					}
				}

				networks = append(networks, ov.Connection{
					ID:            rawNetworkItem["id"].(int),
					Name:          rawNetworkItem["name"].(string),
					FunctionType:  rawNetworkItem["function_type"].(string),
					NetworkURI:    utils.NewNstring(rawNetworkItem["network_uri"].(string)),
					PortID:        rawNetworkItem["port_id"].(string),
					RequestedMbps: rawNetworkItem["requested_mbps"].(string),
					Ipv4:          &ipv4,
					Boot:          &bootOptions,
				})
			}
			if val, ok := d.GetOk("manage_connections"); ok {
				serverProfile.ConnectionSettings.ManageConnections = val.(bool)
				serverProfile.ConnectionSettings.Connections = networks
			}
		}

		if val, ok := d.GetOk("boot_order"); ok {
			rawBootOrder := val.(*schema.Set).List()
			bootOrder := make([]string, len(rawBootOrder))
			for i, raw := range rawBootOrder {
				bootOrder[i] = raw.(string)
			}
			serverProfile.Boot.ManageBoot = true
			serverProfile.Boot.Order = bootOrder
			rawBootMode := d.Get("boot_mode").(*schema.Set).List()[0].(map[string]interface{})
			serverProfile.BootMode = ov.BootModeOption{
				ManageMode:    rawBootMode["manage_mode"].(bool),
				Mode:          rawBootMode["mode"].(string),
				PXEBootPolicy: utils.Nstring(rawBootMode["pxe_boot_policy"].(string)),
			}

		}

		if val, ok := d.GetOk("bios_option"); ok {
			rawBiosOption := val.(*schema.Set).List()
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
				biosOption = ov.BiosOption{
					ManageBios:         rawBiosItem["manage_bios"].(bool),
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

		// Get local storage data if provided
		if _, ok := d.GetOk("local_storage"); ok {
			rawLocalStorage := d.Get("local_storage").(*schema.Set).List()
			localStorage := ov.LocalStorageOptions{}
			for _, raw := range rawLocalStorage {
				localStorageItem := raw.(map[string]interface{})
				localStorage = ov.LocalStorageOptions{
					ManageLocalStorage: localStorageItem["manage_local_storage"].(bool),
					Initialize:         localStorageItem["initialize"].(bool),
				}
			}
			serverProfile.LocalStorage = localStorage
		}

		if _, ok := d.GetOk("logical_drives"); ok {
			rawLogicalDrives := d.Get("logical_drives").(*schema.Set).List()
			logicalDrives := make([]ov.LogicalDrive, 0)
			for _, rawLogicalDrive := range rawLogicalDrives {
				logicalDrivesItem := rawLogicalDrive.(map[string]interface{})
				logicalDrives = append(logicalDrives, ov.LogicalDrive{
					Bootable:  logicalDrivesItem["bootable"].(bool),
					RaidLevel: logicalDrivesItem["raid_level"].(string),
				})
			}
			serverProfile.LocalStorage.LogicalDrives = logicalDrives
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
