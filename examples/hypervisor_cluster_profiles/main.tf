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
  default     = "<Hypervisor-Manager-IP>"
}
data "oneview_hypervisor_manager" "hm" {
  name = var.hm_endpoint
}

data "oneview_server_profile_template" "spt" {
  name = "TestServerProfileTemplate"
}

# Crate Hypervisor Cluster Profile
resource "oneview_hypervisor_cluster_profile" "HypervisorClusterProfile" {
  type                   = "HypervisorClusterProfileV6"
  name                   = "Cluster"
  description            = "cluster profile description"
  hypervisor_type        = "Vmware"
  hypervisor_manager_uri = data.oneview_hypervisor_manager.hm.uri
  hypervisor_cluster_settings {
    type                = "Vmware"
    drs_enabled         = true
    ha_enabled          = false
    multi_nic_v_motion  = false
    virtual_switch_type = "Standard"
  }
  hypervisor_host_profile_template {
    server_profile_template_uri = data.oneview_server_profile_template.spt.uri
    host_prefix                 = "Cluster"
    data_store_name_sync 	= true
    deployment_manager_type = "UserManaged"
   
  }
}

