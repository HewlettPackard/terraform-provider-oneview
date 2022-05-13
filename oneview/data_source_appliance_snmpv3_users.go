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

func dataSourceSNMPv3User() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSNMPv3UserRead,

		Schema: map[string]*schema.Schema{
			"authentication_passphrase": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"authentication_protocol": {
				Type:     schema.TypeString,
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
			"user_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"privacy_passphrase": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"privacy_protocol": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_level": {
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
			"user_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func dataSourceSNMPv3UserRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	username := d.Get("user_name").(string)
	snmpUser, err := config.ovClient.GetSNMPv3UserByUserName(username)
	if err != nil {
		d.SetId("")
		return err
	} else if snmpUser.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.Set("authentication_passphrase", snmpUser.AuthenticationPassphrase)
	d.Set("authentication_protocol", snmpUser.AuthenticationProtocol)
	d.Set("category", snmpUser.Category)
	d.Set("created", snmpUser.Created)
	d.Set("etag", snmpUser.ETAG)
	d.Set("user_id", snmpUser.Id)
	d.Set("modified", snmpUser.Modified)
	d.Set("privacy_passphrase", snmpUser.PrivacyPassphrase)
	d.Set("privacy_protocol", snmpUser.PrivacyProtocol)
	d.Set("security_level", snmpUser.SecurityLevel)
	d.Set("type", snmpUser.Type)
	d.Set("uri", snmpUser.URI.String())
	d.SetId(username)

	return nil
}
