package ov

import (
	"encoding/json"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
	"strconv"
)

type LogicalInterconnectGroup struct {
	Category                string                   `json:"category,omitempty"`               // "category": "logical-interconnect-groups",
	Created                 string                   `json:"created,omitempty"`                // "created": "20150831T154835.250Z",
	Description             utils.Nstring            `json:"description,omitempty"`            // "description": "Logical Interconnect Group 1",
	ETAG                    string                   `json:"eTag,omitempty"`                   // "eTag": "1441036118675/8",
	EnclosureIndexes        []int                    `json:"enclosureIndexes,omitempty"`       // "enclosureIndexes": [1],
	EnclosureType           string                   `json:"enclosureType,omitempty"`          // "enclosureType": "C7000",
	EthernetSettings        *EthernetSettings        `json:"ethernetSettings,omitempty"`       // "ethernetSettings": {...},
	IgmpSettings            *IgmpSettings            `json:"igmpSettings,omitempty"`           // "igmpSettings": {...},
	FabricUri               utils.Nstring            `json:"fabricUri,omitempty"`              // "fabricUri": "/rest/fabrics/9b8f7ec0-52b3-475e-84f4-c4eac51c2c20",
	InterconnectBaySet      int                      `json:"interconnectBaySet,omitempty"`     // "interconnectBaySet": 1,
	InterconnectMapTemplate *InterconnectMapTemplate `json:"interconnectMapTemplate"`          // "interconnectMapTemplate": {...},
	InternalNetworkUris     []utils.Nstring          `json:"internalNetworkUris,omitempty"`    // "internalNetworkUris": []
	Modified                string                   `json:"modified,omitempty"`               // "modified": "20150831T154835.250Z",
	Name                    string                   `json:"name"`                             // "name": "Logical Interconnect Group1",
	QosConfiguration        *QosConfiguration        `json:"qosConfiguration,omitempty"`       // "qosConfiguration": {},
	RedundancyType          string                   `json:"redundancyType,omitempty"`         // "redundancyType": "HighlyAvailable"
	ScopesUri               utils.Nstring            `json:"scopesUri,omitempty"`              // "scopesUri":""
	InitialScopeUris        []utils.Nstring          `json:"initialScopeUris,omitempty"`       // "initialScopeUris":[]
	NextPageUri             string                   `json:"NextPageUri,omitempty"`            // "NextPageUri":""
	PrevPageUri             string                   `json:"prevPageUri,omitempty"`            // "PrevPageUri":""
	Start                   int                      `json:"start,omitempty"`                  // "start":""
	Total                   int                      `json:"total,omitempty"`                  // "total":""
	SnmpConfiguration       *SnmpConfiguration       `json:"snmpConfiguration,omitempty"`      // "snmpConfiguration": {...}
	StackingHealth          string                   `json:"stackingHealth,omitempty"`         //"stackingHealth": "Connected",
	StackingMode            string                   `json:"stackingMode,omitempty"`           //"stackingMode": "Enclosure",
	State                   string                   `json:"state,omitempty"`                  // "state": "Normal",
	Status                  string                   `json:"status,omitempty"`                 // "status": "Critical",
	TelemetryConfiguration  *TelemetryConfiguration  `json:"telemetryConfiguration,omitempty"` // "telemetryConfiguration": {...},
	SflowConfiguration      *SflowConfiguration      `json:"sflowConfiguration,omitempty"`     // "sflowConfiguration": {...},
	Type                    string                   `json:"type"`                             // "type": "logical-interconnect-groupsV3",
	UplinkSets              []UplinkSets             `json:"uplinkSets"`                       // "uplinkSets": {...},
	URI                     utils.Nstring            `json:"uri,omitempty"`                    // "uri": "/rest/logical-interconnect-groups/e2f0031b-52bd-4223-9ac1-d91cb519d548",
}

