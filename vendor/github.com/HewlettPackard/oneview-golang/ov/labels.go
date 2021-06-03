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

package ov

import (
	"encoding/json"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type AllLabels struct {
	Category    string        `json:"category,omitempty"`
	Count       int           `json:"count,omitempty"`
	Created     string        `json:"created,omitempty"`
	ETAG        utils.Nstring `json:"eTag,omitempty"`
	Members     []Member      `json:"members,omitempty"`
	Modified    string        `json:"modified,omitempty"`
	NextPageUri utils.Nstring `json:"nextPageUri,omitempty"`
	PrevPageUri utils.Nstring `json:"prevPageUri,omitempty"`
	Start       int           `json:"start,omitempty"`
	Total       int           `json:"total,omitempty"`
	Type        string        `json:"type,omitempty"`
	Uri         utils.Nstring `json:"uri,omitempty"`
}

type Member struct {
	Category string        `json:"category,omitempty"`
	Created  string        `json:"created,omitempty"`
	ETAG     string        `json:"eTag,omitempty"`
	Modified string        `json:"modified,omitempty"`
	Name     string        `json:"name,omitempty"`
	Type     string        `json:"type,omitempty"`
	Uri      utils.Nstring `json:"uri,omitempty"`
}

type AssignedLabel struct {
	Category    string        `json:"category,omitempty"`
	Created     string        `json:"created,omitempty"`
	ETAG        utils.Nstring `json:"eTag,omitempty"`
	Labels      []Label       `json:"labels,omitempty"`
	Modified    string        `json:"modified,omitempty"`
	ResourceUri utils.Nstring `json:"resourceUri,omitempty"`
	Type        string        `json:"type,omitempty"`
	Uri         utils.Nstring `json:"uri,omitempty"`
}

type Label struct {
	Name string        `json:"name,omitempty"`
	Uri  utils.Nstring `json:"uri,omitempty"`
}

// GetAllLabels- fetches all labels
func (c *OVClient) GetAllLabels(sort string, start string, count string, fields string, namePrefix string) (AllLabels, error) {
	var (
		uri       = "/rest/labels"
		q         = make(map[string]interface{})
		allLabels AllLabels
	)

	if sort != "" {
		q["sort"] = sort
	}

	if fields != "" {
		q["fields"] = fields
	}

	if namePrefix != "" {
		q["namePrefix"] = namePrefix
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
		return allLabels, err
	}

	log.Debugf("GetAllLabels %s", data)
	if err := json.Unmarshal(data, &allLabels); err != nil {
		return allLabels, err
	}
	return allLabels, nil
}

// CreateLabel- creates new labels
func (c *OVClient) CreateLabel(label AssignedLabel) (AssignedLabel, error) {
	log.Infof("Initializing creation of labels for %s.", label.ResourceUri)
	var (
		uri      = "/rest/labels/resources"
		response AssignedLabel
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	log.Debugf("REST : %s \n %+v\n", uri, label)

	data, err := c.RestAPICall(rest.POST, uri, label)
	log.Debugf("Response New Label %s", data)

	if err != nil {
		log.Errorf("Error with the request: %s", err)
		return response, err
	}

	if err := json.Unmarshal(data, &response); err != nil {
		log.Errorf("Error with data un-marshal: %s", err)
		return response, err
	}
	return response, nil
}

// GetAssignedLabels- get labels from the resource
func (c *OVClient) GetAssignedLabels(ResourceUri utils.Nstring) (AssignedLabel, error) {
	var (
		response AssignedLabel
		uri      = "/rest/labels/resources" + ResourceUri
	)

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri.String(), nil)
	if err != nil {
		return response, err
	}

	log.Debugf("GetAssignedLabels %s", data)
	if err := json.Unmarshal([]byte(data), &response); err != nil {
		return response, err
	}
	return response, nil
}

// UpdateAssignedLabels - Set all the labels for a resource.
func (c *OVClient) UpdateAssignedLabels(assignedLabel AssignedLabel) (AssignedLabel, error) {
	log.Infof("Initializing Label updates for %s.", assignedLabel.ResourceUri)
	var (
		uri      = "/rest/labels/resources/" + string(assignedLabel.ResourceUri)
		response AssignedLabel
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	log.Debugf("REST : %s \n %+v\n", uri, assignedLabel)

	data, err := c.RestAPICall(rest.PUT, uri, assignedLabel)

	if err != nil {
		return response, err
	}

	log.Debugf("UpdateAssignedLabels %s", data)
	if err := json.Unmarshal(data, &response); err != nil {
		return response, err
	}

	return response, nil
}

// DeleteAssignedLabel- Delete all the labels for a resource.
func (c *OVClient) DeleteAssignedLabel(resourceUri string) error {
	var (
		uri = "/rest/labels/resources" + resourceUri
	)
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	if resourceUri != "" {
		log.Debugf("REST : %s \n", uri)
		data, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting request: %s", err)
			return err
		}

		log.Debugf("Response delete labels %s", data)
	} else {
		log.Infof("Resource Uri not found to delete labels, skipping delete ...")
	}
	return nil
}

// GetLabelByURI - Get a label.
func (c *OVClient) GetLabelByURI(uri utils.Nstring) (Member, error) {
	var (
		label Member
	)

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri.String(), nil)
	if err != nil {
		return label, err
	}

	log.Debugf("GetLabelByURI %s", data)
	if err := json.Unmarshal([]byte(data), &label); err != nil {
		return label, err
	}
	return label, nil
}
