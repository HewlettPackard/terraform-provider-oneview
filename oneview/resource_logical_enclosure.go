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
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
)

func resourceLogicalEnclosure() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogicalEnclosureCreate,
		Read:   resourceLogicalEnclosureRead,
		Update: resourceLogicalEnclosureUpdate,
		Delete: resourceLogicalEnclosureDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ambient_temperature_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"delete_failed": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"deployment_manager_settings": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deployement_cluster_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"deployment_mode": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"deployment_network_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"manage_os_deployment": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"enclosure_group_uri": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enclosure_uris": {
				Required: true,
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"firmware": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"firmware_baseline_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"firmware_update_on": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"force_install_firmware": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"logical_interconnect_update_mode": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"update_firmware_on_unmanaged_interconnect": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"validate_if_li_firmware_update_is_non_disruptive": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"ip_addressing_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv4_range": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_servers": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"domain": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"gateway": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip_range_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
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
			"logical_interconnect_uris": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"power_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scaling_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scopes_uri": {
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
			"update_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}

}
func resourceLogicalEnclosureCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	logicalEnclosure := ov.LogicalEnclosure{
		Name:              d.Get("name").(string),
		EnclosureGroupUri: utils.NewNstring(d.Get("enclosure_group_uri").(string)),
	}
	enclosureSetCount := d.Get("enclosure_uris.#").(int)
	enclosureUris := make([]utils.Nstring, enclosureSetCount)
	for i := 0; i < enclosureSetCount; i++ {
		enclosureSetPrefix := fmt.Sprintf("enclosure_uris.%d", i)
		if val, ok := d.GetOk(enclosureSetPrefix); ok {
			enclosureUris[i] = utils.NewNstring(val.(string))
		}
	}
	logicalEnclosure.EnclosureUris = enclosureUris

	firmwareList := d.Get("firmware").(*schema.Set).List()
	for _, raw := range firmwareList {
		firmware := raw.(map[string]interface{})
		logicalEnclosureFirmware := ov.LogicalEnclosureFirmware{
			FirmwareBaselineUri:  utils.NewNstring(firmware["firmware_baseline_uri"].(string)),
			ForceInstallFirmware: firmware["force_install_firmware"].(bool),
		}
		logicalEnclosure.Firmware = &logicalEnclosureFirmware
	}
	logicalEnclosureError := config.ovClient.CreateLogicalEnclosure(logicalEnclosure)
	d.SetId(d.Get("name").(string))
	if logicalEnclosureError != nil {
		d.SetId("")
		return logicalEnclosureError
	}
	return resourceLogicalEnclosureRead(d, meta)
}

func resourceLogicalEnclosureRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	logicalEnclosure, err := config.ovClient.GetLogicalEnclosureByName(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}
	d.SetId(logicalEnclosure.Name)
	d.Set("ambient_temperature_mode", logicalEnclosure.AmbientTemperatureMode)
	d.Set("category", logicalEnclosure.Category)
	d.Set("created", logicalEnclosure.Created)
	d.Set("delete_failed", logicalEnclosure.DeleteFailed)
	deploymentManagerSettings := make([]map[string]interface{}, 0, 1)
	deploymentManagerSettings = append(deploymentManagerSettings, map[string]interface{}{
		"deployment_mode":         logicalEnclosure.DeploymentManagerSettings.OsDeploymentSettings.DeploymentModeSettings.DeploymentMode,
		"deployment_network_uri":  logicalEnclosure.DeploymentManagerSettings.OsDeploymentSettings.DeploymentModeSettings.DeploymentNetworkUri,
		"manage_os_deployment":    logicalEnclosure.DeploymentManagerSettings.OsDeploymentSettings.ManageOSDeployment,
		"deployement_cluster_uri": logicalEnclosure.DeploymentManagerSettings.DeploymentClusterUri,
	})
	d.Set("deployment_manager_settings", deploymentManagerSettings)
	d.Set("description", logicalEnclosure.Description)
	d.Set("enclosure_group_uri", logicalEnclosure.EnclosureGroupUri.String())
	d.Set("enclosure_uris", logicalEnclosure.EnclosureUris)
	logicalEnclosureFirmware := make([]map[string]interface{}, 0, 1)
	logicalEnclosureFirmware = append(logicalEnclosureFirmware, map[string]interface{}{
		"firmware_baseline_uri":                            logicalEnclosure.Firmware.FirmwareBaselineUri,
		"firmware_update_on":                               logicalEnclosure.Firmware.FirmwareUpdateOn,
		"force_install_firmware":                           logicalEnclosure.Firmware.ForceInstallFirmware,
		"logical_interconnect_update_mode":                 logicalEnclosure.Firmware.LogicalInterconnectUpdateMode,
		"update_firmware_on_unmanaged_interconnect":        logicalEnclosure.Firmware.UpdateFirmwareOnUnmanagedInterconnect,
		"validate_if_li_firmware_update_is_non_disruptive": logicalEnclosure.Firmware.ValidateIfLIFirmwareUpdateIsNonDisruptive,
	})
	d.Set("firmware", logicalEnclosureFirmware)
	d.Set("ip_addressing_mode", logicalEnclosure.IpAddressingMode)
	logicalEnclosureIpv4Ranges := make([]map[string]interface{}, 0, len(logicalEnclosure.Ipv4Ranges))
	for _, logicalEnclosureIpv4Range := range logicalEnclosure.Ipv4Ranges {
		dnsServerMap := make([]interface{}, len(logicalEnclosureIpv4Range.DnsServers))
		for i, dnsServer := range logicalEnclosureIpv4Range.DnsServers {
			dnsServerMap[i] = dnsServer
		}
		logicalEnclosureIpv4Ranges = append(logicalEnclosureIpv4Ranges, map[string]interface{}{
			"dns_servers":  schema.NewSet(schema.HashString, dnsServerMap),
			"domain":       logicalEnclosureIpv4Range.Domain,
			"gateway":      logicalEnclosureIpv4Range.Gateway,
			"ip_range_uri": logicalEnclosureIpv4Range.IpRangeUri,
			"name":         logicalEnclosureIpv4Range.Name,
			"subnet_mask":  logicalEnclosureIpv4Range.SubnetMask,
		})
	}
	ipv4RangeCount := d.Get("ipv4_range.#").(int)
	oneviewIpv4RangeCount := len(logicalEnclosureIpv4Ranges)
	for i := 0; i < ipv4RangeCount; i++ {
		currIpv4RangeName := d.Get("ipv4_range." + strconv.Itoa(i) + ".name").(string)
		for j := 0; j < oneviewIpv4RangeCount; j++ {
			if currIpv4RangeName == logicalEnclosureIpv4Ranges[j]["name"] && i <= j {
				logicalEnclosureIpv4Ranges[i], logicalEnclosureIpv4Ranges[j] = logicalEnclosureIpv4Ranges[j], logicalEnclosureIpv4Ranges[i]
			}
		}
	}
	d.Set("ipv4_range", logicalEnclosureIpv4Ranges)
	d.Set("logical_interconnect_uris", logicalEnclosure.LogicalInterconnectUris)
	d.Set("modified", logicalEnclosure.Modified)
	d.Set("name", logicalEnclosure.Name)
	d.Set("power_mode", logicalEnclosure.PowerMode)
	d.Set("scaling_state", logicalEnclosure.ScalingState)
	d.Set("scopes_uri", logicalEnclosure.ScopesUri)
	d.Set("status", logicalEnclosure.Status)
	d.Set("type", logicalEnclosure.Type)
	d.Set("uri", logicalEnclosure.URI.String())
	return nil
}

func resourceLogicalEnclosureUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	if val, ok := d.GetOk("update_type"); ok {
		if val.(string) == "updateByGroup" {
			id := d.Id()
			logicalEnclosure, err := config.ovClient.GetLogicalEnclosureByName(id)
			err = config.ovClient.UpdateFromGroupLogicalEnclosure(logicalEnclosure)

			if err != nil {
				return err
			}
			d.SetId(id)

			return resourceLogicalEnclosureRead(d, meta)
		}
	}

	logicalEnclosure := ov.LogicalEnclosure{
		Name: d.Get("name").(string),
		Type: d.Get("type").(string),
	}
	if val, ok := d.GetOk("ambient_temperature_mode"); ok {
		logicalEnclosure.AmbientTemperatureMode = val.(string)
	}
	if val, ok := d.GetOk("delete_failed"); ok {
		logicalEnclosure.DeleteFailed = val.(bool)
	}
	if val, ok := d.GetOk("uri"); ok {
		logicalEnclosure.URI = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk("power_mode"); ok {
		logicalEnclosure.PowerMode = val.(string)
	}
	if val, ok := d.GetOk("scaling_state"); ok {
		logicalEnclosure.PowerMode = val.(string)
	}

	if val, ok := d.GetOk("scopes_uri"); ok {
		logicalEnclosure.ScopesUri = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk("enclosure_group_uri"); ok {
		logicalEnclosure.EnclosureGroupUri = utils.NewNstring(val.(string))
	}

	deploymentManagerSettingsList := d.Get("deployment_manager_settings").(*schema.Set).List()
	for _, raw := range deploymentManagerSettingsList {
		deploymentManagerSetting := raw.(map[string]interface{})
		deploymentModeSettings := ov.DeploymentModeSettings{
			DeploymentMode:       deploymentManagerSetting["deployment_mode"].(string),
			DeploymentNetworkUri: utils.NewNstring(deploymentManagerSetting["deployment_network_uri"].(string)),
		}
		leOsDeploymentSettings := ov.LeOsDeploymentSettings{
			ManageOSDeployment:     deploymentManagerSetting["manage_os_deployment"].(bool),
			DeploymentModeSettings: &deploymentModeSettings,
		}
		deploymentClusterUri := utils.NewNstring("")
		if deploymentManagerSetting["deployment_cluster_uri"] != nil {
			deploymentClusterUri = utils.NewNstring(deploymentManagerSetting["deployment_cluster_uri"].(string))
		}
		deploymentManagerSettings := ov.DeploymentManagerSettings{
			DeploymentClusterUri: deploymentClusterUri,
			OsDeploymentSettings: &leOsDeploymentSettings,
		}
		logicalEnclosure.DeploymentManagerSettings = &deploymentManagerSettings
	}
	enclosureSetCount := d.Get("enclosure_uris.#").(int)
	enclosureUris := make([]utils.Nstring, enclosureSetCount)
	for i := 0; i < enclosureSetCount; i++ {
		enclosureSetPrefix := fmt.Sprintf("enclosure_uris.%d", i)
		if val, ok := d.GetOk(enclosureSetPrefix); ok {
			enclosureUris[i] = utils.NewNstring(val.(string))
		}
	}
	logicalEnclosure.EnclosureUris = enclosureUris
	firmwareList := d.Get("firmware").(*schema.Set).List()
	for _, raw := range firmwareList {
		firmware := raw.(map[string]interface{})
		logicalEnclosureFirmware := ov.LogicalEnclosureFirmware{
			FirmwareBaselineUri:                       utils.NewNstring(firmware["firmware_baseline_uri"].(string)),
			FirmwareUpdateOn:                          firmware["firmware_update_on"].(string),
			ForceInstallFirmware:                      firmware["force_install_firmware"].(bool),
			LogicalInterconnectUpdateMode:             firmware["logical_interconnect_update_mode"].(string),
			UpdateFirmwareOnUnmanagedInterconnect:     firmware["update_firmware_on_unmanaged_interconnect"].(bool),
			ValidateIfLIFirmwareUpdateIsNonDisruptive: firmware["validate_if_li_firmware_update_is_non_disruptive"].(bool),
		}
		logicalEnclosure.Firmware = &logicalEnclosureFirmware
	}
	ipv4rangesList := d.Get("ipv4_range").(*schema.Set).List()
	ipv4rangesCollect := make([]ov.Ipv4Ranges, 0)
	for _, raw := range ipv4rangesList {
		ipv4range := raw.(map[string]interface{})
		dnsServers := make([]string, 0)
		dnsServers = append(dnsServers, ipv4range["dns_servers"].(string))
		ipv4ranges := ov.Ipv4Ranges{
			DnsServers: dnsServers,
			Domain:     ipv4range["domain"].(string),
			Gateway:    ipv4range["gateway"].(string),
			IpRangeUri: utils.NewNstring(ipv4range["ip_range_uri"].(string)),
			Name:       ipv4range["name"].(string),
			SubnetMask: ipv4range["subnet_mask"].(string),
		}
		ipv4rangesCollect = append(ipv4rangesCollect, ipv4ranges)
	}
	logicalEnclosure.Ipv4Ranges = ipv4rangesCollect

	err := config.ovClient.UpdateLogicalEnclosure(logicalEnclosure)
	if err != nil {
		return err
	}
	d.SetId(d.Get("name").(string))

	return resourceLogicalEnclosureRead(d, meta)
}

func resourceLogicalEnclosureDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteLogicalEnclosure(d.Get("name").(string))
	if err != nil {
		return err
	}
	return nil
}
