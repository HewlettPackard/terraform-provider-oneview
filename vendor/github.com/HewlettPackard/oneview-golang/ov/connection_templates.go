package ov

import (
	"encoding/json"
	"fmt"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type BandwidthType struct {
	MaximumBandwidth int `json:"maximumBandwidth,omitempty"`
	TypicalBandwidth int `json:"typicalBandwidth,omitempty"`
}

type ConnectionTemplate struct {
	Bandwidth   BandwidthType `json:"bandwidth,omitempty"`
	Category    string        `json:"category,omitempty"`
	Created     string        `json:"created,omitempty"`
	Description utils.Nstring `json:"description,omitempty"`
	ETAG        string        `json:"eTag,omitempty"`
	Modified    string        `json:"modified,omitempty"`
	Name        string        `json:"name,omitempty"`
	State       string        `json:"state,omitempty"`
	Status      string        `json:"status,omitempty"`
	Type        string        `json:"type,omitempty"`
	URI         utils.Nstring `json:"uri,omitempty"`
}

type ConnectionList struct {
	Category    string               `json:"category,omitempty"`
	Count       int                  `json:"count,omitempty"` // "count": 1,
	Created     string               `json:"created,omitempty"`
	ETAG        string               `json:"eTag,omitempty"`
	Members     []ConnectionTemplate `json:"members,omitempty"`
	Modified    string               `json:"modified,omitempty"`
	NextPageURI utils.Nstring        `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	PrevPageURI utils.Nstring        `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	Start       int                  `json:"start,omitempty"`       // "start": 0,
	Total       int                  `json:"total,omitempty"`       // "total": 1
	Type        string               `json:"type,omitempty"`
	URI         utils.Nstring        `json:"uri,omitempty"`
}

func (c *OVClient) GetConnectionTemplateByName(name string) (ConnectionTemplate, error) {
	conntemplate, err := c.GetConnectionTemplate(fmt.Sprintf("name matches '%s'", name), "name:asc", "", "")
	if conntemplate.Total > 0 {
		return conntemplate.Members[0], err
	}

	return ConnectionTemplate{}, err
}

func (c *OVClient) GetConnectionTemplate(filter string, sort string, start string, count string) (ConnectionList, error) {
	var (
		uri            = "/rest/connection-templates"
		q              = make(map[string]interface{})
		connectionlist ConnectionList
	)
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
		return connectionlist, err
	}

	log.Debugf("GetConnectionTemplate %s", data)
	if err := json.Unmarshal(data, &connectionlist); err != nil {
		return connectionlist, err
	}
	return connectionlist, nil
}

func (c *OVClient) UpdateConnectionTemplate(id string, conntemplate ConnectionTemplate) (ConnectionTemplate, error) {
	log.Infof("Initializing update of Connection Template for %s.", id)
	var (
		uri      = "/rest/connection-templates/" + id
		template ConnectionTemplate
		t        *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()

	log.Debugf("REST : %s \n %+v\n", uri, conntemplate)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, conntemplate)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update Connection Template request: %s", err)
		return template, err
	}

	log.Debugf("Response update Connection Template %s", data)
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return template, err
	}

	if err := json.Unmarshal(data, &template); err != nil {
		return template, err
	}

	return template, nil
}

// GetConnectionTemplateByURI - get a connection template from a uri
func (c *OVClient) GetConnectionTemplateByURI(uri utils.Nstring) (ConnectionTemplate, error) {
	var (
		template ConnectionTemplate
	)

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri.String(), nil)
	if err != nil {
		return template, err
	}

	log.Debugf("GetConnectionTemplateByURI %s", data)
	if err := json.Unmarshal([]byte(data), &template); err != nil {
		return template, err
	}
	return template, nil
}

func (c *OVClient) GetDefaultConnectionTemplate() (ConnectionTemplate, error) {
	var (
		uri               = "/rest/connection-templates/defaultConnectionTemplate"
		defaultConnection ConnectionTemplate
	)

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return defaultConnection, err
	}

	log.Debugf("GetDefaultConnectionTemplate %s", data)
	if err := json.Unmarshal(data, &defaultConnection); err != nil {
		return defaultConnection, err
	}
	return defaultConnection, nil
}
