package oneview

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceStorageVolume() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceStorageVolumeRead,
		Schema: map[string]*schema.Schema{
			"category": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"eTag": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"created": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_volume_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"requesting_refresh": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allocated_capacity": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"initial_scopes": {
				Type:     schema.TypeString,
				Required: true,
			},
			"device_specific_attributes": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"transport": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"iqn": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"number_of_replicas": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"data_protection_level": {
							Type:     schema.TypeString,
							Required: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"copy_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"snapshot_pool_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"volume_template_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_shareable": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"storage_pool_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"storage_system_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"provisioned_capacity": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"properties": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"storage_pool": {
							Type:     schema.TypeString,
							Required: true,
						},
						"size": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"provisioning_type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"template_uri": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_permanent": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"provisioning_type_for_update": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceStorageVolumeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	storageVolume, err := config.ovClient.GetStorageVolumeByName(string(d.Id()))
	if err != nil {
		d.SetId("")
		return nil
	}

	d.SetId(d.Id())
	d.Set("category", storageVolume.Category)
	d.Set("eTag", storageVolume.ETAG)
	d.Set("created", storageVolume.Created)
	d.Set("description", storageVolume.Description)
	d.Set("name", storageVolume.Name)
	d.Set("state", storageVolume.State)
	d.Set("status", storageVolume.Status)
	d.Set("type", storageVolume.Type)
	d.Set("uri", storageVolume.URI)
	d.Set("deviceVolumeName", storageVolume.DeviceVolumeName)
	d.Set("requestingRefresh", storageVolume.RequestingRefresh)
	d.Set("allocatedCapacity", storageVolume.AllocatedCapacity)
	d.Set("initialScopeUris", storageVolume.InitialScopeUris)
	deviceSpecificAttributes := make([]map[string]interface{}, 0)
	deviceSpecificAttributes = append(deviceSpecificAttributes, map[string]interface{}{
		"transport":             storageVolume.DeviceSpecificAttributes.Transport,
		"iqn":                   storageVolume.DeviceSpecificAttributes.Iqn,
		"numberOfReplicas":      storageVolume.DeviceSpecificAttributes.NumberOfReplicas,
		"data_protection_Level": storageVolume.DeviceSpecificAttributes.DataProtectionLevel,
		"id":                    storageVolume.DeviceSpecificAttributes.Id,
		"uri":                   storageVolume.DeviceSpecificAttributes.Uri,
		"copyState":             storageVolume.DeviceSpecificAttributes.CopyState,
		"snapshotPoolUri":       storageVolume.DeviceSpecificAttributes.SnapshotPoolUri})
	d.Set("deviceSpecificAttributes", deviceSpecificAttributes)
	d.Set("volumeTemplateUri", storageVolume.VolumeTemplateUri)
	d.Set("isSheareable", storageVolume.IsShareable)
	d.Set("storagePoolUri", storageVolume.StoragePoolUri)
	d.Set("storageSystemUri", storageVolume.StorageSystemUri)
	d.Set("provisionCapacity", storageVolume.ProvisionedCapacity)
	properties := make([]map[string]interface{}, 0)
	properties = append(properties, map[string]interface{}{
		"name":             storageVolume.Properties.Name,
		"storagePool":      storageVolume.Properties.Storagepool,
		"size":             storageVolume.Properties.Size,
		"provisioningType": storageVolume.Properties.ProvisioningType})
	d.Set("properties", properties)
	d.Set("templateUri", storageVolume.TemplateURI)
	d.Set("isPermanent", storageVolume.IsPermanent)
	d.Set("provisioningType", storageVolume.ProvisioningTypeForUpdate)

	return nil
}
