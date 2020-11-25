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

func dataSourceStorageVolumeTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceStorageVolumeTemplateRead,

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
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_root": {
				Type:     schema.TypeBool,
				Computed: true,
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
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"root_template_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_pool_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"family": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tp_name": {
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"title": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"required": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"min_length": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"max_length": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"enum": {
							Computed: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"minimum": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tp_storage_pool": {
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"title": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"required": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"min_length": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"max_length": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"enum": {
							Computed: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"default": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"minimum": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tp_size": {
				Computed: true,
				Type:     schema.TypeSet,
				//MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"title": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"required": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"enum": {
							Computed: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"minimum": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"maximum": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tp_provisioning_type": {
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"title": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"required": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"min_length": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"max_length": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"enum": {
							Computed: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"minimum": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tp_snapshot_pool": {
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"title": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"required": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"min_length": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"max_length": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"enum": {
							Computed: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"default": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"minimum": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tp_data_transfer_limit": {
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"title": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"required": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"enum": {
							Computed: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"minimum": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"maximum": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tp_is_deduplicated": {
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"title": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"required": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"enum": {
							Computed: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tp_is_encrypted": {
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"title": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"required": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"enum": {
							Computed: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"format": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tp_is_pinned": {
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"title": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"required": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"enum": {
							Computed: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"format": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tp_iops_limit": {
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"title": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"required": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"enum": {
							Computed: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"minimum": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"maximum": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tp_folder": {
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"title": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"required": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"min_length": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"max_length": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"enum": {
							Computed: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"minimum": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tp_template_version": {
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"title": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"required": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"min_length": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"max_length": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"enum": {
							Computed: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"minimum": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tp_performance_policy": {
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"title": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"required": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"min_length": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"max_length": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"enum": {
							Computed: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"minimum": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tp_volume_set": {
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"title": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"required": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"min_length": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"max_length": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"enum": {
							Computed: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"minimum": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tp_description": {
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"title": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"required": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"min_length": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"max_length": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"enum": {
							Computed: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"minimum": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tp_is_adaptive_optimization_enabled": {
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"title": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"required": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"enum": {
							Computed: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"default": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"format": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tp_is_compressed": {
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"title": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"required": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"enum": {
							Computed: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"default": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tp_data_protection_level": {
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"title": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"required": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"min_length": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"max_length": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"enum": {
							Computed: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"default": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"minimum": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"format": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tp_is_shareable": {
				Computed: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"meta_semantic_type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"meta_create_only": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"title": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"required": {
							Computed: true,
							Type:     schema.TypeBool,
						},
						"enum": {
							Computed: true,
							Type:     schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"default": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"format": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceStorageVolumeTemplateRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	id := d.Get("name").(string)

	template, err := config.ovClient.GetStorageVolumeTemplateByName(id)
	if err != nil {
		d.SetId("")
		return nil
	}

	d.SetId(id)
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
		f_item := template.TemplateProperties.Folder
		if f_item != nil {
			folder := make([]map[string]interface{}, 0)
			enumMap := make([]interface{}, len(f_item.Enum))
			for i, enum := range f_item.Enum {
				enumMap[i] = enum
			}

			folder = append(folder, map[string]interface{}{
				"meta_create_only":   f_item.Meta.CreateOnly,
				"meta_locked":        f_item.Meta.Locked,
				"meta_semantic_type": f_item.Meta.SemanticType,
				"type":               f_item.Type,
				"title":              f_item.Title,
				"required":           f_item.Required,
				"max_length":         f_item.Maxlength,
				"min_length":         f_item.Minlength,
				"description":        f_item.Description,
				"enum":               schema.NewSet(schema.HashString, enumMap),
				"default":            f_item.Default,
				"minimum":            f_item.Minimum,
				"format":             f_item.Format})

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
