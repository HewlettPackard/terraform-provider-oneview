package oneview

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceStorageVolume() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceStorageVolumeRead,
		Schema: map[string]*schema.Schema{
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"eTag": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: "Warning: Current value structure is deprecated",
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"device_volume_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"requesting_refresh": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"allocated_capacity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"initial_scopes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"device_specific_attributes": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"transport": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"iqn": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"number_of_replicas": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"data_protection_level": {
							Type:     schema.TypeString,
							Computed: true,
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
							Computed: true,
						},
					},
				},
			},
			"volume_template_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_shareable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"storage_pool_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_system_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"provisioned_capacity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"properties": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"storage_pool": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"provisioning_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"template_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_permanent": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"provisioning_type_for_update": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceStorageVolumeRead(d *schema.ResourceData, meta interface{}) error {
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
	d.Set("device_volume_name", storageVolume.DeviceVolumeName)
	d.Set("requesting_refresh", storageVolume.RequestingRefresh)
	d.Set("allocated_capacity", storageVolume.AllocatedCapacity)
	d.Set("initial_scope_uris", storageVolume.InitialScopeUris)
	deviceSpecificAttributes := make([]map[string]interface{}, 0)
	deviceSpecificAttributes = append(deviceSpecificAttributes, map[string]interface{}{
		"transport":             storageVolume.DeviceSpecificAttributes.Transport,
		"iqn":                   storageVolume.DeviceSpecificAttributes.Iqn,
		"number_of_replicas":    storageVolume.DeviceSpecificAttributes.NumberOfReplicas,
		"data_protection_level": storageVolume.DeviceSpecificAttributes.DataProtectionLevel,
		"id":                    storageVolume.DeviceSpecificAttributes.Id,
		"uri":                   storageVolume.DeviceSpecificAttributes.Uri,
		"copy_state":            storageVolume.DeviceSpecificAttributes.CopyState,
		"snapshot_pool_uri":     storageVolume.DeviceSpecificAttributes.SnapshotPoolUri})
	d.Set("device_specific_attributes", deviceSpecificAttributes)
	d.Set("volume_template_uri", storageVolume.VolumeTemplateUri)
	d.Set("is_sheareable", storageVolume.IsShareable)
	d.Set("storage_pool_uri", storageVolume.StoragePoolUri)
	d.Set("storage_system_uri", storageVolume.StorageSystemUri)
	d.Set("provision_capacity", storageVolume.ProvisionedCapacity)
	properties := make([]map[string]interface{}, 0)
	properties = append(properties, map[string]interface{}{
		"name":              storageVolume.Properties.Name,
		"storage_pool":      storageVolume.Properties.Storagepool,
		"size":              storageVolume.Properties.Size,
		"provisioning_type": storageVolume.Properties.ProvisioningType})
	d.Set("properties", properties)
	d.Set("template_uri", storageVolume.TemplateURI)
	d.Set("is_permanent", storageVolume.IsPermanent)
	d.Set("provisioning_type", storageVolume.ProvisioningTypeForUpdate)

	return nil
}
