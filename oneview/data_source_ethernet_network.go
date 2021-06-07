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
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"io/ioutil"
)

func dataSourceEthernetNetwork() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceEthernetNetworkRead,
		Schema: ethernetSchema().Schema,
	}
}

func dataSourceEthernetNetworkRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	name := d.Get("name").(string)

	if len(name) > 0 {
		eNet, err := config.ovClient.GetEthernetNetworkByName(name)
		if err != nil {
			d.SetId("")
			return nil
		}
		d.Set("name", eNet.Name)
		d.Set("purpose", eNet.Purpose)
		d.Set("vlan_id", eNet.VlanId)
		d.Set("smart_link", eNet.SmartLink)
		d.Set("private_network", eNet.PrivateNetwork)
		d.Set("ethernet_network_type", eNet.EthernetNetworkType)
		d.Set("type", eNet.Type)
		d.Set("created", eNet.Created)
		d.Set("modified", eNet.Modified)
		d.Set("uri", eNet.URI)
		d.Set("connection_template_uri", eNet.ConnectionTemplateUri)
		d.Set("state", eNet.State)
		d.Set("status", eNet.Status)
		d.Set("category", eNet.Category)
		d.Set("fabric_uri", eNet.FabricUri)
		d.Set("etag", eNet.ETAG)
		d.Set("scopesuri", eNet.ScopesUri)
		d.SetId(name)
		return nil
	} else {
		eNetList, err := config.ovClient.GetEthernetNetworks("", "", "", "")

		if err != nil {
			d.SetId("")
			return nil
		} else {
			members := make([]map[string]interface{}, 0, len(eNetList.Members))
			for _, eNet := range eNetList.Members {
				members = append(members, map[string]interface{}{
					"name":                    eNet.Name,
					"purpose":                 eNet.Purpose,
					"vlan_id":                 eNet.VlanId,
					"smart_link":              eNet.SmartLink,
					"private_network":         eNet.PrivateNetwork,
					"ethernet_network_type":   eNet.EthernetNetworkType,
					"type":                    eNet.Type,
					"created":                 eNet.Created,
					"modified":                eNet.Modified,
					"uri":                     eNet.URI,
					"connection_template_uri": eNet.ConnectionTemplateUri,
					"state":                   eNet.State,
					"status":                  eNet.Status,
					"category":                eNet.Category,
					"fabric_uri":              eNet.FabricUri,
					"etag":                    eNet.ETAG,
					"scopesuri":               eNet.ScopesUri,
				})
			}
			file, _ := json.MarshalIndent(members, "", " ")
			_ = ioutil.WriteFile("test.json", file, 0644)

			d.SetId(string(file))
		}
		return nil
	}
}
