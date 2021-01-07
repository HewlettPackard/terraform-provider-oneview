package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		ClientOV           *ov.OVClient
		ethernet_network   = "eth1"
		ethernet_network_1 = "eth77"
		ethernet_network_2 = "eth88"
	)
	apiversion, _ := strconv.Atoi(os.Getenv("ONEVIEW_APIVERSION"))
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		apiversion,
		"")
	ovVer, _ := ovc.GetAPIVersion()
	fmt.Println(ovVer)

	scope := ov.Scope{Name: "ScopeTest", Description: "Test from script", Type: "ScopeV3"}
	_ = ovc.CreateScope(scope)
	scp, _ := ovc.GetScopeByName("ScopeTest")
	initialScopeUris := &[]utils.Nstring{scp.URI}

	fmt.Println("#................... Creating Ethernet Network ...............#")
	ethernetNetwork := ov.EthernetNetwork{Name: ethernet_network, VlanId: 9, Purpose: "General", SmartLink: false, PrivateNetwork: false, ConnectionTemplateUri: "", EthernetNetworkType: "Tagged", Type: "ethernet-networkV4", InitialScopeUris: *initialScopeUris}
	er := ovc.CreateEthernetNetwork(ethernetNetwork)

	fmt.Println("#................... Ethernet Network by Name ...............#")
	ethernet_nw, err := ovc.GetEthernetNetworkByName(ethernet_network)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ethernet_nw)
	}

	sort := "name:desc"
	ethernet_nw_list, err := ovc.GetEthernetNetworks("", "", "", sort)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("# ................... Ethernet Networks List .................#")
		for i := 0; i < len(ethernet_nw_list.Members); i++ {
			fmt.Println(ethernet_nw_list.Members[i].Name)
		}
	}

	ethernet_nw_id := strings.Replace(string(ethernet_nw.URI), "/rest/ethernet-networks/", "", 1)
	fmt.Println("#................... GetAssociatedProfiles ....................#")
	ethernet_nw_ass_pfl, err := ovc.GetAssociatedProfile(ethernet_nw_id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ethernet_nw_ass_pfl)
	}

	fmt.Println("#................... GetAssociatedUplinkGroups ...............#")
	ethernet_nw_uplinkgrps, err := ovc.GetAssociatedUplinkGroup(ethernet_nw_id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ethernet_nw_uplinkgrps)
	}

	bandwidth := ov.Bandwidth{MaximumBandwidth: 10000, TypicalBandwidth: 2000}

	ethernetNetwork = ov.EthernetNetwork{Name: "eth77", VlanId: 10, Purpose: "General", SmartLink: false, PrivateNetwork: false, ConnectionTemplateUri: "", EthernetNetworkType: "Tagged", Type: "ethernet-networkV4", InitialScopeUris: *initialScopeUris}
	er = ovc.CreateEthernetNetwork(ethernetNetwork)

	bulkEthernetNetwork := ov.BulkEthernetNetwork{VlanIdRange: "2-4", Purpose: "General", NamePrefix: "Test_eth", SmartLink: false, PrivateNetwork: false, Bandwidth: bandwidth, Type: "bulk-ethernet-networkV2"}

	er = ovc.CreateEthernetNetwork(ethernetNetwork)
	if er != nil {
		fmt.Println("............... Ethernet Network Creation Failed:", err)
	} else {
		fmt.Println(".... Ethernet Network Created Success")
	}

	err = ovc.CreateBulkEthernetNetwork(bulkEthernetNetwork)
	if err != nil {
		fmt.Println("............. Bulk Ethernet Network Creation Failed:", err)
	} else {
		fmt.Println(".... Bulk Ethernet Network Created Success")
	}

	bulk_ethernet_network_list, err := ovc.GetEthernetNetworks("", "", "", sort)
	for i := 0; i < len(bulk_ethernet_network_list.Members); i++ {
		fmt.Println(i, bulk_ethernet_network_list.Members[i].Name, bulk_ethernet_network_list.Members[i].URI)
	}

	ethernet_ntw, _ := ovc.GetEthernetNetworkByName(ethernet_network_1)
	ethernet_ntw.Name = ethernet_network_2
	err = ovc.UpdateEthernetNetwork(ethernet_ntw)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#.................... Ethernet Network after Updating ...........#")
		ethernet_nw_after_update, err := ovc.GetEthernetNetworks("", "", "", sort)
		if err != nil {
			fmt.Println(err)
		} else {
			for i := 0; i < len(ethernet_nw_after_update.Members); i++ {
				fmt.Println(i, ethernet_nw_after_update.Members[i].Name, ethernet_nw_after_update.Members[i].URI)
			}
		}
	}

	ethernet_ntw_0, _ := ovc.GetEthernetNetworkByName(ethernet_network)
	err = ovc.DeleteEthernetNetwork(ethernet_ntw_0.Name)
	err = ovc.DeleteEthernetNetwork(ethernet_network_2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#...................... Deleted Ethernet Network Successfully .....#")
	}

	ethernet_nw_list, err = ovc.GetEthernetNetworks("", "", "", sort)
	ethernet_ntw2, _ := ovc.GetEthernetNetworkByName("Test_eth_2")
	ethernet_ntw3, _ := ovc.GetEthernetNetworkByName("Test_eth_3")
	ethernet_ntw4, _ := ovc.GetEthernetNetworkByName("Test_eth_4")

	network_uris := &[]utils.Nstring{ethernet_ntw2.URI, ethernet_ntw3.URI, ethernet_ntw4.URI}

	bulkDeleteEthernetNetwork := ov.BulkDelete{NetworkUris: *network_uris}
	err = ovc.DeleteBulkEthernetNetwork(bulkDeleteEthernetNetwork)

	if err != nil {
		fmt.Println("............. Ethernet Network Bulk-Deletion Failed:", err)
	} else {
		fmt.Println("....  Ethernet Network Bulk-Delete is Successful")
	}
}
