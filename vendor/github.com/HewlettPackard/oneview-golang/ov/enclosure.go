package ov

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type Enclosure struct {
	ActiveOaPreferredIP  string               `json:"activeOaPreferredIP,omitempty"`  // "activeOaPreferredIP": "16.124.135.110",
	AssetTag             string               `json:"assetTag,omitempty"`             // "assetTag": "",
	Category             string               `json:"category,omitempty"`             // "category": "enclosures",
	Created              string               `json:"created,omitempty"`              // "created": "20150831T154835.250Z",
	Description          string               `json:"description,omitempty"`          // "description": "Enclosure Group 1",
	DeviceBayCount       int                  `json:"deviceBayCount,omitempty"`       // "deviceBayCount": 16,
	DeviceBays           []DeviceBayMap       `json:"deviceBays,omitempty`            // "deviceBays": [],
	ETAG                 string               `json:"eTag,omitempty"`                 // "eTag": "1441036118675/8",
	EnclosureGroupUri    utils.Nstring        `json:"enclosureGroupUri,omitempty"`    // "enclosureGroupUri": "/rest/enclosure-groups/293e8efe-c6b1-4783-bf88-2d35a8e49071",
	EnclosureType        string               `json:"enclosureType,omitempty"`        // "enclosureType": "BladeSystem c7000 Enclosure",
	FwBaselineName       string               `json:"fwBaselineName,omitempty"`       // "fwBaselineName": null,
	FwBaselineUri        utils.Nstring        `json:"fwBaselineUri,omitempty"`        // "fwBaselineUri": null,
	InterconnectBayCount int                  `json:"interconnectBayCount,omitempty"` // "interconnectBayCount": 8,
	InterconnectBays     []InterconnectBayMap `json:"interconnectBayMappings"`        // "interconnectBays": [],
	IsFwManaged          bool                 `json:"isFwManaged"`                    // "isFwManaged": false,
	LicensingIntent      string               `json:"licensingIntent,omitempty"`      // "licensingIntent": "OneView",
	Modified             string               `json:"modified,omitempty"`             // "modified": "20150831T154835.250Z",
	Name                 string               `json:"name,omitempty"`                 // "name": "e10",
	OA                   []OAMap              `json:"oa,omitempty"`                   // "oa": [],
	OaBayCount           int                  `json:"oaBayCount,omitempty"`           // "oaBayCount": 2,
	PartNumber           string               `json:"partNumber,omitempty"`           // "partNumber": "403320-B21",
	RackName             string               `json:"rackName,omitempty"`             // "rackName": "Rack-Renamed",
	RefreshState         string               `json:"refreshState,omitempty"`         // "refreshState": "NotRefreshing",
	SerialNumber         string               `json:"serialNumber,omitempty"`         // "serialNumber": "USE62519EE",
	StandbyOaPreferredIP string               `json:"standbyOaPreferredIP,omitempty"` // "standbyOaPreferredIP": "",
	State                string               `json:"state,omitempty"`                // "state": "Configured",
	StateReason          string               `json:"StateReason"`                    // "stateReason": "None",
	Status               string               `json:"status,omitempty"`               // "status": "Critical",
	Type                 string               `json:"type,omitempty"`                 // "type": "Enclosure",
	URI                  utils.Nstring        `json:"uri,omitempty"`                  // "uri": "/rest/enclosures/09USE62519EE",
	UUID                 string               `json:"uuid,omitempty"`                 // "uuid": "09USE62519EE",
	VcmDomainId          string               `json:"vcmDomainId,omitempty"`          // "vcmDomainId": "@914ae756bdbce70cf7cbce65d34a23",
	VcmDomainName        string               `json:"vcmDomainName,omitempty"`        // "vcmDomainName": "OneViewDomain",
	VcmMode              bool                 `json:"vcmMode,omitempty"`              // "vcmMode": true,
	VcmUrl               string               `json:"vcmUrl,omitempty"`               // "vcmUrl": "https://16.124.128.80"
}

