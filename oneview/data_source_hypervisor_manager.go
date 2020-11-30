// (C) Copyright 2020 Hewlett Packard Enterprise Development LP
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

func dataSourceHypervisorManager() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHypervisorManagerRead,

		Schema: map[string]*schema.Schema{
			"available_dvs_versions": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hypervisor_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"preferences": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"distributed_switch_usage": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"distributed_switch_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"drs_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"ha_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"multi_nic_v_motion": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"virtual_switch_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"refresh_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_paths": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"actual_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"scopes_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state_reason": {
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
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceHypervisorManagerRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	id := d.Get("name").(string)

	hypMan, err := config.ovClient.GetHypervisorManagerByName(id)
	if err != nil || hypMan.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.SetId(id)
	d.Set("available_dvs_versions", hypMan.AvailableDvsVersions)
	d.Set("category", hypMan.Category)
	d.Set("created", hypMan.Created)
	d.Set("description", hypMan.Description)
	d.Set("display_name", hypMan.DisplayName)
	d.Set("etag", hypMan.ETAG)
	d.Set("hypervisor_type", hypMan.HypervisorType)
	d.Set("modified", hypMan.Modified)
	d.Set("name", hypMan.Name)
	d.Set("password", hypMan.Password)
	d.Set("port", hypMan.Port)
	hypManPreferences := make([]map[string]interface{}, 0, 1)
	hypManPreferences = append(hypManPreferences, map[string]interface{}{
		"type":                       hypMan.Preferences.Type,
		"virtual_switch_type":        hypMan.Preferences.VirtualSwitchType,
		"distributed_switch_version": hypMan.Preferences.DistributedSwitchVersion,
		"distributed_switch_usage":   hypMan.Preferences.DistributedSwitchUsage,
		"multi_nic_v_motion":         hypMan.Preferences.MultiNicVMotion,
		"drs_enabled":                hypMan.Preferences.DrsEnabled,
		"ha_enabled":                 hypMan.Preferences.HaEnabled,
	})
	d.Set("preferences", hypManPreferences)
	d.Set("refresh_state", hypMan.RefreshState)
	hypManResourcePaths := make([]map[string]interface{}, 0, len(hypMan.ResourcePaths))
	for _, hypManResourcePath := range hypMan.ResourcePaths {
		hypManResourcePaths = append(hypManResourcePaths, map[string]interface{}{
			"user_path":   hypManResourcePath.UserPath,
			"actual_path": hypManResourcePath.ActualPath,
		})
	}
	d.Set("resource_paths", hypManResourcePaths)
	d.Set("scopes_uri", hypMan.ScopesUri)
	d.Set("state", hypMan.State)
	d.Set("state_reason", hypMan.StateReason)
	d.Set("status", hypMan.Status)
	d.Set("type", hypMan.Type)
	d.Set("uri", hypMan.URI)
	d.Set("username", hypMan.Username)
	d.Set("uuid", hypMan.UUID)
	d.Set("version", hypMan.Version)
	return nil
}
