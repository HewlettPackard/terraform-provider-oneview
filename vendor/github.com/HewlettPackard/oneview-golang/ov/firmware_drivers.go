package ov

import (
	"encoding/json"
	"strings"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type parentBundle struct {
	ParentBundleName string `json:"parentBundleName,omitempty"`
	ReleaseDate      string `json:"releaseDate,omitempty"`
	Version          string `json:"version,omitempty"`
}

type HotFixes struct {
	HotfixName  string `json:"hotfixName,omitempty"`
	ReleaseDate string `json:"releaseDate,omitempty"`
	ResourceId  string `json:"resourceId,omitempty"`
}

type FWComponents struct {
	ComponentVersion string          `json:"componentVersion,omitempty"`
	FileName         string          `json:"fileName,omitempty"`
	Name             string          `json:"name,omitempty"`
	SwKeyNameList    []utils.Nstring `json:"swKeyNameList,omitempty"`
}

type FirmwareDrivers struct {
	BaselineShortName     string              `json:"baselineShortName,omitempty"`
	BundleSize            int                 `json:"bundleSize,omitempty"`
	BundleType            string              `json:"bundleType,omitempty"`
	Category              string              `json:"category,omitempty"`
	Created               string              `json:"created,omitempty"`
	Description           string              `json:"description,omitempty"`
	ETAG                  string              `json:"eTag,omitempty"`
	EsxiOsDriverMetaData  []utils.Nstring     `json:"esxiOsDriverMetaData,omitempty"`
	FwComponents          []FWComponents      `json:"fwComponents,omitempty"`
	Hotfixes              []HotFixes          `json:"hotfixes,omitempty"`
	HpsumVersion          string              `json:"hpsumVersion,omitempty"`
	IsoFileName           string              `json:"isoFileName,omitempty"`
	LastTaskUri           string              `json:"lastTaskUri,omitempty"`
	Locations             map[string]string   `json:"locations,omitempty"`
	Mirrorlist            map[string][]string `json:"mirrorlist,omitempty"`
	Modified              string              `json:"modified,omitempty"`
	Name                  string              `json:"name,omitempty"`
	ParentBundle          parentBundle        `json:"parentBundle,omitempty"`
	ReleaseDate           string              `json:"releaseDate,omitempty"`
	ResourceId            string              `json:"resourceId,omitempty"`
	ResourceState         string              `json:"resourceState,omitempty"`
	ScopesUri             string              `json:"scopesUri,omitempty"`
	SignatureFileName     string              `json:"signatureFileName,omitempty"`
	SignatureFileRequired bool                `json:"signatureFileRequired,omitempty"`
	State                 string              `json:"state,omitempty"`
	Status                string              `json:"status,omitempty"`
	SupportedLanguages    string              `json:"supportedLanguages,omitempty"`
	SupportedOSList       []utils.Nstring     `json:"supportedOSList,omitempty"`
	SwPackagesFullPath    string              `json:"swPackagesFullPath,omitempty"`
	Type                  string              `json:"type,omitempty"`
	Uri                   utils.Nstring       `json:"uri,omitempty"`
	Uuid                  string              `json:"uuid,omitempty"`
	Version               string              `json:"version,omitempty"`
	XmlKeyName            string              `json:"xmlKeyName,omitempty"`
}

type FirmwareDriversList struct {
	Category    string            `json:"category,omitempty"`
	Count       int               `json:"count,omitempty"`
	Created     string            `json:"created,omitempty"`
	ETAG        string            `json:"eTag,omitempty"`
	Members     []FirmwareDrivers `json:"members,omitempty"`
	Modified    string            `json:"modified,omitempty"`
	NextPageURI utils.Nstring     `json:"nextPageUri,omitempty"`
	PrevPageURI utils.Nstring     `json:"prevPageUri,omitempty"`
	Start       int               `json:"start,omitempty"`
	Total       int               `json:"total,omitempty"`
	Type        string            `json:"type,omitempty"`
	Uri         utils.Nstring     `json:"uri,omitempty"`
}

type CustomServicePack struct {
	BaselineUri        string          `json:"baselineUri,omitempty"`
	CustomBaselineName string          `json:"customBaselineName,omitempty"`
	HotfixUris         []utils.Nstring `json:"hotfixUris,omitempty"`
	InitialScopeUris   []utils.Nstring `json:"initialScopeUris,omitempty"`
}

func (c *OVClient) GetFirmwareBaselineList(sort string, start string, count string) (FirmwareDriversList, error) {
	var (
		uri      = "/rest/firmware-drivers"
		firmware FirmwareDriversList
		q        = make(map[string]interface{})
	)

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
		return firmware, err
	}

	log.Debugf("GetFirmwareBaseline %s", data)
	if err := json.Unmarshal(data, &firmware); err != nil {
		return firmware, err
	}
	return firmware, nil

}

