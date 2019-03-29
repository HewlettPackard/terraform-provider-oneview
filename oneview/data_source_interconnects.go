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
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceInterconnects() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceInterconnectsRead,

		Schema: map[string]*schema.Schema{
			"base_wwn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"category": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_reset_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"eTag": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"edge_virtual_bridging_available": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_cut_through": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_fast_mac_cache_failover": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_igmp_snooping": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_network_loop_protection": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_pause_flood_protection": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_rich_tlv": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_storm_control": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_tagged_lldp": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enclosure_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enclosure_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enclosure_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"firmware_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"igmp_idle_timeout_interval": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"igmp_snooping_vlan_ids": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"initial_scope_uris": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"interconnect_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"interconnect_location": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"location_entries": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"value": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"interconnect_mac": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"interconnect_type_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_address_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ip_address_type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"lldp_ip_address_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lldp_ipv4_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lldp_ipv6_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"logical_interconnect_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_bandwidth": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mgmt_interface": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"migration_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"model": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_loop_protection_interval": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"part_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ports": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"associated_uplink_set_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"available": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"bay_number": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"capability": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						"category": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"config_port_types": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						"connector_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"eTag": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"interconnect_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"lag_id": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"lag_states": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"operational_speed": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"paired_port_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_health_status": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_monitor_config_info": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_running_capability_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_split_mode": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_status": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_status_reason": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port_type_extended": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"state": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"subports": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port_number": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"port_status": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"port_status_reason": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"vendor_specific_port_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlans": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"power_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"roles": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"scopes_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"serial_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"snmp_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"read_community": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"state": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"system_contact": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_access": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						"trap_destination": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"community_string": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"enet_trap_categories": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
									"fc_trap_categories": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
									"vcm_trap_categories": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
									"trap_destination": {
										Type:     schema.TypeString,
										Required: true,
									},
									"trap_format": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "SNMPv1",
									},
									"trap_severities": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
								},
							},
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"v3_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"spare_part_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stacking_domain_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"stacking_domain_role": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stacking_member_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"storm_control_polling_interval": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"storm_control_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sub_port_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"uid_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"unsupported_capabilities": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceInterconnectsRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)

	interconnect, err := config.ovClient.GetInterconnectByName(name)
	if err != nil || interconnect.URI.IsNil() {
		d.SetId("")

		return nil
	}

	d.SetId(name)

	d.Set("base_wwn", interconnect.BaseWWN)
	d.Set("category", interconnect.Category)
	d.Set("description", interconnect.Description)
	d.Set("device_reset_state", interconnect.DeviceResetState)
	d.Set("eTag", interconnect.ETag)
	d.Set("edge_virtual_bridging_available", interconnect.EdgeVirtualBridgingAvailable)
	d.Set("enable_cut_through", interconnect.EnableCutThrough)
	d.Set("enable_fast_mac_cache_failover", interconnect.EnableFastMacCacheFailover)
	d.Set("enable_igmp_snooping", interconnect.EnableIgmpSnooping)
	d.Set("enable_network_loop_protection", interconnect.EnableNetworkLoopProtection)
	d.Set("enable_pause_flood_protection", interconnect.EnablePauseFloodProtection)
	d.Set("enable_rich_tlv", interconnect.EnableRichTLV)
	d.Set("enable_storm_control", interconnect.EnableStormControl)
	d.Set("enable_tagged_lldp", interconnect.EnableTaggedLldp)
	d.Set("enclosure_name", interconnect.EnclosureName)
	d.Set("enclosure_type", interconnect.EnclosureType)
	d.Set("enclosure_uri", interconnect.EnclosureUri.String())
	d.Set("firmware_version", interconnect.FirmwareVersion)
	d.Set("host_name", interconnect.HostName)
	d.Set("igmp_idle_timeout_interval", interconnect.IgmpIdleTimeoutInterval)
	d.Set("igmp_snooping_vlan_ids", interconnect.IgmpSnoopingVlanIds)
	d.Set("interconnect_ip", interconnect.InterconnectIP)
	d.Set("interconnect_mac", interconnect.InterconnectMAC)
	d.Set("interconnect_type_uri", interconnect.InterconnectTypeUri.String())
	d.Set("lldp_ip_address_mode", interconnect.LldpIpAddressMode)
	d.Set("lldp_ipv4_address", interconnect.LldpIpv4Address)
	d.Set("lldp_ipv6_address", interconnect.LldpIpv6Address)
	d.Set("logical_interconnect_uri", interconnect.LogicalInterconnectUri.String())
	d.Set("max_bandwidth", interconnect.MaxBandwidth)
	d.Set("mgmt_interface", interconnect.MgmtInterface)
	d.Set("migration_state", interconnect.MigrationState)
	d.Set("model", interconnect.Model)
	d.Set("name", interconnect.Name)
	d.Set("network_loop_protection_interval", interconnect.NetworkLoopProtectionInterval)
	d.Set("part_number", interconnect.PartNumber)
	d.Set("port_count", interconnect.PortCount)
	d.Set("power_state", interconnect.PowerState)
	d.Set("product_name", interconnect.ProductName)
	d.Set("scopes_uri", interconnect.ScopesUri)
	d.Set("serial_number", interconnect.SerialNumber)
	d.Set("spare_part_number", interconnect.SparePartNumber)
	d.Set("stacking_domain_id", interconnect.StackingDomainId)
	d.Set("stacking_domain_role", interconnect.StackingDomainRole)
	d.Set("stacking_member_id", interconnect.StackingMemberId)
	d.Set("state", interconnect.State)
	d.Set("storm_control_polling_interval", interconnect.StormControlPollingInterval)
	d.Set("storm_control_threshold", interconnect.StormControlThreshold)
	d.Set("sub_port_count", interconnect.SubPortCount)
	d.Set("type", interconnect.Type)
	d.Set("uid_state", interconnect.UidState)
	d.Set("unsupported_capabilities", interconnect.UnsupportedCapabilities)
	d.Set("uri", interconnect.URI.String())

	locationEntries := make([]map[string]interface{}, 0,
		len(interconnect.InterconnectLocation.LocationEntries))
	for _, locationEntry := range interconnect.InterconnectLocation.LocationEntries {
		locationEntries = append(locationEntries, map[string]interface{}{
			"type":  locationEntry.Type,
			"value": locationEntry.Value,
		})
	}
	interconnectLocation := make([]map[string]interface{}, 0, 1)
	interconnectLocation = append(interconnectLocation, map[string]interface{}{
		"location_entries": locationEntries,
	})
	d.Set("interconnect_location", interconnectLocation)

	ipAddressList := make([]map[string]interface{}, 0,
		len(interconnect.IpAddressList))
	for _, ipAddress := range interconnect.IpAddressList {
		ipAddressList = append(ipAddressList, map[string]interface{}{
			"ip_address_type": ipAddress.IpAddressType,
			"ip_address":      ipAddress.IpAddress,
		})
	}

	d.Set("ip_address_list", ipAddressList)

	initialScopeUris := make([]interface{}, len(interconnect.InitialScopeUris))
	for i, initialScopeUri := range interconnect.InitialScopeUris {
		initialScopeUris[i] = initialScopeUri
	}

	d.Set("initial_scope_uris", initialScopeUris)

	trapDestinations := make([]map[string]interface{}, 0, 1)
	for _, trapDestination := range interconnect.SnmpConfiguration.TrapDestinations {

		enetTrapCategories := make([]interface{}, len(trapDestination.EnetTrapCategories))
		for i, enetTrapCategory := range trapDestination.EnetTrapCategories {
			enetTrapCategories[i] = enetTrapCategory
		}

		fcTrapCategories := make([]interface{}, len(trapDestination.FcTrapCategories))
		for i, fcTrapCategory := range trapDestination.FcTrapCategories {
			fcTrapCategories[i] = fcTrapCategory
		}

		vcmTrapCategories := make([]interface{}, len(trapDestination.VcmTrapCategories))
		for i, vcmTrapCategory := range trapDestination.VcmTrapCategories {
			vcmTrapCategories[i] = vcmTrapCategory
		}

		trapSeverities := make([]interface{}, len(trapDestination.TrapSeverities))
		for i, trapSeverity := range trapDestination.TrapSeverities {
			trapSeverities[i] = trapSeverity
		}

		trapDestinations = append(trapDestinations, map[string]interface{}{
			"trap_destination":     trapDestination.TrapDestination,
			"community_string":     trapDestination.CommunityString,
			"trap_format":          trapDestination.TrapFormat,
			"enet_trap_categories": schema.NewSet(schema.HashString, enetTrapCategories),
			"fc_trap_categories":   schema.NewSet(schema.HashString, fcTrapCategories),
			"vcm_trap_categories":  schema.NewSet(schema.HashString, vcmTrapCategories),
			"trap_severities":      schema.NewSet(schema.HashString, trapSeverities),
		})
	}

	snmpAccess := make([]interface{}, len(interconnect.SnmpConfiguration.SnmpAccess))
	for i, snmpAccessIP := range interconnect.SnmpConfiguration.SnmpAccess {
		snmpAccess[i] = snmpAccessIP
	}

	snmpConfiguration := make([]map[string]interface{}, 0, 1)
	snmpConfiguration = append(snmpConfiguration, map[string]interface{}{
		"enabled":          *interconnect.SnmpConfiguration.Enabled,
		"v3_enabled":       *interconnect.SnmpConfiguration.V3Enabled,
		"read_community":   interconnect.SnmpConfiguration.ReadCommunity,
		"snmp_access":      schema.NewSet(schema.HashString, snmpAccess),
		"system_contact":   interconnect.SnmpConfiguration.SystemContact,
		"type":             interconnect.SnmpConfiguration.Type,
		"trap_destination": trapDestinations,
	})
	d.Set("snmp_configuration", snmpConfiguration)

	ports := make([]map[string]interface{}, 0, len(interconnect.Ports))
	for _, port := range interconnect.Ports {
		subports := make([]map[string]interface{}, 0, len(port.SubPorts))
		for _, subport := range port.SubPorts {
			subports = append(subports, map[string]interface{}{
				"port_number":        subport.PortNumber,
				"port_status":        subport.PortStatus,
				"port_status_reason": subport.PortStatusReason,
			})
		}
		ports = append(ports, map[string]interface{}{
			"associated_uplink_set_uri":    port.AssociatedUplinkSetUri.String(),
			"available":                    port.Available,
			"bay_number":                   port.BayNumber,
			"category":                     port.Category,
			"connector_type":               port.ConnectorType,
			"description":                  port.Description,
			"eTag":                         port.ETag,
			"enabled":                      port.Enabled,
			"interconnect_name":            port.InterconnectName,
			"lag_id":                       port.LagId,
			"name":                         port.Name,
			"operational_speed":            port.OperationalSpeed,
			"paired_port_name":             port.PairedPortName,
			"port_health_status":           port.PortHealthStatus,
			"port_id":                      port.PortId,
			"port_monitor_config_info":     port.PortMonitorConfigInfo,
			"port_name":                    port.PortName,
			"port_running_capability_type": port.PortRunningCapabilityType,
			"port_split_mode":              port.PortSplitMode,
			"port_status":                  port.PortStatus,
			"port_status_reason":           port.PortStatusReason,
			"port_type":                    port.PortType,
			"port_type_extended":           port.PortTypeExtended,
			"state":                        port.State,
			"type":                         port.Type,
			"uri":                          port.URI.String(),
			"vendor_specific_port_name":    port.VendorSpecificPortName,
			"vlans":                        port.Vlans,
			"subports":                     subports,
		})
	}
	d.Set("ports", ports)

	return nil
}
