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
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"time"
)

func dataSourceEthernetNetwork() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEthernetNetworkRead,

		Schema: map[string]*schema.Schema{
			"count_": {
                                Type: schema.TypeInt,
                                Computed: true,
                        },
			"total": {
                                Type: schema.TypeInt,
                                Computed: true,
                        },
			"list_uri": {
                                Type: schema.TypeString,
                                Computed: true,
                        },
			"name": {
				Type: schema.TypeString,
                                Optional: true,
                        },
			"get_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"members": {
				Type: schema.TypeList,
				Computed: true,

				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vlan_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"purpose": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_network": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"smart_link": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"ethernet_network_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"connection_template_uri": {
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
						"uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"category": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fabric_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"etag": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"scopesuri": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

		},

	}
}

func dataSourceEthernetNetworkRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	getType := d.Get("get_type").(string)
	name := d.Get("name").(string)

	if getType == "GetById" {
		eNet , err := config.ovClient.GetEthernetNetworkByName(name)
		if err != nil {
                        d.SetId("")
                        return nil
		}
		members := make([]map[string]interface{}, 0, 1)
		members = append(members, map[string]interface{}{
                                        "name":   eNet.Name,
                                        "purpose": eNet.Purpose,
                                        "vlan_id": eNet.VlanId,
                                        "smart_link": eNet.SmartLink,
                                        "private_network": eNet.PrivateNetwork,
                                        "ethernet_network_type": eNet.EthernetNetworkType,
                                        "type": eNet.Type,
                                        "created": eNet.Created,
                                        "modified": eNet.Modified,
                                        "uri": eNet.URI,
                                        "connection_template_uri": eNet.ConnectionTemplateUri,
                                        "state":  eNet.State,
                                        "status": eNet.Status,
                                        "category": eNet.Category,
                                        "fabric_uri": eNet.FabricUri,
                                        "etag": eNet.ETAG,
                                        "scopesuri": eNet.ScopesUri,
                                })
		d.Set("members", members)
		d.SetId(name)
		return nil
	}


	if getType == "GetAll" {
		eNetList, err := config.ovClient.GetEthernetNetworks("", "", "", "")

		if err != nil {
			d.SetId("")
			return nil
		} else {
			d.Set("count_", eNetList.Count)
			d.Set("total", eNetList.Total)
			d.Set("list_uri", eNetList.URI)

			members := make([]map[string]interface{}, 0, len(eNetList.Members))
			for _, eNet := range eNetList.Members {
				members = append(members, map[string]interface{}{
					"name":   eNet.Name,
					"purpose": eNet.Purpose,
					"vlan_id": eNet.VlanId,
					"smart_link": eNet.SmartLink,
					"private_network": eNet.PrivateNetwork,
					"ethernet_network_type": eNet.EthernetNetworkType,
					"type": eNet.Type,
					"created": eNet.Created,
					"modified": eNet.Modified,
					"uri": eNet.URI,
					"connection_template_uri": eNet.ConnectionTemplateUri,
					"state":  eNet.State,
					"status": eNet.Status,
					"category": eNet.Category,
					"fabric_uri": eNet.FabricUri,
					"etag": eNet.ETAG,
					"scopesuri": eNet.ScopesUri,
				})
			}

			d.Set("members", members)
			d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
		}
	}
	return nil

}
