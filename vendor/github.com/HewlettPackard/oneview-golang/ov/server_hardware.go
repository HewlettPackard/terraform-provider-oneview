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
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

// HardwareState
type HardwareState int

const (
	H_UNKNOWN HardwareState = 1 + iota
	H_ADDING
	H_NOPROFILE_APPLIED
	H_MONITORED
	H_UNMANAGED
	H_REMOVING
	H_REMOVE_FAILED
	H_REMOVED
	H_APPLYING_PROFILE
	H_PROFILE_APPLIED
	H_REMOVING_PROFILE
	H_PROFILE_ERROR
	H_UNSUPPORTED
	H_UPATING_FIRMWARE
)

var hardwarestates = [...]string{
	"Unknown",          // not initialized
	"Adding",           // server being added
	"NoProfileApplied", // server successfully added
	"Monitored",        // server being monitored
	"Unmanaged",        // discovered a supported server
	"Removing",         // server being removed
	"RemoveFailed",     // unsuccessful server removal
	"Removed",          // server successfully removed
	"ApplyingProfile",  // profile being applied to server
	"ProfileApplied",   // profile successfully applied
	"RemovingProfile",  // profile being removed
	"ProfileError",     // unsuccessful profile apply or removal
	"Unsupported",      // server model or version not currently supported by the appliance
	"UpdatingFirmware", // server firmware update in progress
}

func (h HardwareState) String() string { return hardwarestates[h-1] }
func (h HardwareState) Equal(s string) bool {
	return (strings.ToUpper(s) == strings.ToUpper(h.String()))
}

