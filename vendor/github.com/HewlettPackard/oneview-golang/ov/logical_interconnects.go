package ov

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
	"strconv"
)

type LogicalInterconnect struct {
	Type                                 string                  `json:"type,omitempty"`                                 //"type": "logical-interconnect",
	URI                                  utils.Nstring           `json:"uri,omitempty"`                                  //"uri": "/rest/logical-interconnects/d4468f89-4442-4324-9c01-624c7382db2d",
	Category                             string                  `json:"category,omitempty"`                             //"category": "logical-interconnects",
	Interconnects                        []utils.Nstring         `json:"interconnects,omitempty"`                        //"interconnects": ["/rest/interconnects/b6b7325f-666f-474f-a8f7-2c32b3c9faab"],
	ConsistencyStatus                    string                  `json:"consistencyStatus,omitempty"`                    //"consistencyStatus": "NOT_CONSISTENT",
	ConsistencyStatusVerificationEnabled bool                    `json:"consistencyStatusVerificationEnabled,omitempty"` //"consistencyStatusVerificationEnabled": true,
	Created                              string                  `json:"created,omitempty"`                              //"created": "2018-08-07T22:48:07.640Z",
	Modified                             string                  `json:"modified,omitempty"`                             //"modified": "2018-12-14T06:36:09.956Z",
	Description                          utils.Nstring           `json:"description,omitempty"`                          //"description": null,
	DomainUri                            utils.Nstring           `json:"domainUri"`                                      //"domainUri": "/rest/domains/7b655ba0-11c1-4e9c-a191-672e1d0bbd8e",
	ETAG                                 string                  `json:"eTag,omitempty"`                                 //"eTag": "bedd16b1-1d27-4ecd-b483-407635de7431",
	EnclosureType                        string                  `json:"enclosureType,omitempty"`                        //"enclosureType": "SY12000",
	InternalNetworkUris                  []utils.Nstring         `json:"internalNetworkUris,omitempty"`                  //"internalNetworkUris": ["/rest/ethernet-networks/cbde97d0-c8f1-4aba-aa86-2b4e5d080401"],
	FabricUri                            utils.Nstring           `json:"fabricUri,omitempty"`                            //"fabricUri": "/rest/fabrics/c2fc09bc-323f-40b4-8a7e-d99ca28df969",
	QosConfiguration                     *QosConfiguration       `json:"qosConfiguration,omitempty"`                     // "qosConfiguration": {},
	IcmLicenses                          *IcmLicenses            `json:"icmLicenses,omitempty"`                          //"icmLicenses":{...},
	EnclosureUris                        []utils.Nstring         `json:"enclosureUris,omitempty"`                        //"enclosureUris": ["/rest/enclosures/013645CN759000AC"],
	FusionDomainUri                      utils.Nstring           `json:"fusionDomainUri,omitempty"`                      //"fusionDomainUri": "/rest/domains/7b655ba0-11c1-4e9c-a191-672e1d0bbd8e",
	EthernetSettings                     *EthernetSettings       `json:"ethernetSettings,omitempty"`                     //"ethernetSettings": {...},
	IgmpSettings                         *IgmpSettings           `json:"igmpSettings,omitempty"`                         //"igmpSettings": [...],
	LogicalInterconnectGroupUri          utils.Nstring           `json:"logicalInterconnectGroupUri,omitempty"`          //"logicalInterconnectGroupUri": "/rest/logical-interconnect-groups/3de143b3-bf77-4944-aa31-18084676b664",
	StackingHealth                       string                  `json:"stackingHealth,omitempty"`                       //"stackingHealth": "BiConnected",
	InitialScopeUris                     []string                `json:"initialScopeUris,omitempty"`                     //"InitialScopeUris": [],
	InterconnectMap                      *InterconnectMap        `json:"interconnectMap,omitempty"`                      //"interconnectMap": {...},
	PortMonitor                          *PortMonitor            `json:"portMonitor,omitempty"`                          //"portMonitor": {...},
	SnmpConfiguration                    *SnmpConfiguration      `json:"snmpConfiguration,omitempty"`                    // "snmpConfiguration": {...},
	TelemetryConfiguration               *TelemetryConfiguration `json:"telemetryConfiguration,omitempty"`               // "telemetryConfiguration": {...},
	ScopeUri                             string                  `json:"ScopeUri,omitempty"`                             //"ScopeUri": null,
	SecurityStandardMode                 string                  `json:"securityStandardMode,omitempty"`                 //"securityStandardMode": null,
	State                                string                  `json:"state,omitempty"`                                // "state": "Normal",
	Status                               string                  `json:"status,omitempty"`                               // "status": "Critical",
	Name                                 string                  `json:"name"`                                           // "name": "Logical Interconnect1",

}

type IcmLicenses struct {
	License []License `json:"License"` //"License": [{...}],
}
type License struct {
	LicenseType   string `json:"licenseType,omitempty"`   //"licenseType": "Synergy 8Gb FC Upgrade",
	State         string `json:"state,omitempty"`         //"state": "Yes",
	ConsumedCount int    `json:"consumedCount,omitempty"` //"consumedCount": 2,
	RequiredCount int    `json:"requiredCount,omitempty"` //"requiredCount": 0
}
type InterconnectMap struct {
	InterconnectMapEntries []InterconnectMapEntries `json:"interconnectMapEntries"` //"interconnectMapEntries": [{...}],
}

