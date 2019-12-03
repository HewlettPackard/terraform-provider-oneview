package ov

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type Location struct {
	LocationEntries []LocationEntries `json:"locationEntries,omitempty"` //"locationEntries":{...},
}

type LocationEntries struct {
	Value string `json:"value,omitempty"` //"value":"2"
	Type  string `json:"type,omitempty"`  //"type":"StackingMemberId",

}

type PortConfigInfos struct {
	DesiredSpeed     string            `json:"desiredSpeed,omitempty"` //"desiredSpeed":"Auto",
	ExpectedNeighbor *ExpectedNeighbor `json:"expectedNeighbor"`       //"expectedNeighbor":"",
	Location         Location          `json:"location"`               //"location":"{...},
	PortUri          string            `json:"portUri"`                //"portUri:"",
}

type ExpectedNeighbor struct {
	RemoteChassisId string `json:"remoteChassisId,omitempty"` //"remoteChassisId":"",
	RemotePortId    string `json:"remotePortId,omitempty"`    //"remotePortId":"",
}

type PrivateVlanDomains struct {
	IsolatedNetwork *VlanAttributes `json:"isolatedNetwork,omitempty"`
	PrimaryNetwork  *VlanAttributes `json:"primaryNetwork,omitempty"`
}

type VlanAttributes struct {
	Name   string `json:"name,omitempty"`
	Uri    string `json:"uri,omitempty"`
	VlanId string `json:"vlanId,omitempty"`
}

type UplinkSet struct {
	Name                           string               `json:"name,omitempty"`                           // "name": "Uplink77",
	LogicalInterconnectURI         utils.Nstring        `json:"logicalInterconnectUri,omitempty"`         // "logicalInterconnectUri": "/rest/logical-interconnects/7769cae0-b680-435b-9b87-9b864c81657f",
	NetworkURIs                    []utils.Nstring      `json:"networkUris,omitempty"`                    // "networkUris": "/rest/ethernet-networks/e2f0031b-52bd-4223-9ac1-d91cb519d548",
	FcNetworkURIs                  []utils.Nstring      `json:"fcNetworkUris"`                            // "fcNetworkUris": "[]",
	FcoeNetworkURIs                []utils.Nstring      `json:"fcoeNetworkUris"`                          // "fcoeNetworkUris": "[]",
	PortConfigInfos                []PortConfigInfos    `json:"portConfigInfos"`                          // "portConfigInfos": "[]",
	ConnectionMode                 string               `json:"connectionMode,omitempty"`                 // "connectionMode":"Auto",
	NetworkType                    string               `json:"networkType,omitempty"`                    // "networkType":"Ethernet",
	EthernetNetworkType            string               `json:"ethernetNetworkType,omitempty"`            // "ethernetNetworkType":"Tagged",
	ManualLoginRedistributionState string               `json:"manualLoginRedistributionState,omitempty"` //"manualLoginRedistributionState":"NotSupported"
	URI                            utils.Nstring        `json:"uri,omitempty"`                            // "uri": "/rest/uplink-sets/"e2f0031b-52bd-4223-9ac1-d91cb519d548",
	Type                           string               `json:"type,omitempty"`                           // "type": "uplink-setV5",
	Category                       string               `json:"category,omitempty"`                       //"category":"uplink-sets",
	Created                        string               `json:"created,omitempty"`                        //"created":"20150831T154835.250Z",
	Description                    utils.Nstring        `json:"description,omitempty"`                    // "description": "Uplink-set 1",
	Etag                           string               `json:"eTag,omitempty"`                           // "eTag": "1441036118675/8",
	Modified                       string               `json:"modified,omitempty"`                       // "modified": "20150831T154835.250Z",
	LacpTimer                      string               `json:"lacpTimer,omitempty"`                      // "lacpTimer": "Long",
	FcMode                         string               `json:"fcMode,omitempty"`                         // "fcMode": "TRUNK",
	NativeNetworkUri               utils.Nstring        `json:"nativeNetworkUri,omitempty"`               // "nativeNetworkUri": null,
	PrimaryPortLocation            *Location            `json:"primaryPort,omitempty"`                    // "primaryPort": {...},
	Reachability                   string               `json:"reachability,omitempty"`                   // "reachability": "Reachable",
	State                          string               `json:"state,omitempty"`                          // "state": "Normal",
	Status                         string               `json:"status,omitempty"`                         // "status": "Critical",
	PrivateVlanDomains             []PrivateVlanDomains `json:"privateVlanDomains,omitempty"`             //"privateVlanDomains: []",

}

