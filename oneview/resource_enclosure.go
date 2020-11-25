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
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceEnclosure() *schema.Resource {
	return &schema.Resource{
		Create: resourceEnclosureCreate,
		Read:   resourceEnclosureRead,
		Update: resourceEnclosureUpdate,
		Delete: resourceEnclosureDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"active_oa_preferred_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"asset_tag": {
				Type:     schema.TypeString,
				Computed: true,
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
			"device_bay_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"device_bays": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"available_for_full_height_profile": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"available_for_half_height_profile": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"bay_number": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"category": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"covered_by_device": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"covered_by_profile": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"device_presence": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"device_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enclosure_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"etag": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"model": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"modified": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"profile_uri": {
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
					},
				},
				Set: schema.HashString,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enclosure_group_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enclosure_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"force_install_firmware": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"firmware_baseline_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"force": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"fw_baseline_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fw_baseline_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"initial_scope_uris": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"interconnect_bay_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"interconnect_bays": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interconnect_bay": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"logical_interconnect_group_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
				Set: schema.HashString,
			},
			"is_fw_managed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"licensing_intent": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"op": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"part_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rack_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"refresh_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"standby_oa_preferred_ip": {
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
			"update_firmware_on": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcm_domain_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcm_domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcm_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcm_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceEnclosureCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	enclosureCreateMap := ov.EnclosureCreateMap{
		EnclosureGroupUri:    utils.NewNstring(d.Get("enclosure_group_uri").(string)),
		Hostname:             d.Get("host_name").(string),
		Username:             d.Get("user_name").(string),
		Password:             d.Get("password").(string),
		LicensingIntent:      d.Get("licensing_intent").(string),
		ForceInstallFirmware: d.Get("force_install_firmware").(bool),
		FirmwareBaselineUri:  d.Get("firmware_baseline_uri").(string),
		Force:                d.Get("force").(bool),
		UpdateFirmwareOn:     d.Get("update_firmware_on").(string),
	}

	rawInitialScopeUris := d.Get("initial_scope_uris").(*schema.Set).List()
	initialScopeUris := make([]string, len(rawInitialScopeUris))
	for i, raw := range rawInitialScopeUris {
		initialScopeUris[i] = raw.(string)
	}

	if val, ok := d.GetOk("initial_scope_uris"); ok {
		rawinitialScopeUris := val.(*schema.Set).List()
		initialScopeUris := make([]utils.Nstring, len(rawinitialScopeUris))
		for i, rawData := range rawinitialScopeUris {
			scope, _ := config.ovClient.GetScopeByName(rawData.(string))
			initialScopeUris[i] = utils.Nstring(scope.URI)
		}
		enclosureCreateMap.InitialScopeUris = initialScopeUris
	}

	enclosureError := config.ovClient.CreateEnclosure(enclosureCreateMap)
	d.SetId(d.Get("name").(string))
	if enclosureError != nil {
		d.SetId("")
		return enclosureError
	}
	return resourceEnclosureRead(d, meta)
}

func resourceEnclosureRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	enclosure, err := config.ovClient.GetEnclosureByName(d.Id())
	if err != nil || enclosure.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.Set("active_oa_preferred_ip", enclosure.ActiveOaPreferredIP)
	d.Set("asset_tag", enclosure.AssetTag)
	d.Set("category", enclosure.Category)
	d.Set("created", enclosure.Created)
	d.Set("description", enclosure.Description)
	d.Set("device_bay_count", enclosure.DeviceBayCount)
	d.Set("etag", enclosure.ETAG)
	d.Set("enclosure_group_uri", enclosure.EnclosureGroupUri.String())
	d.Set("enclosure_type", enclosure.EnclosureType)
	d.Set("fw_baseline_name", enclosure.FwBaselineName)
	d.Set("fw_baseline_uri", enclosure.FwBaselineUri.String())
	d.Set("interconnect_bay_count", enclosure.InterconnectBayCount)
	d.Set("interconnect_bays", enclosure.InterconnectBays)
	d.Set("is_fw_managed", enclosure.IsFwManaged)
	d.Set("licensing_intent", enclosure.LicensingIntent)
	d.Set("modified", enclosure.Modified)
	d.Set("name", enclosure.Name)
	d.Set("part_number", enclosure.PartNumber)
	d.Set("rack_name", enclosure.RackName)
	d.Set("refresh_state", enclosure.RefreshState)
	d.Set("serial_number", enclosure.SerialNumber)
	d.Set("standby_oa_preferred_ip", enclosure.StandbyOaPreferredIP)
	d.Set("state", enclosure.State)
	d.Set("state_reason", enclosure.StateReason)
	d.Set("status", enclosure.Status)
	d.Set("refresh_state", enclosure.RefreshState)
	d.Set("type", enclosure.Type)
	d.Set("uri", enclosure.URI.String())
	d.Set("uuid", enclosure.UUID)
	d.Set("vcm_domain_id", enclosure.VcmDomainId)
	d.Set("vcm_domain_name", enclosure.VcmDomainName)
	d.Set("vcm_mode", enclosure.VcmMode)
	d.Set("vcm_url", enclosure.VcmUrl)
	return nil
}

func resourceEnclosureUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	enclosure := ov.Enclosure{
		URI: utils.NewNstring(d.Get("uri").(string)),
	}

	err := config.ovClient.UpdateEnclosure(d.Get("op").(string), d.Get("path").(string), d.Get("value").(string), enclosure)
	if err != nil {
		return err
	}
	d.SetId(d.Get("name").(string))

	return resourceEnclosureRead(d, meta)
}

func resourceEnclosureDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteEnclosure(d.Get("name").(string))
	if err != nil {
		return err
	}
	return nil
}
