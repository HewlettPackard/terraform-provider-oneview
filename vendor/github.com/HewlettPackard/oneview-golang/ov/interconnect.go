package ov

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

type Interconnect struct {
	BaseWWN                       string               `json:"baseWWN,omitempty"`                       // "baseWWN": "10:00:00:11:0A:06:08:69",
	Category                      string               `json:"category,omitempty"`                      // "category": "interconnects",
	Created                       string               `json:"created,omitempty"`                       // "created": "2018-08-02T15:49:59.963Z",
	Description                   utils.Nstring        `json:"description,omitempty"`                   // "description": null,
	DeviceResetState              string               `json:"deviceResetState,omitempty"`              // "deviceResetState": "Normal",
	EdgeVirtualBridgingAvailable  bool                 `json:"edgeVirtualBridgingAvailable,omitempty"`  // "edgeVirtualBridgingAvailable": false,
	EnableCutThrough              bool                 `json:"enableCutThrough,omitempty"`              // "enableCutThrough": false,
	EnableFastMacCacheFailover    bool                 `json:"enableFastMacCacheFailover,omitempty"`    // "enableFastMacCacheFailover": true,
	EnableIgmpSnooping            bool                 `json:"enableIgmpSnooping,omitempty"`            // "enableIgmpSnooping": false,
	EnableNetworkLoopProtection   bool                 `json:"enableNetworkLoopProtection,omitempty"`   // "enableNetworkLoopProtection": true,
	EnablePauseFloodProtection    bool                 `json:"enablePauseFloodProtection,omitempty"`    // "enablePauseFloodProtection": true,
	EnableRichTLV                 bool                 `json:"enableRichTLV,omitempty"`                 // "enableRichTLV": false,
	EnableStormControl            bool                 `json:"enableStormControl,omitempty"`            // "enableStormControl": false,
	EnableTaggedLldp              bool                 `json:"enableTaggedLldp,omitempty"`              // "enableTaggedLldp": false,
	EnclosureName                 string               `json:"enclosureName,omitempty"`                 // "enclosureName": "SYN03_Frame1",
	EnclosureType                 string               `json:"enclosureType,omitempty"`                 // "enclosureType": "SY12000",
	EnclosureUri                  utils.Nstring        `json:"enclosureUri,omitempty"`                  // "enclosureUri": "/rest/enclosures/013645CN759000AC",
	ETag                          string               `json:"eTag,omitempty"`                          // "eTag": "463bd328-ffc8-40ae-9603-6136fa9e6e58",
	FirmwareVersion               string               `json:"firmwareVersion,omitempty"`               // "firmwareVersion": "1.3.0.1005",
	HostName                      string               `json:"hostName,omitempty"`                      // "hostName": "VC4040F8-2TV5451754",
	IcmLicenses                   IcmLicenses          `json:"icmLicenses,omitempty"`                   // "icmLicenses": {},
	IgmpIdleTimeoutInterval       int                  `json:"igmpIdleTimeoutInterval,omitempty"`       // "igmpIdleTimeoutInterval": 260,
	IgmpSnoopingVlanIds           string               `json:"igmpSnoopingVlanIds,omitempty"`           // "igmpSnoopingVlanIds": "",
	InitialScopeUris              []string             `json:"initialScopeUris,omitempty"`              // "initialScopeUris": [],
	InterconnectIP                string               `json:"interconnectIP,omitempty"`                // "interconnectIP": "fe80::5eb9:1ff:fe47:f5d2",
	InterconnectLocation          InterconnectLocation `json:"interconnectLocation,omitempty"`          // "interconnectLocation": {}
	InterconnectMAC               string               `json:"interconnectMAC,omitempty"`               // "interconnectMAC": "5C:B9:01:47:F5:D2",
	InterconnectTypeUri           utils.Nstring        `json:"interconnectTypeUri,omitempty"`           // "interconnectTypeUri": "/rest/interconnect-types/59080afb-85b5-43ae-8c69-27c08cb91f3a",
	IpAddressList                 []IpAddressList      `json:"ipAddressList,omitempty"`                 // "ipAddressList": []
	LldpIpAddressMode             string               `json:"lldpIpAddressMode,omitempty"`             // "lldpIpAddressMode": "IPV4",
	LldpIpv4Address               string               `json:"lldpIpv4Address,omitempty"`               // "lldpIpv4Address": "",
	LldpIpv6Address               string               `json:"lldpIpv6Address,omitempty"`               // "lldpIpv6Address": "",
	LogicalInterconnectUri        utils.Nstring        `json:"logicalInterconnectUri,omitempty"`        // "logicalInterconnectUri": "/rest/logical-interconnects/d4468f89-4442-4324-9c01-624c7382db2d",
	MaxBandwidth                  string               `json:"maxBandwidth,omitempty"`                  // "maxBandwidth": "Speed_20G",
	MgmtInterface                 string               `json:"mgmtInterface,omitempty"`                 // "mgmtInterface": null,
	MigrationState                string               `json:"migrationState,omitempty"`                // "migrationState": null,
	Model                         string               `json:"model,omitempty"`                         // "model": "Virtual Connect SE 40Gb F8 Module for Synergy",
	Modified                      string               `json:"modified,omitempty"`                      // "modified": "2018-12-03T18:26:43.335Z",
	Name                          string               `json:"name,omitempty"`                          // "name": "SYN03_Frame1, interconnect 3",
	NetworkLoopProtectionInterval int                  `json:"networkLoopProtectionInterval,omitempty"` // "networkLoopProtectionInterval": 5,
	PartNumber                    string               `json:"partNumber,omitempty"`                    // "partNumber": "794502-B23",
	PortCount                     int                  `json:"portCount,omitempty"`                     // "portCount": 42,
	Ports                         []Port               `json:"ports,omitempty"`                         // "ports": [],
	PowerState                    string               `json:"powerState"`                              // "powerState": "On",
	ProductName                   string               `json:"productName,omitempty"`                   // "productName": "Virtual Connect SE 40Gb F8 Module for Synergy"
	QosConfiguration              QosConfiguration     `json:"qosConfiguration,omitempty"`              // "qosConfiguration": {},
	RemoteSupport                 RemoteSupport        `json:"remoteSupport,omitempty"`                 // "remoteSupport": {},
	Roles                         []string             `json:"roles,omitempty"`                         // "roles": []
	ScopesUri                     utils.Nstring        `json:"scopesUri,omitempty"`                     // "scopesUri": "/rest/scopes/resources/rest/interconnects/2b322628-e5a9-4843-b184-08345e7140c3",
	SerialNumber                  string               `json:"serialNumber,omitempty"`                  // "serialNumber": "2TV5451754",
	SnmpConfiguration             SnmpConfiguration    `json:"snmpConfiguration,omitempty"`             // "snmpConfiguration": {},
	SparePartNumber               string               `json:"sparePartNumber,omitempty"`               // "sparePartNumber": "813174-001",
	StackingDomainId              int                  `json:"stackingDomainId,omitempty"`              // "stackingDomainId": 3,
	StackingDomainRole            string               `json:"stackingDomainRole,omitempty"`            // "stackingDomainRole": "Master",
	StackingMemberId              int                  `json:"stackingMemberId,omitempty"`              // "stackingMemberId": 0,
	State                         string               `json:"state,omitempty"`                         // "state": "Configured",
	Status                        string               `json:"status,omitempty"`                        // "status": "OK",
	StormControlPollingInterval   int                  `json:"stormControlPollingInterval,omitempty"`   // "stormControlPollingInterval": 10,
	StormControlThreshold         int                  `json:"stormControlThreshold,omitempty"`         // "stormControlThreshold": 0,
	SubPortCount                  int                  `json:"subPortCount,omitempty"`                  // "subPortCount": 8,
	Type                          string               `json:"type,omitempty"`                          // "type": "InterconnectV4",
	UidState                      string               `json:"uidState,omitempty"`                      // "uidState": "Off",
	UnsupportedCapabilities       string               `json:unsupportedCapabilities,omitempty"`        // "unsupportedCapabilities": null,
	URI                           utils.Nstring        `json:"uri,omitempty"`                           // "uri": "/rest/interconnects/2b322628-e5a9-4843-b184-08345e7140c3"
}

