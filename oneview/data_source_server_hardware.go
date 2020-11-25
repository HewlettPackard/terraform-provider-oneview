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

func dataSourceServerHardware() *schema.Resource {
	return &schema.Resource{
		Read: resourceServerHardwareRead,

		Schema: map[string]*schema.Schema{
			"location_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
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
		},
	}
}

func resourceServerHardwareRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	servHard, err := config.ovClient.GetServerHardwareByName(d.Get("name").(string))
	if err != nil || servHard.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.SetId(d.Get("name").(string))
	d.Set("name", servHard.Name)
	d.Set("location_uri", servHard.LocationURI.String())
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

	return nil
}
