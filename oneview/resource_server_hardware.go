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
	"errors"
	"fmt"
	"log"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServerHardware() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerHardwareCreate,
		Read:   resourceServerHardwareRead,
		Update: resourceServerHardwareUpdate,
		Delete: resourceServerHardwareDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"configuration_state": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"force": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"licensing_intent": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"initial_scope_uris": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"location_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"one_time_boot": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": {
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
				Optional:  true,
			},
			"power_state": {
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
			"server_group_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_hardware_type_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_power_state": {
				Optional: true,
				Computed: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"power_state": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"power_control": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"server_profile_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_serial_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mp_hosts_and_ranges": {
				Optional: true,
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"mp_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mp_firmware_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mp_dns_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uid_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func resourceServerHardwareCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	hardware := ov.ServerHardware{
		Hostname:           d.Get("hostname").(string),
		Username:           d.Get("username").(string),
		Password:           d.Get("password").(string),
		Force:              d.Get("force").(bool),
		LicensingIntent:    d.Get("licensing_intent").(string),
		ConfigurationState: d.Get("configuration_state").(string),
	}
	if val, ok := d.GetOk("initial_scope_uris"); ok {
		rawInitialScopeUris := val.(*schema.Set).List()
		initialScopeUris := make([]utils.Nstring, len(rawInitialScopeUris))
		for i, raw := range rawInitialScopeUris {
			initialScopeUris[i] = utils.Nstring(raw.(string))
		}
		hardware.InitialScopeUris = initialScopeUris
	}

	resourceURI, err := config.ovClient.AddRackServer(hardware)

	if err != nil && resourceURI != "" {
		d.SetId("")
		return err
	}

	d.Set("uri", resourceURI)
	return resourceServerHardwareRead(d, meta)
}

func resourceServerHardwareRead(d *schema.ResourceData, meta interface{}) error {
	var (
		servHard ov.ServerHardware
		err      error
	)
	config := meta.(*Config)

	// fetching server hardware hostname incase it's added
	if _, ok := d.GetOk("uri"); ok {
		servHard, err = config.ovClient.GetServerHardwareByUri(utils.Nstring(d.Get("uri").(string)))
	} else {
		// for refreshing imported server hardware we would need it's name
		if val, ok := d.GetOk("name"); ok {
			servHard, err = config.ovClient.GetServerHardwareByName(val.(string))
		} else {
			// for importing server hardware
			servHard, err = config.ovClient.GetServerHardwareByName(d.Id())
		}
	}

	if err != nil || servHard.URI.IsNil() {
		d.SetId("")
		return fmt.Errorf("unable to retrieve server hardware %s", err)
	}

	// setting UUID as resource Id
	d.SetId(servHard.UUID.String())
	d.Set("configuration_state", d.Get("configuration_state").(string))
	d.Set("hostname", d.Get("hostname").(string))
	//Force option is read from the configuration file
	if val, ok := d.GetOk("force"); ok {
		d.Set("force", val.(bool))
	}
	d.Set("licensing_intent", servHard.LicensingIntent)
	d.Set("maintenance_mode", servHard.MaintenanceMode)
	d.Set("name", servHard.Name)
	d.Set("one_time_boot", servHard.OneTimeBoot)
	d.Set("location_uri", servHard.LocationURI.String())
	d.Set("password", d.Get("password").(string))
	d.Set("power_state", servHard.PowerState)
	d.Set("type", servHard.Type)
	d.Set("uri", servHard.URI.String())
	d.Set("server_group_uri", servHard.ServerGroupURI.String())
	d.Set("server_hardware_type_uri", servHard.ServerHardwareTypeURI.String())
	d.Set("server_profile_uri", servHard.ServerProfileURI.String())
	d.Set("uuid", servHard.UUID.String())
	d.Set("virtual_serial_number", servHard.VirtualSerialNumber.String())
	d.Set("virtual_uuid", servHard.VirtualUUID)
	d.Set("mp_ip_address", servHard.MpIpAddress)
	d.Set("mp_firmware_version", servHard.MpFirwareVersion)
	d.Set("mp_dns_name", servHard.MpDnsName)
	d.Set("uid_state", servHard.UidState)
	d.Set("username", d.Get("username").(string))

	// reads server hardware scopes
	scopes, err := config.ovClient.GetScopeFromResource(servHard.URI.String())
	if err != nil {
		log.Printf("unable to fetch scopes: %s", err)
	} else {
		d.Set("initial_scope_uris", scopes.ScopeUris)
	}

	return nil
}

func resourceServerHardwareUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	if d.HasChange("one_time_boot") {
		err := config.ovClient.SetOneTimeBoot(d.Id(), d.Get("one_time_boot").(string))

		if err != nil {
			d.SetId("")
			return err
		}

	}
	if d.HasChange("maintenance_mode") {
		err := config.ovClient.SetMaintenanceMode(d.Id(), d.Get("maintenance_mode").(string))
		if err != nil {
			d.SetId("")
			return err
		}

	}
	if d.HasChange("uid_state") {
		err := config.ovClient.SetUidState(d.Id(), d.Get("uid_state").(string))
		if err != nil {
			d.SetId("")
			return err
		}
	}
	if d.HasChange("server_power_state") {
		powerMap := make(map[string]interface{})
		powerStates := d.Get("server_power_state").([]interface{})
		for _, powerState := range powerStates {
			powerMap = powerState.(map[string]interface{})
		}

		powerInput := map[string]interface{}{
			"powerState":   powerMap["power_state"],
			"powerControl": powerMap["power_control"],
		}

		err := config.ovClient.SetPowerState(d.Id(), powerInput)
		if err != nil {
			d.SetId("")
			return err
		}
	}
	if d.HasChange("username") || d.HasChange("password") || d.HasChange("configuration_state") || d.HasChange("initial_scope_uris") {
		return errors.New("Fields like username, password, configuration_state and initial_scope_uris cannot be changed")
	}
	d.SetId(d.Id())

	return resourceServerHardwareRead(d, meta)
}

func resourceServerHardwareDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	hardwareType, err := config.ovClient.GetServerHardwareTypeByUri(utils.Nstring(d.Get("server_hardware_type_uri").(string)))
	if err != nil {
		return err
	}
	if hardwareType.Platform == "RackServer" {
		err := config.ovClient.DeleteServerHardware(utils.Nstring(d.Get("uri").(string)))
		if err != nil {
			return err
		}
	} else {
		return errors.New("Deletion of Server hardware is only supported for Rack Servers")
	}
	return nil
}
