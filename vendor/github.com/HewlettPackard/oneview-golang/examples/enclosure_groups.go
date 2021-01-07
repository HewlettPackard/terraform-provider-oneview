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
		clientOV    *ov.OVClient
		eg_name     = "TestEG"
		new_eg_name = "RenamedEnclosureGroup"
		//script      = "#TEST COMMAND"
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

	lig, _ := ovc.GetLogicalInterconnectGroupByName("Auto-LIG")

	ibMappings := new([]ov.InterconnectBayMap)
	interconnectBay1 := ov.InterconnectBayMap{InterconnectBay: 3, LogicalInterconnectGroupUri: lig.URI}
	interconnectBay2 := ov.InterconnectBayMap{InterconnectBay: 6, LogicalInterconnectGroupUri: lig.URI}

	*ibMappings = append(*ibMappings, interconnectBay1)
	*ibMappings = append(*ibMappings, interconnectBay2)

	scp, _ := ovc.GetScopeByName("Auto-Scope")
	initialScopeUris := new([]utils.Nstring)
	*initialScopeUris = append(*initialScopeUris, scp.URI)

	enclosureGroup := ov.EnclosureGroup{Name: eg_name, InterconnectBayMappings: *ibMappings, InitialScopeUris: *initialScopeUris, IpAddressingMode: "External", EnclosureCount: 3}
	enclosureGroup_auto := ov.EnclosureGroup{Name: "Auto-TestEG", InterconnectBayMappings: *ibMappings, InitialScopeUris: *initialScopeUris, IpAddressingMode: "External", EnclosureCount: 3}
	/*
		 This is used for C7000 only
		enclosureGroup := ov.EnclosureGroup{Name: eg_name, InitialScopeUris: *initialScopeUris, InterconnectBayMappings: *ibMappings}
	*/

	err := ovc.CreateEnclosureGroup(enclosureGroup)
	_ = ovc.CreateEnclosureGroup(enclosureGroup_auto)
	if err != nil {
		fmt.Println("Enclosure Group Creation Failed: ", err)
	} else {
		fmt.Println("Enclosure Group created successfully...")
	}

	sort := "name:desc"

	enc_grp_list, err := ovc.GetEnclosureGroups("", "", "", sort, "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("#----------------Enclosure Group List---------------#")

	for i := 0; i < len(enc_grp_list.Members); i++ {
		fmt.Println(enc_grp_list.Members[i].Name)
	}

	if ovc.APIVersion > 500 {
		scope_uri := string(scp.URI)
		enc_grp_list1, err := ovc.GetEnclosureGroups("", "", "", "", scope_uri)
		if err != nil {
			fmt.Println("Error in getting EnclosureGroups by scope URIs:", err)
		}
		fmt.Println("#-----------Enclosure Groups based on Scope URIs----------#")
		for i := 0; i < len(enc_grp_list1.Members); i++ {
			fmt.Println(enc_grp_list1.Members[i].Name)
		}
	}

	enc_grp, err := ovc.GetEnclosureGroupByName(eg_name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("#-------------Enclosure Group by name----------------#")
	fmt.Println(enc_grp)

	uri := enc_grp.URI
	enc_grp, err = ovc.GetEnclosureGroupByUri(uri)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("#----------------Enclosure Group by URI--------------#")
	fmt.Println(enc_grp)

	enc_grp.Name = new_eg_name
	err = ovc.UpdateEnclosureGroup(enc_grp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Enclosure Group got updated")
	}

	enc_grp_list, err = ovc.GetEnclosureGroups("", "", "", sort, "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("#----------------EnclosureList after updating---------#")
	for i := 0; i < len(enc_grp_list.Members); i++ {
		fmt.Println(enc_grp_list.Members[i].Name)
	}

	/* This method is only available on C7000
	update_script, err := ovc.UpdateConfigurationScript(enc_grp.URI, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Update Configuration Script result:", update_script)

	// This method is only available on C7000
	conf_script, err := ovc.GetConfigurationScript(enc_grp.URI)
	if err != nil {
		fmt.Println("Error in getting configuration Script: ", err)
	}
	fmt.Println("Configuation Script: ", conf_script)
	fmt.Println(script) */

	err = ovc.DeleteEnclosureGroup(new_eg_name)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted EnclosureGroup successfully...")
}
