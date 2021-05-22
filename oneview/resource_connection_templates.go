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
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceConnectionTemplates() *schema.Resource {
	return &schema.Resource{
		Read:   resourceConnectionTemplatesRead,
		Update: resourceConnectionTemplatesUpdate,
		Delete: resourceConnectionTemplatesDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"bandwidth": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"maximum_bandwidth": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"typical_bandwidth": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
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
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceConnectionTemplatesRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	id := d.Id()
	cTemplate, err := config.ovClient.GetConnectionTemplateByName(id)
	if err != nil || cTemplate.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.SetId(id)
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
	bandwidthOptions := make([]map[string]interface{}, 0)

	bandwidthOptions = append(bandwidthOptions, map[string]interface{}{
		"maximum_bandwidth": cTemplate.Bandwidth.MaximumBandwidth,
		"typical_bandwidth": cTemplate.Bandwidth.TypicalBandwidth,
	})
	d.Set("bandwidth", bandwidthOptions)
	return nil
}

func resourceConnectionTemplatesUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	id := d.Id()
	updateOptions := ov.ConnectionTemplate{
		Name:        d.Get("name").(string),
		Type:        d.Get("type").(string),
		Category:    d.Get("category").(string),
		Description: utils.NewNstring(d.Get("description").(string)),
		ETAG:        d.Get("etag").(string),
		Modified:    d.Get("modified").(string),
		State:       d.Get("state").(string),
		Status:      d.Get("status").(string),
		URI:         utils.NewNstring(d.Get("uri").(string)),
		Created:     d.Get("created").(string),
	}

	rawBandwidthOptions := d.Get("bandwidth").(*schema.Set).List()
	for _, val := range rawBandwidthOptions {
		rawval := val.(map[string]interface{})
		BandwidthOptions := ov.BandwidthType{
			MaximumBandwidth: rawval["maximum_bandwidth"].(int),
			TypicalBandwidth: rawval["typical_bandwidth"].(int),
		}
		updateOptions.Bandwidth = BandwidthOptions
	}
	_, err := config.ovClient.UpdateConnectionTemplate(id, updateOptions)
	if err != nil {
		return err
	}
	d.SetId(id)

	return resourceConnectionTemplatesRead(d, meta)
}

func resourceConnectionTemplatesDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
