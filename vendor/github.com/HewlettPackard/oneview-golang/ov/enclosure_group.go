package ov

import (
	"encoding/json"
	"fmt"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type EnclosureGroup struct {
	AmbientTemperatureMode              string               `json:"ambientTemperatureMode,omitempty"`       // "ambientTemperatureMode": "Standard"
	AssociatedLogicalInterconnectGroups []string             `json:"associatedInterconnectGroups,omitempty"` // "associatedInterconnectGorups": [],
	Category                            string               `json:"category,omitempty"`                     // "category": "enclosure-groups",
	Created                             string               `json:"created,omitempty"`                      // "created": "20150831T154835.250Z",
	Description                         utils.Nstring        `json:"description,omitempty"`                  // "description": "Enclosure Group 1",
	ETAG                                string               `json:"eTag,omitempty"`                         // "eTag": "1441036118675/8",
	EnclosureCount                      int                  `json:"enclosureCount,omitempty"`               // "enclosureCount": 1,
	EnclosureTypeUri                    utils.Nstring        `json:"enclosureTypeUri,omitempty"`             // "enclosureTypeUri": "/rest/enclosures/e2f0031b-52bd-4223-9ac1-d91cb5219d548"
	InitialScopeUris                    []utils.Nstring      `json:"initialScopeUris,omitempty"`             // "initialScopeUris":[]
	InterconnectBayMappingCount         int                  `json:"interconnectBayMappingCount,omitempty"`  // "interconnectBayMappingCount": 8,
	InterconnectBayMappings             []InterconnectBayMap `json:"interconnectBayMappings"`                // "interconnectBayMappings": [],
	IpAddressingMode                    string               `json:"ipAddressingMode,omitempty"`             // "ipAddressingMode": "DHCP"
	IpRangeUris                         []utils.Nstring      `json:"ipRangeUris,omitempty"`
	Ipv6AddressingMode                  string               `json:"ipv6AddressingMode,omitempty"` // "ipAddressingMode": "DHCP"
	Ipv6RangeUris                       []utils.Nstring      `json:"ipv6RangeUris,omitempty"`
	Modified                            string               `json:"modified,omitempty"`             // "modified": "20150831T154835.250Z",
	Name                                string               `json:"name,omitempty"`                 // "name": "Enclosure Group 1",
	OsDeploymentSettings                *OsDeploymentSetting `json:"osDeploymentSettings,omitempty"` // "osDeploymentSetting": {},
	PortMappingCount                    int                  `json:"portMappingCount,omitempty"`     // "portMappingCount": 1,
	PortMappings                        []PortMap            `json:"portMappings,omitempty"`         // "portMappings": [],
	PowerMode                           string               `json:"powerMode,omitempty"`            // "powerMode": RedundantPowerFeed,
	ScopesUri                           utils.Nstring        `json:"scopesUri,omitempty"`            // "ScopesUri": "/rest/scopes/resources/rest/enclosure-groups/2b322628-e5a9-4843-b184-08345e7140c3",
	StackingMode                        string               `json:"stackingMode,omitempty"`         // "stackingMode": "Enclosure"
	State                               string               `json:"state,omitempty"`                // "state": "Normal",
	Status                              string               `json:"status,omitempty"`               // "status": "Critical",
	Type                                string               `json:"type,omitempty"`                 // "type": "EnclosureGroupV200",
	URI                                 utils.Nstring        `json:"uri,omitempty"`                  // "uri": "/rest/enclosure-groups/e2f0031b-52bd-4223-9ac1-d91cb519d548"
}

type EnclosureGroupList struct {
	Total       int              `json:"total,omitempty"`       // "total": 1,
	Count       int              `json:"count,omitempty"`       // "count": 1,
	Start       int              `json:"start,omitempty"`       // "start": 0,
	PrevPageURI utils.Nstring    `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring    `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         utils.Nstring    `json:"uri,omitempty"`         // "uri": "/rest/server-profiles?filter=connectionTemplateUri%20matches%7769cae0-b680-435b-9b87-9b864c81657fsort=name:asc"
	Members     []EnclosureGroup `json:"members,omitempty"`     // "members":[]
}

type InterconnectBayMap struct {
	EnclosureIndex              int           `json:"enclosureIndex,omitempty"`              // "enclosureIndex": 0,
	InterconnectBay             int           `json:"interconnectBay,omitempty"`             // "interconnectBay": 0,
	LogicalInterconnectGroupUri utils.Nstring `json:"logicalInterconnectGroupUri,omitempty"` // "logicalInterconnectGroupUri": "",
}

type PortMap struct {
	InterconnectBay int `json:"interconnectBay,omitempty"` // "interconnectBay": 1,
	MidplanePort    int `json:"midplanePort,omitempty"`    // "midplanePort": 1,
}

type OsDeploymentSetting struct {
	DeploymentModeSettings DeploymentModeSetting `json:"deploymentModeSettings,omitempty"` // "deploymentModeSettings": {},
	ManageOSDeployment     bool                  `json:"manageOSDeployment,omitempty"`     // "manageOSDeployment": false,
}

type DeploymentModeSetting struct {
	DeploymentMode       string `json:"deploymentMode,omitempty"`       // "deploymentMode": "None",
	DeploymentNetworkUri string `json:"deploymentNetworkUri,omitempty"` // "deploymentNetworkUri": null,
}

func (c *OVClient) GetEnclosureGroupByName(name string) (EnclosureGroup, error) {
	var (
		enclosureGroup EnclosureGroup
	)
	enclosureGroups, err := c.GetEnclosureGroups("", "", fmt.Sprintf("name matches '%s'", name), "name:asc", "")
	if enclosureGroups.Total > 0 {
		return enclosureGroups.Members[0], err
	} else {
		return enclosureGroup, err
	}
}

func (c *OVClient) GetEnclosureGroupByUri(uri utils.Nstring) (EnclosureGroup, error) {
	var (
		enclosureGroup EnclosureGroup
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri.String(), nil)
	if err != nil {
		return enclosureGroup, err
	}
	log.Debugf("GetEnclosureGroup %s", data)
	if err := json.Unmarshal([]byte(data), &enclosureGroup); err != nil {
		return enclosureGroup, err
	}
	return enclosureGroup, nil
}

func (c *OVClient) GetEnclosureGroups(start string, count string, filter string, sort string, scopeUris string) (EnclosureGroupList, error) {
	var (
		uri             = "/rest/enclosure-groups"
		q               map[string]interface{}
		enclosureGroups EnclosureGroupList
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
		return enclosureGroups, err
	}

	log.Debugf("GetEnclosureGroups %s", data)
	if err := json.Unmarshal([]byte(data), &enclosureGroups); err != nil {
		return enclosureGroups, err
	}
	return enclosureGroups, nil
}

func (c *OVClient) CreateEnclosureGroup(eGroup EnclosureGroup) error {
	log.Infof("Initializing creation of enclosure group for %s.", eGroup.Name)
	var (
		uri = "/rest/enclosure-groups"
		t   *Task
	)

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	data, err := c.RestAPICall(rest.POST, uri, eGroup)
	if err != nil {
		log.Errorf("Error submitting new enclosure group request: %s", err)
		return err
	}

	log.Debugf("Response New EnclosureGroup %s", data)
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return err
	}

	return nil
}

