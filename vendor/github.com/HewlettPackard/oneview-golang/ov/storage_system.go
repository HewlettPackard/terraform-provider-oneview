/*
(c) Copyright [2018] Hewlett Packard Enterprise Development LP
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

package ov

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type StorageSystem struct {
	Hostname                              string                                 `json:"hostname,omitempty"`
	Username                              string                                 `json:"username,omitempty"`
	Password                              string                                 `json:"password,omitempty"`
	Credentials                           *Credentials                           `json:"credentials,omitempty"`
	Category                              string                                 `json:"category,omitempty"`
	ETAG                                  string                                 `json:"eTag,omitempty"`
	Name                                  string                                 `json:"name,omitempty"`
	Description                           utils.Nstring                          `json:"description,omitempty"`
	State                                 string                                 `json:"state,omitempty"`
	Status                                string                                 `json:"status,omitempty"`
	Type                                  string                                 `json:"type,omitempty"`
	URI                                   utils.Nstring                          `json:"uri,omitempty"`
	Family                                string                                 `json:"family,omitempty"`
	StoragePoolsUri                       utils.Nstring                          `json:"storagePoolsUri,omitempty"`
	TotalCapacity                         string                                 `json:"totalCapacity,omitempty"`
	Mode                                  string                                 `json:"mode,omitempty"`
	Ports                                 []Ports                                `json:"ports,omitempty"`
	StorageSystemDeviceSpecificAttributes *StorageSystemDeviceSpecificAttributes `json:"deviceSpecificAttributes,omitempty"`
}

type Credentials struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type Ports struct {
	PortDeviceSpecificAttributes PortDeviceSpecificAttributes `json:"deviceSpecificAttributes,omitempty"`
	Id                           string                       `json:"id,omitempty"`
	Mode                         string                       `json:"mode,omitempty"`
}

type PortDeviceSpecificAttributes struct {
	PartnerPort string `json:"partnerport,omitempty"`
}

type StorageSystemDeviceSpecificAttributes struct {
	Firmware                  string                `json:"firmware,omitempty"`
	Model                     string                `json:"model,omitempty"`
	ManagedPools              []ManagedPools        `json:"managedPools,omitempty"`
	ManagedDomain             string                `json:"managedDomain,omitempty"`
	DefaultEncryptionCipher   string                `json:"defaultEncryptionCipher,omitempty"`
	IsDefaultEncryptionForced bool                  `json:"isDefaultEncryptionForced,omitempty"`
	PerformancePolicies       []PerformancePolicies `json:"performancePolicies,omitempty"`
	ProtectionTemplates       []ProtectionTemplates `json:"protectionTemplates,omitempty"`
	SoftwareVersion           string                `json:"softwareVersion,omitempty"`
}

type PerformancePolicies struct {
	ApplicationCategory string `json:"applicationCategory,omitempty"`
	BlockSize           int    `json:",blockSize,omitempty"`
	Name                string `json:"name,omitempty"`
	IsCompressed        bool   `json:"isCompressed,omitempty"`
	SpacePolicy         string `json:"spacePolicy,omitempty"`
}

type ProtectionTemplates struct {
	AppSync string `json:"appSync,omitempty"`
	Name    string `json:"name,omitempty"`
}

type ManagedPools struct {
	Name          string `json:"name,omitempty"`
	Domain        string `json:"domain,omitempty"`
	DeviceType    string `json:"deviceType,omitempty"`
	FreeCapacity  string `json:"freeCapacity,omitempty"`
	RaidLevel     string `json:"raidLevel,omitempty"`
	Totalcapacity string `json:"totalCapacity,omitempty"`
}

type StorageSystemsList struct {
	Total       int             `json:"total,omitempty"`       // "total": 1,
	Count       int             `json:"count,omitempty"`       // "count": 1,
	Start       int             `json:"start,omitempty"`       // "start": 0,
	PrevPageURI utils.Nstring   `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring   `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         utils.Nstring   `json:"uri,omitempty"`         // "uri": "/rest/storage-systems"
	Members     []StorageSystem `json:"members,omitempty"`     // "members":[]
}

type ReachablePortsList struct {
	Category    string           `json:"category,omitempty"`
	Members     []ReachablePorts `json:"members,omitempty"`
	Total       int              `json:"total,omitempty"`
	Count       int              `json:"count,omitempty"`
	Start       int              `json:"start,omitempty"`
	PrevPageURI utils.Nstring    `json:"prevPageUri,omitempty"`
	NextPageURI utils.Nstring    `json:"nextPageUri,omitempty"`
	URI         utils.Nstring    `json:"uri,omitempty"`
}

type ReachablePorts struct {
	ReachableNetworks utils.Nstring `json:"reachableNetworks,omitempty"`
}

type VolumeSetList struct {
	Category    string        `json:"category,omitempty"`
	Members     []VolumeSet   `json:"members,omitempty"`
	Total       int           `json:"total,omitempty"`
	Count       int           `json:"count,omitempty"`
	Start       int           `json:"start,omitempty"`
	PrevPageURI utils.Nstring `json:"prevPageUri,omitempty"`
	NextPageURI utils.Nstring `json:"nextPageUri,omitempty"`
	URI         utils.Nstring `json:"uri,omitempty"`
}

type VolumeSet struct {
	TotalVolumes int      `json:"totalVolumes,omitempty"`
	VolumeURIs   []string `json:"volumeUris,omitempty"`
	State        string   `json:"state,omitempty"`
}

func (c *OVClient) GetStorageSystemByName(name string) (StorageSystem, error) {
	var (
		sSystem StorageSystem
	)
	sSystems, err := c.GetStorageSystems(fmt.Sprintf("name matches '%s'", name), "name:asc")
	if sSystems.Total > 0 {
		return sSystems.Members[0], err
	} else {
		return sSystem, err
	}
}

func (c *OVClient) GetStorageSystems(filter string, sort string) (StorageSystemsList, error) {
	var (
		uri     = "/rest/storage-systems"
		q       map[string]interface{}
		sSystem StorageSystemsList
	)
	q = make(map[string]interface{})
	if len(filter) > 0 {
		q["filter"] = filter
	}

	if sort != "" {
		q["sort"] = sort
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
		return sSystem, err
	}

	log.Debugf("GetStorageSystems %s", data)
	if err := json.Unmarshal([]byte(data), &sSystem); err != nil {
		return sSystem, err
	}
	return sSystem, nil
}

func (c *OVClient) CreateStorageSystem(sSystem StorageSystem) error {
	log.Infof("Initializing creation of storage system for %s.", sSystem.Name)
	var (
		uri = "/rest/storage-systems"
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, sSystem)
	log.Debugf("task -> %+v", t)

	data, err := c.RestAPICall(rest.POST, uri, sSystem)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new storage system request: %s", err)
		return err
	}

	log.Debugf("Response New StorageSystem %s", data)
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

func (c *OVClient) DeleteStorageSystem(name string) error {
	var (
		sSystem StorageSystem
		err     error
		t       *Task
		uri     string
	)

	sSystem, err = c.GetStorageSystemByName(name)
	if err != nil {
		return err
	}
	if sSystem.Name != "" {
		t = t.NewProfileTask(c)
		t.ResetTask()
		log.Debugf("REST : %s \n %+v\n", sSystem.URI, sSystem)
		log.Debugf("task -> %+v", t)
		uri = sSystem.URI.String()
		if uri == "" {
			log.Warn("Unable to post delete, no uri found.")
			t.TaskIsDone = true
			return err
		}
		data, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete storage volume request: %s", err)
			t.TaskIsDone = true
			return err
		}

		log.Debugf("Response delete storage volume %s", data)
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
		log.Infof("StorageSystem could not be found to delete, %s, skipping delete ...", name)
	}
	return nil
}

func (c *OVClient) UpdateStorageSystem(sSystem StorageSystem) error {
	log.Infof("Initializing update of storage volume for %s.", sSystem.Name)
	var (
		uri = sSystem.URI.String()
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, sSystem)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, sSystem)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update StorageSystem request: %s", err)
		return err
	}

	log.Debugf("Response update StorageSystem %s", data)
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

func (c *OVClient) GetReachablePorts(uri utils.Nstring) (ReachablePortsList, error) {
	var (
		reachable_ports ReachablePortsList
		main_uri        = uri.String()
	)
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	main_uri = main_uri + "/reachable-ports"
	data, err := c.RestAPICall(rest.GET, main_uri, nil)
	if err != nil {
		log.Errorf("Error in getting reachable ports: %s", err)
		return reachable_ports, err
	}
	log.Debugf("Reachable ports %s", data)
	if err := json.Unmarshal([]byte(data), &reachable_ports); err != nil {
		return reachable_ports, err
	}
	return reachable_ports, nil
}

func (c *OVClient) GetVolumeSets(uri utils.Nstring) (VolumeSetList, error) {
	var (
		volume_sets VolumeSetList
		main_uri    = uri.String()
	)
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	main_uri = main_uri + "/storage-volume-sets"
	data, err := c.RestAPICall(rest.GET, main_uri, nil)
	if err != nil {
		log.Errorf("Error in getting volume sets: %s", err)
		return volume_sets, err
	}
	log.Debugf("Volume Sets %s", data)
	if err := json.Unmarshal([]byte(data), &volume_sets); err != nil {
		return volume_sets, err
	}
	return volume_sets, nil
}