// ServerHardware get server hardware from ov
type ServerHardware struct {
	ServerHardwarev200
	AssetTag              string          `json:"assetTag,omitempty"`           // "assetTag": "[Unknown]",
	Category              string          `json:"category,omitempty"`           // "category": "server-hardware",
	ConfigurationState    string          `json:"configurationState,omitempty"` // "configurationState": "Managed",
	Created               string          `json:"created,omitempty"`            // "created": "2015-08-14T21:02:01.537Z",
	Description           utils.Nstring   `json:"description,omitempty"`        // "description": null,
	ETAG                  string          `json:"eTag,omitempty"`               // "eTag": "1441147370086",
	Force                 bool            `json:"force,omitempty"`              // "force": false,
	FormFactor            string          `json:"formFactor,omitempty"`         // "formFactor": "HalfHeight",
	Generation            string          `json:"generation,omitempty"`         // "generation":,
	Hostname              string          `json:"hostname,omitempty"`           // "hostname": "17.1.1.1"
	HostOsType            int             `json:"hostOsType,omitempty"`
	InitialScopeUris      []utils.Nstring `json:"initialScopeUris,omitempty"`
	LicensingIntent       string          `json:"licensingIntent,omitempty"`    // "licensingIntent": "OneView",
	LocationURI           utils.Nstring   `json:"locationUri,omitempty"`        // "locationUri": "/rest/enclosures/092SN51207RR",
	MaintenanceMode       bool            `json:"maintenanceMode,omitempty"`    // "maintenanceMode": false,
	MemoryMb              int             `json:"memoryMb,omitempty"`           // "memoryMb": 262144,
	Model                 string          `json:"model,omitempty"`              // "model": "ProLiant BL460c Gen9",
	Modified              string          `json:"modified,omitempty"`           // "modified": "2015-09-01T22:42:50.086Z",
	MpHostsAndRanges      []utils.Nstring `json:"mpHostsAndRanges,omitempty"`   // "mpHostsAndRanges": ["172.1.1.1-172.1.1.10"],
	MpFirwareVersion      string          `json:"mpFirmwareVersion,omitempty"`  // "mpFirmwareVersion": "2.03 Nov 07 2014",
	MpModel               string          `json:"mpModel,omitempty"`            // "mpModel": "iLO4",
	Name                  string          `json:"name,omitempty"`               // "name": "se05, bay 16",
	OneTimeBoot           string          `json:"oneTimeBoot,omitempty"`        // "oneTimeBoot": "USB",
	PartNumber            string          `json:"partNumber,omitempty"`         // "partNumber": "727021-B21",
	Password              string          `json:"password,omitempty"`           // "password":,
	Position              int             `json:"position,omitempty"`           // "position": 16,
	PowerLock             bool            `json:"powerLock,omitempty"`          // "powerLock": false,
	PowerState            string          `json:"powerState,omitempty"`         // "powerState": "Off",
	ProcessorCoreCount    int             `json:"processorCoreCount,omitempty"` // "processorCoreCount": 14,
	ProcessorCount        int             `json:"processorCount,omitempty"`     // "processorCount": 2,
	ProcessorSpeedMhz     int             `json:"processorSpeedMhz,omitempty"`  // "processorSpeedMhz": 2300,
	ProcessorType         string          `json:"processorType,omitempty"`      // "processorType": "Intel(R) Xeon(R) CPU E5-2695 v3 @ 2.30GHz",
	RefreshState          string          `json:"refreshState,omitempty"`       // "refreshState": "NotRefreshing",
	RomVersion            string          `json:"romVersion,omitempty"`         // "romVersion": "I36 11/03/2014",
	ScopesUri             string          `json:"scopesUri,omitempty"`
	ScopeUris             []utils.Nstring `json:"scopeUris,omitempty"`
	SerialNumber          utils.Nstring   `json:"serialNumber,omitempty"`          // "serialNumber": "2M25090RMW",
	ServerGroupURI        utils.Nstring   `json:"serverGroupUri,omitempty"`        // "serverGroupUri": "/rest/enclosure-groups/56ad0069-8362-42fd-b4e3-f5c5a69af039",
	ServerHardwareTypeURI utils.Nstring   `json:"serverHardwareTypeUri,omitempty"` // "serverHardwareTypeUri": "/rest/server-hardware-types/DB7726F7-F601-4EA8-B4A6-D1EE1B32C07C",
	ServerName            string          `json:"servername,omitempty"`            // "serverName": "hostname.hpe.com",
	ServerProfileURI      utils.Nstring   `json:"serverProfileUri,omitempty"`      // "serverProfileUri": "/rest/server-profiles/9979b3a4-646a-4c3e-bca6-80ca0b403a93",
	ShortModel            string          `json:"shortModel,omitempty"`            // "shortModel": "BL460c Gen9",
	State                 string          `json:"state,omitempty"`                 // "state": "ProfileApplied",
	StateReason           string          `json:"stateReason,omitempty"`           // "stateReason": "NotApplicable",
	Status                string          `json:"status,omitempty"`                // "status": "Warning",
	Type                  string          `json:"type,omitempty"`                  // "type": "server-hardware-3",
	URI                   utils.Nstring   `json:"uri,omitempty"`                   // "uri": "/rest/server-hardware/30373237-3132-4D32-3235-303930524D57",
	Username              string          `json:"username,omitempty"`              // "username": "dcs",
	UUID                  utils.Nstring   `json:"uuid,omitempty"`                  // "uuid": "30373237-3132-4D32-3235-303930524D57",
	UidState              string          `json:"uidState,omitempty"`              // "uidState": "Off",
	VirtualSerialNumber   utils.Nstring   `json:"VirtualSerialNumber,omitempty"`   // "virtualSerialNumber": "",
	VirtualUUID           string          `json:"virtualUuid,omitempty"`           // "virtualUuid": "00000000-0000-0000-0000-000000000000"
	// v1 properties
	MpDnsName   string `json:"mpDnsName,omitempty"`   // "mpDnsName": "ILO2M25090RMW",
	MpIpAddress string `json:"mpIpAddress,omitempty"` // make this private to force calls to GetIloIPAddress() "mpIpAddress": "172.28.3.136",
	// extra client struct
	Client *OVClient `json:"-"`
}

