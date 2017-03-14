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
	"fmt"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceServerProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerProfileCreate,
		Read:   resourceServerProfileRead,
		Update: resourceServerProfileUpdate,
		Delete: resourceServerProfileDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ServerProfileV5",
			},
			"template": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ilo_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hardware_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hardware_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_connection": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_mac": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_slot_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceServerProfileCreate(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*Config)

	serverProfileTemplate, err := config.ovClient.GetProfileTemplateByName(d.Get("template").(string))
	if err != nil || serverProfileTemplate.URI.IsNil() {
		return fmt.Errorf("Could not find Server Profile Template\n%+v", d.Get("template").(string))
	}
	var serverHardware ov.ServerHardware
	if val, ok := d.GetOk("hardware_name"); ok {
		serverHardware, err = config.ovClient.GetServerHardwareByName(val.(string))
		if err != nil {
			return err
		}
	} else {
		serverHardware, err = getServerHardware(config, serverProfileTemplate)
		if err != nil {
			return err
		}
	}

	profileType := d.Get("type")
	if profileType == "ServerProfileV6" {
		err = config.ovClient.CreateProfileFromTemplateWithI3S(d.Get("name").(string), serverProfileTemplate, serverHardware)
		d.SetId(d.Get("name").(string))

		if err != nil {
			d.SetId("")
			return err
		}
	} else {
		err = config.ovClient.CreateProfileFromTemplate(d.Get("name").(string), serverProfileTemplate, serverHardware)
		d.SetId(d.Get("name").(string))

		if err != nil {
			d.SetId("")
			return err
		}
	}

	return resourceServerProfileRead(d, meta)
}

func resourceServerProfileRead(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*Config)

	serverProfile, err := config.ovClient.GetProfileByName(d.Id())
	if err != nil || serverProfile.URI.IsNil() {
		d.SetId("")
		return nil
	}

	serverHardware, err := config.ovClient.GetServerHardware(serverProfile.ServerHardwareURI)
	if err != nil {
		return err
	}

	d.Set("ilo_ip", serverHardware.GetIloIPAddress())
	d.Set("serial_number", serverProfile.SerialNumber.String())

	if val, ok := d.GetOk("public_connection"); ok {
		publicConnection, err := serverProfile.GetConnectionByName(val.(string))
		if err != nil {
			return err
		}
		d.Set("public_mac", publicConnection.MAC)
		d.Set("public_slot_id", publicConnection.ID)
	}

	return nil
}

func resourceServerProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serverProfile, err := config.ovClient.GetProfileByName(d.Id())
	if err != nil || serverProfile.URI.IsNil() {
		d.SetId("")
		return nil
	}

	return resourceServerProfileRead(d, meta)
}

func resourceServerProfileDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	err := config.ovClient.DeleteProfile(d.Get("name").(string))
	if err != nil {
		return err
	}
	return nil
}

func getServerHardware(config *Config, serverProfileTemplate ov.ServerProfile) (hw ov.ServerHardware, err error) {

	var availableHardware ov.ServerHardware
	ovMutexKV.Lock(serverProfileTemplate.EnclosureGroupURI.String())
	defer ovMutexKV.Unlock(serverProfileTemplate.EnclosureGroupURI.String())

	for availableHardware.Created == "" {
		serverHardware, err := config.ovClient.GetAvailableHardware(serverProfileTemplate.ServerHardwareTypeURI, serverProfileTemplate.EnclosureGroupURI)
		if err != nil {
			return availableHardware, err
		}
		if !serverHardwareURIs[serverHardware.URI.String()] {
			availableHardware = serverHardware
			serverHardwareURIs[serverHardware.URI.String()] = true
		}
	}
	return availableHardware, nil
}