type InterconnectMapEntries struct {
	Location                     Location      `json:"location,omitempty"`                     //"location":{...},
	LogicalDownlinkUri           utils.Nstring `json:"logicalDownlinkUri,omitempty"`           //"logicalDownlinkUri": "/rest/logical-downlinks/126ee6e0-db82-4fac-a4ce-c8ee423b77dc",
	PermittedInterconnectTypeUri utils.Nstring `json:"permittedInterconnectTypeUri,omitempty"` //"permittedInterconnectTypeUri": "/rest/interconnect-types/59080afb-85b5-43ae-8c69-27c08cb91f3a",
	InterconnectUri              utils.Nstring `json:"interconnectUri,omitempty"`              //"interconnectUri": "/rest/interconnects/aca6687f-1370-46cd-b832-7e3192dbddfd",
	EnclosureIndex               int           `json:"enclosureIndex,omitempty"`               //"enclosureIndex": 2
}

type IgmpSettingsUpdate struct {
	ConsistencyChecking     string        `json:"consistencyChecking,omitempty"`     // "consistencyChecking":"ExactMatch"
	Created                 string        `json:"created,omitempty"`                 // "created": "20150831T154835.250Z",
	DependentResourceUri    string        `json:"dependentResourceUri,omitempty"`    // "dependentResourceUri": "/rest/logical-interconnect-groups/b7b144e9-1f5e-4d52-8534-2e39280f9e86",
	EnableIgmpSnooping      *bool         `json:"enableIgmpSnooping,omitempty"`      // "enableIgmpSnooping": true,
	ID                      string        `json:"id,omitempty"`                      // "id": "0c398238-2d35-48eb-9eb5-7560d59f94b3",
	IgmpIdleTimeoutInterval int           `json:"igmpIdleTimeoutInterval,omitempty"` // "igmpIdleTimeoutInterval": 260,
	Modified                string        `json:"modified,omitempty"`                // "modified": "20150831T154835.250Z",
	Name                    string        `json:"name,omitempty"`                    // "name": "IgmpSettings 1",
	Type                    string        `json:"type,omitempty"`                    // "type": "IgmpSettings"
	URI                     utils.Nstring `json:"uri,omitempty"`                     // "uri": "/rest/logical-interconnect-groups/b7b144e9-1f5e-4d52-8534-2e39280f9e86/igmpSettings"
}

type PortMonitor struct {
	Type              string            `json:"type,omitempty"`              //type": "port-monitor",
	URI               utils.Nstring     `json:"uri,omitempty"`               //"uri": "/rest/logical-interconnects/d4468f89-4442-4324-9c01-624c7382db2d/port-monitor",
	Category          string            `json:"category,omitempty"`          //"category": "port-monitor",
	ETAG              string            `json:"eTag,omitempty"`              //"eTag": "8a302a85-ec4d-4214-a3e0-10ef71d28769",
	Created           string            `json:"created,omitempty"`           //"created": null,
	Modified          string            `json:"modified,omitempty"`          //"modified": null,
	EnablePortMonitor bool              `json:"enablePortMonitor,omitempty"` //"enablePortMonitor": false,
	AnalyzerPort      MonitorPortInfo   `json:"analyzerPort,omitempty"`      //"analyzerPort": null,
	MonitoredPorts    []MonitorPortInfo `json:"monitoredPorts,omitempty"`    //"monitoredPorts": [],
	Description       utils.Nstring     `json:"description,omitempty"`       //"description": null,
	State             string            `json:"state,omitempty"`             //"state": null,
	Status            string            `json:"status,omitempty"`            //"status": null,
	Name              string            `json:"name,omitempty"`              //"name": "name2095641007-1533682087640"
}

type MonitorPortInfo struct {
	BayNumber             string `json:"bayNumber,omitempty"`
	InterconnectName      string `json:"interconnectName,omitempty"`
	InterconnectUri       string `json:"interconnectUri,omitempty"`
	PortHealthStatus      string `json:"portHealthStatus,omitempty"`
	PortMonitorConfigInfo string `json:"portMonitorConfigInfo,omitempty"`
	PortName              string `json:"portName,omitempty"`
	PortStatus            string `json:"portStatus,omitempty"`
	PortUri               string `json:"portUri,omitempty"`
}

type PortMonitorPortCollection struct {
	Type        string            `json:"type,omitempty"`        //type": "PortMonitorPortCollection",
	URI         utils.Nstring     `json:"uri,omitempty"`         //"uri":null
	Category    string            `json:"category,omitempty"`    //"category": null,
	ETAG        string            `json:"eTag,omitempty"`        //"eTag": null,
	Created     string            `json:"created,omitempty"`     // "created": null,
	Modified    string            `json:"modified,omitempty"`    // "modified": null,
	Start       int               `json:"start,omitempty"`       // "start": 0,
	Count       int               `json:"count,omitempty"`       //"count": 90,
	Total       int               `json:"total,omitempty"`       //"total": 90,
	PrevPageUri string            `json:"prevPageUri,omitempty"` //"prevPageUri": null,
	NextPageUri string            `json:"nextPageUri,omitempty"` //"nextPageUri": null,
	Members     []PortMonitorPort `json:"members,omitempty"`     //"members":[],
}

