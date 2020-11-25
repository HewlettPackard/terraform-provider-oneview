// (C) Copyright 2019 Hewlett Packard Enterprise Development LP
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

func dataSourceDeploymentPlan() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDeploymentPlanRead,

		Schema: map[string]*schema.Schema{
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_attributes": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"constraints": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"editable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"visible": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"golden_image_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hp_provided": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"dp_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"oe_build_plan_uri": {
				Type:     schema.TypeString,
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
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceDeploymentPlanRead(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*Config)

	id := d.Get("name").(string)
	deploymentPlan, err := config.i3sClient.GetDeploymentPlanByName(id)
	if err != nil || deploymentPlan.URI.IsNil() {
		d.SetId("")
		return nil
	}

	d.SetId(id)
	d.Set("category", deploymentPlan.Category)
	d.Set("custom_attributes", deploymentPlan.CustomAttributes)
	d.Set("description", deploymentPlan.Description.String())
	d.Set("etag", deploymentPlan.ETAG)
	d.Set("golden_image_uri", deploymentPlan.GoldenImageUri)
	d.Set("hp_provided", deploymentPlan.HPProvided)
	d.Set("dp_id", deploymentPlan.ID)
	d.Set("name", deploymentPlan.Name)
	d.Set("oe_build_plan_uri", deploymentPlan.OEBuildPlanURI.String())
	d.Set("status", deploymentPlan.Status)
	d.Set("type", deploymentPlan.Type)
	d.Set("uri", deploymentPlan.URI.String())

	return nil
}
