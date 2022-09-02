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

func dataSourceRelativeValue() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRelativeValueRead,

		Schema: map[string]*schema.Schema{
			"port_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"interconnect_type_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port_num": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceRelativeValueRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	port_name := d.Get("port_name").(string)
	interrconnect_type_name := d.Get("interconnect_type_name").(string)
	interrconnect_type, _ := config.ovClient.GetInterconnectTypeByName(interrconnect_type_name)
	port_num, err := config.ovClient.GetRelativeValue(port_name, interrconnect_type.URI)

	if err != nil {
		d.SetId("")
		return err
	}

	d.SetId(port_name)

	d.Set("port_name", port_name)
	d.Set("port_num", port_num)
	d.Set("interrconnect_type", interrconnect_type)

	return nil
}