type EthernetSettings struct {
	Category                           utils.Nstring `json:"category,omitempty"`                           // "category": null,
	ConsistencyChecking                string        `json:"consistencyChecking,omitempty"`                // "consistencyChecking":"ExactMatch"
	Created                            string        `json:"created,omitempty"`                            // "created": "20150831T154835.250Z",
	DependentResourceUri               utils.Nstring `json:"dependentResourceUri,omitempty"`               // "dependentResourceUri": "/rest/logical-interconnect-groups/b7b144e9-1f5e-4d52-8534-2e39280f9e86",
	Description                        utils.Nstring `json:"description,omitempty"`                        // "description": "Ethernet Settings",
	DomainName                         string        `json:"domainName,omitempty"           `              // "domainName": "ADHS",
	ETAG                               utils.Nstring `json:"eTag,omitempty"`                               // "eTag": "1441036118675/8",
	EnableCutThrough                   *bool         `json:"enableCutThrough,omitempty"`                   // "enableCutThrough": false,
	EnableDdns                         *bool         `json:"enableDdns,omitempty"`                         // "enableCutThrough": false,
	EnableFastMacCacheFailover         *bool         `json:"enableFastMacCacheFailover,omitempty"`         // "enableFastMacCacheFailover": false,
	EnableInterconnectUtilizationAlert *bool         `json:"enableInterconnectUtilizationAlert,omitempty"` // "enableInterconnectUtilizationAlert": false,
	EnableNetworkLoopProtection        *bool         `json:"enableNetworkLoopProtection,omitempty"`        // "enableNetworkLoopProtection": false,
	EnablePauseFloodProtection         *bool         `json:"enablePauseFloodProtection,omitempty"`         // "enablePauseFloodProtection": false,
	EnableRichTLV                      *bool         `json:"enableRichTLV,omitempty"`                      // "enableRichTLV": false,
	EnableStormControl                 *bool         `json:"enableStormControl,omitempty"`                 // "enableStormControl": false,
	EnableTaggedLldp                   *bool         `json:"enableTaggedLldp,omitempty"`                   // "enableTaggedLldp": false,
	ID                                 string        `json:"id,omitempty"`                                 // "id": "0c398238-2d35-48eb-9eb5-7560d59f94b3",
	LldpIpAddressMode                  string        `json:"lldpIpAddressMode,omitempty"`                  // "lldpIpAddressMode": "IPV4",
	LldpIpv4Address                    string        `json:"lldpIpv4Address,omitempty"`                    // "lldpIpv4Address": "",
	LldpIpv6Address                    string        `json:"lldpIpv6Address,omitempty"`                    // "lldpIpv6Address": "",
	InterconnectType                   string        `json:"interconnectType,omitempty"`                   // "interconnectType": "Ethernet",
	MacRefreshInterval                 int           `json:"macRefreshInterval,omitempty"`                 // "macRefreshInterval": 5,
	Modified                           string        `json:"modified,omitempty"`                           // "modified": "20150831T154835.250Z",
	Name                               string        `json:"name,omitempty"`                               // "name": "ethernetSettings 1",
	State                              string        `json:"state,omitempty"`                              // "state": "Normal",
	Status                             string        `json:"status,omitempty"`                             // "status": "Critical",
	StormControlPollingInterval        int           `json:"stormControlPollingInterval,omitempty"`        // "stormControlPollingInterval": 10,
	StormControlThreshold              int           `json:"stormControlThreshold,,omitempty"`             // "stormControlThreshold": 0,
	Type                               string        `json:"type,omitempty"`                               // "type": "EthernetInterconnectSettingsV5",
	URI                                utils.Nstring `json:"uri,omitempty"`                                // "uri": "/rest/logical-interconnect-groups/b7b144e9-1f5e-4d52-8534-2e39280f9e86/ethernetSettings"
}

type IgmpSettings struct {
	Category                utils.Nstring `json:"category,omitempty"`                // "category": null,
	ConsistencyChecking     string        `json:"consistencyChecking,omitempty"`     // "consistencyChecking":"ExactMatch"
	Created                 string        `json:"created,omitempty"`                 // "created": "20150831T154835.250Z",
	DependentResourceUri    string        `json:"dependentResourceUri,omitempty"`    // "dependentResourceUri": "/rest/logical-interconnect-groups/b7b144e9-1f5e-4d52-8534-2e39280f9e86",
	Description             string        `json:"description,omitempty"`             // "description": "Igmp Settings",
	ETAG                    utils.Nstring `json:"eTag,omitempty"`                    // "eTag": "1441036118675/8",
	EnableIgmpSnooping      *bool         `json:"enableIgmpSnooping,omitempty"`      // "enableIgmpSnooping": true,
	EnablePreventFlooding   *bool         `json:"enablePreventFlooding,omitempty"`   // "enablePreventFlooding": false,
	EnableProxyReporting    *bool         `json:"enableProxyReporting,omitempty"`    // "enableProxyReporting": false,
	ID                      string        `json:"id,omitempty"`                      // "id": "0c398238-2d35-48eb-9eb5-7560d59f94b3",
	IgmpIdleTimeoutInterval int           `json:"igmpIdleTimeoutInterval,omitempty"` // "igmpIdleTimeoutInterval": 260,
	IgmpSnoopingVlanIds     string        `json:"igmpSnoopingVlanIds,omitempty"`     // "igmpSnoopingVlanIds": "",
	Modified                string        `json:"modified,omitempty"`                // "modified": "20150831T154835.250Z",
	Name                    string        `json:"name,omitempty"`                    // "name": "IgmpSettings 1",
	State                   string        `json:"state,omitempty"`                   // "state": "Normal",
	Status                  string        `json:"status,omitempty"`                  // "status": "Critical",
	Type                    string        `json:"type,omitempty"`                    // "type": "IgmpSettings"
	URI                     utils.Nstring `json:"uri,omitempty"`                     // "uri": "/rest/logical-interconnect-groups/b7b144e9-1f5e-4d52-8534-2e39280f9e86/igmpSettings"
}

