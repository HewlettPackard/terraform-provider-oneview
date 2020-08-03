provider "oneview" {
  ov_username   = "<ov_username>"
  ov_password   = "<ov_password>"
  ov_endpoint   = "<ov_endpoint>"
  ov_sslverify  = false
  ov_apiversion = <ov_apiversion>
  ov_ifmatch    = "*"
}


resource "oneview_storage_system" "ss_inst" {
        hostname = "<storage_system_ip>"
        username = "<storage_system_username>"
        password = "<storage_system_password>"
        family   = "StoreServ"
}


// Uncomment the following resource to update.
/*resource "oneview_storage_system" "ss_inst" {
  credentials = [
    {
      username = "<storage_system_username>"
      password = "<storage_system_password>"
    },
  ]

  hostname = "<storage_system_ip>"
  name     = "ThreePAR-2"

  storage_system_device_specific_attributes = {
    managed_domain = "TestDomain"
  }

  eTag        = "--"
  description = "TestStorageSystem"
  uri         = "/rest/storage-systems/TXQ1010307"
}
*/
/*
// Testing the data source
data "oneview_storage_system" "ss_int" {
        name = "ThreePAR-2"
}

output "oneview_storage_system_value" {
        value = "${data.oneview_storage_system.ss_int.uri}"
}
// Testing import of existing resource
*/
/*resource "oneview_storage_system" "ss_import"{
}*/

