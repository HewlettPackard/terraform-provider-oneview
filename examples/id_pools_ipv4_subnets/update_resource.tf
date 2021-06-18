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
  name = "RenamedSF"
  network_id="<network_id>"
  subnet_mask="<subnet_mask>"
  gateway="<gateway>"
}

# Below resources are prerequisite for update_resource_allocator.tf
# which allocates IPs from subnet to the resource

# Gets subnet details using network id
data "oneview_id_pools_ipv4_subnets" "ipv4_subnets_data" {
  network_id = "<network_id>"
}

# Creates Range of Ip Addresses for the subnet 
# To allocate ips from subnet to associated resource
resource "oneview_ipv4_range" "ipv4range" {
  name = "IpRange"
  subnet_uri = data.oneview_id_pools_ipv4_subnets.ipv4_subnets_data.uri
  start_stop_fragments {
    start_address = "<startAddress>"
    end_address = "<endAddress>"
  }
}

# Associate Ethernet Resource with subnet
resource "oneview_ethernet_network" "ethernetnetwork" {
  name    = "SubnetEthernet"
  type    = "ethernet-networkV4"
  vlan_id = 157
  subnet_uri = data.oneview_id_pools_ipv4_subnets.ipv4_subnets_data.uri
}
