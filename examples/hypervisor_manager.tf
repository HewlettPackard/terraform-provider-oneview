provider "oneview" {
  ov_username =   "${var.username}"
  ov_password =   "${var.password}"
  ov_endpoint =   "${var.endpoint}"
  ov_sslverify =  "${var.ssl_enabled}"
  ov_apiversion = 2400
  ov_ifmatch = "*"
}

variable "hm_endpoint" {
 type = "string"
 description = "Hypervisor Manager IP"
 default = "<HM-IP>"
}

variable "hm_username" {
 type = "string"
 description = "Hypervisor Manager Username"
 default = "<HM-Username>"
}

variable "hm_password" {
 type = "string"
 description = "Hypervisor Manager Password"
 default = "<HM-Password>"
}

data "oneview_scope" "scope_obj" {
        name = "testing"
}

/*
# Fetching Server Certificate
data "oneview_server_certificate" "sc" {
         remote_ip = "${var.hm_endpoint}"
}
#<hypervisor_manager_port_num>
# Importing Server Certificate from Hypervisor Manager
resource "oneview_server_certificate" "ServerCertificate" {
    certificate_details = [{
                        base64_data="${data.oneview_server_certificate.sc.certificate_details.0.base64_data}"
                        type="CertificateDetailV2"
                        alias_name = "HM-ServerCertificate"
                        }]

}

# Adding Hypervisor Manager
resource "oneview_hypervisor_manager" "HypervisorManager" {
	display_name = "TestHypervisorManager"
        name = "${var.hm_endpoint}"
        username= "${var.hm_username}"
        password = "${var.hm_password}"
        port = 443
	type = "HypervisorManagerV2"
	initial_scope_uris = ["${data.oneview_scope.scope_obj.uri}"]
}
*/

/*
# Uncomment following to run update
resource "oneview_hypervisor_manager" "HypervisorManager" {
	display_name = "TestHypervisorManager_Renamed"
        name = "${var.hm_endpoint}"
	type = "HypervisorManagerV2"
        username= "${var.hm_username}"
        hypervisor_type=  "Vmware"
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
        name = "${var.hm_endpoint}"
}

#output "oneview_hypervisor_manager_value" {
#        value = "${data.oneview_hypervisor_manager.HypervisorManager.type}"
#}


//Import existing resource
/*resource "oneview_hypervisor_manager" "HypervisorManager" {
}
*/
