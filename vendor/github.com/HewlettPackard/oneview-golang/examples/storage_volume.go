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
		st_vol_template = "template"
		new_volume      = "TestVolume"
		name_to_update  = "UpdatedName"
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

	// Create storage volume with name <new_volume>
	st_pool, _ := ovc.GetStoragePoolByName("CPG-SSD")
	properties := &ov.Properties{
		Name:             new_volume,
		Storagepool:      st_pool.URI,
		Size:             268435456,
		ProvisioningType: "Thin",
		//	DataProtectionLevel: "NetworkRaid10Mirror2Way",
	}
	trueVal := true
	vol_template, err := ovc.GetStorageVolumeTemplateByName(st_vol_template)
	if err != nil {
		fmt.Println(err)
	}
	storageVolume := ov.StorageVolume{TemplateURI: vol_template.URI, Properties: properties, IsPermanent: &trueVal}

	err = ovc.CreateStorageVolume(storageVolume)
	if err != nil {
		fmt.Println("Could not create the volume", err)
	}

	// Update the given storage volume
	update_vol, _ := ovc.GetStorageVolumeByName(new_volume)

	updated_storage_volume := ov.StorageVolume{
		ProvisioningTypeForUpdate: update_vol.ProvisioningTypeForUpdate,
		IsPermanent:               update_vol.IsPermanent,
		IsShareable:               update_vol.IsShareable,
		Name:                      name_to_update,
		ProvisionedCapacity:       "107374741824",
		DeviceSpecificAttributes:  update_vol.DeviceSpecificAttributes,
		URI:                       update_vol.URI,
		ETAG:                      update_vol.ETAG,
		Description:               "empty",
		TemplateVersion:           "1.1",
		VolumeTemplateUri:         update_vol.VolumeTemplateUri,
	}

	err = ovc.UpdateStorageVolume(updated_storage_volume)
	if err != nil {
		fmt.Println("Could not update the volume", err)
	}

	// Get All the volumes present
	fmt.Println("\nGetting all the storage volumes present in the system: \n")
	sort := "name:desc"
	vol_list, err := ovc.GetStorageVolumes("", sort)
	if err != nil {
		fmt.Println("Error Getting the storage volumes ", err)
	}
	for i := 0; i < len(vol_list.Members); i++ {
		fmt.Println(vol_list.Members[i].Name)
	}

	// Get volume by name
	fmt.Println("\nGetting details of volume with name: ", name_to_update)
	vol_by_name, _ := ovc.GetStorageVolumeByName(name_to_update)
	fmt.Println(vol_by_name)

	// Delete the created volume
	fmt.Println("\nDeleting the volume with name : UpdatedName")
	err = ovc.DeleteStorageVolume(name_to_update)
	if err != nil {
		fmt.Println("Delete Unsuccessful", err)
	}
}
