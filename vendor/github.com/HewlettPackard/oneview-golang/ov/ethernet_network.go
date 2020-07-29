package ov

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type EthernetNetwork struct {
	Category              string          `json:"category,omitempty"`              // "category": "ethernet-networks",
	ConnectionTemplateUri utils.Nstring   `json:"connectionTemplateUri,omitempty"` // "connectionTemplateUri": "/rest/connection-templates/7769cae0-b680-435b-9b87-9b864c81657f",
	Created               string          `json:"created,omitempty"`               // "created": "20150831T154835.250Z",
	Description           utils.Nstring   `json:"description,omitempty"`           // "description": "Ethernet network 1",
	ETAG                  string          `json:"eTag,omitempty"`                  // "eTag": "1441036118675/8",
	EthernetNetworkType   string          `json:"ethernetNetworkType,omitempty"`   // "ethernetNetworkType": "Tagged",
	FabricUri             utils.Nstring   `json:"fabricUri,omitempty"`             // "fabricUri": "/rest/fabrics/9b8f7ec0-52b3-475e-84f4-c4eac51c2c20",
	Modified              string          `json:"modified,omitempty"`              // "modified": "20150831T154835.250Z",
	Name                  string          `json:"name,omitempty"`                  // "name": "Ethernet Network 1",
	PrivateNetwork        bool            `json:"privateNetwork"`                  // "privateNetwork": false,
	Purpose               string          `json:"purpose,omitempty"`               // "purpose": "General",
	SmartLink             bool            `json:"smartLink"`                       // "smartLink": false,
	State                 string          `json:"state,omitempty"`                 // "state": "Normal",
	Status                string          `json:"status,omitempty"`                // "status": "Critical",
	SubnetUri             utils.Nstring   `json:"subnetUri,omitempty"`             // "subnetUri": "",
	Type                  string          `json:"type,omitempty"`                  // "type": "ethernet-networkV3",
	URI                   utils.Nstring   `json:"uri,omitempty"`                   // "uri": "/rest/ethernet-networks/e2f0031b-52bd-4223-9ac1-d91cb519d548"
	VlanId                int             `json:"vlanId,omitempty"`                // "vlanId": 1,
	ScopesUri             utils.Nstring   `json:"scopesUri,omitempty"`             // "scopesUri":
	InitialScopeUris      []utils.Nstring `json:"initialScopeUris,omitempty"`      // "initialScopUris":
}

type EthernetNetworkList struct {
	Total       int               `json:"total,omitempty"`       // "total": 1,
	Count       int               `json:"count,omitempty"`       // "count": 1,
	Start       int               `json:"start,omitempty"`       // "start": 0,
	PrevPageURI utils.Nstring     `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring     `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         utils.Nstring     `json:"uri,omitempty"`         // "uri": "/rest/ethernet-networks?filter=connectionTemplateUri%20matches%7769cae0-b680-435b-9b87-9b864c81657fsort=name:asc"
	Members     []EthernetNetwork `json:"members,omitempty"`     // "members":[]
}

type Bandwidth struct {
	MaximumBandwidth int `json:"maximumBandwidth"` //"maximumBandwidth":10000
	TypicalBandwidth int `json:"typicalBandwidth"` //"typicalBandwidth":2000
}

type BulkEthernetNetwork struct {
	VlanIdRange    string    `json:"vlanIdRange"`    // "vlanIdRange":"1-500",
	Purpose        string    `json:"purpose"`        // "purpose":"General",
	NamePrefix     string    `json:"namePrefix"`     // "namePrefix":"TestNetwork",
	SmartLink      bool      `json:"smartLink"`      // "smartLink":false,
	PrivateNetwork bool      `json:"privateNetwork"` // "privateNetwork":false,
	Bandwidth      Bandwidth `json:"bandwidth"`      // "bandwidth":10000,2000
	Type           string    `json:"type"`           // "type":"bulk-ethernet-network",
}

type BulkDelete struct {
	NetworkUris []utils.Nstring `json:"networkUris",omitempty` // "networkUris": [ "/rest/ethernet-networks/e2f0031b-52bd-4223-9ac1-d91cb519d548", "/rest/ethernet-networks/f2f0031b-52bd-4223-9ac1-d91cb519d549"]
}

