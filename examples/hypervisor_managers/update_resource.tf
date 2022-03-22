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

variable "hm_username" {
  type        = string
  description = "Hypervisor Manager Username"
  default     = "<HM-Username>"
}
# For API >= 2600 we can set force as a query parameter to true or false, for API <2600 user can send an empty string.
# Update the resource Post applying main.tf  
resource "oneview_hypervisor_manager" "HypervisorManager" {
  force           = "true"
  display_name    = "TestHypervisorManager_Renamed"
  name            = var.hm_endpoint
  type            = "HypervisorManagerV2"
  username        = var.hm_username
  hypervisor_type = "Vmware"
  preferences {
    type                       = "Vmware"
    virtual_switch_type        = "Standard"
    distributed_switch_version = ""
    distributed_switch_usage   = ""
    multi_nic_v_motion         = "false"
    drs_enabled                = "true"
    ha_enabled                 = "false"
  }
}

