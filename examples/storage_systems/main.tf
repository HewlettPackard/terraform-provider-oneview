provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
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
  default = "StoreServ"
}

# Extracting Server Certificate
data "oneview_server_certificate" "sc" {
  remote_ip = var.hostname
}

# Importing Server Certificate for adding storage system
resource "oneview_server_certificate" "ServerCertificate" {
  certificate_details {
    base64_data = element(tolist(data.oneview_server_certificate.sc.certificate_details[*].base64_data), 0)
    type        = "CertificateDetailV2"
    alias_name  = "TestServerCertificate"
  }
}

# Adds Storage System to OneView
resource "oneview_storage_system" "ss_inst" {
  hostname   = var.hostname
  username   = var.ss_username
  password   = var.ss_password
  family     = var.ss_family
  depends_on = [oneview_server_certificate.ServerCertificate]
}

