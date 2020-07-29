package ov

import (
	"encoding/json"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type Ipv4Range struct {
	AllocatedFragmentUri utils.Nstring         `json:"allocatedFragmentUri,omitempty"`
	AllocatedIdCount     int                   `json:"allocatedIdCount,omitempty"`
	AllocatorUri         utils.Nstring         `json:"allocatorUri,omitempty"`
	AssociatedResources  []AssociatedResources `json:"associatedResources,omitempty"`
	Category             string                `json:"category,omitempty"`
	CollectorUri         utils.Nstring         `json:"collectorUri"`
	Created              string                `json:"created,omitempty"`
	DefaultRange         bool                  `json:"defaultRange"`
	ETAG                 string                `json:"eTag,omitempty"`
	Modified             string                `json:"modified,omitempty"`
	Enabled              bool                  `json:"enabled,omitempty"`
	Name                 string                `json:"name,omitempty"`
	EndAddress           utils.Nstring         `json:"endAddress,omitempty"`
	FreeFragmentUri      utils.Nstring         `json:"freeFragmentUri,omitempty"`
	URI                  utils.Nstring         `json:"uri,omitempty"`
	Prefix               utils.Nstring         `json:"prefix,omitempty"`
	RangeCategory        utils.Nstring         `json:"rangeCategory,omitempty"`
	ReservedIdCount      int                   `json:"reservedIdCount,omitempty"`
	StartAddress         utils.Nstring         `json:"startAddress,omitempty"`
	StartStopFragments   []StartStopFragments  `json:"startStopFragments,omitempty"`
	SubnetUri            utils.Nstring         `json:"subnetUri,omitempty"`
	TotalCount           int                   `json:"totalCount,omitempty"`
	Type                 string                `json:"type,omitempty"`
}

type CreateIpv4Range struct {
	Name               string               `json:"name,omitempty"`
	StartStopFragments []StartStopFragments `json:"startStopFragments,omitempty"`
	SubnetUri          utils.Nstring        `json:"subnetUri,omitempty"`
	Type               string               `json:"type,omitempty"`
}

type AssociatedResources struct {
	AssociationType  string        `json:"associationType,omitempty"`
	ResourceCategory string        `json:"resourceCategory,omitempty"`
	ResourceName     string        `json:"resourceName,omitempty"`
	ResourceUri      utils.Nstring `json:"resourceUri,omitempty"`
}

type StartStopFragments struct {
	StartAddress utils.Nstring `json:"startAddress,omitempty"`
	EndAddress   utils.Nstring `json:"endAddress,omitempty"`
	FragmentType string        `json:"fragmentType,omitempty"`
}

type FragmentsList struct {
	Category    string               `json:"category,omitempty"`
	Count       int                  `json:"count,omitempty"`
	ETAG        string               `json:"eTag,omitempty"`
	Created     string               `json:"created,omitempty"`
	Modified    string               `json:"modified,omitempty"`
	Total       int                  `json:"total,omitempty"`
	Start       int                  `json:"start,omitempty"`
	PrevPageURI utils.Nstring        `json:"prevPageUri,omitempty"`
	NextPageURI utils.Nstring        `json:"nextPageUri,omitempty"`
	URI         utils.Nstring        `json:"uri,omitempty"`
	Members     []StartStopFragments `json:"members,omitempty"`
}

type UpdateAllocatorList struct {
	Count  int             `json:"count,omitempty"`
	ETAG   string          `json:"eTag,omitempty"`
	Valid  bool            `json:"valid,omitempty"`
	IdList []utils.Nstring `json:"idList,omitempty"`
}

type UpdateCollectorList struct {
	ETAG   string          `json:"eTag,omitempty"`
	IdList []utils.Nstring `json:"idList,omitempty"`
}

type UpdateIpv4 struct {
	Enabled bool   `json:"enabled,omitempty"`
	Type    string `json:"type,omitempty"`
}

func (c *OVClient) GetIPv4RangebyId(id string) (Ipv4Range, error) {
	var (
		uri       = "/rest/id-pools/ipv4/ranges/"
		ipv4Range Ipv4Range
	)

	uri = uri + id
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return ipv4Range, err
	}

	log.Debugf("GetIpv4Ranges %s", data)
	if err := json.Unmarshal([]byte(data), &ipv4Range); err != nil {
		return ipv4Range, err
	}
	return ipv4Range, nil
}

func (c *OVClient) GetAllocatedFragments(filter string, sort string, start string, count string, id string) (FragmentsList, error) {
	var (
		uri                = "/rest/id-pools/ipv4/ranges/" + id + "/allocated-fragments"
		q                  = make(map[string]interface{})
		allocatedFragments FragmentsList
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
		return allocatedFragments, err
	}

	log.Debugf("GetallocatedFragments %s", data)
	if err := json.Unmarshal(data, &allocatedFragments); err != nil {
		return allocatedFragments, err
	}
	return allocatedFragments, nil
}

func (c *OVClient) GetFreeFragments(filter string, sort string, start string, count string, id string) (FragmentsList, error) {
	var (
		uri           = "/rest/id-pools/ipv4/ranges/" + id + "/free-fragments"
		q             = make(map[string]interface{})
		freeFragments FragmentsList
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
		return freeFragments, err
	}

	log.Debugf("GetfreeFragments %s", data)
	if err := json.Unmarshal(data, &freeFragments); err != nil {
		return freeFragments, err
	}
	return freeFragments, nil
}

func (c *OVClient) CreateIPv4Range(ipv4 CreateIpv4Range) error {
	log.Infof("Initializing creation of ipv4Range for %s.", ipv4.Name)
	var (
		uri = "/rest/id-pools/ipv4/ranges/"
		t   = (&Task{}).NewProfileTask(c)
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, ipv4)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, ipv4)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new ipv4Range creation request: %s", err)
		return err
	}

	log.Debugf("Response New ipv4Range %s", data)
	if err := json.Unmarshal(data, &t); err != nil {
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

func (c *OVClient) DeleteIpv4Range(id string) error {
	var (
		ipv4 Ipv4Range
		err  error
		t    *Task
		uri  string
	)

	ipv4, err = c.GetIPv4RangebyId(id)
	if err != nil {
		return err
	}
	if ipv4.Name != "" {
		t = t.NewProfileTask(c)
		t.ResetTask()
		log.Debugf("REST : %s \n %+v\n", ipv4.URI, ipv4)
		log.Debugf("task -> %+v", t)
		uri = ipv4.URI.String()
		if uri == "" {
			log.Warn("Unable to post delete, no uri found.")
			t.TaskIsDone = true
			return err
		}
		data, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting new ipv4 delete request: %s", err)
			t.TaskIsDone = true
			return err
		}

		log.Debugf("Response delete ipv4 Range %s", data)
		if err := json.Unmarshal(data, &t); err != nil {
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
		log.Infof("ipv4 Range could not be found to delete, %s, skipping delete ...", ipv4.Name)
	}
	return nil
}

func (c *OVClient) UpdateIpv4Range(id string, ipv4 UpdateIpv4) error {
	log.Infof("Initializing update of ipv4 Range")
	var (
		uri = "/rest/id-pools/ipv4/ranges/" + id
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()

	log.Debugf("REST : %s \n %+v\n", uri, ipv4)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, ipv4)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update ipv4 Range request: %s", err)
		return err
	}

	log.Debugf("Response update ipv4 Range %s", data)
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return err
	}

	return nil
}
