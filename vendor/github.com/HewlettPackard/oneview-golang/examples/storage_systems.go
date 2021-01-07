package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
	"strconv"
)

func main() {

	var (
		ClientOV        *ov.OVClient
		name2_to_create = "ThreePAR-2"
		name_to_create  = "ThreePAR-1"
		managed_domain  = "TestDomain" //Variable to update the managedDomain
		username        = "<storage_password>"
		password        = "<storage_username>"
		host_ip         = "<storage_IP>"
		host2_ip        = "<another_Storage_IP>"
		family          = "StoreServ"
		//		description    = ""
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

	// Create storage system
	storageSystem := ov.StorageSystem{Hostname: host_ip, Username: username, Password: password, Family: family}
	storageSystem_2 := ov.StorageSystem{Hostname: host2_ip, Username: username, Password: password, Family: family}
	err := ovc.CreateStorageSystem(storageSystem)
	err = ovc.CreateStorageSystem(storageSystem_2)
	if err != nil {
		fmt.Println("Could not create the system", err)
	}

	//The below example is to update a storeServ system.
	//Please refer the API reference for fields to update a storeVirtual system.

	// Get the storage system to be updated
	update_system, _ := ovc.GetStorageSystemByName(name_to_create)

	// Update the given storage system
	//Managed domain is mandatory attribute for update
	DeviceSpecificAttributesForUpdate := update_system.StorageSystemDeviceSpecificAttributes
	DeviceSpecificAttributesForUpdate.ManagedDomain = managed_domain

	updated_storage_system := ov.StorageSystem{
		Name:                                  name_to_create,
		StorageSystemDeviceSpecificAttributes: DeviceSpecificAttributesForUpdate,
		URI:                                   update_system.URI,
		ETAG:                                  update_system.ETAG,
		Description:                           "Updated the storage system",
		Credentials:                           update_system.Credentials,
		Hostname:                              update_system.Hostname,
		Ports:                                 update_system.Ports,
	}

	err = ovc.UpdateStorageSystem(updated_storage_system)
	if err != nil {
		fmt.Println("Could not update the system", err)
	}

	// Get All the systems present
	fmt.Println("\nGetting all the storage systems present in the appliance: \n")
	sort := "name:desc"
	system_list, err := ovc.GetStorageSystems("", sort)
	if err != nil {
		fmt.Println("Error Getting the storage systems ", err)
	} else {
		for i := 0; i < len(system_list.Members); i++ {
			fmt.Println(system_list.Members[i].Name)
		}
	}

	// Get reachable ports
	fmt.Println("\n Getting rechable ports of:", name_to_create)
	reachable_ports, _ := ovc.GetReachablePorts(update_system.URI)
	fmt.Println(reachable_ports.Members)

	// Get volume sets
	fmt.Println("\n Getting volume sets of:", name_to_create)
	volume_sets, _ := ovc.GetVolumeSets(update_system.URI)
	fmt.Println(volume_sets.Members)

	// Delete the created system
	fmt.Println("\nDeleting the system with name : ", name2_to_create)
	err = ovc.DeleteStorageSystem(name2_to_create)
	if err != nil {
		fmt.Println("Delete Unsuccessful", err)
	}
}
