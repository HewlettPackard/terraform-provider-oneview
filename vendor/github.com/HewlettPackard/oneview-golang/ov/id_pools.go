package ov

import (
	"encoding/json"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type IdPool struct {
	AllocatedCount int           `json:"allocatedCount,omitempty"`
	Category       string        `json:"category,omitempty"`
	Created        string        `json:"created,omitempty"`
	ETAG           utils.Nstring `json:"etag,omitempty"`
	Enabled        *bool         `json:"enabled,omitempty"`
	FreeCount      int           `json:"freeCount,omitempty"`
	Modified       string        `json:"modified,omitempty"`
	Name           string        `json:"name,omitempty"`
	PoolType       string        `json:"poolType,omitempty"`
	Prefix         string        `json:"prefix,omitempty"`
	RangeUris      []string      `json:"rangeUris,omitempty"`
	TotalCount     int           `json:"totalCount,omitempty"`
	Type           string        `json:"type,omitempty"`
	URI            string        `json:"uri,omitempty"`
}

// Checks the range availability in the ID pool.
func (c *OVClient) GetRangeAvailibility(poolType string, ids []string) (UpdateAllocatorList, error) {
	var (
		uri    = "/rest/id-pools/" + poolType + "/checkrangeavailability"
		idList UpdateAllocatorList
	)

	if poolType == "" {
		log.Errorf("Error submitting update request. Please provide a valid Pool Type")
	}

	q := make(map[string]interface{})
	if len(ids) > 0 {
		q["idList"] = ids
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
		return idList, err
	}

	log.Debugf("Get Id Lists %s", data)
	if err := json.Unmarshal(data, &idList); err != nil {
		return idList, err
	}
	return idList, nil
}

// Validates a set of IDs to reserve in the pool.
func (c *OVClient) GetValidateIds(poolType string, ids []string) (UpdateAllocatorList, error) {
	var (
		uri    = "/rest/id-pools/" + poolType + "/validate"
		idList UpdateAllocatorList
	)

	if poolType == "" {
		log.Errorf("Error submitting update request. Please provide a valid Pool Type")
	}

	q := make(map[string]interface{})
	if len(ids) > 0 {
		q["idList"] = ids
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
		return idList, err
	}

	log.Debugf("Get Id Lists %s", data)
	if err := json.Unmarshal(data, &idList); err != nil {
		return idList, err
	}
	return idList, nil
}

// Gets a pool along with the list of ranges present in it
func (c *OVClient) GetPoolType(poolType string) (IdPool, error) {
	var (
		uri    = "/rest/id-pools/"
		idPool IdPool
	)

	if poolType == "" {
		log.Errorf("Error submitting update request. Please provide a valid Pool Type")
	}

	uri = uri + poolType

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return idPool, err
	}

	log.Debugf("GetPoolType %s", data)
	if err := json.Unmarshal(data, &idPool); err != nil {
		return idPool, err
	}
	return idPool, nil
}

// Generates and returns a random range.
// Used to generate a range for validation prior to actually creating it. This API is not applicable for the IPv4 and IPv6 IDs.
func (c *OVClient) Generate(poolType string) (StartStopFragments, error) {
	var (
		uri = "/rest/id-pools/"
		ids StartStopFragments
	)

	if poolType == "" {
		log.Errorf("Error submitting update request. Please provide a valid Pool Type")
	}

	uri = uri + poolType + "/generate"

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return ids, err
	}

	log.Debugf("GenerateIds %s", data)
	if err := json.Unmarshal(data, &ids); err != nil {
		return ids, err
	}
	return ids, nil
}

// Enables or disables the pool
func (c *OVClient) UpdatePoolType(idPool IdPool, poolType string) (IdPool, error) {

	log.Infof("Initializing update of pool type for %s.", poolType)
	var (
		uri      = "/rest/id-pools/"
		t        *Task
		response IdPool
	)

	if poolType == "" {
		log.Errorf("Error submitting update request. Please provide a valid Pool Type")
	}

	uri = uri + poolType

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()

	log.Debugf("REST : %s \n %+v\n", uri, idPool)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, idPool)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update request: %s", err)
		return response, err
	}

	log.Debugf("Response update Id Pool %s", data)

	dataResponse := data
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return response, err
	}
	if err := json.Unmarshal([]byte(dataResponse), &response); err != nil {
		log.Errorf("Error with task un-marshal: %s", err)
		return response, err
	}
	return response, nil
}

// Allocates one or more IDs from a pool.
func (c *OVClient) Allocator(allocateIds UpdateAllocatorList, poolType string) (UpdateAllocatorList, error) {

	log.Infof("Initializing update of pool type for %s.", poolType)
	var (
		uri      = "/rest/id-pools/"
		t        *Task
		response UpdateAllocatorList
	)

	if poolType == "" {
		log.Errorf("Error submitting update request. Please provide a valid Pool Type")
	}

	uri = uri + poolType + "/allocator"

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()

	log.Debugf("REST : %s \n %+v\n", uri, allocateIds)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, allocateIds)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update request: %s", err)
		return response, err
	}

	log.Debugf("Response allocate Ids %s", data)

	if err := json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return response, err
	}
	if err := json.Unmarshal(data, &response); err != nil {
		log.Errorf("Error with task un-marshal: %s", err)
		return response, err
	}

	return response, nil
}

// Collects one or more IDs to be returned to a pool.
func (c *OVClient) Collector(idList UpdateCollectorList, poolType string) (UpdateCollectorList, error) {

	log.Infof("Initializing update of pool type for %s.", poolType)
	var (
		uri      = "/rest/id-pools/"
		t        *Task
		response UpdateCollectorList
	)

	if poolType == "" {
		log.Errorf("Error submitting update request. Please provide a valid Pool Type")
	}

	uri = uri + poolType + "/collector"

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()

	log.Debugf("REST : %s \n %+v\n", uri, idList)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, idList)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update request: %s", err)
		return response, err
	}

	log.Debugf("Response Collector Ids %s", data)

	if err := json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return response, err
	}
	if err := json.Unmarshal(data, &response); err != nil {
		log.Errorf("Error with task un-marshal: %s", err)
		return response, err
	}

	return response, nil
}

func (c *OVClient) UpdateValidateIds(Ids UpdateAllocatorList, poolType string) (UpdateAllocatorList, error) {

	log.Infof("Initializing update of pool type for %s.", poolType)
	var (
		uri      = "/rest/id-pools/"
		t        *Task
		response UpdateAllocatorList
	)

	if poolType == "" {
		log.Errorf("Error submitting update request. Please provide a valid Pool Type")
	}

	uri = uri + poolType + "/validate"

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()

	log.Debugf("REST : %s \n %+v\n", uri, Ids)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, Ids)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update request: %s", err)
		return response, err
	}

	log.Debugf("Response Ids %s", data)

	if err := json.Unmarshal([]byte(data), &t); err != nil {
		t.TaskIsDone = true
		log.Errorf("Error with task un-marshal: %s", err)
		return response, err
	}
	if err := json.Unmarshal(data, &response); err != nil {
		log.Errorf("Error with task un-marshal: %s", err)
		return response, err
	}

	return response, nil
}
