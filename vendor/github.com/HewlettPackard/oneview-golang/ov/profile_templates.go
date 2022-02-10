/*
(c) Copyright [2021] Hewlett Packard Enterprise Development LP

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package ov -
package ov

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/HewlettPackard/oneview-golang/liboneview"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/docker/machine/libmachine/log"
)

// introduced in v200 for oneview, allows for an easier method
// to build the profiles for servers and associate them.
// we don't operate on a new struct, we simply use the ServerProfile struct

// ProfileTemplatesNotSupported - determine these functions are supported
func (c *OVClient) ProfileTemplatesNotSupported() bool {
	var currentversion liboneview.Version
	var asc liboneview.APISupport
	currentversion = currentversion.CalculateVersion(c.APIVersion, 108) // force icsp to 108 version since icsp version doesn't matter
	asc = asc.NewByName("profile_templates.go")
	if !asc.IsSupported(currentversion) {
		log.Debugf("ProfileTemplates client version not supported: %+v", currentversion)
		return true
	}
	return false
}

// IsProfileTemplates - returns true when we should use GetProfileTemplate...
func (c *OVClient) IsProfileTemplates() bool {
	return !c.ProfileTemplatesNotSupported()
}

// GetProfileTemplateByName gets a server profile template by name
func (c *OVClient) GetProfileTemplateByName(name string) (ServerProfile, error) {
	var (
		profile ServerProfile
	)
	// v2 way to get ServerProfile
	if c.IsProfileTemplates() {
		profiles, err := c.GetProfileTemplates("", "", fmt.Sprintf("name matches '%s'", name), "name:asc", "")
		if profiles.Total > 0 {
			return profiles.Members[0], err
		} else {
			return profile, err
		}
	} else {

		// v1 way to get a ServerProfile
		profiles, err := c.GetProfiles("", "", fmt.Sprintf("name matches '%s'", name), "name:asc", "")
		if profiles.Total > 0 {
			return profiles.Members[0], err
		} else {
			return profile, err
		}
	}

}

// GetProfileTemplates gets a server profiles
func (c *OVClient) GetProfileTemplates(start string, count string, filter string, sort string, scopeUris string) (ServerProfileList, error) {
	var (
		uri      = "/rest/server-profile-templates"
		q        map[string]interface{}
		profiles ServerProfileList
	)
	q = make(map[string]interface{})
	if len(filter) > 0 {
		q["filter"] = filter
	}

	if sort != "" {
		q["sort"] = sort
	}

	if start != "" {
		q["start"] = start
	}

	if count != "" {
		q["count"] = count
	}

	if scopeUris != "" {
		q["scopeUris"] = scopeUris
	}

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	data, err := c.RestAPICall(rest.GET, uri, nil, q)
	if err != nil {
		return profiles, err
	}

	log.Debugf("GetProfileTemplates %s", data)
	if err := json.Unmarshal([]byte(data), &profiles); err != nil {
		return profiles, err
	}
	return profiles, nil
}

// IsZeroOfUnderlyingType returns true if a value is initialized.
func IsZeroOfUnderlyingType(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

// SetMp maps ManagementProcessors to IntManagementProcessor struct.
func SetMp(shtgen string, mp ManagementProcessors) IntManagementProcessor {
	mps := make([]MpSetting, 0)
	var emptyMpSettings MpSettings
	if !reflect.DeepEqual(mp.MpSetting, emptyMpSettings) {
		var emptyAdminAcc AdministratorAccount
		if mp.MpSetting.AdministratorAccount != emptyAdminAcc {
			args := make(map[string]interface{})
			v := reflect.ValueOf(mp.MpSetting.AdministratorAccount)
			typeOfS := v.Type()
			// iterate through Administrative Account Fields
			for i := 0; i < v.NumField(); i++ {
				// only adds fields that are initialized in order to bypass adding default/uninitialized values.
				if !IsZeroOfUnderlyingType(v.Field(i).Interface()) {
					args[strings.ToLower(string(typeOfS.Field(i).Name[0]))+typeOfS.Field(i).Name[1:]] = v.Field(i).Interface()
				}
			}
			mps = append(mps, MpSetting{
				SettingType: "AdministratorAccount",
				Args:        args,
			})
		}
		var emptyDirectory Directory
		if !reflect.DeepEqual(mp.MpSetting.Directory, emptyDirectory) {
			args := make(map[string]interface{})
			v := reflect.ValueOf(mp.MpSetting.Directory)
			typeOfS := v.Type()
			// iterate through Directory Fields.
			for i := 0; i < v.NumField(); i++ {
				// only adds fields that are initialized in order to bypass adding default/uninitialized values.
				if !IsZeroOfUnderlyingType(v.Field(i).Interface()) {
					args[strings.ToLower(string(typeOfS.Field(i).Name[0]))+typeOfS.Field(i).Name[1:]] = v.Field(i).Interface()
				}
			}
			mps = append(mps, MpSetting{
				SettingType: "Directory",
				Args:        args,
			})
		}

		var emptyHostname IloHostName
		if !reflect.DeepEqual(mp.MpSetting.IloHostName, emptyHostname) {
			args := make(map[string]interface{})
			v := reflect.ValueOf(mp.MpSetting.IloHostName)
			typeOfS := v.Type()
			// iterate through IloHostName Fields.
			for i := 0; i < v.NumField(); i++ {
				// only adds fields that are initialized in order to bypass adding default/uninitialized values.
				if !IsZeroOfUnderlyingType(v.Field(i).Interface()) {
					args[strings.ToLower(string(typeOfS.Field(i).Name[0]))+typeOfS.Field(i).Name[1:]] = v.Field(i).Interface()
				}
			}
			mps = append(mps, MpSetting{
				SettingType: "Hostname",
				Args:        args,
			})
		}

		var emptyKeyManager KeyManager
		if !reflect.DeepEqual(mp.MpSetting.KeyManager, emptyKeyManager) {
			args := make(map[string]interface{})
			v := reflect.ValueOf(mp.MpSetting.KeyManager)
			typeOfS := v.Type()
			// iterate through KeyManager Fields
			for i := 0; i < v.NumField(); i++ {
				// only adds fields that are initialized in order to bypass adding default/uninitialized values.
				if !IsZeroOfUnderlyingType(v.Field(i).Interface()) {
					args[strings.ToLower(string(typeOfS.Field(i).Name[0]))+typeOfS.Field(i).Name[1:]] = v.Field(i).Interface()
				}
			}
			mps = append(mps, MpSetting{
				SettingType: "KeyManager",
				Args:        args,
			})
		}

		if len(mp.MpSetting.DirectoryGroups) > 0 {
			var ags []interface{}
			// iterate through Directory Groups
			for _, i := range mp.MpSetting.DirectoryGroups {
				v := reflect.ValueOf(i)
				typeOfS := v.Type()
				arg := make(map[string]interface{})
				// iterate through Directory Group fields.
				for j := 0; j < v.NumField(); j++ {
					// only adds fields that are initialized in order to bypass adding default/uninitialized values.
					if !IsZeroOfUnderlyingType(v.Field(j).Interface()) {
						arg[strings.ToLower(string(typeOfS.Field(j).Name[0]))+typeOfS.Field(j).Name[1:]] = v.Field(j).Interface()
					}
				}
				ags = append(ags, arg)
			}
			args := map[string]interface{}{
				"directoryGroupAccounts": ags,
			}
			mps = append(mps, MpSetting{
				SettingType: "DirectoryGroups",
				Args:        args,
			})

		}

		if len(mp.MpSetting.LocalAccounts) > 0 {
			var ags []interface{}
			// iterate through localAccounts
			for _, i := range mp.MpSetting.LocalAccounts {
				v := reflect.ValueOf(i)
				typeOfS := v.Type()
				arg := make(map[string]interface{})
				// iterate through fields in local accounts
				for j := 0; j < v.NumField(); j++ {
					// only adds fields that are initialized in order to bypass adding default/uninitialized values.
					if !IsZeroOfUnderlyingType(v.Field(j).Interface()) {
						arg[strings.ToLower(string(typeOfS.Field(j).Name[0]))+typeOfS.Field(j).Name[1:]] = v.Field(j).Interface()
					}
				}
				//Check generation of server. Gen 9 and below does support some iLO attributes.
				if shtgen != "Gen10" && shtgen != "Gen11" {
					delete(arg, "loginPriv")
					delete(arg, "hostBIOSConfigPriv")
					delete(arg, "hostNICConfigPriv")
					delete(arg, "hostStorageConfigPriv")
				}
				ags = append(ags, arg)
			}

			args := map[string]interface{}{
				"localAccounts": ags,
			}

			mps = append(mps, MpSetting{
				SettingType: "LocalAccounts",
				Args:        args,
			})
		}
	}
	result := IntManagementProcessor{
		ComplianceControl: mp.ComplianceControl,
		ManageMp:          mp.ManageMp,
		MpSettings:        mps,
		ReapplyState:      mp.ReapplyState,
	}
	return result
}

func (c *OVClient) CreateProfileTemplate(serverProfileTemplate ServerProfile) error {
	log.Infof("Initializing creation of server profile template for %s.", serverProfileTemplate.Name)
	var (
		uri = "/rest/server-profile-templates"
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, serverProfileTemplate)
	log.Debugf("task -> %+v", t)
	serverHardwareType, err := c.GetServerHardwareTypeByUri(serverProfileTemplate.ServerHardwareTypeURI)
	if err != nil {
		log.Warnf("Error getting server hardware type %s", err)
	}
	serverHardwareTypeGen := serverHardwareType.Generation

	var emptyMgmtProcessorsStruct ManagementProcessors
	if !reflect.DeepEqual(serverProfileTemplate.ManagementProcessors, emptyMgmtProcessorsStruct) {
		mp := SetMp(serverHardwareTypeGen, serverProfileTemplate.ManagementProcessors)
		serverProfileTemplate.ManagementProcessor = mp
	}

	data, err := c.RestAPICall(rest.POST, uri, serverProfileTemplate)

	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new server profile template request: %s", err)
		return err
	}

	log.Debugf("Response New server profile template %s", data)
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return err
	}

	err = t.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (c *OVClient) DeleteProfileTemplate(name string) error {
	var (
		serverProfileTemplate ServerProfile
		err                   error
		t                     *Task
		uri                   string
	)

	serverProfileTemplate, err = c.GetProfileTemplateByName(name)
	if err != nil {
		return err
	}
	if serverProfileTemplate.Name != "" {
		t = t.NewProfileTask(c)
		t.ResetTask()
		log.Debugf("REST : %s \n %+v\n", serverProfileTemplate.URI, serverProfileTemplate)
		log.Debugf("task -> %+v", t)
		uri = serverProfileTemplate.URI.String()
		if uri == "" {
			log.Warn("Unable to post delete, no uri found.")
			t.TaskIsDone = true
			return err
		}
		data, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete server profile template request: %s", err)
			t.TaskIsDone = true
			return err
		}
		log.Debugf("Response delete profile_template %s", data)
		if err := json.Unmarshal([]byte(data), &t); err != nil {
			t.TaskIsDone = true
			log.Errorf("Error with task un-marshal: %s", err)
			return err
		}
		err = t.Wait()
		if err != nil {
			return err
		}

		return nil
	} else {
		log.Infof("ServerProfileTemplate could not be found to delete, %s, skipping delete ...", name)
	}
	return nil
}

func (c *OVClient) UpdateProfileTemplate(serverProfileTemplate ServerProfile) error {
	log.Infof("Initializing update of server profile template for %s.", serverProfileTemplate.Name)
	var (
		uri = serverProfileTemplate.URI.String()
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, serverProfileTemplate)
	log.Debugf("task -> %+v", t)

	serverHardwareType, err := c.GetServerHardwareTypeByUri(serverProfileTemplate.ServerHardwareTypeURI)
	if err != nil {
		log.Warnf("Error getting server hardware type %s", err)
	}
	serverHardwareTypeGen := serverHardwareType.Generation

	var emptyMgmtProcessorsStruct ManagementProcessors
	if !reflect.DeepEqual(serverProfileTemplate.ManagementProcessors, emptyMgmtProcessorsStruct) {
		mp := SetMp(serverHardwareTypeGen, serverProfileTemplate.ManagementProcessors)
		serverProfileTemplate.ManagementProcessor = mp
	}

	data, err := c.RestAPICall(rest.PUT, uri, serverProfileTemplate)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update server profile template request: %s", err)
		return err
	}

	log.Debugf("Response update ServerProfileTemplate %s", data)
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return err
	}

	err = t.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (c *OVClient) PatchServerProfileTemplate(p ServerProfile, request []Options) error {

	log.Infof("Initializing update of server profile for %s.", p.Name)

	var (
		uri = p.URI.String()
		t   *Task
	)

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, request)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PATCH, uri, request)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update server profile request: %s", err)
		return err
	}
	log.Debugf("Response update ServerProfile Template %s", data)
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return err
	}

	err = t.Wait()
	if err != nil {
		return err
	}

	return nil
}
