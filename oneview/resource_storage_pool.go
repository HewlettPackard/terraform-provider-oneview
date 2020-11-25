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
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceStoragePool() *schema.Resource {
	return &schema.Resource{
		Read:   resourceStoragePoolRead,
		Update: resourceStoragePoolUpdate,
		Delete: resourceStoragePoolDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"category": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allocated_capacity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"storage_system_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_capacity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"free_capacity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_managed": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"storage_pool_device_specific_attributes": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"capacity_limit": {
							Type:     schema.TypeString,
							Required: true,
						},
						"device_speed": {
							Type:     schema.TypeString,
							Required: true,
						},
						"domain": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"supported_raid_level": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uuid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_deduplication_capable": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceStoragePoolRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	storagePool, err := config.ovClient.GetStoragePoolByName(d.Id())
	if err != nil || storagePool.URI.IsNil() {
		d.SetId("")
		return nil
	}

	d.Set("category", storagePool.Category)
	d.Set("etag", storagePool.ETAG)
	d.Set("name", storagePool.Name)
	d.Set("description", storagePool.Description.String())
	d.Set("state", storagePool.State)
	d.Set("status", storagePool.Status)
	d.Set("type", storagePool.Type)
	d.Set("uri", storagePool.URI.String())
	d.Set("allocated_capacity", storagePool.AllocatedCapacity)
	d.Set("total_capacity", storagePool.TotalCapacity)
	d.Set("free_capacity", storagePool.FreeCapacity)
	d.Set("storage_system_uri", storagePool.StorageSystemUri.String())
	d.Set("is_managed", storagePool.IsManaged)

	rawdevspecificattributes := storagePool.DeviceSpecificAttributes
	devspecificattributes := make([]map[string]interface{}, 0)
	devspecificattributes = append(devspecificattributes, map[string]interface{}{
		"device_id":                rawdevspecificattributes.DeviceID,
		"capacity_limit":           rawdevspecificattributes.CapacityLimit,
		"device_speed":             rawdevspecificattributes.DeviceSpeed,
		"domain":                   rawdevspecificattributes.Domain,
		"supported_raid_level":     rawdevspecificattributes.SupportedRaidLevel,
		"uuid":                     rawdevspecificattributes.Uuid,
		"is_deduplication_capable": rawdevspecificattributes.IsDeduplicationCapable,
	})
	d.Set("storage_pool_device_specific_attributes", devspecificattributes)

	return nil
}

func resourceStoragePoolUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	storagePool := ov.StoragePool{
		URI:  utils.NewNstring(d.Get("uri").(string)),
		Name: d.Get("name").(string),
	}

	rawDeviceSpecificAttributes := d.Get("storage_pool_device_specific_attributes").(*schema.Set).List()
	deviceSpecificAttributes := ov.DeviceSpecificAttributesStoragePool{}

	for _, rawData := range rawDeviceSpecificAttributes {
		deviceSpecificAttributesItem := rawData.(map[string]interface{})
		deviceSpecificAttributes = ov.DeviceSpecificAttributesStoragePool{
			DeviceID:               deviceSpecificAttributesItem["device_id"].(string),
			CapacityLimit:          deviceSpecificAttributesItem["capacity_limit"].(string),
			DeviceSpeed:            deviceSpecificAttributesItem["device_speed"].(string),
			Domain:                 deviceSpecificAttributesItem["domain"].(string),
			SupportedRaidLevel:     deviceSpecificAttributesItem["supported_raid_level"].(string),
			IsDeduplicationCapable: deviceSpecificAttributesItem["is_deduplication_capabale"].(bool),
		}
	}

	storagePool.DeviceSpecificAttributes = &deviceSpecificAttributes

	if val, ok := d.GetOk("category"); ok {
		storagePool.Category = val.(string)
	}

	if val, ok := d.GetOk("description"); ok {
		storagePool.Description = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk("etag"); ok {
		storagePool.ETAG = val.(string)
	}

	if val, ok := d.GetOk("free_capacity"); ok {
		storagePool.FreeCapacity = val.(string)
	}

	if val, ok := d.GetOk("is_managed"); ok {
		storagePool.IsManaged = val.(bool)
	}

	if val, ok := d.GetOk("state"); ok {
		storagePool.State = val.(string)
	}

	if val, ok := d.GetOk("status"); ok {
		storagePool.Status = val.(string)
	}

	if val, ok := d.GetOk("storage_system_uri"); ok {
		storagePool.StorageSystemUri = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk("total_capacity"); ok {
		storagePool.TotalCapacity = val.(string)
	}

	if val, ok := d.GetOk("type"); ok {
		storagePool.Type = val.(string)
	}

	err := config.ovClient.UpdateStoragePool(storagePool)
	if err != nil {
		return err
	}
	d.SetId(d.Get("name").(string))

	return resourceStoragePoolRead(d, meta)
}

func resourceStoragePoolDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