type IpAddressList struct {
	IpAddress     string `json:"ipAddress,omitempty"`     // "ipAddress": "10.50.4.125",
	IpAddressType string `json:"ipAddressType,omitempty"` // "ipAddressType": "Ipv4Static"
}

type Port struct {
	AssociatedUplinkSetUri    utils.Nstring    `json:"associatedUplinkSetUri,omitempty"`    // "associatedUplinkSetUri": "/rest/uplink-sets/34d55132-fbb8-4ebc-aa3e-8164180ce845",
	Available                 bool             `json:"available,omitempty"`                 // "available": true,
	BayNumber                 int              `json:"bayNumber,omitempty"`                 // "bayNumber": 3,
	Capability                []string         `json:"capability,omitempty"`                // "capability": [],
	Category                  string           `json:"category,omitempty"`                  // "category": "ports",
	ConfigPortTypes           []string         `json:"configPortTypes,omitempty"`           // "configPortTypes": [],
	ConnectorType             string           `json:"connectorType,omitempty"`             // "connectorType": "QSFP+CR4",
	Created                   string           `json:"created,omitempty"`                   // "created": null,
	DcbxInfo                  DcbxInfo         `json:"dcbxInfo,omitempty"`                  // "dcbxInfo": {},
	Description               string           `json:"description,omitempty"`               // "description": null,
	Enabled                   bool             `json:"enabled,omitempty"`                   // "enabled": true,
	ETag                      string           `json:"eTag,omitempty"`                      // "eTag": null,
	FcPortProperties          FcPortProperties `json:"fcPortProperties,omitempty"`          // fcPortProperties: {}
	InterconnectName          string           `json:"interconnectName,omitempty"`          // "interconnectName": "SYN03_Frame1, interconnect 3",
	LagId                     int              `json:"lagId,omitempty"`                     // "lagId": 2,
	LagStates                 []string         `json:"lagStates,omitempty"`                 // "lagStates": [],
	Modified                  string           `json:"modified,omitempty"`                  // "modified": null,
	Name                      string           `json:"name,omitempty"`                      // "name": "Q3",
	Neighbor                  Neighbor         `json:"neighbor,omitempty"`                  // "neighbor": {},
	OperationalSpeed          string           `json:"operationalSpeed,omitempty"`          // "operationalSpeed": "Speed40G",
	PairedPortName            string           `json:"pairedPortName,omitempty"`            // "pairedPortName": null,
	PortHealthStatus          string           `json:"portHealthStatus,omitempty"`          // "portHealthStatus": "Normal",
	PortId                    string           `json:"portId,omitempty"`                    // "portId": "2b322628-e5a9-4843-b184-08345e7140c3:Q3",
	PortMonitorConfigInfo     string           `json:"portMonitorConfigInfo,omitempty"`     // "portMonitorConfigInfo": "NotMonitored",
	PortName                  string           `json:"portName,omitempty"`                  // "portName": "Q3",
	PortRunningCapabilityType string           `json:"portRunningCapabilityType,omitempty"` // "portRunningCapabilityType": null,
	PortSplitMode             string           `json:"portSplitMode,omitempty"`             // "portSplitMode": "Unsplit",
	PortStatus                string           `json:"portStatus,omitempty"`                // "portStatus": "Linked",
	PortStatusReason          string           `json:"portStatusReason,omitempty"`          // "portStatusReason": "Active",
	PortType                  string           `json:"portType,omitempty"`                  // "portType": "Uplink",
	PortTypeExtended          string           `json:"portTypeExtended,omitempty"`          // "portTypeExtended": "External",
	State                     string           `json:"state,omitempty"`                     // "state": null,
	Status                    string           `json:"status,omitempty"`                    // "status": "OK",
	SubPorts                  []SubPort        `json:"subports,omitempty"`                  // "subports": null,
	Type                      string           `json:"type,omitempty"`                      // "type": "port",
	URI                       utils.Nstring    `json:"uri,omitempty"`                       // "uri": "/rest/interconnects/2b322628-e5a9-4843-b184-08345e7140c3/ports/2b322628-e5a9-4843-b184-08345e7140c3:Q3",
	VendorSpecificPortName    string           `json:"vendorSpecificPortName,omitempty"`    // "vendorSpecificPortName": null,
	Vlans                     string           `json:"vlans,omitempty"`                     // "vlans": null
}

