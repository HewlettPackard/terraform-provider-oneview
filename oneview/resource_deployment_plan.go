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
	"github.com/HewlettPackard/oneview-golang/i3s"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDeploymentPlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeploymentPlanCreate,
		Read:   resourceDeploymentPlanRead,
		Update: resourceDeploymentPlanUpdate,
		Delete: resourceDeploymentPlanDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_attributes": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"constraints": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"editable": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"visible": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"golden_image_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hp_provided": {
				Type:     schema.TypeBool,
				Optional: true,
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
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceDeploymentPlanCreate(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*Config)

	deploymentPlan := i3s.DeploymentPlan{
		Name:           d.Get("name").(string),
		OEBuildPlanURI: utils.NewNstring(d.Get("oe_build_plan_uri").(string)),
		Type:           d.Get("type").(string)}

	if val, ok := d.GetOk("description"); ok {
		deploymentPlan.Description = utils.NewNstring(val.(string))
	}
	if val, ok := d.GetOk("custom_attributes"); ok {
		rawCutsomAttributes := val.(*schema.Set).List()
		customAttributes := make([]i3s.CustomAttribute, 0)
		for _, rawData := range rawCutsomAttributes {
			item := rawData.(map[string]interface{})
			customAttribute := i3s.CustomAttribute{}

			if item["constraints"] != nil {
				customAttribute.Constraints = item["constraints"].(string)
			}
			if item["description"] != nil {
				customAttribute.Description = item["description"].(string)
			}
			if item["editable"] != nil {
				customAttribute.Editable = item["editable"].(bool)
			}
			if item["id"] != nil {
				customAttribute.ID = item["id"].(string)
			}
			if item["name"] != nil {
				customAttribute.Name = item["name"].(string)
			}
			if item["type"] != nil {
				customAttribute.Type = item["type"].(string)
			}
			if item["value"] != nil {
				customAttribute.Value = item["value"].(string)
			}
			if item["visible"] != nil {
				customAttribute.Visible = item["visible"].(bool)
			}

			customAttributes = append(customAttributes, customAttribute)
		}
		deploymentPlan.CustomAttributes = customAttributes
	}

	if val, ok := d.GetOk("golden_image_uri"); ok {
		deploymentPlan.GoldenImageUri = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk("hp_provided"); ok {
		deploymentPlan.HPProvided = val.(bool)
	}

	err := config.i3sClient.CreateDeploymentPlan(deploymentPlan)
	d.SetId(d.Get("name").(string))
	if err != nil {
		d.SetId("")
		return err
	}

	return resourceDeploymentPlanRead(d, meta)
}

func resourceDeploymentPlanRead(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*Config)

	deploymentPlan, err := config.i3sClient.GetDeploymentPlanByName(d.Id())
	if err != nil || deploymentPlan.URI.IsNil() {
		d.SetId("")
		return nil
	}

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

func resourceDeploymentPlanUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	deploymentPlan := i3s.DeploymentPlan{
		Name:           d.Get("name").(string),
		OEBuildPlanURI: utils.NewNstring(d.Get("oe_build_plan_uri").(string)),
		Type:           d.Get("type").(string),
		URI:            utils.NewNstring(d.Get("uri").(string))}

	if val, ok := d.GetOk("description"); ok {
		deploymentPlan.Description = utils.NewNstring(val.(string))
	}
	if val, ok := d.GetOk("custom_attributes"); ok {
		rawCutsomAttributes := val.(*schema.Set).List()
		customAttributes := make([]i3s.CustomAttribute, 0)
		for _, rawData := range rawCutsomAttributes {
			item := rawData.(map[string]interface{})
			var customAttribute i3s.CustomAttribute

			if item["constraints"] != nil {
				customAttribute.Constraints = item["constraints"].(string)
			}
			if item["description"] != nil {
				customAttribute.Description = item["description"].(string)
			}
			if item["editable"] != nil {
				customAttribute.Editable = item["editable"].(bool)
			}
			if item["id"] != nil {
				customAttribute.ID = item["id"].(string)
			}
			if item["name"] != nil {
				customAttribute.Name = item["name"].(string)
			}
			if item["type"] != nil {
				customAttribute.Type = item["type"].(string)
			}
			if item["value"] != nil {
				customAttribute.Value = item["value"].(string)
			}
			if item["visible"] != nil {
				customAttribute.Visible = item["visible"].(bool)
			}

			customAttributes = append(customAttributes, customAttribute)
		}
		deploymentPlan.CustomAttributes = customAttributes
	}

	if val, ok := d.GetOk("golden_image_uri"); ok {
		deploymentPlan.GoldenImageUri = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk("hp_provided"); ok {
		deploymentPlan.HPProvided = val.(bool)
	}

	err := config.i3sClient.UpdateDeploymentPlan(deploymentPlan)
	if err != nil {
		d.SetId("")
		return err
	}

	d.SetId(d.Get("name").(string))
	return resourceDeploymentPlanRead(d, meta)
}

func resourceDeploymentPlanDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.i3sClient.DeleteDeploymentPlan(d.Get("name").(string))
	if err != nil {
		return err
	}
	return nil
}
