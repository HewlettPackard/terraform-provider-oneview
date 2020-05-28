package ov

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
	"strconv"
)

type HypervisorClusterProfile struct {
	AddHostRequests               []string                       `json:"addHostRequests,omitempty"`               //"addHostRequests":"[]"
	Category                      string                         `json:"category,omitempty"`                      //"category":"hypervisor-cluster-profiles"
	ComplianceState               string                         `json:"complianceState,omitempty"`               //"complianceState":"Consistent"
	Created                       string                         `json:"created,omitempty"`                       //"created":"2020-04-13T16:28:44.234Z"
	Description                   utils.Nstring                  `json:"description,omitempty"`                   //"description":""
	ETag                          string                         `json:"eTag,omitempty"`                          //"eTag":"1586795326281/1586795326281"
	HypervisorClusterSettings     *HypervisorClusterSettings     `json:"hypervisorClusterSettings,omitempty"`     //"hypervisorClusterSettings":""
	HypervisorClusterUri          string                         `json:"hypervisorClusterUri,omitempty"`          //"hypervisorClusterUri":"/rest/hypervisor-clusters/a2c4c63e-f96e-4dc1-976d-12a677ba5306"
	HypervisorHostProfileTemplate *HypervisorHostProfileTemplate `json:"hypervisorHostProfileTemplate,omitempty"` //"hypervisorHostProfileTemplate":""
	HypervisorHostProfileUris     utils.Nstring                  `json:"hypervisorHostProfileUris,omitempty"`     //"hypervisorHostProfileUris":"null"
	HypervisorManagerUri          utils.Nstring                  `json:"hypervisorManagerUri,omitempty"`          //"hypervisorManagerUri":"/rest/hypervisor-managers/1ded903a-ac66-41cf-ba57-1b9ded9359b6"
	HypervisorType                string                         `json:"hypervisorType,omitempty"`                //"hypervisorType":"Vmware"
	IpPools                       []utils.Nstring                `json:"ipPools,omitempty"`                       //"ipPools":"[]"
	MgmtIpSettingsOverride        string                         `json:"mgmtIpSettingsOverride,omitempty"`        //"mgmtIpSettingsOverride":"null"
	Modified                      string                         `json:"modified,omitempty"`                      //"modified":"2020-04-13T16:28:46.281Z"
	Name                          string                         `json:"name,omitempty"`                          //"name":"HCP
	Path                          string                         `json:"path,omitempty"`                          //"path":"DC1"
	RefreshState                  string                         `json:"refreshState,omitempty"`                  //"refreshState":"NotRefreshing"
	ScopesUri                     string                         `json:"scopesUri,omitempty"`                     //"scopesUri":"/rest/scopes/resources/rest/hypervisor-cluster-profiles/4340293c-0701-4773"
	SharedStorageVolumes          []utils.Nstring                `json:"sharedStorageVolumes,omitempty"`          //"sharedStorageVolumes":"[]"
	State                         string                         `json:"state,omitempty"`                         //"state":"Active"
	StateReason                   string                         `json:"stateReason,omitempty"`                   //"stateReason":"None"
	Status                        string                         `json:"status,omitempty"`                        //"status":"OK"
	Type                          string                         `json:"type,omitempty"`                          //"type":"HypervisorClusterProfileV3"
	URI                           utils.Nstring                  `json:"uri,omitempty"`                           //"uri":"/rest/hypervisor-cluster-profiles/4340293c-0701-4773-b863-32854b0f7d29"
}

type HypervisorClusterSettings struct {
	DistributedSwitchUsage   string `json:"distributedSwitchUsage,omitempty"`   //"distributedSwitchUsage":"null"
	DistributedSwitchVersion string `json:"distributedSwitchVersion,omitempty"` //"distributedSwitchVersion":"null"
	DrsEnabled               bool   `json:"drsEnabled,omitempty"`               //"drsEnabled":"true"
	HaEnabled                bool   `json:"haEnabled,omitempty"`                //"haEnabled":"false"
	MultiNicVMotion          bool   `json:"multiNicVMotion"`                    //"multiNicVMotion":"false"
	Type                     string `json:"type,omitempty"`                     //"type":"Vmware"
	VirtualSwitchType        string `json:"virtualSwitchType,omitempty"`        //"virtualSwitchType":"Standard"
}

