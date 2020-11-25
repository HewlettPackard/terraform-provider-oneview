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
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceHypervisorClusterProfile() *schema.Resource {
	return &schema.Resource{
		Read: datasourceHypervisorClusterProfileRead,

		Schema: map[string]*schema.Schema{

			"add_host_requests": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString},
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
				Computed: true,
			},

			"e_tag": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"hypervisor_cluster_settings": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"distributed_switch_usage": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"distributed_switch_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"drs_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"ha_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},

						"multi_nic_v_motion": {
							Type:     schema.TypeBool,
							Computed: true,
						},

						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"virtual_switch_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"hypervisor_host_profile_template": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deployment_manager_type": {
							Type:     schema.TypeString,
							Computed: true},
						"deployment_plan": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"deployment_custom_args": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"deployment_plan_description": {
										Type:     schema.TypeString,
										Computed: true},
									"deployment_plan_uri": {
										Type:     schema.TypeString,
										Computed: true},
									"name": {
										Type:     schema.TypeString,
										Computed: true},
									"server_password": {
										Type:     schema.TypeString,
										Computed: true},
								}}},
						"host_config_policy": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"leave_host_in_maintenance": {
										Type:     schema.TypeBool,
										Computed: true},
									"use_host_prefix_as_hostname": {
										Type:     schema.TypeBool,
										Computed: true},
									"use_hostname_to_register": {
										Type:     schema.TypeBool,
										Computed: true},
								}}},
						"host_prefix": {
							Type:     schema.TypeString,
							Computed: true},
						"server_profile_template_uri": {
							Type:     schema.TypeString,
							Computed: true},
						"virtual_switch_config_policy": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"configure_port_group": {
										Type:     schema.TypeBool,
										Computed: true},
									"custom_virtual_switches": {
										Type:     schema.TypeBool,
										Computed: true},
									"manage_virtual_switches": {
										Type:     schema.TypeBool,
										Computed: true},
								}}},
						"virtual_switches": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:     schema.TypeString,
										Computed: true},
									"name": {
										Type:     schema.TypeString,
										Computed: true},
									"network_uris": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString},
									},
									"version": {
										Type:     schema.TypeString,
										Computed: true},
									"virtual_switch_port_groups": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"action": {
													Type:     schema.TypeString,
													Computed: true},
												"name": {
													Type:     schema.TypeString,
													Computed: true},
												"network_uris": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString},
												},
												"virtual_switch_ports": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"action": {
																Type:     schema.TypeString,
																Computed: true},
															"dhcp": {
																Type:     schema.TypeBool,
																Computed: true},
															"ip_address": {
																Type:     schema.TypeString,
																Computed: true},
															"subnet_mask": {
																Type:     schema.TypeString,
																Computed: true},
															"virtual_port_purpose": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString},
															},
														},
													}},
												"vlan": {
													Type:     schema.TypeString,
													Computed: true},
											}},
									},
									"virtual_switch_type": {
										Type:     schema.TypeString,
										Computed: true},
									"virtual_switch_uplinks": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"action": {
													Type:     schema.TypeString,
													Computed: true},
												"active": {
													Type:     schema.TypeBool,
													Computed: true},
												"mac": {
													Type:     schema.TypeString,
													Computed: true},
												"name": {
													Type:     schema.TypeString,
													Computed: true},
												"vmnic": {
													Type:     schema.TypeString,
													Computed: true},
											},
										},
									},
								},
							}},
					}}},
			"hypervisor_cluster_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"hypervisor_host_profile_uris": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"hypervisor_manager_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"hypervisor_type": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"ip_pools": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString},
				Set: schema.HashString,
			},

			"mgmt_ip_settings_override": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"refresh_state": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"scopes_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"shared_storage_volumes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString},
				Set: schema.HashString,
			},

			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"state_reason": {
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
				Optional: true,
			},
		},
	}
}

func datasourceHypervisorClusterProfileRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	hypCP, err := config.ovClient.GetHypervisorClusterProfileByName(d.Get("name").(string))
	if err != nil || hypCP.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.SetId(d.Get("name").(string))
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
	for i, deploymentCustomArg := range hypCP.HypervisorHostProfileTemplate.DeploymentPlan.DeploymentCustomArgs {
		deploymentCustomArgs[i] = deploymentCustomArg.String()
	}
	dplist := make([]map[string]interface{}, 0, 1)
	dplist = append(dplist, map[string]interface{}{

		"deployment_custom_args":      deploymentCustomArgs,
		"deployment_plan_description": hypCP.HypervisorHostProfileTemplate.DeploymentPlan.DeploymentPlanDescription,
		"deployment_plan_uri":         hypCP.HypervisorHostProfileTemplate.DeploymentPlan.DeploymentPlanUri.String(),
		"name":                        hypCP.HypervisorHostProfileTemplate.DeploymentPlan.Name,
		"server_password":             hypCP.HypervisorHostProfileTemplate.DeploymentPlan.ServerPassword,
	})
	hostConfigPolicylist := make([]map[string]interface{}, 0, 1)
	hostConfigPolicylist = append(hostConfigPolicylist, map[string]interface{}{
		"leave_host_in_maintenance":   hypCP.HypervisorHostProfileTemplate.HostConfigPolicy.LeaveHostInMaintenance,
		"use_host_prefix_as_hostname": hypCP.HypervisorHostProfileTemplate.HostConfigPolicy.LeaveHostInMaintenance,
		"use_hostname_to_register":    hypCP.HypervisorHostProfileTemplate.HostConfigPolicy.UseHostnameToRegister,
	})

	virtualSwitchConfigPolicylist := make([]map[string]interface{}, 0, 1)
	virtualSwitchConfigPolicylist = append(virtualSwitchConfigPolicylist, map[string]interface{}{
		"configure_port_group":    hypCP.HypervisorHostProfileTemplate.VirtualSwitchConfigPolicy.ConfigurePortGroups,
		"custom_virtual_switches": hypCP.HypervisorHostProfileTemplate.VirtualSwitchConfigPolicy.CustomVirtualSwitches,
		"manage_virtual_switches": hypCP.HypervisorHostProfileTemplate.VirtualSwitchConfigPolicy.ManageVirtualSwitches,
	})

	/**************** virtual switches*******************************/

	virtualSwitches := make([]map[string]interface{}, 0, len(hypCP.HypervisorHostProfileTemplate.VirtualSwitches))
	for _, virtualSwitch := range hypCP.HypervisorHostProfileTemplate.VirtualSwitches {

		/***********virtualswicth port group*****************************/

		virtualSwitchPortGroups := make([]map[string]interface{}, 0, len(virtualSwitch.VirtualSwitchPortGroups))
		for _, virtualSwitchPortGroup := range virtualSwitch.VirtualSwitchPortGroups {
			vspgnetworkUris := make([]interface{}, len(virtualSwitchPortGroup.NetworkUris))
			for i, vspgnetworkURI := range virtualSwitchPortGroup.NetworkUris {
				vspgnetworkUris[i] = vspgnetworkURI.String()
			}
			/***********vritual switch ports*********************/

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
			/*************virtual switch ports ends********************/
			virtualSwitchPortGroups = append(virtualSwitchPortGroups, map[string]interface{}{
				"action":               virtualSwitchPortGroup.Action,
				"name":                 virtualSwitchPortGroup.Name,
				"network_uris":         vspgnetworkUris,
				"virtual_switch_ports": virtualSwitchPorts,
				"vlan":                 virtualSwitchPortGroup.Vlan,
			})
		}

		/**********virtual switch port group ends*********/

		/**********Virtual switch uplink***********/
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

		/**********virtual switch upnlinks end************/

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

	/*****************virtual switch ends*************************************/

	hypCPHHPTList := make([]map[string]interface{}, 0, 1)
	hypCPHHPTList = append(hypCPHHPTList, map[string]interface{}{
		"deployment_manager_type":      hypCP.HypervisorHostProfileTemplate.DeploymentManagerType,
		"deployment_plan":              dplist,
		"host_config_policy":           hostConfigPolicylist,
		"host_prefix":                  hypCP.HypervisorHostProfileTemplate.Hostprefix,
		"server_profile_template_uri":  hypCP.HypervisorHostProfileTemplate.ServerProfileTemplateUri.String(),
		"virtual_switch_config_policy": virtualSwitchConfigPolicylist,
		"virtual_switches":             virtualSwitches,
	})

	d.Set("hypervisor_host_profile_template", hypCPHHPTList)
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
