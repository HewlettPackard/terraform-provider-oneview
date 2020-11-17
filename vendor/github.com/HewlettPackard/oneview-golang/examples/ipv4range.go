package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
	"strconv"
)

func main() {
	var (
		ClientOV *ov.OVClient
		id       = "a257c58c-bbe9-4174-b2a3-eada622fc555"
	)
	apiversion, _ := strconv.Atoi("ONEVIEW_APIVERSION")
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		apiversion,
		"*")
	fragments := new([]ov.StartStopFragments)
	fragment1 := ov.StartStopFragments{EndAddress: "10.16.0.100", StartAddress: "10.16.0.10"}
	*fragments = append(*fragments, fragment1)
	ipV4Range := ov.Createipv4Range{
		Type:               "Range",
		Name:               "test",
		StartStopFragments: *fragments,
		SubnetUri:          "/rest/id-pools/ipv4/subnets/d1f095c9-014c-43db-a65d-8236aea70b21",
	}

	// Gets an IPv4 range
	ipv4Range_by_id, err := ovc.GetIPv4RangebyId(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------Get Ipv4Range by id----------------#")
		fmt.Println(ipv4Range_by_id)
	}

	// Creates an IPv4 range
	fmt.Println(ipV4Range)
	err = ovc.CreateIPv4Range(ipV4Range)
	if err != nil {
		fmt.Println("IPv4 Range Creation Failed: ", err)
	} else {
		fmt.Println("IPv4 Range created successfully...")
	}

	// Gets all allocated fragments in an IPv4 range
	allocatedFragments, err := ovc.GetAllocatedFragments("", "", "", "", id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Get allocatedFragments ----#")
		for i := range allocatedFragments.Members {
			fmt.Println(allocatedFragments.Members[i])
		}
	}

	// Gets all free fragments in an IPv4 range
	freeFragments, err := ovc.GetFreeFragments("", "", "", "", id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Get freeGegments ----#")
		for i := range freeFragments.Members {
			fmt.Println(freeFragments.Members[i])
		}
	}

	// Perform either of the following operations on a Range i.e., Enable Range or Edit Range
	updateIpv4Range := ov.Updateipv4{Type: "Range", Enabled: false}
	err = ovc.UpdateIpv4Range("a257c58c-bbe9-4174-b2a3-eada622fc555", updateIpv4Range)
	if err != nil {
		panic(err)
	} else {
		if updateIpv4Range.Enabled == false {
			fmt.Println("Ipv4Range has disabled successfully ")
		} else {
			fmt.Println("Ipv4Range has enabled successfully")
		}
	}

	// Deletes an IPv4 range.
	err = ovc.DeleteIpv4Range(id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Deleted Ipv4Range successfully...")
	}
}