type PortMonitorPort struct {
	InterconnectName string `json:"interconnectName,omitempty"` //"interconnectName": "SYN03_Frame2, interconnect 6",
	PortName         string `json:"portName,omitempty"`         //"portName": "Q1:2",
	URI              string `json:"uri,omitempty"`              //"uri": "/rest/interconnects/aca6687f-1370-46cd-b832-7e3192dbddfd/ports/aca6687f-1370-46cd-b832-7e3192dbddfd:Q1:2
}
type LogicalInterconnectList struct {
	Type        string                `json:"type,omitempty"`        //"type": "LogicalInterconnectCollectionV5",
	Category    string                `json:"category,omitempty"`    //"category": "logical-interconnects",
	ETAG        string                `json:"eTag,omitempty"`        //"eTag": null,
	Created     string                `json:"created,omitempty"`     //"created": null,
	Modified    string                `json:"modified,omitempty"`    //"modified": null,
	Total       int                   `json:"total,omitempty"`       // "total": 1,
	Count       int                   `json:"count,omitempty"`       // "count": 1,
	Start       int                   `json:"start,omitempty"`       // "start": 0,
	PrevPageURI utils.Nstring         `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring         `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         utils.Nstring         `json:"uri,omitempty"`         // "uri": "/rest/logical-interconnects?start=0&count=1",
	Members     []LogicalInterconnect `json:"members,omitempty"`     // "members":[]
}
type Firmware struct {
	Command                 string                 `json:"command"`                           //command": "Update",
	FwBaseline              string                 `json:"fwBaseline,omitempty"`              //"fwBaseline": "1.16",
	SppUri                  utils.Nstring          `json:"sppUri"`                            //"sppUri": "/rest/firmware-drivers/SPP_2018_06_20180709_for_HPE_Synergy_Z7550-96524",
	SppName                 string                 `json:"sppName,omitempty"`                 //"sppName": "HPE Synergy Custom SPP 2018 06 19  2018 07 09 2018.07.09.00",
	Interconnects           []FirmwareInterconnect `json:"interconnects,omitempty"`           //"interconnects":[]
	State                   string                 `json:"state,omitempty"`                   // "state": "Activated",
	Force                   bool                   `json:"force,omitempty"`                   //"force": false,
	EthernetActivationType  string                 `json:"ethernetActivationType,omitempty"`  //"ethernetActivationType": "None",
	EthernetActivationDelay int                    `json:"ethernetActivationDelay,omitempty"` //"ethernetActivationDelay": 0,
	FcActivationType        string                 `json:"fcActivationType,omitempty"`        //"fcActivationType": "None",
	FcActivationDelay       int                    `json:"fcActivationDelay,omitempty"`       //"fcActivationDelay": 0,
	ValidationType          string                 `json:"validationType,omitempty"`          //"validationType": "None",
	LogicalSwitchId         string                 `json:"logicalSwitchId,omitempty"`         //"logicalSwitchId": "d4468f89-4442-4324-9c01-624c7382db2d"
}

type FirmwareInterconnect struct {
	DesiredFw        string `json:"desiredFw,omitempty"`        //"desiredFw": "1.16",
	DeviceType       string `json:"deviceType,omitempty"`       //"deviceType": "Synergy 20Gb Interconnect Link Module",
	InstalledFw      string `json:"installedFw,omitempty"`      //"installedFw": "1.16",
	InterconnectName string `json:"interconnectName,omitempty"` //"interconnectName": "SYN03_Frame2, interconnect 3",
	InterconnectUri  string `json:"interconnectUri,omitempty"`  //"interconnectUri": "/rest/interconnects/394a88fb-b6b0-430a-907f-aed03275b0ef",
	SppName          string `json:"sppName,omitempty"`          //"sppName": "HPE Synergy Custom SPP 2018 06 19  2018 07 09 2018.07.09.00",
	SppURI           string `json:"sppUri,omitempty"`           //"sppUri": "/rest/firmware-drivers/SPP_2018_06_20180709_for_HPE_Synergy_Z7550-96524",
	State            string `json:"state,omitempty"`            //"state": "Active",
	UpdateFlagDesc   string `json:"updateFlagDesc,omitempty"`   //"updateFlagDesc": "Interconnect already has same firmware version",
}

type InterconnectFibData struct {
	Count   int                        `json:"count,omitempty"`   //"count": "1",
	Members []InterconnectFibDataEntry `json:"members,omitempty"` //"members":[],
}