func (c *OVClient) GetEthernetNetworkByName(name string) (EthernetNetwork, error) {
	var (
		eNet EthernetNetwork
	)
	eNets, err := c.GetEthernetNetworks("", "", fmt.Sprintf("name matches '%s'", name), "name:asc")
	if eNets.Total > 0 {
		return eNets.Members[0], err
	} else {
		return eNet, err
	}
}

func (c *OVClient) GetEthernetNetworks(start string, count string, filter string, sort string) (EthernetNetworkList, error) {
	var (
		uri              = "/rest/ethernet-networks"
		q                map[string]interface{}
		ethernetNetworks EthernetNetworkList
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
		return ethernetNetworks, err
	}

	log.Debugf("GetEthernetNetworks %s", data)
	if err := json.Unmarshal([]byte(data), &ethernetNetworks); err != nil {
		return ethernetNetworks, err
	}
	return ethernetNetworks, nil
}

func (c *OVClient) GetAssociatedProfile(id string) ([]string, error) {
	var (
		uri            = "/rest/ethernet-networks/"
		serverProfiles = new([]string)
	)
	uri = uri + id + "/associatedProfiles"
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return *serverProfiles, err
	}
	log.Infof("GetAssociatedProfile %s", data)
	if err := json.Unmarshal([]byte(data), serverProfiles); err != nil {
		return *serverProfiles, err
	}
	return *serverProfiles, nil
}

func (c *OVClient) GetAssociatedUplinkGroup(id string) ([]string, error) {
	var (
		uri          = "/rest/ethernet-networks/"
		uplinkGroups = new([]string)
	)
	uri = uri + id + "/associatedUplinkGroups"
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return *uplinkGroups, err
	}
	log.Infof("GetAssociatedUplinkGroups %s", data)
	if err := json.Unmarshal([]byte(data), uplinkGroups); err != nil {
		return *uplinkGroups, err
	}
	return *uplinkGroups, nil
}

func (c *OVClient) CreateEthernetNetwork(eNet EthernetNetwork) error {
	log.Infof("Initializing creation of ethernet network for %s.", eNet.Name)
	var (
		uri = "/rest/ethernet-networks"
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, eNet)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, eNet)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new ethernet network request: %s", err)
		return err
	}

	log.Debugf("Response New EthernetNetwork %s", data)
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

func (c *OVClient) CreateBulkEthernetNetwork(eNet BulkEthernetNetwork) error {
	log.Infof("Initializing creation of bulk ethernet network")
	var (
		uri = "rest/ethernet-networks/bulk"
		t   *Task
	)
	//refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST :%s \n %+v\n", uri, eNet)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, eNet)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new bulk ethernet network request: %s", err)
		return err
	}

	log.Debugf("Response New Bulk EthernetNetwork %s", data)
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

func (c *OVClient) DeleteEthernetNetwork(name string) error {
	var (
		eNet EthernetNetwork
		err  error
		t    *Task
		uri  string
	)

	eNet, err = c.GetEthernetNetworkByName(name)
	if err != nil {
		return err
	}
	if eNet.Name != "" {
		t = t.NewProfileTask(c)
		t.ResetTask()
		log.Debugf("REST : %s \n %+v\n", eNet.URI, eNet)
		log.Debugf("task -> %+v", t)
		uri = eNet.URI.String()
		if uri == "" {
			log.Warn("Unable to post delete, no uri found.")
			t.TaskIsDone = true
			return err
		}
		data, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete ethernet network request: %s", err)
			t.TaskIsDone = true
			return err
		}

		log.Debugf("Response delete ethernet network %s", data)
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
		log.Infof("EthernetNetwork could not be found to delete, %s, skipping delete ...", name)
	}
	return nil
}

func (c *OVClient) DeleteBulkEthernetNetwork(eNet BulkDelete) error {
	log.Infof("Initializing deletion of bulk ethernet network")
	var (
		uri = "rest/ethernet-networks/bulk-delete"
		t   *Task
	)
	//refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST :%s \n %+v\n", uri, eNet)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, eNet)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new bulk delete ethernet network request: %s", err)
		return err
	}

	log.Debugf("Response of Bulk Delete for EthernetNetwork %s", data)
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

func (c *OVClient) UpdateEthernetNetwork(eNet EthernetNetwork) error {
	log.Infof("Initializing update of ethernet network for %s.", eNet.Name)
	var (
		uri = eNet.URI.String()
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, eNet)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, eNet)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update ethernet network request: %s", err)
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
