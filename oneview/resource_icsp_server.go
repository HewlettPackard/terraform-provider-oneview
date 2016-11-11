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
  "github.com/hashicorp/terraform/helper/schema"
  "fmt"
  "github.com/HewlettPackard/oneview-golang/icsp"
)

func resourceIcspServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceIcspServerCreate,
		Read: resourceIcspServerRead,
		Update: resourceIcspServerUpdate,
		Delete: resourceIcspServerDelete,

		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
			},
			"ilo_ip": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type: schema.TypeInt,
				Optional: true,
				Default: 443,
			},
			"user_name": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"mid": &schema.Schema{
				Type: schema.TypeString,
				Computed: true,
			},
			"serial_number": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"build_plans": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"public_mac": &schema.Schema{
				Type: schema.TypeString,
			 	Optional: true,
			},
			"public_slot_id": &schema.Schema{
				Type: schema.TypeInt,
				Optional: true,
			},
			"host_name": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
			},
			"custom_attribute": &schema.Schema{
				Optional: true,
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scope": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Default: "server",
						},
						"key": &schema.Schema{
							Type: schema.TypeString,
							Required: true,
						},
						"value": &schema.Schema{
							Type: schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"public_ipv4": &schema.Schema{
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIcspServerCreate(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*Config)

	
    csa := icsp.CustomServerAttributes{}
    initCsa := csa.New()

    customizeServer := icsp.CustomizeServer{
            SerialNumber: d.Get("serial_number").(string),
            ILoUser: d.Get("user_name").(string),
            IloIPAddress: d.Get("ilo_ip").(string),
            IloPort: d.Get("port").(int),
            IloPassword: d.Get("password").(string),
            ServerProperties: initCsa,
    }

    if _, ok := d.GetOk("build_plans"); ok {
            rawBuildPlans := d.Get("build_plans").([]interface{})
            buildPlans := make([]string, len(rawBuildPlans))
            for i, raw := range rawBuildPlans {
                    buildPlans[i] = raw.(string)
            }
            customizeServer.OSBuildPlans = buildPlans
    }

    if _, ok := d.GetOk("custom_attribute"); ok{
            customAttributeCount := d.Get("custom_attribute.#").(int)
            for i := 0; i < customAttributeCount; i++ {
                customAttributePrefix := fmt.Sprintf("custom_attribute.%d", i)
                initCsa.Set(d.Get(customAttributePrefix + ".key").(string), d.Get(customAttributePrefix + ".value").(string))
            }

    }

    if val, ok := d.GetOk("host_name"); ok {
            customizeServer.HostName = val.(string)
    }

    if val, ok := d.GetOk("public_slot_id"); ok {
            customizeServer.PublicSlotID = val.(int)
    }

    if val, ok := d.GetOk("public_mac"); ok {
            customizeServer.PublicMAC = val.(string)
    }

    err := config.icspClient.CustomizeServer(customizeServer)
    d.SetId(d.Get("ilo_ip").(string))
    if err != nil {
            return err
    }
       
	return resourceIcspServerRead(d, meta)
	/*
	config := meta.(*Config)
	

        csa := icsp.CustomServerAttributes{}

	customizeServer := icsp.CustomizeServer{
		SerialNumber: d.Get("serial_number").(string),
		ILoUser: d.Get("user_name").(string),
		IloIPAddress: d.Get("ilo_ip").(string),
		IloPort: d.Get("port").(int),
		IloPassword: d.Get("password").(string),
                ServerProperties: csa.New(),
	}

	if _, ok := d.GetOk("build_plans"); ok {
		rawBuildPlans := d.Get("build_plans").([]interface{})
		buildPlans := make([]string, len(rawBuildPlans))
		for i, raw := range rawBuildPlans {
			buildPlans[i] = raw.(string)
		}
		customizeServer.OSBuildPlans = buildPlans
		/*_, err = config.icspClient.ApplyDeploymentJobsWithUri(buildPlans, nil, server)
		if err != nil {
			return err
		}
	}

	if val, ok := d.GetOk("host_name"); ok {
		customizeServer.HostName = val.(string)
	}

	if val, ok := d.GetOk("public_slot_id"); ok {
		customizeServer.PublicSlotID = val.(int)
	}

	if val, ok := d.GetOk("public_mac"); ok {
		customizeServer.PublicMAC = val.(string)
	}

	err := config.icspClient.CustomizeServer(customizeServer)
	d.SetId(d.Get("ilo_ip").(string))
        if err != nil {
		return err
	}
	/*
	serverError := config.icspClient.CreateServer(d.Get("user_name").(string),d.Get("password").(string),d.Get("ilo_ip").(string),d.Get("port").(int))

        	
	if(serverError != nil){
		d.SetId("")
		return serverError
	}

	server, err := config.icspClient.GetServerBySerialNumber(d.Get("serial_number").(string))
	if err != nil {
		return err
	}
    
	if _, ok := d.GetOk("custom_attribute"); ok{

		customAttributeCount := d.Get("custom_attribute.#").(int)
		for i := 0; i < customAttributeCount; i++ {
			customAttributePrefix := fmt.Sprintf("custom_attribute.%d", i)
			
			values := make([]icsp.ValueItem, 0)
			values = append(values, icsp.ValueItem {
				Scope: d.Get(customAttributePrefix + ".scope").(string),
				Value: d.Get(customAttributePrefix + ".value").(string),
			})

			customAttribute := icsp.CustomAttribute{
				Key: d.Get(customAttributePrefix + ".key").(string),
				Values: values,
			}
			
			server.CustomAttributes = append(server.CustomAttributes, customAttribute)
		}
		
	}
	
	server, err = config.icspClient.SaveServer(server)
	
	if err != nil {
		return err
	}

	if _, ok := d.GetOk("build_plans"); ok {
		rawBuildPlans := d.Get("build_plans").([]interface{})
		buildPlans := make([]utils.Nstring, len(rawBuildPlans))
		for i, raw := range rawBuildPlans {
			buildPlans[i] = utils.NewNstring(raw.(string))
		}
		_, err = config.icspClient.ApplyDeploymentJobsWithUri(buildPlans, nil, server)
		if err != nil {
			return err
		}
	}

	d.SetId(d.Get("ilo_ip").(string))

    return resourceIcspServerRead(d, meta)*/
}

func resourceIcspServerRead(d *schema.ResourceData, meta interface{}) error {

    config := meta.(*Config)
  
	server, err := config.icspClient.GetServerByIP(d.Id())
	if err != nil {
		d.SetId("")
		return err
	}

	d.Set("ilo_ip", server.ILO.IPAddress)
	d.Set("mid", server.MID)
	d.Set("host_name", server.HostName)

	if publicMac, ok := d.GetOk("public_mac"); ok {
		for _, network := range server.Interfaces {
			if network.MACAddr == publicMac.(string) {
				d.Set("public_ipv4", network.IPV4Addr)
			}
		}
	}
    return nil
}

func resourceIcspServerUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

    csa := icsp.CustomServerAttributes{}
    initCsa := csa.New()

    customizeServer := icsp.CustomizeServer{
            SerialNumber: d.Get("serial_number").(string),
            ILoUser: d.Get("user_name").(string),
            IloIPAddress: d.Get("ilo_ip").(string),
            IloPort: d.Get("port").(int),
            IloPassword: d.Get("password").(string),
            ServerProperties: initCsa,
    }

    if _, ok := d.GetOk("build_plans"); ok {
            rawBuildPlans := d.Get("build_plans").([]interface{})
            buildPlans := make([]string, len(rawBuildPlans))
            for i, raw := range rawBuildPlans {
                    buildPlans[i] = raw.(string)
            }
            customizeServer.OSBuildPlans = buildPlans
    }

    if _, ok := d.GetOk("custom_attribute"); ok{
            customAttributeCount := d.Get("custom_attribute.#").(int)
            for i := 0; i < customAttributeCount; i++ {
                customAttributePrefix := fmt.Sprintf("custom_attribute.%d", i)
                initCsa.Set(d.Get(customAttributePrefix + ".key").(string), d.Get(customAttributePrefix + ".value").(string))
            }

    }

    if val, ok := d.GetOk("host_name"); ok {
            customizeServer.HostName = val.(string)
    }

    if val, ok := d.GetOk("public_slot_id"); ok {
            customizeServer.PublicSlotID = val.(int)
    }

    if val, ok := d.GetOk("public_mac"); ok {
            customizeServer.PublicMAC = val.(string)
    }

    err := config.icspClient.CustomizeServer(customizeServer)
    d.SetId(d.Get("ilo_ip").(string))
    if err != nil {
            return err
    }
       
	return resourceIcspServerRead(d, meta)
}

func resourceIcspServerDelete(d *schema.ResourceData, meta interface{}) error {
    config := meta.(*Config)
    
    isDel, error := config.icspClient.DeleteServer(d.Get("mid").(string))
    if error != nil {
      return error
    }
    if !isDel {
      return fmt.Errorf("Could not delete server")
    }
    return nil
}
