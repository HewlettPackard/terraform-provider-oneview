package ov

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
	"strconv"
)

type VirtualPort struct {
	Capabilities []string `json:"capabilities,omitempty"` // "capabilities": {}
	PortFunction string   `json:"portFunction,omitempty"` // "portFunction": "a"
	PortNumber   int      `json:"portNumber,omitempty"`   // "portNumber": 2
}
type SHTPort struct {
	Mapping               int           `json:"mapping,omitempty"`               // "mapping": 3
	MaxSpeedMbps          int           `json:"maxSpeedMbps,omitempty"`          // "maxSpeedMbps":3553
	MaxVFsSupported       int           `json:"maxVFsSupported,omitempty"`       // "maxVFsSupported": 23
	Number                int           `json:"number,omitempty"`                // "number": 1
	PhysicalFunctionCount int           `json:"physicalFunctionCount,omitempty"` // "physicalFunctionCount": 4
	SupportedFcGbps       []string      `json:"supportedFcGbps,omitempty"`       // "supportedFcGbps": ""
	Type                  string        `json:"type,omitempty"`                  // "type": "Ethernet"
	VirtualPort           []VirtualPort `json:"virtualPorts,omitempty"`          // "virtualPorts": {}
}
type Adapter struct {
	Capabilities        []string           `json:"capabilities,omitempty"`        // "capabilities": {}
	DeviceNumber        int                `json:"deviceNumber,omitempty"`        // "deviceNumber": 123
	DeviceType          string             `json:"deviceType,omitempty"`          // "deviceType":"Ethernet"
	Location            string             `json:"location,omitempty"`            // "location": "Mezz"
	MaxVFsSupported     int                `json:"maxVFsSupported,omitempty"`     // "maxVFsSupported": 256
	MinVFsIncrement     int                `json:"minVFsIncrement,omitempty"`     // "minVFsIncrement": 8
	Model               string             `json:"model,omitempty"`               // "model": ""
	Ports               []SHTPort          `json:"ports,omitempty"`               // "ports": {}
	Slot                int                `json:"slot,omitempty"`                // "slot": 3
	StorageCapabilities *StorageCapability `json:"storageCapabilities,omitempty"` // "storageCapabilities":{..,},
}

type Dependency struct {
	Type string `json:"type,omitempty"` // "type": "Map"
}

type OptionLink struct {
	Action    string `json:"action,omitempty"`
	OptionID  string `json:"optionId,omitempty"`
	SettingID string `json:"settingId,omitempty"`
}

type Option struct {
	ID          string       `json:"id,omitempty"`          //"id":"sc3d"
	Name        string       `json:"name,omitempty"`        // "name":"COM"
	OptionLinks []OptionLink `json:"optionLinks,omitempty"` // "optionLinks":{}
}
type BiosSetting struct {
	Category        string       `json:"category,omitempty"`        // "category":""
	DefaultValue    string       `json:"defaultValue,omitempty"`    // "defaultValue":"Cpm21Trq"
	DependencyList  []Dependency `json:"dependencyList,omitempty"`  //"dependencyList:" {}
	HelpText        string       `json:"helpText,omitempty"`        // "helpText": ""
	ID              string       `json:"id,omitempty"`              // "id":""
	LowerBound      int          `json:"lowerBound,omitempty"`      // "lowerBound": 3
	Name            string       `json:"name,omitempty"`            // "name": ""
	Options         []Option     `json:"options,omitempty"`         // "options":{}
	ScalarIncrement int          `json:"scalarIncrement,omitempty"` // "scalarIncrement":4
	StringMaxLength int          `json:"stringMaxLength,omitempty"` // "stringMaxLength":24
	StringMinLength int          `json:"stringMinLength,omitempty"` // "stringMinLength":2
	Type            string       `json:"type,omitempty"`            // "type":""
	UpperBound      int          `json:"upperBound,omitempty"`      // "upperBound": 78
	ValueExpression string       `json:"valueExpression,omitempty"` // "valueExpression":""
	WarningText     string       `json:"warningText,omitempty"`     // "warningText":""
}

