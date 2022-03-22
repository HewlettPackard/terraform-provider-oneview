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
  default     = "<HM-IP>"
}

# Testing data source
data "oneview_hypervisor_manager" "HypervisorManager" {
  name = var.hm_endpoint
}

output "oneview_hypervisor_manager_value" {
  value = data.oneview_hypervisor_manager.HypervisorManager.type
}

#Import existing resource
#resource "oneview_hypervisor_manager" "HypervisorManager" {
# }
