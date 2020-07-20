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

type StoragePool struct {
	Category                 string                               `json:"category,omitempty"`
	Created                  string                               `json:"created,omitempty"`
	Description              utils.Nstring                        `json:"description"`
	ETAG                     string                               `json:"eTag,omitempty"`
	Name                     string                               `json:"name,omitempty"`
	State                    string                               `json:"state,omitempty"`
	Status                   string                               `json:"status,omitempty"`
	Type                     string                               `json:"type,omitempty"`
	URI                      utils.Nstring                        `json:"uri,omitempty"`
	AllocatedCapacity        string                               `json:"allocatedCapacity,omitempty"`
	InitialScopeUris         utils.Nstring                        `json:"initialScopeUris,omitempty"`
	DeviceSpecificAttributes *DeviceSpecificAttributesStoragePool `json:"deviceSpecificAttributes,omitempty"`
	StorageSystemUri         utils.Nstring                        `json:"storageSystemUri,omitempty"`
	TotalCapacity            string                               `json:"totalCapacity,omitempty"`
	FreeCapacity             string                               `json:"freeCapacity,omitempty"`
	IsManaged                bool                                 `json:"isManaged"`
}

type DeviceSpecificAttributesStoragePool struct {
	DeviceID               string                `json:"deviceId,omitempty"`
	Folders                []Folders             `json:"folders,omitempty"`
	IsDeduplicationCapable bool                  `json:"isDeduplicationCapable,omitempty"`
	AllocatedCapacity      *AllocatedCapacity    `json:"allocatedCapacity,omitempty"`
	CapacityLimit          string                `json:"capacityLimit,omitempty"`
	CapacityWarningLimit   string                `json:"capacityWarningLimit,omitempty"`
	DeviceSpeed            string                `json:"deviceSpeed,omitempty"`
	Domain                 string                `json:"domain,omitempty"`
	SupportedRaidLevel     string                `json:"supportedRaidLevel,omitempty"`
	Uuid                   string                `json:"uuid,omitempty"`
	VolumeCreationSpace    []VolumeCreationSpace `json:"volumeCreationSpace,omitempty"`
}

type Folders struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type AllocatedCapacity struct {
	SnapshotAllocatedCapacity string `json:"snapshotAllocatedCapacity,omitempty"`
	TotalAllocatedCapacity    string `json:"totalAllocatedCapacity,omitempty"`
	SharedAllocatedCapacity   string `json:"sharedAllocatedCapacity,omitempty"`
	VolumeAllocatedCapacity   string `json:"volumeAllocatedCapacity,omitempty"`
}

type VolumeCreationSpace struct {
	AvailableSpace   int `json:"availableSpace,omitempty"`
	ReplicationLevel int `json:"replicationLevel,omitempty"`
}

type StoragePoolsList struct {
	Total       int           `json:"total,omitempty"`       // "total": 1,
	Count       int           `json:"count,omitempty"`       // "count": 1,
	Start       int           `json:"start,omitempty"`       // "start": 0,
	PrevPageURI utils.Nstring `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         utils.Nstring `json:"uri,omitempty"`         // "uri": "/rest/storage-pools?filter=connectionTemplateUri%20matches%7769cae0-b680-435b-9b87-9b864c81657fsort=name:asc"
	Members     []StoragePool `json:"members,omitempty"`     // "members":[]
}

func (c *OVClient) GetStoragePoolByName(name string) (StoragePool, error) {
	var (
		sPool StoragePool
	)
	sPools, err := c.GetStoragePools(fmt.Sprintf("name matches '%s'", name), "name:asc", "", "")
	if sPools.Total > 0 {
		return sPools.Members[0], err
	} else {
		return sPool, err
	}
}

func (c *OVClient) GetStoragePools(filter string, sort string, start string, count string) (StoragePoolsList, error) {
	var (
		uri    = "/rest/storage-pools"
		q      map[string]interface{}
		sPools StoragePoolsList
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
		return sPools, err
	}

	log.Debugf("GetStoragePools %s", data)
	if err := json.Unmarshal([]byte(data), &sPools); err != nil {
		return sPools, err
	}
	return sPools, nil
}

func (c *OVClient) UpdateStoragePool(sPool StoragePool) error {
	log.Infof("Initializing update of storage volume for %s.", sPool.Name)
	var (
		uri = sPool.URI.String()
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, sPool)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, sPool)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update StoragePool request: %s", err)
		return err
	}
	log.Debugf("Response update StoragePool %s", data)
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