type InterconnectFibDataEntry struct {
	InterconnectName string        `json:"interconnectName,omitempty"` //"interconnectName": "SYN03_Frame2, interconnect 6",
	InterconnectURI  utils.Nstring `json:"interconnectUri,omitempty"`  //"interconnectUri": "/rest/interconnects/aca6687f-1370-46cd-b832-7e3192dbddfd",
	NetworkInterface string        `json:"networkInterface,omitempty"` //"networkInterface": "Q3",
	MacAddress       string        `json:"macAddress,omitempty"`       //"macAddress": "94:57:A5:67:2C:BE",
	EntryType        string        `json:"entryType,omitempty"`        //"entryType": "Learned",
	NetworkName      string        `json:"networkName,omitempty"`      //"networkName": "vlan504",
	NetworkUri       utils.Nstring `json:"networkUri,omitempty"`       //"networkUri": "/rest/ethernet-networks/6de2920a-8ad4-4cd8-865c-1907d3b4682e",
	ExternalVlan     string        `json:"externalVlan,omitempty"`     //"externalVlan": "504",
	InternalVlan     string        `json:"internalVlan,omitempty"`     //"internalVlan": "504"
}
type InternalVlanAssociationCollection struct {
	Type        string                    `json:"type"`                  //"type": "InternalVlanAssociationCollection",
	URI         string                    `json:"uri,omitempty"`         //"uri": null,
	Category    string                    `json:"category,omitempty"`    //"category": "logical-interconnects",
	ETAG        string                    `json:"eTag,omitempty"`        // "eTag": null,
	Created     string                    `json:"created,omitempty"`     // "created": null,
	Modified    string                    `json:"modified,omitempty"`    // "modified": null,
	Start       int                       `json:"start,omitempty"`       // "start": 0,
	Count       int                       `json:"count,omitempty"`       //"count": 7,
	Total       int                       `json:"total,omitempty"`       //"total": 7,
	PrevPageUri string                    `json:"prevPageUri,omitempty"` //"prevPageUri": null,
	NextPageUri string                    `json:"nextPageUri,omitempty"` //"nextPageUri": null,
	Members     []InternalVlanAssociation `json:"members,omitempty"`     //"members":[],
}

type InternalVlanAssociation struct {
	Type                   string `json:"type"`                             //"type": "internal-vlan-association",
	URI                    string `json:"uri,omitempty"`                    //"uri": null,
	Category               string `json:"category,omitempty"`               //"category": null,
	ETAG                   string `json:"eTag,omitempty"`                   //"eTag": null,
	Created                string `json:"created,omitempty"`                //"created": "2018-11-26T09:16:20.707Z",
	Modified               string `json:"modified,omitempty"`               //"modified": "2018-11-26T09:16:20.725Z",
	InternalVlanId         int    `json:"internalVlanId,omitempty"`         // "internalVlanId": -1,
	GeneralNetworkUri      string `json:"generalNetworkUri,omitempty"`      //"generalNetworkUri": "/rest/ethernet-networks/cbde97d0-c8f1-4aba-aa86-2b4e5d080401",
	LogicalInterconnectUri string `json:"logicalInterconnectUri,omitempty"` //"logicalInterconnectUri": "/rest/logical-interconnects/d4468f89-4442-4324-9c01-624c7382db2d",
	Description            string `json:"description,omitempty"`            // "description": null,
	State                  string `json:"state,omitempty"`                  // "state": null,
	Status                 string `json:"status,omitempty"`                 // "status": null,
	Name                   string `json:"name,omitempty"`                   // "name": null
}

type LogicalInterconnectCompliance struct {
	Type                    string          `json:"type"`                              //"type":"li-compliance",
	URI                     utils.Nstring   `json:"uri,omitempty"`                     //"uri": "/rest/logical-interconnects/d4468f89-4442-4324-9c01-624c7382db2d",
	Category                string          `json:"category,omitempty"`                //"category": "logical-interconnects",
	Created                 string          `json:"created,omitempty"`                 //"created": "2018-08-07T22:48:07.640Z",
	Modified                string          `json:"modified,omitempty"`                //"modified": "2018-12-14T06:36:09.956Z",
	Description             utils.Nstring   `json:"description,omitempty"`             //"description": null,
	EnclosureGroupUri       utils.Nstring   `json:"enclosureGroupUri,omitempty"`       //"enclosureGroupUri": "/rest/domains/7b655ba0-11c1-4e9c-a191-672e1d0bbd8e",
	EnclosureUris           []utils.Nstring `json:"enclosureUris,omitempty"`           //"enclosureUris": ["/rest/enclosures/013645CN759000AC"],
	ETAG                    string          `json:"eTag,omitempty"`                    //"eTag": "bedd16b1-1d27-4ecd-b483-407635de7431",
	LogicalInterconnectUris []utils.Nstring `json:"logicalInterconnectUris,omitempty"` //"logicalInterconnectUris":["/rest/logical-interconnects/d4468f89-4442-4324-9c01-624c7382db2d"]
	State                   string          `json:"state,omitempty"`                   // "state": null,
	Status                  string          `json:"status,omitempty"`                  // "status": null,
	Name                    string          `json:"name,omitempty"`                    // "name": null

}

