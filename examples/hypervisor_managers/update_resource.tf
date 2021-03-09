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

variable "hm_username" {
  type        = string
  description = "Hypervisor Manager Username"
  default     = "dcs" #"<HM-Username>"
}

# Update the resource Post applying main.tf  
resource "oneview_hypervisor_manager" "HypervisorManager" {
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

