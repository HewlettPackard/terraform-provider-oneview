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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFirmwareDrivers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirmwareDriversRead,

		Schema: map[string]*schema.Schema{

			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri": {
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
			"baseline_short_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bundle_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"bundle_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"esxi_os_driver_meta_data": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"fw_components": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"component_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"file_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sw_key_name_list": {
							Computed: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
					},
				},
			},
			"hotfixes": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hotfix_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"release_data": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"hpsum_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"iso_file_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_task_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mirror_list": {
				Computed: true,
				Type:     schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeList,
				},
			},
			"locations": {
				Computed: true,
				Type:     schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"parent_bundle": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"parent_bundle_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"release_data": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"release_data": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scope_uri": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"signature_file_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"signature_file_required": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"supported_languages": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"supported_os_list": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"sw_packages_full_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"xml_key_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceFirmwareDriversRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)
	version := d.Get("version").(string)
	id := name + version
	firmware, err := config.ovClient.GetFirmwareBaselineByNameandVersion(name, version)

	if err != nil {
		d.SetId("")
		return err
	} else if firmware.Uri.IsNil() {
		d.SetId("")
		return nil
	}
	d.Set("name", firmware.Name)
	d.Set("type", firmware.Type)
	d.Set("created", firmware.Created)
	d.Set("modified", firmware.Modified)
	d.Set("uri", firmware.Uri.String())
	d.Set("status", firmware.Status)
	d.Set("category", firmware.Category)
	d.Set("state", firmware.State)
	d.Set("etag", firmware.ETAG)
	d.Set("description", firmware.Description)
	d.Set("baseline_short_name", firmware.BaselineShortName)
	d.Set("bundle_size", firmware.BundleSize)
	d.Set("bundle_type", firmware.BundleType)
	d.Set("esxi_os_driver_meta_data", firmware.EsxiOsDriverMetaData)
	d.Set("hpsum_version", firmware.HpsumVersion)
	d.Set("iso_file_name", firmware.IsoFileName)
	d.Set("last_task_uri", firmware.LastTaskUri)
	d.Set("release_data", firmware.ReleaseDate)
	d.Set("resource_id", firmware.ResourceId)
	d.Set("resource_state", firmware.ResourceState)
	d.Set("scope_uri", firmware.ScopesUri)
	d.Set("signature_file_name", firmware.SignatureFileName)
	d.Set("signature_file_required", firmware.SignatureFileRequired)
	d.Set("supported_languages", firmware.SupportedLanguages)
	d.Set("supported_os_list", firmware.SupportedOSList)
	d.Set("sw_packages_full_path", firmware.SwPackagesFullPath)
	d.Set("uuid", firmware.Uuid)
	d.Set("version", firmware.Version)
	d.Set("xml_key_name", firmware.XmlKeyName)

	fwcomponent := make([]map[string]interface{}, 0, len(firmware.FwComponents))
	for _, component := range firmware.FwComponents {
		fwcomponent = append(fwcomponent, map[string]interface{}{
			"component_version": component.ComponentVersion,
			"file_name":         component.FileName,
			"name":              component.Name,
			"sw_key_name_list":  component.SwKeyNameList,
		})
	}
	d.Set("fw_components", fwcomponent)

	hotFixes := make([]map[string]interface{}, 0, len(firmware.Hotfixes))
	for _, hotfix := range firmware.Hotfixes {
		hotFixes = append(hotFixes, map[string]interface{}{
			"hotfix_name":  hotfix.HotfixName,
			"release_data": hotfix.ReleaseDate,
			"resource_id":  hotfix.ResourceId,
		})
	}
	d.Set("hotfixes", hotFixes)

	parentBundle := make([]map[string]interface{}, 0, 1)
	parentBundle = append(parentBundle, map[string]interface{}{
		"parent_bundle_name": firmware.ParentBundle.ParentBundleName,
		"release_data":       firmware.ParentBundle.ReleaseDate,
		"version":            firmware.ParentBundle.Version,
	})

	d.Set("parent_bundle", parentBundle)

	d.Set("locations", firmware.Locations)
	d.Set("mirror_list", firmware.Mirrorlist)

	d.SetId(id)
	return nil
}
