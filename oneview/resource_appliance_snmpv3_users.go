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
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSNMPv3User() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSNMPv3UserRead,
		Create: resourceSNMPv3UserCreate,
		Update: resourceSNMPv3UserUpdate,
		Delete: resourceSNMPv3UserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"authentication_passphrase": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"authentication_protocol": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"category": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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
				Optional: true,
			},
			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"privacy_passphrase": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"privacy_protocol": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"security_level": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSNMPv3UserCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	createOptions := ov.SNMPv3User{
		AuthenticationProtocol:   d.Get("authentication_protocol").(string),
		AuthenticationPassphrase: d.Get("authentication_passphrase").(string),
		Category:                 d.Get("category").(string),
		Id:                       d.Get("user_id").(string),
		PrivacyPassphrase:        d.Get("privacy_passphrase").(string),
		PrivacyProtocol:          d.Get("privacy_protocol").(string),
		SecurityLevel:            d.Get("security_level").(string),
		Type:                     d.Get("type").(string),
		UserName:                 d.Get("user_name").(string),
	}

	_, err := config.ovClient.CreateSNMPv3Users(createOptions)
	if err != nil {
		d.SetId("")
		return err
	}

	d.SetId(d.Get("user_name").(string))
	return resourceSNMPv3UserRead(d, meta)
}

func resourceSNMPv3UserRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	id := d.Id()
	snmpUser, err := config.ovClient.GetSNMPv3UserByUserName(id)
	if err != nil || snmpUser.URI.IsNil() {
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
	d.SetId(id)
	return nil
}

func resourceSNMPv3UserUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	updateOptions := ov.SNMPv3User{
		AuthenticationProtocol:   d.Get("authentication_protocol").(string),
		AuthenticationPassphrase: d.Get("authentication_passphrase").(string),
		Category:                 d.Get("category").(string),
		ETAG:                     d.Get("etag").(string),
		Id:                       d.Get("user_id").(string),
		SecurityLevel:            d.Get("security_level").(string),
		PrivacyProtocol:          d.Get("privacy_protocol").(string),
		PrivacyPassphrase:        d.Get("privacy_passphrase").(string),
		Type:                     d.Get("type").(string),
		UserName:                 d.Get("user_name").(string),
		URI:                      utils.NewNstring(d.Get("uri").(string)),
	}

	_, err := config.ovClient.UpdateSNMPv3User(updateOptions, d.Get("user_id").(string))
	if err != nil {
		return err
	}
	d.SetId(d.Get("user_name").(string))

	return resourceSNMPv3UserRead(d, meta)
}

func resourceSNMPv3UserDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteSNMPv3UserByName(d.Id())
	if err != nil {
		return err
	}
	return nil
}
