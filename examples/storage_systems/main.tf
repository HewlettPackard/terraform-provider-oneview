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

# Extracting Server Certificate
data "oneview_server_certificate" "sc" {
  remote_ip = "172.18.11.12" //"${var.hostname}"
}

# Importing Server Certificate for adding storage system
resource "oneview_server_certificate" "ServerCertificate" {
    certificate_details = [{
                        base64_data="${data.oneview_server_certificate.sc.certificate_details.0.base64_data}"
                        type="CertificateDetailV2"
                        alias_name = "StorageSystemCertificate"
                        }]
}

# Adds Storage System to OneView
resource "oneview_storage_system" "ss_inst" {
   hostname = "${var.hostname}"
   username = "${var.ss_username}"
   password = "${var.ss_password}"
   family   = "${var.ss_family}"
   depends_on = ["oneview_server_certificate.ServerCertificate"]
}
