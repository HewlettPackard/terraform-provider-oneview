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

func dataSourceStorageAttachment() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceStorageAttachmentRead,

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
			"storage_system_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_volume_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host": {
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"os": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"paths": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"transport": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceStorageAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	id := d.Get("name").(string)

	storageAttachment, err := config.ovClient.GetStorageAttachmentById(id)
	if err != nil || storageAttachment.URI.IsNil() {
		d.SetId("")
		return nil
	}

	d.SetId(id)
	d.Set("category", storageAttachment.Category)
	d.Set("etag", storageAttachment.ETAG)
	d.Set("name", storageAttachment.Name)
	d.Set("description", storageAttachment.Description.String())
	d.Set("state", storageAttachment.State)
	d.Set("status", storageAttachment.Status)
	d.Set("etag", storageAttachment.ETAG)
	d.Set("type", storageAttachment.Type)
	d.Set("uri", storageAttachment.URI.String())
	d.Set("storage_system_uri", storageAttachment.StorageSystemUri.String())
	d.Set("storage_volume_uri", storageAttachment.StorageVolumeUri.String())

	rawpaths := storageAttachment.Paths
	paths := make([]map[string]interface{}, 0, len(rawpaths))
	for _, path := range rawpaths {
		paths = append(paths, map[string]interface{}{
			"connection_name": path.ConnectionName,
			"is_enabled":      path.IsEnabled,
			"transport":       path.Transport,
		})
	}
	d.Set("paths", paths)

	rawhost := storageAttachment.Host
	hosts := make([]map[string]interface{}, 0)
	hosts = append(hosts, map[string]interface{}{
		"name": rawhost.Name,
		"os":   rawhost.Os,
	})
	d.Set("host", hosts)

	return nil
}