func (c *OVClient) DeleteEnclosureGroup(name string) error {
	var (
		enclosureGroup EnclosureGroup
		err            error
		t              *Task
		uri            string
	)

	enclosureGroup, err = c.GetEnclosureGroupByName(name)
	if err != nil {
		return err
	}
	if enclosureGroup.Name != "" {
		t = t.NewProfileTask(c)
		t.ResetTask()
		log.Debugf("REST : %s \n %+v\n", enclosureGroup.URI, enclosureGroup)
		log.Debugf("task -> %+v", t)
		uri = enclosureGroup.URI.String()
		if uri == "" {
			log.Warn("Unable to post delete, no uri found.")
			t.TaskIsDone = true
			return err
		}
		_, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete enclosure group request: %s", err)
			t.TaskIsDone = true
			return err
		}

		return nil
	} else {
		log.Infof("EnclosureGroup could not be found to delete, %s, skipping delete ...", name)
	}
	return nil
}

func (c *OVClient) UpdateEnclosureGroup(enclosureGroup EnclosureGroup) error {
	log.Infof("Initializing update of enclosure group for %s.", enclosureGroup.Name)
	var (
		uri = enclosureGroup.URI.String()
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, enclosureGroup)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, enclosureGroup)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update enclosure group request: %s", err)
		return err
	}

	log.Debugf("Response update EnclosureGroup %s", data)
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return err
	}

	return nil
}

func (c *OVClient) GetConfigurationScript(uri utils.Nstring) (string, error) {
	var (
		configuration_script string
		main_uri             = uri.String()
	)
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	main_uri = main_uri + "/script"
	script, err := c.RestAPICall(rest.GET, main_uri, nil)
	if err != nil {
		log.Errorf("Error in getting the configuration script: %s", err)
		return "", err
	}
	configuration_script = string(script)
	log.Debugf("ConfigurationScript %s", configuration_script)
	return configuration_script, nil
}

func (c *OVClient) UpdateConfigurationScript(uri utils.Nstring, body string) (string, error) {
	var (
		main_uri = uri.String()
	)

	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	main_uri = main_uri + "/script"
	log.Debugf("REST : %s \n %s\n", main_uri, body)
	data, err := c.RestAPICall(rest.PUT, main_uri, body)
	if err != nil {
		log.Errorf("Error submitting update enclosure group configure script request: %s", err)
		return "", err
	}

	log.Debugf("Response update Configuration Script %s", data)
	return string(data), nil
}
