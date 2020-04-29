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
        name = "172.18.13.11"
        username= "dcs"
        password = "dcs"
        port = 443
	type = "HypervisorManagerV2"
	initial_scope_uris = ["${data.oneview_scope.scope_obj.uri}"]


}/* Uncomment following to run update
resource "oneview_hypervisor_manager" "HypervisorManager" {
	display_name = "TestHypervisorManager_Renamed"
        name = "172.18.13.11"
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
data "oneview_hypervisor_manager" "HypervisorManager" {
        name = "172.18.13.11"
}

output "oneview_hypervisor_manager_value" {
        value = "${data.oneview_hypervisor_manager.HypervisorManager.type}"
}

