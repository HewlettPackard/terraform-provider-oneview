package ov

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type HypervisorManager struct {
	Category             string          `json:"category,omitempty"`             // "category": "hypervisor-managers",
	AvailableDvsVersions []utils.Nstring `json:"availableDvsVersions,omitempty"` // "availableDvsVersions": "",
	Created              string          `json:"created,omitempty"`              // "created": "20150831T154835.250Z",
	Description          utils.Nstring   `json:"description,omitempty"`          // "description": "Hypervisor Manager 1",
	DisplayName          string          `json:"displayName,omitempty"`          // "displayName": "HypervisorManager1",
	ETAG                 string          `json:"eTag,omitempty"`                 // "eTag": "1441036118675/8",
	HypervisorType       string          `json:"hypervisorType,omitempty"`       // "hypervisorType": "HyperV","Vmware",
	Modified             string          `json:"modified,omitempty"`             // "modified": "20150831T154835.250Z",
	Name                 string          `json:"name,omitempty"`                 // "name": "hostname or IP",
	Password             string          `json:"password,omitempty"`             // "password": "",
	Port                 int             `json:"port,omitempty"`                 // "port": 443,
	Preferences          *Preference     `json:"preferences"`                    // "preferences": HypervisorClusterSettings,
	RefreshState         string          `json:"refreshState,omitempty"`         // "refreshState": "NotRfreshing",
	ResourcePaths        []ResourcePath  `json:"resourcePaths,omitempty"`        //"resourcePaths":""
	ScopesUri            utils.Nstring   `json:"scopesUri,omitempty"`            // "scopesUri":
	State                string          `json:"state,omitempty"`                // "state": "Connected",
	StateReason          string          `json:"stateReason,omitempty"`          // "state": "",
	Status               string          `json:"status,omitempty"`               // "status": "Critical",
	Type                 string          `json:"type,omitempty"`                 // "type": "HypervisorManagerV2",
	URI                  utils.Nstring   `json:"uri,omitempty"`                  // "uri": "/rest/hypervisor-managers/e2f0031b-52bd-4223-9ac1-d91cb519d548"
	UUID                 utils.Nstring   `json:"uuid,omitempty"`                 // "UUID":"60FB5CB3-FF04-400A-BEC8-E7920CB4193"
	Username             string          `json:"username,omitempty"`             // "username": "name1",
	Version              string          `json:"version,omitempty"`              // "version": ""
	InitialScopeUris     []utils.Nstring `json:"initialScopeUris,omitempty"`     // "initialScopUris":
}

type ResourcePath struct {
	UserPath   string `json:"userPath,omitempty"`   // "userPath":"DC1"
	ActualPath string `json:"actualPath,omitempty"` // "actualPath":"DC1/host"
}

type Preference struct {
	Type                     string `json:"type"`                               //"type":"Vmware"
	VirtualSwitchType        string `json:"virtualSwitchType"`                  // "virtualSwitchType":"Standard"
	DistributedSwitchVersion string `json:"distributedSwitchVersion,omitempty"` //"distributedSwitchVersion":null
	DistributedSwitchUsage   string `json:"distributedSwitchUsage,omitempty"`   //"distributedSwitchUsage":null
	MultiNicVMotion          bool   `json:"multiNicVMotion"`                    //"multiNicVMotion":false
	DrsEnabled               bool   `json:"drsEnabled"`                         //"drsEnabled":true
	HaEnabled                bool   `json:"haEnabled"`                          //"haEnabled":false

}
type HypervisorManagerList struct {
	Total       int                 `json:"total,omitempty"`       // "total": 1,
	Count       int                 `json:"count,omitempty"`       // "count": 1,
	Start       int                 `json:"start,omitempty"`       // "start": 0,
	PrevPageURI utils.Nstring       `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring       `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         utils.Nstring       `json:"uri,omitempty"`         // "uri": "/rest/hypervisor-managers?filter=connectionTemplateUri%20matches%7769cae0-b680-435b-9b87-9b864c81657fsort=name:asc"
	Members     []HypervisorManager `json:"members,omitempty"`     // "members":[]
}

func (c *OVClient) GetHypervisorManagerByName(name string) (HypervisorManager, error) {
	var (
		hypM HypervisorManager
	)
	hypMs, err := c.GetHypervisorManagers("", "", fmt.Sprintf("name matches '%s'", name), "name:asc")
	if hypMs.Total > 0 {
		return hypMs.Members[0], err
	} else {
		return hypM, err
	}
}

func (c *OVClient) GetHypervisorManagers(start string, count string, filter string, sort string) (HypervisorManagerList, error) {
	var (
		uri                = "/rest/hypervisor-managers"
		q                  map[string]interface{}
		hypervisorManagers HypervisorManagerList
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

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	// Setup query
	if len(q) > 0 {
		c.SetQueryString(q)
	}

	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return hypervisorManagers, err
	}

	log.Debugf("GetHypervisorManagers %s", data)
	if err := json.Unmarshal([]byte(data), &hypervisorManagers); err != nil {
		return hypervisorManagers, err
	}
	return hypervisorManagers, nil
}

func (c *OVClient) CreateHypervisorManager(hypM HypervisorManager) error {
	log.Infof("Initializing adding of HypervisorManager %s.", hypM.Name)
	var (
		uri = "/rest/hypervisor-managers"
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Infof("REST : %s \n %+v\n", uri, hypM)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, hypM)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new add HypervisorManager request: %s", err)
		return err
	}

	log.Debugf("Response New HypervisorManager %s", data)
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

func (c *OVClient) DeleteHypervisorManager(name string) error {
	var (
		hypM HypervisorManager
		err  error
		t    *Task
		uri  string
	)

	hypM, err = c.GetHypervisorManagerByName(name)
	if err != nil {
		return err
	}
	if hypM.Name != "" {
		t = t.NewProfileTask(c)
		t.ResetTask()
		log.Debugf("REST : %s \n %+v\n", hypM.URI, hypM)
		log.Debugf("task -> %+v", t)
		uri = hypM.URI.String()
		if uri == "" {
			log.Warn("Unable to post delete, no uri found.")
			t.TaskIsDone = true
			return err
		}
		data, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete hypervisor manager request: %s", err)
			t.TaskIsDone = true
			return err
		}

		log.Debugf("Response delete hypervisor manager %s", data)
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
		log.Infof("Hypervisor Manager could not be found to delete, %s, skipping delete ...", name)
	}
	return nil
}

func (c *OVClient) UpdateHypervisorManager(hypM HypervisorManager) error {
	log.Infof("Initializing update of hypervisor manager for %s.", hypM.Name)
	var (
		uri = hypM.URI.String()
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, hypM)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, hypM)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update hypervisor manager request: %s", err)
		return err
	}

	log.Debugf("Response update EthernetNetwork %s", data)
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
