provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2200
  ov_ifmatch    = "*"
}

variable "hostname" {
  type    = string
  default = "<storage_system_ip>"
}

variable "ss_username" {
  type    = string
  default = "<storage_system_username>"
}

variable "ss_password" {
  type    = string
  default = "<storage_system_password>"
}

variable "ss_family" {
  type    = string
  default = "<storage_system_family>"
}

# Extracting Server Certificate
data "oneview_server_certificate" "sc" {
  remote_ip = var.hostname
}

# Importing Server Certificate for adding storage system
resource "oneview_server_certificate" "ServerCertificate" {
  certificate_details {
    base64_data = data.oneview_server_certificate.sc.certificate_details[0].base64_data
    type        = "CertificateDetailV2"
    alias_name  = "TestServerCertificate"
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
data "oneview_storage_system" "ss_inst_data" {
  name = "ThreePAR-2"
}

output "oneview_storage_system_value" {
  value = data.oneview_storage_system.ss_int.uri
}

# Uncomment the following resource to update.
resource "oneview_storage_system" "ss_inst" {
  credentials {
    username = var.ss_username
    password = var.ss_password
  }

  hostname = var.hostname
  name     = "ThreePAR-2"

  storage_system_device_specific_attributes {
    managed_domain = "TestDomain"
  }

  etag        = "--"
  description = "TestStorageSystem"
  uri         = data.oneview_storage_system.ss_inst_data.uri
}

# Testing import of existing resource

resource "oneview_storage_system" "ss_import" {
}

