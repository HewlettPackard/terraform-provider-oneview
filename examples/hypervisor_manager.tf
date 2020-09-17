provider "oneview" {
	ov_username = "<username>"
	ov_password = "<password>"
	ov_endpoint = "<endpoint>"
	ov_sslverify = false
	ov_apiversion = <ov-apiversion>
	ov_ifmatch = "*"
}

data "oneview_scope" "scope_obj" {
        name = "test_scope"
}

resource "oneview_hypervisor_manager" "HypervisorManager" {
	display_name = "TestHypervisorManager"
        name = "<hypervisor_manager_hostname>"
        username= "<hypervisor_manager_username>"
        password = "<hypervisor_manager_password>"
        port = <hypervisor_manager_port_num>
	type = "HypervisorManagerV2"
	initial_scope_uris = ["${data.oneview_scope.scope_obj.uri}"]


}/* Uncomment following to run update
resource "oneview_hypervisor_manager" "HypervisorManager" {
	display_name = "TestHypervisorManager_Renamed"
        name = "<hypervisor_manager_hostname>"
	type = "HypervisorManagerV2"
        username= "dcs1"
        preferences= {   type= "Vmware"
			 virtual_switch_type= "Standard"
			 distributed_switch_version=""
			 distributed_switch_usage=""
			 multi_nic_v_motion="false"
			 drs_enabled="true"
			 ha_enabled="false"
	} 
}
*/
/* Testing data source*/
/*data "oneview_hypervisor_manager" "HypervisorManager" {
        name = "<hypervisor_manager_hostname>"
}

output "oneview_hypervisor_manager_value" {
        value = "${data.oneview_hypervisor_manager.HypervisorManager.type}"
}
*/
//Import existing resource
/*resource "oneview_hypervisor_manager" "HypervisorManager" {
}
*/


