provider "oneview" {
        ov_username = "<ov-username>"
        ov_password = "<ov-password>"
        ov_endpoint = "<ov-endpoint>"
        ov_sslverify = false
        ov_apiversion = <ov-apiversion>
        ov_ifmatch = "*"
}

resource "oneview_storage_system" "ss_inst" {
        hostname = "<hostname>"
        username = "<user>"
        password = "<password>"
        family   = "<family>"
}



/*
// Uncomment the following resource to update.
resource "oneview_storage_system" "ss_inst"{
        credentials = [
        {
		username = "<user>"
	        password = "<password>"
        }
        ]
        hostname = "<hostname>"
        name = "<name>"
        ports = [
        {
                mode = "<mode>"
                id= "<id>"
        }
        ]
        managed_pool = []
        eTag = "--"
        description = "TestStorageSystem"
        uri = "/rest/storage-systems/<id>"
}
*/

/* Testing the data source
data "oneview_storage_system" "storage_system" {
        name = "Cluster-2"
}

output "oneview_storage_system_value" {
        value = "${data.oneview_storage_system.storage_system.uri}"
}*/
