/*
(c) Copyright [2020] Hewlett Packard Enterprise Development LP

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

// Package ov
package ov

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
	"os"
)

// FirmwareOption structure for firware settings
type FirmwareOption struct {
	ConsistencyState         string        `json:"consistencyState,omitempty"` //Consistency state of the firmware component.
	FirmwareActivationType   string        `json:"firmwareActivationType,omitempty"`
	FirmwareBaselineUri      utils.Nstring `json:"firmwareBaselineUri,omitempty"`      // "firmwareBaselineUri": null,
	FirmwareInstallType      string        `json:"firmwareInstallType,omitempty"`      // Specifies the way a Service Pack for ProLiant (SPP) is installed. This field is used if the 'manageFirmware' field is true. Possible values are
	FirmwareScheduleDateTime string        `json:"firmwareScheduleDateTime,omitempty"` // Identifies the date and time the Service Pack for Proliant (SPP) will be activated.
	ForceInstallFirmware     bool          `json:"forceInstallFirmware"`               // "forceInstallFirmware": false,
	ManageFirmware           bool          `json:"manageFirmware"`                     // "manageFirmware": false
	ReapplyState             string        `json:"reapplyState,omitempty"`             //Current reapply state of the firmware component.
}

// BootModeOption mode option
type BootModeOption struct {
	ManageMode    bool          `json:"manageMode"`              // "manageMode": true,
	Mode          string        `json:"mode,omitempty"`          // "mode": "BIOS",
	PXEBootPolicy utils.Nstring `json:"pxeBootPolicy,omitempty"` // "pxeBootPolicy": null
	secureBoot    string        `json:"secureBoot,omitempty"`    // Enable or disable UEFI Secure Boot
}

// BootManagement management
type BootManagement struct {
	ManageBoot bool     `json:"manageBoot,omitempty"` // "manageBoot": true,
	Order      []string `json:"order,omitempty"`      // "order": ["CD","USB","HardDisk","PXE"]
}

// BiosSettings structure
type BiosSettings struct {
	ID    string `json:"id,omitempty"`    // id
	Value string `json:"value,omitempty"` // value
}

// BiosOption - bios options
type BiosOption struct {
	ConsistencyState   string         `json:"consistencyState,omitempty"`   //Consistency state of the BIOS component
	ManageBios         bool           `json:"manageBios"`                   // "manageBios": false,
	OverriddenSettings []BiosSettings `json:"overriddenSettings,omitempty"` // "overriddenSettings": []
	ReapplyState       string         `json:"reapplyState,omitempty"`       //Current reapply state of the BIOS component.
}

type ConnectionSettings struct {
	ComplianceControl string       `json:"complianceControl,omitempty"` // "complianceControl": "Checked",
	ManageConnections bool         `json:"manageConnections,omitempty"` // "manageConnections": false,
	Connections       []Connection `json:"connections,omitempty"`
	ReapplyState      string       `json:"reapplyState,omitempty"` //Current reapply state of the connection downlinks associated with the server profile
}

type Options struct {
	Op    string `json:"op,omitempty"`    // "op": "replace",
	Path  string `json:"path,omitempty"`  // "path": "/templateCompliance",
	Value string `json:"value,omitempty"` // "value": "Compliant",
}

type Servers struct {
	EnclosureGroupName     string   `json:"enclosureGroupName, omitempty"`
	EnclosureName          string   `json:"enclosureGroupName, omitempty"`
	EnclosureUri           string   `json:"enclosureGroupName, omitempty"`
	EnclosureBay           int      `json:"enclosureBay, omitempty"`
	ServerHardwareName     string   `json:"serverHardwareName, omitempty"`
	ServerHardwareUri      string   `json:"serverHardwareUri, omitempty"`
	ServerHardwareTypeName string   `json:"serverHardwareTypeName, omitempty"`
	ServerHardwareTypeUri  string   `json:"serverHardwareTypeUri, omitempty"`
	EnclosureGroupUri      string   `json:"enclosuregroupUri, omitempty"`
	PowerState             string   `json:"powerState, omitempty"`
	FormFactor             []string `json:"formFactor, omitempty"`
	ServerHardwareStatus   string   `json:"serverHardwareStatus, omitempty"`
}

type AvailableTarget struct {
	Type    string    `json:"type, omitempty"`
	Members []Servers `json:"targets, omitempty"`
}

type ManagementProcessor struct {
	ManageMp     bool         `json:"manageMp,omitempty"`
	MpSettings   []mpSettings `json:"mpSettings,omitempty"`
	ReapplyState string       `json:"reapplyState,omitempty"`
}
type mpSettings struct {
	Args        string `json:"args, omitempty"`
	SettingType string `json:"settingType, omitempty"`
}

// ServerProfile - server profile object for ov
type ServerProfile struct {
	Affinity                   string               `json:"affinity,omitempty"`         // "affinity": "Bay",
	AssociatedServer           utils.Nstring        `json:"associatedServer,omitempty"` // "associatedServer": null,
	Bios                       *BiosOption          `json:"bios,omitempty"`             // "bios": {	},
	Boot                       BootManagement       `json:"boot,omitempty"`             // "boot": { },
	BootMode                   BootModeOption       `json:"bootMode,omitempty"`         // "bootMode": {},
	Category                   string               `json:"category,omitempty"`         // "category": "server-profiles",
	ConnectionSettings         ConnectionSettings   `json:"connectionSettings,omitempty"`
	Created                    string               `json:"created,omitempty"`                    // "created": "20150831T154835.250Z",
	Description                string               `json:"description,omitempty"`                // "description": "Docker Machine Bay 16",
	ETAG                       string               `json:"eTag,omitempty"`                       // "eTag": "1441036118675/8"
	EnclosureBay               int                  `json:"enclosureBay,omitempty"`               // "enclosureBay": 16,
	EnclosureGroupURI          utils.Nstring        `json:"enclosureGroupUri,omitempty"`          // "enclosureGroupUri": "/rest/enclosure-groups/56ad0069-8362-42fd-b4e3-f5c5a69af039",
	EnclosureURI               utils.Nstring        `json:"enclosureUri,omitempty"`               // "enclosureUri": "/rest/enclosures/092SN51207RR",
	Firmware                   FirmwareOption       `json:"firmware,omitempty"`                   // "firmware": { },
	HideUnusedFlexNics         bool                 `json:"hideUnusedFlexNics"`                   // "hideUnusedFlexNics": false,
	InProgress                 bool                 `json:"inProgress,omitempty"`                 // "inProgress": false,
	InitialScopeUris           []utils.Nstring      `json:"initialScopeUris,omitempty"`           // "initialScopeUris":[],
	IscsiInitiatorName         string               `json:"iscsiInitiatorName,omitempty"`         //When iscsiInitatorNameType is set to UserDefined
	IscsiInitiatorNameType     string               `json:"iscsiInitiatorNameType,omitempty"`     //When set to UserDefined, the value of iscsiInitatorName is used as provided
	LocalStorage               LocalStorageOptions  `json:"localStorage,omitempty"`               // "localStorage": {},
	MACType                    string               `json:"macType,omitempty"`                    // "macType": "Physical",
	ManagementProcessor        *ManagementProcessor `json:"managementProcessor,omitempty"`        //
	Modified                   string               `json:"modified,omitempty"`                   // "modified": "20150902T175611.657Z",
	Name                       string               `json:"name,omitempty"`                       // "name": "Server_Profile_scs79",
	OSDeploymentSettings       OSDeploymentSettings `json:"osDeploymentSettings,omitempty"`       // "osDeploymentSettings": {...},
	ProfileUUID                utils.Nstring        `json:"profileUUID,omitempty"`                //The automatically generated 36-byte Universally Unique ID of the server profile.
	RefreshState               string               `json:"refreshState,omitempty"`               //Current refresh State of this Server Profile
	SanStorage                 SanStorageOptions    `json:"sanStorage,omitempty"`                 // "sanStorage": {},
	ScopesUri                  utils.Nstring        `json:"scopesUri,omitempty"`                  // "scopesUri": "/rest/scopes/resources/rest/server-profiles/DB7726F7-F601-4EA8-B4A6-D1EE1B32C07C",
	SerialNumber               utils.Nstring        `json:"serialNumber,omitempty"`               // "serialNumber": "2M25090RMW",
	SerialNumberType           string               `json:"serialNumberType,omitempty"`           // "serialNumberType": "Physical",
	ServerHardwareReapplyState string               `json:"serverHardwareReapplyState,omitempty"` //Current reapply state of the server that is associated with this server profile
	ServerHardwareTypeURI      utils.Nstring        `json:"serverHardwareTypeUri,omitempty"`      // "serverHardwareTypeUri": "/rest/server-hardware-types/DB7726F7-F601-4EA8-B4A6-D1EE1B32C07C",
	ServerHardwareURI          utils.Nstring        `json:"serverHardwareUri,omitempty"`          // "serverHardwareUri": "/rest/server-hardware/30373237-3132-4D32-3235-303930524D57",
	ServerProfileTemplateURI   utils.Nstring        `json:"serverProfileTemplateUri,omitempty"`   // undocmented option
	ServiceManager             string               `json:"serviceManager,omitempty"`             //Name of a service manager that is designated owner of the profile
	State                      string               `json:"state,omitempty"`                      // "state": "Normal",
	Status                     string               `json:"status,omitempty"`                     // "status": "Critical",
	TaskURI                    utils.Nstring        `json:"taskUri,omitempty"`                    // "taskUri": "/rest/tasks/6F0DF438-7D30-41A2-A36D-62AB866BC7E8",
	TemplateCompliance         string               `json:"templateCompliance,omitempty"`         // v2 Compliant, NonCompliant, Unknown
	Type                       string               `json:"type,omitempty"`                       // "type": "ServerProfileV4",
	URI                        utils.Nstring        `json:"uri,omitempty"`                        // "uri": "/rest/server-profiles/9979b3a4-646a-4c3e-bca6-80ca0b403a93",
	UUID                       utils.Nstring        `json:"uuid,omitempty"`                       // "uuid": "30373237-3132-4D32-3235-303930524D57",
	WWNType                    string               `json:"wwnType,omitempty"`                    // "wwnType": "Physical",
}

// GetConnectionByName gets the connection from a profile with a given name
func (s ServerProfile) GetConnectionByName(name string) (Connection, error) {
	var connection Connection
	for _, c := range s.ConnectionSettings.Connections {
		if c.Name == name {
			return c, nil
		}
	}
	return connection, errors.New("Error connection not found on server profile, please try a different connection name.")
}

// ServerProfileList a list of ServerProfile objects
// TODO: missing properties, need to think how we can make a higher lvl structure like an OVList
// Then things like Members are inherited
type ServerProfileList struct {
	Total       int             `json:"total,omitempty"`       // "total": 1,
	Count       int             `json:"count,omitempty"`       // "count": 1,
	Start       int             `json:"start,omitempty"`       // "start": 0,
	PrevPageURI utils.Nstring   `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring   `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         utils.Nstring   `json:"uri,omitempty"`         // "uri": "/rest/server-profiles?filter=serialNumber%20matches%20%272M25090RMW%27&sort=name:asc"
	Members     []ServerProfile `json:"members,omitempty"`     // "members":[]
}

// GetProfileByName gets a server profile by name
func (c *OVClient) GetProfileByName(name string) (ServerProfile, error) {
	var (
		profile ServerProfile
	)
	profiles, err := c.GetProfiles("", "", fmt.Sprintf("name matches '%s'", name), "name:asc", "")
	if profiles.Total > 0 {
		return profiles.Members[0], err
	} else {
		return profile, err
	}
}

// GetProfileBySN  accepts serial number
func (c *OVClient) GetProfileBySN(serialnum string) (ServerProfile, error) {
	var (
		profile ServerProfile
	)
	profiles, err := c.GetProfiles("", "", fmt.Sprintf("serialNumber matches '%s'", serialnum), "name:asc", "")
	if profiles.Total > 0 {
		return profiles.Members[0], err
	} else {
		return profile, err
	}
}

// GetProfiles - get a server profiles
func (c *OVClient) GetProfiles(start string, count string, filter string, sort string, scopeUris string) (ServerProfileList, error) {
	var (
		uri      = "/rest/server-profiles"
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
	// Setup query
	if len(q) > 0 {
		c.SetQueryString(q)
	}
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return profiles, err
	}

	log.Debugf("GetProfiles %s", data)
	if err := json.Unmarshal([]byte(data), &profiles); err != nil {
		return profiles, err
	}
	return profiles, nil
}

// GetProfileByURI - get the profile from a uri
func (c *OVClient) GetProfileByURI(uri utils.Nstring) (ServerProfile, error) {
	var (
		profile ServerProfile
	)

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri.String(), nil)
	if err != nil {
		return profile, err
	}

	log.Debugf("GetProfileByURI %s", data)
	if err := json.Unmarshal([]byte(data), &profile); err != nil {
		return profile, err
	}
	return profile, nil
}

// GetAvailableServers - To fetch available server hardwares
func (c *OVClient) GetAvailableServers(ServerHardwareUri string) (bool, error) {
	var (
		hardwareUri         = "/rest/server-profiles/available-targets"
		isHardwareAvailable = false
		profiles            AvailableTarget
	)

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	sh_data, err := c.RestAPICall(rest.GET, hardwareUri, nil)
	if err != nil {
		return isHardwareAvailable, err
	}

	if err := json.Unmarshal([]byte(sh_data), &profiles); err != nil {
		return isHardwareAvailable, err
	}

	for i := 0; i < len(profiles.Members); i++ {
		if profiles.Members[i].ServerHardwareUri == ServerHardwareUri {
			isHardwareAvailable = true
		}
	}
	return isHardwareAvailable, nil
}

// SubmitNewProfile - submit new profile template
func (c *OVClient) SubmitNewProfile(p ServerProfile) (err error) {
	log.Infof("Initializing creation of server profile for %s.", p.Name)
	var (
		uri    = "/rest/server-profiles"
		server ServerHardware
		t      *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, p)
	log.Debugf("task -> %+v", t)

	// Get available server hardwares to assign it to SP
	isHardwareAvailable, err := c.GetAvailableServers(p.ServerHardwareURI.String())
	if err != nil || isHardwareAvailable == false {
		log.Errorf("Error getting available Hardware: %s", p.ServerHardwareURI.String())
		os.Exit(1)
	}

	server, err = c.GetServerHardwareByUri(p.ServerHardwareURI)
	if err != nil {
		log.Warnf("Problem getting server hardware, %s", err)
	}

	// power off the server so that we can add to SP
	if server.Name != "" {
		server.PowerOff()
	}

	data, err := c.RestAPICall(rest.POST, uri, p)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new profile request: %s", err)
		return err
	}

	log.Debugf("Response New Profile %s", data)
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

// create profile from template
func (c *OVClient) CreateProfileFromTemplate(name string, template ServerProfile, blade ServerHardware) error {
	log.Debugf("TEMPLATE : %+v\n", template)
	var (
		new_template ServerProfile
		err          error
	)

	//GET on /rest/server-profile-templates/{id}new-profile
	log.Debugf("getting profile by URI %+v, v2", template.URI)
	new_template, err = c.GetProfileByURI(template.URI)
	if err != nil {
		return err
	}
	if c.APIVersion == 200 {
		new_template.Type = "ServerProfileV5"
	} else if c.APIVersion == 300 {
		new_template.Type = "ServerProfileV6"
	} else if c.APIVersion == 500 {
		new_template.Type = "ServerProfileV7"
	} else if c.APIVersion == 600 {
		new_template.Type = "ServerProfileV8"
	} else if c.APIVersion == 800 {
		new_template.Type = "ServerProfileV9"
	} else if c.APIVersion == 1000 {
		new_template.Type = "ServerProfileV10"
	} else if c.APIVersion == 1200 {
		new_template.Type = "ServerProfileV11"
	} else if c.APIVersion >= 1600 {
		new_template.Type = "ServerProfileV12"
	}
	new_template.ServerProfileTemplateURI = template.URI // create relationship
	new_template.ConnectionSettings = ConnectionSettings{
		Connections: template.ConnectionSettings.Connections,
	}
	log.Debugf("new_template -> %+v", new_template)
	new_template.ServerHardwareURI = blade.URI
	new_template.Description += " " + name
	new_template.Name = name
	log.Debugf("new_template -> %+v", new_template)

	err = c.SubmitNewProfile(new_template)
	return err
}

// submit new profile template
func (c *OVClient) SubmitDeleteProfile(p ServerProfile) (t *Task, err error) {
	var (
		uri = p.URI.String()
	)

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, p)
	log.Debugf("task -> %+v", t)
	if uri == "" {
		log.Warn("Unable to post delete, no uri found.")
		t.TaskIsDone = true
		return t, err
	}
	data, err := c.RestAPICall(rest.DELETE, uri, nil)
	if err != nil {
		log.Errorf("Error submitting new profile request: %s", err)
		t.TaskIsDone = true
		return t, err
	}

	log.Debugf("Response delete profile %s", data)
	if err := json.Unmarshal(data, &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return t, err
	}

	return t, err
}

// delete a profile, assign the server and remove the profile from the system
func (c *OVClient) DeleteProfile(name string) error {
	// get the profile for this server
	var (
		servernamemsg string
		server        ServerHardware
		profile       ServerProfile
		err           error
	)

	servernamemsg = "'no server'"
	profile, err = c.GetProfileByName(name)
	if err != nil {
		return err
	}

	if profile.Name != "" {
		if profile.ServerHardwareURI != "" {
			server, err = c.GetServerHardwareByUri(profile.ServerHardwareURI)
			if err != nil {
				log.Warnf("Problem getting server hardware, %s", err)
			} else {
				if server.Name != "" {
					servernamemsg = server.Name
				}
			}
		}
		log.Infof("Delete server profile %s from oneview, %s will be unassigned.", profile.Name, servernamemsg)

		// power off the server so that we can remove it
		if server.Name != "" {
			server.PowerOff()
		}

		// submit delete task
		t, err := c.SubmitDeleteProfile(profile)
		if c.APIVersion < 1000 {
			err = t.Wait()
			if err != nil {
				return err
			}
		}

		// check for task execution

	} else {
		log.Infof("Profile could not be found to delete, %s, skipping delete ...", name)
	}
	return nil
}

func (c *OVClient) UpdateServerProfile(p ServerProfile) error {
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
	log.Debugf("REST : %s \n %+v\n", uri, p)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, p)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update server profile request: %s", err)
		return err
	}

	log.Debugf("Response update ServerProfile %s", data)
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

func (c *OVClient) PatchServerProfile(p ServerProfile, request []Options) error {

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
	log.Debugf("Response update ServerProfile %s", data)
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
