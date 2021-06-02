package ov

import (
	"encoding/json"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type ApplianceSshAccess struct {
	AllowSshAccess bool          `json:"allowSshAccess"`
	Category       string        `json:"category,omitempty"`
	Created        string        `json:"created,omitempty"`
	ETAG           string        `json:"eTag,omitempty"`
	Modified       string        `json:"modified,omitempty"`
	Type           string        `json:"type,omitempty"`
	URI            utils.Nstring `json:"uri,omitempty"`
}

func (c *OVClient) GetSshAccess() (ApplianceSshAccess, error) {
	var (
		uri          = "/rest/appliance/ssh-access"
		getsshaccess ApplianceSshAccess
	)

	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return getsshaccess, err
	}

	log.Debugf("GetSSHAccess %s", data)
	if err := json.Unmarshal(data, &getsshaccess); err != nil {
		return getsshaccess, err
	}
	return getsshaccess, nil
}

func (c *OVClient) SetSshAccess(sshaccess ApplianceSshAccess) error {
	log.Infof("Initializing setting of appliance SSH access.")
	var (
		uri = "/rest/appliance/ssh-access"
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, sshaccess)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, sshaccess)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting set appliance ssh access request: %s", err)
		return err
	}

	log.Infof("Response set timelocalework %s", data)
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