type DcbxInfo struct {
	DcbxApReason  string `json:"dcbxApReason,omitempty"`  // "dcbxApReason": "Disabled",
	DcbxPfcReason string `json:"dcbxPfcReason,omitempty"` // "dcbxPfcReason": "Disabled",
	DcbxPgReason  string `json:"dcbxPgReason,omitempty"`  // "dcbxPgReason": "Disabled",
	DcbxStatus    string `json:"dcbxStatus,omitempty"`    // "dcbxStatus": "NotApplicable",
}

type FcPortProperties struct {
	FcfMac                        string   `json:"fcfMac,omitempty"`                        // "fcfMac": "",
	Logins                        string   `json:"logins,omitempty"`                        // "logins": "",
	LoginsCount                   int      `json:"loginsCount,omitempty"`                   // "loginsCount": 0,
	NeighborInterconnectName      string   `json:"neighborInterconnectName,omitempty"`      // "neighborInterconnectName": "",
	OpOnline                      bool     `json:"opOnline,omitempty"`                      // "opOnline": false,
	OpOnlineReason                string   `json:"opOnlineReason,omitempty"`                // "opOnlineReason": "",
	PrincipleInterconnectName     string   `json:"principleInterconnectName,omitempty"`     // "principleInterconnectName": "",
	PrincipleInterconnectNameList []string `json:"principleInterconnectNameList,omitempty"` // "principleInterconnectNameList": [],
	TrunkMaster                   string   `json:"trunkMaster,omitempty"`                   // "trunkMaster": "",
	WWNN                          string   `json:"wwnn,omitempty"`                          // "wwnn": "",
	WWPN                          string   `json:"wwpn,omitempty"`                          // "wwpn": "",
}

