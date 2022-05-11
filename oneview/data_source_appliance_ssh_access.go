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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSSHAccess() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSSHAccessRead,

		Schema: map[string]*schema.Schema{
			"allow_ssh_access": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified": {
				Type:     schema.TypeString,
				Computed: true,
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

func dataSourceSSHAccessRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	sshAccess, err := config.ovClient.GetSshAccess()
	if err != nil {
		d.SetId("")
		return err
	} else if sshAccess.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.SetId(sshAccess.Type)
	d.Set("allow_ssh_access", sshAccess.AllowSshAccess)
	d.Set("category", sshAccess.Category)
	d.Set("created", sshAccess.Created)
	d.Set("etag", sshAccess.ETAG)
	d.Set("modified", sshAccess.Modified)
	d.Set("type", sshAccess.Type)
	d.Set("uri", sshAccess.URI.String())
	return nil
}
