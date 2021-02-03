package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
	"strconv"
)

func main() {
	var (
		clientOV *ov.OVClient
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

	sort := "name:desc"
	interconnect_list, err := ovc.GetInterconnects("", "", "", sort)
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("#----------------Interconnect List---------------#")

		for i := 0; i < len(interconnect_list.Members); i++ {
			fmt.Println(interconnect_list.Members[i].Name)
		}

		interconnect, err := ovc.GetInterconnectByName(interconnect_list.Members[0].Name)
		if err != nil {
			fmt.Println(err)
		} else {

			fmt.Println("#-------------Interconnect by Name----------------#")
			fmt.Println(interconnect.Name)

			uri := interconnect.URI
			interconnect, err = ovc.GetInterconnectByUri(uri)
			if err != nil {
				fmt.Println(err)
			} else {

				fmt.Println("#----------------Interconnect by URI--------------#")
				fmt.Println(interconnect.Name)
			}
		}
	}
}
