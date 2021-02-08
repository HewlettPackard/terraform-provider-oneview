package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
	"strconv"
)

func main() {
	var (
		ClientOV     *ov.OVClient
		networkset_2 = "updatednetworkset"
		networkset_3 = "creatednetworkset"
	)

	apiversion, _ := strconv.Atoi(os.Getenv("ONEVIEW_APIVERSION"))
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		apiversion,
		"*")
	ovVer, _ := ovc.GetAPIVersion()
	fmt.Println(ovVer)

	networkUris := new([]utils.Nstring)

	// Append all your networks to networkUris
	ethernetNetwork := ov.EthernetNetwork{Name: "test_eth-1", VlanId: 9, Purpose: "General", SmartLink: false, PrivateNetwork: false, ConnectionTemplateUri: "", EthernetNetworkType: "Tagged", Type: "ethernet-networkV4"}
	_ = ovc.CreateEthernetNetwork(ethernetNetwork)
	ethernet_ntw1, _ := ovc.GetEthernetNetworkByName("test_eth-1")

	*networkUris = append(*networkUris, ethernet_ntw1.URI)

	NetworkSet := ov.NetworkSet{Name: networkset_3,
		NativeNetworkUri:      "",
		NetworkUris:           *networkUris,
		ConnectionTemplateUri: "",
		Type:                  "network-setV5",
		NetworkSetType:        "Large",
	}
	err := ovc.CreateNetworkSet(NetworkSet)
	if err != nil {
		fmt.Println("............... NetworkSet Creation Failed:", err)
	} else {
		fmt.Println(".... NetworkSet Created Success.......")
	}

	fmt.Println("#...................NetworkSet by Name ...............#")
	net_set, err := ovc.GetNetworkSetByName(networkset_3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(net_set)
	}

	sort := "name:desc"
	networkset_list, err := ovc.GetNetworkSets("", sort)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("# ................... NetworkSet List .................#")
		for i := 0; i < len(networkset_list.Members); i++ {
			fmt.Println(networkset_list.Members[i].Name)
		}
	}

	net_set, err = ovc.GetNetworkSetByName(networkset_3)
	net_set.Name = networkset_2
	err = ovc.UpdateNetworkSet(net_set)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#.................... NetworkSet after Updating ...........#")
		networksets_after_update, err := ovc.GetNetworkSets("", sort)
		if err != nil {
			fmt.Println(err)
		} else {
			for i := 0; i < len(networksets_after_update.Members); i++ {
				fmt.Println(networksets_after_update.Members[i].Name)
			}
		}
	}

	err = ovc.DeleteNetworkSet(networkset_2)
	_ = ovc.DeleteEthernetNetwork("test_eth-1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#...................... Deleted Network Set Successfully .....#")
	}

}
