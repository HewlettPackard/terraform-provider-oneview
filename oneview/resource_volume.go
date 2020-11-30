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
)

func resourceVolume() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolumeCreate,
		Read:   resourceVolumeRead,
		Update: resourceVolumeUpdate,
		Delete: resourceVolumeDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_volume_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"device_specific_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"copy_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_compressed": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"is_deduplicated": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"snapshot_pool_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"properties": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"storage_pool": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"size": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"provisioning_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"data_transfer_limit": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"is_deduplicated": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"is_encrypted": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"is_pinned": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"is_compressed": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"data_protection_level": {
							Type:     schema.TypeString,
							Optional: true,
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
				Optional: true,
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
				Optional: true,
			},
			"provisioning_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allocated_capacity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"requesting_refresh": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"template_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
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
			"volume_template_uri": {
				Type:     schema.TypeString,
				Computed: true,
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
		},
	}
}

func resourceVolumeCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	volume := ov.StorageVolume{}

	properties := d.Get("properties").(*schema.Set).List()[0].(map[string]interface{})
	volumeProperties := ov.Properties{
		Storagepool:         utils.NewNstring(properties["storage_pool"].(string)),
		Name:                d.Get("name").(string),
		Size:                properties["size"].(int),
		ProvisioningType:    properties["provisioning_type"].(string),
		DataTransferLimit:   properties["data_transfer_limit"].(int),
		DataProtectionLevel: properties["data_protection_level"].(string),
		IsDeduplicated:      properties["is_deduplicated"].(bool),
		IsEncrypted:         properties["is_encrypted"].(bool),
		IsPinned:            properties["is_pinned"].(bool),
		IsCompressed:        properties["is_compressed"].(bool),
	}
	volume.Properties = &volumeProperties
	volume.TemplateURI = utils.NewNstring(d.Get("template_uri").(string))

	if value, exist := d.GetOk("is_permanent"); exist {
		val := value.(bool)
		volume.IsPermanent = &val
	}
	if val, ok := d.GetOk("initial_scope_uris"); ok {
		rawInitialScopeUris := val.(*schema.Set).List()
		initialScopeUris := make([]utils.Nstring, len(rawInitialScopeUris))
		for i, raw := range rawInitialScopeUris {
			initialScopeUris[i] = utils.Nstring(raw.(string))
		}
		volume.InitialScopeUris = initialScopeUris
	}
	err := config.ovClient.CreateStorageVolume(volume)
	d.SetId(d.Get("name").(string))

	if err != nil {
		d.SetId("")
		return err
	}

	return resourceVolumeRead(d, meta)

}

func resourceVolumeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	storageVolume, err := config.ovClient.GetStorageVolumeByName(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	if storageVolume.URI.IsNil() {
		d.SetId("")
		return nil
	}

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
	d.Set("state", storageVolume.State)
	d.Set("status", storageVolume.Status)
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

func resourceVolumeUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	volume := ov.StorageVolume{}
	isPermanent := d.Get("is_permanent").(bool)
	isShareable := d.Get("is_shareable").(bool)

	volume.Name = d.Get("name").(string)
	volume.ProvisioningTypeForUpdate = d.Get("provisioning_type").(string)
	volume.Description = utils.NewNstring(d.Get("description").(string))
	volume.IsPermanent = &isPermanent
	volume.URI = utils.NewNstring(d.Get("uri").(string))
	volume.IsShareable = &isShareable

	deviceSpecificAttributesTemplate := d.Get("device_specific_attributes").(*schema.Set).List()[0].(map[string]interface{})

	deviceSpecificAttributes := ov.DeviceSpecificAttributes{
		CopyState:      deviceSpecificAttributesTemplate["copy_state"].(string),
		IsCompressed:   deviceSpecificAttributesTemplate["is_compressed"].(bool),
		IsDeduplicated: deviceSpecificAttributesTemplate["is_deduplicated"].(bool),
	}
	if val, exist := deviceSpecificAttributesTemplate["snapshot_pool_uri"]; exist {
		if val.(string) != "null" {
			deviceSpecificAttributes.SnapshotPoolUri = utils.NewNstring(val.(string))
		}
	}
	volume.DeviceSpecificAttributes = &deviceSpecificAttributes
	volume.Category = d.Get("category").(string)
	volume.Type = d.Get("type").(string)
	volume.ETAG = d.Get("etag").(string)
	volume.ProvisionedCapacity = d.Get("provisioned_capacity").(string)
	volume.TemplateVersion = d.Get("template_version").(string)
	volume.VolumeTemplateUri = utils.NewNstring(d.Get("volume_template_uri").(string))

	err := config.ovClient.UpdateStorageVolume(volume)
	d.SetId(d.Get("name").(string))

	if err != nil {
		d.SetId("")
		return err
	}

	return resourceVolumeRead(d, meta)
}

func resourceVolumeDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteStorageVolume(d.Get("name").(string))
	if err != nil {
		return err
	}
	return nil
}
