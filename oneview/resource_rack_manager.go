// (C) Copyright 2022 Hewlett Packard Enterprise Development LP
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
	"errors"
	"log"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRackManager() *schema.Resource {
	return &schema.Resource{
		Create: resourceRackManagerCreate,
		Read:   resourceRackManagerRead,
		Update: resourceRackManagerUpdate,
		Delete: resourceRackManagerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
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
			"force": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"licensing_intent": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"model": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"part_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"refresh_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_support_uri": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scopes_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"initial_scope_uris": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"support_data_collection_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"support_data_collection_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"support_data_collections_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"support_state": {
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
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRackManagerCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	rm := ov.RackManager{
		Hostname: utils.Nstring(d.Get("hostname").(string)),
		UserName: d.Get("username").(string),
		Password: utils.Nstring(d.Get("password").(string)),
		Force:    d.Get("force").(bool),
	}

	if val, ok := d.GetOk("initial_scope_uris"); ok {
		rawInitialScopeUris := val.(*schema.Set).List()
		initialScopeUris := make([]utils.Nstring, len(rawInitialScopeUris))
		for i, raw := range rawInitialScopeUris {
			initialScopeUris[i] = utils.Nstring(raw.(string))
		}
		rm.InitialScopeUris = initialScopeUris
	}
	rmURI, rmtError := config.ovClient.AddRackManager(rm)

	if rmtError != nil {
		d.SetId("")
		return rmtError
	}
	d.Set("uri", rmURI)

	return resourceRackManagerRead(d, meta)
}

func resourceRackManagerRead(d *schema.ResourceData, meta interface{}) error {
	var (
		rm  ov.RackManager
		err error
	)
	config := meta.(*Config)
	if _, ok := d.GetOk("uri"); ok {

		rm, err = config.ovClient.GetRackManagerById(d.Get("id").(string))
	} else {

		// for importing by name
		rm, err = config.ovClient.GetRackManagerByName(d.Id())
	}
	if err != nil || rm.URI.IsNil() {
		d.SetId("")
		return nil
	}
	// setting ID as resource Id
	d.SetId(rm.Id)
	d.Set("category", rm.Category)
	d.Set("created", rm.Created)
	d.Set("etag", rm.ETAG)
	d.Set("hostname", d.Get("hostname").(string))
	d.Set("licensing_intent", rm.LicensingIntent)
	d.Set("location", rm.Location)
	d.Set("model", rm.Model)
	d.Set("modified", rm.Modified)
	d.Set("name", rm.Name)
	d.Set("part_number", rm.PartNumber)
	d.Set("refresh_state", rm.RefreshState)
	d.Set("remote_support_uri", rm.RemoteSupportUri)
	d.Set("password", d.Get("password").(string))
	d.Set("scopes_uri", rm.ScopesUri)
	d.Set("serial_number", rm.SerialNumber)
	d.Set("state", rm.State)
	d.Set("status", rm.Status)
	d.Set("support_data_collection_state", rm.SupportDataCollectionState)
	d.Set("support_data_collection_type", rm.SupportDataCollectionType)
	d.Set("support_data_collections_uri", rm.SupportDataCollectionsUri)
	d.Set("type", rm.Type)
	d.Set("uri", rm.URI.String())
	d.Set("username", d.Get("username").(string))

	// reads scopes from rack manager
	scopes, err := config.ovClient.GetScopeFromResource(rm.URI.String())
	if err != nil {
		log.Printf("unable to fetch scopes: %s", err)
	} else {
		d.Set("initial_scope_uris", scopes.ScopeUris)
	}

	return nil
}

func resourceRackManagerDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteRackManager(d.Get("name").(string))
	if err != nil {
		return err
	}
	return nil
}

func resourceRackManagerUpdate(d *schema.ResourceData, meta interface{}) error {
	return errors.New("update is not permitted for rack manager")

}
