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

func dataSourceInterconnects() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceInterconnectsRead,

		Schema: map[string]*schema.Schema{
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enclosure_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enclosure_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enclosure_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"interconnect_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"interconnect_location": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"location_entries": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"interconnect_mac": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"interconnect_type_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_address_list": {
				Type:     schema.TypeList,
				Computed: true,
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
				Computed: true,
			},
			"lldp_ipv4_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lldp_ipv6_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"logical_interconnect_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"model": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"product_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"roles": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"scopes_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"snmp_configuration": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_access": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						"trap_destinations": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"community_string": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"enet_trap_categories": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
									"fc_trap_categories": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
									"vcm_trap_categories": {
										Type:     schema.TypeSet,
										Computed: true,
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
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
								},
							},
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"v3_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
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

	d.Set("category", interconnect.Category)
	d.Set("description", interconnect.Description)
	d.Set("enclosure_name", interconnect.EnclosureName)
	d.Set("enclosure_type", interconnect.EnclosureType)
	d.Set("enclosure_uri", interconnect.EnclosureUri.String())
	d.Set("host_name", interconnect.HostName)
	d.Set("interconnect_ip", interconnect.InterconnectIP)
	d.Set("interconnect_mac", interconnect.InterconnectMAC)
	d.Set("interconnect_type_uri", interconnect.InterconnectTypeUri.String())
	d.Set("lldp_ip_address_mode", interconnect.LldpIpAddressMode)
	d.Set("lldp_ipv4_address", interconnect.LldpIpv4Address)
	d.Set("lldp_ipv6_address", interconnect.LldpIpv6Address)
	d.Set("logical_interconnect_uri", interconnect.LogicalInterconnectUri.String())
	d.Set("model", interconnect.Model)
	d.Set("name", interconnect.Name)
	d.Set("product_name", interconnect.ProductName)
	d.Set("scopes_uri", interconnect.ScopesUri)
	d.Set("type", interconnect.Type)
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
		"category":          interconnect.SnmpConfiguration.Category,
		"description":       interconnect.SnmpConfiguration.Description,
		"name":              interconnect.SnmpConfiguration.Name,
		"uri":               interconnect.SnmpConfiguration.URI,
		"enabled":           interconnect.SnmpConfiguration.Enabled,
		"v3_enabled":        interconnect.SnmpConfiguration.V3Enabled,
		"snmp_access":       schema.NewSet(schema.HashString, snmpAccess),
		"type":              interconnect.SnmpConfiguration.Type,
		"trap_destinations": trapDestinations,
	})
	d.Set("snmp_configuration", snmpConfiguration)

	return nil
}