type BulkInconsistencyValidation struct {
	AllowUpdateFromGroup                     bool                        `json:"allowUpdateFromGroup,omitempty"`                     // "allowUpdateFromGroup": false,
	LogicalInterconnectsReport               []LogicalInterconnectReport `json:"logicalInterconnectsReport,omitempty"`               // "LogicalInterconnectsReport":[]
	MaximumLimitOnUpdates                    int                         `json:"maximumLimitOnUpdates,omitempty"`                    // "maximumLimitOnUpdates": 1,
	NoOfConsistentLIs                        int                         `json:"noOfConsistentLIs,omitempty"`                        // "noOfConsistentLIs": 1,
	NoOfInconsistentLIs                      int                         `json:"noOfInconsistentLIs,omitempty"`                      // "noOfInconsistentLIs": 1,
	PortMonitoringConfigurationCumulative    bool                        `json:"portMonitoringConfigurationCumulative,omitempty"`    // "portMonitoringConfigurationCumulative": false,
	PossibleTrafficLossTransactionCumulative bool                        `json:"possibleTrafficLossTransactionCumulative,omitempty"` // "possibleTrafficLossTransactionCumulative": false,
	TotalNumberOfUpdates                     int                         `json:"totalNumberOfUpdates,omitempty"`                     // "totalNumberOfUpdates": 1,
}

type LogicalInterconnectReport struct {
	DeploymentClusterAffected   DepClusterAffected `json:"deploymentClusterAffected,omitempty"`   // "deploymentClusterAffected":[]
	Name                        string             `json:"name,omitempty"`                        // "name": null,
	NoOfProfilesAffected        int                `json:"noOfProfilesAffected,omitempty"`        // "noOfProfilesAffected": 1,
	NoOfUpdates                 int                `json:"noOfUpdates,omitempty"`                 // "noOfUpdates": 1,
	PortMonitoringConfiguration bool               `json:"portMonitoringConfiguration,omitempty"` // "portMonitoringConfiguration": false,
	ProfilesAffected            []ProfileAffected  `json:"profilesAffected,omitempty"`            // "profilesAffected":[]
	Report                      Reports            `json:"report,omitempty"`                      // "report":[]
	URI                         utils.Nstring      `json:"uri,omitempty"`                         // "uri": null,
}

type DepClusterAffected struct {
	ApplianceUri          utils.Nstring `json:"applianceUri,omitempty"`          // "applianceUri": null,
	DeploymentClusterName string        `json:"deploymentClusterName,omitempty"` // "deploymentClusterName": null,
	PrimaryClusterName    string        `json:"primaryClusterName,omitempty"`    // "primaryClusterName": null,
	ServerIP              utils.Nstring `json:"serverIP,omitempty"`              // "serverIP": null,
	URI                   utils.Nstring `json:"uri,omitempty"`                   // "uri": null,
}

type ProfileAffected struct {
	Name   string        `json:"name,omitempty"`   // "name": null,
	Status string        `json:"status,omitempty"` // "status": null,
	URI    utils.Nstring `json:"uri,omitempty"`    // "uri": null,
}

type Reports struct {
	AutomaticUpdates               []string        `json:"automaticUpdates,omitempty"`               // "automaticUpdates": [],
	ConnectivityReport             map[string]bool `json:"connectivityReport,omitempty"`             // "connectivityReport:"{},
	ManualUpdates                  []string        `json:"manualUpdates,omitempty"`                  // "manualUpdates": [],
	PossibleTrafficLossTransaction bool            `json:"possibleTrafficLossTransaction,omitempty"` // "possibleTrafficLossTransaction": false,
}

func (c *OVClient) BulkInconsistencyValidations(ligUris map[string][]utils.Nstring) (BulkInconsistencyValidation, error) {
	var (
		uri     = "/rest/logical-interconnects/bulk-inconsistency-validation"
		payload BulkInconsistencyValidation
	)

	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	log.Debugf("REST : %s \n %+v\n", uri, ligUris)
	data, err := c.RestAPICall(rest.POST, uri, ligUris)

	if err != nil {

		log.Errorf("Error submitting new BulkInconsistencyValidations request: %s", err)
		return payload, err
	}
	err = json.Unmarshal([]byte(data), &payload)
	if err != nil {
		log.Errorf("Error with payload un-marshal: %s", err)
		return payload, err
	}
	log.Debugf("Response New BulkInconsistencyValidations %s", data)

	return payload, nil
}

func (c *OVClient) GetLogicalInterconnectById(Id string) (LogicalInterconnect, error) {
	var (
		uri                 = "/rest/logical-interconnects/"
		logicalInterconnect LogicalInterconnect
	)
	uri += Id
	logicalInterconnect, err := c.GetLogicalInterconnectByUri(uri)

	return logicalInterconnect, err
}

func (c *OVClient) GetUnassignedPortsForPortMonitor(Id string) PortMonitorPortCollection {
	var (
		uri                       = "/rest/logical-interconnects/"
		portMonitorPortCollection PortMonitorPortCollection
	)
	uri = uri + Id + "/unassignedPortsForPortMonitor"
	//var retValue = c.GetReturn(uri, portMonitorPortCollection, "PortMonitorPortCollection")
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return portMonitorPortCollection
	}
	log.Debugf("GetLogicalInterconnect %s", data)
	if err := json.Unmarshal([]byte(data), &portMonitorPortCollection); err != nil {
		return portMonitorPortCollection
	}
	return portMonitorPortCollection
}
func (c *OVClient) GetUnassignedUplinkPortsForPortMonitor(Id string) (PortMonitorPortCollection, error) {
	var (
		uri                       = "/rest/logical-interconnects/"
		portMonitorPortCollection PortMonitorPortCollection
	)
	uri = uri + Id + "/unassignedUplinkPortsForPortMonitor"
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return portMonitorPortCollection, err
	}
	log.Debugf("GetLogicalInterconnect %s", data)
	if err := json.Unmarshal([]byte(data), &portMonitorPortCollection); err != nil {
		return portMonitorPortCollection, err
	}
	return portMonitorPortCollection, nil
}

