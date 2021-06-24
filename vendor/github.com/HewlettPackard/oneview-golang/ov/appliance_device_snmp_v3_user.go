package ov

import (
	"encoding/json"
	"fmt"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type SNMPv3User struct {
	AuthenticationPassphrase string        `json:"authenticationPassphrase,omitempty"`
	AuthenticationProtocol   string        `json:"authenticationProtocol,omitempty"`
	Category                 string        `json:"category,omitempty"`
	Created                  string        `json:"created,omitempty"`
	ETAG                     string        `json:"eTag,omitempty"`
	Id                       string        `json:"id,omitempty"`
	Modified                 string        `json:"modified,omitempty"`
	PrivacyPassphrase        string        `json:"privacyPassphrase,omitempty"`
	PrivacyProtocol          string        `json:"privacyProtocol,omitempty"`
	SecurityLevel            string        `json:"securityLevel,omitempty"`
	Type                     string        `json:"type,omitempty"`
	URI                      utils.Nstring `json:"uri,omitempty"`
	UserName                 string        `json:"userName,omitempty"`
}

type SNMPv3UserList struct {
	Category    string        `json:"category,omitempty"`
	Count       int           `json:"count,omitempty"`
	Created     string        `json:"created,omitempty"`
	ETAG        string        `json:"eTag,omitempty"`
	Members     []SNMPv3User  `json:"members,omitempty"`     // "members":[]
	Modified    string        `json:"modified,omitempty"`    // "modified":null"
	NextPageURI utils.Nstring `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	PrevPageURI utils.Nstring `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	Start       int           `json:"start,omitempty"`       // "start": 0,
	Total       int           `json:"total,omitempty"`       // "total": 1,
	Type        string        `json:"type,omitempty"`
	URI         utils.Nstring `json:"uri,omitempty"`
}

func (c *OVClient) CreateSNMPv3Users(snmpv3User SNMPv3User) (SNMPv3User, error) {
	log.Infof("Initializing creation of  USM user for %s.", snmpv3User.UserName)
	var (
		uri = "/rest/appliance/snmpv3-trap-forwarding/users"
		t   = (&Task{}).NewProfileTask(c)
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, snmpv3User)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, snmpv3User)
	if err != nil {

		log.Errorf("Error submitting new snmp v3 user request: %s", err)
		return snmpv3User, err
	}

	log.Debugf("Response New snmpv3User %s", data)
	if err := json.Unmarshal(data, &t); err != nil {

		log.Errorf("Error with task un-marshal: %s", err)
		return snmpv3User, err
	}

	if err != nil {
		return snmpv3User, err
	}

	return snmpv3User, nil
}

func (c *OVClient) GetSNMPv3Users(start string, count string, filter string, sort string) (SNMPv3UserList, error) {
	var (
		uri            = "/rest/appliance/snmpv3-trap-forwarding/users"
		q              = make(map[string]interface{})
		snmpv3userlist SNMPv3UserList
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
		return snmpv3userlist, err
	}

	log.Debugf("Get v1 Trap list: %s", data)
	if err := json.Unmarshal(data, &snmpv3userlist); err != nil {
		return snmpv3userlist, err
	}
	return snmpv3userlist, nil

}

func (c *OVClient) GetSNMPv3UserById(id string) (SNMPv3User, error) {
	var (
		uri          = "/rest/appliance/snmpv3-trap-forwarding/users/" + id
		snmpv3userid SNMPv3User
	)

	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return snmpv3userid, err
	}

	log.Debugf("Get SNMP users By Id %s", data)
	if err := json.Unmarshal(data, &snmpv3userid); err != nil {
		return snmpv3userid, err
	}
	return snmpv3userid, nil
}

func (c *OVClient) GetSNMPv3UserByUserName(username string) (SNMPv3User, error) {
	var (
		snmpv3userusername SNMPv3User
	)

	snmpv3userusernames, err := c.GetSNMPv3Users("", "", fmt.Sprintf("userName matches '%s'", username), "userName:asc")
	if snmpv3userusernames.Total > 0 {
		return snmpv3userusernames.Members[0], err
	} else {
		return snmpv3userusername, err
	}

}

func (c *OVClient) UpdateSNMPv3User(updateOption SNMPv3User, id string) (SNMPv3User, error) {
	var (
		uri            = "/rest/appliance/snmpv3-trap-forwarding/users/" + id
		updateResponse SNMPv3User
	)

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	log.Debugf("REST : %s \n %+v\n", uri, updateOption)
	data, err := c.RestAPICall(rest.PUT, uri, updateOption)
	if err != nil {
		log.Errorf("Error submitting update SNMPv3 User  request: %s", err)
		return updateResponse, err
	}

	log.Debugf("Response SNMPv3 User %s", data)

	if err := json.Unmarshal(data, &updateResponse); err != nil {
		return updateResponse, err
	}

	return updateResponse, nil
}

func (c *OVClient) DeleteSNMPv3UserById(id string) error {
	var (
		snmpv3User SNMPv3User
		err        error
		uri        string
	)

	snmpv3User, err = c.GetSNMPv3UserById(id)

	if err != nil {
		return err
	}
	if snmpv3User.URI != "" {
		log.Debugf("REST : %s \n %+v\n", snmpv3User.URI, snmpv3User)

		uri = snmpv3User.URI.String()
		if uri == "" {
			log.Warn("Unable to post delete, no uri found.")
			return err
		}
		data, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete snmpv3 user request: %s", err)
			return err
		}
		log.Debugf("Response delete  snmpv3 user %s", data)

		return nil
	} else {
		log.Infof("SNMP V3 Usercould not be found to delete, %s, skipping delete ...", id)
	}
	return nil
}

func (c *OVClient) DeleteSNMPv3UserByName(username string) error {
	var (
		uri = "/rest/appliance/snmpv3-trap-forwarding/users"
		q   = make(map[string]interface{})
	)

	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	filter := "'username'=" + username
	if filter != "" {
		q["filter"] = filter
	}

	if len(q) > 0 {
		c.SetQueryString(q)
	}
	data, err := c.RestAPICall(rest.DELETE, uri, nil)
	if err != nil {
		log.Errorf("Error submitting delete snmpv3 user request: %s", err)

		return err
	}
	log.Debugf("Response delete snmpv3 user network %s", data)

	return nil

}