type InterconnectMapTemplate struct {
	InterconnectMapEntryTemplates []InterconnectMapEntryTemplate `json:"interconnectMapEntryTemplates"` // "interconnectMapEntryTemplates": {...},
}

type InterconnectMapEntryTemplate struct {
	EnclosureIndex               int             `json:"enclosureIndex,omitempty"`               // "enclosureIndex": 1,
	LogicalDownlinkUri           utils.Nstring   `json:"logicalDownlinkUri,omitempty"`           // "logicalDownlinkUri": "/rest/logical-downlinks/5b33fec1-63e8-40e1-9e3d-3af928917b2f",
	LogicalLocation              LogicalLocation `json:"logicalLocation,omitempty"`              // "logicalLocation": {...},
	PermittedInterconnectTypeUri utils.Nstring   `json:"permittedInterconnectTypeUri,omitempty"` //"permittedSwitchTypeUri": "/rest/switch-types/a2bc8f42-8bb8-4560-b80f-6c3c0e0d66e0",
}

type LogicalLocation struct {
	LocationEntries []LocationEntry `json:"locationEntries,omitempty"` // "locationEntries": {...}
}

type LocationEntry struct {
	RelativeValue int    `json:"relativeValue,omitempty"` //"relativeValue": 2,
	Type          string `json:"type,omitempty"`          //"type": "StackingMemberId",
}

type QosConfiguration struct {
	ActiveQosConfig          ActiveQosConfig           `json:"activeQosConfig,omitempty"`          //"activeQosConfig": {...},
	Category                 string                    `json:"category,omitempty"`                 // "category": "qos-aggregated-configuration",
	Created                  string                    `json:"created,omitempty"`                  // "created": "20150831T154835.250Z",
	Description              utils.Nstring             `json:"description,omitempty,omitempty"`    // "description": null,
	ETAG                     string                    `json:"eTag,omitempty"`                     // "eTag": "1441036118675/8",
	InactiveFCoEQosConfig    *InactiveFCoEQosConfig    `json:"inactiveFCoEQosConfig,omitempty"`    // "inactiveFCoEQosConfig": {...},
	InactiveNonFCoEQosConfig *InactiveNonFCoEQosConfig `json:"inactiveNonFCoEQosConfig,omitempty"` // "inactiveNonFCoEQosConfig": {...},
	Modified                 string                    `json:"modified,omitempty"`                 // "modified": "20150831T154835.250Z",
	Name                     string                    `json:"name,omitempty"`                     // "name": "Qos Config 1",
	State                    string                    `json:"state,omitempty"`                    // "state": "Normal",
	Status                   string                    `json:"status,omitempty"`                   // "status": "Critical",
	Type                     string                    `json:"type,omitempty"`                     // "qos-aggregated-configuration",
	URI                      utils.Nstring             `json:"uri,omitempty"`                      // "uri": null
}

type ActiveQosConfig struct {
	Category                   utils.Nstring          `json:"category,omitempty"`                   // "category": "null",
	ConfigType                 string                 `json:"configType,omitempty"`                 // "configType": "CustomWithFCoE",
	Created                    string                 `json:"created,omitempty"`                    // "created": "20150831T154835.250Z",
	Description                utils.Nstring          `json:"description,omitempty,omitempty"`      // "description": "Ethernet Settings",
	DownlinkClassificationType string                 `json:"downlinkClassificationType,omitempty"` //"downlinkClassifcationType": "DOT1P_AND_DSCP",
	ETAG                       string                 `json:"eTag,omitempty"`                       // "eTag": "1441036118675/8",
	Modified                   string                 `json:"modified,omitempty"`                   // "modified": "20150831T154835.250Z",
	Name                       string                 `json:"name,omitempty"`                       // "name": "active QOS Config 1",
	QosTrafficClassifiers      []QosTrafficClassifier `json:"qosTrafficClassifiers"`                // "qosTrafficClassifiers": {...},
	State                      string                 `json:"state,omitempty"`                      // "state": "Normal",
	Status                     string                 `json:"status,omitempty"`                     // "status": "Critical",
	Type                       string                 `json:"type,omitempty"`                       // "type": "QosConfiguration",
	UplinkClassificationType   string                 `json:"uplinkClassificationType,omitempty"`   // "uplinkClassificationType": "DOT1P"
	URI                        utils.Nstring          `json:"uri,omitempty"`                        // "uri": null
}