func (c *OVClient) GetTelemetryConfigurations(Id string, TCId string) (TelemetryConfiguration, error) {
	var (
		uri                    = "/rest/logical-interconnects/"
		telemetryConfiguration TelemetryConfiguration
	)
	uri = uri + Id + "/telemetry-configurations/" + TCId
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return telemetryConfiguration, err
	}
	log.Debugf("GetLogicalInterconnect %s", data)
	if err := json.Unmarshal([]byte(data), &telemetryConfiguration); err != nil {
		return telemetryConfiguration, err
	}
	return telemetryConfiguration, nil
}

func (c *OVClient) GetLogicalInternalVlans(Id string) (InternalVlanAssociationCollection, error) {
	var (
		uri           = "/rest/logical-interconnects/"
		internalVlans InternalVlanAssociationCollection
	)
	uri = uri + Id + "/internalVlans"
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return internalVlans, err
	}
	log.Debugf("GetLogicalInterconnect %s", data)
	if err := json.Unmarshal([]byte(data), &internalVlans); err != nil {
		return internalVlans, err
	}
	return internalVlans, nil
}

func (c *OVClient) GetLogicalQosAggregatedConfiguration(Id string, fields string, view string) (QosConfiguration, error) {
	var (
		uri              = "/rest/logical-interconnects/"
		q                map[string]interface{}
		qosConfiguration QosConfiguration
	)
	uri = uri + Id + "/qos-aggregated-configuration"
	q = make(map[string]interface{})
	if fields != "" {
		q["fields"] = fields
	}
	if view != "" {
		q["view"] = view
	}
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	if len(q) > 0 {
		c.SetQueryString(q)
	}
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return qosConfiguration, err
	}
	log.Debugf("GetLogicalInterconnect QOS Configuration %s", data)
	if err := json.Unmarshal([]byte(data), &qosConfiguration); err != nil {
		return qosConfiguration, err
	}
	return qosConfiguration, nil
}

func (c *OVClient) GetLogicalInterconnectPortMonitor(Id string) (PortMonitor, error) {
	var (
		uri         = "/rest/logical-interconnects/"
		portMonitor PortMonitor
	)
	uri = uri + Id + "/port-monitor"
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return portMonitor, err
	}
	log.Debugf("GetLogicalInterconnect %s", data)
	if err := json.Unmarshal([]byte(data), &portMonitor); err != nil {
		return portMonitor, err
	}
	return portMonitor, nil
}

func (c *OVClient) GetLogicalInterconnectIgmpSettings(Id string) (IgmpSettings, error) {
	var (
		uri           = "/rest/logical-interconnects/"
		igmp_settings IgmpSettings
	)
	uri = uri + Id + "/igmpSettings"
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return igmp_settings, err
	}
	log.Debugf("GetLogicalInterconnect %s", data)
	if err := json.Unmarshal([]byte(data), &igmp_settings); err != nil {
		return igmp_settings, err
	}
	return igmp_settings, nil
}

func (c *OVClient) UpdateLogicalInterconnectIgmpSettings(IgmpConfig IgmpSettingsUpdate, Id string) error {
	var (
		uri = "/rest/logical-interconnects/"
		t   *Task
	)
	uri = uri + Id + "/igmpSettings"
	//	IgmpConfig["DependentResourceUri"] = uri + Id
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Infof("REST : %s \n %+v\n", uri, IgmpConfig)
	log.Infof("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, IgmpConfig)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error updating logicalInterConnect Igmp update request: %s", err)
		return err
	}

	log.Debugf("Response update LogicalInterConnect Igmp Settings %s", data)
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

func (c *OVClient) GetLogicalInterconnectEthernetSettings(Id string) (EthernetSettings, error) {
	var (
		uri              = "/rest/logical-interconnects/"
		ethernetSettings EthernetSettings
	)
	uri = uri + Id + "/ethernetSettings"
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return ethernetSettings, err
	}
	log.Debugf("GetLogicalInterconnect EthernetSettings %s", data)
	if err := json.Unmarshal([]byte(data), &ethernetSettings); err != nil {
		return ethernetSettings, err
	}
	return ethernetSettings, nil
}

func (c *OVClient) GetLogicalInterconnectFirmware(Id string) (Firmware, error) {
	var (
		uri      = "/rest/logical-interconnects/"
		firmware Firmware
	)
	uri = uri + Id + "/firmware"
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return firmware, err
	}
	log.Debugf("GetLogicalInterconnect Firmware %s", data)
	if err := json.Unmarshal([]byte(data), &firmware); err != nil {
		return firmware, err
	}
	return firmware, nil
}