type HypervisorHostProfileTemplate struct {
	DeploymentManagerType     string                     `json:"deploymentManagerType,omitempty"`     //"deploymentManagerType":"I3S"
	DeploymentPlan            *DeploymentPlan            `json:"deploymentPlan,omitempty"`            //"deploymentPlan":""
	HostConfigPolicy          *HostConfigPolicy          `json:"hostConfigPolicy,omitempty"`          //"hostConfigPolicy":""
	Hostprefix                string                     `json:"hostprefix,omitempty"`                //"hostprefix":"HCP"
	ServerProfileTemplateUri  utils.Nstring              `json:"serverProfileTemplateUri,omitempty"`  //"serverProfileTemplateUri":"/rest/server-profile-templates/278cadfb-2e86-4a05-8932-972553518259"
	VirtualSwitchConfigPolicy *VirtualSwitchConfigPolicy `json:"virtualSwitchConfigPolicy,omitempty"` //"virtualSwitchConfigPolicy":""
	VirtualSwitches           []VirtualSwitches          `json:"virtualSwitches,omitempty"`           //"virtualSwitches":""
}

type DeploymentPlan struct {
	DeploymentCustomArgs      []utils.Nstring `json:"deploymentCustomArgs,omitempty"`      //"deploymentCustomArgs":"[]"
	DeploymentPlanDescription string          `json:"deploymentPlanDescription,omitempty"` //"deploymentPlanDescription":"null"
	DeploymentPlanUri         utils.Nstring   `json:"deploymentPlanUri,omitempty"`         //"deploymentPlanUri":"null"
	Name                      string          `json:"name,omitempty"`                      //"name":"null"
	ServerPassword            string          `json:"serverPassword,omitempty"`            //"serverPassword":"null"
}

type HostConfigPolicy struct {
	LeaveHostInMaintenance  bool `json:"leaveHostInMaintenance,omitempty"`  //"leaveHostInMaintenance":"false"
	UseHostPrefixAsHostname bool `json:"useHostPrefixAsHostname,omitempty"` //"useHostPrefixAsHostname":"false"
	UseHostnameToRegister   bool `json:"useHostnameToRegister,omitempty"`   //"useHostnameToRegister":"false"
}
type VirtualSwitchConfigPolicy struct {
	ConfigurePortGroups   bool `json:"configurePortGroups,omitempty"`   //"configurePortGroups":"true"
	CustomVirtualSwitches bool `json:"customVirtualSwitches,omitempty"` //"customVirtualSwitches":"false"
	ManageVirtualSwitches bool `json:"manageVirtualSwitches,omitempty"` //"manageVirtualSwitches":"true"
}
type VirtualSwitches struct {
	Action                  string                    `json:"action,omitempty"`                  //"action":"NONE"
	Name                    string                    `json:"name,omitempty"`                    //"name":"mgmt"
	NetworkUris             []utils.Nstring           `json:"networkUris,omitempty"`             //"networkUris":""
	Version                 string                    `json:"version,omitempty"`                 //"version":"null"
	VirtualSwitchPortGroups []VirtualSwitchPortGroups `json:"virtualSwitchPortGroups,omitempty"` //"virtualSwitchPortGroups":""
	VirtualSwitchType       string                    `json:"virtualSwitchType,omitempty"`       //"virtualSwitchType":"Standard"
	VirtualSwitchUplinks    []VirtualSwitchUplinks    `json:"virtualSwitchUplinks,omitempty"`    //"virtualSwitchUplinks":""
}
type VirtualSwitchPortGroups struct {
	Action             string               `json:"action,omitempty"`             //"action":"NONE"
	Name               string               `json:"name,omitempty"`               //"name":"mgmt"
	NetworkUris        []utils.Nstring      `json:"networkUris,omitempty"`        //"networkUris":""
	VirtualSwitchPorts []VirtualSwitchPorts `json:"virtualSwitchPorts,omitempty"` //"virtualSwitchPorts":""
	Vlan               string               `json:"vlan,omitempty"`               //"vlan":"0"
}
type VirtualSwitchPorts struct {
	Action             string          `json:"action,omitempty"`             //"action":"NONE"
	Dhcp               bool            `json:"dhcp,omitempty"`               //"dhcp":"false"
	IpAddress          string          `json:"ipAddress,omitempty"`          //"ipAddress":"null"
	SubnetMask         string          `json:"subnetMask,omitempty"`         //"subnetMask":"null"
	VirtualPortPurpose []utils.Nstring `json:"virtualPortPurpose,omitempty"` //"virtualPortPurpose":""

}
type VirtualSwitchUplinks struct {
	Action string `json:"action,omitempty"` //"action":"NONE"
	Active bool   `json:"active,omitempty"` //"active":"false"
	Mac    string `json:"mac,omitempty"`    //"mac":"null"
	Name   string `json:"name,omitempty"`   //"name":"Mezz 3:1-c"
	Vmnic  string `json:"vmnic,omitempty"`  //"vmnic":"null"
}
type HypervisorClusterProfileList struct {
	Total       int                        `json:"total,omitempty"`       // "total": 1,
	Count       int                        `json:"count,omitempty"`       // "count": 1,
	Start       int                        `json:"start,omitempty"`       // "start": 0,
	PrevPageURI utils.Nstring              `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring              `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         utils.Nstring              `json:"uri,omitempty"`         // "uri": "/rest/interconnects?start=2&count=2",
	Members     []HypervisorClusterProfile `json:"members,omitempty"`     // "members":[]
}

