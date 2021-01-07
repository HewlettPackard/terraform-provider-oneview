package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
	"strconv"
)

func main() {
	var (
		clientOV           *ov.OVClient
		new_enclosure_name = "Renamed_Enclosure"
		path               = "/name"
		op                 = "replace"
		//		eg_name            = "Auto-EnclosureGroup"
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

	/*	enc_grp, _ := ovc.GetEnclosureGroupByName(eg_name)
		enclosure_create_map := ov.EnclosureCreateMap{
			EnclosureGroupUri: enc_grp.URI,
			Hostname:          os.Getenv("ENCLOSURE_HOSTNAME"),
			Username:          os.Getenv("ENCLOSURE_USERNAME"),
			Password:          os.Getenv("ENCLOSURE_PASSWORD"),
			LicensingIntent:   "OneView",
			InitialScopeUris:  make([]string, 0),
		}

		fmt.Println("#----------------Create Enclosure---------------#")

		err := ovc.CreateEnclosure(enclosure_create_map)
		if err != nil {
			fmt.Println("Enclosure Creation Failed: ", err)
		} else {
			fmt.Println("Enclosure created successfully...")
		}
	*/
	sort := ""

	enc_list, err := ovc.GetEnclosures("", "", "", sort, "")
	if err != nil {
		fmt.Println("Enclosure Retrieval Failed: ", err)
	} else {
		fmt.Println("#----------------Enclosure List---------------#")

		for i := 0; i < len(enc_list.Members); i++ {
			fmt.Println(enc_list.Members[i].Name)
		}
	}

	enclosure, err := ovc.GetEnclosureByName(enc_list.Members[0].Name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#----------------Enclosure by Name----------------#")
		fmt.Println(enclosure.Name)
	}

	uri := enclosure.URI
	enclosure, err = ovc.GetEnclosurebyUri(uri)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#----------------Enclosure by URI--------------#")
		fmt.Println(enclosure.Name)
	}

	err = ovc.UpdateEnclosure(op, path, new_enclosure_name, enclosure)
	if err != nil {
		fmt.Println(err)
	}

	enc_list, err = ovc.GetEnclosures("", "", "", sort, "")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#----------------Enclosure List after Updating---------#")
		for i := 0; i < len(enc_list.Members); i++ {
			fmt.Println(enc_list.Members[i].Name)
		}
	}

	/*	err = ovc.DeleteEnclosure(new_enclosure_name)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Deleted Enclosure successfully...")
		}
	*/
}
