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
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceStorageSystem() *schema.Resource {
	return &schema.Resource{
		Create: resourceStorageSystemCreate,
		Read:   resourceStorageSystemRead,
		Update: resourceStorageSystemUpdate,
		Delete: resourceStorageSystemDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"credentials": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type:     schema.TypeString,
							Required: true,
						},
						"password": {
							Type:     schema.TypeString,
							Optional: true,
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
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"family": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"storage_pools_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_capacity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ports": {
				Optional: true,
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"partner_port": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"mode": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
					},
				},
			},
			"storage_system_device_specific_attributes": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"firmware": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"model": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"managed_domain": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"managed_pool": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"domain": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"device_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"free_capacity": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"raid_level": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"total_capacity": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceStorageSystemCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	storageSystem := ov.StorageSystem{
		Hostname: d.Get("hostname").(string),
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		Family:   d.Get("family").(string),
	}

	if val, ok := d.GetOk("name"); ok {
		storageSystem.Name = val.(string)
	}

	storageSystemError := config.ovClient.CreateStorageSystem(storageSystem)
	d.SetId(d.Get("hostname").(string))
	if storageSystemError != nil {
		d.SetId("")
		return storageSystemError
	}
	return resourceStorageSystemRead(d, meta)
}

func resourceStorageSystemRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	storageSystemList, err := config.ovClient.GetStorageSystems(fmt.Sprintf("hostname matches '%s'", d.Id()), "")
	if err != nil || len(storageSystemList.Members) < 1 {
		d.SetId("")
		return nil
	}

	storageSystem := storageSystemList.Members[0]
	if storageSystem.URI.IsNil() {
		d.SetId("")
		return nil
	}

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
	d.Set("mode", storageSystem.Mode)
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

func resourceStorageSystemUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	storageSystem := ov.StorageSystem{
		Hostname: d.Get("hostname").(string),
		URI:      utils.NewNstring(d.Get("uri").(string)),
		Name:     d.Get("name").(string),
	}

	rawCredentials := d.Get("credentials").(*schema.Set).List()
	credentials := ov.Credentials{}
	for _, raw := range rawCredentials {
		credentialsItem := raw.(map[string]interface{})
		credentials = ov.Credentials{
			Username: credentialsItem["username"].(string),
			Password: credentialsItem["password"].(string)}
	}

	storageSystem.Credentials = &credentials

	rawManagedPools := d.Get("managed_pool").(*schema.Set).List()
	managedPools := make([]ov.ManagedPools, 0)

	for _, rawMP := range rawManagedPools {
		managedPoolItem := rawMP.(map[string]interface{})
		managedPools = append(managedPools, ov.ManagedPools{
			Name:          managedPoolItem["name"].(string),
			Domain:        managedPoolItem["domain"].(string),
			DeviceType:    managedPoolItem["device_type"].(string),
			FreeCapacity:  managedPoolItem["free_capacity"].(string),
			RaidLevel:     managedPoolItem["raid_level"].(string),
			Totalcapacity: managedPoolItem["total_capacity"].(string)})
	}

	rawDeviceSpecificAttributes := d.Get("storage_system_device_specific_attributes").(*schema.Set).List()
	deviceSpecificAttributes := ov.StorageSystemDeviceSpecificAttributes{}

	for _, rawData := range rawDeviceSpecificAttributes {
		deviceSpecificAttributesItem := rawData.(map[string]interface{})
		deviceSpecificAttributes = ov.StorageSystemDeviceSpecificAttributes{
			Firmware:      deviceSpecificAttributesItem["firmware"].(string),
			Model:         deviceSpecificAttributesItem["model"].(string),
			ManagedPools:  managedPools,
			ManagedDomain: deviceSpecificAttributesItem["managed_domain"].(string)}
	}

	storageSystem.StorageSystemDeviceSpecificAttributes = &deviceSpecificAttributes

	rawPorts := d.Get("ports").(*schema.Set).List()
	ports := make([]ov.Ports, 0)
	for _, rawPort := range rawPorts {
		portsItem := rawPort.(map[string]interface{})
		ports = append(ports, ov.Ports{
			Id:   portsItem["id"].(string),
			Mode: portsItem["mode"].(string),
			PortDeviceSpecificAttributes: ov.PortDeviceSpecificAttributes{
				PartnerPort: portsItem["partner_port"].(string)}})
	}

	storageSystem.Ports = ports

	if val, ok := d.GetOk("category"); ok {
		storageSystem.Category = val.(string)
	}

	if val, ok := d.GetOk("description"); ok {
		storageSystem.Description = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk("etag"); ok {
		storageSystem.ETAG = val.(string)
	}

	if val, ok := d.GetOk("family"); ok {
		storageSystem.Family = val.(string)
	}

	if val, ok := d.GetOk("state"); ok {
		storageSystem.State = val.(string)
	}

	if val, ok := d.GetOk("status"); ok {
		storageSystem.Status = val.(string)
	}

	if val, ok := d.GetOk("storage_pools_uri"); ok {
		storageSystem.StoragePoolsUri = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk("total_capacity"); ok {
		storageSystem.TotalCapacity = val.(string)
	}

	if val, ok := d.GetOk("type"); ok {
		storageSystem.Type = val.(string)
	}

	err := config.ovClient.UpdateStorageSystem(storageSystem)
	if err != nil {
		return err
	}
	d.SetId(d.Get("hostname").(string))

	return resourceStorageSystemRead(d, meta)
}

func resourceStorageSystemDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteStorageSystem(d.Get("name").(string))
	if err != nil {
		return err
	}
	return nil
}
