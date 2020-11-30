provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2200
  ov_ifmatch    = "*"
}

variable "hm_endpoint" {
  type        = string
  description = "Hypervisor Manager IP"
  default     = "<Hypervisor-Manager-IP>"
}

variable "deployment_plan_password" {
  type        = string
  description = "Deployement Plan Server Password"
  default     = "<password>"
}

data "oneview_hypervisor_manager" "hm" {
  name = var.hm_endpoint
}

data "oneview_server_profile_template" "spt" {
  name = "VolAtSPT"
}

data "oneview_ethernet_network" "eth_nw" {
  name = "ESX vMotion"
}

# Crate Hypervisor Cluster Profile
resource "oneview_hypervisor_cluster_profile" "HypervisorClusterProfile" {
  type                   = "HypervisorClusterProfileV4"
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
    deployment_plan {
      server_password = var.deployment_plan_password
    }
  }
}

# Update Oneview Hypervisor Cluster Profile
resource "oneview_hypervisor_cluster_profile" "HypervisorClusterProfile" {
  type                   = "HypervisorClusterProfileV4"
  name                   = "Cluster-renamed"
  description            = "cluster profile description"
  hypervisor_type        = "Vmware"
  hypervisor_manager_uri = data.oneview_hypervisor_manager.hm.uri
  path                   = "DC2"
  hypervisor_cluster_settings {
    type                = "Vmware"
    drs_enabled         = true
    ha_enabled          = false
    multi_nic_v_motion  = false
    virtual_switch_type = "Standard"
  }
  hypervisor_host_profile_template {
    server_profile_template_uri = data.oneview_server_profile_template.spt.uri
    host_prefix                 = "Cluster-renamed"
    deployment_plan {
      server_password = var.deployment_plan_password
    }
  }
  host_config_policy {
    leave_host_in_maintenance = false
    use_hostname_to_register  = false
  }
  virtual_switch_config_policy {
    manage_virtual_switches = true
    configure_port_groups   = true
  }
  virtual_switches {
    name                = "mgmt"
    virtual_switch_type = "Standard"
    virtual_switch_port_groups {
      name         = "mgmt"
      network_uris = [data.oneview_ethernet_network.eth_nw.uri]
      vlan         = "0"
      virtual_switch_ports {
        virtual_port_purpose = ["Management"]
        dhcp                 = false
        action               = "NONE"
      }
      action = "NONE"
    }
    virtual_switch_uplinks {
      name   = "Mezz 3:1-c"
      active = false
      action = "NONE"
    }
    action       = "NONE"
    network_uris = [data.oneview_ethernet_network.eth_nw.uri]
  }
}

# Testing data source
data "oneview_hypervisor_cluster_profile" "hcp" {
  name = "cluster7"
}

output "oneview_hypervisor_cluster_profile_value" {
  value = data.oneview_hypervisor_cluster_profile.hcp.type
}

