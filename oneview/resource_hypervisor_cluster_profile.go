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
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceHypervisorClusterProfile() *schema.Resource {

	return &schema.Resource{
		Read:   resourceHypervisorClusterProfileRead,
		Create: resourceHypervisorClusterProfileCreate,
		Update: resourceHypervisorClusterProfileUpdate,
		Delete: resourceHypervisorClusterProfileDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"add_host_requests": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},

			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"compliance_state": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"e_tag": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"hypervisor_cluster_settings": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"distributed_switch_usage": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"distributed_switch_version": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"drs_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"ha_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"multi_nic_v_motion": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"virtual_switch_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"hypervisor_host_profile_template": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deployment_manager_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"deployment_plan": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"deployment_custom_args": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"deployment_plan_description": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"deployment_plan_uri": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"server_password": {
										Type:     schema.TypeString,
										Optional: true,
									},
								}}},

						"host_prefix": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"server_profile_template_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
					}}},
			"host_config_policy": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"leave_host_in_maintenance": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"use_host_prefix_as_hostname": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"use_hostname_to_register": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					}}},
			"virtual_switch_config_policy": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"configure_port_groups": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"custom_virtual_switches": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"manage_virtual_switches": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					}}},
			"hypervisor_cluster_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_switches": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"network_uris": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"virtual_switch_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"virtual_switch_uplinks": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"active": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"mac": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"vmnic": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"virtual_switch_port_groups": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"network_uris": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"virtual_switch_ports": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"action": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"dhcp": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"ip_address": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"subnet_mask": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"virtual_port_purpose": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										}},
									"vlan": {
										Type:     schema.TypeString,
										Optional: true,
									},
								}},
						},
					},
				}},

			"hypervisor_host_profile_uris": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"hypervisor_manager_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"hypervisor_type": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"ip_pools": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString},
				Set: schema.HashString,
			},

			"mgmt_ip_settings_override": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"path": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"refresh_state": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"scopes_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"shared_storage_volumes": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},

			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"state_reason": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

