package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
	"path/filepath"
	"strconv"
)

func main() {

	var (
		ClientOV    *ov.OVClient
		name_to_get = "<vol_attachment_name>"
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

	// Get All the attachments present
	fmt.Println("\nGetting all the storage attachments present in the system: \n")
	sort := "name:desc"
	attachment_list, err := ovc.GetStorageAttachments("", sort, "", "")
	if err != nil {
		fmt.Println("Error Getting the storage attachments ", err)
	}
	for i := 0; i < len(attachment_list.Members); i++ {
		fmt.Println(attachment_list.Members[i].URI)
	}
	if len(attachment_list.Members) != 0 {
		id := filepath.Base(string(attachment_list.Members[0].URI))
		// Get volume attachment by id
		fmt.Println("\nGetting details of volume attachment with id: ", id)
		volAttach_by_id, err_id := ovc.GetStorageAttachmentById(id)
		if err_id != nil {
			fmt.Println("Error Getting the storage attachments ", err_id)
		}
		fmt.Println(volAttach_by_id)
	}
	// Get volume attachment by name
	fmt.Println("\nGetting details of volume attachment with name: ", name_to_get)
	volAttach_by_name, err_name := ovc.GetStorageAttachmentByName(name_to_get)
	if err_name != nil {
		fmt.Println("Error Getting the storage attachments ", err_name)
	}

	fmt.Println(volAttach_by_name)
}
