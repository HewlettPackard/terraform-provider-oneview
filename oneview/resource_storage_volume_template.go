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
	"io/ioutil"
	//"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform/helper/schema"
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
				Optional: true,
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
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
			},
			"family": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tp_name": {
				Optional: true,
				Type:     schema.TypeSet,
				//MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"meta_semanctic_type": {
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
						"meta_semanctic_type": {
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
					},
				},
			},
			"tp_size": {
				Optional: true,
				Type:     schema.TypeSet,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"meta_locked": {
							Optional: true,
							Type:     schema.TypeBool,
						},
						"meta_semanctic_type": {
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
							Type:     schema.TypeInt,
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
						"meta_semanctic_type": {
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
						"meta_semanctic_type": {
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
						"meta_semanctic_type": {
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
							Type:     schema.TypeBool,
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
						"meta_semanctic_type": {
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
							Type:     schema.TypeBool,
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
						"meta_semanctic_type": {
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
							Type:     schema.TypeBool,
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
						"meta_semanctic_type": {
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
						"meta_semanctic_type": {
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
						"meta_semanctic_type": {
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
						"meta_semanctic_type": {
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
						"meta_semanctic_type": {
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
						"meta_semanctic_type": {
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
						"meta_semanctic_type": {
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
							Type:     schema.TypeBool,
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
						"meta_semanctic_type": {
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
							Type:     schema.TypeBool,
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
						"meta_semanctic_type": {
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
						"meta_semanctic_type": {
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
							Type:     schema.TypeBool,
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
	name_item := ov.TemplatePropertyDatatypeStruct{}
	if val, ok := d.GetOk("tp_name"); ok {
		rawName := val.(*schema.Set).List()
		name := make([]ov.TemplatePropertyDatatypeStruct, 1)
		for _, rawData := range rawName {
			item := rawData.(map[string]interface{})
			//			name_item := ov.TemplatePropertyDatatypeStruct{}
			meta := ov.Meta{}

			b := []byte(item["type"].(string) + "::" + item["title"].(string))
			ioutil.WriteFile("fields.txt", b, 0644)
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
			if item["meta_semanctic_type"] != nil {
				meta.SemanticType = item["meta_semanctic_type"].(string)
			}
			if item["enum"] != nil {
				rawEnums := item["enum"].(*schema.Set).List()
				enums := make([]string, len(rawEnums))
				for i, rawData := range rawEnums {
					enums[i] = rawData.(string)
				}
				name_item.Enum = enums
			}
			//if meta != nil {
			name_item.Meta = &meta
			//}
			file, _ := json.MarshalIndent(name_item, "", " ")
			ioutil.WriteFile("name_item.txt", file, 0644)
			name = append(name, name_item)

		}
		properties[0].Name = &name_item //&name[0]
		file, _ := json.MarshalIndent(properties[0], "", " ")
		ioutil.WriteFile("name_o.txt", file, 0644)
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
	//	file,_ := json.MarshalIndent(template, "", " ")
	//	ioutil.WriteFile("in.txt", file, 0644)

	//bb := []byte(err.Error())
	//	ioutil.WriteFile("err.txt",bb, 0644)
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
		//for _,item := range rawName {
		var meta_create_only, meta_locked, required bool
		var meta_semanctic_type, title, description, tp_type, tp_default, format string
		var enums []string
		var max_length, min_length, minimum int

		if item.Meta != nil {
			if len(item.Meta.SemanticType) != 0 {
				meta_semanctic_type = item.Meta.SemanticType
			}
			meta_create_only = item.Meta.CreateOnly
			meta_locked = item.Meta.Locked
		}
		if len(item.Type) != 0 {
			tp_type = item.Type
		}
		if len(item.Title) != 0 {
			title = item.Title
		}
		required = item.Required
		max_length = item.Maxlength
		min_length = item.Minlength
		minimum = item.Minimum
		//if item.Description != nil  {
		description = item.Description.String()
		//}
		if item.Enum != nil {
			enums = item.Enum
		}
		if len(item.Default) != 0 {
			tp_default = item.Default
		}
		if len(item.Format) != 0 {
			format = item.Format
		}
		name = append(name, map[string]interface{}{
			"meta_create_only":    meta_create_only,
			"meta_locked":         meta_locked,
			"meta_semanctic_type": meta_semanctic_type,
			"type":                tp_type,
			"title":               title,
			"required":            required,
			"max_length":          max_length,
			"min_length":          min_length,
			"description":         description,
			"enum":                enums,
			"default":             tp_default,
			"minimum":             minimum,
			"format":              format})
		file, _ := json.MarshalIndent(name, "", " ")
		ioutil.WriteFile("name_check.txt", file, 0644)
		b := []byte("Enetered if")
		ioutil.WriteFile("ifcond.txt", b, 0644)
		d.Set("tp_name", name)
		file, _ = json.MarshalIndent(d.Get("tp_name"), "", " ")
		ioutil.WriteFile("ifcond1.txt", file, 0644)
	}

	return nil
}

func resourceStorageVolumeTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceStorageVolumeTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
