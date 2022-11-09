/*
(c) Copyright [2022] Hewlett Packard Enterprise Development LP

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
	"fmt"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

// RackManager object for OV
type RackManager struct {
	Category                   string          `json:"category,omitempty"`
	Created                    string          `json:"created,omitempty"` // "created": "20150831T154835.250Z",
	ETAG                       string          `json:"eTag,omitempty"`    // "eTag": "1441036118675/8"
	Force                      bool            `json:"force,omitempty"`
	Hostname                   utils.Nstring   `json:"hostname,omitempty"`
	Id                         string          `json:"id,omitempty"` // "43A48BDB-5FAC-42F5-9D1C-A280F48246AD"
	InitialScopeUris           []utils.Nstring `json:"initialScopeUris,omitempty"`
	LicensingIntent            string          `json:"licensingIntent,omitempty"`  //OneViewNoiLO"
	Location                   string          `json:"location,omitempty"`         // null
	Model                      string          `json:"model,omitempty"`            // null
	Modified                   string          `json:"modified,omitempty"`         // "modified": "20150902T175611.657Z",
	Name                       string          `json:"name,omitempty"`             //
	PartNumber                 string          `json:"partNumber,omitempty"`       // "affinity": "Bay",
	RefreshState               string          `json:"refreshState,omitempty"`     //Current refresh State of this Server Profile
	RemoteSupportUri           utils.Nstring   `json:"remoteSupportUri,omitempty"` // "
	Password                   utils.Nstring   `json:"password,omitempty"`
	ScopesUri                  utils.Nstring   `json:"scopesUri,omitempty"`                  // "scopesUri":
	SerialNumber               utils.Nstring   `json:"serialNumber,omitempty"`               // "serialNumber": "2M25090RMW",
	State                      string          `json:"state,omitempty"`                      // "state": "Normal",
	Status                     string          `json:"status,omitempty"`                     // "status": "Critical",
	SubResources               *SubResource    `json:"subResources,omitempty"`               // "subResources":[]
	SupportDataCollectionState string          `json:"supportDataCollectionState,omitempty"` //supportDataCollectionState
	SupportDataCollectionType  string          `json:"supportDataCollectionType,omitempty"`
	SupportDataCollectionsUri  string          `json:"supportDataCollectionsUri,omitempty"`
	SupportState               string          `json:"supportState,omitempty"`
	Type                       string          `json:"type,omitempty"` // "type": "ServerProfileV4",
	URI                        utils.Nstring   `json:"uri,omitempty"`  // "uri": "/rest/server-profiles/9979b3a4-
	UserName                   string          `json:"username,omitempty"`
}

type SubResource struct {
	Chassis               subresourceDat
	Partition             subresourceDat
	Managers              subresourceDat
	RemoteSupportSettings subresourceDat
	FwInventories         subresourceDat
}

type subresourceDat struct {
	Type     string        `json:"type,omitempty"`
	URI      string        `json:"uri,omitempty"`
	Modified string        `json:"modified,omitempty"`
	State    string        `json:"state,omitempty"`
	Etag     string        `json:"etag,omitempty"`
	Count    int           `json:"count,omitempty"`
	Data     []interface{} `json:"data,omitempty"`
}

type RackManagerList struct {
	Type        string        `json:"type,omitempty"`        // "type": "server-hardware-list-3",
	Category    string        `json:"category,omitempty"`    // "category": "server-hardware",
	Count       int           `json:"count,omitempty"`       // "count": 15,
	Created     string        `json:"created,omitempty"`     // "created": "2015-09-08T04:58:21.489Z",
	ETAG        string        `json:"eTag,omitempty"`        // "eTag": "1441688301489",
	Modified    string        `json:"modified,omitempty"`    // "modified": "2015-09-08T04:58:21.489Z",
	NextPageURI utils.Nstring `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	PrevPageURI utils.Nstring `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	Start       int           `json:"start,omitempty"`       // "start": 0,
	Total       int           `json:"total,omitempty"`       // "total": 15,
	URI         string        `json:"uri,omitempty"`         // "uri": "/rest/server-hardware/*/firmware?filter=serverModel='ProLiant DL380 Gen10'"
	Members     []RackManager `json:"members,omitempty"`     // "members":[]
}

// GetRack Manager By Name
func (c *OVClient) GetRackManagerByName(name string) (RackManager, error) {
	var (
		rm RackManager
	)
	rmList, err := c.GetRackManagerList("", "", fmt.Sprintf("name matches '%s'", name), "name:asc", "")
	if rmList.Total > 0 {
		return rmList.Members[0], err
	} else {
		return rm, err
	}
}

func (c *OVClient) GetRackManagerById(Id string) (RackManager, error) {
	var (
		rm  RackManager
		uri = "/rest/rack-managers/" + Id
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return rm, err
	}
	log.Debugf("GetRackManager %s", data)
	if err := json.Unmarshal([]byte(data), &rm); err != nil {
		return rm, err
	}
	return rm, nil
}

// GetRackMansgers - get all the rack managers
func (c *OVClient) GetRackManagerList(start string, count string, filter string, sort string, scopeUris string) (RackManagerList, error) {
	var (
		uri = "/rest/rack-managers"
		q   map[string]interface{}
		rms RackManagerList
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
		return rms, err
	}

	log.Debugf("Get Rack Managers %s", data)
	if err := json.Unmarshal([]byte(data), &rms); err != nil {
		return rms, err
	}
	return rms, nil
}

// Add rack manager- Add  new rack manager
func (c *OVClient) AddRackManager(rm RackManager) (resourceId string, err error) {

	log.Debugf("Initializing adding of RackManager %s.", rm.Hostname)
	var (
		uri   = "/rest/rack-managers"
		rmUri = ""
		t     *Task
	)

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	// log.Infof("REST : %s \n %+v\n", uri, rm)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, rm)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new add RackManager request: %s", err)
		return rmUri, err
	}

	log.Debugf("Response New RackrManager %s", data)
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return rmUri, err
	}

	err = t.Wait()
	if err != nil {
		return rmUri, err
	}
	rmUri = string(t.AssociatedRes.ResourceURI)

	return rmUri, nil
}

// delete a profile, assign the server and remove the profile from the system
func (c *OVClient) DeleteRackManager(name string) error {
	var (
		rm  RackManager
		err error
		t   *Task
		uri string
	)

	rm, err = c.GetRackManagerByName(name)
	if err != nil {
		return err
	}
	if rm.Name != "" {
		t = t.NewProfileTask(c)
		t.ResetTask()
		log.Debugf("REST : %s \n %+v\n", rm.URI, rm)
		log.Debugf("task -> %+v", t)
		uri = rm.URI.String()
		if uri == "" {
			log.Warn("Unable to post delete, no uri found.")
			t.TaskIsDone = true
			return err
		}
		data, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete rack manager request: %s", err)
			t.TaskIsDone = true
			return err
		}

		log.Debugf("Response delete rack manager %s", data)
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
		log.Infof("Rack Manager could not be found to delete, %s, skipping delete ...", name)
	}
	return nil
}