type InterconnectLocation struct {
	LocationEntries []InterconnectLocationEntry `json:"locationEntries,omitempty"` // "locationEntries": []
}

type InterconnectLocationEntry struct {
	Type  string `json:"type"`  // "type": null
	Value string `json:"value"` // "value": null
}

type SubPort struct {
	PortNumber       int    `json:"portNumber"`
	PortStatus       string `json:"portStatus"`
	PortStatusReason string `json:"portStatusReason"`
}

type RemoteSupport struct {
	RemoteSupportUri           string          `json:"remoteSupportUri,omitempty"`           // "remoteSupportUri": "/rest/support/interconnects/2b322628-e5a9-4843-b184-08345e7140c3",
	SupportDataCollectionState string          `json:"supportDataCollectionState,omitempty"` // "supportDataCollectionState": ,
	SupportDataCollectionType  string          `json:"supportDataCollectionType,omitempty"`  // "supportDataCollectionType": ,
	SupportDataCollectionsUri  string          `json:"supportDataCollectionsUri,omitempty"`  // "supportDataCollectionsUri": "/rest/support/data-collections?deviceID=2b322628-e5a9-4843-b184-08345e7140c3&category=interconnects",
	SupportSettings            SupportSettings `json:"supportSettings,omitempty"`            // "supportSettings": {},
	SupportState               string          `json:"supportState,omitempty"`               // "supportState": "Disabled",
}

