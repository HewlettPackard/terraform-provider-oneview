provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

variable "hm_endpoint" {
  type        = string
  description = "Hypervisor Manager IP"
  default     = "<HM-IP>" #

variable "hm_username" {
  type        = string
  description = "Hypervisor Manager Username"
  default     = "<HM-Username>"

variable "hm_password" {
  type        = string
  description = "Hypervisor Manager Password"
  default     = "<HM-Password>"
}

data "oneview_scope" "scope_obj" {
  name = "Auto-Scope"
}

# Fetching Server Certificate
data "oneview_server_certificate" "sc" {
  remote_ip = var.hm_endpoint
}

# Importing Server Certificate from Hypervisor Manager
resource "oneview_server_certificate" "ServerCertificate" {
  certificate_details {
    //base64_data = data.oneview_server_certificate.sc.certificate_details[0].base64_data
    base64_data = element(tolist(data.oneview_server_certificate.sc.certificate_details[*].base64_data), 0)

    type        = "CertificateDetailV2"
    alias_name  = "HM-ServerCertificate"
  }
}

# Adding Hypervisor Manager
resource "oneview_hypervisor_manager" "HypervisorManager" {
  display_name       = "TestHypervisorManager"
  name               = var.hm_endpoint
  query_params       = "true"
  username           = var.hm_username
  password           = var.hm_password
  port               = 443
  type               = "HypervisorManagerV2"
  initial_scope_uris = [data.oneview_scope.scope_obj.uri]
}

