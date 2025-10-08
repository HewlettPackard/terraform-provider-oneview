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
	"fmt"
	"strings"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceServerProfileAsyncTask() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServerProfileAsyncTaskRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"task_id", "task_uri_input"},
				Description:   "Name of the server profile to find associated task",
			},
			"task_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"name", "task_uri_input"},
				Description:   "Task ID to look up directly",
			},
			"task_uri_input": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"name", "task_id"},
				Description:   "Full task URI to look up directly (e.g., /rest/tasks/task-id)",
			},
			"task_uri": {
				Type:     schema.TypeString,
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
			"task_category": {
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
			"user": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"percent_complete": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"expected_duration": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"computed_percent_complete": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"progress_updates": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"timestamp": {
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
			"associated_resources": {
				Type:     schema.TypeList,
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
			"server_profile_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_profile_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceServerProfileAsyncTaskRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)
	taskID := d.Get("task_id").(string)
	taskURI := d.Get("task_uri_input").(string)

	// Ensure one of name, task_id, or task_uri_input is provided
	if name == "" && taskID == "" && taskURI == "" {
		return fmt.Errorf("Either 'name' (server profile name), 'task_id', or 'task_uri_input' must be provided")
	}

	var task ov.Task
	var err error

	if taskURI != "" {
		// Extract task ID from URI (e.g., /rest/tasks/abc-123 -> abc-123)
		// Split by "/" and take the last part
		parts := strings.Split(taskURI, "/")
		if len(parts) == 0 {
			return fmt.Errorf("Invalid task URI format: %s", taskURI)
		}
		extractedTaskID := parts[len(parts)-1]
		
		// Direct task lookup by extracted ID
		task, err = config.ovClient.GetTasksById("", "", "", "", extractedTaskID)
		if err != nil {
			return fmt.Errorf("Could not find task with URI '%s': %s", taskURI, err)
		}
	} else if taskID != "" {
		// Direct task lookup by ID
		task, err = config.ovClient.GetTasksById("", "", "", "", taskID)
		if err != nil {
			return fmt.Errorf("Could not find task with ID '%s': %s", taskID, err)
		}
	} else {
		// Lookup by server profile name
		serverProfile, err := config.ovClient.GetProfileByName(name)
		if err != nil || serverProfile.URI.IsNil() {
			return fmt.Errorf("Could not find server profile with name '%s': %s", name, err)
		}

		// Set server profile information
		d.Set("server_profile_uri", serverProfile.URI.String())
		d.Set("server_profile_name", serverProfile.Name)

		// Get tasks related to this server profile
		// We'll look for tasks that have this server profile as an associated resource
		taskCollection, err := config.ovClient.GetTasks("", "", "", "", "", "name:desc")
		if err != nil {
			return fmt.Errorf("Error getting tasks: %s", err)
		}

		var relatedTask *ov.Task
		// Find the most recent task related to this server profile
		for _, t := range taskCollection.Members {
			if t.AssociatedRes.ResourceURI.String() == serverProfile.URI.String() {
				relatedTask = &t
				break
			}
		}

		if relatedTask == nil {
			return fmt.Errorf("No task found for server profile '%s'", name)
		}
		task = *relatedTask
	}

	// Set task information
	d.Set("task_uri", task.URI.String())
	d.Set("task_state", task.TaskState)
	d.Set("task_status", task.TaskStatus)
	d.Set("task_type", task.Type)
	d.Set("task_category", task.Category)
	d.Set("created", task.Created)
	d.Set("modified", task.Modified)
	d.Set("user", task.User)
	d.Set("owner", task.Owner)
	d.Set("percent_complete", task.PercentComplete)
	d.Set("expected_duration", task.ExpectedDuration)
	d.Set("computed_percent_complete", task.ComputedPercentComplete)

	// Set progress updates
	progressUpdates := make([]map[string]interface{}, 0)
	for _, update := range task.ProgressUpdates {
		progressUpdate := map[string]interface{}{
			"timestamp":     update.TimeStamp,
			"status_update": update.StatusUpdate,
			"id":            update.ID,
		}
		progressUpdates = append(progressUpdates, progressUpdate)
	}
	d.Set("progress_updates", progressUpdates)

	// Set associated resources
	associatedResources := make([]map[string]interface{}, 0)
	assocResource := map[string]interface{}{
		"association_type":  task.AssociatedRes.AssociationType,
		"resource_category": task.AssociatedRes.ResourceCateogry,
		"resource_name":     task.AssociatedRes.ResourceName,
		"resource_uri":      task.AssociatedRes.ResourceURI,
	}
	associatedResources = append(associatedResources, assocResource)
	d.Set("associated_resources", associatedResources)

	// If we looked up by task ID, try to get server profile info from the task
	if taskID != "" && task.AssociatedRes.ResourceURI.String() != "" {
		d.Set("server_profile_uri", task.AssociatedRes.ResourceURI.String())
		d.Set("server_profile_name", task.AssociatedRes.ResourceName)
	}

	d.SetId(task.URI.String())

	return nil
}