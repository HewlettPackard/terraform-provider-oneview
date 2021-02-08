package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
	"strconv"
	//	"encoding/json"
)

func main() {
	var (
		ClientOV *ov.OVClient
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

	sort := "name:desc"
	interconnect_type_list, err := ovc.GetInterconnectTypes("", "", "", sort)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("# ................... Interconnect Type List .................#")
		for i := 0; i < len(interconnect_type_list.Members); i++ {
			fmt.Println(interconnect_type_list.Members[i].Name)
			fmt.Println(interconnect_type_list.Members[i].URI)
		}
	}
	/*
		fmt.Println("#................... Interconnect Type by Name ...............#")
		interconnect, err := ovc.GetInterconnectTypeByName(string(interconnect_type_list.Members[0].Name))
		if err != nil {
			fmt.Println(err)
		} else {
			interconnect, _ := json.MarshalIndent(interconnect, "", "\t");
			fmt.Print(string(interconnect))
		}

		fmt.Println("#................... Interconnect Type by Uri ....................#")
		int_uri, err := ovc.GetInterconnectTypeByUri(interconnect.URI)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(int_uri)
		}*/
}
