// (C) Copyright 2020 Hewlett Packard Enterprise Development LP
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

func resourceStorageVolumeTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceStorageVolumeTemplateCreate,
		Read:   resourceStorageVolumeTemplateRead,
		Update: resourceStorageVolumeTemplateUpdate,
		Delete: resourceStorageVolumeTemplateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compatible_storage_systems_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_root": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"scopes_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"state": {
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
			"initial_scope_uris": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"root_template_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"storage_pool_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"family": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tp_name": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"title": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"required": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"min_length": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"max_length": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"enum": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"minimum": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tp_storage_pool": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"title": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"required": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"min_length": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"max_length": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"enum": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"default": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"minimum": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tp_size": {
				Optional: true,
				Type:     schema.TypeSet,
				//MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"title": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"required": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"enum": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"minimum": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"maximum": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tp_provisioning_type": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"title": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"required": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"min_length": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"max_length": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"enum": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"minimum": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tp_snapshot_pool": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"title": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"required": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"min_length": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"max_length": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"enum": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"default": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"minimum": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tp_data_transfer_limit": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"title": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"required": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"enum": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"minimum": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"maximum": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tp_is_deduplicated": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"title": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"required": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"enum": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tp_is_encrypted": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"title": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"required": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"enum": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tp_is_pinned": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"title": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"required": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"enum": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tp_iops_limit": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"title": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"required": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"enum": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"minimum": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"maximum": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tp_folder": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"title": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"required": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"min_length": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"max_length": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"enum": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"minimum": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tp_template_version": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"title": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"required": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"min_length": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"max_length": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"enum": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"minimum": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tp_performance_policy": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"title": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"required": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"min_length": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"max_length": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"enum": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"minimum": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tp_volume_set": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"title": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"required": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"min_length": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"max_length": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"enum": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"minimum": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tp_description": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"title": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"required": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"min_length": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"max_length": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"enum": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"minimum": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tp_is_adaptive_optimization_enabled": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"title": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"required": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"enum": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"default": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tp_is_compressed": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"title": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"required": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"enum": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tp_data_protection_level": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"title": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"required": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"min_length": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"max_length": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"enum": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"default": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"minimum": {
							Optional: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tp_is_shareable": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"title": {
							Optional: true,
							Type:     schema.TypeString,
						},
						"required": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"enum": {
							Optional: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"default": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"format": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceStorageVolumeTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	template := ov.StorageVolumeTemplate{
		Name:            d.Get("name").(string),
		Description:     utils.NewNstring(d.Get("description").(string)),
		RootTemplateUri: utils.NewNstring(d.Get("root_template_uri").(string))}

	if val, ok := d.GetOk("initial_scope_uris"); ok {
		rawinitialScopeUris := val.(*schema.Set).List()
		initialScopeUris := make([]utils.Nstring, len(rawinitialScopeUris))
		for i, rawData := range rawinitialScopeUris {
			initialScopeUris[i] = utils.Nstring(rawData.(string))
		}
		template.InitialScopeUris = initialScopeUris
	}

	properties := make([]ov.TemplateProperties, 1)
	if val, ok := d.GetOk("tp_name"); ok {
		name_item := ov.TemplatePropertyDatatypeStructString{}
		rawName := val.(*schema.Set).List()
		for _, rawData := range rawName {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				name_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				name_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				name_item.Required = item["required"].(bool)
			}
			if item["max_length"] != nil {
				name_item.Maxlength = item["max_length"].(int)
			}
			if item["min_length"] != nil {
				name_item.Minlength = item["min_length"].(int)
			}
			if item["description"] != nil {
				name_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				name_item.Default = item["default"].(string)
			}
			if item["minimum"] != nil {
				name_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				name_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				name_item.Enum = enums
			}
			name_item.Meta = &meta

		}
		properties[0].Name = &name_item
	}

	//Storage Pool
	if val, ok := d.GetOk("tp_storage_pool"); ok {
		rawStoragePool := val.(*schema.Set).List()
		sp_item := ov.TemplatePropertyDatatypeStructString{}
		for _, rawData := range rawStoragePool {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				sp_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				sp_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				sp_item.Required = item["required"].(bool)
			}
			if item["max_length"] != nil {
				sp_item.Maxlength = item["max_length"].(int)
			}
			if item["min_length"] != nil {
				sp_item.Minlength = item["min_length"].(int)
			}
			if item["description"] != nil {
				sp_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				sp_item.Default = item["default"].(string)
			}
			if item["minimum"] != nil {
				sp_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				sp_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				sp_item.Enum = enums
			}
			sp_item.Meta = &meta
		}
		properties[0].StoragePool = &sp_item
	}

	// Size
	if val, ok := d.GetOk("tp_size"); ok {
		rawSize := val.(*schema.Set).List()
		size_item := ov.TemplatePropertyDatatypeStructInt{}
		for _, rawData := range rawSize {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				size_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				size_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				size_item.Required = item["required"].(bool)
			}
			if item["description"] != nil {
				size_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				size_item.Default = item["default"].(int)
			}
			if item["minimum"] != nil {
				size_item.Minimum = item["minimum"].(int)
			}
			if item["maximum"] != nil {
				size_item.Maximum = item["maximum"].(int)
			}
			if item["format"] != nil {
				size_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				size_item.Enum = enums
			}
			size_item.Meta = &meta
		}
		properties[0].Size = &size_item
	}

	// Provisioning type
	if val, ok := d.GetOk("tp_provisioning_type"); ok {
		rawPT := val.(*schema.Set).List()
		pt_item := ov.TemplatePropertyDatatypeStructString{}
		for _, rawData := range rawPT {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				pt_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				pt_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				pt_item.Required = item["required"].(bool)
			}
			if item["max_length"] != nil {
				pt_item.Maxlength = item["max_length"].(int)
			}
			if item["min_length"] != nil {
				pt_item.Minlength = item["min_length"].(int)
			}
			if item["description"] != nil {
				pt_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				pt_item.Default = item["default"].(string)
			}
			if item["minimum"] != nil {
				pt_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				pt_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				pt_item.Enum = enums
			}
			pt_item.Meta = &meta
		}
		properties[0].ProvisioningType = &pt_item
	}

	//Snapshot Pool
	if val, ok := d.GetOk("tp_snapshot_pool"); ok {
		rawStoragePool := val.(*schema.Set).List()
		sp_item := ov.TemplatePropertyDatatypeStructString{}
		for _, rawData := range rawStoragePool {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				sp_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				sp_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				sp_item.Required = item["required"].(bool)
			}
			if item["max_length"] != nil {
				sp_item.Maxlength = item["max_length"].(int)
			}
			if item["min_length"] != nil {
				sp_item.Minlength = item["min_length"].(int)
			}
			if item["description"] != nil {
				sp_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				sp_item.Default = item["default"].(string)
			}
			if item["minimum"] != nil {
				sp_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				sp_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				sp_item.Enum = enums
			}
			sp_item.Meta = &meta
		}
		properties[0].SnapshotPool = &sp_item
	}

	//DataTransferLimit
	if val, ok := d.GetOk("tp_data_transfer_limit"); ok {
		rawDTL := val.(*schema.Set).List()
		stl_item := ov.TemplatePropertyDatatypeStructInt{}
		for _, rawData := range rawDTL {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				stl_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				stl_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				stl_item.Required = item["required"].(bool)
			}
			if item["maximum"] != nil {
				stl_item.Maximum = item["maximum"].(int)
			}
			if item["description"] != nil {
				stl_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				stl_item.Default = item["default"].(int)
			}
			if item["minimum"] != nil {
				stl_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				stl_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				stl_item.Enum = enums
			}
			stl_item.Meta = &meta
		}
		properties[0].DataTransferLimit = &stl_item
	}

	//IsDeduplicated
	if val, ok := d.GetOk("tp_is_deduplicated"); ok {
		rawIDD := val.(*schema.Set).List()
		idd_item := ov.TemplatePropertyDatatypeStructBool{}
		for _, rawData := range rawIDD {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				idd_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				idd_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				idd_item.Required = item["required"].(bool)
			}
			if item["description"] != nil {
				idd_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				idd_item.Default = item["default"].(bool)
			}
			if item["format"] != nil {
				idd_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				idd_item.Enum = enums
			}
			idd_item.Meta = &meta
		}
		properties[0].IsDeduplicated = &idd_item
	}

	//IsEncrypted
	if val, ok := d.GetOk("tp_is_encrypted"); ok {
		rawIE := val.(*schema.Set).List()
		ie_item := ov.TemplatePropertyDatatypeStructBool{}
		for _, rawData := range rawIE {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				ie_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				ie_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				ie_item.Required = item["required"].(bool)
			}
			if item["description"] != nil {
				ie_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				ie_item.Default = item["default"].(bool)
			}
			if item["format"] != nil {
				ie_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				ie_item.Enum = enums
			}
			ie_item.Meta = &meta
		}
		properties[0].IsEncrypted = &ie_item
	}

	//IsPinned

	if val, ok := d.GetOk("tp_is_pinned"); ok {
		rawIsPinned := val.(*schema.Set).List()
		ip_item := ov.TemplatePropertyDatatypeStructBool{}
		for _, rawData := range rawIsPinned {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				ip_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				ip_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				ip_item.Required = item["required"].(bool)
			}
			if item["description"] != nil {
				ip_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				ip_item.Default = item["default"].(bool)
			}
			if item["format"] != nil {
				ip_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				ip_item.Enum = enums
			}
			ip_item.Meta = &meta
		}
		properties[0].IsPinned = &ip_item
	}

	//IopsLimit
	if val, ok := d.GetOk("tp_iops_limit"); ok {
		rawIopsLimit := val.(*schema.Set).List()
		il_item := ov.TemplatePropertyDatatypeStructInt{}
		for _, rawData := range rawIopsLimit {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				il_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				il_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				il_item.Required = item["required"].(bool)
			}
			if item["description"] != nil {
				il_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				il_item.Default = item["default"].(int)
			}
			if item["minimum"] != nil {
				il_item.Minimum = item["minimum"].(int)
			}
			if item["maximum"] != nil {
				il_item.Maximum = item["maximum"].(int)
			}
			if item["format"] != nil {
				il_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				il_item.Enum = enums
			}
			il_item.Meta = &meta
		}
		properties[0].IopsLimit = &il_item
	}

	//Folder
	if val, ok := d.GetOk("tp_folder"); ok {
		rawFolder := val.(*schema.Set).List()
		folder_item := ov.TemplatePropertyDatatypeStructString{}
		for _, rawData := range rawFolder {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				folder_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				folder_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				folder_item.Required = item["required"].(bool)
			}
			if item["max_length"] != nil {
				folder_item.Maxlength = item["max_length"].(int)
			}
			if item["min_length"] != nil {
				folder_item.Minlength = item["min_length"].(int)
			}
			if item["description"] != nil {
				folder_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				folder_item.Default = item["default"].(string)
			}
			if item["minimum"] != nil {
				folder_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				folder_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				folder_item.Enum = enums
			}
			folder_item.Meta = &meta
		}
		properties[0].Folder = &folder_item
	}

	//TemplateVersion
	if val, ok := d.GetOk("tp_template_version"); ok {
		rawTV := val.(*schema.Set).List()
		templateVersion_item := ov.TemplatePropertyDatatypeStructString{}
		for _, rawData := range rawTV {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				templateVersion_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				templateVersion_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				templateVersion_item.Required = item["required"].(bool)
			}
			if item["max_length"] != nil {
				templateVersion_item.Maxlength = item["max_length"].(int)
			}
			if item["min_length"] != nil {
				templateVersion_item.Minlength = item["min_length"].(int)
			}
			if item["description"] != nil {
				templateVersion_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				templateVersion_item.Default = item["default"].(string)
			}
			if item["minimum"] != nil {
				templateVersion_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				templateVersion_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				templateVersion_item.Enum = enums
			}
			templateVersion_item.Meta = &meta
		}
		properties[0].TemplateVersion = &templateVersion_item
	}

	//PerformancePolicy
	if val, ok := d.GetOk("tp_performance_policy"); ok {
		rawPP := val.(*schema.Set).List()
		pp_item := ov.TemplatePropertyDatatypeStructString{}
		for _, rawData := range rawPP {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				pp_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				pp_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				pp_item.Required = item["required"].(bool)
			}
			if item["max_length"] != nil {
				pp_item.Maxlength = item["max_length"].(int)
			}
			if item["min_length"] != nil {
				pp_item.Minlength = item["min_length"].(int)
			}
			if item["description"] != nil {
				pp_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				pp_item.Default = item["default"].(string)
			}
			if item["minimum"] != nil {
				pp_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				pp_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				pp_item.Enum = enums
			}
			pp_item.Meta = &meta
		}
		properties[0].PerformancePolicy = &pp_item
	}

	//VolumetSet
	if val, ok := d.GetOk("tp_volume_set"); ok {
		rawVS := val.(*schema.Set).List()
		vs_item := ov.TemplatePropertyDatatypeStructString{}
		for _, rawData := range rawVS {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				vs_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				vs_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				vs_item.Required = item["required"].(bool)
			}
			if item["max_length"] != nil {
				vs_item.Maxlength = item["max_length"].(int)
			}
			if item["min_length"] != nil {
				vs_item.Minlength = item["min_length"].(int)
			}
			if item["description"] != nil {
				vs_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				vs_item.Default = item["default"].(string)
			}
			if item["minimum"] != nil {
				vs_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				vs_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				vs_item.Enum = enums
			}
			vs_item.Meta = &meta
		}
		properties[0].VolumetSet = &vs_item
	}

	//Description
	if val, ok := d.GetOk("tp_description"); ok {
		rawDesc := val.(*schema.Set).List()
		desc_item := ov.TemplatePropertyDatatypeStructString{}
		for _, rawData := range rawDesc {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				desc_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				desc_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				desc_item.Required = item["required"].(bool)
			}
			if item["max_length"] != nil {
				desc_item.Maxlength = item["max_length"].(int)
			}
			if item["min_length"] != nil {
				desc_item.Minlength = item["min_length"].(int)
			}
			if item["description"] != nil {
				desc_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				desc_item.Default = item["default"].(string)
			}
			if item["minimum"] != nil {
				desc_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				desc_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				desc_item.Enum = enums
			}
			desc_item.Meta = &meta
		}
		properties[0].Description = &desc_item
	}

	//IsAdaptiveOptimizationEnabled
	if val, ok := d.GetOk("tp_is_adaptive_optimization_enabled"); ok {
		rawIAOE := val.(*schema.Set).List()
		iaoe_item := ov.TemplatePropertyDatatypeStructBool{}
		for _, rawData := range rawIAOE {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				iaoe_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				iaoe_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				iaoe_item.Required = item["required"].(bool)
			}
			if item["description"] != nil {
				iaoe_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				iaoe_item.Default = item["default"].(bool)
			}
			if item["format"] != nil {
				iaoe_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				iaoe_item.Enum = enums
			}
			iaoe_item.Meta = &meta
		}
		properties[0].IsAdaptiveOptimizationEnabled = &iaoe_item
	}

	//IsCompressed
	if val, ok := d.GetOk("tp_is_compressed"); ok {
		rawIC := val.(*schema.Set).List()
		ic_item := ov.TemplatePropertyDatatypeStructBool{}
		for _, rawData := range rawIC {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				ic_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				ic_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				ic_item.Required = item["required"].(bool)
			}
			if item["description"] != nil {
				ic_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				ic_item.Default = item["default"].(bool)
			}
			if item["format"] != nil {
				ic_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				ic_item.Enum = enums
			}
			ic_item.Meta = &meta
		}
		properties[0].IsCompressed = &ic_item
	}

	//DataProtectionLevel
	if val, ok := d.GetOk("tp_data_protection_level"); ok {
		rawDPL := val.(*schema.Set).List()
		dpl_item := ov.TemplatePropertyDatatypeStructString{}
		for _, rawData := range rawDPL {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				dpl_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				dpl_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				dpl_item.Required = item["required"].(bool)
			}
			if item["max_length"] != nil {
				dpl_item.Maxlength = item["max_length"].(int)
			}
			if item["min_length"] != nil {
				dpl_item.Minlength = item["min_length"].(int)
			}
			if item["description"] != nil {
				dpl_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				dpl_item.Default = item["default"].(string)
			}
			if item["minimum"] != nil {
				dpl_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				dpl_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				dpl_item.Enum = enums
			}
			dpl_item.Meta = &meta
		}
		properties[0].DataProtectionLevel = &dpl_item
	}

	//IsShareable
	if val, ok := d.GetOk("tp_is_shareable"); ok {
		rawIS := val.(*schema.Set).List()
		is_item := ov.TemplatePropertyDatatypeStructBool{}
		for _, rawData := range rawIS {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				is_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				is_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				is_item.Required = item["required"].(bool)
			}
			if item["description"] != nil {
				is_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				is_item.Default = item["default"].(bool)
			}
			if item["format"] != nil {
				is_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				is_item.Enum = enums
			}
			is_item.Meta = &meta
		}
		properties[0].IsShareable = &is_item
	}

	template.TemplateProperties = &properties[0]

	err := config.ovClient.CreateStorageVolumeTemplate(template)
	d.SetId(d.Get("name").(string))
	if err != nil {
		d.SetId("")
		return err
	}
	return resourceStorageVolumeTemplateRead(d, meta)
}

func resourceStorageVolumeTemplateRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	template, err := config.ovClient.GetStorageVolumeTemplateByName(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("category", template.Category)
	d.Set("compatible_storage_systems_uri", template.CompatibleStorageSystemsUri.String())
	d.Set("description", template.Description.String())
	d.Set("etag", template.ETAG)
	d.Set("is_root", template.IsRoot)
	d.Set("scopes_uri", template.ScopesURI.String())
	d.Set("name", template.Name)
	d.Set("state", template.State)
	d.Set("status", template.Status)
	d.Set("type", template.Type)
	d.Set("uri", template.URI.String())
	d.Set("initial_scope_uris", template.InitialScopeUris)
	d.Set("root_template_uri", template.RootTemplateUri.String())
	d.Set("storage_pool_uri", template.StoragePoolUri.String())
	d.Set("version", template.Version)
	d.Set("uuid", template.Uuid)
	d.Set("family", template.Family)

	if template.TemplateProperties != nil {
		item := template.TemplateProperties.Name
		if item != nil {
			name := make([]map[string]interface{}, 0)
			enumMap := make([]interface{}, len(item.Enum))
			for i, enum := range item.Enum {
				enumMap[i] = enum
			}

			name = append(name, map[string]interface{}{
				"meta_create_only":   item.Meta.CreateOnly,
				"meta_locked":        item.Meta.Locked,
				"meta_semantic_type": item.Meta.SemanticType,
				"type":               item.Type,
				"title":              item.Title,
				"required":           item.Required,
				"max_length":         item.Maxlength,
				"min_length":         item.Minlength,
				"description":        item.Description,
				"enum":               schema.NewSet(schema.HashString, enumMap),
				"default":            item.Default,
				"minimum":            item.Minimum,
				"format":             item.Format})

			d.Set("tp_name", name)
		}

		//Storage Pool
		sp_item := template.TemplateProperties.StoragePool
		if sp_item != nil {
			storagePool := make([]map[string]interface{}, 0)
			enumMap := make([]interface{}, len(sp_item.Enum))
			for i, enum := range sp_item.Enum {
				enumMap[i] = enum
			}

			storagePool = append(storagePool, map[string]interface{}{
				"meta_create_only":   sp_item.Meta.CreateOnly,
				"meta_locked":        sp_item.Meta.Locked,
				"meta_semantic_type": sp_item.Meta.SemanticType,
				"type":               sp_item.Type,
				"title":              sp_item.Title,
				"required":           sp_item.Required,
				"max_length":         sp_item.Maxlength,
				"min_length":         sp_item.Minlength,
				"description":        sp_item.Description,
				"enum":               schema.NewSet(schema.HashString, enumMap),
				"default":            sp_item.Default,
				"minimum":            sp_item.Minimum,
				"format":             sp_item.Format})

			d.Set("tp_storage_pool", storagePool)
		}

		//Size
		size_item := template.TemplateProperties.Size
		if size_item != nil {
			size := make([]map[string]interface{}, 0)
			enumMap := make([]interface{}, len(size_item.Enum))
			for i, enum := range size_item.Enum {
				enumMap[i] = enum
			}

			size = append(size, map[string]interface{}{
				"meta_create_only":   size_item.Meta.CreateOnly,
				"meta_locked":        size_item.Meta.Locked,
				"meta_semantic_type": size_item.Meta.SemanticType,
				"type":               size_item.Type,
				"title":              size_item.Title,
				"required":           size_item.Required,
				"description":        size_item.Description,
				"enum":               schema.NewSet(schema.HashString, enumMap),
				"default":            size_item.Default,
				"minimum":            size_item.Minimum,
				"maximum":            size_item.Maximum,
				"format":             size_item.Format})

			d.Set("tp_size", size)
		}

		//ProvisioningType
		pt_item := template.TemplateProperties.ProvisioningType
		if pt_item != nil {
			provisioningType := make([]map[string]interface{}, 0)
			enumMap := make([]interface{}, len(pt_item.Enum))
			for i, enum := range pt_item.Enum {
				enumMap[i] = enum
			}

			provisioningType = append(provisioningType, map[string]interface{}{
				"meta_create_only":   pt_item.Meta.CreateOnly,
				"meta_locked":        pt_item.Meta.Locked,
				"meta_semantic_type": pt_item.Meta.SemanticType,
				"type":               pt_item.Type,
				"title":              pt_item.Title,
				"required":           pt_item.Required,
				"max_length":         pt_item.Maxlength,
				"min_length":         pt_item.Minlength,
				"description":        pt_item.Description,
				"enum":               schema.NewSet(schema.HashString, enumMap),
				"default":            pt_item.Default,
				"minimum":            pt_item.Minimum,
				"format":             pt_item.Format})

			d.Set("tp_provisioning_type", provisioningType)
		}

		//SnapShotPool
		ssp_item := template.TemplateProperties.SnapshotPool
		if ssp_item != nil {
			snapshotPool := make([]map[string]interface{}, 0)
			enumMap := make([]interface{}, len(ssp_item.Enum))
			for i, enum := range ssp_item.Enum {
				enumMap[i] = enum
			}

			snapshotPool = append(snapshotPool, map[string]interface{}{
				"meta_create_only":   ssp_item.Meta.CreateOnly,
				"meta_locked":        ssp_item.Meta.Locked,
				"meta_semantic_type": ssp_item.Meta.SemanticType,
				"type":               ssp_item.Type,
				"title":              ssp_item.Title,
				"required":           ssp_item.Required,
				"max_length":         ssp_item.Maxlength,
				"min_length":         ssp_item.Minlength,
				"description":        ssp_item.Description,
				"enum":               schema.NewSet(schema.HashString, enumMap),
				"default":            ssp_item.Default,
				"minimum":            ssp_item.Minimum,
				"format":             ssp_item.Format})

			d.Set("tp_snapshot_pool", snapshotPool)
		}

		//DataTransferLimit
		dtl_item := template.TemplateProperties.DataTransferLimit
		if dtl_item != nil {
			dtl := make([]map[string]interface{}, 0)
			enumMap := make([]interface{}, len(dtl_item.Enum))
			for i, enum := range dtl_item.Enum {
				enumMap[i] = enum
			}

			dtl = append(dtl, map[string]interface{}{
				"meta_create_only":   dtl_item.Meta.CreateOnly,
				"meta_locked":        dtl_item.Meta.Locked,
				"meta_semantic_type": dtl_item.Meta.SemanticType,
				"type":               dtl_item.Type,
				"title":              dtl_item.Title,
				"required":           dtl_item.Required,
				"description":        dtl_item.Description,
				"enum":               schema.NewSet(schema.HashString, enumMap),
				"default":            dtl_item.Default,
				"minimum":            dtl_item.Minimum,
				"maximum":            dtl_item.Maximum,
				"format":             dtl_item.Format})

			d.Set("tp_data_transfer_limit", dtl)
		}

		//IsDeduplicated
		idd_item := template.TemplateProperties.IsDeduplicated
		if idd_item != nil {
			isDeDuplicated := make([]map[string]interface{}, 0)
			enumMap := make([]interface{}, len(idd_item.Enum))
			for i, enum := range idd_item.Enum {
				enumMap[i] = enum
			}

			isDeDuplicated = append(isDeDuplicated, map[string]interface{}{
				"meta_create_only":   idd_item.Meta.CreateOnly,
				"meta_locked":        idd_item.Meta.Locked,
				"meta_semantic_type": idd_item.Meta.SemanticType,
				"type":               idd_item.Type,
				"title":              idd_item.Title,
				"required":           idd_item.Required,
				"description":        idd_item.Description,
				"enum":               schema.NewSet(schema.HashString, enumMap),
				"default":            idd_item.Default,
				"format":             idd_item.Format})

			d.Set("tp_is_deduplicated", isDeDuplicated)
		}

		//IsEncrypted
		ie_item := template.TemplateProperties.IsEncrypted
		if ie_item != nil {
			isEncrypted := make([]map[string]interface{}, 0)
			enumMap := make([]interface{}, len(ie_item.Enum))
			for i, enum := range ie_item.Enum {
				enumMap[i] = enum
			}

			isEncrypted = append(isEncrypted, map[string]interface{}{
				"meta_create_only":   ie_item.Meta.CreateOnly,
				"meta_locked":        ie_item.Meta.Locked,
				"meta_semantic_type": ie_item.Meta.SemanticType,
				"type":               ie_item.Type,
				"title":              ie_item.Title,
				"required":           ie_item.Required,
				"description":        ie_item.Description,
				"enum":               schema.NewSet(schema.HashString, enumMap),
				"default":            ie_item.Default,
				"format":             ie_item.Format})

			d.Set("tp_is_encrypted", isEncrypted)
		}

		//IsPinned
		ip_item := template.TemplateProperties.IsPinned
		if ip_item != nil {
			isPinned := make([]map[string]interface{}, 0)
			enumMap := make([]interface{}, len(ip_item.Enum))
			for i, enum := range ip_item.Enum {
				enumMap[i] = enum
			}

			isPinned = append(isPinned, map[string]interface{}{
				"meta_create_only":   ip_item.Meta.CreateOnly,
				"meta_locked":        ip_item.Meta.Locked,
				"meta_semantic_type": ip_item.Meta.SemanticType,
				"type":               ip_item.Type,
				"title":              ip_item.Title,
				"required":           ip_item.Required,
				"description":        ip_item.Description,
				"enum":               schema.NewSet(schema.HashString, enumMap),
				"default":            ip_item.Default,
				"format":             ip_item.Format})

			d.Set("tp_is_pinned", isPinned)
		}

		//Folder
		folder_item := template.TemplateProperties.Folder
		if folder_item != nil {
			folder := make([]map[string]interface{}, 0)
			enumMap := make([]interface{}, len(folder_item.Enum))
			for i, enum := range folder_item.Enum {
				enumMap[i] = enum
			}

			folder = append(folder, map[string]interface{}{
				"meta_create_only":   folder_item.Meta.CreateOnly,
				"meta_locked":        folder_item.Meta.Locked,
				"meta_semantic_type": folder_item.Meta.SemanticType,
				"type":               folder_item.Type,
				"title":              folder_item.Title,
				"required":           folder_item.Required,
				"max_length":         folder_item.Maxlength,
				"min_length":         folder_item.Minlength,
				"description":        folder_item.Description,
				"enum":               schema.NewSet(schema.HashString, enumMap),
				"default":            folder_item.Default,
				"minimum":            folder_item.Minimum,
				"format":             folder_item.Format})

			d.Set("tp_folder", folder)
		}

		//IopsLimit
		il_item := template.TemplateProperties.IopsLimit
		if il_item != nil {
			iopsLimit := make([]map[string]interface{}, 0)
			enumMap := make([]interface{}, len(il_item.Enum))
			for i, enum := range il_item.Enum {
				enumMap[i] = enum
			}

			iopsLimit = append(iopsLimit, map[string]interface{}{
				"meta_create_only":   il_item.Meta.CreateOnly,
				"meta_locked":        il_item.Meta.Locked,
				"meta_semantic_type": il_item.Meta.SemanticType,
				"type":               il_item.Type,
				"title":              il_item.Title,
				"required":           il_item.Required,
				"description":        il_item.Description,
				"enum":               schema.NewSet(schema.HashString, enumMap),
				"default":            il_item.Default,
				"minimum":            il_item.Minimum,
				"maximum":            il_item.Maximum,
				"format":             il_item.Format})

			d.Set("tp_iops_limit", iopsLimit)
		}

		//TemplateVersion
		tv_item := template.TemplateProperties.TemplateVersion
		if tv_item != nil {
			templateVersion := make([]map[string]interface{}, 0)
			enumMap := make([]interface{}, len(tv_item.Enum))
			for i, enum := range tv_item.Enum {
				enumMap[i] = enum
			}

			templateVersion = append(templateVersion, map[string]interface{}{
				"meta_create_only":   tv_item.Meta.CreateOnly,
				"meta_locked":        tv_item.Meta.Locked,
				"meta_semantic_type": tv_item.Meta.SemanticType,
				"type":               tv_item.Type,
				"title":              tv_item.Title,
				"required":           tv_item.Required,
				"max_length":         tv_item.Maxlength,
				"min_length":         tv_item.Minlength,
				"description":        tv_item.Description,
				"enum":               schema.NewSet(schema.HashString, enumMap),
				"default":            tv_item.Default,
				"minimum":            tv_item.Minimum,
				"format":             tv_item.Format})

			d.Set("tp_template_version", templateVersion)
		}

		//PerformancePolicy
		pp_item := template.TemplateProperties.PerformancePolicy
		if pp_item != nil {
			pp := make([]map[string]interface{}, 0)
			enumMap := make([]interface{}, len(pp_item.Enum))
			for i, enum := range pp_item.Enum {
				enumMap[i] = enum
			}

			pp = append(pp, map[string]interface{}{
				"meta_create_only":   pp_item.Meta.CreateOnly,
				"meta_locked":        pp_item.Meta.Locked,
				"meta_semantic_type": pp_item.Meta.SemanticType,
				"type":               pp_item.Type,
				"title":              pp_item.Title,
				"required":           pp_item.Required,
				"max_length":         pp_item.Maxlength,
				"min_length":         pp_item.Minlength,
				"description":        pp_item.Description,
				"enum":               schema.NewSet(schema.HashString, enumMap),
				"default":            pp_item.Default,
				"minimum":            pp_item.Minimum,
				"format":             pp_item.Format})

			d.Set("tp_performance_policy", pp)
		}

		//VolumetSet
		vs_item := template.TemplateProperties.VolumetSet
		if vs_item != nil {
			volumeSet := make([]map[string]interface{}, 0)
			enumMap := make([]interface{}, len(vs_item.Enum))
			for i, enum := range vs_item.Enum {
				enumMap[i] = enum
			}

			volumeSet = append(volumeSet, map[string]interface{}{
				"meta_create_only":   vs_item.Meta.CreateOnly,
				"meta_locked":        vs_item.Meta.Locked,
				"meta_semantic_type": vs_item.Meta.SemanticType,
				"type":               vs_item.Type,
				"title":              vs_item.Title,
				"required":           vs_item.Required,
				"max_length":         vs_item.Maxlength,
				"min_length":         vs_item.Minlength,
				"description":        vs_item.Description,
				"enum":               schema.NewSet(schema.HashString, enumMap),
				"default":            vs_item.Default,
				"minimum":            vs_item.Minimum,
				"format":             vs_item.Format})

			d.Set("tp_volume_set", volumeSet)
		}

		//Description
		desc_item := template.TemplateProperties.Description
		if desc_item != nil {
			desc := make([]map[string]interface{}, 0)
			enumMap := make([]interface{}, len(desc_item.Enum))
			for i, enum := range desc_item.Enum {
				enumMap[i] = enum
			}

			desc = append(desc, map[string]interface{}{
				"meta_create_only":   desc_item.Meta.CreateOnly,
				"meta_locked":        desc_item.Meta.Locked,
				"meta_semantic_type": desc_item.Meta.SemanticType,
				"type":               desc_item.Type,
				"title":              desc_item.Title,
				"required":           desc_item.Required,
				"max_length":         desc_item.Maxlength,
				"min_length":         desc_item.Minlength,
				"description":        desc_item.Description,
				"enum":               schema.NewSet(schema.HashString, enumMap),
				"default":            desc_item.Default,
				"minimum":            desc_item.Minimum,
				"format":             desc_item.Format})

			d.Set("tp_description", desc)
		}

		//IsAdaptiveOptimizationEnabled
		iaoe_item := template.TemplateProperties.IsAdaptiveOptimizationEnabled
		if iaoe_item != nil {
			iaoe := make([]map[string]interface{}, 0)
			enumMap := make([]interface{}, len(iaoe_item.Enum))
			for i, enum := range iaoe_item.Enum {
				enumMap[i] = enum
			}

			iaoe = append(iaoe, map[string]interface{}{
				"meta_create_only":   iaoe_item.Meta.CreateOnly,
				"meta_locked":        iaoe_item.Meta.Locked,
				"meta_semantic_type": iaoe_item.Meta.SemanticType,
				"type":               iaoe_item.Type,
				"title":              iaoe_item.Title,
				"required":           iaoe_item.Required,
				"description":        iaoe_item.Description,
				"enum":               schema.NewSet(schema.HashString, enumMap),
				"default":            iaoe_item.Default,
				"format":             iaoe_item.Format})

			d.Set("tp_is_adaptive_optimization_enabled", iaoe)
		}

		//IsCompressed
		ic_item := template.TemplateProperties.IsCompressed
		if ic_item != nil {
			isCompressed := make([]map[string]interface{}, 0)
			enumMap := make([]interface{}, len(ic_item.Enum))
			for i, enum := range ic_item.Enum {
				enumMap[i] = enum
			}

			isCompressed = append(isCompressed, map[string]interface{}{
				"meta_create_only":   ic_item.Meta.CreateOnly,
				"meta_locked":        ic_item.Meta.Locked,
				"meta_semantic_type": ic_item.Meta.SemanticType,
				"type":               ic_item.Type,
				"title":              ic_item.Title,
				"required":           ic_item.Required,
				"description":        ic_item.Description,
				"enum":               schema.NewSet(schema.HashString, enumMap),
				"default":            ic_item.Default,
				"format":             ic_item.Format})

			d.Set("tp_is_compressed", isCompressed)
		}

		//DataProtectionLevel
		dpl_item := template.TemplateProperties.DataProtectionLevel
		if dpl_item != nil {
			dpl := make([]map[string]interface{}, 0)
			enumMap := make([]interface{}, len(dpl_item.Enum))
			for i, enum := range dpl_item.Enum {
				enumMap[i] = enum
			}

			dpl = append(dpl, map[string]interface{}{
				"meta_create_only":   dpl_item.Meta.CreateOnly,
				"meta_locked":        dpl_item.Meta.Locked,
				"meta_semantic_type": dpl_item.Meta.SemanticType,
				"type":               dpl_item.Type,
				"title":              dpl_item.Title,
				"required":           dpl_item.Required,
				"max_length":         dpl_item.Maxlength,
				"min_length":         dpl_item.Minlength,
				"dplription":         dpl_item.Description,
				"enum":               schema.NewSet(schema.HashString, enumMap),
				"default":            dpl_item.Default,
				"minimum":            dpl_item.Minimum,
				"format":             dpl_item.Format})

			d.Set("tp_data_protection_level", dpl)
		}

		//IsShareable
		is_item := template.TemplateProperties.IsShareable
		if is_item != nil {
			isShareable := make([]map[string]interface{}, 0)
			enumMap := make([]interface{}, len(is_item.Enum))
			for i, enum := range is_item.Enum {
				enumMap[i] = enum
			}

			isShareable = append(isShareable, map[string]interface{}{
				"meta_create_only":   is_item.Meta.CreateOnly,
				"meta_locked":        is_item.Meta.Locked,
				"meta_semantic_type": is_item.Meta.SemanticType,
				"type":               is_item.Type,
				"title":              is_item.Title,
				"required":           is_item.Required,
				"description":        is_item.Description,
				"enum":               schema.NewSet(schema.HashString, enumMap),
				"default":            is_item.Default,
				"format":             is_item.Format})

			d.Set("tp_is_shareable", isShareable)
		}
	}

	return nil
}

func resourceStorageVolumeTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	template := ov.StorageVolumeTemplate{
		Name: d.Get("name").(string),
		ETAG: d.Get("etag").(string),
		URI:  utils.NewNstring(d.Get("uri").(string))}

	if val, ok := d.GetOk("description"); ok {
		template.Description = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk("scopes_uri"); ok {
		template.ScopesURI = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk("family"); ok {
		template.Family = val.(string)
	}

	if val, ok := d.GetOk("is_root"); ok {
		template.IsRoot = val.(bool)
	}

	if val, ok := d.GetOk("storage_pool_uri"); ok {
		template.StoragePoolUri = utils.NewNstring(val.(string))
	}

	if val, ok := d.GetOk("root_template_uri"); ok {
		template.RootTemplateUri = utils.NewNstring(val.(string))
	}

	properties := make([]ov.TemplateProperties, 1)
	if val, ok := d.GetOk("tp_name"); ok {
		name_item := ov.TemplatePropertyDatatypeStructString{}
		rawName := val.(*schema.Set).List()
		for _, rawData := range rawName {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				name_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				name_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				name_item.Required = item["required"].(bool)
			}
			if item["max_length"] != nil {
				name_item.Maxlength = item["max_length"].(int)
			}
			if item["min_length"] != nil {
				name_item.Minlength = item["min_length"].(int)
			}
			if item["description"] != nil {
				name_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				name_item.Default = item["default"].(string)
			}
			if item["minimum"] != nil {
				name_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				name_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				name_item.Enum = enums
			}
			name_item.Meta = &meta
		}
		properties[0].Name = &name_item
	}

	//Storage Pool
	if val, ok := d.GetOk("tp_storage_pool"); ok {
		rawStoragePool := val.(*schema.Set).List()
		sp_item := ov.TemplatePropertyDatatypeStructString{}
		for _, rawData := range rawStoragePool {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				sp_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				sp_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				sp_item.Required = item["required"].(bool)
			}
			if item["max_length"] != nil {
				sp_item.Maxlength = item["max_length"].(int)
			}
			if item["min_length"] != nil {
				sp_item.Minlength = item["min_length"].(int)
			}
			if item["description"] != nil {
				sp_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				sp_item.Default = item["default"].(string)
			}
			if item["minimum"] != nil {
				sp_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				sp_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				sp_item.Enum = enums
			}
			sp_item.Meta = &meta
		}
		properties[0].StoragePool = &sp_item
	}

	// Size
	if val, ok := d.GetOk("tp_size"); ok {
		rawSize := val.(*schema.Set).List()
		size_item := ov.TemplatePropertyDatatypeStructInt{}
		for _, rawData := range rawSize {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				size_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				size_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				size_item.Required = item["required"].(bool)
			}
			if item["description"] != nil {
				size_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				size_item.Default = item["default"].(int)
			}
			if item["minimum"] != nil {
				size_item.Minimum = item["minimum"].(int)
			}
			if item["maximum"] != nil {
				size_item.Maximum = item["maximum"].(int)
			}
			if item["format"] != nil {
				size_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				size_item.Enum = enums
			}
			size_item.Meta = &meta
		}
		properties[0].Size = &size_item
	}

	// Provisioning type
	if val, ok := d.GetOk("tp_provisioning_type"); ok {
		rawPT := val.(*schema.Set).List()
		pt_item := ov.TemplatePropertyDatatypeStructString{}
		for _, rawData := range rawPT {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				pt_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				pt_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				pt_item.Required = item["required"].(bool)
			}
			if item["max_length"] != nil {
				pt_item.Maxlength = item["max_length"].(int)
			}
			if item["min_length"] != nil {
				pt_item.Minlength = item["min_length"].(int)
			}
			if item["description"] != nil {
				pt_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				pt_item.Default = item["default"].(string)
			}
			if item["minimum"] != nil {
				pt_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				pt_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				pt_item.Enum = enums
			}
			pt_item.Meta = &meta
		}
		properties[0].ProvisioningType = &pt_item
	}

	//Snapshot Pool
	if val, ok := d.GetOk("tp_snapshot_pool"); ok {
		rawStoragePool := val.(*schema.Set).List()
		sp_item := ov.TemplatePropertyDatatypeStructString{}
		for _, rawData := range rawStoragePool {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				sp_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				sp_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				sp_item.Required = item["required"].(bool)
			}
			if item["max_length"] != nil {
				sp_item.Maxlength = item["max_length"].(int)
			}
			if item["min_length"] != nil {
				sp_item.Minlength = item["min_length"].(int)
			}
			if item["description"] != nil {
				sp_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				sp_item.Default = item["default"].(string)
			}
			if item["minimum"] != nil {
				sp_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				sp_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				sp_item.Enum = enums
			}
			sp_item.Meta = &meta
		}
		properties[0].SnapshotPool = &sp_item
	}

	//DataTransferLimit
	if val, ok := d.GetOk("tp_data_transfer_limit"); ok {
		rawDTL := val.(*schema.Set).List()
		stl_item := ov.TemplatePropertyDatatypeStructInt{}
		for _, rawData := range rawDTL {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				stl_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				stl_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				stl_item.Required = item["required"].(bool)
			}
			if item["maximum"] != nil {
				stl_item.Maximum = item["maximum"].(int)
			}
			if item["description"] != nil {
				stl_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				stl_item.Default = item["default"].(int)
			}
			if item["minimum"] != nil {
				stl_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				stl_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				stl_item.Enum = enums
			}
			stl_item.Meta = &meta
		}
		properties[0].DataTransferLimit = &stl_item
	}

	//IsDeduplicated
	if val, ok := d.GetOk("tp_is_deduplicated"); ok {
		rawIDD := val.(*schema.Set).List()
		idd_item := ov.TemplatePropertyDatatypeStructBool{}
		for _, rawData := range rawIDD {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				idd_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				idd_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				idd_item.Required = item["required"].(bool)
			}
			if item["description"] != nil {
				idd_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				idd_item.Default = item["default"].(bool)
			}
			if item["format"] != nil {
				idd_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				idd_item.Enum = enums
			}
			idd_item.Meta = &meta
		}
		properties[0].IsDeduplicated = &idd_item
	}

	//IsEncrypted
	if val, ok := d.GetOk("tp_is_encrypted"); ok {
		rawIE := val.(*schema.Set).List()
		ie_item := ov.TemplatePropertyDatatypeStructBool{}
		for _, rawData := range rawIE {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				ie_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				ie_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				ie_item.Required = item["required"].(bool)
			}
			if item["description"] != nil {
				ie_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				ie_item.Default = item["default"].(bool)
			}
			if item["format"] != nil {
				ie_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				ie_item.Enum = enums
			}
			ie_item.Meta = &meta
		}
		properties[0].IsEncrypted = &ie_item
	}

	//IsPinned

	if val, ok := d.GetOk("tp_is_pinned"); ok {
		rawIsPinned := val.(*schema.Set).List()
		ip_item := ov.TemplatePropertyDatatypeStructBool{}
		for _, rawData := range rawIsPinned {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				ip_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				ip_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				ip_item.Required = item["required"].(bool)
			}
			if item["description"] != nil {
				ip_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				ip_item.Default = item["default"].(bool)
			}
			if item["format"] != nil {
				ip_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				ip_item.Enum = enums
			}
			ip_item.Meta = &meta
		}
		properties[0].IsPinned = &ip_item
	}

	//IopsLimit
	if val, ok := d.GetOk("tp_iops_limit"); ok {
		rawIopsLimit := val.(*schema.Set).List()
		il_item := ov.TemplatePropertyDatatypeStructInt{}
		for _, rawData := range rawIopsLimit {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				il_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				il_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				il_item.Required = item["required"].(bool)
			}
			if item["description"] != nil {
				il_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				il_item.Default = item["default"].(int)
			}
			if item["minimum"] != nil {
				il_item.Minimum = item["minimum"].(int)
			}
			if item["maximum"] != nil {
				il_item.Maximum = item["maximum"].(int)
			}
			if item["format"] != nil {
				il_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				il_item.Enum = enums
			}
			il_item.Meta = &meta
		}
		properties[0].IopsLimit = &il_item
	}

	//Folder
	if val, ok := d.GetOk("tp_folder"); ok {
		rawFolder := val.(*schema.Set).List()
		folder_item := ov.TemplatePropertyDatatypeStructString{}
		for _, rawData := range rawFolder {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				folder_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				folder_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				folder_item.Required = item["required"].(bool)
			}
			if item["max_length"] != nil {
				folder_item.Maxlength = item["max_length"].(int)
			}
			if item["min_length"] != nil {
				folder_item.Minlength = item["min_length"].(int)
			}
			if item["description"] != nil {
				folder_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				folder_item.Default = item["default"].(string)
			}
			if item["minimum"] != nil {
				folder_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				folder_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				folder_item.Enum = enums
			}
			folder_item.Meta = &meta
		}
		properties[0].Folder = &folder_item
	}

	//TemplateVersion
	if val, ok := d.GetOk("tp_template_version"); ok {
		rawTV := val.(*schema.Set).List()
		templateVersion_item := ov.TemplatePropertyDatatypeStructString{}
		for _, rawData := range rawTV {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				templateVersion_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				templateVersion_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				templateVersion_item.Required = item["required"].(bool)
			}
			if item["max_length"] != nil {
				templateVersion_item.Maxlength = item["max_length"].(int)
			}
			if item["min_length"] != nil {
				templateVersion_item.Minlength = item["min_length"].(int)
			}
			if item["description"] != nil {
				templateVersion_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				templateVersion_item.Default = item["default"].(string)
			}
			if item["minimum"] != nil {
				templateVersion_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				templateVersion_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				templateVersion_item.Enum = enums
			}
			templateVersion_item.Meta = &meta
		}
		properties[0].TemplateVersion = &templateVersion_item
	}

	//PerformancePolicy
	if val, ok := d.GetOk("tp_performance_policy"); ok {
		rawPP := val.(*schema.Set).List()
		pp_item := ov.TemplatePropertyDatatypeStructString{}
		for _, rawData := range rawPP {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				pp_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				pp_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				pp_item.Required = item["required"].(bool)
			}
			if item["max_length"] != nil {
				pp_item.Maxlength = item["max_length"].(int)
			}
			if item["min_length"] != nil {
				pp_item.Minlength = item["min_length"].(int)
			}
			if item["description"] != nil {
				pp_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				pp_item.Default = item["default"].(string)
			}
			if item["minimum"] != nil {
				pp_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				pp_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				pp_item.Enum = enums
			}
			pp_item.Meta = &meta
		}
		properties[0].PerformancePolicy = &pp_item
	}

	//VolumetSet
	if val, ok := d.GetOk("tp_volume_set"); ok {
		rawVS := val.(*schema.Set).List()
		vs_item := ov.TemplatePropertyDatatypeStructString{}
		for _, rawData := range rawVS {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				vs_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				vs_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				vs_item.Required = item["required"].(bool)
			}
			if item["max_length"] != nil {
				vs_item.Maxlength = item["max_length"].(int)
			}
			if item["min_length"] != nil {
				vs_item.Minlength = item["min_length"].(int)
			}
			if item["description"] != nil {
				vs_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				vs_item.Default = item["default"].(string)
			}
			if item["minimum"] != nil {
				vs_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				vs_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				vs_item.Enum = enums
			}
			vs_item.Meta = &meta
		}
		properties[0].VolumetSet = &vs_item
	}

	//Description
	if val, ok := d.GetOk("tp_description"); ok {
		rawDesc := val.(*schema.Set).List()
		desc_item := ov.TemplatePropertyDatatypeStructString{}
		for _, rawData := range rawDesc {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				desc_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				desc_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				desc_item.Required = item["required"].(bool)
			}
			if item["max_length"] != nil {
				desc_item.Maxlength = item["max_length"].(int)
			}
			if item["min_length"] != nil {
				desc_item.Minlength = item["min_length"].(int)
			}
			if item["description"] != nil {
				desc_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				desc_item.Default = item["default"].(string)
			}
			if item["minimum"] != nil {
				desc_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				desc_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				desc_item.Enum = enums
			}
			desc_item.Meta = &meta
		}
		properties[0].Description = &desc_item
	}

	//IsAdaptiveOptimizationEnabled
	if val, ok := d.GetOk("tp_is_adaptive_optimization_enabled"); ok {
		rawIAOE := val.(*schema.Set).List()
		iaoe_item := ov.TemplatePropertyDatatypeStructBool{}
		for _, rawData := range rawIAOE {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				iaoe_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				iaoe_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				iaoe_item.Required = item["required"].(bool)
			}
			if item["description"] != nil {
				iaoe_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				iaoe_item.Default = item["default"].(bool)
			}
			if item["format"] != nil {
				iaoe_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				iaoe_item.Enum = enums
			}
			iaoe_item.Meta = &meta
		}
		properties[0].IsAdaptiveOptimizationEnabled = &iaoe_item
	}

	//IsCompressed
	if val, ok := d.GetOk("tp_is_compressed"); ok {
		rawIC := val.(*schema.Set).List()
		ic_item := ov.TemplatePropertyDatatypeStructBool{}
		for _, rawData := range rawIC {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				ic_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				ic_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				ic_item.Required = item["required"].(bool)
			}
			if item["description"] != nil {
				ic_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				ic_item.Default = item["default"].(bool)
			}
			if item["format"] != nil {
				ic_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				ic_item.Enum = enums
			}
			ic_item.Meta = &meta
		}
		properties[0].IsCompressed = &ic_item
	}

	//DataProtectionLevel
	if val, ok := d.GetOk("tp_data_protection_level"); ok {
		rawDPL := val.(*schema.Set).List()
		dpl_item := ov.TemplatePropertyDatatypeStructString{}
		for _, rawData := range rawDPL {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				dpl_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				dpl_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				dpl_item.Required = item["required"].(bool)
			}
			if item["max_length"] != nil {
				dpl_item.Maxlength = item["max_length"].(int)
			}
			if item["min_length"] != nil {
				dpl_item.Minlength = item["min_length"].(int)
			}
			if item["description"] != nil {
				dpl_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				dpl_item.Default = item["default"].(string)
			}
			if item["minimum"] != nil {
				dpl_item.Minimum = item["minimum"].(int)
			}
			if item["format"] != nil {
				dpl_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				dpl_item.Enum = enums
			}
			dpl_item.Meta = &meta
		}
		properties[0].DataProtectionLevel = &dpl_item
	}

	//IsShareable
	if val, ok := d.GetOk("tp_is_shareable"); ok {
		rawIS := val.(*schema.Set).List()
		is_item := ov.TemplatePropertyDatatypeStructBool{}
		for _, rawData := range rawIS {
			item := rawData.(map[string]interface{})
			meta := ov.Meta{}

			if item["type"] != nil {
				is_item.Type = item["type"].(string)
			}
			if item["title"] != nil {
				is_item.Title = item["title"].(string)
			}
			if item["required"] != nil {
				is_item.Required = item["required"].(bool)
			}
			if item["description"] != nil {
				is_item.Description = utils.Nstring(item["description"].(string))
			}
			if item["default"] != nil {
				is_item.Default = item["default"].(bool)
			}
			if item["format"] != nil {
				is_item.Format = item["format"].(string)
			}
			if item["meta_locked"] != nil {
				meta.Locked = item["meta_locked"].(bool)
			}
			if item["meta_create_only"] != nil {
				meta.CreateOnly = item["meta_create_only"].(bool)
			}
			if item["meta_semantic_type"] != nil {
				meta.SemanticType = item["meta_semantic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, 0)
				for _, rawData := range rawEnums {
					enums = append(enums, rawData.(string))
				}
				is_item.Enum = enums
			}
			is_item.Meta = &meta
		}
		properties[0].IsShareable = &is_item
	}

	template.TemplateProperties = &properties[0]

	err := config.ovClient.UpdateStorageVolumeTemplate(template)
	d.SetId(d.Get("name").(string))
	if err != nil {
		d.SetId("")
		return err
	}
	return resourceStorageVolumeTemplateRead(d, meta)
}

func resourceStorageVolumeTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteStorageVolumeTemplate(d.Get("name").(string))
	if err != nil {
		return err
	}
	return nil
}
