provider "oneview" {
	ov_username = "Administrator"
	ov_password = "admin123"
	ov_endpoint = "https://10.50.9.41"
	ov_sslverify = false
	ov_apiversion = 1600
	ov_ifmatch = "*"
}

data "oneview_scope" "scope_obj" {
        name = "test_scope"
}

resource "oneview_hypervisor_cluster_profile" "HypervisorClusterProfile"{  
    "type"="HypervisorClusterProfileV3",
    "name"="Cluster",
    "description"="asa",
    "hypervisor_type"="Vmware",
    "hypervisor_manager_uri"="/rest/hypervisor-managers/4d47b4ca-1c40-47ad-b170-099b04e3691f",
    "hypervisor_cluster_settings"={  
                                  "type"="Vmware",
                                  "drs_enabled"=true,
                                  "ha_enabled"=false,
                                  "multi_nic_v_motion"=false,
                                  "virtual_switch_type"="Standard"
                               },
    "hypervisor_host_profile_template"={  
                                  "server_profile_template_uri"="/rest/server-profile-templates/87afd6f4-8b24-4182-8f06-8b5cc5b40786"
                                  "host_prefix"="Cluster",

     }
}
/*

resource "oneview_hypervisor_cluster_profile" "HypervisorClusterProfile"{  
    "type"="HypervisorClusterProfileV3",
    "name"="Cluster-renamed",
    "description"="asa",
    "hypervisor_type"="Vmware",
    "hypervisor_manager_uri"="/rest/hypervisor-managers/4d47b4ca-1c40-47ad-b170-099b04e3691f",
    "path"="DC2",
    "hypervisor_cluster_settings"={  
                                  "type"="Vmware",
                                  "drs_enabled"=true,
                                  "ha_enabled"=false,
                                  "multi_nic_v_motion"=false,
                                  "virtual_switch_type"="Standard"
                               },
    "hypervisor_host_profile_template"={  
    "server_profile_template_uri"="/rest/server-profile-templates/87afd6f4-8b24-4182-8f06-8b5cc5b40786",
    "host_prefix"="Cluster-renamed",

     },
     "host_config_policy"= {
            "leave_host_in_maintenance"= false,
            "use_hostname_to_register"= false
        },
        "virtual_switch_config_policy"= {
            "manage_virtual_switches"= true,
            "configure_port_groups"= true,
        },
"virtual_switches"=[{
        "name"="mgmt",
        "virtual_switch_type"="Standard",

        "virtual_switch_port_groups"=[{
            "name"="mgmt",
            "network_uris"=["/rest/ethernet-networks/8fa842d8-558e-4958-80cf-116bad88e68b"],
            "vlan"="0",
            "virtual_switch_ports"=[{
                "virtual_port_purpose"=["Management"],
                "dhcp"=false,
                "action"="NONE"
              }],
            "action"="NONE"
          }],
        "virtual_switch_uplinks"=[{
            "name"="Mezz 3:1-c",
            "active"=false,
            "action"="NONE"
          }],
        "action"="NONE",
        "network_uris"=["/rest/ethernet-networks/8fa842d8-558e-4958-80cf-116bad88e68b"]
      }],
}/*

/* Testing data source

data "oneview_hypervisor_cluster_profile" "hcp" {
         name = "cluster7"
}
output "oneview_hypervisor_cluster_profile_value" {
        value = "${data.oneview_hypervisor_cluster_profile.hcp.type}"
}
*/
`