type StorageCapability struct {
	ConotrollerCapabilities    []string `json:"controllerCapabilities,omitempty"`     // "controllerCapabilities":{}
	ControllerModes            []string `json:"controllerModes,omitempty"`            // "controllerModes":{}
	DedicatedSpareSupported    bool     `json:"dedicatedSpareSupported,omitempty"`    // "dedicatedSpareSupported":false
	DriveTechnologies          []string `json:"driveTechnologies,omitempty"`          // "driveTechnologies":{}
	DriveWriteCacheSupported   bool     `json:"driveWriteCacheSupported,omitempty"`   // "driveWriteCacheSupported":false
	MaximumDrives              int      `json:"maximumDrives,omitempty"`              // "maximumDrives":5
	NvmeBackplaneCapable       bool     `json:"nvmeBackplaneCapable,omitempty"`       // "nvmeBackplaneCapable":true
	RaidLevels                 []string `json:"raidLevels,omitempty"`                 // "raidLevels":{}
	StandupControllerSupported bool     `json:"standupControllerSupported,omitempty"` // "standupControllerSupported,":false
}
type ServerHardwareType struct {
	Adapters            []Adapter          `json:"adapters,omitempty"`            // "adapters": {},
	BiosSettings        []BiosSetting      `json:"biosSettings,omitempty"`        //"biosSettings": {},
	BootCapabilities    []string           `json:"bootCapabilities,omitempty"`    // "bootCapabilities":{},
	BootModes           []string           `json:"bootModes,omitempty"`           // "bootModes":{},
	Capabilities        []string           `json:"capabilities,omitempty"`        // "capabilities": {},
	Category            string             `json:"category,omitempty"`            // "category": "server-hardware",
	Created             string             `json:"created,omitempty"`             // "created": "20150831T154835.250Z",
	Description         utils.Nstring      `json:"description,omitempty"`         // "description": "ServerHardware",
	ETAG                string             `json:"eTag,omitempty"`                // "eTag": "1441036118675/8",
	Family              string             `json:"family,omitempty"`              // "family":"135",
	FormFactor          string             `json:"formFactor,omitempty"`          // "formFactor":"HalfHeight",
	Model               string             `json:"model,omitempty"`               // "model":"",
	Modified            string             `json:"modified,omitempty"`            // "modified": "20150831T154835.250Z",
	Name                string             `json:"name,omitempty"`                // "name": "ServerHardware 1",
	PxeBootPolicies     []string           `json:"pxeBootPolicies,omitempty"`     // "pxeBootPolicies":{},
	StorageCapabilities *StorageCapability `json:"storageCapabilities,omitempty"` // "storageCapabilities":{..,},
	Type                string             `json:"type,omitempty"`                // "type": "server-hardware-type-10",
	URI                 utils.Nstring      `json:"uri,omitempty"`                 // "uri": "/rest/server-hardware-types/e2f0031b-52bd-4223-9ac1-d91cb519d548"
	UefiClass           string             `json:"uefiClass,omitempty"`           // "uefiClass":"2"
}

type ServerHardwareTypeList struct {
	Total       int                  `json:"total,omitempty"`       // "total": 1,
	Count       int                  `json:"count,omitempty"`       // "count": 1,
	Start       int                  `json:"start,omitempty"`       // "start": 0,
	PrevPageURI utils.Nstring        `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring        `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         utils.Nstring        `json:"uri,omitempty"`         // "uri": "/rest/server-hardware-types?filter=connectionTemplateUri%20matches%7769cae0-b680-435b-9b87-9b864c81657fsort=name:asc"
	Members     []ServerHardwareType `json:"members,omitempty"`     // "members":[]
}

func (c *OVClient) GetServerHardwareTypeByName(name string) (ServerHardwareType, error) {
	var (
		serverHardwareType ServerHardwareType
	)
	serverHardwareTypes, err := c.GetServerHardwareTypes(0, 0, fmt.Sprintf("name matches '%s'", name), "name:asc")
	if serverHardwareTypes.Total > 0 {
		return serverHardwareTypes.Members[0], err
	} else {
		return serverHardwareType, fmt.Errorf("Could not find Server Hardware Type: %s", name)
	}
}

func (c *OVClient) GetServerHardwareTypeByUri(uri utils.Nstring) (ServerHardwareType, error) {
	var (
		serverHardwareType ServerHardwareType
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri.String(), nil)
	if err != nil {
		return serverHardwareType, err
	}
	log.Debugf("GetServerHardwareType %s", data)
	if err := json.Unmarshal([]byte(data), &serverHardwareType); err != nil {
		return serverHardwareType, err
	}
	return serverHardwareType, nil
}

func (c *OVClient) GetServerHardwareTypes(start int, count int, filter string, sort string) (ServerHardwareTypeList, error) {
	var (
		uri                 = "/rest/server-hardware-types"
		q                   map[string]interface{}
		serverHardwareTypes ServerHardwareTypeList
	)
	q = make(map[string]interface{})
	if start > 0 {
		q["start"] = strconv.Itoa(start)
	}

	if count > 0 {
		q["count"] = strconv.Itoa(count)
	}

	if len(filter) > 0 {
		q["filter"] = filter
	}

	if sort != "" {
		q["sort"] = sort
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
		return serverHardwareTypes, err
	}

	log.Debugf("GetServerHardwareTypes %s", data)
	if err := json.Unmarshal([]byte(data), &serverHardwareTypes); err != nil {
		return serverHardwareTypes, err
	}
	return serverHardwareTypes, nil
}
