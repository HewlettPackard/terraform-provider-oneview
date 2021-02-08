package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
	"strconv"
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
		lig_name     = "TestLIG-GO"
		lig_type     = "logical-interconnect-groupV8"
		new_lig_name = "RenamedLogicalInterConnectGroup"
	)
	apiversion, _ := strconv.Atoi(os.Getenv("ONEVIEW_APIVERSION"))

	ovc := clientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		apiversion,
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

	locationEntry_five := ov.LocationEntry{Type: "Bay", RelativeValue: 3}
	locationEntry_six := ov.LocationEntry{Type: "Enclosure", RelativeValue: 2}
	locationEntries3 := new([]ov.LocationEntry)
	*locationEntries3 = append(*locationEntries3, locationEntry_five)
	*locationEntries3 = append(*locationEntries3, locationEntry_six)

	locationEntry_seven := ov.LocationEntry{Type: "Bay", RelativeValue: 6}
	locationEntry_eight := ov.LocationEntry{Type: "Enclosure", RelativeValue: 2}
	locationEntries4 := new([]ov.LocationEntry)
	*locationEntries4 = append(*locationEntries4, locationEntry_seven)
	*locationEntries4 = append(*locationEntries4, locationEntry_eight)

	logicalLocation3 := ov.LogicalLocation{LocationEntries: *locationEntries3}
	logicalLocation4 := ov.LogicalLocation{LocationEntries: *locationEntries4}

	locationEntry_nine := ov.LocationEntry{Type: "Bay", RelativeValue: 3}
	locationEntry_ten := ov.LocationEntry{Type: "Enclosure", RelativeValue: 3}
	locationEntries5 := new([]ov.LocationEntry)
	*locationEntries5 = append(*locationEntries5, locationEntry_nine)
	*locationEntries5 = append(*locationEntries5, locationEntry_ten)

	locationEntry_eleven := ov.LocationEntry{Type: "Bay", RelativeValue: 6}
	locationEntry_twelle := ov.LocationEntry{Type: "Enclosure", RelativeValue: 3}
	locationEntries6 := new([]ov.LocationEntry)
	*locationEntries6 = append(*locationEntries6, locationEntry_eleven)
	*locationEntries6 = append(*locationEntries6, locationEntry_twelle)

	logicalLocation5 := ov.LogicalLocation{LocationEntries: *locationEntries5}
	logicalLocation6 := ov.LogicalLocation{LocationEntries: *locationEntries6}

	interconnect1, err := ovc.GetInterconnectTypeByName("Virtual Connect SE 40Gb F8 Module for Synergy")
	interconnect2, err := ovc.GetInterconnectTypeByName("Synergy 10Gb Interconnect Link Module")

	if err != nil {
		fmt.Println(err)
	}

	interconnectMapEntryTemplate1 := ov.InterconnectMapEntryTemplate{LogicalLocation: logicalLocation1,
		PermittedInterconnectTypeUri: interconnect1.URI,
		EnclosureIndex:               1}
	interconnectMapEntryTemplate2 := ov.InterconnectMapEntryTemplate{LogicalLocation: logicalLocation2,
		PermittedInterconnectTypeUri: interconnect1.URI,
		EnclosureIndex:               1}

	interconnectMapEntryTemplate3 := ov.InterconnectMapEntryTemplate{LogicalLocation: logicalLocation3,
		PermittedInterconnectTypeUri: interconnect2.URI,
		EnclosureIndex:               2}
	interconnectMapEntryTemplate4 := ov.InterconnectMapEntryTemplate{LogicalLocation: logicalLocation4,
		PermittedInterconnectTypeUri: interconnect2.URI,
		EnclosureIndex:               2}

	interconnectMapEntryTemplate5 := ov.InterconnectMapEntryTemplate{LogicalLocation: logicalLocation5,
		PermittedInterconnectTypeUri: interconnect2.URI,
		EnclosureIndex:               3}
	interconnectMapEntryTemplate6 := ov.InterconnectMapEntryTemplate{LogicalLocation: logicalLocation6,
		PermittedInterconnectTypeUri: interconnect2.URI,
		EnclosureIndex:               3}

	interconnectMapEntryTemplates := new([]ov.InterconnectMapEntryTemplate)
	*interconnectMapEntryTemplates = append(*interconnectMapEntryTemplates, interconnectMapEntryTemplate1)
	*interconnectMapEntryTemplates = append(*interconnectMapEntryTemplates, interconnectMapEntryTemplate2)
	*interconnectMapEntryTemplates = append(*interconnectMapEntryTemplates, interconnectMapEntryTemplate3)
	*interconnectMapEntryTemplates = append(*interconnectMapEntryTemplates, interconnectMapEntryTemplate4)
	*interconnectMapEntryTemplates = append(*interconnectMapEntryTemplates, interconnectMapEntryTemplate5)
	*interconnectMapEntryTemplates = append(*interconnectMapEntryTemplates, interconnectMapEntryTemplate6)

	interconnectMapTemplate := ov.InterconnectMapTemplate{InterconnectMapEntryTemplates: *interconnectMapEntryTemplates}
	fmt.Println(&interconnectMapTemplate)

	enclosureIndexes := []int{1, 2, 3}

	ethernetSettings := ov.EthernetSettings{Type: "EthernetInterconnectSettingsV7",
		URI:                                "/settings",
		Name:                               "defaultEthernetSwitchSettings",
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
		QosConfiguration:        &qosConfig}
	er := ovc.CreateLogicalInterconnectGroup(logicalInterconnectGroup)

	logicalInterconnectGroupAuto := ov.LogicalInterconnectGroup{Type: lig_type,
		EthernetSettings:        &ethernetSettings,
		IgmpSettings:            &igmpSettings,
		Name:                    "Auto-LIG",
		TelemetryConfiguration:  &telemetryConfig,
		InterconnectMapTemplate: &interconnectMapTemplate,
		EnclosureType:           "SY12000",
		EnclosureIndexes:        enclosureIndexes,
		InterconnectBaySet:      3,
		RedundancyType:          "Redundant",
		SnmpConfiguration:       &snmpConfig,
		QosConfiguration:        &qosConfig}
	er = ovc.CreateLogicalInterconnectGroup(logicalInterconnectGroupAuto)

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
	err = ovc.UpdateLogicalInterconnectGroup(lig_uri)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(".....Updated Logical Interconnect Group Successfully....")
	}
	fmt.Println("... Deleting LogicalInterconnectGroup ...")
	del_err := ovc.DeleteLogicalInterconnectGroup(lig_uri.Name)
	if del_err != nil {
		panic(del_err)
	} else {
		fmt.Println(".....Deleted Logical Interconnect Group Successfully....")
	}
}
