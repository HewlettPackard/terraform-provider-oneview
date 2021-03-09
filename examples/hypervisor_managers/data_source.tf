provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2600
  ov_ifmatch    = "*"
}

variable "hm_endpoint" {
  type        = string
  description = "Hypervisor Manager IP"
  default     = "172.18.13.11" #"<HM-IP>"
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
