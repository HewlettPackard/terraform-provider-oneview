package ov

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

// The Marshal() will omit the bool attibutes below if they are false.
// Please remove the omitempty option and use it as and when required.

type StorageAttachment struct {
	Category         string        `json:"category,omitempty"`
	Created          string        `json:"created,omitempty"`
	Description      string        `json:"description,omitempty"`
	ETAG             string        `json:"eTag,omitempty"`
	Host             *Host         `json:"host,omitempty"`
	Paths            *Paths        `json:"paths,omitempty"`
	Name             string        `json:"name,omitempty"`
	State            string        `json:"state,omitempty"`
	Status           string        `json:"status,omitempty"`
	Type             string        `json:"type,omitempty"`
	URI              utils.Nstring `json:"uri,omitempty"`
	StorageSystemUri utils.Nstring `json:"storageSystemUri,omitempty"`
	StorageVolumeUri utils.Nstring `json:"storageVolumeUri,omitempty"`
}

type Host struct {
	Name string `json:"name,omitempty"`
	Os   string `json:"os,omitempty"`
}

type Paths struct {
	ConnectionName     string     `json:"connectionName,omitempty"`
	ExpectedNetworkUri string     `json:"expectedNetworkUri,omitempty"`
	Initiator          *Initiator `json:"initiator,omitempty"`
	IsEnabled          bool       `json:"isEnabled,omitempty"`
	Transport          string     `json:"transport,omitempty"`
}

type Initiator struct {
	Chap       *Chap       `json:"chap,omitempty"`
	MutualChap *MutualChap `json:"mutualChap,omitempty"`
	Name       string      `json:"name,omitempty"`
}

type Chap struct {
	Name   string `json:"name,omitempty"`
	secret string `json:"secret,omitempty"`
}

type MutualChap struct {
	Name   string `json:"name,omitempty"`
	secret string `json:"secret,omitempty"`
}

type StorageAttachmentsListV2 struct {
	Total       int                 `json:"total,omitempty"`       // "total": 1,
	Count       int                 `json:"count,omitempty"`       // "count": 1,
	Start       int                 `json:"start,omitempty"`       // "start": 0,
	PrevPageURI utils.Nstring       `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring       `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         utils.Nstring       `json:"uri,omitempty"`         // "uri": "/rest/storage-pools?filter=connectionTemplateUri%20matches%7769cae0-b680-435b-9b87-9b864c81657fsort=name:asc"
	Members     []StorageAttachment `json:"members,omitempty"`     // "members":[]
}

func (c *OVClient) GetStorageAttachmentByName(name string) (StorageAttachment, error) {
	var (
		sAttachment StorageAttachment
	)
	sAttachments, err := c.GetStorageAttachments(fmt.Sprintf("name matches '%s'", name), "name:asc", "", "")
	if sAttachments.Total > 0 {
		return sAttachments.Members[0], err
	} else {
		return sAttachment, err
	}
}

func (c *OVClient) GetStorageAttachments(filter string, sort string, start string, count string) (StorageAttachmentsListV2, error) {
	var (
		uri          = "/rest/storage-volume-attachments"
		q            map[string]interface{}
		sAttachments StorageAttachmentsListV2
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
		return sAttachments, err
	}

	log.Debugf("GetStorageAttachments %s", data)
	if err := json.Unmarshal([]byte(data), &sAttachments); err != nil {
		return sAttachments, err
	}
	return sAttachments, nil
}