// server hardware list, simillar to ServerProfileList with a TODO
type ServerHardwareList struct {
	Type        string           `json:"type,omitempty"`        // "type": "server-hardware-list-3",
	Category    string           `json:"category,omitempty"`    // "category": "server-hardware",
	Count       int              `json:"count,omitempty"`       // "count": 15,
	Created     string           `json:"created,omitempty"`     // "created": "2015-09-08T04:58:21.489Z",
	ETAG        string           `json:"eTag,omitempty"`        // "eTag": "1441688301489",
	Modified    string           `json:"modified,omitempty"`    // "modified": "2015-09-08T04:58:21.489Z",
	NextPageURI utils.Nstring    `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	PrevPageURI utils.Nstring    `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	Start       int              `json:"start,omitempty"`       // "start": 0,
	Total       int              `json:"total,omitempty"`       // "total": 15,
	URI         string           `json:"uri,omitempty"`         // "uri": "/rest/server-hardware?sort=name:asc&filter=serverHardwareTypeUri=%27/rest/server-hardware-types/DB7726F7-F601-4EA8-B4A6-D1EE1B32C07C%27&filter=serverGroupUri=%27/rest/enclosure-groups/56ad0069-8362-42fd-b4e3-f5c5a69af039%27&start=0&count=100"
	Members     []ServerHardware `json:"members,omitempty"`     // "members":[]
}

// ServerFirmware get server firmware from ov
type ServerFirmware struct {
	Category          string        `json:"category,omitempty"`          // "category": "server-hardware",
	Components        []Component   `json:"components,omitempty"`        // "components": [],
	Created           string        `json:"created,omitempty"`           // "created": "2019-02-11T16:01:30.321Z",
	ETAG              string        `json:"eTag,omitempty"`              // "eTag": null,
	Modified          string        `json:"modified,omitempty"`          // "modified": "2019-02-11T16:01:30.324Z",
	ServerHardwareURI utils.Nstring `json:"serverHardwareUri,omitempty"` // "serverHardwareUri": "/rest/server-hardware/31393736-3831-4753-4831-30305837524E",
	ServerModel       string        `json:"serverModel,omitempty"`       // "serverModel": "ProLiant BL660c Gen9",
	ServerName        string        `json:"serverName,omitempty"`        // "serverName": "Encl1, bay 1",
	State             string        `json:"state,omitempty"`             // "state": "Supported",
	Type              string        `json:"type,omitempty"`              // "type": "server-hardware-firmware-1",
	URI               utils.Nstring `json:"uri,omitempty"`               // "uri": "/rest/server-hardware/31393736-3831-4753-4831-30305837524E/firmware"
	Client            *OVClient     `json:"-"`
}

// server firmware list, simillar to ServerProfileList with a TODO
type ServerFirmwareList struct {
	Type        string           `json:"type,omitempty"`        // "type": "server-hardware-list-3",
	Category    string           `json:"category,omitempty"`    // "category": "server-hardware",
	Count       int              `json:"count,omitempty"`       // "count": 15,
	Created     string           `json:"created,omitempty"`     // "created": "2015-09-08T04:58:21.489Z",
	ETAG        string           `json:"eTag,omitempty"`        // "eTag": "1441688301489",
	Modified    string           `json:"modified,omitempty"`    // "modified": "2015-09-08T04:58:21.489Z",
	NextPageURI utils.Nstring    `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	PrevPageURI utils.Nstring    `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	Start       int              `json:"start,omitempty"`       // "start": 0,
	Total       int              `json:"total,omitempty"`       // "total": 15,
	URI         string           `json:"uri,omitempty"`         // "uri": "/rest/server-hardware/*/firmware?filter=serverModel='ProLiant DL380 Gen10'"
	Members     []ServerFirmware `json:"members,omitempty"`     // "members":[]
}

// Component Firmware Map from ServerFirmware
type Component struct {
	ComponentKey      string `json:"componentKey,omitempty"`      // "componentKey": "TBD",
	ComponentLocation string `json:"componentLocation,omitempty"` // "componentLocation": "System Board",
	ComponentName     string `json:"componentName,omitempty"`     // "componentName": "Power Management Controller Firmware",
	ComponentVersion  string `json:"componentVersion,omitempty"`  // "componentVersion": "1.0"
}

// Patch Operation attributes
type PatchData struct {
	Op    string `json:"op,omitempty"`    //"op": "replace",
	Path  string `json:"path,omitempty"`  //"path": "/uidState",
	Value string `json:"value,omitempty"` //"value": "On"
}

type PatchPowerData struct {
	Op    string                 `json:"op,omitempty"`    //"op": "replace",
	Path  string                 `json:"path,omitempty"`  //"path": "/uidState",
	Value map[string]interface{} `json:"value,omitempty"` //"value": "On"
}

