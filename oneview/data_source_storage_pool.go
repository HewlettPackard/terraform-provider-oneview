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

func dataSourceStoragePool() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceStoragePoolRead,

		Schema: map[string]*schema.Schema{
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
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
			"allocated_capacity": {
				Type:     schema.TypeString,
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
				Computed: true,
			},
			"storage_pool_device_specific_attributes": {
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"capacity_limit": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"device_speed": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"supported_raid_level": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_deduplication_capable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceStoragePoolRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	storagePool, err := config.ovClient.GetStoragePoolByName(d.Get("name").(string))
	if err != nil || storagePool.URI.IsNil() {
		d.SetId("")
		return nil
	}

	d.SetId(d.Get("name").(string))
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
