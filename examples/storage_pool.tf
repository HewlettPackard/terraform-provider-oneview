provider "oneview" {
        ov_username = "<ov-username>"
        ov_password = "<ov-password>"
        ov_endpoint = "<ov-endpoint>"
        ov_sslverify = false
        ov_apiversion = <ov-apiversion>
        ov_ifmatch = "*"
}

// Resource updation. Make sure the storage pool to be updated is first imported in terraform before performing the update
// Use `terraform import oneview_storage_pool.<instance name> <name of the resource>
// Eg. terraform import oneview_storage_pool.storage_pool CPG-SSD
resource "oneview_storage_pool" "storage_pool" {
	name = "CPG-SSD"
	is_managed = false
}

// Data source example
/*
data "oneview_storage_pool" "storage_pool" {
        name = "Cluster-1"
}

output "oneview_storage_pool_value" {
        value = "${data.oneview_storage_pool.storage_pool.uri}"
}
*/
