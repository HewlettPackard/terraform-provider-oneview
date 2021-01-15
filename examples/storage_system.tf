provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password 
  ov_endpoint   = var.endpoint 
  ov_sslverify  = var.ssl_enabled 
  ov_apiversion = 2400
  ov_ifmatch    = "*"
}

variable "hostname" {
  type    = string
  default = "172.18.11.11"
}

variable "ss_username" {
  type    = string
  default = "dcs"
}

variable "ss_password" {
  type    = string
  default = "dcs"
}

variable "ss_family" {
  type    = string
  default = "StoreServ"
}

# Extracting Server Certificate
data "oneview_server_certificate" "sc" {
  remote_ip = "172.18.11.11"
}
output "oneview_server_certificate_value" {
  #value = data.oneview_server_certificate.sc.certificate_details{0}.base64_data
  value = element(tolist(data.oneview_server_certificate.sc.certificate_details[*].base64_data), 0)
}
# Importing Server Certificate for adding storage system
resource "oneview_server_certificate" "ServerCertificate" {
  certificate_details {
    base64_data = element(tolist(data.oneview_server_certificate.sc.certificate_details[*].base64_data), 0)
    type        = "CertificateDetailV2"
    alias_name  = "TestServerCertificate1"
  }
 
}

resource "oneview_storage_system" "ss_inst" {
  hostname   = var.hostname
  username   = var.ss_username
  password   = var.ss_password
  family     = var.ss_family
  depends_on = [oneview_server_certificate.ServerCertificate]
}


# Extracting Storage System
data "oneview_storage_system" "ss_inst" {
  name = "ThreePAR-1"
}

output "oneview_storage_system_value" {
  value = data.oneview_storage_system.ss_inst.uri
}
# Uncomment the following resource to update.
/*
resource "oneview_storage_system" "ss_inst" {
  credentials {
    username = var.ss_username
    password = var.ss_password
  }

  hostname = var.hostname
  name     = "ThreePAR-1"

  storage_system_device_specific_attributes {
    managed_domain = "TestDomain"
  }

  etag        = "--"
  description = "TestStorageSystem"
  uri         = data.oneview_storage_system.ss_inst.uri
}

# Testing import of existing resource

resource "oneview_storage_system" "ss_import"{
}
*/
