package ov

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type LogicalEnclosure struct {
	AmbientTemperatureMode    string                     `json:"ambientTemperatureMode,omitempty"`    // "ambientTemperatureMode": "Standard",
	Category                  string                     `json:"category,omitempty"`                  // "category": "logical-enclosures",
	Created                   string                     `json:"created,omitempty"`                   // "created": "20150831T154835.250Z",
	DeleteFailed              bool                       `json:"deleteFailed,omitempty"`              // "deleteFailed": true,
	DeploymentManagerSettings *DeploymentManagerSettings `json:"deploymentManagerSettings,omitempty"` // "deploymentManagerSettings": "",
	Description               utils.Nstring              `json:"description,omitempty"`               // "description": "Logical Enclosure 1",
	Etag                      string                     `json:"eTag,omitempty"`                      // "eTag": "1441036118675/8",
	EnclosureGroupUri         utils.Nstring              `json:"enclosureGroupUri,omitempty"`         // "enclosureGroupUri": "/rest/enclosure-groups/9b8f7ec0-52b3-475e-84f4-c4eac51c2c20",
	EnclosureUris             []utils.Nstring            `json:"enclosureUris,omitempty"`             // "enclosureUris":""
	Enclosures                map[string]Enclosures      `json:"enclosures,omitempty"`                // "enclosures":"[]",
	Firmware                  *LogicalEnclosureFirmware  `json:"firmware,omitempty"`                  // "firmware":"",
	IpAddressingMode          string                     `json:"ipAddressingMode,omitempty"`          // "ipAddressingMode":"DHCP",
	Ipv4Ranges                []Ipv4Ranges               `json:"ipv4Ranges,omitempty"`                //"ipv4Ranges":"[]"
	LogicalInterconnectUris   []utils.Nstring            `json:"logicalInterconnectUris,omitempty"`   //"logicalInterconnectUris":"[]",
	Modified                  string                     `json:"modified,omitempty"`                  // "modified": "20150831T154835.25Z",
	Name                      string                     `json:"name,omitempty"`                      // "name": "Ethernet Network 1",
	PowerMode                 string                     `json:"powerMode,omitempty"`                 // "powerMode": "RedundantPowerFeed",
	ScalingState              string                     `json:"scalingState,omitempty"`              // "scalingState": "Growing",
	ScopesUri                 utils.Nstring              `json:"scopesUri,omitempty"`                 // "scopesUri":
	State                     string                     `json:"state,omitempty"`                     // "state": "Creating",
	Status                    string                     `json:"status,omitempty"`                    // "status": "Critical",
	Type                      string                     `json:"type,omitempty"`                      // "type": "LogicalEnclosureV4",
	URI                       utils.Nstring              `json:"uri,omitempty"`                       // "uri": "/rest/logical-enclosures/e2f0031b-52bd-4223-9ac1-d91cb519d548"

}

type LogicalEnclosureList struct {
	Total       int                `json:"total,omitempty"`       // "total": 1,
	Count       int                `json:"count,omitempty"`       // "count": 1,
	Start       int                `json:"start,omitempty"`       // "start": 0,
	PrevPageURI utils.Nstring      `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring      `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         utils.Nstring      `json:"uri,omitempty"`         // "uri": "/rest/ethernet-networks?filter=connectionTemplateUri%20matches%7769cae0-b680-435b-9b87-9b864c81657fsort=name:asc"
	Members     []LogicalEnclosure `json:"members,omitempty"`     // "members":[]
}

type DeploymentModeSettings struct {
	DeploymentMode       string        `json:"deploymentMode,omitempty"`       //"deploymentMode":"None"
	DeploymentNetworkUri utils.Nstring `json:"deploymentNetworkUri,omitempty"` //"deploymentNetworkUri":"/rest/ethernet-networks/e2f0031b-52bd-4223-9ac1-d91cb519d548"
}

type LeOsDeploymentSettings struct {
	DeploymentModeSettings *DeploymentModeSettings `json:"deploymentModeSettings,omitempty"`
	ManageOSDeployment     bool                    `json:"manageOSDeployment,omitempty"`
}

type DeploymentManagerSettings struct {
	DeploymentClusterUri utils.Nstring           `json:"deploymentClusterUri,omitempty"` //"deploymentClusterUri":""
	OsDeploymentSettings *LeOsDeploymentSettings `json:"osDeploymentSettings,omitempty"` //"OsdeploymentSettings":""
}

type Enclosures struct {
	EnclosureUri     utils.Nstring      `json:"enclosureUri,omitempty"`     //"enclosureUri":"",
	InterconnectBays []InterconnectBays `json:"interconnectBays,omitempty"` //"interconnectBays":"[]",
}

type InterconnectBays struct {
	BayNumber      int             `json:"bayNumber,omitempty"`      //"bayNumber":"3",
	LicenseIntents *LicenseIntents `json:"licenseIntents,omitempty"` //"licenseIntent":"",
}

type LicenseIntents struct {
	FCUpgrade string `json:"FCUpgrade,omitempty"` //"FCUpgrade":"Automatic",
}

type Ipv4Ranges struct {
	DnsServers []string      `json:"dnsServers,omitempty"` //"dnsServers":"",
	Domain     string        `json:"domain,omitempty"`     //"domain":"",
	Gateway    string        `json:"gateway,omitempty"`    //"gateway":"",
	IpRangeUri utils.Nstring `json:"ipRangeUri,omitempty"` //"ipRangeUri":"",
	Name       string        `json:"name,omitempty"`       //"name":"",
	SubnetMask string        `json:"subnetMask,omitempty"` //"subnetMask":""
}

