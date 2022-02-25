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

func resourceLabel() *schema.Resource {
	return &schema.Resource{
		Create: resourceLabelCreate,
		Read:   resourceLabelRead,
		Update: resourceLabelUpdate,
		Delete: resourceLabelDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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
				Optional: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceLabelCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	AssignedLabel := ov.AssignedLabel{
		ResourceUri: utils.Nstring(d.Get("resource_uri").(string)),
	}

	if val, ok := d.GetOk("labels"); ok {
		rawlabels := val.([]interface{})
		label := []ov.Label{}
		for _, rawlabel := range rawlabels {
			rlabel := rawlabel.(map[string]interface{})
			label = append(label, ov.Label{
				Name: rlabel["name"].(string),
				Uri:  utils.Nstring(rlabel["uri"].(string)),
			})
		}
		AssignedLabel.Labels = label
	}
	_, err := config.ovClient.CreateLabel(AssignedLabel)
	d.SetId(d.Get("resource_uri").(string))
	if err != nil {
		d.SetId("")
		return err
	}
	return resourceLabelRead(d, meta)
}

func resourceLabelRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	label, err := config.ovClient.GetAssignedLabels(utils.Nstring(d.Id()))
	if err != nil || label.Uri.IsNil() {
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

	if len(label.Labels) != 0 {
		rlabel := make([]map[string]interface{}, len(label.Labels), len(label.Labels))
		for i, raw := range label.Labels {
			rlabel[i] = map[string]interface{}{
				"name": raw.Name,
				"uri":  raw.Uri,
			}
		}
		d.Set("labels", rlabel)
	}
	return nil
}

func resourceLabelUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	assignedLabel := ov.AssignedLabel{
		ResourceUri: utils.Nstring(d.Get("resource_uri").(string)),
		Category:    d.Get("category").(string),
		Created:     d.Get("created").(string),
		ETAG:        utils.Nstring(d.Get("etag").(string)),
		Modified:    d.Get("modified").(string),
		Uri:         utils.Nstring(d.Get("uri").(string)),
		Type:        d.Get("type").(string),
	}

	if val, ok := d.GetOk("labels"); ok {
		rawlabels := val.([]interface{})
		label := []ov.Label{}
		for _, rawlabel := range rawlabels {
			rlabel := rawlabel.(map[string]interface{})
			label = append(label, ov.Label{
				Name: rlabel["name"].(string),
				Uri:  utils.Nstring(rlabel["uri"].(string)),
			})
		}
		assignedLabel.Labels = label
	}

	response, err := config.ovClient.UpdateAssignedLabels(assignedLabel)
	if err != nil {
		d.SetId("")
		return err
	}
	d.SetId(response.ResourceUri.String())
	return resourceLabelRead(d, meta)
}

func resourceLabelDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteAssignedLabel(d.Id())
	if err != nil {
		return err
	}
	return nil
}