// server hardware power off
func (s ServerHardware) PowerOff() error {
	var pt *PowerTask
	pt = pt.NewPowerTask(s)
	return pt.PowerExecutor(P_OFF)
}

// server hardware power on
func (s ServerHardware) PowerOn() error {
	var pt *PowerTask
	pt = pt.NewPowerTask(s)
	return pt.PowerExecutor(P_ON)
}

// GetPowerState gets the power state
func (s ServerHardware) GetPowerState() (PowerState, error) {
	var pt *PowerTask
	pt = pt.NewPowerTask(s)
	if err := pt.GetCurrentPowerState(); err != nil {
		return P_UKNOWN, err
	}
	return pt.State, nil
}

// Add single rack server to the appliance
func (c *OVClient) AddRackServer(rackServer ServerHardware) (utils.Nstring, error) {
	log.Infof("Adding rack server %s.", rackServer.Hostname)
	var (
		uri = "/rest/server-hardware"
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, rackServer)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, rackServer)

	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new rack server addition: %s", err)
		return t.AssociatedRes.ResourceURI, err
	}

	log.Debugf("Response New Rackserver %s", data)
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return t.AssociatedRes.ResourceURI, err
	}

	err = t.Wait()

	if err != nil {
		return t.AssociatedRes.ResourceURI, err
	}

	return t.AssociatedRes.ResourceURI, nil
}

// Add multiple rack servers
func (c *OVClient) AddMultipleRackServers(rackServer ServerHardware) error {
	log.Infof("Adding multiple rack servers %s.", rackServer.Hostname)
	var (
		uri = "/rest/server-hardware/discovery"
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	log.Debugf("REST : %s \n %+v\n", uri, rackServer)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, rackServer)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting multiple rack servers addition: %s", err)
		return err
	}

	log.Debugf("Response for multiple Rackservers %s", data)
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

// GetIloIPAddress - Use MpIpAddress for v1 and
// For v2 check MpHostInfo is not nil , loop through MpHostInfo.MpIPAddress[],
// and return the first nonzero address
func (h ServerHardware) GetIloIPAddress() string {
	if h.Client.IsHardwareSchemaV2() {
		if h.MpHostInfo != nil {
			log.Debug("working on getting IloIPAddress from MpHostInfo using v2")
			for _, MpIPObj := range h.MpHostInfo.MpIPAddresses {
				if len(MpIPObj.Address) > 0 &&
					(MpDHCP.Equal(MpIPObj.Type) ||
						MpStatic.Equal(MpIPObj.Type) ||
						MpUndefined.Equal(MpIPObj.Type)) {
					return MpIPObj.Address
				}
			}
		}
	} else {
		log.Debug("working on getting IloIPAddress from MpIpAddress")
		return h.MpIpAddress
	}
	return ""
}

// GetServerHardwareByUri gets a server hardware with uri
func (c *OVClient) GetServerHardwareByUri(uri utils.Nstring) (ServerHardware, error) {

	var hardware ServerHardware
	// refresh login

	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	// rest call
	data, err := c.RestAPICall(rest.GET, uri.String(), nil)
	if err != nil {
		return hardware, err
	}

	log.Debugf("GetServerHardware %s", data)
	if err := json.Unmarshal([]byte(data), &hardware); err != nil {
		return hardware, err
	}
	hardware.Client = c
	return hardware, nil
}

// GetServerHardwareByName gets a server hardware with uri
func (c *OVClient) GetServerHardwareByName(name string) (ServerHardware, error) {

	var (
		serverHardware ServerHardware
	)

	filters := []string{fmt.Sprintf("name matches '%s'", name)}
	serverHardwareList, err := c.GetServerHardwareList(filters, "name:asc", "", "", "")
	if serverHardwareList.Total > 0 {
		serverHardwareList.Members[0].Client = c
		return serverHardwareList.Members[0], err
	} else {
		return serverHardware, err
	}
}