type InactiveFCoEQosConfig struct {
	Category                   utils.Nstring          `json:"category,omitempty"`                   // "category": "null",
	ConfigType                 string                 `json:"configType,omitempty"`                 // "configType": "CustomWithFCoE",
	Created                    string                 `json:"created,omitempty"`                    // "created": "20150831T154835.250Z",
	Description                utils.Nstring          `json:"description,omitempty,omitempty"`      // "description": "Ethernet Settings",
	DownlinkClassificationType string                 `json:"downlinkClassificationType,omitempty"` //"downlinkClassifcationType": "DOT1P_AND_DSCP",
	ETAG                       string                 `json:"eTag,omitempty"`                       // "eTag": "1441036118675/8",
	Modified                   string                 `json:"modified,omitempty"`                   // "modified": "20150831T154835.250Z",
	Name                       string                 `json:"name,omitempty"`                       // "name": "active QOS Config 1",
	QosTrafficClassifiers      []QosTrafficClassifier `json:"qosTrafficClassifiers,omitempty"`      // "qosTrafficClassifiers": {...},
	State                      string                 `json:"state,omitempty"`                      // "state": "Normal",
	Status                     string                 `json:"status,omitempty"`                     // "status": "Critical",
	Type                       string                 `json:"type,omitempty"`                       // "type": "QosConfiguration",
	UplinkClassificationType   string                 `json:"uplinkClassificationType,omitempty"`   // "uplinkClassificationType": "DOT1P"
	URI                        utils.Nstring          `json:"uri,omitempty"`                        // "uri": null
}

type InactiveNonFCoEQosConfig struct {
	Category                   utils.Nstring          `json:"category,omitempty"`                   // "category": "null",
	ConfigType                 string                 `json:"configType,omitempty"`                 // "configType": "CustomWithFCoE",
	Created                    string                 `json:"created,omitempty"`                    // "created": "20150831T154835.250Z",
	Description                utils.Nstring          `json:"description,omitempty,omitempty"`      // "description": "Ethernet Settings",
	DownlinkClassificationType string                 `json:"downlinkClassificationType,omitempty"` //"downlinkClassifcationType": "DOT1P_AND_DSCP",
	ETAG                       string                 `json:"eTag,omitempty"`                       // "eTag": "1441036118675/8",
	Modified                   string                 `json:"modified,omitempty"`                   // "modified": "20150831T154835.250Z",
	Name                       string                 `json:"name,omitempty"`                       // "name": "active QOS Config 1",
	QosTrafficClassifiers      []QosTrafficClassifier `json:"qosTrafficClassifiers,omitempty"`      // "qosTrafficClassifiers": {...},
	State                      string                 `json:"state,omitempty"`                      // "state": "Normal",
	Status                     string                 `json:"status,omitempty"`                     // "status": "Critical",
	Type                       string                 `json:"type,omitempty"`                       // "type": "QosConfiguration",
	UplinkClassificationType   string                 `json:"uplinkClassificationType,omitempty"`   // "uplinkClassificationType": "DOT1P"
	URI                        utils.Nstring          `json:"uri,omitempty"`                        // "uri": null
}

type QosTrafficClassifier struct {
	QosClassificationMapping *QosClassificationMap `json:"qosClassificationMapping"`  // "qosClassificationMapping": {...},
	QosTrafficClass          QosTrafficClass       `json:"qosTrafficClass,omitempty"` // "qosTrafficClass": {...},
}

type QosClassificationMap struct {
	Dot1pClassMapping []int    `json:"dot1pClassMapping"` // "dot1pClassMapping": [3],
	DscpClassMapping  []string `json:"dscpClassMapping"`  // "dscpClassMapping": [],
}

type QosTrafficClass struct {
	BandwidthShare   string `json:"bandwidthShare,omitempty"` // "bandwidthShare": "fcoe",
	ClassName        string `json:"className"`                // "className": "FCoE lossless",
	EgressDot1pValue int    `json:"egressDot1pValue"`         // "egressDot1pValue": 3,
	Enabled          *bool  `json:"enabled,omitempty"`        // "enabled": true,
	MaxBandwidth     int    `json:"maxBandwidth"`             // "maxBandwidth": 100,
	RealTime         *bool  `json:"realTime,omitempty"`       // "realTime": true,
}

