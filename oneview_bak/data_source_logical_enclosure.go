// (C) Copyright 2018 Hewlett Packard Enterprise Development LP
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

func dataSourceLogicalEnclosure() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLogicalEnclosureRead,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ambient_temperature_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"delete_failed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"deployment_manager_settings": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deployement_cluster_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"deployment_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"deployment_network_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"manage_os_deployment": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"enclosure_group_uri": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enclosure_uris": {
				Optional: true,
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"firmware": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"firmware_baseline_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"firmware_update_on": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"force_install_firmware": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"logical_interconnect_update_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"update_firmware_on_unmanaged_interconnect": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"validate_if_li_firmware_update_is_non_disruptive": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},

			"ip_addressing_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv4_range": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_servers": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"gateway": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_range_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subnet_mask": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"logical_interconnect_uris": {
				Computed: true,
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
				Computed: true,
			},
			"scaling_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scopes_uri": {
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
		},
	}
}

func dataSourceLogicalEnclosureRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	logicalEnclosure, err := config.ovClient.GetLogicalEnclosureByName(d.Get("name").(string))
	if err != nil {
		d.SetId("")
		return nil
	}
	d.SetId("name")
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