// Set Query params for GetServerHardwareList and GetServerFirmwareList
func (c *OVClient) SetQueryParams(filters []string, sort string, start string, count string, expand string) map[string]interface{} {
	q := make(map[string]interface{})

	if len(filters) > 0 {
		q["filter"] = filters
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

	if expand != "" {
		q["expand"] = expand
	}
	return q
}

// GetServerHardwareList gets a server hardware with filters
func (c *OVClient) GetServerHardwareList(filters []string, sort string, start string, count string, expand string) (ServerHardwareList, error) {
	var (
		uri        = "/rest/server-hardware"
		q          map[string]interface{}
		serverlist ServerHardwareList
	)

	q = c.SetQueryParams(filters, sort, start, count, expand)

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	data, err := c.RestAPICall(rest.GET, uri, nil, q)
	if err != nil {
		return serverlist, err
	}

	log.Debugf("GetServerHardwareList %s", data)

	if err := json.Unmarshal([]byte(data), &serverlist); err != nil {
		return serverlist, err
	}

	if count == "" {
		total := strconv.Itoa(serverlist.Total)
		return c.GetServerHardwareList(filters, "", "", total, "")
	}

	for i := 0; i < serverlist.Count; i++ {
		serverlist.Members[i].Client = c
	}

	return serverlist, nil
}

// GetAvailableHardware gets available server
// blades = rest_api(:oneview, :get, "/rest/server-hardware?sort=name:asc&filter=serverHardwareTypeUri='#{server_hardware_type_uri}'&filter=serverGroupUri='#{enclosure_group_uri}'")
func (c *OVClient) GetAvailableHardware(hardwareTypeUri utils.Nstring, enclosureGroupUri utils.Nstring) (hw ServerHardware, err error) {
	var (
		hwlist ServerHardwareList
		f      = []string{"serverHardwareTypeUri='" + hardwareTypeUri.String() + "'",
			"serverGroupUri='" + enclosureGroupUri.String() + "'"}
	)
	if hwlist, err = c.GetServerHardwareList(f, "name:desc", "", "", ""); err != nil {
		return hw, err
	}
	if !(len(hwlist.Members) > 0) {
		return hw, errors.New("Error! No available blades that are compatible with the server profile!")
	}

	// pick an available blade
	for _, blade := range hwlist.Members {
		if H_NOPROFILE_APPLIED.Equal(blade.State) {
			hw = blade
			break
		}
	}
	if hw.Name == "" {
		return hw, errors.New("No more blades are available for provisioning!")
	}
	return hw, nil
}

// GetServerFirmwareByUri gets firmware for a server hardware with uri
func (c *OVClient) GetServerFirmwareByUri(uri utils.Nstring) (ServerFirmware, error) {

	var (
		firmware ServerFirmware
		main_uri = uri.String()
	)

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	// Get firmware
	main_uri = main_uri + "/firmware"

	// rest call
	data, err := c.RestAPICall(rest.GET, main_uri, nil)
	if err != nil {
		return firmware, err
	}

	log.Debugf("GetServerHardware %s", data)
	if err := json.Unmarshal([]byte(data), &firmware); err != nil {
		return firmware, err
	}
	return firmware, nil
}

// GetServerFirmwareByUri gets firmware for a server hardware with uri
func (c *OVClient) GetServerFirmwareList(filters []string, sort string, start string, count string) (ServerFirmwareList, error) {
	var (
		uri          = "/rest/server-hardware/*/firmware"
		q            map[string]interface{}
		firmwareList ServerFirmwareList
	)

	q = c.SetQueryParams(filters, sort, start, count, "")

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	data, err := c.RestAPICall(rest.GET, uri, nil, q)
	if err != nil {
		return firmwareList, err
	}

	log.Debugf("GetServerFirmwareList %s", data)

	if err := json.Unmarshal([]byte(data), &firmwareList); err != nil {
		return firmwareList, err
	}

	if count == "" {
		total := strconv.Itoa(firmwareList.Total)
		return c.GetServerFirmwareList(filters, "", "", total)
	}

	for i := 0; i < firmwareList.Count; i++ {
		firmwareList.Members[i].Client = c
	}

	return firmwareList, nil
}

// Refresh server hardware configuration
func (c *OVClient) RefreshServerHardware(id string, hardware ServerHardware) error {
	var (
		uri = "/rest/server-hardware/" + id + "/refreshState"
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, hardware)
	data, err := c.RestAPICall(rest.PUT, uri, hardware)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error while refreshing server hardware: %s", err)
		return err
	}

	log.Debugf("Response of server hardware refresh %s", data)
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

// Updates iLO Firmware Version to minimum firmware version
// supported by Oneview appliance
func (c *OVClient) UpdateiLOFirmwareVersion(id string) error {
	var (
		uri = "/rest/server-hardware/" + id + "/mpFirmwareVersion"
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, nil)
	data, err := c.RestAPICall(rest.PUT, uri, nil)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error while updating firmware version: %s", err)
		return err
	}

	log.Debugf("Response of firmware version update %s", data)
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

func (c *OVClient) SetOneTimeBoot(serverHardwareId string, value string) error {
	patchOperation := PatchData{
		Op:    "replace",
		Path:  "/oneTimeBoot",
		Value: value,
	}
	operation := []PatchData{patchOperation}
	fmt.Println("Update server Hardware's oneTimeBoot\n")
	err := c.Patch(serverHardwareId, operation)
	if err != nil {
		log.Errorf("Error while updating oneTimeBoot: %s", err)
		return err
	}

	return nil
}

func (c *OVClient) PatchPowerState(id string, operation []PatchPowerData) error {
	var (
		uri = "/rest/server-hardware/" + id
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, operation)
	data, err := c.RestAPICall(rest.PATCH, uri, operation)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error while doing Patch: %s", err)
		return err
	}

	log.Debugf("Response of Patch %s", data)
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

// Performs patch operation to update SH attributes
func (c *OVClient) Patch(id string, operation []PatchData) error {
	var (
		uri = "/rest/server-hardware/" + id
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, operation)
	data, err := c.RestAPICall(rest.PATCH, uri, operation)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error while doing Patch: %s", err)
		return err
	}

	log.Debugf("Response of Patch %s", data)
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

// Update the server into/out of maintenance mode
func (c *OVClient) SetMaintenanceMode(serverHardwareId string, value string) error {
	patchOperation := PatchData{
		Op:    "replace",
		Path:  "/maintenanceMode",
		Value: value,
	}
	operation := []PatchData{patchOperation}
	log.Debugf("Update server Hardware's maintenance mode\n")
	err := c.Patch(serverHardwareId, operation)
	if err != nil {
		log.Errorf("Error while updating maintenance mode: %s", err)
		return err
	}

	return nil
}

// Turn the server UID light On/Off
func (c *OVClient) SetUidState(serverHardwareId string, value string) error {
	patchOperation := PatchData{
		Op:    "replace",
		Path:  "/uidState",
		Value: value,
	}
	operation := []PatchData{patchOperation}
	log.Debugf("Update server Hardware's uidState\n")
	err := c.Patch(serverHardwareId, operation)
	if err != nil {
		log.Errorf("Error while updating uidState: %s", err)
		return err
	}

	return nil
}

// Update the power state of the server
func (c *OVClient) SetPowerState(serverHardwareId string, powerState map[string]interface{}) error {
	patchOperation := PatchPowerData{
		Op:    "replace",
		Path:  "/powerState",
		Value: powerState,
	}
	operation := []PatchPowerData{patchOperation}
	log.Debugf("Update server Hardware's power state\n")
	err := c.PatchPowerState(serverHardwareId, operation)
	if err != nil {
		log.Errorf("Error while updating power state: %s", err)
		return err
	}

	return nil
}

// Delete the rack server using URI
func (c *OVClient) DeleteServerHardware(uri utils.Nstring) error {
	var (
		hardware ServerHardware
		err      error
		t        *Task
	)

	hardware, err = c.GetServerHardwareByUri(uri)
	if err != nil {
		return err
	}

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Infof("URI:%s", hardware.URI)
	log.Debugf("REST : %s \n %+v\n", hardware.URI, hardware)

	if hardware.URI == "" {
		t.TaskIsDone = true
		log.Warn("Unable to post delete, no uri found.")
		return err
	}

	data, err := c.RestAPICall(rest.DELETE, hardware.URI.String(), nil)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting server hardware delete request: %s", err)
		return err
	}

	log.Debugf("Response delete server hardware %s", data)
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
