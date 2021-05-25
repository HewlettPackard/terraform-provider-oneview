package ov

import (
	"encoding/json"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type Locales struct {
	DisplayName string `json:"displayname,omitempty"`
	Locale      string `json:"locale,omitempty"`
	LocaleName  string `json:"localename,omitempty"`
}

type LocalesList struct {
	Count       int           `json:"count,omitempty"`
	Category    string        `json:"category,omitempty"`
	Created     string        `json:"created,omitempty"`
	ETAG        string        `json:"eTag,omitempty"`
	Members     []Locales     `json:"members,omitempty"`     // "members":[]
	Modified    string        `json:"modified,omitempty"`    // "modified":null"
	NextPageURI utils.Nstring `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	PrevPageURI utils.Nstring `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	Start       int           `json:"start,omitempty"`       // "start": 0,
	Total       int           `json:"total,omitempty"`       // "total": 1,
	Type        string        `json:"type,omitempty"`
	URI         utils.Nstring `json:"uri,omitempty"`
}

func (c *OVClient) GetLocales() (LocalesList, error) {
	var (
		uri        = "/rest/appliance/configuration/timeconfig/locales"
		localelist LocalesList
	)

	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return localelist, err
	}

	log.Debugf("GetLocalelist %s", data)
	if err := json.Unmarshal(data, &localelist); err != nil {
		return localelist, err
	}
	return localelist, nil
}