func (c *OVClient) GetFirmwareBaselineById(id string) (FirmwareDrivers, error) {
	var (
		uri        = "/rest/firmware-drivers/" + id
		firmwareId FirmwareDrivers
	)
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return firmwareId, err
	}

	log.Debugf("GetFirmwareBaseline %s", data)
	if err := json.Unmarshal(data, &firmwareId); err != nil {
		return firmwareId, err
	}
	return firmwareId, nil
}

func (c *OVClient) GetFirmwareBaselineByNameandVersion(name, version string) (FirmwareDrivers, error) {

	firmwareList, err := c.GetFirmwareBaselineList("", "", "")

	if firmwareList.Total > 0 {

		for i := range firmwareList.Members {
			if version != "" {
				if strings.EqualFold(strings.TrimSpace(firmwareList.Members[i].Name), strings.TrimSpace(name)) &&
					strings.EqualFold(strings.TrimSpace(firmwareList.Members[i].Version), strings.TrimSpace(version)) {
					return firmwareList.Members[i], err
				}

			} else {
				if strings.EqualFold(strings.TrimSpace(firmwareList.Members[i].Name), strings.TrimSpace(name)) {
					return firmwareList.Members[i], err

				}
			}

		}

	}
	return FirmwareDrivers{}, err
}

func (c *OVClient) CreateCustomServicePack(sp CustomServicePack, force string) error {
	var (
		uri = "/rest/firmware-drivers/"
		t   *Task
	)
	q := make(map[string]interface{})
	if force != "" {
		q["force"] = force
	}
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	if len(q) > 0 {
		c.SetQueryString(q)
	}
	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("task -> %+v", t)

	data, err := c.RestAPICall(rest.POST, uri, sp)
	if err != nil {
		log.Errorf("Error submitting create firmware baseline request: %s", err)
		t.TaskIsDone = true
		return err
	}

	log.Debugf("CreateFirmwareBaseline")
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

func (c *OVClient) DeleteFirmwareBaseline(id string, force string) error {
	var (
		firmware FirmwareDrivers
		err      error
		t        *Task
		uri      = "/rest/firmware-drivers/" + id
	)

	firmware, err = c.GetFirmwareBaselineById(id)
	if err != nil {
		return err
	}
	if firmware.Name != "" {
		q := make(map[string]interface{})
		if force != "" {
			q["force"] = force
		}
		if len(q) > 0 {
			c.SetQueryString(q)
		}
		t = t.NewProfileTask(c)
		t.ResetTask()
		log.Debugf("REST : %s \n %+v\n", firmware.Uri, firmware)
		log.Debugf("task -> %+v", t)
		uri = firmware.Uri.String()
		if uri == "" {
			log.Warn("Unable to post delete, no uri found.")
			t.TaskIsDone = true
			return err
		}
		data, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete firmware baseline request: %s", err)
			t.TaskIsDone = true
			return err
		}

		log.Debugf("Response firmware baseline network %s", data)
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
		log.Infof("Firmware Baseline could not be found to delete, %s, skipping delete ...", id)
	}
	return nil
}
