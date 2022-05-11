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
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLabel() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLabelRead,
		Schema: map[string]*schema.Schema{
			"resource_uri": {
				Type:     schema.TypeString,
				Required: true,
			},
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified": {
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
			"labels": {
				Computed: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceLabelRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	label, err := config.ovClient.GetAssignedLabels(utils.Nstring(d.Get("resource_uri").(string)))
	if err != nil {
		d.SetId("")
		return err
	} else if label.Uri.IsNil() {
		d.SetId("")
		return nil
	}
	d.Set("category", label.Category)
	d.Set("created", label.Created)
	d.Set("etag", label.ETAG.String())
	d.Set("modfied", label.Modified)
	d.Set("type", label.Type)
	d.Set("uri", label.Uri.String())
	d.Set("resource_uri", label.ResourceUri.String())
	// Sets Labels
	if len(label.Labels) != 0 {
		rlabel := make([]map[string]interface{}, 0, len(label.Labels))
		for _, raw := range label.Labels {
			rlabel = append(rlabel, map[string]interface{}{
				"name": raw.Name,
				"uri":  raw.Uri,
			})
		}
		d.Set("labels", rlabel)
	}
	d.SetId(label.ResourceUri.String())
	return nil
}
