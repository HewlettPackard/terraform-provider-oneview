provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2800
  ov_ifmatch    = "*"
}

# Collects allocated ips from the Resource
resource "oneview_id_pools_ipv4_subnets" "ipv4_subnets" {
  collector_id_list = ["192.169.1.10", "192.169.1.11"]
}

# Retaining dependent resource state for the above update operation
# TF destroys when resource instance not in .tf file

data "oneview_id_pools_ipv4_subnets" "ipv4_subnets_data" {
  subnet_id = "5698662c-235e-473e-828a-73716731fc75"
}

resource "oneview_ipv4_range" "ipv4range" {
  name = "IpRange"
  subnet_uri = data.oneview_id_pools_ipv4_subnets.ipv4_subnets_data.uri
  start_stop_fragments {
    start_address = "192.169.1.10"
    end_address = "192.169.1.20"
  }
}

# Associate Ethernet Resource with subnet
resource "oneview_ethernet_network" "ethernetnetwork" {
  name    = "SubnetEthernet"
  type    = "ethernet-networkV4"
  vlan_id = 157
  subnet_uri = data.oneview_id_pools_ipv4_subnets.ipv4_subnets_data.uri
}
