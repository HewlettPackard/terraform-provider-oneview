package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
	"strconv"
)

func main() {
	var (
		ClientOV              *ov.OVClient
		existing_interconnect = "Virtual Connect SE 40Gb F8 Module for Synergy"
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

	fmt.Println("#................... Interconnect Type by Name ...............#")
	interconnect, err := ovc.GetInterconnectTypeByName(existing_interconnect)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(interconnect)
	}

	fmt.Println("#................... Interconnect Type by Uri ....................#")
	int_uri, err := ovc.GetInterconnectTypeByUri(interconnect.URI)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(int_uri)
	}

	sort := "name:desc"
	interconnect_type_list, err := ovc.GetInterconnectTypes("", "", "", sort)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("# ................... Interconnect Type List .................#")
		for i := 0; i < len(interconnect_type_list.Members); i++ {
			fmt.Println(interconnect_type_list.Members[i].Name)
		}
	}
}
