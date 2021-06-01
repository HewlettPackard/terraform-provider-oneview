package ov

import (
	"encoding/json"
	"errors"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type SNMPv1Trap struct {
	CommunityString string        `json:"communityString, omitempty"`
	Destination     string        `json:"destination,omitempty"`
	Port            int           `json:"port,omitempty"`
	URI             utils.Nstring `json:"uri,omitempty"`
}

type SNMPv1TrapList struct {
	Count       int           `json:"count,omitempty"`
	Category    string        `json:"category,omitempty"`
	Created     string        `json:"created,omitempty"`
	ETAG        string        `json:"eTag,omitempty"`
	Members     []SNMPv1Trap  `json:"members,omitempty"`     // "members":[]
	Modified    string        `json:"modified,omitempty"`    // "modified":null"
	NextPageURI utils.Nstring `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	PrevPageURI utils.Nstring `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	Start       int           `json:"start,omitempty"`       // "start": 0,
	Total       int           `json:"total,omitempty"`       // "total": 1,
	Type        string        `json:"type,omitempty"`
	URI         utils.Nstring `json:"uri,omitempty"`
}

type Trapv1ValidationAddress struct {
	CommunityString string        `json:"communityString, omitempty"`
	Destination     string        `json:"destination,omitempty"`
	URI             utils.Nstring `json:"uri,omitempty"`
}

func (c *OVClient) Trapv1ValidateDestinationAddress(validate Trapv1ValidationAddress) error {
	var (
		uri = "/rest/appliance/trap-destinations/validation"
	)

	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	log.Debugf("REST: %s \n %+v\n", uri, validate.Destination)
	_, err := c.RestAPICall(rest.POST, uri, validate)

	if err != nil {
		log.Errorf("Error submitting validating destination address request: %s", err)
		return err
	}

	return nil
}

func (c *OVClient) CreateSNMPv1TrapDestinations(trapOption SNMPv1Trap, id string) error {
	log.Infof("Initializing creation of SNMPv1 Trap Destinations for %s.", trapOption.CommunityString)
	var (
		uri      = "/rest/appliance/trap-destinations/" + id
		trapdata SNMPv1Trap
	)

	//validating the Destination Address
	validate := Trapv1ValidationAddress{
		CommunityString: trapOption.CommunityString,
		Destination:     trapOption.Destination,
		URI:             utils.Nstring(uri),
	}
	log.Infof("Validating SNMPv1 Trap Destinations Address %s.", validate)
	err := c.Trapv1ValidateDestinationAddress(validate)
	if err != nil {
		return errors.New("Invalid Destination Address")
	}
	log.Infof("Successfully validated the Destination Address.")
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	log.Debugf("REST : %s \n %+v\n", uri, trapOption)
	data, err := c.RestAPICall(rest.POST, uri, trapOption)
	if err != nil {
		log.Errorf("Error submitting new snmpv1 trap destinations request: %s", err)
		return err
	}

	log.Debugf("Response New SNMPv1 Trap Destinations %s", data)
	if err := json.Unmarshal(data, &trapdata); err != nil {
		log.Errorf("Error with task un-marshal: %s", err)
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func (c *OVClient) GetSNMPv1TrapDestinations(filter string, sort string, start string, count string) (SNMPv1TrapList, error) {
	var (
		uri      = "/rest/appliance/trap-destinations/"
		q        = make(map[string]interface{})
		traplist SNMPv1TrapList
	)

	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

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
	if len(q) > 0 {
		c.SetQueryString(q)
	}
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return traplist, err
	}

	log.Debugf("Get v1 Trap list: %s", data)
	if err := json.Unmarshal(data, &traplist); err != nil {
		return traplist, err
	}
	return traplist, nil

}

func (c *OVClient) GetSNMPv1TrapDestinationsById(id string) (SNMPv1Trap, error) {
	var (
		uri    = "/rest/appliance/trap-destinations/" + id
		trapId SNMPv1Trap
	)

	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return trapId, err
	}

	log.Debugf("Get Trap By Id %s", data)
	if err := json.Unmarshal(data, &trapId); err != nil {
		return trapId, err
	}
	return trapId, nil
}

func (c *OVClient) UpdateSNMPv1TrapDestinations(updateOption SNMPv1Trap, id string) (SNMPv1Trap, error) {
	var (
		uri            = "/rest/appliance/trap-destinations/" + id
		updateResponse SNMPv1Trap
	)

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	log.Debugf("REST : %s \n %+v\n", uri, updateOption)
	data, err := c.RestAPICall(rest.PUT, uri, updateOption)
	if err != nil {
		log.Errorf("Error submitting update SNMPv1 Trap Destination request: %s", err)
		return updateResponse, err
	}

	log.Debugf("Response SNMPv1 Trap Destination %s", data)

	if err := json.Unmarshal(data, &updateResponse); err != nil {
		return updateResponse, err
	}

	return updateResponse, nil
}

func (c *OVClient) DeleteSNMPv1TrapDestinations(id string) error {
	var (
		trap SNMPv1Trap
		err  error
		uri  = "/rest/appliance/trap-destinations/" + id
	)

	trap, err = c.GetSNMPv1TrapDestinationsById(id)
	if err != nil {
		return err
	}
	if trap.Destination != "" {
		log.Debugf("REST : %s \n %+v\n", trap.CommunityString, trap)
		_, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete snmp trap destination request: %s", err)
			return err
		}
	} else {
		log.Debugf("SNMPv1TrapDestination not found, %s, skipping delete ...", id)
	}
	return nil
}
