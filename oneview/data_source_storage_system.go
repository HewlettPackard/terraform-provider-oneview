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

func dataSourceStorageSystem() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceStorageSystemRead,

		Schema: map[string]*schema.Schema{
			"hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"credentials": {
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"password": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
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
			"family": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_pools_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_capacity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ports": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"partner_port": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"storage_system_device_specific_attributes": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"firmware": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"model": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"managed_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"managed_pool": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"device_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"free_capacity": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"raid_level": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"total_capacity": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceStorageSystemRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	id := d.Get("name").(string)

	storageSystem, err := config.ovClient.GetStorageSystemByName(id)
	if err != nil || storageSystem.URI.IsNil() {
		d.SetId("")
		return nil
	}

	d.SetId(id)
	d.Set("hostname", storageSystem.Hostname)
	d.Set("category", storageSystem.Category)
	d.Set("etag", storageSystem.ETAG)
	d.Set("name", storageSystem.Name)
	d.Set("description", storageSystem.Description.String())
	d.Set("state", storageSystem.State)
	d.Set("status", storageSystem.Status)
	d.Set("type", storageSystem.Type)
	d.Set("uri", storageSystem.URI.String())
	d.Set("family", storageSystem.Family)
	d.Set("storage_pools_uri", storageSystem.StoragePoolsUri.String())
	d.Set("total_capacity", storageSystem.TotalCapacity)

	rawcredentials := storageSystem.Credentials
	credentials := make([]map[string]interface{}, 0)
	credentials = append(credentials, map[string]interface{}{
		"username": rawcredentials.Username,
		"password": rawcredentials.Password})
	d.Set("credentials", credentials)

	rawports := storageSystem.Ports
	ports := make([]map[string]interface{}, 0, len(rawports))
	for _, port := range rawports {
		ports = append(ports, map[string]interface{}{
			"id":           port.Id,
			"mode":         port.Mode,
			"partner_port": port.PortDeviceSpecificAttributes.PartnerPort})
	}
	d.Set("ports", ports)

	rawmp := storageSystem.StorageSystemDeviceSpecificAttributes.ManagedPools
	managedPools := make([]map[string]interface{}, 0)
	for _, mp := range rawmp {
		managedPools = append(managedPools, map[string]interface{}{
			"name":           mp.Name,
			"domain":         mp.Domain,
			"device_type":    mp.DeviceType,
			"free_capacity":  mp.FreeCapacity,
			"raid_level":     mp.RaidLevel,
			"total_capacity": mp.Totalcapacity})
	}
	d.Set("managed_pool", managedPools)

	rawssda := storageSystem.StorageSystemDeviceSpecificAttributes
	deviceSpecificAttributes := make([]map[string]interface{}, 0)
	deviceSpecificAttributes = append(deviceSpecificAttributes, map[string]interface{}{
		"firmware":       rawssda.Firmware,
		"model":          rawssda.Model,
		"managed_domain": rawssda.ManagedDomain})
	d.Set("storage_system_device_specific_attributes", deviceSpecificAttributes)

	return nil
}