//TODO SNMPConfiguration
type SnmpConfiguration struct {
	Category         utils.Nstring     `json:"category,omitempty"`         // "category": "snmp-configuration",
	Created          string            `json:"created,omitempty"`          // "created": "20150831T154835.250Z",
	Description      utils.Nstring     `json:"description,omitempty"`      // "description": null,
	ETAG             string            `json:"eTag,omitempty"`             // "eTag": "1441036118675/8",
	Enabled          *bool             `json:"enabled,omitempty"`          // "enabled": true,
	Modified         string            `json:"modified,omitempty"`         // "modified": "20150831T154835.250Z",
	Name             string            `json:"name,omitempty"`             // "name": "Snmp Config",
	ReadCommunity    string            `json:"readCommunity,omitempty"`    // "readCommunity": "public",
	SnmpAccess       []string          `json:"snmpAccess,omitempty"`       // "snmpAccess": [],
	snmpUsers        []Snmpv3User      `json:"snmpUsers,omitempty"`        // "snmpUsers": []
	State            string            `json:"state,omitempty"`            // "state": "Normal",
	Status           string            `json:"status,omitempty"`           // "status": "Critical",
	SystemContact    string            `json:"systemContact,omitempty"`    // "systemContact": "",
	TrapDestinations []TrapDestination `json:"trapDestinations,omitempty"` // "trapDestinations": {...}
	Type             string            `json:"type,omitempty"`             // "type": "snmp-configuration",
	URI              utils.Nstring     `json:"uri,omitempty"`              // "uri": null,
	V3Enabled        *bool             `json:"v3Enabled,omitempty"`        // "v3Enabled": true
}

type Snmpv3User struct {
	SnmpV3UserName    string             `json:"snmpV3UserName,omitempty"`    //"snmpV3UserName":"",
	UserCredentials   []ExtentedProperty `json:"userCredentials,omitempty"`   //"UserCredentials":"",
	V3AuthProtocol    string             `json:"v3AuthProtocol,omitempty"`    // "v3AuthProtocol":"",
	V3PrivacyProtocol string             `json:"v3PrivacyProtocol,omitempty"` // "v3PrivacyProtocol":""
}

type ExtentedProperty struct {
	PropertyName string `json:"propertyName"` //"propertyName":"",
	Value        string `json:"value"`        //"value":"",
	ValueFormat  string `json:"valueFormat"`  //"valueFormat":"",
	ValueType    string `json:"valueType"`    //"valueType":"",

}
type TrapDestination struct {
	CommunityString    string   `json:"communityString,omitempty"`    //"communityString": "public",
	EnetTrapCategories []string `json:"enetTrapCategories,omitempty"` //"enetTrapCategories": ["PortStatus", "Other"],
	FcTrapCategories   []string `json:"fcTrapCategories,omitempty"`   //"fcTrapCategories": ["PortStatus", "Other"]
	TrapDestination    string   `json:"trapDestination,omitempty"`    //"trapDestination": "127.0.0.1",
	TrapFormat         string   `json:"trapFormat,omitempty"`         //"trapFormat", "SNMPv1",
	TrapSeverities     []string `json:"trapSeverities,omitempty"`     //"trapSeverities": "Info",
	VcmTrapCategories  []string `json:"vcmTrapCategories,omitempty"`  // "vcmTrapCategories": ["Legacy"],
}

type TelemetryConfiguration struct {
	Category            string        `json:"category,omitempty"`            // "category": "telemetry-configuration",
	Created             string        `json:"created,omitempty"`             // "created": "20150831T154835.250Z",
	Description         utils.Nstring `json:"description,omitempty"`         // "description": null,
	ETAG                string        `json:"eTag,omitempty"`                // "eTag": "1441036118675/8",
	EnableTelemetry     *bool         `json:"enableTelemetry,omitempty"`     // "enableTelemetry": false,
	Modified            string        `json:"modified,omitempty"`            // "modified": "20150831T154835.250Z",
	Name                string        `json:"name,omitempty"`                // "name": "telemetry configuration",
	SampleCount         int           `json:"sampleCount,omitempty"`         // "sampleCount": 12
	SampleInterval      int           `json:"sampleInterval,omitempty"`      // "sampleInterval": 300,
	State               string        `json:"state,omitempty"`               // "state": "Normal",
	Status              string        `json:"status,omitempty"`              // "status": "Critical",
	Type                string        `json:"type,omitempty"`                // "type": "telemetry-configuration",
	URI                 utils.Nstring `json:"uri,omitempty"`                 // "uri": null,
	VcfcsampleIntervals string        `json:"vcfcsampleIntervals,omitempty"` // "vcfcsampleIntervals": "NOT_APPLICABLE"
}

