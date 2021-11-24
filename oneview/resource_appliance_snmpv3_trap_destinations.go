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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSNMPv3TrapDestination() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSNMPv3TrapDestinationRead,
		Create: resourceSNMPv3TrapDestinationCreate,
		Update: resourceSNMPv3TrapDestinationUpdate,
		Delete: resourceSNMPv3TrapDestinationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"trap_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"engine_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"destination_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id_field": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"category": {
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

func resourceSNMPv3TrapDestinationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	createOptions := ov.SNMPv3Trap{
		UserID:             d.Get("user_id").(string),
		Port:               d.Get("port").(int),
		DestinationAddress: d.Get("destination_address").(string),
	}

	response, err := config.ovClient.CreateSNMPv3TrapDestinations(createOptions)
	if err != nil {
		d.SetId("")
		return err
	}

	d.SetId(response.ID)
	return resourceSNMPv3TrapDestinationRead(d, meta)
}

func resourceSNMPv3TrapDestinationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	snmpTrap, err := config.ovClient.GetSNMPv3TrapDestinationsById(d.Id())
	if err != nil || snmpTrap.URI.IsNil() {
		d.SetId("")
		return nil
	}

	d.Set("type", snmpTrap.Type)
	d.Set("created", snmpTrap.Created)
	d.Set("modified", snmpTrap.Modified)
	d.Set("uri", snmpTrap.URI.String())
	d.Set("destination_address", snmpTrap.DestinationAddress)
	d.Set("category", snmpTrap.Category)
	d.Set("user_uri", snmpTrap.UserURI)
	d.Set("etag", snmpTrap.ETAG)
	d.Set("user_id", snmpTrap.UserID)
	d.Set("id_field", snmpTrap.ID)
	d.Set("port", snmpTrap.Port)
	d.Set("engine_id", snmpTrap.EngineID)
	d.Set("trap_type", snmpTrap.TrapType)
	return nil
}

func resourceSNMPv3TrapDestinationUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	updateOptions := ov.SNMPv3Trap{
		ID:                 d.Id(),
		UserID:             d.Get("user_id").(string),
		Port:               d.Get("port").(int),
		DestinationAddress: d.Get("destination_address").(string),
	}

	trap, err := config.ovClient.UpdateSNMPv3TrapDestinations(updateOptions)
	if err != nil {
		return err
	}
	d.SetId(trap.ID)

	return resourceSNMPv3TrapDestinationRead(d, meta)
}

func resourceSNMPv3TrapDestinationDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteSNMPv3TrapDestinations(d.Get("id").(string))
	if err != nil {
		return err
	}
	return nil
}