type EnclosureList struct {
	Total       int           `json:"total,omitempty"`       // "total": 1,
	Count       int           `json:"count,omitempty"`       // "count": 1,
	Start       int           `json:"start,omitempty"`       // "start": 0,
	PrevPageURI utils.Nstring `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         utils.Nstring `json:"uri,omitempty"`         // "uri": "/rest/enclosures?sort=name:asc"
	Members     []Enclosure   `json:"members,omitempty"`     // "members":[]
}

type DeviceBayMap struct {
	AvailableForFullHeightProfile bool          `json:"availableForFullHeightProfile"` // "availableForFullHeightProfile": false,
	AvailableForHalfHeightProfile bool          `json:"availableForHalfHeightProfile"` // "availableForHalfHeightProfile": true,
	BayNumber                     int           `json:"bayNumber"`                     // "bayNumber": 1,
	Category                      string        `json:"category,omitempty"`            // "category": "device-bays",
	CoveredByDevice               utils.Nstring `json:"coveredByDevice,omitempty"`     // "coveredByDevice": "/rest/server-hardware/30373237-3132-4D32-3236-303730344E54",
	CoveredByProfile              string        `json:"coveredByProfile,omitempty"`    // "coveredByProfile": null,
	Created                       string        `json:"created,omitempty"`             // "created": null,
	DevicePresence                string        `json:"devicePresence,omitempty"`      // "devicePresence": "Present",
	DeviceUri                     utils.Nstring `json:"deviceUri,omitempty"`           // "deviceUri": "/rest/server-hardware/30373237-3132-4D32-3236-303730344E54",
	EnclosureUri                  utils.Nstring `json:"enclosureUri,omitempty"`        // "enclosureUri": null,
	ETAG                          string        `json:"eTag,omitempty"`                // "eTag": null,
	Model                         string        `json:"model,omitempty"`               // "model": null,
	Modified                      string        `json:"modified,omitempty"`            // "modified": null,
	ProfileUri                    utils.Nstring `json:"profileUri,omitempty"`          // "profileUri": null,
	Type                          string        `json:"type,omitempty"`                // "type": "DeviceBay",
	URI                           utils.Nstring `json:"uri,omitempty"`                 // "uri": "/rest/enclosures/09USE62519EE/device-bays/1"
}

type OAMap struct {
	BayNumber      int             `json:"bayNumber"`               // "bayNumber": 1,
	DhcpEnable     bool            `json:"dhcpEnable"`              // "dhcpEnable": false,
	DhcpIpv6Enable bool            `json:"dhcpIpv6Enable"`          // "dhcpIpv6Enable": false,
	FqdnHostName   string          `json:"fqdnHostName,omitempty"`  // "fqdnHostName": "e10-oa.vse.rdlabs.hpecorp.net",
	FwBuildDate    string          `json:"fwBuildDate,omitempty"`   // "fwBuildDate": "Jun 17 2016",
	FwVersion      string          `json:"fwVersion,omitempty"`     // "fwVersion": "4.60",
	IpAddress      string          `json:"ipAddress,omitempty"`     // "ipAddress": "16.124.135.110",
	Ipv6Addresses  []Ipv6Addresses `json:"ipv6Addresses,omitempty"` // "ipv6Addresses": []
	Role           string          `json:"role,omitempty"`          // "role": "Active",
	State          string          `json:"state,omitempty"`         // "state": null

}

type Ipv6Addresses struct {
	Address string `json:"address,omitempty"` // "address": "",
	Type    string `json:"type,omitempty"`    // "type": "NotSet"
}

type EnclosurePatchMap struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}

type EnclosureCreateMap struct {
	EnclosureGroupUri    utils.Nstring `json:"enclosureGroupUri"`
	Hostname             string        `json:"hostname"`
	Username             string        `json:"username"`
	Password             string        `json:"password"`
	LicensingIntent      string        `json:"licensingIntent"`
	ForceInstallFirmware bool          `json:"forceInstallFirmware,omitempty"`
	FirmwareBaselineUri  string        `json:"firmwareBaselineUri,omitempty"`
	Force                bool          `json:"force,omitempty"`
	InitialScopeUris     []string      `json:"initialScopeUris"`
	UpdateFirmwareOn     string        `json:"updateFirmwareOn,omitempty"`
}

func (c *OVClient) GetEnclosureByName(name string) (Enclosure, error) {
	var (
		enclosure Enclosure
	)
	enclosures, err := c.GetEnclosures("", "", fmt.Sprintf("name matches '%s'", name), "name:asc", "")
	if enclosures.Total > 0 {
		return enclosures.Members[0], err
	} else {
		return enclosure, err
	}
}

func (c *OVClient) GetEnclosurebyUri(uri utils.Nstring) (Enclosure, error) {
	var (
		enclosure Enclosure
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri.String(), nil)
	if err != nil {
		return enclosure, err
	}
	log.Debugf("GetEnclosure %s", data)
	if err := json.Unmarshal([]byte(data), &enclosure); err != nil {
		return enclosure, err
	}
	return enclosure, nil
}

func (c *OVClient) GetEnclosures(start string, count string, filter string, sort string, scopeUris string) (EnclosureList, error) {
	var (
		uri        = "/rest/enclosures"
		q          map[string]interface{}
		enclosures EnclosureList
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
		return enclosures, err
	}
	log.Debugf("GetEnclosures %s", data)
	if err := json.Unmarshal([]byte(data), &enclosures); err != nil {
		return enclosures, err
	}
	return enclosures, nil
}

func (c *OVClient) CreateEnclosure(enclosure_create_map EnclosureCreateMap) error {
	log.Debugf("Initializing creation of enclosure")
	var (
		uri = "/rest/enclosures"
		t   *Task
	)

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	t.ExpectedDuration = 30000

	data, err := c.RestAPICall(rest.POST, uri, enclosure_create_map)
	if err != nil {
		log.Errorf("Error submitting new enclosure request: %s", err)
		return err
	}

	log.Debugf("Response New Enclosure %s", data)
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

func (c *OVClient) DeleteEnclosure(name string) error {
	var (
		enclosure Enclosure
		err       error
		t         *Task
		uri       string
	)

	enclosure, err = c.GetEnclosureByName(name)
	if err != nil {
		return err
	}
	if enclosure.Name != "" {
		t = t.NewProfileTask(c)
		t.ResetTask()
		log.Debugf("REST : %s \n %+v\n", enclosure.URI, enclosure)
		log.Debugf("task -> %+v", t)
		uri = enclosure.URI.String()
		if uri == "" {
			log.Warn("Unable to post delete, no uri found.")
			t.TaskIsDone = true
			return err
		}
		_, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete enclosure request: %s", err)
			t.TaskIsDone = true
			return err
		}

		return nil
	} else {
		log.Debugf("Enclosure could not be found to delete, %s, skipping delete ...", name)
	}
	return nil
}

func (c *OVClient) UpdateEnclosure(op string, path string, value string, enclosure Enclosure) error {
	log.Debugf("Initializing update of enclosure for %s.", enclosure.Name)
	var (
		uri          = enclosure.URI.String()
		t            *Task
		enc_pat_reqs [1]EnclosurePatchMap
	)
	enc_pat_reqs[0] = EnclosurePatchMap{
		Op:    op,
		Path:  path,
		Value: value,
	}

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, enc_pat_reqs)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PATCH, uri, enc_pat_reqs)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update enclosure request: %s", err)
		return err
	}

	log.Debugf("Response Update Enclosure %s")
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
