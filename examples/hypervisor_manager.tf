provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2400
  ov_ifmatch    = "*"
}

data "oneview_scope" "scope_obj" {
        name = "test_scope"
}

resource "oneview_hypervisor_manager" "HypervisorManager" {
  display_name       = "TestHypervisorManager"
  name               = "<hypervisor_manager_hostname>"
  username           = "<hypervisor_manager_username>"
  password           = "<hypervisor_manager_password>"
  port               = "<hypervisor_manager_port_num>"
  type               = "HypervisorManagerV2"
  initial_scope_uris = [data.oneview_scope.scope_obj.uri]
}
/*
resource "oneview_hypervisor_manager" "HypervisorManager" {
  display_name = "TestHypervisorManager_Renamed"
  name         = "172.18.13.11"
  query_params = "true"
  type         = "HypervisorManagerV2"
  username     = "dcs1"
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
*/
/* Testing data source*/
/*data "oneview_hypervisor_manager" "HypervisorManager" {
  name = "<hypervisor_manager_hostname>"
}

output "oneview_hypervisor_manager_value" {
  value = data.oneview_hypervisor_manager.HypervisorManager.type
}
*/
//Import existing resource
/*resource "oneview_hypervisor_manager" "HypervisorManager" {
}
*/


