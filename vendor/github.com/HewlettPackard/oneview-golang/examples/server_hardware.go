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
		server_1 = "<server_hardware_name>"
		server_2 = "<server_hardware_name>"
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

	fmt.Println("Get server hardware list by name")
	serverName, err := ovc.GetServerHardwareByName(server_1)
	if err != nil {
		fmt.Println("Failed to fetch server hardware name: ", err)
	} else {

		fmt.Println("******************")
		fmt.Println("Server hardware details ")
		fmt.Println(serverName.Model)
		fmt.Println(serverName.URI)
		fmt.Println(serverName.UUID)

		fmt.Println("******************")
		fmt.Println("Server hardware IloIPAddress ")
		fmt.Println(serverName.GetIloIPAddress())

		fmt.Println("******************")
		fmt.Println("Server hardware MpIPAddresses ")
		if ovc.IsHardwareSchemaV2() {
			for i := 0; i < len(serverName.MpHostInfo.MpIPAddresses); i++ {
				fmt.Println(serverName.MpHostInfo.MpIPAddresses[i].Address)
			}
		}
	}
	fmt.Println("Get server hardware list by url")

	fmt.Println("******************")

	ServerId, err := ovc.GetServerHardwareByUri(serverName.URI)
	if err == nil {
		fmt.Println(ServerId.URI)
	} else {
		fmt.Println("Failed to fetch server hardware : ", err)
	}

	fmt.Println("Get server-hardware list statistics specifying parameters")
	fmt.Println("******************")

	filters := []string{"name matches 'SY 660 Gen9 2'"}
	ServerList, err := ovc.GetServerHardwareList(filters, "", "", "", "")
	if err == nil {
		fmt.Println("Total server list :", ServerList.Total)
	} else {
		fmt.Println("Failed to fetch server List : ", err)
	}

	fmt.Println("Get firmware inventory for a server-hardware")
	fmt.Println("******************")
	firmware, err := ovc.GetServerFirmwareByUri(ServerId.URI)

	if err == nil {
		fmt.Println(firmware.ServerName)
	} else {
		fmt.Println("Failed to fetch Firmwares: ", err)
	}

	fmt.Println("******************")
	server2, err := ovc.GetServerHardwareByName(server_1)
	server1, err := ovc.GetServerHardwareByName(server_2)
	serverHarware, err := ovc.GetAvailableHardware(string(server2.URI), string(server1.URI))

	if err == nil {
		fmt.Println(serverHarware.Type)
	} else {
		fmt.Println("Failed to fetch server hardware by URI: ", err)
	}

	fmt.Println("Get power status of a server ")
	fmt.Println("******************")
	powerState, err := serverName.GetPowerState()
	if err == nil {
		fmt.Println("Power state of the machine is ", powerState)
	} else {
		fmt.Println("Failed to fetch powerstate of the server: ", err)
	}
	//Trying to touggle power state of the server using put
	if powerState.String() == "On" {
		fmt.Println("Server is in powered on state ")
		serverName.PowerOff()
		powerState, _ := serverName.GetPowerState()
		fmt.Println("Power state of the machine is ", powerState)

	} else {

		fmt.Println("Server is in powered off state ")
		serverName.PowerOn()
		powerState, _ := serverName.GetPowerState()
		fmt.Println("Power state of the machine is ", powerState)

	}

	fmt.Println("Get ilo ipaddress of a server ")
	fmt.Println("******************")
	iloIpaddress := serverName.GetIloIPAddress()
	fmt.Println("ilo ip address of an server is =", iloIpaddress)
}
