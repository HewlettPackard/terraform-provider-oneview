provider "oneview" {
  ov_username   = "administrator"
  ov_password   = "admin123"
  ov_endpoint   = "https://10.50.9.41"
  ov_sslverify  = false
  ov_apiversion = 1800
  ov_ifmatch    = "*"
}


resource "oneview_storage_system" "ss_inst" {
        hostname = "172.18.11.12"
        username = "dcs"
        password = "dcs"
        family   = "StoreServ"
}


// Uncomment the following resource to update.
/*resource "oneview_storage_system" "ss_inst" {
  credentials = [
    {
      username = "dcs"
      password = "dcs"
    },
  ]

  hostname = "172.18.11.12"
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