type HypervisorClusterProfileCompliancePreview struct {
	ClusterComplianceDetails               *ClusterComplianceDetails                `json:clusterComplianceDetails,omitempty"`                //"ClusterComplianceDetails":" "
	HypervisorHostProfileComplianceDetails []HypervisorHostProfileComplianceDetails `"json:hypervisorHostProfileComplianceDetails,omitempty"` //"hypervisorHostProfileComplianceDetails":""
}

type ClusterComplianceDetails struct {
	AutomaticUpdates []string `json:"automaticUpdates,omitempty"` //"automaticUpdates":""
	ManualUpdates    []string `json:"manualUpdates,omitempty"`    //"manualUpdates":""
}
type HypervisorHostProfileComplianceDetails struct {
	HostProfileName                    string                              `json:"hostProfileName,omitempty"`                    //"hostProfileName":""
	HostProfileUri                     string                              `json:"hostProfileUri,omitempty"`                     //"hostProfileUri":""
	HypervisorProfileComplianceDetails *HypervisorProfileComplianceDetails `json:"hypervisorProfileComplianceDetails,omitempty"` //"hypervisorProfileComplianceDetail":
	IsOnlineUpdate                     bool                                `json:"isOnlineUpdate,omitempty"`                     //"isOnlineUpdat":""
	ServerProfileComplianceDetails     *ServerProfileComplianceDetails     `json:"serverProfileComplianceDetails,omitempty"`     //"serverProfileComplianceDetail":
}
type HypervisorProfileComplianceDetails struct {
	AutomaticUpdates []string `json:"automaticUpdates,omitempty"` //"automaticUpdates":""
	ManualUpdates    []string `json:"manualUpdates,omitempty"`    //"manualUpdates":""

}
type ServerProfileComplianceDetails struct {
	AutomaticUpdates []string `json:"automaticUpdates,omitempty"` //"automaticUpdates":""
	ManualUpdates    []string `json:"manualUpdates,omitempty"`    //"manualUpdates":""

}

type VirtualSwitchLayout struct {
	ServerProfileTemplateUri utils.Nstring `json:"serverProfileTemplateUri"` //"ServerProfileTemplateUri":""
	HypervisorManagerUri     utils.Nstring `json:"hypervisorManagerUri"`     //"HypervisorManagerUri":""
}

func (c *OVClient) GetHypervisorClusterProfileById(id string) (HypervisorClusterProfile, error) {
	var (
		uri                      = "/rest/hypervisor-cluster-profiles/"
		hypervisorclusterprofile HypervisorClusterProfile
	)

	uri = uri + id
	hypervisorclusterprofile, err := c.GetHypervisorClusterProfileByUri(uri)

	return hypervisorclusterprofile, err
}
func (c *OVClient) GetHypervisorClusterProfileByName(name string) (HypervisorClusterProfile, error) {
	var (
		hypervisorclusterprofile HypervisorClusterProfile
	)
	hypervisorclusterprofiles, err := c.GetHypervisorClusterProfiles("", "", fmt.Sprintf("name matches '%s'", name), "name:asc")
	if hypervisorclusterprofiles.Total > 0 {
		return hypervisorclusterprofiles.Members[0], err
	} else {
		return hypervisorclusterprofile, err
	}
}
func (c *OVClient) GetHypervisorClusterProfileByUri(uri string) (HypervisorClusterProfile, error) {
	var (
		hypervisorClusterProfile HypervisorClusterProfile
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return hypervisorClusterProfile, err
	}
	log.Debugf("GetHypervisorClusterProfile %s", data)
	if err := json.Unmarshal([]byte(data), &hypervisorClusterProfile); err != nil {
		return hypervisorClusterProfile, err
	}
	return hypervisorClusterProfile, nil
}

func (c *OVClient) GetHypervisorClusterProfiles(start string, count string, filter string, sort string) (HypervisorClusterProfileList, error) {
	var (
		uri                       = "/rest/hypervisor-cluster-profiles"
		q                         map[string]interface{}
		hypervisorClusterProfiles HypervisorClusterProfileList
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
	if len(q) > 1 {
		c.SetQueryString(q)
	}

	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return hypervisorClusterProfiles, err
	}

	log.Debugf("GetHypervisorClusterProfiles %s", data)
	if err := json.Unmarshal([]byte(data), &hypervisorClusterProfiles); err != nil {
		return hypervisorClusterProfiles, err
	}
	return hypervisorClusterProfiles, nil
}