type UplinkSetList struct {
	Total       int           `json:"total,omitempty"`       // "total": 1,
	Count       int           `json:"count,omitempty"`       // "count": 1,
	Start       int           `json:"start,omitempty"`       // "start": 0,
	PrevPageURI utils.Nstring `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         string        `json:"uri,omitempty"`         // "uri": "/rest/uplink-sets?start=0&count=10"
	Members     []UplinkSet   `json:"members,omitempty"`     // "members":[]
}

func (c *OVClient) GetUplinkSetByName(name string) (UplinkSet, error) {
	var (
		upSet UplinkSet
	)
	upSets, err := c.GetUplinkSets("", "", fmt.Sprintf("name matches '%s'", name), "name:asc")
	if upSets.Total > 0 {
		return upSets.Members[0], err
	} else {
		return upSet, err
	}
}

func (c *OVClient) GetUplinkSets(start string, count string, filter string, sort string) (UplinkSetList, error) {
	var (
		uri        = "/rest/uplink-sets"
		q          map[string]interface{}
		uplinkSets UplinkSetList
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
		return uplinkSets, err
	}

	log.Debugf("GetUplinkSets %s", data)
	if err := json.Unmarshal([]byte(data), &uplinkSets); err != nil {
		return uplinkSets, err
	}
	return uplinkSets, nil
}

func (c *OVClient) GetUplinkSetById(id string) ([]string, error) {
	var (
		uri         = "/rest/uplink-sets/"
		uplinkSetId = new([]string)
	)
	uri = uri + id
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return *uplinkSetId, err
	}
	if err := json.Unmarshal([]byte(data), uplinkSetId); err != nil {
		return *uplinkSetId, err
	}
	return *uplinkSetId, nil
}

func (c *OVClient) CreateUplinkSet(upSet UplinkSet) error {
	log.Infof("Initializing creation of uplink-set for %s.", upSet.Name)
	var (
		uri = "/rest/uplink-sets"
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, upSet)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, upSet)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new Uplink Set request: %s", err)
		return err
	}

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

func (c *OVClient) DeleteUplinkSet(name string) error {
	var (
		upSet UplinkSet
		err   error
		t     *Task
		uri   string
	)

	upSet, err = c.GetUplinkSetByName(name)
	if err != nil {
		return err
	}
	if upSet.Name != "" {
		t = t.NewProfileTask(c)
		t.ResetTask()
		log.Debugf("REST : %s \n %+v\n", upSet.URI, upSet)
		log.Debugf("task -> %+v", t)
		uri = upSet.URI.String()
		if uri == "" {
			log.Warn("Unable to post delete, no uri found.")
			t.TaskIsDone = true
			return err
		}
		data, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete uplink-set  request: %s", err)
			t.TaskIsDone = true
			return err
		}

		log.Debugf("Response delete uplink-set %s", data)
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
		log.Infof("uplink-set could not be found to delete, %s, skipping delete ...", name)
	}
	return nil
}

func (c *OVClient) UpdateUplinkSet(upSet UplinkSet) error {
	log.Infof("Initializing update of uplink-set for %s.", upSet.Name)
	var (
		uri = upSet.URI.String()
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, upSet)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, upSet)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update uplink-set request: %s", err)
		return err
	}

	log.Debugf("Response update Uplink-set %s", data)
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
