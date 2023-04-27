provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Updates Subnet Name
resource "oneview_id_pools_ipv4_subnets" "ipv4_subnets" {
  name          = "RenamedSF"
  network_id    = "10.1.0.0"
  subnet_mask   = "255.255.192.0"
  gateway       = "10.1.0.1"
}

# Below resources are prerequisite for update_resource_allocator.tf
# which allocates IPs from subnet to the resource

# Gets subnet details using network id
data "oneview_id_pools_ipv4_subnets" "ipv4_subnets_data" {
  network_id = "10.1.0.0"
}

# Creates Range of Ip Addresses for the subnet 
# To allocate ips from subnet to associated resource
resource "oneview_ipv4_range" "ipv4range" {
  name = "IpRange"
  subnet_uri = data.oneview_id_pools_ipv4_subnets.ipv4_subnets_data.uri
  start_stop_fragments {
    start_address = "10.1.19.56"
    end_address = "10.1.19.61"
  }
}

# Associate Ethernet Resource with subnet
resource "oneview_ethernet_network" "ethernetnetwork" {
  name                  = "SubnetEthernet"
  type                  = "ethernet-networkV4"
  ethernet_network_type = "Tagged"
  vlan_id               = 215
  subnet_uri = data.oneview_id_pools_ipv4_subnets.ipv4_subnets_data.uri
}

