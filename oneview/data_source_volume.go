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
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceVolume() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolumeRead,

		Schema: map[string]*schema.Schema{
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"device_specific_attributes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"copy_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_compressed": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_deduplicated": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"snapshot_pool_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_permanent": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_shareable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"provisioned_capacity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"provisioning_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"requesting_refresh": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"storage_pool_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"template_version": {
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
			"volume_set_uris": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"volume_template_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceVolumeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	storageVolume, err := config.ovClient.GetStorageVolumeByName(d.Get("name").(string))
	if err != nil {
		d.SetId("")
		return nil
	}

	if storageVolume.URI.IsNil() {
		d.SetId("")
		return nil
	}

	d.SetId(d.Get("name").(string))
	d.Set("category", storageVolume.Category)
	d.Set("description", storageVolume.Description)
	d.Set("allocated_capacity", storageVolume.AllocatedCapacity)
	d.Set("device_volume_name", storageVolume.DeviceVolumeName)

	deviceSpecificAttributesTemplates := make([]map[string]interface{}, 0, 1)
	deviceSpecificAttributesTemplates = append(deviceSpecificAttributesTemplates, map[string]interface{}{
		"copy_state":        storageVolume.DeviceSpecificAttributes.CopyState,
		"is_compressed":     storageVolume.DeviceSpecificAttributes.IsCompressed,
		"is_deduplicated":   storageVolume.DeviceSpecificAttributes.IsDeduplicated,
		"snapshot_pool_uri": storageVolume.DeviceSpecificAttributes.SnapshotPoolUri.String(),
	})

	d.Set("device_specific_attributes", &deviceSpecificAttributesTemplates)
	d.Set("etag", storageVolume.ETAG)
	d.Set("is_permanent", storageVolume.IsPermanent)
	d.Set("is_shareable", storageVolume.IsShareable)
	d.Set("name", storageVolume.Name)
	d.Set("provisioned_capacity", storageVolume.ProvisionedCapacity)
	d.Set("provisioning_type", storageVolume.ProvisioningTypeForUpdate)
	d.Set("requesting_refresh", storageVolume.RequestingRefresh)
	d.Set("storage_pool_uri", storageVolume.StoragePoolUri.String())
	d.Set("template_version", storageVolume.TemplateVersion)
	d.Set("type", storageVolume.Type)
	d.Set("uri", storageVolume.URI.String())
	d.Set("volume_template_uri", storageVolume.VolumeTemplateUri.String())

	return nil
}
