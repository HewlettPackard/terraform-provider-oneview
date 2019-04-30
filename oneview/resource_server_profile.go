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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceServerProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerProfileCreate,
		Read:   resourceServerProfileRead,
		Update: resourceServerProfileUpdate,
		Delete: resourceServerProfileDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"template": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ServerProfileV9",
			},
			"hw_filter": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"hardware_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_connection": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ilo_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hardware_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial_number": {
				Type:     schema.TypeString,
				Computed: true,
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

	serverProfile := ov.ServerProfile{
		Type: d.Get("type").(string),
		Name: d.Get("name").(string),
	}

	if val, ok := d.GetOk("hardware_name"); ok {
		serverHardware, err := config.ovClient.GetServerHardwareByName(val.(string))
		if err != nil {
			return fmt.Errorf("could not find resource", err)
		}
		if serverHardware.PowerState != "off" {
			return fmt.Errorf("Server Hardware must be powered off to assign to the server profile")
		}
		serverProfile.ServerHardwareURI = serverHardware.URI
	}

	if val, ok := d.GetOk("template"); ok {
		serverProfileTemplate, err := config.ovClient.GetProfileTemplateByName(val.(string))
		if err != nil || serverProfileTemplate.URI.IsNil() {
			return fmt.Errorf("Could not find Server Profile Template\n%+v", val.(string))
		}
		serverProfile.ServerProfileTemplateURI = serverProfileTemplate.URI
	}

	err := config.ovClient.SubmitNewProfile(serverProfile)
	d.SetId(d.Get("name").(string))

	if err != nil {
		d.SetId("")
		return err
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

	serverHardware, err := config.ovClient.GetServerHardwareByUri(serverProfile.ServerHardwareURI)
	if err != nil {
		return err
	}

	d.Set("hardware_uri", serverHardware.URI.String())
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

	d.Set("name", serverProfile.Name)
	d.Set("type", serverProfile.Type)
	d.Set("uri", serverProfile.URI.String())

	return nil
}

func resourceServerProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	serverProfile := ov.ServerProfile{
		Type: d.Get("type").(string),
		Name: d.Get("name").(string),
		URI:  utils.NewNstring(d.Get("uri").(string)),
	}

	if val, ok := d.GetOk("hardware_name"); ok {
		serverHardware, err := config.ovClient.GetServerHardwareByName(val.(string))
		if err != nil {
			return fmt.Errorf("couldnt not find resource", err)
		}
		if serverHardware.PowerState != "off" {
			return fmt.Errorf("Server Hardware must be powered off to assign to server profile")
		}
		serverProfile.ServerHardwareURI = serverHardware.URI
	}

	if val, ok := d.GetOk("template"); ok {
		serverProfileTemplate, err := config.ovClient.GetProfileTemplateByName(val.(string))
		if err != nil || serverProfileTemplate.URI.IsNil() {
			return fmt.Errorf("Could not find Server Profile Template\n%+v", val.(string))
		}
		serverProfile.ServerProfileTemplateURI = serverProfileTemplate.URI
	}

	err := config.ovClient.UpdateServerProfile(serverProfile)
	if err != nil {
		return err
	}
	d.SetId(d.Get("name").(string))

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

func getServerHardware(config *Config, serverProfileTemplate ov.ServerProfile, filters []string) (hw ov.ServerHardware, err error) {
	ovMutexKV.Lock(serverProfileTemplate.EnclosureGroupURI.String())
	defer ovMutexKV.Unlock(serverProfileTemplate.EnclosureGroupURI.String())

	var (
		hwlist ov.ServerHardwareList
		f      = []string{"serverHardwareTypeUri='" + serverProfileTemplate.ServerHardwareTypeURI.String() + "'",
			"serverGroupUri='" + serverProfileTemplate.EnclosureGroupURI.String() + "'",
			"state='NoProfileApplied'"}
	)

	f = append(f, filters...)

	if hwlist, err = config.ovClient.GetServerHardwareList(f, "name:desc"); err != nil {
		if _, ok := err.(*json.SyntaxError); ok && len(filters) > 0 {
			return hw, fmt.Errorf("%s. It's likely your hw_filter(s) are incorrectly formatted", err)
		}
		return hw, err
	}
	for _, h := range hwlist.Members {
		if _, reserved := serverHardwareURIs[h.URI.String()]; !reserved {
			serverHardwareURIs[h.URI.String()] = true // Mark as reserved
			h.Client = config.ovClient                // The SDK GetServerHardwareList method doesn't set the
			// client, so we need to do it here. See https://github.com/HewlettPackard/oneview-golang/issues/103
			return h, nil
		}
	}

	return hw, errors.New("No blades that are compatible with the template are available!")
}
