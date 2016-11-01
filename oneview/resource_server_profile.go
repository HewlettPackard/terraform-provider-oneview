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
  "github.com/hashicorp/terraform/helper/schema"
  "github.com/HewlettPackard/oneview-golang/ov"
)

func resourceServerProfile() *schema.Resource {
  return &schema.Resource{
    Create: resourceServerProfileCreate,
    Read:   resourceServerProfileRead,
    Update: resourceServerProfileUpdate,
    Delete: resourceServerProfileDelete,

    Schema: map[string]*schema.Schema{
      "name": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      }, 
      "type": &schema.Schema{
        Type: schema.TypeString,
        Optional: true,
        Default: "ServerProfileV5",
      },
      "server_template": &schema.Schema{
        Type: schema.TypeString,
        Required: true,
      },
      "ilo_ip": &schema.Schema{
        Type: schema.TypeString,
        Computed: true,
      },
      "blade_bay": &schema.Schema{
        Type: schema.TypeString,
        Optional: true,
      },
      "os_deployment_plan": &schema.Schema { 
        Type: schema.TypeString,
        Optional: true,
      },
      "deployment_attribute": &schema.Schema{
        Optional: true,
        Type: schema.TypeList,
        Elem: &schema.Resource{
          Schema: map[string]*schema.Schema{
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
      "server_hardware_uri": &schema.Schema{
        Type: schema.TypeString,
        Computed: true,
      },
      "serial_number": &schema.Schema{
        Type: schema.TypeString,
        Computed: true,
      },
      "public_connection": &schema.Schema{
        Type: schema.TypeString,
        Optional: true,
      },
      "public_mac": &schema.Schema{
        Type: schema.TypeString,
        Computed: true,
      }, 
      "public_slot_id": &schema.Schema{
        Type: schema.TypeInt,
        Computed: true,
      },
    },
  }
}

func resourceServerProfileCreate(d *schema.ResourceData, meta interface{}) error {

    config := meta.(*Config)

    serverProfileTemplate, error := config.ovClient.GetProfileTemplateByName(d.Get("server_template").(string))
    if error != nil || serverProfileTemplate.URI.IsNil() {
      return fmt.Errorf("Could not find Server Profile Template\n%+v", d.Get("server_template").(string))
    }
    var serverHardware ov.ServerHardware
    if val, ok := d.GetOk("blade_bay"); ok {
      serverHardware, error = config.ovClient.GetServerHardwareByName(val.(string))
      if(error != nil){
        return error
      }
    } else {
      serverHardware, error = getServerHardware(config, serverProfileTemplate)
      if(error != nil){
        return error
      }
    }

    if val, ok := d.GetOk("os_deployment_plan"); ok {

      deploymentAttributes := make(map[string]string)
      osDeploymentPlan, error := config.ovClient.GetOSDeploymentPlanByName(val.(string))
      if error != nil || osDeploymentPlan.URI.IsNil() {
        return fmt.Errorf("Count not find osDeploymentPlan: %s", val.(string))
      }

      if _, ok := d.GetOk("deployment_attribute"); ok {
        deploymentAttributeCount := d.Get("deployment_attribute.#").(int)
        for i := 0; i < deploymentAttributeCount; i++ {
          deploymentAttributePrefix := fmt.Sprintf("deployment_attribute.%d", i)
          deploymentAttributes[d.Get(deploymentAttributePrefix + ".key").(string)] = d.Get(deploymentAttributePrefix + ".value").(string)
        }
      }

      SPerror := config.ovClient.CreateProfileFromTemplateWithI3S(d.Get("name").(string), serverProfileTemplate, serverHardware, osDeploymentPlan, deploymentAttributes)
      d.SetId(d.Get("name").(string))

      if SPerror != nil {
        d.SetId("")
        return SPerror
      }
    } else {
      SPerror := config.ovClient.CreateProfileFromTemplate(d.Get("name").(string), serverProfileTemplate, serverHardware)
      d.SetId(d.Get("name").(string))

      if SPerror != nil {
        d.SetId("")
        return SPerror
      }
    }

    return resourceServerProfileRead(d, meta)
}

func resourceServerProfileRead(d *schema.ResourceData, meta interface{}) error{
  
  config := meta.(*Config)
  
  serverProfile, err := config.ovClient.GetProfileByName(d.Id())
  if err != nil || serverProfile.URI.IsNil(){
    d.SetId("")
    return nil
  }

  serverHardware, err := config.ovClient.GetServerHardware(serverProfile.ServerHardwareURI)
  if(err != nil){
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

func resourceServerProfileUpdate(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceServerProfileDelete(d *schema.ResourceData, meta interface{}) error {
    config := meta.(*Config)
    
    error := config.ovClient.DeleteProfile(d.Get("name").(string))
    if error != nil {
      return error
    }
    return nil
}

func getServerHardware(config *Config, serverProfileTemplate ov.ServerProfile) (hw ov.ServerHardware, err error){

  var availableHardware ov.ServerHardware
  ovMutexKV.Lock(serverProfileTemplate.EnclosureGroupURI.String())
  defer ovMutexKV.Unlock(serverProfileTemplate.EnclosureGroupURI.String())

  for availableHardware.Created == "" {
    serverHardware, error := config.ovClient.GetAvailableHardware(serverProfileTemplate.ServerHardwareTypeURI, serverProfileTemplate.EnclosureGroupURI)
    if error != nil {
      return availableHardware, error
    }
    if ( ! serverHardwareURIs[serverHardware.URI.String()]){
      availableHardware = serverHardware
      serverHardwareURIs[serverHardware.URI.String()] = true
    }
  }
  return availableHardware, nil
}
