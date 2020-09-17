package ov

import (
	"encoding/json"
	"fmt"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type Enclosure struct {
	ActiveOaPreferredIP                       string                `json:"activeOaPreferredIP,omitempty"`                       // "activeOaPreferredIP": "16.124.135.110",
	ApplianceBayCount                         int                   `json:"applianceBayCount,omitempty"`                         // "applianceBayCount": 16,
	ApplianceBays                             []ApplianceBay        `json:"applianceBays,omitempty"`                             // "applianceBays": [],
	AssetTag                                  string                `json:"assetTag,omitempty"`                                  // "assetTag": "",
	Category                                  string                `json:"category,omitempty"`                                  // "category": "enclosures",
	Created                                   string                `json:"created,omitempty"`                                   // "created": "20150831T154835.250Z",
	CrossBars                                 []CrossBar            `json:"crossBars,omitempty"`                                 // "crossBars": {},
	Description                               utils.Nstring         `json:"description,omitempty"`                               // "description": "Enclosure Group 1",
	DeviceBayCount                            int                   `json:"deviceBayCount,omitempty"`                            // "deviceBayCount": 16,
	DeviceBays                                []DeviceBayMap        `json:"deviceBays,omitempty`                                 // "deviceBays": [],
	DeviceBayWatts                            int                   `json:"deviceBayWatts,omitempty"`                            // "deviceBayWatts": 16,
	ETAG                                      string                `json:"eTag,omitempty"`                                      // "eTag": "1441036118675/8",
	EmBays                                    int                   `json:"emBays,omitempty"`                                    // "emBays": 16,
	EnclosureGroupUri                         utils.Nstring         `json:"enclosureGroupUri,omitempty"`                         // "enclosureGroupUri": "/rest/enclosure-groups/293e8efe-c6b1-4783-bf88-2d35a8e49071",
	EnclosureModel                            string                `json:"enclosureModel,omitempty"`                            // "enclosureModel": "Enclosure Group 1",
	EnclosureType                             string                `json:"enclosureType,omitempty"`                             // "enclosureType": "BladeSystem c7000 Enclosure",
	EnclosureTypeUri                          utils.Nstring         `json:"enclosureTypeUriomitempty"`                           // "enclosureTypeUri": "/rest/enclosure-groups/293e8efe-c6b1-4783-bf88-2d35a8e49071",
	FanBayCount                               int                   `json:"fanBayCount,omitempty"`                               // "fanBayCount": 16,
	FanBays                                   []FanBay              `json:"fanBays,omitempty`                                    // "fanBays": [],
	FanAndManagementDevicesWatts              int                   `json:"fanAndManagementDevicesWatts,omitempty"`              // "fanAndManagementDevicesWatts": 16,
	ForceInstallFirmware                      bool                  `json:"forceInstallFirmware,omitempty"`                      // "forceInstallFirmware": true
	FrameLinkModuleDomain                     string                `json:"frameLinkModuleDomain,omitempty"`                     // "frameLinkModuleDomain": "",
	FwBaselineName                            string                `json:"fwBaselineName,omitempty"`                            // "fwBaselineName": null,
	FwBaselineUri                             utils.Nstring         `json:"fwBaselineUri,omitempty"`                             // "fwBaselineUri": null,
	InterconnectBayCount                      int                   `json:"interconnectBayCount,omitempty"`                      // "interconnectBayCount": 8,
	InterconnectBays                          []InterconnectBay     `json:"interconnectBays"`                                    // "interconnectBays": [],
	InterconnectBayWatts                      int                   `json:"interconnectBayWatts,omitempty"`                      // "interconnectBayWatts": 8,
	IsFwManaged                               bool                  `json:"isFwManaged"`                                         // "isFwManaged": false,
	LicensingIntent                           string                `json:"licensingIntent,omitempty"`                           // "licensingIntent": "OneView",
	LogicalEnclosureUri                       utils.Nstring         `json:"logicalEnclosureUri,omitempty"`                       // "logicalEnclosureUri": null,
	ManagerBays                               []ManagerBay          `json:"managerBays,omitempty"`                               // "managerBays": [],
	MinimumPowerSupplies                      int                   `json:"minimumPowerSupplies,omitempty"`                      // "minimumPowerSupplies": 8,
	MinimumPowerSuppliesForRedundantPowerFeed int                   `json:"minimumPowerSuppliesForRedundantPowerFeed,omitempty"` // "minimumPowerSuppliesForRedundantPowerFeed": 8,
	Modified                                  string                `json:"modified,omitempty"`                                  // "modified": "20150831T154835.250Z",
	Name                                      string                `json:"name,omitempty"`                                      // "name": "e10",
	OaBays                                    int                   `json:"oaBays,omitempty"`                                    // "oaBays": 2,
	PartNumber                                string                `json:"partNumber,omitempty"`                                // "partNumber": "403320-B21",
	Partitions                                []Partition           `json:"partitions,omitempty"`                                // "partitions": [],
	PowerAllocatedWatts                       int                   `json:"powerAllocatedWatts,omitempty"`                       // "powerAllocatedWatts": ""
	PowerAvailableWatts                       int                   `json:"powerCapacityWatts,omitempty"`                        // "powerCapacityBoostWatts": ""
	PowerCapacityBoostWatts                   int                   `json:"powerCapacityBoostWatts,omitempty"`                   // "powerCapacityBoostWatts": ""
	PowerMode                                 string                `json:"powerMode,omitempty"`                                 // "powerMode": ""
	PowerSupplyBayCount                       int                   `json:"powerSupplyBayCount,omitempty"`                       // "powerSupplyBayCount": 1
	PowerSupplyBays                           []PowerSupplyBay      `json:"powerSupplyBay,omitempty"`                            // "powerSupplyBay": ""
	RackName                                  string                `json:"rackName,omitempty"`                                  // "rackName": "Rack-Renamed",
	ReconfigurationState                      string                `json:"reconfigurationState,omitempty"`                      // "reconfigurationState": "Pending"
	RefreshState                              string                `json:"refreshState,omitempty"`                              // "refreshState": "NotRefreshing",
	RemoteSupportSettings                     RemoteSupportSettings `json:"remoteSupportSettings,omitempty"`                     // "remoteSupportSettings": {},
	RemoteSupportUri                          utils.Nstring         `json:"remoteSupportUri,omitempty"`                          // "remoteSupportUri": "/rest/support/resources/enclosures/09USE62519EE",
	ScopesUri                                 utils.Nstring         `json:"scopesUri,omitempty"`                                 // "scopesUri": "/rest/scopes/resources/rest/server-profiles/DB7726F7-F601-4EA8-B4A6-D1EE1B32C07C",
	SerialNumber                              string                `json:"serialNumber,omitempty"`                              // "serialNumber": "USE62519EE",
	StandbyOaPreferredIP                      string                `json:"standbyOaPreferredIP,omitempty"`                      // "standbyOaPreferredIP": "",
	State                                     string                `json:"state,omitempty"`                                     // "state": "Configured",
	StateReason                               string                `json:"stateReason"`                                         // "stateReason": "None",
	Status                                    string                `json:"status,omitempty"`                                    // "status": "Critical",
	SupportDataCollectionState                string                `json:"supportDataCollectionState,omitempty"`                // "supportDataCollectionState": "PendingEnable",
	SupportDataCollectionType                 string                `json:"supportDataCollectionType,omitempty"`                 // "supportDataCollectionType": "",
	SupportDataCollectionUri                  utils.Nstring         `json:"supportDataCollectionUri,omitempty"`                  // "supportDataCollectionUri": "/rest/support/data-collections",
	SupportState                              string                `json:"supportState,omitempty"`                              // "supportState": "PendingEnable",
	Type                                      string                `json:"type,omitempty"`                                      // "type": "Enclosure",
	UIDState                                  string                `json:"uidState,omitempty"`                                  // "uidState": "Blink",
	URI                                       utils.Nstring         `json:"uri,omitempty"`                                       // "uri": "/rest/enclosures/09USE62519EE",
	UUID                                      string                `json:"uuid,omitempty"`                                      // "uuid": "09USE62519EE",
	VcmDomainId                               string                `json:"vcmDomainId,omitempty"`                               // "vcmDomainId": "@914ae756bdbce70cf7cbce65d34a23",
	VcmDomainName                             string                `json:"vcmDomainName,omitempty"`                             // "vcmDomainName": "OneViewDomain",
	VcmMode                                   bool                  `json:"vcmMode,omitempty"`                                   // "vcmMode": true,
	VcmUrl                                    string                `json:"vcmUrl,omitempty"`                                    // "vcmUrl": "https://16.124.128.80"
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

type ApplianceBay struct {
	BayNumber       int    `json:"bayNumber"`                 // "bayNumber": 1,
	BayPowerState   string `json:"bayPowerState,omitempty"`   // "bayPowerState": "Unknown",
	DevicePresence  string `json:"devicePresence,omitempty"`  // "devicePresence": "Present",
	Model           string `json:"model,omitempty"`           // "model": null,
	PartNumber      string `json:"partNumber,omitempty"`      // "partNumber": ""
	PoweredOn       bool   `json:"poweredOn,omitempty"`       // "poweredOn": true
	SerialNumber    string `json:"serialNumber,omitempty"`    // "serialNumber": "",
	SparePartNumber string `json:"sparePartNumber,omitempty"` // "sparePartNumber": ""
	Status          string `json:"status,omitempty"`          // "status": ""
}

type CrossBar struct {
	BayNumber    int    `json:"bayNumber,omitempty"`    // "bayNumber": 1,
	HwVersion    string `json:"hwVersion,omitempty"`    // "hwVersion": "",
	Manufacturer string `json:"manufacturer,omitempty"` // "manufacturer": ""
	PartNumber   string `json:"partNumber,omitempty"`   // "partNumber": ""
	Presence     string `json:"presence,omitempty"`     // "presence": ""
	SerialNumber string `json:"serialNumber,omitempty"` // "serialNumber": "",
	Status       string `json:"status,omitempty"`       // "status": ""

}

type DeviceBayMap struct {
	AvailableForFullHeightProfile           bool          `json:"availableForFullHeightProfile"`           // "availableForFullHeightProfile": false,
	AvailableForFullHeightDoubleWideProfile bool          `json:"availableForHalfHeightDoubleWideProfile"` // "availableForHalfHeightDoubleWideProfile": true,
	AvailableForHalfHeightProfile           bool          `json:"availableForHalfHeightProfile"`           // "availableForHalfHeightProfile": true,
	AvailableForHalfHeightDoubleWideProfile bool          `json:"availableForHalfHeightDoubleWideProfile"` // "availableForHalfHeightDoubleWideProfile": true,
	BayNumber                               int           `json:"bayNumber"`                               // "bayNumber": 1,
	BayPowerState                           string        `json:"bayPowerState,omitempty"`                 // "bayPowerState": "Unknown",
	Category                                string        `json:"category,omitempty"`                      // "category": "device-bays",
	ChangeState                             string        `json:"changeState,omitempty"`                   // "changeState": "None",
	CoveredByDevice                         utils.Nstring `json:"coveredByDevice,omitempty"`               // "coveredByDevice": "/rest/server-hardware/30373237-3132-4D32-3236-303730344E54",
	CoveredByProfile                        string        `json:"coveredByProfile,omitempty"`              // "coveredByProfile": null,
	Created                                 string        `json:"created,omitempty"`                       // "created": null,
	DeviceFormFactor                        string        `json:"deviceFormFactor,omitempty"`              // "deviceFormFactor": "SingleHeightSingleWide",
	DevicePresence                          string        `json:"devicePresence,omitempty"`                // "devicePresence": "Present",
	DeviceUri                               utils.Nstring `json:"deviceUri,omitempty"`                     // "deviceUri": "/rest/server-hardware/30373237-3132-4D32-3236-303730344E54",
	EnclosureUri                            utils.Nstring `json:"enclosureUri,omitempty"`                  // "enclosureUri": null,
	ETAG                                    string        `json:"eTag,omitempty"`                          // "eTag": null,
	Ipv4Setting                             Ipv4Setting   `json:"ipv4Setting,omitempty"`                   // "ipv4Setting": {},
	Model                                   string        `json:"model,omitempty"`                         // "model": null,
	Modified                                string        `json:"modified,omitempty"`                      // "modified": null,
	PowerAllocationWatts                    int           `json:"powerAllocationWatts,omitempty"`          // "powerAllocationWatts": 1,
	ProfileUri                              utils.Nstring `json:"profileUri,omitempty"`                    // "profileUri": null,
	SerialConsole                           bool          `json:"serialConsole,omitempty"`                 // "serialConsole": true,
	SerialNumber                            string        `json:"serialNumber,omitempty"`                  // "serialNumber": "",
	Type                                    string        `json:"type,omitempty"`                          // "type": "DeviceBay",
	URI                                     utils.Nstring `json:"uri,omitempty"`                           // "uri": "/rest/enclosures/09USE62519EE/device-bays/1"
}

type FanBay struct {
	BayNumber       int    `json:"bayNumber,omitempty"`       // "bayNumber": 1,
	ChangeState     string `json:"changeState,omitempty"`     // "changeState": "None",
	DevicePresence  string `json:"devicePresence,omitempty"`  // "devicePresence": "Present",
	DeviceRequired  bool   `json:"deviceRequired,omitempty"`  // "deviceRequired": false,
	Model           string `json:"model,omitempty"`           // "model": null,
	PartNumber      string `json:"partNumber,omitempty"`      // "partNumber": ""
	SerialNumber    string `json:"serialNumber,omitempty"`    // "serialNumber": "",
	SparePartNumber string `json:"sparePartNumber,omitempty"` // "sparePartNumber": ""
	State           string `json:"state,omitempty"`           // "state": null
	Status          string `json:"status,omitempty"`          // "status": ""
}

type InterconnectBay struct {
	BayNumber              int           `json:"bayNumber,omitempty"`              // "bayNumber": 1,
	BayPowerState          string        `json:"bayPowerState,omitempty"`          // "bayPowerState": "Unknown",
	ChangeState            string        `json:"changeState,omitempty"`            // "changeState": "None",
	Empty                  bool          `json:"empty,omitempty"`                  // "empty": false,
	EnclosureUri           utils.Nstring `json:"enclosureUri,omitempty"`           // "enclosureUri": "/rest/enclosures/013645CN759000AD",
	InterconnectBayType    string        `json:"interconnectBayType,omitempty"`    // "interconnectBayType": "SY12000InterconnectBay",
	InterconnectModel      string        `json:"interconnectModel,omitempty"`      // "interconnectModel": "Synergy 12Gb SAS Connection Module",
	InterconnectReady      bool          `json:"interconnectReady,omitempty"`      // "interconnectReady": true,
	InterconnectUri        utils.Nstring `json:"interconnectUri,omitempty"`        // "interconnectUri": "/rest/sas-interconnects/TWT546W04N",
	Ipv4Setting            Ipv4Setting   `json:"ipv4Setting,omitempty"`            // "ipv4Setting": {},
	LogicalInterconnectUri utils.Nstring `json:"logicalInterconnectUri,omitempty"` // "logicalInterconnectUri": "/rest/sas-logical-interconnects/23868fa4-b773-4fea-a1ab-44c0d30b2d50",
	OriginOfCondition      string        `json:"originOfCondition,omitempty"`      // "originOfCondition": "/rest/v1/InterconnectManager/1",
	PartNumber             string        `json:"partNumber,omitempty"`             // "partNumber": "755985-B21",
	PowerAllocationWatts   int           `json:"powerAllocationWatts,omitempty"`   // "powerAllocationWatts": 32,
	SerialConsole          bool          `json:"serialConsole,omitempty"`          // "serialConsole": true,
	SerialNumber           string        `json:"serialNumber,omitempty"`           // "serialNumber": "TWT546W04N",
}

type Ipv4Setting struct {
	IpAddress         string `json:"ipAddress,omitempty"`         // "ipAddress": "",
	IpAssignmentState string `json:"ipAssignmentState,omitempty"` // "ipAssignmentState": "",
	IpRangeUri        string `json:"ipRangeUri,omitempty"`        // "ipRangeUri": "",
	Mode              string `json:"mode,omitempty"`              // "mode": "",
}

type ManagerBay struct {
	BayNumber                  int              `json:"bayNumber,omitempty"`                  // "bayNumber": 1,
	BayPowerState              string           `json:"bayPowerState,omitempty"`              // "bayPowerState": "Unknown",
	ChangeState                string           `json:"changeState,omitempty"`                // "changeState": "None",
	DevicePresence             string           `json:"devicePresence,omitempty"`             // "devicePresence": "Present",
	EnclosureUri               utils.Nstring    `json:"enclosureUri,omitempty"`               // "enclosureUri": "/rest/enclosures/013645CN759000AC",
	FwBuildDate                string           `json:"fwBuildDate,omitempty"`                // "fwBuildDate": "05/23/2018,17:52:54",
	FwVersion                  string           `json:"fwVersion,omitempty"`                  // "fwVersion": "2.02.03",
	IpAddress                  string           `json:"ipAddress,omitempty"`                  // "ipAddress": "fe80::9657:a5ff:fe56:6b30",
	LinkedEnclosure            LinkedEnclosure  `json:"linkedEnclosure,omitempty"`            // "linkedEnclosure": {},
	LinkPortIsolated           bool             `json:"linkPortIsolated,omitempty"`           // "linkPortIsolated": false,
	LinkPortSpeedGbs           string           `json:"linkPortSpeedGbs,omitempty"`           // "linkPortSpeedGbs": "10",
	LinkPortState              string           `json:"linkPortState,omitempty"`              // "linkPortState": "Linked",
	LinkPortStatus             string           `json:"linkPortStatus,omitempty"`             // "linkPortStatus": "OK",
	ManagerType                string           `json:"managerType,omitempty"`                // "managerType": "EnclosureManager",
	MgmtPortLinkState          string           `json:"mgmtPortLinkState,omitempty"`          // "mgmtPortLinkState": "Linked",
	MgmtPortNeighbor           MgmtPortNeighbor `json:"mgmtPortNeighbor,omitempty"`           // "mgmtPortNeighbor": {},
	MgmtPortSpeedGbs           string           `json:"mgmtPortSpeedGbs,omitempty"`           // "mgmtPortSpeedGbs": "10",
	MgmtPortState              string           `json:"mgmtPortState,omitempty"`              // "mgmtPortState": "Standby",
	MgmtPortStatus             string           `json:"mgmtPortStatus,omitempty"`             // "mgmtPortStatus": "OK",
	Model                      string           `json:"model,omitempty"`                      // "model": "Synergy Frame Link Module",
	NegotiatedLinkPortSpeedGbs int              `json:"negotiatedLinkPortSpeedGbs,omitempty"` // "negotiatedLinkPortSpeedGbs": 10,
	NegotiatedMgmtPortSpeedGbs int              `json:"negotiatedMgmtPortSpeedGbs,omitempty"` // "negotiatedMgmtPortSpeedGbs": 10,
	PartNumber                 string           `json:"partNumber,omitempty"`                 // "partNumber": "802341-B21",
	Role                       string           `json:"role,omitempty"`                       // "role": "Standby",
	SerialNumber               string           `json:"serialNumber,omitempty"`               // "serialNumber": "CN7613V053",
	SparePartNumber            string           `json:"sparePartNumber,omitempty"`            // "sparePartNumber": "807963-001",
	Status                     string           `json:"status,omitempty"`                     // "status": "OK",
	UidState                   string           `json:"uidState,omitempty"`                   // "uidState": "Off"
}

type LinkedEnclosure struct {
	BayNumber    int    `json:"bayNumber,omitempty"`    // "bayNumber": 1,
	SerialNumber string `json:"serialNumber,omitempty"` // "serialNumber": "CN7613V053",
}

type MgmtPortNeighbor struct {
	Description string        `json:"description,omitempty"` // "description": "HPE Comware Platform Software",
	IpAddress   string        `json:"ipAddress,omitempty"`   // "ipAddress": null,
	MacAddress  string        `json:"macAddress,omitempty"`  // "macAddress": "5C:8A:38:4E:F2:4F",
	Port        string        `json:"port,omitempty"`        // "port": "Ten-GigabitEthernet1/1/1",
	ResourceUri utils.Nstring `json:"resourceUri,omitempty"` // "resourceUri": null
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

type Partition struct {
	AssociatedDevices string        `json:"associatedDevices,omitempty"` // "associatedDevices": "",
	DeviceCount       int           `json:"deviceCount,omitempty"`       // "deviceCount": 1,
	MemoryMb          int           `json:"memoryMb,omitempty"`          // "memoryMb": 1,
	MonarchDevice     int           `json:"monarchDevice,omitempty"`     // "monarchDevice": 1,
	PartitionHealth   string        `json:"partitionHealth,omitempty"`   // "partitionHealth": "",
	PartitionID       int           `json:"partitionID,omitempty"`       // "partitionID": "",
	PartitionName     string        `json:"partitionName,omitempty"`     // "partitionName": "",
	PartitionStatus   string        `json:"partitionStatus,omitempty"`   // "partitionStatus": "",
	ParttionUUID      string        `json:"parttionUUID,omitempty"`      // "parttionUUID": "",
	PendingChange     bool          `json:"pendingChange,omitempty"`     // "pendingChange": true,
	ProcessorCount    int           `json:"processorCount,omitempty"`    // "processorCount": 1,
	RunState          string        `json:"runState,omitempty"`          // "runState": "",
	ServerHardwareUri utils.Nstring `json:"serverHardwareUri,omitempty"` // "serverHardwareUri": "",
}

type PowerSupplyBay struct {
	BayNumber           int    `json:"bayNumber,omitempty"`           // "bayNumber": 1
	ChangeState         string `json:"changeState,omitempty"`         // "changeState": "None"
	DevicePresence      string `json:"devicePresence,omitempty"`      // "devicePresence": "Absent"
	Model               string `json:"model,omitempty"`               // "model": ""
	OutputCapacityWatts int    `json:"outputCapacityWatts,omitempty"` // "outputCapacityWatts": ""
	PartNumber          string `json:"partNumber,omitempty"`          // "partNumber": ""
	SerialNumber        string `json:"serialNumber,omitempty"`        // "serialNumber": ""
	SparePartNumber     string `json:"sparePartNumber,omitempty"`     // "sparePartNumber": ""
	Status              string `json:"status,omitempty"`              // "status": ""
}

type RemoteSupportSettings struct {
	Destination               string `json:"destination,omitemty"`                // "destination": ""
	RemoteSupportCurrentState string `json:"remoteSupportCurrentState,omitempty"` // "remoteSupportCurrentState": ""
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
	InitialScopeUris     []utils.Nstring      `json:"initialScopeUris"`
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
