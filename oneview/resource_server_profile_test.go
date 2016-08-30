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
/*
import (
  "fmt"
  "testing"
  
  "github.com/hashicorp/terraform/helper/resource"
  "github.com/hashicorp/terraform/terraform"
  "github.com/HewlettPackard/oneview-golang/ov" 
)

func TestAccServerProfile_1(t *testing.T){
  var serverProfile ov.ServerProfile 

  resource.Test(t, resource.TestCase{
    PreCheck: func() { testAccPreCheck(t) },
    Providers: testAccProviders,
    CheckDestroy: testAccCheckServerProfileDestroy,
    Steps: []resource.TestStep{
      resource.TestStep{
        Config: testAccServerProfile,
        Check: resource.ComposeTestCheckFunc(
	  testAccCheckServerProfileExists(
     "oneview_machine.test", &serverProfile),
          resource.TestCheckResourceAttr(
 	    "oneview_machine.test", "name", "test"),
        ),
      },
    },
  })
}

func testAccCheckServerProfileExists(n string, serverProfile *ov.ServerProfile) resource.TestCheckFunc {
  return func(s *terraform.State) error {
    rs, ok := s.RootModule().Resources[n]
    if !ok {
      return fmt.Errorf("Not found :%v", n)
    }
 
    if rs.Primary.ID == "" {
      return fmt.Errorf("No ID is set")
    }

    config, err := testProviderConfig()
    if err != nil {
      return err
    }

    testServerProfile, err := config.ovClient.GetProfileByName(rs.Primary.ID)
    if err != nil{
      return err
    }
    if testServerProfile.Name != rs.Primary.ID {
      return fmt.Errorf("Instance not found")
    }
    *serverProfile = testServerProfile
    return nil
  }
}

func testAccCheckServerProfileDestroy(s *terraform.State) error {
  config := testAccProvider.Meta().(*Config)
  for _, rs := range s.RootModule().Resources {
    if rs.Type != "oneview_machine" {
      continue
    }
    
    _, err := config.ovClient.GetProfileByName(rs.Primary.ID)

    if err == nil {
      return fmt.Errorf("Instance still exists")
    }
  }
 
  return nil
}

var testAccServerProfile = 
  `
  resource "oneview_machine" "test" {
    name             = "test"
    server_template  = "Web Server Template"
  }`

*/