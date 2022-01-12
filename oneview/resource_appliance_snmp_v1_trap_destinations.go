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
	"math/rand"
	"strconv"
)

func resourceSNMPv1TrapDestination() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceSNMPv1TrapDestinationRead,
		Create: resourceSNMPv1TrapDestinationCreate,
		Update: resourceSNMPv1TrapDestinationUpdate,
		Delete: resourceSNMPv1TrapDestinationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"community_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"destination": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"destination_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func resourceSNMPv1TrapDestinationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	id := d.Get("destination_id").(string)

	// Generates random integer when id is null
	if id == "" {
		id = strconv.Itoa(rand.Intn(100))
	}
	options := ov.SNMPv1Trap{
		CommunityString: d.Get("community_string").(string),
		Destination:     d.Get("destination").(string),
		Port:            d.Get("port").(int),
	}

	err := config.ovClient.CreateSNMPv1TrapDestinations(options, id)
	if err != nil {
		d.SetId("")
		return err
	}

	d.SetId(id)
	return resourceSNMPv1TrapDestinationRead(d, meta)
}

func resourceSNMPv1TrapDestinationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	snmpTrap, err := config.ovClient.GetSNMPv1TrapDestinationsById(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}
	d.Set("community_string", snmpTrap.CommunityString)
	d.Set("destination", snmpTrap.Destination)
	d.Set("destination_id", d.Id())
	d.Set("port", snmpTrap.Port)
	d.Set("uri", snmpTrap.URI.String())
	return nil
}

func resourceSNMPv1TrapDestinationUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	updateOptions := ov.SNMPv1Trap{
		CommunityString: d.Get("community_string").(string),
		Port:            d.Get("port").(int),
		Destination:     d.Get("destination").(string),
	}

	_, err := config.ovClient.UpdateSNMPv1TrapDestinations(updateOptions, d.Id())
	if err != nil {
		return err
	}
	d.SetId(d.Id())

	return resourceSNMPv1TrapDestinationRead(d, meta)
}

func resourceSNMPv1TrapDestinationDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteSNMPv1TrapDestinations(d.Id())
	if err != nil {
		return err
	}
	return nil
}
