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
		testName    = "TestFCNetworkGOsdk"
		new_fc_name = "RenamedFCNetwork"
		falseVar    = false
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
	scope := ov.Scope{Name: "ScopTest", Description: "Test from script", Type: "ScopeV3"}
	_ = ovc.CreateScope(scope)
	scp, _ := ovc.GetScopeByName("ScopTest")
	initialScopeUris := &[]utils.Nstring{scp.URI}

	fcNetwork := ov.FCNetwork{
		AutoLoginRedistribution: falseVar,
		Description:             "Test FC Network",
		LinkStabilityTime:       30,
		FabricType:              "FabricAttach",
		Name:                    testName,
		Type:                    "fc-networkV4",    //The Type value is for API>500.
		InitialScopeUris:        *initialScopeUris, //added for API>500
	}
	fmt.Println(fcNetwork)
	err := ovc.CreateFCNetwork(fcNetwork)
	if err != nil {
		fmt.Println("Fc Network Creation Failed: ", err)
	} else {
		fmt.Println("Fc Network created successfully...")
	}

	sort := "name:desc"
	fcNetworks, err := ovc.GetFCNetworks("", sort, "", "")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Get Fc Networks sorted by name in descending order----#")
		for i := range fcNetworks.Members {
			fmt.Println(fcNetworks.Members[i].Name)
		}
	}
	fcNetwork2, err := ovc.GetFCNetworkByName(testName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------Get FCNetworks by name----------------#")
		fmt.Println(fcNetwork2)
	}
	fcNetwork2.Name = new_fc_name
	err = ovc.UpdateFcNetwork(fcNetwork2)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("FCNetwork has been updated with name: " + fcNetwork2.Name)
	}
	err = ovc.DeleteFCNetwork(new_fc_name)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Deleted FCNetworks successfully...")
	}

	fcNetwork01 := ov.FCNetwork{
		AutoLoginRedistribution: falseVar,
		Description:             "Test FC Network 1",
		LinkStabilityTime:       30,
		FabricType:              "FabricAttach",
		Name:                    "testName1",
		Type:                    "fc-networkV4",
	}
	err = ovc.CreateFCNetwork(fcNetwork01)

	fcNetwork02 := ov.FCNetwork{
		AutoLoginRedistribution: falseVar,
		Description:             "Test FC Network 2",
		LinkStabilityTime:       30,
		FabricType:              "FabricAttach",
		Name:                    "testName2",
		Type:                    "fc-networkV4",
	}
	err = ovc.CreateFCNetwork(fcNetwork02)

	fcNetwork1, err := ovc.GetFCNetworkByName("testName1")
	fcNetwork2, err = ovc.GetFCNetworkByName("testName2")

	network_uris := &[]utils.Nstring{fcNetwork1.URI, fcNetwork2.URI}
	bulkDeleteFCNetwork := ov.FCNetworkBulkDelete{FCNetworkUris: *network_uris}
	err = ovc.DeleteScope("ScopTest")
	err = ovc.DeleteBulkFcNetwork(bulkDeleteFCNetwork)

	if err != nil {
		fmt.Println("............. FC Network Bulk-Deletion Failed:", err)
	} else {
		fmt.Println(".... FC Network Bulk-Delete is Successful")
	}

}
