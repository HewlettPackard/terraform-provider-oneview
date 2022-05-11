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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConnectionTemplates() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceConnectionTemplatesRead,

		Schema: map[string]*schema.Schema{
			"bandwidth": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"maximum_bandwidth": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"typical_bandwidth": {
							Type:     schema.TypeInt,
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
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Optional: true,
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
		},
	}
}

func dataSourceConnectionTemplatesRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	var (
		cTemplate ov.ConnectionTemplate
		err       error
	)

	// reads connection template via uri or name
	if name, ok := d.Get("name").(string); ok {
		cTemplate, err = config.ovClient.GetConnectionTemplateByName(name)
	} else if uri, ok := d.Get("uri").(string); ok {
		cTemplate, err = config.ovClient.GetConnectionTemplateByURI(utils.Nstring(uri))
	}
	if err != nil {
		d.SetId("")
		return err
	} else if cTemplate.URI.IsNil() {
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
	d.SetId(d.Get("uri").(string))
	return nil
}