type SflowConfiguration struct {
	Category        string           `json:"category,omitempty"`        // "category": "sflow-configuration",
	Created         string           `json:"created,omitempty"`         // "created": "20150831T154835.250Z",
	Description     utils.Nstring    `json:"description,omitempty"`     // "description": null,
	ETAG            string           `json:"eTag,omitempty"`            // "eTag": "1441036118675/8",
	Enabled         *bool            `json:"enabled,omitempty"`         // "enabled": false,
	Modified        string           `json:"modified,omitempty"`        // "modified": "20150831T154835.250Z",
	Name            string           `json:"name,omitempty"`            // "name": "sflow configuration",
	SflowAgents     []SflowAgent     `json:"sflowAgents,omitempty"`     // "sflowAgents": {...},
	SflowCollectors []SflowCollector `json:"sflowCollectors,omitempty"` // "sflowCollectors": {...},
	SflowNetwork    *SflowNetwork    `json:"sflowNetwork,omitempty"`    // "sflowNetwork": {...},
	SflowPorts      []SflowPort      `json:"sflowPorts,omitempty"`      // "sflowPorts": {...},
	State           string           `json:"state,omitempty"`           // "state": "Normal",
	Status          string           `json:"status,omitempty"`          // "status": "Critical",
	Type            string           `json:"type,omitempty"`            // "type": "sflow-configuration",
	URI             utils.Nstring    `json:"uri,omitempty"`             // "uri": null
}

type SflowAgent struct {
	BayNumber      int    `json:"bayNumber,omitempty"`      // "bayNumber": 1,
	EnclosureIndex int    `json:"enclosureIndex,omitempty"` // "enclosureIndex": 1,
	IpAddr         string `json:"ipAddr,omitempty"`         // "ipAddr": "172.18.1.11",
	IpMode         string `json:"ipMode,omitempty"`         // "ipAddr": "SflowAgent",
	Status         string `json:"status,omitempty"`         // "status": "Enabled",
	SubnetMask     string `json:"subnetMask,omitempty"`     // "subnetMask": "",
}

type SflowCollector struct {
	CollectorEnabled *bool  `json:"collectorEnabled,omitempty"` // "collectorEnabled": false,
	CollectorId      int    `json:"collectorId,omitempty"`      // "collectorId": 1,
	IPAddress        string `json:"ipAddress"`                  //"ipAddress": "172.18.1.11",
	MaxDatagramSize  int    `json:"maxDatagramSize,omitempty"`  // "maxDatagramSize": 1400,
	MaxHeaderSize    int    `json:"maxHeaderSize,omitempty"`    // "maxHeaderSize": 128,
	Name             string `json:"name,omitempty"`             // "name": "",
	Port             int    `json:"port,omitempty"`             // "port": 6343,
}

type SflowNetwork struct {
	Name   string        `json:"name,omitempty"`   // "name": "",
	URI    utils.Nstring `json:"uri,omitempty"`    // "uri": null,
	VlanId int           `json:"vlanId,omitempty"` // "vlanId": 1
}

type SflowPort struct {
	BayNumber               int                      `json:"bayNumber,omitempty"`               // "bayNumber": 1,
	CollectorId             int                      `json:"collectorId,omitempty"`             // "collectorId": 1,
	EnclosureIndex          int                      `json:"enclosureIndex,omitempty"`          // "enclosureIndex": 1,
	IcmName                 string                   `json:"icmName,omitempty"`                 // "icmName": "",
	PortName                string                   `json:"portName,omitempty"`                // "portName": "",
	SflowConfigurationModes []SflowConfigurationMode `json:"sflowConfigurationModes,omitempty"` // "sflowConfigurationModes": {...},
}

type SflowConfigurationMode struct {
	configurationMode string `json:"configurationMode,omitempty"` // "configurationMode": "SflowConfigurationMode",
}
type UplinkSets struct {
	EthernetNetworkType    string                  `json:"ethernetNetworkType,omitempty"` // "ethernetNetworkType": "Tagged",
	LacpTimer              string                  `json:"lacpTimer,omitempty"`           // "lacpTimer": "Long",
	LogicalPortConfigInfos []LogicalPortConfigInfo `json:"logicalPortConfigInfos"`        // "logicalPortConfigInfos": {...},
	Mode                   string                  `json:"mode,omitempty"`                // "mode": "Auto",
	FcMode                 string                  `json:"fcMode,omitempty"`              //"fcMode": "NA",
	LoadBalancingMode      string                  `json:"loadBalancingMode,omitempty"`   //"loadBalancingMode": "None",
	Name                   string                  `json:"name,omitempty"`                // "name": "Uplink 1",
	NativeNetworkUri       utils.Nstring           `json:"nativeNetworkUri,omitempty"`    // "nativeNetworkUri": null,
	NetworkType            string                  `json:"networkType,omitempty"`         // "networkType": "Ethernet",
	NetworkUris            []utils.Nstring         `json:"networkUris"`                   // "networkUris": ["/rest/ethernet-networks/f1e38895-721b-4204-8395-ae0caba5e163"]
	PrimaryPort            *LogicalLocation        `json:"primaryPort,omitempty"`         // "primaryPort": {...},
	PrivateVlanDomains     []PrivateVlanDomain     `json:"privateVlanDomains,omitempty"`  // "privateVlanDomains": {...}
	Reachability           string                  `json:"reachability,omitempty"`        // "reachability": "Reachable",
}

