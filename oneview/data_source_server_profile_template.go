// (C) Copyright 2019 Hewlett Packard Enterprise Development LP
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
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceServerProfileTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServerProfileTemplateRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"boot_order": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
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
						// "compliance_control":{
						// 	Type:     schema.TypeString,
						// 	Required: true,
						// },
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
			"network": {
				Optional: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"function_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"port_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"requested_mbps": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
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
			"controllers": {
				Optional: true,
				Type:     schema.TypeSet,
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
					},
				},
			},
			"logical_drives": {
				Optional: true,
				Type:     schema.TypeSet,
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
			"logical_jbod": {
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
						"num_physical_drives": {
							Type:     schema.TypeString,
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
				"id":             rawNet.ID,
			})
		}
		d.Set("network", networks)
	}

	if spt.Boot.ManageBoot {
		bootOrder := make([]interface{}, len(spt.Boot.Order))
		for i, currBoot := range spt.Boot.Order {
			bootOrder[i] = currBoot
		}
		d.Set("boot_order", bootOrder)
	}

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

	localStorages := make([]map[string]interface{}, 0, 1)
	localStorages = append(localStorages, map[string]interface{}{
		//"manage_local_storage": serverProfile.LocalStorage.ManageLocalStorage,
		//"initialize":  serverProfile.LocalStorage.Initialize,
		"reapply_state": spt.LocalStorage.ReapplyState,
	})

	d.Set("local_storage", localStorages)
	controllers := make([]map[string]interface{}, 0, len(spt.LocalStorage.Controllers))
	for _, controller := range spt.LocalStorage.Controllers {

		controllers = append(controllers, map[string]interface{}{
			"device_slot":          controller.DeviceSlot,
			"drive_write_cache":    controller.DriveWriteCache,
			"import_configuration": controller.ImportConfiguration,
			"initialize":           controller.Initialize,
			"mode":                 controller.Mode,
			"predictive_spare_rebuild": controller.PredictiveSpareRebuild,
		})
		logicaldrives := make([]map[string]interface{}, 0, len(controller.LogicalDrives))
		for _, logicaldrive := range controller.LogicalDrives {

			logicaldrives = append(logicaldrives, map[string]interface{}{
				"accelerator":         logicaldrive.Accelerator,
				"bootable":            logicaldrive.Bootable,
				"drive_number":        logicaldrive.DriveNumber,
				"drive_technology":    logicaldrive.DriveTechnology,
				"name":                logicaldrive.Name,
				"num_physical_drives": logicaldrive.NumPhysicalDrives,
				"num_spare_drives":    logicaldrive.NumSpareDrives,
				"raid_level":          logicaldrive.RaidLevel,
				"sas_logical_jbod_id": logicaldrive.SasLogicalJBODId,
			})
		}

		d.Set("logical_drives", logicaldrives)

	}

	d.Set("controllers", controllers)

	saslogicaljbods := make([]map[string]interface{}, 0, len(spt.LocalStorage.SasLogicalJBODs))
	for _, saslogicaljbod := range spt.LocalStorage.SasLogicalJBODs {

		saslogicaljbods = append(saslogicaljbods, map[string]interface{}{
			"description":          saslogicaljbod.Description,
			"device_slot":          saslogicaljbod.DeviceSlot,
			"drive_max_size_gb":    saslogicaljbod.DriveMaxSizeGB,
			"drive_min_size_gb":    saslogicaljbod.DriveMinSizeGB,
			"drive_technology":     saslogicaljbod.DriveTechnology,
			"erase_data":           saslogicaljbod.EraseData,
			"id":                   saslogicaljbod.ID,
			"name":                 saslogicaljbod.Name,
			"num_physical_drives":  saslogicaljbod.NumPhysicalDrives,
			"persistent":           saslogicaljbod.Persistent,
			"sas_logical_jbod_uri": saslogicaljbod.SasLogicalJBODUri,
			"status":               saslogicaljbod.Status,
		})
	}
	d.Set("logical_jbod", saslogicaljbods)

	return nil
}