func (c *OVClient) GetLogicalInterconnectSNMPConfiguration(Id string) (SnmpConfiguration, error) {
	var (
		uri               = "/rest/logical-interconnects/"
		snmpConfiguration SnmpConfiguration
	)
	uri = uri + Id + "/snmp-configuration"
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return snmpConfiguration, err
	}
	log.Debugf("GetLogicalInterconnect SNMP Configuration %s", data)
	if err := json.Unmarshal([]byte(data), &snmpConfiguration); err != nil {
		return snmpConfiguration, err
	}
	return snmpConfiguration, nil
}
func (c *OVClient) GetLogicalInterconnectForwardingInformationByMacAddress(MacAddress string, Id string) (InterconnectFibDataEntry, error) {
	var (
		InterconnectFibDataEntry InterconnectFibDataEntry
	)
	filter := make([]string, 1)
	filter[0] = fmt.Sprintf("macAddress='%s'", MacAddress)
	InterconnectFibData, err := c.GetLogicalInterconnectForwardingInformation(filter, Id)
	if InterconnectFibData.Count > 0 {
		return InterconnectFibData.Members[0], err
	} else {
		return InterconnectFibDataEntry, err
	}

}
func (c *OVClient) GetLogicalInterconnectForwardingInformationByInternalVlan(InternalVlan string, Id string) (InterconnectFibDataEntry, error) {
	var (
		InterconnectFibDataEntry InterconnectFibDataEntry
	)
	filter := make([]string, 1)
	filter[0] = fmt.Sprintf("internalVlan='%s'", InternalVlan)
	InterconnectFibData, err := c.GetLogicalInterconnectForwardingInformation(filter, Id)
	if InterconnectFibData.Count > 0 {
		return InterconnectFibData.Members[0], err
	} else {
		return InterconnectFibDataEntry, err
	}
}
func (c *OVClient) GetLogicalInterconnectForwardingInformationByInterconnectAndExternalVlan(InterconnectURI string, ExternalVlan string, Id string) (InterconnectFibDataEntry, error) {
	var (
		InterconnectFibDataEntry InterconnectFibDataEntry
	)
	filter := make([]string, 2)
	filter[0] = fmt.Sprintf("interconnectUri='%s'", InterconnectURI)
	filter[1] = fmt.Sprintf("externalVlan='%s'", ExternalVlan)
	InterconnectFibData, err := c.GetLogicalInterconnectForwardingInformation(filter, Id)
	if InterconnectFibData.Count > 0 {
		return InterconnectFibData.Members[0], err
	} else {
		return InterconnectFibDataEntry, err
	}
}

func (c *OVClient) GetLogicalInterconnectForwardingInformation(filter []string, Id string) (InterconnectFibData, error) {
	var (
		uri                 = "/rest/logical-interconnects/"
		q                   map[string]interface{}
		interconnectFibData InterconnectFibData
	)
	uri = uri + Id + "/forwarding-information-base"
	q = make(map[string]interface{})
	q["filter"] = filter
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	// Setup query
	if len(q) > 0 {
		c.SetQueryString(q)
	}
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return interconnectFibData, err
	}
	log.Debugf("GetLogicalInterconnects Forwarding Information %s", data)
	if err := json.Unmarshal([]byte(data), &interconnectFibData); err != nil {
		return interconnectFibData, err
	}
	return interconnectFibData, nil
}