type PrivateVlanDomain struct {
	IsolatedNetwork *NetworkLite `json:"isolatedNetwork,omitempty"` // "isolatedNetwork": {...},
	PrimaryNetwork  *NetworkLite `json:"primaryNetwork,omitempty"`  // "primaryNetwork": {...},
}
type NetworkLite struct {
	Name   string        `json:"name,omitempty"`   // "name": "Uplink 1",
	URI    utils.Nstring `json:"uri,omitempty"`    // "uri": null,
	VlanId int           `json:"vlanId,omitempty"` // "vlanId": 100,
}
type LogicalPortConfigInfo struct {
	DesiredSpeed    string          `json:"desiredSpeed,omitempty"`    // "desiredSpeed": "Auto",
	LogicalLocation LogicalLocation `json:"logicalLocation,omitempty"` // "logicalLocation": {...},
}

type LogicalInterconnectGroupList struct {
	Total       int                        `json:"total,omitempty"`       // "total": 1,
	Count       int                        `json:"count,omitempty"`       // "count": 1,
	Start       int                        `json:"start,omitempty"`       // "start": 0,
	PrevPageURI utils.Nstring              `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI utils.Nstring              `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         utils.Nstring              `json:"uri,omitempty"`         // "uri": "/rest/server-profiles?filter=connectionTemplateUri%20matches%7769cae0-b680-435b-9b87-9b864c81657fsort=name:asc"
	Members     []LogicalInterconnectGroup `json:"members,omitempty"`     // "members":[]
}

type LogicalInterconnectGroupDefaultSettings struct {
	Type             string            `json:"type"`                       // "type": "InterconnectSettingsV4",
	URI              utils.Nstring     `json:"uri,omitempty"`              // "uri": null,
	Category         string            `json:"category,omitempty"`         // "category": null
	ETAG             string            `json:"eTag,omitempty"`             //"eTag": null,
	Created          string            `json:"created,omitempty"`          //"created": null,
	Modified         string            `json:"modified,omitempty"`         //"modified": null,
	EthernetSettings *EthernetSettings `json:"ethernetSettings,omitempty"` // "ethernetSettings": {...},
	IgmpSettings     *IgmpSettings     `json:"igmpSettings,omitempty"`     // "igmpSettings": {...},
	FcoeSettings     string            `json:"fcoeSettings,omitempty"`     // "fcoeSettings": null,
	Description      string            `json:"description,omitempty"`      // "description": null,
	State            string            `json:"state,omitempty"`            //  "state": null,
	Status           string            `json:"status,omitempty"`           // "status": null,
	Name             string            `json:"name,omitempty"`             // "name": null
}

func (c *OVClient) GetLogicalInterconnectGroupDefaultSettings() (LogicalInterconnectGroupDefaultSettings, error) {
	var (
		uri   = "/rest/logical-interconnect-groups/defaultSettings"
		ligDS LogicalInterconnectGroupDefaultSettings
	)
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return ligDS, err
	}
	log.Debugf("GetLogicalInterconnectGroup %s", data)
	if err := json.Unmarshal([]byte(data), &ligDS); err != nil {
		return ligDS, err
	}
	return ligDS, nil
}

func (c *OVClient) GetLogicalInterconnectGroupSettings(uri string) (LogicalInterconnectGroupDefaultSettings, error) {
	var (
		ligDS LogicalInterconnectGroupDefaultSettings
	)
	uri = uri + "/settings"
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return ligDS, err
	}
	log.Debugf("GetLogicalInterconnectGroup %s", data)
	if err := json.Unmarshal([]byte(data), &ligDS); err != nil {
		return ligDS, err
	}
	return ligDS, nil
}