type SupportSettings struct {
	Destination         string `json:"destination,omitempty"`         // "destination": "",
	SupportCurrentState string `json:"supportCurrentState,omitempty"` // "supportCurrentState": "Unknown",

}

type Neighbor struct {
	LinkLabel                string `json:"linkLabel"`                // "linkLabel": null,
	LinkUri                  string `json:"linkUri"`                  // "linkUri": null,
	RemoteChassisId          string `json:"remoteChassisId"`          // "remoteChassisId": "5c:8a:38:4e:f2:4f",
	RemoteChassisIdType      string `json:"remoteChassisIdType"`      // "remoteChassisIdType": "macAddress",
	RemoteMgmtAddress        string `json:"remoteMgmtAddress"`        // "remoteMgmtAddress": "5c:8a:38:4e:f2:a0",
	RemoteMgmtAddressType    string `json:"remoteMgmtAddressType"`    // "remoteMgmtAddressType": "all802",
	RemotePortDescription    string `json:"remotePortDescription"`    // "remotePortDescription": "FortyGigE1/2/1 Interface",
	RemotePortId             string `json:"remotePortId"`             // "remotePortId": "FortyGigE1/2/1",
	RemotePortIdType         string `json:"remotePortIdType"`         // "remotePortIdType": "interfaceName",
	RemoteSystemCapabilities string `json:"remoteSystemCapabilities"` // "remoteSystemCapabilities": "Bridge, Router",
	RemoteSystemDescription  string `json:"remoteSystemDescription"`  // "remoteSystemDescription": "HPE Comware Platform Software, Software Version 7.1.070, Release 2612\r\nHPE FF 5930-2Slot+2QSFP+Switch\r\nCopyright (c) 2010-2018 Hewlett Packard Enterprise Development LP",
	RemoteSystemName         string `json:"remoteSystemName"`         // "remoteSystemName": "eco1FORGbE",
	RemoteType               string `json:"remoteType"`               // "remoteType": "external",
}

type InterconnectList struct {
	Total       int            `json:"total,omitempty"`       // "total": 1,
	Count       int            `json:"count,omitempty"`       // "count": 1,
	Start       int            `json:"start,omitempty"`       // "start": 0,
	PrevPageURI utils.Nstring  `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring  `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         utils.Nstring  `json:"uri,omitempty"`         // "uri": "/rest/interconnects?start=2&count=2",
	Members     []Interconnect `json:"members,omitempty"`     // "members":[]
}

func (c *OVClient) GetInterconnects(start string, count string, filter string, sort string) (InterconnectList, error) {
	var (
		uri           = "/rest/interconnects"
		q             map[string]interface{}
		interconnects InterconnectList
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
		return interconnects, err
	}

	log.Debugf("GetInterconnects %s", data)
	if err := json.Unmarshal([]byte(data), &interconnects); err != nil {
		return interconnects, err
	}
	return interconnects, nil
}

func (c *OVClient) GetInterconnectByName(name string) (Interconnect, error) {
	var (
		interconnect Interconnect
	)
	interconnects, err := c.GetInterconnects("", "", fmt.Sprintf("name matches '%s'", name), "name:asc")
	if interconnects.Total > 0 {
		return interconnects.Members[0], err
	} else {
		return interconnect, err
	}
}

func (c *OVClient) GetInterconnectByUri(uri utils.Nstring) (Interconnect, error) {
	var (
		interconnect Interconnect
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri.String(), nil)
	if err != nil {
		return interconnect, err
	}
	log.Debugf("GetEnclosureGroup %s", data)
	if err := json.Unmarshal([]byte(data), &interconnect); err != nil {
		return interconnect, err
	}
	return interconnect, nil
}