func (c *OVClient) UpdateLogicalInterconnectConsistentState(liCompliance LogicalInterconnectCompliance) error {
	var (
		uri = "/rest/logical-interconnects/compliance"
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Infof("REST : %s \n %+v\n", uri, liCompliance)
	log.Infof("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, liCompliance)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error updating logicalInterConnectCompliance request: %s", err)
		return err
	}

	log.Debugf("Response update LogicalInterConnectCompliance %s", data)
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

func (c *OVClient) UpdateLogicalInterconnectConsistentStateById(Id string) error {
	var (
		uri = "/rest/logical-interconnects/"
		t   *Task
	)

	uri = uri + Id + "/compliance"
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Infof("REST : %s \n %+v\n", uri, nil)
	log.Infof("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, nil)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error updating logicalInterConnectCompliance request: %s", err)
		return err
	}

	log.Debugf("Response update LogicalInterConnectCompliance %s", data)
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

func (c *OVClient) UpdateLogicalInterconnectEthernetSettings(ethernetSetting EthernetSettings, Id string) error {
	err_ethernet := c.UpdateLogicalInterconnectEthernetSettingsForce(ethernetSetting, Id, false)
	return err_ethernet
}

func (c *OVClient) UpdateLogicalInterconnectEthernetSettingsForce(ethernetSetting EthernetSettings, Id string, force bool) error {
	var (
		uri = "/rest/logical-interconnects/"
		q   map[string]interface{}
		t   *Task
	)
	q = make(map[string]interface{})
	q["force"] = strconv.FormatBool(force)
	uri = uri + Id + "/ethernetSettings"
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	c.SetQueryString(q)
	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Infof("REST : %s \n %+v\n", uri, ethernetSetting)
	log.Infof("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, ethernetSetting)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error updating logicalInterConnect EthernetSetting request: %s", err)
		return err
	}

	log.Debugf("Response update LogicalInterConnect EthernetSetting %s", data)
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

func (c *OVClient) UpdateLogicalInterconnectFirmware(firmware Firmware, Id string) error {
	err_firmware := c.UpdateLogicalInterconnectFirmwareForce(firmware, Id, false)
	return err_firmware
}

func (c *OVClient) UpdateLogicalInterconnectFirmwareForce(firmware Firmware, Id string, force bool) error {
	var (
		uri = "/rest/logical-interconnects/"
		q   map[string]interface{}
		t   *Task
	)
	uri = uri + Id + "/firmware"
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	q = make(map[string]interface{})
	q["force"] = strconv.FormatBool(force)
	c.SetQueryString(q)
	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Infof("REST : %s \n %+v\n", uri, firmware)
	log.Infof("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, firmware)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error updating logicalInterConnect Firmware request: %s", err)
		return err
	}

	log.Debugf("Response update LogicalInterConnect Firmware %s", data)
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

func (c *OVClient) UpdateLogicalInterconnectInternalNetworks(internalNetworks []utils.Nstring, Id string) error {
	err_networks := c.UpdateLogicalInterconnectInternalNetworksForce(internalNetworks, Id, false)
	return err_networks
}

func (c *OVClient) UpdateLogicalInterconnectInternalNetworksForce(internalNetworks []utils.Nstring, Id string, force bool) error {
	var (
		uri = "/rest/logical-interconnects/"
		q   map[string]interface{}
		t   *Task
	)
	uri = uri + Id + "/internalNetworks"
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	q = make(map[string]interface{})
	q["force"] = strconv.FormatBool(force)
	c.SetQueryString(q)
	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Infof("REST : %s \n %+v\n", uri, internalNetworks)
	log.Infof("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, internalNetworks)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error updating logicalInterConnect InternalNetwork request: %s", err)
		return err
	}

	log.Debugf("Response update LogicalInterConnect InternalNetwork %s", data)
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

func (c *OVClient) UpdateLogicalInterconnectQosConfigurations(qosConfig QosConfiguration, Id string) error {
	var (
		uri = "/rest/logical-interconnects/"
		t   *Task
	)
	uri = uri + Id + "/qos-aggregated-configuration"
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Infof("REST : %s \n %+v\n", uri, qosConfig)
	log.Infof("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, qosConfig)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error updating logicalInterConnect InternalNetwork request: %s", err)
		return err
	}

	log.Debugf("Response update LogicalInterConnect InternalNetwork %s", data)
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

func (c *OVClient) UpdateLogicalInterconnectSNMPConfigurations(snmpConfig SnmpConfiguration, Id string) error {
	var (
		uri = "/rest/logical-interconnects/"
		t   *Task
	)
	uri = uri + Id + "/snmp-configuration"
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Infof("REST : %s \n %+v\n", uri, snmpConfig)
	log.Infof("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, snmpConfig)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error updating logicalInterConnect InternalNetwork request: %s", err)
		return err
	}

	log.Debugf("Response update LogicalInterConnect InternalNetwork %s", data)
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

func (c *OVClient) UpdateLogicalInterconnectTelemetryConfigurations(TConfig TelemetryConfiguration, Id string, TcId string) error {
	var (
		uri = "/rest/logical-interconnects/"
		t   *Task
	)
	uri = uri + Id + "/telemetry-configuration/" + TcId
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Infof("REST : %s \n %+v\n", uri, TConfig)
	log.Infof("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, TConfig)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error updating logicalInterConnect InternalNetwork request: %s", err)
		return err
	}

	log.Debugf("Response update LogicalInterConnect InternalNetwork %s", data)
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

func (c *OVClient) UpdateLogicalInterconnectPortMonitor(PMConfig PortMonitor, Id string) error {
	var (
		uri = "/rest/logical-interconnects/"
		t   *Task
	)
	uri = uri + Id + "/port-monitor"
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Infof("REST : %s \n %+v\n", uri, PMConfig)
	log.Infof("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, PMConfig)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error updating logicalInterConnect InternalNetwork request: %s", err)
		return err
	}

	log.Debugf("Response update LogicalInterConnect InternalNetwork %s", data)
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

func (c *OVClient) UpdateLogicalInterconnectConfigurations(Id string) error {
	var (
		uri = "/rest/logical-interconnects/"
		t   *Task
	)
	uri = uri + Id + "/configuration"
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Infof("REST : %s \n %+v\n", uri, nil)
	log.Infof("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, nil)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error updating logicalInterConnect InternalNetwork request: %s", err)
		return err
	}

	log.Debugf("Response update LogicalInterConnect InternalNetwork %s", data)
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

func (c *OVClient) GetLogicalInterconnectByUri(uri string) (LogicalInterconnect, error) {
	var (
		logicalInterconnect LogicalInterconnect
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return logicalInterconnect, err
	}
	log.Debugf("GetLogicalInterconnect %s", data)
	if err := json.Unmarshal([]byte(data), &logicalInterconnect); err != nil {
		return logicalInterconnect, err
	}
	return logicalInterconnect, nil
}

func (c *OVClient) GetLogicalInterconnects(sort string, start string, count string) (LogicalInterconnectList, error) {
	var (
		uri                     = "/rest/logical-interconnects"
		q                       map[string]interface{}
		logicalInterconnectList LogicalInterconnectList
	)
	q = make(map[string]interface{})
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
		return logicalInterconnectList, err
	}
	log.Debugf("GetLogicalInterconnects %s", data)
	if err := json.Unmarshal([]byte(data), &logicalInterconnectList); err != nil {
		return logicalInterconnectList, err
	}
	return logicalInterconnectList, nil
}