func (c *OVClient) GetLogicalInterconnectGroupByName(name string) (LogicalInterconnectGroup, error) {
	var (
		logicalInterconnectGroup LogicalInterconnectGroup
	)
	logicalInterconnectGroups, err := c.GetLogicalInterconnectGroups(0, fmt.Sprintf("name matches '%s'", name), "", "name:asc", 0)
	if logicalInterconnectGroups.Total > 0 {
		return logicalInterconnectGroups.Members[0], err
	} else {
		return logicalInterconnectGroup, err
	}
}

func (c *OVClient) GetLogicalInterconnectGroupByUri(uri utils.Nstring) (LogicalInterconnectGroup, error) {
	var (
		lig LogicalInterconnectGroup
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri.String(), nil)
	if err != nil {
		return lig, err
	}
	log.Debugf("GetLogicalInterconnectGroup %s", data)
	if err := json.Unmarshal([]byte(data), &lig); err != nil {
		return lig, err
	}
	return lig, nil
}

func (c *OVClient) GetLogicalInterconnectGroups(count int, filter string, scopeUris string, sort string, start int) (LogicalInterconnectGroupList, error) {
	var (
		uri                       = "/rest/logical-interconnect-groups"
		q                         map[string]interface{}
		logicalInterconnectGroups LogicalInterconnectGroupList
	)
	q = make(map[string]interface{})
	if len(filter) > 0 {
		q["filter"] = filter
	}

	if sort != "" {
		q["sort"] = sort
	}
	if count != 0 {
		q["count"] = strconv.Itoa(count)
	}
	if scopeUris != "" {
		q["scopeUris"] = scopeUris
	}
	if start >= 0 {
		q["start"] = strconv.Itoa(start)
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
		return logicalInterconnectGroups, err
	}

	log.Debugf("GetLogicalInterconnectGroups %s", data)
	if err := json.Unmarshal([]byte(data), &logicalInterconnectGroups); err != nil {
		return logicalInterconnectGroups, err
	}
	return logicalInterconnectGroups, nil
}

func (c *OVClient) CreateLogicalInterconnectGroup(logicalInterconnectGroup LogicalInterconnectGroup) error {
	log.Infof("Initializing creation of logicalInterconnectGroup for %s.", logicalInterconnectGroup.Name)
	var (
		uri = "/rest/logical-interconnect-groups"
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()

	log.Debugf("REST : %s \n %+v\n", uri, logicalInterconnectGroup)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.POST, uri, logicalInterconnectGroup)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting new logical interconnect group request: %s", err)
		return err
	}

	log.Debugf("Response New LogicalInterconnectGroup %s", data)
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

func (c *OVClient) DeleteLogicalInterconnectGroup(name string) error {
	var (
		logicalInterconnectGroup LogicalInterconnectGroup
		err                      error
		t                        *Task
		uri                      string
	)

	logicalInterconnectGroup, err = c.GetLogicalInterconnectGroupByName(name)
	if err != nil {
		return err
	}
	if logicalInterconnectGroup.Name != "" {
		t = t.NewProfileTask(c)
		t.ResetTask()
		log.Debugf("REST : %s \n %+v\n", logicalInterconnectGroup.URI, logicalInterconnectGroup)
		log.Debugf("task -> %+v", t)
		uri = logicalInterconnectGroup.URI.String()
		if uri == "" {
			log.Warn("Unable to post delete, no uri found.")
			t.TaskIsDone = true
			return err
		}
		data, err := c.RestAPICall(rest.DELETE, uri, nil)
		if err != nil {
			log.Errorf("Error submitting delete logicalInterconnectGroup request: %s", err)
			t.TaskIsDone = true
			return err
		}

		log.Debugf("Response delete logicalInterconnectGroup %s", data)
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
		log.Infof("LogicalInterconnectGroup could not be found to delete, %s, skipping delete ...", name)
	}
	return nil
}

func (c *OVClient) UpdateLogicalInterconnectGroup(logicalInterconnectGroup LogicalInterconnectGroup) error {
	log.Infof("Initializing update of logicalInterConnectGroup for %s.", logicalInterconnectGroup.Name)
	var (
		uri = logicalInterconnectGroup.URI.String()
		t   *Task
	)
	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())

	t = t.NewProfileTask(c)
	t.ResetTask()
	log.Debugf("REST : %s \n %+v\n", uri, logicalInterconnectGroup)
	log.Debugf("task -> %+v", t)
	data, err := c.RestAPICall(rest.PUT, uri, logicalInterconnectGroup)
	if err != nil {
		t.TaskIsDone = true
		log.Errorf("Error submitting update logicalInterConnectGroup request: %s", err)
		return err
	}

	log.Debugf("Response update LogicalInterConnectGroup %s", data)
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
