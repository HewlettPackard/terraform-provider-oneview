provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

/*
# Extracting Scope 
data "oneview_scope" "scope_obj" {
        name = "Auto-Scope"
}

# Extracting Storage Pool
data "oneview_storage_pool" "st_pool" {
        name = "CPG-SSD-AO"
}

# Creating a storage volume template
resource "oneview_storage_volume_template" "svt" {
        name = "DemoStorageTemplate"
        description = "Testing creation of storage volume template"
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
                minimum = 268435456 
		maximum = 17592186044416
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
//		meta_semantic_type = "device-provisioningType"

        }]
//	tp_data_protection_level=[
//	{
//		meta_locked = false
//		meta_semantic_type =  "device-dataProtectionLevel"
//		enum = ["NetworkRaid0None",
//                        "NetworkRaid5SingleParity",
//                       "NetworkRaid10Mirror2Way",
//                       "NetworkRaid10Mirror3Way",
//                       "NetworkRaid10Mirror4Way",
//                       "NetworkRaid6DualParity"]
//       	type =      "string"
//               title =       "Data Protection Level"
//	        default =     "NetworkRaid10Mirror2Way"
//               description = "Indicates the number and configuration of data copies in the Storage Pool"
//               required = true
//       }]
//       tp_is_adaptive_optimization_enabled = [
//       {
//       	meta_locked = true
//       	description = ""
//       	default = true
//       	required = false
//       	title = "Adaptive Optimization"
//       	type = "boolean"
//	}]
}
*/

# Fetching Existing Template for update
data "oneview_storage_volume_template" "d_svt" {
  name = "DemoStorageTemplate"
}

# Update the storage volume template
resource "oneview_storage_volume_template" "svt" {
  name              = "RenameDemoStorageTemplate"
  description       = "Testing update of storage volume template"
  family            = data.oneview_storage_volume_template.d_svt.family
  storage_pool_uri  = data.oneview_storage_volume_template.d_svt.storage_pool_uri
  root_template_uri = data.oneview_storage_volume_template.d_svt.root_template_uri
  tp_name {
    meta_locked = false
    type        = "string"
    title       = "Volume name"
    required    = true
    max_length  = 100
    min_length  = 1
    description = "A volume name between 1 and 100 characters"
  }
  tp_size {
    meta_locked        = false
    meta_semantic_type = "capacity"
    type               = "integer"
    title              = "Capacity"
    default            = 268435456
    required           = true
    minimum            = 268435456
    maximum            = 17592186044416
    description        = "The capacity of the volume in bytes"
  }
  tp_description {
    meta_locked = false
    type        = "string"
    title       = "Description"
    default     = "A description for the volume"
    required    = false
    max_length  = 2000
    min_length  = 1
    description = "A description for the volume"
  }
  tp_is_shareable {
    meta_locked = false
    type        = "boolean"
    title       = "Is Shareable"
    default     = false
    required    = false
    description = "The shareability of the volume"
  }
  tp_storage_pool {
    meta_locked        = false
    meta_semantic_type = "device-storage-pool"
    meta_create_only   = true
    type               = "string"
    title              = "Storage Pool"
    format             = "x-uri-reference"
    default            = data.oneview_storage_volume_template.d_svt.storage_pool_uri
    required           = true
    description        = "A common provisioning group URI reference"
  }
  tp_snapshot_pool {
    meta_locked        = true
    meta_semantic_type = "device-snapshot-storage-pool"
    type               = "string"
    title              = "Snapshot Pool"
    format             = "x-uri-reference"
    default            = data.oneview_storage_volume_template.d_svt.storage_pool_uri
    required           = false
    description        = "A URI reference to the common provisioning group used to create snapshots"
  }
  tp_is_deduplicated {
    meta_locked = true
    type        = "boolean"
    title       = "Is Deduplicated"
    default     = false
    required    = false
    description = "Enables or disables deduplication of the volume"
  }
  tp_template_version {
    meta_locked = true
    type        = "string"
    title       = "Template version"
    default     = "1.1"
    required    = true
    description = "Version of the template"
  }
  tp_provisioning_type {
    meta_locked      = true
    meta_create_only = true
    enum             = ["Thin", "Full"]
    type             = "string"
    title            = "Provisioning Type"
    default          = "Thin"
    required         = false
    description      = "The provisioning type for the volume"
    //		meta_semantic_type = "device-provisioningType"
  }
  //	tp_data_protection_level=[
  //       {
  //       	meta_locked = false
  //       	meta_semantic_type =  "device-dataProtectionLevel"
  //       	enum = ["NetworkRaid0None",
  //                       "NetworkRaid5SingleParity",
  //                       "NetworkRaid10Mirror2Way",
  //                       "NetworkRaid10Mirror3Way",
  //                       "NetworkRaid10Mirror4Way",
  //                       "NetworkRaid6DualParity"]
  //       	type =      "string"
  //               title =       "Data Protection Level"
  //               default =     "NetworkRaid10Mirror2Way"
  //               description = "Indicates the number and configuration of data copies in the Storage Pool"
  //               required = true
  //       }]
  //       tp_is_adaptive_optimization_enabled = [
  //       {
  //       	meta_locked = true
  //       	description = ""
  //       	default = true
  //       	required = false
  //       	title = "Adaptive Optimization"
  //       	type = "boolean"
  //       }]
}

# Testing data source
/*
data "oneview_storage_volume_template" "d_svt" {
  name = "DemoStorageTemplate"
}

output "oneview_svt_value" {
  value = "${data.oneview_storage_volume_template.d_svt.root_template_uri}"
}
*/
# Importing an existing resource from the appliance.
/*
resource "oneview_storage_volume_template" "svt" {
}
*/
