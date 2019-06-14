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
        username = "<username>"
        password = "<password>"
        family   = "StoreServ"
}

// Uncomment the following resource to update.
/*resource "oneview_storage_system" "ss_inst"{
        credentials = [
        {
                username = "<username>"
                password = "<password>"
        }
        ]
        hostname = "<hostname>"
        name = "ThreePAR-2"
        ports = [
        {
                id = "ea0b2d3d-098c-4f95-ac08-aa6100a80de7"
                mode = "AutoSelectExpectedSan"
                partner_port = "1:1:1"
        },
        {
                id = "5a1469d8-a925-4b9a-a87b-aa6100a80de7"
                mode = "AutoSelectExpectedSan"
                partner_port = "1:1:2"
        },
        {
                id = "4f55584b-2abb-47f5-95cc-aa6100a80de7"
                mode = "AutoSelectExpectedSan"
                partner_port = "1:1:3"
        },
	]
	managed_pool = []
        storage_system_device_specific_attributes = [
        {
                managed_domain = "TestDomain"
                firmware = "3.2.1.292"
                model = "HP_3PAR 7200"
        }
        ]
        eTag = "10"
        description = "TestStorageSystem"
        uri = "/rest/storage-systems/TXQ1010307"
}*/

/* Testing the data source
data "oneview_storage_system" "storage_system" {
        name = "Cluster-2"
}

output "oneview_storage_system_value" {
        value = "${data.oneview_storage_system.storage_system.uri}"
}*/
