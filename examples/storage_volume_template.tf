provider "oneview" {
  ov_username   = "<ov_username>"
  ov_password   = "<ov_password>"
  ov_endpoint   = "<ov_endpoint>"
  ov_sslverify  = false
  ov_apiversion =<ov_apiversion>
  ov_ifmatch    = "*"
}

data "oneview_scope" "scope_obj" {
        name = "test"
}

data "oneview_storage_pool" "st_pool" {
        name = "CPG_FC-AO"
}

//Creating a storage volume template
resource "oneview_storage_volume_template" "svt" {
        name = "DemoStorageTemplate"
        description = "Testing creation of storage volume template"
        root_template_uri = "/rest/storage-volume-templates/8b1d18f7-3bce-4cd4-80ae-ac2700f27c16"
        initial_scope_uris = ["${data.oneview_scope.scope_obj.uri}"]        
        tp_name =[
	{
                meta_locked = false
                type = "string"
                title = "Volume name"
                required = true
                max_length = 100
                min_length = 1
                description = "A volume name between 1 and 100 characters"
        }]
	tp_size = [
	{
                meta_locked = false
                meta_semantic_type = "capacity"
                type = "integer"
                title = "Capacity"
                default = 268435456
                required = true
                maximum = 1368744177664
                minimum = 268435456 
                description = "The capacity of the volume in bytes"
        }]
        tp_description = [
	{
                meta_locked = false
                type = "string"
                title = "Description"
                default = "A description for the volume"
                required = false
                max_length = 2000
		min_length = 1
                description = "A description for the volume"
        }]
        tp_is_shareable = [
	{
                meta_locked = false
                type = "boolean"
                title = "Is Shareable"
                default = false
                required = false
                description = "The shareability of the volume"
        }]
        tp_storage_pool = [
	{
                meta_locked = false
                meta_semantic_type = "device-storage-pool"
		meta_create_only = true
                type = "string"
                title = "Storage Pool"
                format = "x-uri-reference"
                default = "${data.oneview_storage_pool.st_pool.uri}"
                required = true
                description = "A common provisioning group URI reference"
        }]
	tp_snapshot_pool = [
	{
                meta_locked = true
                meta_semantic_type = "device-snapshot-storage-pool"
                type = "string"
                title = "Snapshot Pool"
                format = "x-uri-reference"
                default = "${data.oneview_storage_pool.st_pool.uri}"
                required = false
                description = "A URI reference to the common provisioning group used to create snapshots"
        }]
        tp_is_deduplicated = [
	{
                meta_locked = true
                type = "boolean"
                title = "Is Deduplicated"
                default = false
                required = false
                description = "Enables or disables deduplication of the volume"
        }]
        tp_template_version = [
	{
                meta_locked = true
                type = "string"
                title = "Template version"
                default = "1.1"
                required = true
                description = "Version of the template"
        }]
        tp_provisioning_type = [
	{
                meta_locked = true
                meta_create_only = true
                enum = ["Thin","Full"]
                type = "string"
                title = "Provisioning Type"
                default = "Thin"
                required = false
                description = "The provisioning type for the volume"
        }]

}

//Update the storage volume template
/*resource "oneview_storage_volume_template" "svt" {
        name = "RenameDemoStorageTemplate"
        description = "Testing update of storage volume template"
        root_template_uri = "/rest/storage-volume-templates/96196d4c-3cac-4d6b-ab6b-a93c0143ac75"
	family = "StoreServ"
	version = "2.0"
	storage_pool_uri = "/rest/storage-pools/547F8659-BD66-4775-9943-A93C0143AC70"
        tp_name =[
	{
                meta_locked = false
                type = "string"
                title = "Volume name"
                required = true
                max_length = 100
                min_length = 1
                description = "A volume name between 1 and 100 characters"
        }]
	tp_size = [
	{
                meta_locked = false
                meta_semantic_type = "capacity"
                type = "integer"
                title = "Capacity"
                default = 268435456
                required = true
                maximum = 70368744177664
                minimum = 268435456 
                description = "The capacity of the volume in bytes"
        }]
        tp_description = [
	{
                meta_locked = false
                type = "string"
                title = "Description"
                default = "A description for the volume"
                required = false
                max_length = 2000
		min_length = 1
                description = "A description for the volume"
        }]
        tp_is_shareable = [
	{
                meta_locked = false
                type = "boolean"
                title = "Is Shareable"
                default = false
                required = false
                description = "The shareability of the volume"
        }]
        tp_storage_pool = [
	{
                meta_locked = false
                meta_semantic_type = "device-storage-pool"
		meta_create_only = true
                type = "string"
                title = "Storage Pool"
                format = "x-uri-reference"
                default = "/rest/storage-pools/547F8659-BD66-4775-9943-A93C0143AC70"
                required = true
                description = "A common provisioning group URI reference"
        }]
	tp_snapshot_pool = [
	{
                meta_locked = true
                meta_semantic_type = "device-snapshot-storage-pool"
                type = "string"
                title = "Snapshot Pool"
                format = "x-uri-reference"
                default = "/rest/storage-pools/547F8659-BD66-4775-9943-A93C0143AC70"
                required = false
                description = "A URI reference to the common provisioning group used to create snapshots"
        }]
        tp_is_deduplicated = [
	{
                meta_locked = true
                type = "boolean"
                title = "Is Deduplicated"
                default = false
                required = false
                description = "Enables or disables deduplication of the volume"
        }]
        tp_template_version = [
	{
                meta_locked = true
                type = "string"
                title = "Template version"
                default = "2.0"
                required = true
                description = "Version of the template"
        }]
        tp_provisioning_type = [
	{
                meta_locked = true
                meta_create_only = true
                enum = ["Thin","Full"]
                type = "string"
                title = "Provisioning Type"
                default = "Thin"
                required = false
                description = "The provisioning type for the volume"
        }]
}*/

/*
// Importing an existing resource from the appliance
resource "oneview_storage_volume_template" "st" {
}
*/
//Tetsing data source
/*data "oneview_storage_volume_template" "d_svt" {
  name = "vt"
}

output "oneview_svt_value" {
  value = "${data.oneview_storage_volume_template.d_svt.root_template_uri}"
}*/
