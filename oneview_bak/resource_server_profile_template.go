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
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ServerProfileTemplateV1",
			},
			"server_hardware_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enclosure_group": {
				Type:     schema.TypeString,
				Required: true,
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
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
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
			"os_deployment_settings": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"os_deployment_plan_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"os_volume_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
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

func resourceServerProfileTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	serverProfileTemplate := ov.ServerProfile{
		Name:               d.Get("name").(string),
		Type:               d.Get("type").(string),
		Affinity:           d.Get("affinity").(string),
		SerialNumberType:   d.Get("serial_number_type").(string),
		WWNType:            d.Get("wwn_type").(string),
		MACType:            d.Get("mac_type").(string),
		HideUnusedFlexNics: d.Get("hide_unused_flex_nics").(bool),
	}

	enclosureGroup, err := config.ovClient.GetEnclosureGroupByName(d.Get("enclosure_group").(string))
	if err != nil {
		return err
	}
	serverProfileTemplate.EnclosureGroupURI = enclosureGroup.URI

	serverHardwareType, err := config.ovClient.GetServerHardwareTypeByName(d.Get("server_hardware_type").(string))
	if err != nil {
		return err
	}
	serverProfileTemplate.ServerHardwareTypeURI = serverHardwareType.URI

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
		serverProfileTemplate.ConnectionSettings.ManageConnections = val.(bool)
		serverProfileTemplate.ConnectionSettings.Connections = networks
	}

	if val, ok := d.GetOk("boot_order"); ok {
		rawBootOrder := val.(*schema.Set).List()
		bootOrder := make([]string, len(rawBootOrder))
		for i, raw := range rawBootOrder {
			bootOrder[i] = raw.(string)
		}
		serverProfileTemplate.Boot.ManageBoot = true
		serverProfileTemplate.Boot.Order = bootOrder
		rawBootMode := d.Get("boot_mode").(*schema.Set).List()[0].(map[string]interface{})
		serverProfileTemplate.BootMode = ov.BootModeOption{
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
		serverProfileTemplate.Firmware = firmware
	}
	// Get local storage data if provided
	rawLocalStorage := d.Get("local_storage").(*schema.Set).List()
	localStorage := ov.LocalStorageOptions{}
	for _, raw := range rawLocalStorage {
		localStorageItem := raw.(map[string]interface{})
		localStorage = ov.LocalStorageOptions{
			ManageLocalStorage: localStorageItem["manage_local_storage"].(bool),
			Initialize:         localStorageItem["initialize"].(bool),
		}
	}
	serverProfileTemplate.LocalStorage = localStorage

	rawLogicalDrives := d.Get("logical_drives").(*schema.Set).List()
	logicalDrives := make([]ov.LogicalDrive, 0)
	for _, rawLogicalDrive := range rawLogicalDrives {
		logicalDrivesItem := rawLogicalDrive.(map[string]interface{})
		logicalDrives = append(logicalDrives, ov.LogicalDrive{
			Bootable:  logicalDrivesItem["bootable"].(bool),
			RaidLevel: logicalDrivesItem["raid_level"].(string),
		})
	}
	serverProfileTemplate.LocalStorage.LogicalDrives = logicalDrives

	// get SAN storage data if provided
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
	serverProfileTemplate.SanStorage = sanStorage

	// Get volume attachment data for san storage
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
		serverProfileTemplate.SanStorage.VolumeAttachments = volumeAttachments
	}
	rawOsDeploySetting := d.Get("os_deployment_settings").(*schema.Set).List()
	osDeploySetting := ov.OSDeploymentSettings{}
	for _, raw := range rawOsDeploySetting {
		osDeploySettingItem := raw.(map[string]interface{})
		osDeploymentPlan, err := config.ovClient.GetOSDeploymentPlanByName(osDeploySettingItem["os_deployment_plan_name"].(string))
		if err != nil {
			return err
		}
		if osDeploymentPlan.URI == "" {
			return fmt.Errorf("Could not find deployment plan by name: %s", osDeploySettingItem["os_deployment_plan_name"].(string))
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
			OSDeploymentPlanUri: osDeploymentPlan.URI,
			OSVolumeUri:         utils.NewNstring(osDeploySettingItem["os_volume_uri"].(string)),
			OSCustomAttributes:  osCustomAttributes,
		}
	}

	serverProfileTemplate.OSDeploymentSettings = osDeploySetting

	sptError := config.ovClient.CreateProfileTemplate(serverProfileTemplate)
	d.SetId(d.Get("name").(string))
	if sptError != nil {
		d.SetId("")
		return sptError
	}
	return resourceServerProfileTemplateRead(d, meta)
}

func resourceServerProfileTemplateRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	spt, err := config.ovClient.GetProfileTemplateByName(d.Id())
	if err != nil || spt.URI.IsNil() {
		d.SetId("")
		return nil
	}

	d.Set("name", spt.Name)
	d.Set("type", spt.Type)

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
	d.Set("affinity", spt.Affinity)
	d.Set("uri", spt.URI.String())
	d.Set("etag", spt.ETAG)
	d.Set("serial_number_type", spt.SerialNumberType)
	d.Set("wwn_type", spt.WWNType)
	d.Set("mac_type", spt.MACType)
	d.Set("hide_unused_flex_nics", spt.HideUnusedFlexNics)

	var connections []ov.Connection
	if len(spt.ConnectionSettings.Connections) != 0 {
		connections = spt.ConnectionSettings.Connections
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
	if spt.Boot.ManageBoot {
		bootOrder := make([]interface{}, 0)
		for _, currBoot := range spt.Boot.Order {
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

func resourceServerProfileTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	serverProfileTemplate := ov.ServerProfile{
		Name:               d.Get("name").(string),
		Type:               d.Get("type").(string),
		Affinity:           d.Get("affinity").(string),
		URI:                utils.NewNstring(d.Get("uri").(string)),
		ETAG:               d.Get("etag").(string),
		SerialNumberType:   d.Get("serial_number_type").(string),
		WWNType:            d.Get("wwn_type").(string),
		MACType:            d.Get("mac_type").(string),
		HideUnusedFlexNics: d.Get("hide_unused_flex_nics").(bool),
	}

	enclosureGroup, err := config.ovClient.GetEnclosureGroupByName(d.Get("enclosure_group").(string))
	if err != nil {
		return err
	}
	serverProfileTemplate.EnclosureGroupURI = enclosureGroup.URI

	serverHardwareType, err := config.ovClient.GetServerHardwareTypeByName(d.Get("server_hardware_type").(string))
	if err != nil {
		return err
	}
	serverProfileTemplate.ServerHardwareTypeURI = serverHardwareType.URI

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

	if val, ok := d.GetOk("boot_order"); ok {
		rawBootOrder := val.(*schema.Set).List()
		bootOrder := make([]string, len(rawBootOrder))
		for i, raw := range rawBootOrder {
			bootOrder[i] = raw.(string)
		}
		serverProfileTemplate.Boot.ManageBoot = true
		serverProfileTemplate.Boot.Order = bootOrder
	}

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
		serverProfileTemplate.Firmware = firmware
	}
	// Get local storage data if provided
	rawLocalStorage := d.Get("local_storage").(*schema.Set).List()
	localStorage := ov.LocalStorageOptions{}
	for _, raw := range rawLocalStorage {
		localStorageItem := raw.(map[string]interface{})
		localStorage = ov.LocalStorageOptions{
			ManageLocalStorage: localStorageItem["manage_local_storage"].(bool),
			Initialize:         localStorageItem["initialize"].(bool),
		}
	}
	serverProfileTemplate.LocalStorage = localStorage

	rawLogicalDrives := d.Get("logical_drives").(*schema.Set).List()
	logicalDrives := make([]ov.LogicalDrive, 0)
	for _, rawLogicalDrive := range rawLogicalDrives {
		logicalDrivesItem := rawLogicalDrive.(map[string]interface{})
		logicalDrives = append(logicalDrives, ov.LogicalDrive{
			Bootable:  logicalDrivesItem["bootable"].(bool),
			RaidLevel: logicalDrivesItem["raid_level"].(string),
		})
	}
	serverProfileTemplate.LocalStorage.LogicalDrives = logicalDrives

	// get SAN storage data if provided
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
	serverProfileTemplate.SanStorage = sanStorage

	// Get volume attachment data for san storage
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
		serverProfileTemplate.SanStorage.VolumeAttachments = volumeAttachments
	}

	rawOsDeploySetting := d.Get("os_deployment_settings").(*schema.Set).List()
	osDeploySetting := ov.OSDeploymentSettings{}
	for _, raw := range rawOsDeploySetting {
		osDeploySettingItem := raw.(map[string]interface{})
		osDeploymentPlan, err := config.ovClient.GetOSDeploymentPlanByName(osDeploySettingItem["os_deployment_plan_name"].(string))
		if err != nil {
			return err
		}
		if osDeploymentPlan.URI == "" {
			return fmt.Errorf("Could not find deployment plan by name: %s", osDeploySettingItem["os_deployment_plan_name"].(string))
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
			OSDeploymentPlanUri: osDeploymentPlan.URI,
			OSVolumeUri:         utils.NewNstring(osDeploySettingItem["os_volume_uri"].(string)),
			OSCustomAttributes:  osCustomAttributes,
		}
	}

	serverProfileTemplate.OSDeploymentSettings = osDeploySetting

	err = config.ovClient.UpdateProfileTemplate(serverProfileTemplate)
	if err != nil {
		return err
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
