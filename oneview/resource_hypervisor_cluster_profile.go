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
				Required: true,
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

func resourceHypervisorClusterProfileCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	hypCP := ov.HypervisorClusterProfile{
		Category:                  d.Get("category").(string),
		ComplianceState:           d.Get("compliance_state").(string),
		Created:                   d.Get("created").(string),
		Description:               utils.Nstring(d.Get("description").(string)),
		ETag:                      d.Get("e_tag").(string),
		HypervisorClusterUri:      d.Get("hypervisor_cluster_uri").(string),
		HypervisorHostProfileUris: utils.Nstring(d.Get("hypervisor_host_profile_uris").(string)),
		HypervisorManagerUri:      utils.Nstring(d.Get("hypervisor_manager_uri").(string)),
		HypervisorType:            d.Get("hypervisor_type").(string),
		MgmtIpSettingsOverride:    d.Get("mgmt_ip_settings_override").(string),
		Modified:                  d.Get("modified").(string),
		Name:                      d.Get("name").(string),
		Path:                      d.Get("path").(string),
		RefreshState:              d.Get("refresh_state").(string),
		ScopesUri:                 d.Get("scopes_uri").(string),
		State:                     d.Get("state").(string),
		StateReason:               d.Get("state_reason").(string),
		Status:                    d.Get("status").(string),
		Type:                      d.Get("type").(string),
	}

	/*********************Add hostrequests start***********************/
	rawAddHostRequests := d.Get("add_host_requests").(*schema.Set).List()
	addHostRequests := make([]string, len(rawAddHostRequests))
	for i, rawadd := range rawAddHostRequests {
		addHostRequests[i] = rawadd.(string)
	}
	hypCP.AddHostRequests = addHostRequests
	/*********************Add hostrequests end***********************/

	/*********************IP Pools start***********************/
	rawIPPools := d.Get("ip_pools").(*schema.Set).List()
	ipPools := make([]utils.Nstring, len(rawIPPools))
	for i, rawip := range rawIPPools {
		ipPools[i] = utils.Nstring(rawip.(string))
	}
	hypCP.IpPools = ipPools
	/*********************Ip Pools end***********************/
	/*********************Shared Storage volumes start***********************/
	rawSharedStorageVolumes := d.Get("ip_pools").(*schema.Set).List()
	sharedStorageVolumes := make([]utils.Nstring, len(rawSharedStorageVolumes))
	for i, rawshsvol := range rawSharedStorageVolumes {
		sharedStorageVolumes[i] = utils.Nstring(rawshsvol.(string))
	}
	hypCP.SharedStorageVolumes = sharedStorageVolumes
	/*********************Shared Storage Volumes end***********************/
	/*********************Hypervisor cluster settings start***********************/

	hypClusterSettings := ov.HypervisorClusterSettings{}
	HypervisorClusterSettingslist := d.Get("hypervisor_cluster_settings").(*schema.Set).List()
	for _, raw := range HypervisorClusterSettingslist {
		hypervisorClusterSettings := raw.(map[string]interface{})

		hypClusterSettings = ov.HypervisorClusterSettings{
			DistributedSwitchUsage:   hypervisorClusterSettings["distributed_switch_usage"].(string),
			DistributedSwitchVersion: hypervisorClusterSettings["distributed_switch_version"].(string),
			DrsEnabled:               hypervisorClusterSettings["drs_enabled"].(bool),
			HaEnabled:                hypervisorClusterSettings["ha_enabled"].(bool),
			MultiNicVMotion:          hypervisorClusterSettings["multi_nic_v_motion"].(bool),
			Type:                     hypervisorClusterSettings["type"].(string),
			VirtualSwitchType:        hypervisorClusterSettings["virtual_switch_type"].(string),
		}
	}
	hypCP.HypervisorClusterSettings = &hypClusterSettings
	/*********************Hypervisor cluster settings ends***********************/
	var virtualSwitchesCollect []ov.VirtualSwitches
	/*****************************virtual switch********************************/
	if val, ok := d.GetOk("virtual_switches"); ok {
		rawVirtualSwitches := val.(*schema.Set).List()
		virtualSwitchesCollect = make([]ov.VirtualSwitches, 0)
		virtualSwitchUplinks := ov.VirtualSwitchUplinks{}

		virtualSwitches := ov.VirtualSwitches{}
		for _, rawvs := range rawVirtualSwitches {

			rawVirtualSwitchesItem := rawvs.(map[string]interface{})
			/*********************virtual switch networkUris start***********************/
			var networkUris []utils.Nstring
			if rawVirtualSwitchesItem["network_uris"] != nil {
				rawNetworkUris := rawVirtualSwitchesItem["network_uris"].([]interface{})
				networkUris = make([]utils.Nstring, len(rawNetworkUris))
				for i, rawvsnw := range rawNetworkUris {
					networkUris[i] = utils.Nstring(rawvsnw.(string))
				}
			}
			/******************virtual switch network uris end*************************/
			/***************virtual switch uplinks start***************************/
			{
				var virtualSwitchUplinksCollect []ov.VirtualSwitchUplinks
				if rawVirtualSwitchesItem["virtual_switch_uplinks"] != nil {

					rawVirtualSwitchUplinks := rawVirtualSwitchesItem["virtual_switch_uplinks"].([]interface{})
					virtualSwitchUplinksCollect = make([]ov.VirtualSwitchUplinks, 0)

					for _, rawvsul := range rawVirtualSwitchUplinks {
						rawVirtualSwitchUplinksItem := rawvsul.(map[string]interface{})
						virtualSwitchUplinks = ov.VirtualSwitchUplinks{
							Action: rawVirtualSwitchUplinksItem["action"].(string),
							Active: rawVirtualSwitchUplinksItem["active"].(bool),
							Mac:    rawVirtualSwitchUplinksItem["mac"].(string),
							Name:   rawVirtualSwitchUplinksItem["name"].(string),
							Vmnic:  rawVirtualSwitchUplinksItem["vmnic"].(string),
						}
						virtualSwitchUplinksCollect = append(virtualSwitchUplinksCollect, virtualSwitchUplinks)
					}
				}
				/***************virtual switch uplinks ends***************************/
				/****************virtual port groups***********************************/

				rawVirtualSwitchPortGroups := rawVirtualSwitchesItem["virtual_switch_port_groups"].([]interface{})
				virtualSwitchPortGroupsCollect := make([]ov.VirtualSwitchPortGroups, 0)
				virtualSwitchPortGroups := ov.VirtualSwitchPortGroups{}
				virtualSwitchPorts := ov.VirtualSwitchPorts{}

				for _, rawvspg := range rawVirtualSwitchPortGroups {
					rawVirtualSwitchPortGroupsItem := rawvspg.(map[string]interface{})
					/*********************virtual switch port group networkUris start***********************/
					rawNetworkUris := rawVirtualSwitchPortGroupsItem["network_uris"].([]interface{})
					networkUris := make([]utils.Nstring, len(rawNetworkUris))

					for i, rawvspgnw := range rawNetworkUris {
						networkUris[i] = utils.Nstring(rawvspgnw.(string))
					}
					/******************virtual switch port groups network uris end*************************/
					/*************virtual switch ports start******************************************/
					rawVirtualSwitchPorts := rawVirtualSwitchPortGroupsItem["virtual_switch_ports"].([]interface{})
					virtualSwitchPortsCollect := make([]ov.VirtualSwitchPorts, 0)

					for _, rawvsp := range rawVirtualSwitchPorts {
						rawVirtualSwitchPortsItem := rawvsp.(map[string]interface{})
						/*********************virtual switch port purpose start***********************/
						rawVirtualPortPurpose := rawVirtualSwitchPortsItem["virtual_port_purpose"].([]interface{})
						virtualPortPurpose := make([]utils.Nstring, len(rawVirtualPortPurpose))
						for i, rawpp := range rawVirtualPortPurpose {
							virtualPortPurpose[i] = utils.Nstring(rawpp.(string))
						}
						/******************virtual switch port purpose end*************************/
						virtualSwitchPorts = ov.VirtualSwitchPorts{
							Action:             rawVirtualSwitchPortsItem["action"].(string),
							Dhcp:               rawVirtualSwitchPortsItem["dhcp"].(bool),
							IpAddress:          rawVirtualSwitchPortsItem["ip_address"].(string),
							SubnetMask:         rawVirtualSwitchPortsItem["subnet_mask"].(string),
							VirtualPortPurpose: virtualPortPurpose,
						}
						virtualSwitchPortsCollect = append(virtualSwitchPortsCollect, virtualSwitchPorts)
					}
					/***************virtual switch ports ends***************************/

					virtualSwitchPortGroups = ov.VirtualSwitchPortGroups{
						Action:             rawVirtualSwitchPortGroupsItem["action"].(string),
						Name:               rawVirtualSwitchPortGroupsItem["name"].(string),
						NetworkUris:        networkUris,
						VirtualSwitchPorts: virtualSwitchPortsCollect,
						Vlan:               rawVirtualSwitchPortGroupsItem["vlan"].(string),
					}
					virtualSwitchPortGroupsCollect = append(virtualSwitchPortGroupsCollect, virtualSwitchPortGroups)
				}
				/****************virtual port groups  ends***********************************/
				virtualSwitches = ov.VirtualSwitches{
					Action:                  rawVirtualSwitchesItem["action"].(string),
					Name:                    rawVirtualSwitchesItem["name"].(string),
					NetworkUris:             networkUris,
					Version:                 rawVirtualSwitchesItem["version"].(string),
					VirtualSwitchPortGroups: virtualSwitchPortGroupsCollect,
					VirtualSwitchType:       rawVirtualSwitchesItem["virtual_switch_type"].(string),
					VirtualSwitchUplinks:    virtualSwitchUplinksCollect,
				}
				virtualSwitchesCollect = append(virtualSwitchesCollect, virtualSwitches)

			}
		}
	}

	/*************virtual switch ends******************************/
	/*****************switch config policy**************************/
	virtualSwitchConfigPolicy := ov.VirtualSwitchConfigPolicy{}

	rawVirtualSwitchConfigPolicy := d.Get("virtual_switch_config_policy").(*schema.Set).List()

	for _, rawscp := range rawVirtualSwitchConfigPolicy {
		rawVirtualSwitchConfigPolicyItem := rawscp.(map[string]interface{})

		virtualSwitchConfigPolicy = ov.VirtualSwitchConfigPolicy{
			ConfigurePortGroups:   rawVirtualSwitchConfigPolicyItem["configure_port_groups"].(bool),
			CustomVirtualSwitches: rawVirtualSwitchConfigPolicyItem["custom_virtual_switches"].(bool),
			ManageVirtualSwitches: rawVirtualSwitchConfigPolicyItem["manage_virtual_switches"].(bool),
		}
	}

	/*****************switch config policy ends**************************/
	/*****************host config policy**************************/
	hostConfigPolicy := ov.HostConfigPolicy{}
	rawHostConfigPolicy := d.Get("host_config_policy").(*schema.Set).List()

	for _, rawhcp := range rawHostConfigPolicy {
		rawHostConfigPolicyItem := rawhcp.(map[string]interface{})

		hostConfigPolicy = ov.HostConfigPolicy{
			LeaveHostInMaintenance:  rawHostConfigPolicyItem["leave_host_in_maintenance"].(bool),
			UseHostPrefixAsHostname: rawHostConfigPolicyItem["use_host_prefix_as_hostname"].(bool),
			UseHostnameToRegister:   rawHostConfigPolicyItem["use_hostname_to_register"].(bool),
		}
	}

	/*****************host config policy**************************/
	/*********************Hypervisor Host Profile Template Start***********************/
	rawHypervisorHostProfileTemplate := d.Get("hypervisor_host_profile_template").(*schema.Set).List()
	hypervisorProfileTemplate := ov.HypervisorHostProfileTemplate{}

	for _, rawht := range rawHypervisorHostProfileTemplate {
		rawHostProfileTemplateItem := rawht.(map[string]interface{})
		deploymentPlan := ov.DeploymentPlan{}

		/******************* deployment plan start********************/
		rawDeploymentPlan := rawHostProfileTemplateItem["deployment_plan"].([]interface{})

		for _, rawdp := range rawDeploymentPlan {
			rawDeploymentPlanItem := rawdp.(map[string]interface{})
			if val, ok := rawDeploymentPlanItem["deployment_custom_args"]; ok {
				dpCustomArgsOrder := val.([]interface{})
				dpCustomArgs := make([]utils.Nstring, len(dpCustomArgsOrder))
				for i, rawCustomArgs := range dpCustomArgsOrder {
					dpCustomArgs[i] = utils.Nstring(rawCustomArgs.(string))
				}

				deploymentPlan.DeploymentCustomArgs = dpCustomArgs
			}
			deploymentPlan = ov.DeploymentPlan{
				DeploymentPlanDescription: rawDeploymentPlanItem["deployment_plan_description"].(string),
				DeploymentPlanUri:         utils.Nstring(rawDeploymentPlanItem["deployment_plan_uri"].(string)),
				Name:                      rawDeploymentPlanItem["name"].(string),
				ServerPassword:            rawDeploymentPlanItem["server_password"].(string),
			}
		}
		/******************* deployment plan end********************/

		hypervisorProfileTemplate = ov.HypervisorHostProfileTemplate{
			DeploymentManagerType:     rawHostProfileTemplateItem["deployment_manager_type"].(string),
			DeploymentPlan:            &deploymentPlan,
			Hostprefix:                rawHostProfileTemplateItem["host_prefix"].(string),
			ServerProfileTemplateUri:  utils.Nstring(rawHostProfileTemplateItem["server_profile_template_uri"].(string)),
			VirtualSwitchConfigPolicy: &virtualSwitchConfigPolicy,
			HostConfigPolicy:          &hostConfigPolicy,
			VirtualSwitches:           virtualSwitchesCollect,
		}

	}

	/*********************Hypervisor Host Profile Template end***********************/

	hypCP.HypervisorHostProfileTemplate = &hypervisorProfileTemplate
	hypCPError := config.ovClient.CreateHypervisorClusterProfile(hypCP)
	d.SetId(d.Get("name").(string))
	if hypCPError != nil {
		d.SetId("")
		return hypCPError
	}
	return resourceHypervisorClusterProfileRead(d, meta)

}

func resourceHypervisorClusterProfileRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	hypCP, err := config.ovClient.GetHypervisorClusterProfileByName(d.Id())
	if err != nil || hypCP.URI.IsNil() {
		d.SetId("")
		return nil
	}
	addHostRequests := make([]interface{}, len(hypCP.AddHostRequests))
	for i, addHostRequest := range hypCP.AddHostRequests {
		addHostRequests[i] = addHostRequest
	}
	d.Set("add_host_requests", addHostRequests)
	d.Set("category", hypCP.Category)
	d.Set("compliance_state", hypCP.ComplianceState)
	d.Set("created", hypCP.Created)
	d.Set("description", hypCP.Description.String())
	d.Set("e_tag", hypCP.ETag)
	hypCPCSList := make([]map[string]interface{}, 0, 1)
	hypCPCSList = append(hypCPCSList, map[string]interface{}{
		"distributed_switch_version": hypCP.HypervisorClusterSettings.DistributedSwitchVersion,
		"distributed_switch_usage":   hypCP.HypervisorClusterSettings.DistributedSwitchUsage,
		"drs_enabled":                hypCP.HypervisorClusterSettings.DrsEnabled,
		"ha_enabled":                 hypCP.HypervisorClusterSettings.HaEnabled,
		"multi_nic_v_motion":         hypCP.HypervisorClusterSettings.MultiNicVMotion,
		"type":                       hypCP.HypervisorClusterSettings.Type,
		"virtual_switch_type":        hypCP.HypervisorClusterSettings.VirtualSwitchType,
	})

	d.Set("hypervisor_cluster_settings", hypCPCSList)

	d.Set("hypervisor_cluster_uri", hypCP.HypervisorClusterUri)
	deploymentCustomArgs := make([]interface{}, len(hypCP.HypervisorHostProfileTemplate.DeploymentPlan.DeploymentCustomArgs))
	dplist := make([]map[string]interface{}, 0, 1)
	if hypCP.HypervisorHostProfileTemplate.DeploymentPlan != nil {
		for i, deploymentCustomArg := range hypCP.HypervisorHostProfileTemplate.DeploymentPlan.DeploymentCustomArgs {
			deploymentCustomArgs[i] = deploymentCustomArg
		}
		dplist = append(dplist, map[string]interface{}{

			"deployment_custom_args":      deploymentCustomArgs,
			"deployment_plan_description": hypCP.HypervisorHostProfileTemplate.DeploymentPlan.DeploymentPlanDescription,
			"deployment_plan_uri":         hypCP.HypervisorHostProfileTemplate.DeploymentPlan.DeploymentPlanUri.String(),
			"name":                        hypCP.HypervisorHostProfileTemplate.DeploymentPlan.Name,
			"server_password":             hypCP.HypervisorHostProfileTemplate.DeploymentPlan.ServerPassword,
		})
	}
	hostConfigPolicylist := make([]map[string]interface{}, 0, 1)
	hostConfigPolicylist = append(hostConfigPolicylist, map[string]interface{}{
		"leave_host_in_maintenance":   hypCP.HypervisorHostProfileTemplate.HostConfigPolicy.LeaveHostInMaintenance,
		"use_host_prefix_as_hostname": hypCP.HypervisorHostProfileTemplate.HostConfigPolicy.UseHostPrefixAsHostname,
		"use_hostname_to_register":    hypCP.HypervisorHostProfileTemplate.HostConfigPolicy.UseHostnameToRegister,
	})
	d.Set("host_config_policy", hostConfigPolicylist)

	virtualSwitchConfigPolicylist := make([]map[string]interface{}, 0, 1)
	virtualSwitchConfigPolicylist = append(virtualSwitchConfigPolicylist, map[string]interface{}{
		"configure_port_group":    hypCP.HypervisorHostProfileTemplate.VirtualSwitchConfigPolicy.ConfigurePortGroups,
		"custom_virtual_switches": hypCP.HypervisorHostProfileTemplate.VirtualSwitchConfigPolicy.CustomVirtualSwitches,
		"manage_virtual_switches": hypCP.HypervisorHostProfileTemplate.VirtualSwitchConfigPolicy.ManageVirtualSwitches,
	})
	d.Set("virtual_switch_config_policy", virtualSwitchConfigPolicylist)

	hypCPHHPTList := make([]map[string]interface{}, 0, 1)
	hypCPHHPTList = append(hypCPHHPTList, map[string]interface{}{
		"deployment_manager_type":     hypCP.HypervisorHostProfileTemplate.DeploymentManagerType,
		"deployment_plan":             dplist,
		"host_prefix":                 hypCP.HypervisorHostProfileTemplate.Hostprefix,
		"server_profile_template_uri": hypCP.HypervisorHostProfileTemplate.ServerProfileTemplateUri.String(),
	})
	d.Set("hypervisor_host_profile_template", hypCPHHPTList)

	//###########################virtual switches###########################

	virtualSwitches := make([]map[string]interface{}, 0, len(hypCP.HypervisorHostProfileTemplate.VirtualSwitches))
	for _, virtualSwitch := range hypCP.HypervisorHostProfileTemplate.VirtualSwitches {

		//#########################Virtual switch uplink####################
		virtualSwitchPortUplinks := make([]map[string]interface{}, 0, len(virtualSwitch.VirtualSwitchUplinks))
		for _, virtualSwitchPortUplink := range virtualSwitch.VirtualSwitchUplinks {
			virtualSwitchPortUplinks = append(virtualSwitchPortUplinks, map[string]interface{}{
				"action": virtualSwitchPortUplink.Action,
				"active": virtualSwitchPortUplink.Active,
				"mac":    virtualSwitchPortUplink.Mac,
				"name":   virtualSwitchPortUplink.Name,
				"vmnic":  virtualSwitchPortUplink.Vmnic,
			})
		}

		//#####################virtual switch upnlinks end#######################
		//####################virtualswicth port group##########################

		virtualSwitchPortGroups := make([]map[string]interface{}, 0, len(virtualSwitch.VirtualSwitchPortGroups))
		for _, virtualSwitchPortGroup := range virtualSwitch.VirtualSwitchPortGroups {
			vspgnetworkUris := make([]interface{}, len(virtualSwitchPortGroup.NetworkUris))
			for i, vspgnetworkURI := range virtualSwitchPortGroup.NetworkUris {
				vspgnetworkUris[i] = vspgnetworkURI.String()
			}
			//########################vritual switch ports####################################

			virtualSwitchPorts := make([]map[string]interface{}, 0, len(virtualSwitchPortGroup.VirtualSwitchPorts))
			for _, virtualSwitchPort := range virtualSwitchPortGroup.VirtualSwitchPorts {
				virtualPortPurposes := make([]interface{}, len(virtualSwitchPort.VirtualPortPurpose))
				for i, virtualPortPurpose := range virtualSwitchPort.VirtualPortPurpose {
					virtualPortPurposes[i] = virtualPortPurpose
				}
				virtualSwitchPorts = append(virtualSwitchPorts, map[string]interface{}{
					"action":               virtualSwitchPort.Action,
					"dhcp":                 virtualSwitchPort.Dhcp,
					"ip_address":           virtualSwitchPort.IpAddress,
					"subnet_mask":          virtualSwitchPort.SubnetMask,
					"virtual_port_purpose": virtualPortPurposes,
				})
			}
			//#########################virtual switch ports ends#############################
			virtualSwitchPortGroups = append(virtualSwitchPortGroups, map[string]interface{}{
				"action":               virtualSwitchPortGroup.Action,
				"name":                 virtualSwitchPortGroup.Name,
				"network_uris":         vspgnetworkUris,
				"virtual_switch_ports": virtualSwitchPorts,
				"vlan":                 virtualSwitchPortGroup.Vlan,
			})
		}

		//#############################virtual switch port group ends##########################

		networkUris := make([]interface{}, len(virtualSwitch.NetworkUris))
		for i, networkURI := range virtualSwitch.NetworkUris {
			networkUris[i] = networkURI
		}

		virtualSwitches = append(virtualSwitches, map[string]interface{}{
			"action":                     virtualSwitch.Action,
			"name":                       virtualSwitch.Name,
			"network_uris":               networkUris,
			"version":                    virtualSwitch.Version,
			"virtual_switch_port_groups": virtualSwitchPortGroups,
			"virtual_switch_type":        virtualSwitch.VirtualSwitchType,
			"virtual_switch_uplinks":     virtualSwitchPortUplinks,
		})

	}
	d.Set("virtual_switches", virtualSwitches)
	//#########################virtual switch ends############################

	d.Set("hypervisor_host_profile_uris", hypCP.HypervisorHostProfileUris)
	d.Set("hypervisor_manager_uri", hypCP.HypervisorManagerUri)
	d.Set("hypervisor_type", hypCP.HypervisorType)
	ipPools := make([]interface{}, len(hypCP.IpPools))
	for i, ipPool := range hypCP.IpPools {
		ipPools[i] = ipPool
	}
	d.Set("ip_pools", ipPools)
	d.Set("mgmt_ip_settings_override", hypCP.MgmtIpSettingsOverride)
	d.Set("modified", hypCP.Modified)
	d.Set("name", hypCP.Name)
	d.Set("path", hypCP.Path)
	d.Set("refresh_state", hypCP.RefreshState)
	d.Set("scopes_uri", hypCP.ScopesUri)
	sharedStorageVolumes := make([]interface{}, len(hypCP.SharedStorageVolumes))
	for i, sharedStorageVolume := range hypCP.SharedStorageVolumes {
		sharedStorageVolumes[i] = sharedStorageVolume
	}
	d.Set("shared_storage_volumes", sharedStorageVolumes)
	d.Set("state", hypCP.State)
	d.Set("state_reason", hypCP.StateReason)
	d.Set("status", hypCP.Status)
	d.Set("type", hypCP.Type)
	d.Set("uri", hypCP.URI)

	return nil

}

func resourceHypervisorClusterProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	hypCP := ov.HypervisorClusterProfile{
		Category:                  d.Get("category").(string),
		ComplianceState:           d.Get("compliance_state").(string),
		Created:                   d.Get("created").(string),
		Description:               utils.Nstring(d.Get("description").(string)),
		HypervisorClusterUri:      d.Get("hypervisor_cluster_uri").(string),
		HypervisorHostProfileUris: utils.Nstring(d.Get("hypervisor_host_profile_uris").(string)),
		HypervisorManagerUri:      utils.Nstring(d.Get("hypervisor_manager_uri").(string)),
		HypervisorType:            d.Get("hypervisor_type").(string),
		MgmtIpSettingsOverride:    d.Get("mgmt_ip_settings_override").(string),
		Name:                      d.Get("name").(string),
		Path:                      d.Get("path").(string),
		ScopesUri:                 d.Get("scopes_uri").(string),
		StateReason:               d.Get("state_reason").(string),
		Type:                      d.Get("type").(string),
		URI:                       utils.Nstring(d.Get("uri").(string)),
	}
	hypClusterSettings := ov.HypervisorClusterSettings{}
	HypervisorClusterSettingslist := d.Get("hypervisor_cluster_settings").(*schema.Set).List()
	for _, raw := range HypervisorClusterSettingslist {
		hypervisorClusterSettings := raw.(map[string]interface{})

		hypClusterSettings = ov.HypervisorClusterSettings{
			DistributedSwitchUsage:   hypervisorClusterSettings["distributed_switch_usage"].(string),
			DistributedSwitchVersion: hypervisorClusterSettings["distributed_switch_version"].(string),
			DrsEnabled:               hypervisorClusterSettings["drs_enabled"].(bool),
			HaEnabled:                hypervisorClusterSettings["ha_enabled"].(bool),
			MultiNicVMotion:          hypervisorClusterSettings["multi_nic_v_motion"].(bool),
			Type:                     hypervisorClusterSettings["type"].(string),
			VirtualSwitchType:        hypervisorClusterSettings["virtual_switch_type"].(string),
		}
	}
	hypCP.HypervisorClusterSettings = &hypClusterSettings

	var virtualSwitchesCollect []ov.VirtualSwitches
	/*****************************virtual switch********************************/
	if val, ok := d.GetOk("virtual_switches"); ok {
		rawVirtualSwitches := val.(*schema.Set).List()
		virtualSwitchesCollect = make([]ov.VirtualSwitches, 0)
		virtualSwitchUplinks := ov.VirtualSwitchUplinks{}

		virtualSwitches := ov.VirtualSwitches{}
		for _, rawvs := range rawVirtualSwitches {

			rawVirtualSwitchesItem := rawvs.(map[string]interface{})
			/*********************virtual switch networkUris start***********************/
			var networkUris []utils.Nstring
			if rawVirtualSwitchesItem["network_uris"] != nil {
				rawNetworkUris := rawVirtualSwitchesItem["network_uris"].([]interface{})
				networkUris = make([]utils.Nstring, len(rawNetworkUris))
				for i, rawvsnw := range rawNetworkUris {
					networkUris[i] = utils.Nstring(rawvsnw.(string))
				}
			}
			/******************virtual switch network uris end*************************/
			/***************virtual switch uplinks start***************************/
			{
				var virtualSwitchUplinksCollect []ov.VirtualSwitchUplinks
				if rawVirtualSwitchesItem["virtual_switch_uplinks"] != nil {

					rawVirtualSwitchUplinks := rawVirtualSwitchesItem["virtual_switch_uplinks"].([]interface{})
					virtualSwitchUplinksCollect = make([]ov.VirtualSwitchUplinks, 0)

					for _, rawvsul := range rawVirtualSwitchUplinks {
						rawVirtualSwitchUplinksItem := rawvsul.(map[string]interface{})
						virtualSwitchUplinks = ov.VirtualSwitchUplinks{
							Action: rawVirtualSwitchUplinksItem["action"].(string),
							Active: rawVirtualSwitchUplinksItem["active"].(bool),
							Mac:    rawVirtualSwitchUplinksItem["mac"].(string),
							Name:   rawVirtualSwitchUplinksItem["name"].(string),
							Vmnic:  rawVirtualSwitchUplinksItem["vmnic"].(string),
						}
						virtualSwitchUplinksCollect = append(virtualSwitchUplinksCollect, virtualSwitchUplinks)
					}
				}
				/***************virtual switch uplinks ends***************************/
				/****************virtual port groups***********************************/

				rawVirtualSwitchPortGroups := rawVirtualSwitchesItem["virtual_switch_port_groups"].([]interface{})
				virtualSwitchPortGroupsCollect := make([]ov.VirtualSwitchPortGroups, 0)
				virtualSwitchPortGroups := ov.VirtualSwitchPortGroups{}
				virtualSwitchPorts := ov.VirtualSwitchPorts{}

				for _, rawvspg := range rawVirtualSwitchPortGroups {
					rawVirtualSwitchPortGroupsItem := rawvspg.(map[string]interface{})
					/*********************virtual switch port group networkUris start***********************/
					rawNetworkUris := rawVirtualSwitchPortGroupsItem["network_uris"].([]interface{})
					networkUris := make([]utils.Nstring, len(rawNetworkUris))

					for i, rawvspgnw := range rawNetworkUris {
						networkUris[i] = utils.Nstring(rawvspgnw.(string))
					}
					/******************virtual switch port groups network uris end*************************/
					/*************virtual switch ports start******************************************/
					rawVirtualSwitchPorts := rawVirtualSwitchPortGroupsItem["virtual_switch_ports"].([]interface{})
					virtualSwitchPortsCollect := make([]ov.VirtualSwitchPorts, 0)

					for _, rawvsp := range rawVirtualSwitchPorts {
						rawVirtualSwitchPortsItem := rawvsp.(map[string]interface{})
						/*********************virtual switch port purpose start***********************/
						rawVirtualPortPurpose := rawVirtualSwitchPortsItem["virtual_port_purpose"].([]interface{})
						virtualPortPurpose := make([]utils.Nstring, len(rawVirtualPortPurpose))
						for i, rawpp := range rawVirtualPortPurpose {
							virtualPortPurpose[i] = utils.Nstring(rawpp.(string))
						}
						/******************virtual switch port purpose end*************************/
						virtualSwitchPorts = ov.VirtualSwitchPorts{
							Action:             rawVirtualSwitchPortsItem["action"].(string),
							Dhcp:               rawVirtualSwitchPortsItem["dhcp"].(bool),
							IpAddress:          rawVirtualSwitchPortsItem["ip_address"].(string),
							SubnetMask:         rawVirtualSwitchPortsItem["subnet_mask"].(string),
							VirtualPortPurpose: virtualPortPurpose,
						}
						virtualSwitchPortsCollect = append(virtualSwitchPortsCollect, virtualSwitchPorts)
					}
					/***************virtual switch ports ends***************************/

					virtualSwitchPortGroups = ov.VirtualSwitchPortGroups{
						Action:             rawVirtualSwitchPortGroupsItem["action"].(string),
						Name:               rawVirtualSwitchPortGroupsItem["name"].(string),
						NetworkUris:        networkUris,
						VirtualSwitchPorts: virtualSwitchPortsCollect,
						Vlan:               rawVirtualSwitchPortGroupsItem["vlan"].(string),
					}
					virtualSwitchPortGroupsCollect = append(virtualSwitchPortGroupsCollect, virtualSwitchPortGroups)
				}
				/****************virtual port groups  ends***********************************/
				virtualSwitches = ov.VirtualSwitches{
					Action:                  rawVirtualSwitchesItem["action"].(string),
					Name:                    rawVirtualSwitchesItem["name"].(string),
					NetworkUris:             networkUris,
					Version:                 rawVirtualSwitchesItem["version"].(string),
					VirtualSwitchPortGroups: virtualSwitchPortGroupsCollect,
					VirtualSwitchType:       rawVirtualSwitchesItem["virtual_switch_type"].(string),
					VirtualSwitchUplinks:    virtualSwitchUplinksCollect,
				}
				virtualSwitchesCollect = append(virtualSwitchesCollect, virtualSwitches)

			}
		}
	}

	/*************virtual switch ends******************************/
	/*****************switch config policy**************************/
	virtualSwitchConfigPolicy := ov.VirtualSwitchConfigPolicy{}

	rawVirtualSwitchConfigPolicy := d.Get("virtual_switch_config_policy").(*schema.Set).List()

	for _, rawscp := range rawVirtualSwitchConfigPolicy {
		rawVirtualSwitchConfigPolicyItem := rawscp.(map[string]interface{})

		virtualSwitchConfigPolicy = ov.VirtualSwitchConfigPolicy{
			ConfigurePortGroups:   rawVirtualSwitchConfigPolicyItem["configure_port_groups"].(bool),
			CustomVirtualSwitches: rawVirtualSwitchConfigPolicyItem["custom_virtual_switches"].(bool),
			ManageVirtualSwitches: rawVirtualSwitchConfigPolicyItem["manage_virtual_switches"].(bool),
		}
	}

	/*****************switch config policy ends**************************/
	/*****************host config policy**************************/
	hostConfigPolicy := ov.HostConfigPolicy{}
	rawHostConfigPolicy := d.Get("host_config_policy").(*schema.Set).List()

	for _, rawhcp := range rawHostConfigPolicy {
		rawHostConfigPolicyItem := rawhcp.(map[string]interface{})

		hostConfigPolicy = ov.HostConfigPolicy{
			LeaveHostInMaintenance:  rawHostConfigPolicyItem["leave_host_in_maintenance"].(bool),
			UseHostPrefixAsHostname: rawHostConfigPolicyItem["use_host_prefix_as_hostname"].(bool),
			UseHostnameToRegister:   rawHostConfigPolicyItem["use_hostname_to_register"].(bool),
		}
	}

	/*****************host config policy**************************/
	/*********************Hypervisor Host Profile Template Start***********************/
	rawHypervisorHostProfileTemplate := d.Get("hypervisor_host_profile_template").(*schema.Set).List()
	hypervisorProfileTemplate := ov.HypervisorHostProfileTemplate{}

	for _, rawht := range rawHypervisorHostProfileTemplate {
		rawHostProfileTemplateItem := rawht.(map[string]interface{})
		deploymentPlan := ov.DeploymentPlan{}

		/******************* deployment plan start********************/
		rawDeploymentPlan := rawHostProfileTemplateItem["deployment_plan"].([]interface{})

		for _, rawdp := range rawDeploymentPlan {
			rawDeploymentPlanItem := rawdp.(map[string]interface{})
			if val, ok := rawDeploymentPlanItem["deployment_custom_args"]; ok {
				dpCustomArgsOrder := val.([]interface{})
				dpCustomArgs := make([]utils.Nstring, len(dpCustomArgsOrder))
				for i, rawCustomArgs := range dpCustomArgsOrder {
					dpCustomArgs[i] = utils.Nstring(rawCustomArgs.(string))
				}

				deploymentPlan.DeploymentCustomArgs = dpCustomArgs
			}
			deploymentPlan = ov.DeploymentPlan{
				DeploymentPlanDescription: rawDeploymentPlanItem["deployment_plan_description"].(string),
				DeploymentPlanUri:         utils.Nstring(rawDeploymentPlanItem["deployment_plan_uri"].(string)),
				Name:                      rawDeploymentPlanItem["name"].(string),
				ServerPassword:            rawDeploymentPlanItem["server_password"].(string),
			}
		}
		/******************* deployment plan end********************/

		hypervisorProfileTemplate = ov.HypervisorHostProfileTemplate{
			DeploymentManagerType:     rawHostProfileTemplateItem["deployment_manager_type"].(string),
			DeploymentPlan:            &deploymentPlan,
			Hostprefix:                rawHostProfileTemplateItem["host_prefix"].(string),
			ServerProfileTemplateUri:  utils.Nstring(rawHostProfileTemplateItem["server_profile_template_uri"].(string)),
			VirtualSwitchConfigPolicy: &virtualSwitchConfigPolicy,
			HostConfigPolicy:          &hostConfigPolicy,
			VirtualSwitches:           virtualSwitchesCollect,
		}

	}

	/*********************Hypervisor Host Profile Template end***********************/

	hypCP.HypervisorHostProfileTemplate = &hypervisorProfileTemplate
	hypCPError := config.ovClient.UpdateHypervisorClusterProfile(hypCP)
	d.SetId(d.Get("name").(string))
	if hypCPError != nil {
		d.SetId("")
		return hypCPError
	}
	return resourceHypervisorClusterProfileRead(d, meta)
}

func resourceHypervisorClusterProfileDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	var softdelete = true
	if config.OVAPIVersion == 1600 {

		err := config.ovClient.DeleteHypervisorClusterProfileSoftDelete(d.Id(), softdelete)
		if err != nil {
			return err
		}
	} else {
		err := config.ovClient.DeleteHypervisorClusterProfile(d.Id())
		if err != nil {
			return err
		}
	}
	return nil
}
