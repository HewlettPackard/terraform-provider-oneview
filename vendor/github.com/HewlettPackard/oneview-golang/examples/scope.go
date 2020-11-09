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
		ClientOV    *ov.OVClient
		scp_name    = "ScopeTest"
		new_scope   = "new-scope"
		upd_scope   = "update-scope"
		eth_network = "testing"
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

	fmt.Println("#................... Scope by Name ...............#")
	scp, scperr := ovc.GetScopeByName(scp_name)
	if scperr != nil {
		fmt.Println(scperr)
	}
	fmt.Println(scp)

	sort := "name:desc"
	scp_list, err := ovc.GetScopes("", "", "", "", sort)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("# ................... Scopes List .................#")
	for i := 0; i < len(scp_list.Members); i++ {
		fmt.Println(scp_list.Members[i].Name)
	}
	eth_uri, err := ovc.GetEthernetNetworkByName(eth_network)
	if err != nil {
		fmt.Println(err)
	}

	initialScopeUris := &[]utils.Nstring{(scp.URI)}
	addedResourceUris := &[]utils.Nstring{(eth_uri.URI)}
	scope := ov.Scope{Name: new_scope, Description: "Test from script", Type: "ScopeV3", InitialScopeUris: *initialScopeUris, AddedResourceUris: *addedResourceUris}

	er := ovc.CreateScope(scope)
	if er != nil {
		fmt.Println("............... Scope Creation Failed:", err)
	} else {
		fmt.Println("# ................... Scope Created Successfully.................#")
	}

	new_scp, err := ovc.GetScopeByName(new_scope)
	if err != nil {
		fmt.Println(err)
	} else {
		new_scp.Name = upd_scope
		err = ovc.UpdateScope(new_scp)

		if err != nil {
			fmt.Println("#.................... Scope Updation failed ...........#")
			panic(err)
		} else {
			fmt.Println("#.................... Scope after Updating ...........#")
		}
	}
	up_list, err := ovc.GetScopes("", "", "", "", sort)
	for i := 0; i < len(up_list.Members); i++ {
		fmt.Println(up_list.Members[i].Name)
	}

	err = ovc.DeleteScope(upd_scope)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#...................... Deleted Scope Successfully .....#")
	}
	scp_list, err = ovc.GetScopes("", "", "", "", sort)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("# ................... Scopes List .................#")
	for i := 0; i < len(scp_list.Members); i++ {
		fmt.Println(scp_list.Members[i].Name)
	}
}
