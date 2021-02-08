package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
	"os"
)

func newTrue() *bool {
	b := true
	return &b
}
func newFalse() *bool {
	b := false
	return &b
}

func main() {
	var (
		clientOV     *ov.OVClient
		lig_name     = "LIG-FlexFabric3"
		lig_type     = "logical-interconnect-groupV8"
		new_lig_name = "RenamedLogicalInterConnectGroup"
		nw_set_name  = "Prod"
	)
	ovc := clientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		2000,
		"*")

	fmt.Println("#..........Creating Logical Interconnect Group.....#")
	locationEntry_first := ov.LocationEntry{Type: "Bay", RelativeValue: 3}
	locationEntry_second := ov.LocationEntry{Type: "Enclosure", RelativeValue: 1}
	locationEntries1 := new([]ov.LocationEntry)
	*locationEntries1 = append(*locationEntries1, locationEntry_first)
	*locationEntries1 = append(*locationEntries1, locationEntry_second)

	locationEntry_third := ov.LocationEntry{Type: "Bay", RelativeValue: 6}
	locationEntry_four := ov.LocationEntry{Type: "Enclosure", RelativeValue: 1}
	locationEntries2 := new([]ov.LocationEntry)
	*locationEntries2 = append(*locationEntries2, locationEntry_third)
	*locationEntries2 = append(*locationEntries2, locationEntry_four)

	logicalLocation1 := ov.LogicalLocation{LocationEntries: *locationEntries1}
	logicalLocation2 := ov.LogicalLocation{LocationEntries: *locationEntries2}

	// interconnect1, err := ovc.GetInterconnectTypeByName("Synergy 40Gb F8 Switch Module")
	// interconnect2, err := ovc.GetInterconnectTypeByName("Synergy 40Gb F8 Switch Module")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	interconnectMapEntryTemplate1 := ov.InterconnectMapEntryTemplate{LogicalLocation: logicalLocation1,
		PermittedInterconnectTypeUri: "/rest/interconnect-types/c657e79f-464f-422c-b556-8b92ef6da6a4",
		EnclosureIndex:               1}
	interconnectMapEntryTemplate2 := ov.InterconnectMapEntryTemplate{LogicalLocation: logicalLocation2,
		PermittedInterconnectTypeUri: "/rest/interconnect-types/c657e79f-464f-422c-b556-8b92ef6da6a4",
		EnclosureIndex:               1}
	interconnectMapEntryTemplates := new([]ov.InterconnectMapEntryTemplate)
	*interconnectMapEntryTemplates = append(*interconnectMapEntryTemplates, interconnectMapEntryTemplate1)

	*interconnectMapEntryTemplates = append(*interconnectMapEntryTemplates, interconnectMapEntryTemplate2)

	interconnectMapTemplate := ov.InterconnectMapTemplate{InterconnectMapEntryTemplates: *interconnectMapEntryTemplates}
	fmt.Println(&interconnectMapTemplate)

	enclosureIndexes := []int{1}

	ethernetSettings := ov.EthernetSettings{Type: "EthernetInterconnectSettingsV7",
		URI:                                "/settings",
		Name:                               "defaultEthernetSwitchSettings",
		ID:                                 "cf3509e5-5330-4464-8d4c-fc679bc3ad0b",
		InterconnectType:                   "Ethernet",
		EnableInterconnectUtilizationAlert: newFalse(),
		EnableFastMacCacheFailover:         newTrue(),
		MacRefreshInterval:                 5,
		EnableNetworkLoopProtection:        newTrue(),
		EnablePauseFloodProtection:         newTrue(),
		EnableRichTLV:                      newFalse()}

	igmpSettings := ov.IgmpSettings{Type: "IgmpSettings",
		Name:                    "defaultIgmpSettings",
		EnableIgmpSnooping:      newTrue(),
		ConsistencyChecking:     "ExactMatch",
		IgmpIdleTimeoutInterval: 260,
		EnableProxyReporting:    newTrue()}

	telemetryConfig := ov.TelemetryConfiguration{Type: "telemetry-configuration",
		EnableTelemetry: newTrue(),
		SampleCount:     12,
		SampleInterval:  300,
	}
	snmpConfig := ov.SnmpConfiguration{Type: "snmp-configuration",
		Enabled:   newFalse(),
		Category:  "snmp-configuration",
		V3Enabled: newTrue()}
	qosActiveConfig := ov.ActiveQosConfig{Type: "QosConfiguration",
		Category:   "qos-aggregated-configuration",
		ConfigType: "Passthrough"}
	qosConfig := ov.QosConfiguration{ActiveQosConfig: qosActiveConfig,
		Type:     "qos-aggregated-configuration",
		Category: "qos-aggregated-configuration"}

	le1 := ov.LocationEntry{
		RelativeValue: 1,
		Type:          "Enclosure",
	}
	le2 := ov.LocationEntry{
		RelativeValue: 3,
		Type:          "Bay",
	}
	le3 := ov.LocationEntry{
		RelativeValue: 62,
		Type:          "Port",
	}
	locationentries1 := new([]ov.LocationEntry)
	*locationentries1 = append(*locationentries1, le1)
	*locationentries1 = append(*locationentries1, le2)
	*locationentries1 = append(*locationentries1, le3)

	le4 := ov.LocationEntry{
		RelativeValue: 1,
		Type:          "Enclosure",
	}
	le5 := ov.LocationEntry{
		RelativeValue: 6,
		Type:          "Bay",
	}
	le6 := ov.LocationEntry{
		RelativeValue: 67,
		Type:          "Port",
	}
	locationentries2 := new([]ov.LocationEntry)
	*locationentries2 = append(*locationentries2, le4)
	*locationentries2 = append(*locationentries2, le5)
	*locationentries2 = append(*locationentries2, le6)
	ll1 := ov.LogicalLocation{LocationEntries: *locationentries1}
	ll2 := ov.LogicalLocation{LocationEntries: *locationentries2}

	lcp1 := ov.LogicalPortConfigInfo{DesiredSpeed: "Auto",
		LogicalLocation: ll1,
	}
	lcp2 := ov.LogicalPortConfigInfo{DesiredSpeed: "Auto",
		LogicalLocation: ll2,
	}
	logicalPortConfigInfos := new([]ov.LogicalPortConfigInfo)
	*logicalPortConfigInfos = append(*logicalPortConfigInfos, lcp1)
	*logicalPortConfigInfos = append(*logicalPortConfigInfos, lcp2)
	networkuris := &[]utils.Nstring{utils.NewNstring("/rest/ethernet-networks/1642c0c9-4e0d-4710-a25d-35180711720b")}
	up1 := ov.UplinkSets{
		EthernetNetworkType:    "Tagged",
		LacpTimer:              "Short",
		LogicalPortConfigInfos: *logicalPortConfigInfos,
		Mode:                   "Auto",
		LoadBalancingMode:      "SourceAndDestinationMac",
		Name:                   "up1",
		NativeNetworkUri:       "",
		NetworkType:            "Ethernet",
		NetworkUris:            *networkuris,
	}
	uplinksets := new([]ov.UplinkSets)
	*uplinksets = append(*uplinksets, up1)

	logicalInterconnectGroup := ov.LogicalInterconnectGroup{Type: lig_type,
		EthernetSettings:        &ethernetSettings,
		IgmpSettings:            &igmpSettings,
		Name:                    lig_name,
		TelemetryConfiguration:  &telemetryConfig,
		InterconnectMapTemplate: &interconnectMapTemplate,
		EnclosureType:           "SY12000",
		EnclosureIndexes:        enclosureIndexes,
		InterconnectBaySet:      3,
		RedundancyType:          "Redundant",
		SnmpConfiguration:       &snmpConfig,
		QosConfiguration:        &qosConfig,
		UplinkSets:              *uplinksets,
	}
	log.Infof("%#v", logicalInterconnectGroup)

	er := ovc.CreateLogicalInterconnectGroup(logicalInterconnectGroup)
	if er != nil {
		fmt.Println("........Logical Interconnect Group Creation failed:", er)
	} else {
		fmt.Println(".....Logical Interconnect Group Creation Success....")
	}

	fmt.Println("#..........Getting Logical Interconnect Group Collection.....")
	sort := "name:desc"
	logicalInterconnectGroupList, _ := ovc.GetLogicalInterconnectGroups(10, "", "", sort, 0)
	fmt.Println(logicalInterconnectGroupList)

	fmt.Println("....  Logical Interconnect Group by Name.....")
	lig, _ := ovc.GetLogicalInterconnectGroupByName(lig_name)

	fmt.Println(lig)

	fmt.Println("... Logical Interconnect Group by URI ....")
	uri := lig.URI
	lig_uri, _ := ovc.GetLogicalInterconnectGroupByUri(uri)
	fmt.Println(lig_uri)

	fmt.Println("... Getting setting for the specified Logical Interconnect Group ....")
	lig_s, _ := ovc.GetLogicalInterconnectGroupSettings(uri.String())
	fmt.Println(lig_s)

	fmt.Println("...Listing Logical Interconnect Group Default Settings .. ")
	lig_ds, _ := ovc.GetLogicalInterconnectGroupDefaultSettings()
	fmt.Println(lig_ds)

	fmt.Println("... Updating LogicalInterconnectGroup ...")
	fmt.Println("")
	lig_uri.Name = new_lig_name
	err1 := ovc.UpdateLogicalInterconnectGroup(lig_uri)
	if err1 != nil {
		panic(err1)
	} else {
		fmt.Println(".....Updated Logical Interconnect Group Successfully....")
	}
	lig1, _ := ovc.GetLogicalInterconnectGroupByName(new_lig_name)

	fmt.Println("... Logical Interconnect Group by URI ....")
	uri1 := lig1.URI
	lig_uri1, _ := ovc.GetLogicalInterconnectGroupByUri(uri1)

	fmt.Println("... Updating LogicalInterconnectGroup uplink set...")
	fmt.Println("")
	nw_set, _ := ovc.GetNetworkSetByName(nw_set_name)
	fmt.Println(nw_set.NetworkUris)
	//Adding the network set to first uplink set
	lig_uri1.UplinkSets[0].NetworkUris = append(lig_uri1.UplinkSets[0].NetworkUris, nw_set.NetworkUris...)

	err2 := ovc.UpdateLogicalInterconnectGroup(lig_uri1)
	if err2 != nil {
		panic(err2)
	} else {
		fmt.Println(".....Updated Logical Interconnect Group uplink set Successfully....")
	}

	fmt.Println("... Deleting LogicalInterconnectGroup ...")
	del_err := ovc.DeleteLogicalInterconnectGroup(lig_uri.Name)
	if del_err != nil {
		panic(del_err)
	} else {
		fmt.Println(".....Deleted Logical Interconnect Group Successfully....")
	}
}