type LogicalEnclosureFirmware struct {
	FirmwareBaselineUri                       utils.Nstring `json:"firmwareBaselineUri,omitempty"`                       //"firmwareBaselineUri":"",
	FirmwareUpdateOn                          string        `json:"firmwareUpdateOn,omitempty"`                          //"firmwareUpdateOn":"EnclosureOnly",
	ForceInstallFirmware                      bool          `json:"forceInstallFirmware,omitempty"`                      //"forceInstallFirmware":true,
	LogicalInterconnectUpdateMode             string        `json:"logicalInterconnectUpdateMode,omitempty"`             //"logicalInterconnectUpdateMode":"Parallel",
	UpdateFirmwareOnUnmanagedInterconnect     bool          `json:"updateFirmwareOnUnmanagedInterconnect,omitempty"`     //"updateFirmwareOnUnmanagedInterconnect":true,
	ValidateIfLIFirmwareUpdateIsNonDisruptive bool          `json:"validateIfLIFirmwareUpdateIsNonDisruptive,omitempty"` //"validateIfLIFirmwareUpdateIsNonDisruptive":false,
}

type SupportDumps struct {
	Encrypt                 utils.Nstring   `json:"encrypt,omitempty"`                 //""encrypt":"true",
	ErrorCode               string          `json:"errorCode,omitempty"`               //""errorCode":"MyDump16",
	ExcludeApplianceDump    bool            `json:"excludeApplianceDump,omitempty"`    //""excludeApplianceDump":false,
	LogicalInterconnectUris []utils.Nstring `json:"logicalInterconnectUris,omitempty"` //"logicalInterconnectUris":"",
}

func (c *OVClient) GetLogicalEnclosureByName(name string) (LogicalEnclosure, error) {
	var (
		logEn LogicalEnclosure
	)
	scopeUris := []string{}
	logEns, err := c.GetLogicalEnclosures("", "", fmt.Sprintf("name matches '%s'", name), scopeUris, "name:asc")
	if logEns.Total > 0 {
		return logEns.Members[0], err
	} else {
		return logEn, err
	}
}

func (c *OVClient) GetLogicalEnclosures(start string, count string, filter string, scopeUris []string, sort string) (LogicalEnclosureList, error) {
	var (
		uri               = "/rest/logical-enclosures"
		q                 map[string]interface{}
		logicalEnclosures LogicalEnclosureList
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

	if len(scopeUris) != 0 {
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
		return logicalEnclosures, err
	}

	log.Debugf("GetLogicalEnclosures %s", data)
	if err := json.Unmarshal([]byte(data), &logicalEnclosures); err != nil {
		return logicalEnclosures, err
	}
	return logicalEnclosures, nil
}
func (c *OVClient) CreateSupportDump(supportdump SupportDumps, id string) (map[string]string, error) {
	var (
		uri = "/rest/logical-enclosures/"
		t   *Task
	)
	uri = uri + id + "/support-dumps"

	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, supportdump)
	log.Debugf("task -> %+v", t)

	data, err := c.RestAPICall(rest.POST, uri, supportdump)
	payload := make(map[string]string)

	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new logical Enclosure Support Dump request: %s", err)
		return payload, err
	}

	err = json.Unmarshal([]byte(data), &payload)

	if err != nil {
		log.Errorf("Error with payload un-marshal: %s", err)
		return payload, err
	}

	log.Debugf("Response New Support Dump for LogicalEnclosure %s", data)
	if err = json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return payload, err
	}

	err = t.Wait()
	if err != nil {
		return payload, err
	}

	return payload, nil
}

func (c *OVClient) CreateLogicalEnclosure(logEn LogicalEnclosure) error {
	log.Infof("Initializing creation of logical enclosure for %s.", logEn.Name)
	var (
		uri = "/rest/logical-enclosures"
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, logEn)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, logEn)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new logical Enclosure request: %s", err)
		return err
	}

	log.Debugf("Response New LogicalEnclosure %s", data)
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

func (c *OVClient) DeleteLogicalEnclosure(name string) error {
	var (
		logEn LogicalEnclosure
		err   error
		t     *Task
		uri   string
	)

	logEn, err = c.GetLogicalEnclosureByName(name)
	if err != nil {
		return err
	}
	if logEn.Name != "" {
		t = t.NewProfileTask(c)
		t.ResetTask()
		log.Debugf("REST : %s \n %+v\n", logEn.URI, logEn)
		log.Debugf("task -> %+v", t)
		uri = logEn.URI.String()
		if uri == "" {
			log.Warn("Unable to post delete, no uri found.")
			t.TaskIsDone = true
			return err
		}
		data, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete logical Enclosure request: %s", err)
			t.TaskIsDone = true
			return err
		}

		log.Debugf("Response delete Logical Enclosure %s", data)
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
		log.Infof("LogicalEnclosure could not be found to delete, %s, skipping delete ...", name)
	}
	return nil
}

func (c *OVClient) UpdateLogicalEnclosure(logEn LogicalEnclosure) error {
	log.Infof("Initializing update of logical enclosure for %s.", logEn.Name)
	var (
		uri = logEn.URI.String()
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	// reset query
	c.SetQueryString(make(map[string]interface{}))

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, logEn)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, logEn)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update logical enclosure request: %s", err)
		return err
	}

	log.Debugf("Response update LogicalEnclosure %s", data)
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

func (c *OVClient) UpdateFromGroupLogicalEnclosure(logEn LogicalEnclosure) error {
	log.Infof("Initializing updateFromGroup of logical enclosure for %s.", logEn.Name)
	var (
		uri = logEn.URI.String() + "/updateFromGroup"
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	// reset query
	c.SetQueryString(make(map[string]interface{}))

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, nil)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, nil)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting updateFromGroup logical enclosure request: %s", err)
		return err
	}

	log.Debugf("Response updateFromGroup LogicalEnclosure %s", data)
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
