package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
	"strconv"
	"strings"
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
		clientOV         *ov.OVClient
		ethernet_network = "Auto-ethernet_network"
		tcId             = "1"
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

	logicalInterconnect, err := ovc.GetLogicalInterconnects("", "", "")
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("#-------------All Logical Interconnect Names----------------#")
		fmt.Println(logicalInterconnect.Members[0].Name)
	}
	interconnectURI := string(logicalInterconnect.Members[0].URI)
	id := strings.Replace(string(logicalInterconnect.Members[0].URI), "/rest/logical-interconnects/", "", 1)

	fmt.Println("....  Logical Interconnects Collection .....")
	logicalInterconnectList, _ := ovc.GetLogicalInterconnects("", "0", "10")
	fmt.Println(logicalInterconnectList)

	fmt.Println("....  Logical Interconnect by Id.....")
	lig, _ := ovc.GetLogicalInterconnectById(id)
	fmt.Println(lig)

	fmt.Println("....  Logical Interconnect PortMonitor.....")
	portMonitor, _ := ovc.GetLogicalInterconnectPortMonitor(id)
	fmt.Println(portMonitor)

	fmt.Println("....  Logical Interconnect EthernetSettings.....")
	ethernetSettings, _ := ovc.GetLogicalInterconnectEthernetSettings(id)
	fmt.Println(ethernetSettings)

	fmt.Println("....  Logical Interconnect Firmware.....")
	firmware, _ := ovc.GetLogicalInterconnectFirmware(id)
	fmt.Println(firmware)

	fmt.Println("....  Logical Interconnect SNMPConfiguration.....")
	snmpconfig, _ := ovc.GetLogicalInterconnectSNMPConfiguration(id)
	fmt.Println(snmpconfig)

	fmt.Println("....  Logical Interconnect Forwarding Information.....")
	var filter []string
	fi, er := ovc.GetLogicalInterconnectForwardingInformation(filter, id)
	fmt.Println(fi)
	if er != nil {
		fmt.Println("....  Logical Interconnect Forwarding Information By Mac Address.....")
		fi_mac, _ := ovc.GetLogicalInterconnectForwardingInformationByMacAddress(fi.Members[0].MacAddress, id)
		fmt.Println(fi_mac)

		fmt.Println("....  Logical Interconnect Forwarding Information By Internal Vlan.....")
		fi_intern_vlan, _ := ovc.GetLogicalInterconnectForwardingInformationByInternalVlan(fi.Members[0].InternalVlan, id)
		fmt.Println(fi_intern_vlan)

		fmt.Println("....  Logical Interconnect Forwarding Information By Interconnect URI and ExternalVlan.....")
		fi_interconnect_external, _ := ovc.GetLogicalInterconnectForwardingInformationByInterconnectAndExternalVlan(interconnectURI, fi.Members[0].ExternalVlan, id)
		fmt.Println(fi_interconnect_external)
	}

	fmt.Println("....  Logical Interconnect Internal VLAN IDs for the provisioned networks.....")
	fi_internal_vlan, _ := ovc.GetLogicalInternalVlans(id)
	fmt.Println(fi_internal_vlan)

	fmt.Println("....  Logical Interconnect QOS Configuration.....")
	fi_qos_config, _ := ovc.GetLogicalQosAggregatedConfiguration(id, "", "expand")
	fmt.Println(fi_qos_config)

	fmt.Println("....  Logical Interconnect Unassigned Ports for Port Monitor.....")
	port_monitor_ports := ovc.GetUnassignedPortsForPortMonitor(id)
	fmt.Println(port_monitor_ports)

	fmt.Println("....  Logical Interconnect Unassigned Uplink Ports for Port Monitor.....")
	uplink_port_monitor_ports, _ := ovc.GetUnassignedUplinkPortsForPortMonitor(id)
	fmt.Println(uplink_port_monitor_ports)

	fmt.Println("....  Logical Interconnect Telemetry Configuration.....")
	telemetry_config, _ := ovc.GetTelemetryConfigurations(id, "1")
	fmt.Println(telemetry_config)

	fmt.Println("....  Logical Interconnect IgmpSettings.....")
	igmpSettings, _ := ovc.GetLogicalInterconnectIgmpSettings(id)
	fmt.Println(igmpSettings)

	fmt.Println("....  Updating Logical Interconnect Consistent State.....")
	var liUris []utils.Nstring
	liUris = append(liUris, logicalInterconnect.Members[0].URI)
	liCompliance := ov.LogicalInterconnectCompliance{Type: "li-compliance",
		LogicalInterconnectUris: liUris,
		Description:             ""}
	err_compliance := ovc.UpdateLogicalInterconnectConsistentState(liCompliance)
	if err_compliance != nil {
		fmt.Println("Could not update ConsistentState of Logical Interconnect", err_compliance)
	}

	fmt.Println(".... Updating Logical Interconnect Consistent State by ID ....")
	err_update_compliance := ovc.UpdateLogicalInterconnectConsistentStateById(id)
	if err_update_compliance != nil {
		fmt.Println("Could not update ConsistentState of Logical Interconnect", err_update_compliance)
	}

	fmt.Println("....  Updating Logical Interconnect EthernetSetting .....")
	liEthernetSettings := ov.EthernetSettings{Type: "EthernetInterconnectSettingsV7",
		InterconnectType: "Ethernet",
		URI:              ethernetSettings.URI,
		ID:               ethernetSettings.ID}
	err_ethernet := ovc.UpdateLogicalInterconnectEthernetSettings(liEthernetSettings, id)
	if err_ethernet != nil {
		fmt.Println("Could not update Ethernet Settings of Logical Interconnect", err_ethernet)
	}

	fmt.Println("....  Updating Logical Interconnect EthernetSetting Force.....")
	err_ethernet_force := ovc.UpdateLogicalInterconnectEthernetSettingsForce(liEthernetSettings, id, true)
	if err_ethernet_force != nil {
		fmt.Println("Could not update Ethernet Settings of Logical Interconnect", err_ethernet_force)
	}

	fmt.Println("....  Updating Logical Interconnect Firmware.....")
	liFirmware := ov.Firmware{Command: "Update",
		EthernetActivationDelay: 5,
		EthernetActivationType:  "Parallel",
		FcActivationDelay:       5,
		FcActivationType:        "Parallel",
		Force:                   false,
		SppUri:                  utils.NewNstring("/rest/firmware-drivers/SPP_2018_06_20180709_for_HPE_Synergy_Z7550-96524")}
	err_firmware := ovc.UpdateLogicalInterconnectFirmware(liFirmware, id)
	if err_firmware != nil {
		fmt.Println("Could not update Firmware of Logical Interconnect", err_firmware)
	}

	fmt.Println("....  Updating Logical Interconnect Firmware Force.....")
	err_firmware_force := ovc.UpdateLogicalInterconnectFirmwareForce(liFirmware, id, true)
	if err_firmware_force != nil {
		fmt.Println("Could not update Firmware of Logical Interconnect", err_firmware_force)
	}

	fmt.Println("....  Updating Logical Interconnect InternalNetworks.....")
	eth_nw, _ := ovc.GetEthernetNetworkByName(ethernet_network)
	var internalNetworks []utils.Nstring
	internalNetworks = append(internalNetworks, eth_nw.URI)
	err_networks := ovc.UpdateLogicalInterconnectInternalNetworks(internalNetworks, id)
	if err_networks != nil {
		fmt.Println("Could not update Internal Networks of Logical Interconnect", err_networks)
	}
	fmt.Println("....  Updating Logical Interconnect InternalNetworks Force.....")
	err_networks_force := ovc.UpdateLogicalInterconnectInternalNetworksForce(internalNetworks, id, true)
	if err_networks_force != nil {
		fmt.Println("Could not update Internal Networks of Logical Interconnect", err_networks_force)
	}

	fmt.Println("....  Updating Logical Interconnect QOS Configuration.....")
	liActiveQosConfig := ov.ActiveQosConfig{Type: "QosConfiguration", Category: "qos-aggregated-configuration", ConfigType: "Passthrough"}
	liQosConfig := ov.QosConfiguration{Type: "qos-aggregated-configuration", Category: "qos-aggregated-configuration", ActiveQosConfig: liActiveQosConfig}

	err_qos := ovc.UpdateLogicalInterconnectQosConfigurations(liQosConfig, id)
	if err_qos != nil {
		fmt.Println("Could not update QOS Configuration of Logical Interconnect", err_qos)
	}

	fmt.Println("....  Updating Logical Interconnect SNMP Configuration.....")
	liSNMPConfig := ov.SnmpConfiguration{Type: "snmp-configuration", Category: "snmp-configuration", V3Enabled: newTrue()}

	err_snmp := ovc.UpdateLogicalInterconnectSNMPConfigurations(liSNMPConfig, id)
	if err_snmp != nil {
		fmt.Println("Could not update SNMP Configuration of Logical Interconnect", err_snmp)
	}

	fmt.Println("....  Updating Logical Interconnect Configuration.....")
	err_conf := ovc.UpdateLogicalInterconnectConfigurations(id)
	if err_conf != nil {
		fmt.Println("Could not update Configuration of Logical Interconnect", err_conf)
	}
	/*
		fmt.Println("....  Updating Logical Interconnect Port Monitor Configuration.....")
		liPMConfig := ov.PortMonitor{Type: "port-monitorV1", Category: "port-monitor", ETAG: "08d7c29-0c0e-4231-bf44-78bf96686455", Name: "name677351721-1594111671299"}

		err_pm := ovc.UpdateLogicalInterconnectPortMonitor(liPMConfig, id)
		if err_pm != nil {
			fmt.Println("Could not update PortMonitor Configuration of Logical Interconnect", err_pm)
		}
	*/
	li_uris := make(map[string][]utils.Nstring)
	li_uris["logicalInterconnectUris"] = []utils.Nstring{lig.URI}
	data, err := ovc.BulkInconsistencyValidations(li_uris)
	if err == nil {
		fmt.Println(data)
	} else {
		fmt.Println("Bulk Inconsistency Validation Failed", err)
	}

	fmt.Println("....  Updating Logical Interconnect Telemetry  Configuration.....")
	liTMConfig := ov.TelemetryConfiguration{Type: "telemetry-configurations", EnableTelemetry: newTrue(), SampleInterval: 300, SampleCount: 12, Name: "name-1563511630-1594111846712"}

	err_tm := ovc.UpdateLogicalInterconnectTelemetryConfigurations(liTMConfig, id, tcId)
	if err_tm != nil {
		fmt.Println("Could not update PortMonitor Configuration of Logical Interconnect", err_tm)
	}

	fmt.Println("....  Updating Logical Interconnect Igmp Configuration.....")
	dru := "/rest/logical-interconnects/" + id
	liIgmpConfig := ov.IgmpSettingsUpdate{Type: "IgmpSettings", Name: "Igmp-Update", EnableIgmpSnooping: newTrue(), IgmpIdleTimeoutInterval: 200, ID: id, DependentResourceUri: dru}

	err_is := ovc.UpdateLogicalInterconnectIgmpSettings(liIgmpConfig, id)
	if err_is != nil {
		fmt.Println("Could not update Igmp Configuration of Logical Interconnect", err_is)
	}

	fmt.Println("....  Updating Logical Interconnect PortFlap Configuration.....")
	portFlapConfig := ov.PortFlapProtection{Name: "portFlapProtection-Update", DetectionInterval: 30, PortFlapThresholdPerInterval: 3, NoOfSamplesDeclareFailures: 3, ConsistencyChecking: "ExactMatch"}

	err_pf := ovc.UpdateLogicalInterconnectPortFlapSettings(portFlapConfig, id)
	if err_pf != nil {
		fmt.Println("Could not update PortFlap Configuration of Logical Interconnect", err_pf)
	}

}
