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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTask() *schema.Resource {
	return &schema.Resource{
		Read:   resourceTaskRead,
		Update: resourceTaskUpdate,
		Delete: resourceTaskDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"associated_resources": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"association_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_category": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"associated_task_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"completed_steps": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"computed_percent_complete": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"created": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"data": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"task_category": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"expected_duration": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hidden": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_cancellable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_task_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"percent_complete": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"progress_updates": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"time_stamp": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status_update": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"state_reason": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"task_errors": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"error_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"message": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_source": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"task_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"task_is_done": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"task_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"task_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"task_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_steps": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_initiated": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"wait_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceTaskRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	task, err := config.ovClient.GetTasksById("", "", "", "", d.Id())

	if err != nil || task.URI.IsNil() {
		d.SetId("")
		return nil
	}

	associatedRes := make([]map[string]interface{}, 0, 1)
	associatedRes = append(associatedRes, map[string]interface{}{
		"association_type":  task.AssociatedRes.AssociationType,
		"resource_category": task.AssociatedRes.ResourceCateogry,
		"resource_name":     task.AssociatedRes.ResourceName,
		"resource_uri":      task.AssociatedRes.ResourceURI,
	})

	progressUpdates := make([]map[string]interface{}, 0, len(task.ProgressUpdates))
	for _, update := range task.ProgressUpdates {
		progressUpdates = append(progressUpdates, map[string]interface{}{
			"time_stamp":    update.TimeStamp,
			"status_update": update.StatusUpdate,
			"id":            update.ID,
		})
	}

	data := make([]map[string]interface{}, 0, 1)
	data = append(data, map[string]interface{}{
		"task_category": task.Data.TaskCategory,
	})

	taskErrors := make([]map[string]interface{}, 0, len(task.TaskErrors))
	for _, taskError := range task.TaskErrors {
		taskErrors = append(taskErrors, map[string]interface{}{
			"error_code":   taskError.ErrorCode,
			"details":      taskError.Details,
			"message":      taskError.Message,
			"error_source": taskError.ErrorSource,
		})
	}

	d.Set("associated_resources", associatedRes)
	d.Set("progress_updates", progressUpdates)
	d.Set("data", data)
	d.Set("task_errors", taskErrors)
	d.Set("type", task.Type)
	d.Set("category", task.Category)
	d.Set("computed_percent_complete", task.ComputedPercentComplete)
	d.Set("is_cancellable", task.IsCancellable)
	d.Set("state_reason", task.StateReason)
	d.Set("percentComplete", task.PercentComplete)
	d.Set("task_state", task.TaskState)
	d.Set("task_type", task.TaskType)
	d.Set("user_initiated", task.UserInitiated)
	d.Set("name", task.Name)
	d.Set("owner", task.Owner)
	d.Set("etag", task.ETAG)
	d.Set("created", task.Created)
	d.Set("modified", task.Modified)
	d.Set("uri", task.URI)
	d.Set("task_is_done", task.TaskIsDone)
	d.Set("timeout", task.Timeout)
	d.Set("wait_time", task.WaitTime)
	d.SetId(d.Id())
	return nil
}

func resourceTaskUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.PatchTask(d.Get("uri").(string))

	if err != nil {
		d.SetId("")
		return err
	}

	d.SetId(d.Id())

	return resourceTaskRead(d, meta)
}

func resourceTaskDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
