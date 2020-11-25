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
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"io/ioutil"
)

func resourceHypervisorManager() *schema.Resource {
	return &schema.Resource{
		Create: resourceHypervisorManagerCreate,
		Read:   resourceHypervisorManagerRead,
		Update: resourceHypervisorManagerUpdate,
		Delete: resourceHypervisorManagerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"available_dvs_versions": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hypervisor_type": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  443,
			},
			"preferences": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"distributed_switch_usage": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"distributed_switch_version": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"drs_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"ha_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"multi_nic_v_motion": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"virtual_switch_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"refresh_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_paths": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"actual_path": {
							Type:     schema.TypeString,
							Optional: true,
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
				Required: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"initial_scope_uris": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
		},
	}
}

func resourceHypervisorManagerCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	hypMan := ov.HypervisorManager{
		Name:        d.Get("name").(string),
		DisplayName: d.Get("display_name").(string),
		Username:    d.Get("username").(string),
		Password:    d.Get("password").(string),
		Port:        d.Get("port").(int),
		Type:        d.Get("type").(string),
	}

	if val, ok := d.GetOk("initial_scope_uris"); ok {
		rawInitialScopeUris := val.(*schema.Set).List()
		initialScopeUris := make([]utils.Nstring, len(rawInitialScopeUris))
		for i, raw := range rawInitialScopeUris {
			initialScopeUris[i] = utils.Nstring(raw.(string))
		}
		hypMan.InitialScopeUris = initialScopeUris
	}

	preferencesList := d.Get("preferences").(*schema.Set).List()
	for _, raw := range preferencesList {
		preferences := raw.(map[string]interface{})
		hypervisorManagerPreferences := ov.Preference{
			Type:                     preferences["type"].(string),
			VirtualSwitchType:        preferences["virtual_switch_type"].(string),
			DistributedSwitchVersion: preferences["distributed_switch_version"].(string),
			DistributedSwitchUsage:   preferences["distributed_switch_usage"].(string),
			MultiNicVMotion:          preferences["multi_nic_v_motion"].(bool),
			DrsEnabled:               preferences["drs_enabled"].(bool),
			HaEnabled:                preferences["ha_enabled"].(bool),
		}

		hypMan.Preferences = &hypervisorManagerPreferences
	}

	resourcePathList := d.Get("resource_paths").(*schema.Set).List()
	resourcePathCollect := make([]ov.ResourcePath, 0)
	for _, raw := range resourcePathList {
		resourcePaths := raw.(map[string]interface{})
		hypervisorManagerResourcePaths := ov.ResourcePath{
			UserPath:   resourcePaths["user_path"].(string),
			ActualPath: resourcePaths["actual_path"].(string),
		}
		resourcePathCollect = append(resourcePathCollect, hypervisorManagerResourcePaths)
		hypMan.ResourcePaths = resourcePathCollect
	}

	hypManError := config.ovClient.CreateHypervisorManager(hypMan)
	d.SetId(d.Get("name").(string))
	if hypManError != nil {
		d.SetId("")
		return hypManError
	}
	return resourceHypervisorManagerRead(d, meta)
}

func resourceHypervisorManagerRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	hypMan, err := config.ovClient.GetHypervisorManagerByName(d.Id())
	if err != nil || hypMan.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.Set("available_dvs_versions", hypMan.AvailableDvsVersions)
	d.Set("category", hypMan.Category)
	d.Set("created", hypMan.Created)
	d.Set("description", hypMan.Description)
	d.Set("display_name", hypMan.DisplayName)
	d.Set("e_tag", hypMan.ETAG)
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
	_ = ioutil.WriteFile("error.txt", []byte(string(len(hypMan.ResourcePaths))), 0644)

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
	d.Set("initial_scope_uris", hypMan.InitialScopeUris)
	return nil
}

func resourceHypervisorManagerUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	hypMan := ov.HypervisorManager{
		ETAG:        d.Get("etag").(string),
		URI:         utils.NewNstring(d.Get("uri").(string)),
		DisplayName: d.Get("display_name").(string),
		Name:        d.Get("name").(string),
		Username:    d.Get("username").(string),
		Password:    d.Get("password").(string),
		Port:        d.Get("port").(int),
	}

	if val, ok := d.GetOk("initial_scope_uris"); ok {
		rawInitialScopeUris := val.(*schema.Set).List()
		initialScopeUris := make([]utils.Nstring, len(rawInitialScopeUris))
		for i, raw := range rawInitialScopeUris {
			initialScopeUris[i] = utils.Nstring(raw.(string))
		}
		hypMan.InitialScopeUris = initialScopeUris
	}

	if val, ok := d.GetOk("preferences"); ok {
		preferencesList := val.(*schema.Set).List()
		for _, raw := range preferencesList {
			preferences := raw.(map[string]interface{})

			hypervisorManagerPreferences := ov.Preference{
				Type:                     preferences["type"].(string),
				VirtualSwitchType:        preferences["virtual_switch_type"].(string),
				DistributedSwitchVersion: preferences["distributed_switch_version"].(string),
				DistributedSwitchUsage:   preferences["distributed_switch_usage"].(string),
				MultiNicVMotion:          preferences["multi_nic_v_motion"].(bool),
				DrsEnabled:               preferences["drs_enabled"].(bool),
				HaEnabled:                preferences["ha_enabled"].(bool),
			}

			hypMan.Preferences = &hypervisorManagerPreferences
		}

	}
	resourcePathList := d.Get("resource_paths").(*schema.Set).List()
	resourcePathCollect := make([]ov.ResourcePath, 0)
	for _, raw := range resourcePathList {
		resourcePaths := raw.(map[string]interface{})
		hypervisorManagerResourcePaths := ov.ResourcePath{
			UserPath:   resourcePaths["user_path"].(string),
			ActualPath: resourcePaths["actual_path"].(string),
		}
		resourcePathCollect = append(resourcePathCollect, hypervisorManagerResourcePaths)
		hypMan.ResourcePaths = resourcePathCollect
	}
	err := config.ovClient.UpdateHypervisorManager(hypMan)
	if err != nil {
		return err
	}
	d.SetId(d.Get("name").(string))

	return resourceHypervisorManagerRead(d, meta)
}

func resourceHypervisorManagerDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteHypervisorManager(d.Get("name").(string))
	if err != nil {
		return err
	}
	return nil
}
