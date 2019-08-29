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
	"encoding/json"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"reflect"
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
			"eTag": {
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
				}
				is_item.Enum = enums
			}
			is_item.Meta = &meta
		}
		properties[0].IsShareable = &is_item
	}

	template.TemplateProperties = &properties[0]

	file, _ := json.MarshalIndent(template, "", " ")
	ioutil.WriteFile("ian.txt", file, 0644)
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
	d.Set("eTag", template.ETAG)
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
		name := make([]map[string]interface{}, 0)

		name = append(name, map[string]interface{}{
			"meta_create_only":   item.Meta.CreateOnly,//meta_create_only,
			"meta_locked":        item.Meta.Locked,//meta_locked,
			"meta_semantic_type": item.Meta.SemanticType,//meta_semantic_type,
			"type":               item.Type,//tp_type,
			"title":              item.Title,//title,
			"required":           item.Required,//required,
			"max_length":         item.Maxlength,//max_length,
			"min_length":         item.Minlength,//min_length,
			"description":        item.Description,//description,
			"enum":               item.Enum,//enums,
			"default":            item.Default,//tp_default,
			"minimum":            item.Minimum,//minimum,
			"format":             item.Format})//format})
		x := []byte("Reached")
		ioutil.WriteFile("sample.txt",x, 0644)

		file, _ := json.MarshalIndent(name, "", " ")
		ioutil.WriteFile("name_check.txt", file, 0644)
		d.Set("tp_name", name)

		tp_name_type := "tp_name Type:: "+reflect.TypeOf(d.Get("tp_name")).String()
		bb := []byte(tp_name_type)
		ioutil.WriteFile("tp_name.txt", bb, 0644)

		type_name := "Name Type: "+reflect.TypeOf(name).String()
		bb = []byte(type_name)
		ioutil.WriteFile("type_name.txt", bb, 0644)

		ite := template.TemplateProperties.Size
                size := make([]map[string]interface{}, 0)

                size = append(size, map[string]interface{}{
                        "meta_create_only":   ite.Meta.CreateOnly,
                        "meta_locked":        ite.Meta.Locked,
                        "meta_semantic_type": ite.Meta.SemanticType,
                        "type":               ite.Type,
                        "title":              ite.Title,
                        "required":           ite.Required,
                        "description":        ite.Description,
                        "enum":               ite.Enum,
                        "default":            ite.Default,
                        "minimum":            ite.Minimum,
                        "maximum":            ite.Maximum,
                        "format":             ite.Format})
		d.Set("tp_size", size)

	}

	return nil
}

func resourceStorageVolumeTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	template := ov.StorageVolumeTemplate{
		Name: d.Get("name").(string),
		ETAG: d.Get("eTag").(string),
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
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