func (c *OVClient) GetHypervisorClusterProfileCompliancePreview(id string) (HypervisorClusterProfileCompliancePreview, error) {
	var (
		uri                                       = "/rest/hypervisor-cluster-profiles/"
		hypervisorClusterProfileCompliancePreview HypervisorClusterProfileCompliancePreview
	)

	uri = uri + id + "/compliance-preview"

	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return hypervisorClusterProfileCompliancePreview, err
	}
	log.Debugf("GetHypervisorClusterProfileCompliancePreview %s", data)
	if err := json.Unmarshal([]byte(data), &hypervisorClusterProfileCompliancePreview); err != nil {
		return hypervisorClusterProfileCompliancePreview, err
	}
	return hypervisorClusterProfileCompliancePreview, nil
}

func (c *OVClient) CreateHypervisorClusterProfile(hyClustProf HypervisorClusterProfile) error {
	log.Infof("Initializing creation of hypervisor cluster profile for %s.", hyClustProf.Name)
	var (
		uri = "/rest/hypervisor-cluster-profiles"
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, hyClustProf)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, hyClustProf)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new hypervisor cluster profile request: %s", err)
		return err
	}

	log.Debugf("Response New HypervisorClusterProfile %s", data)
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

func (c *OVClient) CreateVirtualSwitchLayout(virtualswitchlayout VirtualSwitchLayout) error {
	var (
		uri = "/rest/hypervisor-cluster-profiles/virtualswitch-layout"
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	log.Debugf("REST : %s \n %+v\n", uri, virtualswitchlayout)
	data, err := c.RestAPICall(rest.POST, uri, virtualswitchlayout)
	if err != nil {
		log.Errorf("Error submitting virtual switch layout request: %s", err)
		return err
	} else {
		log.Infof("Virtual switch layout creation successfule %s", data)
	}

	return nil
}

func (c *OVClient) DeleteHypervisorClusterProfile(name string) error {
	var (
		hyClustProf HypervisorClusterProfile
		err         error
		t           *Task
		uri         string
	)

	hyClustProf, err = c.GetHypervisorClusterProfileByName(name)
	if err != nil {
		return err
	}
	if hyClustProf.Name != "" {
		t = t.NewProfileTask(c)
		t.ResetTask()
		log.Debugf("REST : %s \n %+v\n", hyClustProf.URI, hyClustProf)
		log.Debugf("task -> %+v", t)
		uri = hyClustProf.URI.String()
		if uri == "" {
			log.Warn("Unable to post delete, no uri found.")
			t.TaskIsDone = true
			return err
		}
		data, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete hypervisor cluster profile request: %s", err)
			t.TaskIsDone = true
			return err
		}

		log.Debugf("Response delete hypervisor cluster profile %s", data)
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
		log.Infof("HypervisorClusterProfile could not be found to delete, %s, skipping delete ...", name)
	}
	return nil
}

func (c *OVClient) DeleteHypervisorClusterProfileSoftDelete(name string, soft_delete bool) error {
	err_softdelete := c.DeleteHypervisorClusterProfileSoftDeleteForce(name, soft_delete, false)
	return err_softdelete
}

func (c *OVClient) DeleteHypervisorClusterProfileSoftDeleteForce(name string, soft_delete bool, force bool) error {
	var (
		hyClustProf HypervisorClusterProfile
		err         error
		t           *Task
		uri         string
		q           map[string]interface{}
	)

	q = make(map[string]interface{})
	q["softDelete"] = strconv.FormatBool(soft_delete)
	q["force"] = strconv.FormatBool(force)

	hyClustProf, err = c.GetHypervisorClusterProfileByName(name)
	if err != nil {
		return err
	}
	if hyClustProf.Name != "" {
		t = t.NewProfileTask(c)
		t.ResetTask()
		log.Debugf("REST : %s \n %+v\n", hyClustProf.URI, hyClustProf)
		log.Debugf("task -> %+v", t)
		uri = hyClustProf.URI.String()
		if uri == "" {
			log.Warn("Unable to post delete, no uri found.")
			t.TaskIsDone = true
			return err
		}

		// refresh login
		c.RefreshLogin()
		c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
		// Setup query
		c.SetQueryString(q)

		data, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete hypervisor cluster profile request: %s", err)
			t.TaskIsDone = true
			return err
		}

		log.Debugf("Response delete hypervisor cluster profile %s", data)
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
		log.Infof("HypervisorClusterProfile could not be found to delete, %s, skipping delete ...", name)
	}
	return nil
}

func (c *OVClient) UpdateHypervisorClusterProfile(hyClustProf HypervisorClusterProfile) error {
	log.Infof("Initializing update of hypervisor cluster profile for %s.", hyClustProf.Name)
	var (
		uri = hyClustProf.URI.String()
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, hyClustProf)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, hyClustProf)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update hypervisor cluster profile request: %s", err)
		return err
	}

	log.Debugf("Response update HypervisorClusterProfile %s", data)
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
