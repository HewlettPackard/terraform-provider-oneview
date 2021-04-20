provider "oneview" {
  ov_username   = "${var.username}"
  ov_password   = "${var.password}"
  ov_endpoint   = "${var.endpoint}"
  ov_sslverify  = "${var.ssl_enabled}"
  ov_apiversion = 2800
  ov_ifmatch    = "*"
}

variable "hostname" {
  type    = "string"
  default = "172.18.11.11" //"<storage_system_ip>"
}
variable "ss_username" {
  type    = "string"
  default = "dcs" //"<storage_system_username>"
}
variable "ss_password" {
  type    = "string"
  default = "dcs" //"<storage_system_password>"
}
variable "ss_family" {
  type    = "string"
  default = "StoreServ" //"<storage_system_family>"
}

# Extracting Storage System
data "oneview_storage_system" "ss_inst" {
   name = "ThreePAR-1"
}

# Updates the resource created from main.tf 
resource "oneview_storage_system" "ss_inst" {
  credentials = [
    {
      username = "${var.ss_username}"
      password = "${var.ss_password}"
    },
  ]

  hostname =  "${var.hostname}"
  name     = "ThreePAR-1"

  storage_system_device_specific_attributes = {
    managed_domain = "TestDomain"
  }

  eTag        = "--"
  description = "TestStorageSystem"
  uri         = "${data.oneview_storage_system.ss_inst.uri}"
}
