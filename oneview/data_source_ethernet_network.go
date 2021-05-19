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
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/HewlettPackard/oneview-golang/utils"
	"encoding/json"
	"io/ioutil"
	"fmt"
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
						"description": {
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
					},
				},
			},

		},

	}
}

func dataSourceEthernetNetworkRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	eNetList, err := config.ovClient.GetEthernetNetworks("", "", "", "")

	if err != nil {
		d.SetId("")
		return nil
	} else {
		d.Set("count_", eNetList.Count)
		d.Set("total", eNetList.Total)
		d.Set("list_uri", eNetList.URI)

		elist := [] ov.EthernetNetwork{}
		for i := range eNetList.Members {
			ethernet := ov.EthernetNetwork{}
			ethernet.Name = eNetList.Members[i].Name
			ethernet.VlanId =  eNetList.Members[i].VlanId
			ethernet.Purpose = eNetList.Members[i].Purpose
			ethernet.SmartLink = eNetList.Members[i].SmartLink
			ethernet.PrivateNetwork = eNetList.Members[i].PrivateNetwork
			ethernet.EthernetNetworkType = eNetList.Members[i].EthernetNetworkType
			ethernet.Type = eNetList.Members[i].Type
			ethernet.Created = eNetList.Members[i].Created
			ethernet.Modified = eNetList.Members[i].Modified
			ethernet.URI = utils.Nstring(eNetList.Members[i].URI.String())
			ethernet.ConnectionTemplateUri = utils.Nstring(eNetList.Members[i].ConnectionTemplateUri.String())
			ethernet.State =  eNetList.Members[i].State
			ethernet.Status = eNetList.Members[i].Status
			ethernet.Category = eNetList.Members[i].Category
			ethernet.FabricUri = utils.Nstring(eNetList.Members[i].FabricUri.String())
			ethernet.ETAG = eNetList.Members[i].ETAG
			ethernet.ScopesUri = utils.Nstring(eNetList.Members[i].ScopesUri.String())

			initialScopeUris := make([]utils.Nstring, len(eNetList.Members[i].InitialScopeUris))
			for _, scope := range eNetList.Members[i].InitialScopeUris {
				initialScopeUri := utils.Nstring(scope.String())
				initialScopeUris = append(initialScopeUris, initialScopeUri)
			}

			ethernet.InitialScopeUris = initialScopeUris
			elist = append(elist, ethernet)

		}
		file, _ := json.MarshalIndent(elist, "", " ")
                file_name  := fmt.Sprintf("test_check_%d.json", 5)
                _ = ioutil.WriteFile(file_name, file, 0644)

		d.Set("members", elist)
		d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
		
	}
	return nil

}
