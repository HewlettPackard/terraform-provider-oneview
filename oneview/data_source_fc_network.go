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

func dataSourceFCNetwork() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFCNetworkRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fabric_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"link_stability_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"auto_login_redistribution": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"connection_template_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"managed_san_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
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
			"etag": {
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
			"scopesuri": {
				Computed: true,
				Type:     schema.TypeString,
			},
		},
	}
}

func dataSourceFCNetworkRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	id := d.Get("name").(string)

	fcNet, err := config.ovClient.GetFCNetworkByName(id)
	if err != nil || fcNet.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.SetId(id)
	d.Set("name", fcNet.Name)
	d.Set("fabric_type", fcNet.FabricType)
	d.Set("link_stability_time", fcNet.LinkStabilityTime)
	d.Set("auto_login_redistribution", fcNet.AutoLoginRedistribution)
	d.Set("description", fcNet.Description.String())
	d.Set("type", fcNet.Type)
	d.Set("uri", fcNet.URI.String())
	d.Set("connection_template_uri", fcNet.ConnectionTemplateUri.String())
	d.Set("managed_san_uri", fcNet.ManagedSanURI.String())
	d.Set("status", fcNet.Status)
	d.Set("category", fcNet.Category)
	d.Set("state", fcNet.State)
	d.Set("fabric_uri", fcNet.FabricUri.String())
	d.Set("created", fcNet.Created)
	d.Set("modified", fcNet.Modified)
	d.Set("etag", fcNet.ETAG)
	d.Set("scopesuri", fcNet.ScopesUri.String())
	return nil
}
