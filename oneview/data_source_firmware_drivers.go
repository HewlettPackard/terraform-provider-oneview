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
				Computed: true,
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
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"locations": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
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
				Type:     schema.TypeSet,
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
				Computed: true,
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
	cTemplate, err := config.ovClient.GetConnectionTemplateByName(name)
	if err != nil || cTemplate.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.Set("name", cTemplate.Name)
	d.Set("type", cTemplate.Type)
	d.Set("created", cTemplate.Created)
	d.Set("modified", cTemplate.Modified)
	d.Set("uri", cTemplate.URI.String())
	d.Set("status", cTemplate.Status)
	d.Set("category", cTemplate.Category)
	d.Set("state", cTemplate.State)
	d.Set("etag", cTemplate.ETAG)
	d.Set("description", cTemplate.Description)
	bandwidth := make([]map[string]interface{}, 0, 1)

	bandwidth = append(bandwidth, map[string]interface{}{
		"maximum_bandwidth": cTemplate.Bandwidth.MaximumBandwidth,
		"typical_bandwidth": cTemplate.Bandwidth.TypicalBandwidth,
	})

	d.Set("bandwidth", bandwidth)
	d.SetId(name)
	return nil
}
