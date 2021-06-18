provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Collects allocated ips from the Resource
resource "oneview_id_pools_ipv4_subnets" "ipv4_subnets" {
  collector_id_list = ["<ipAddress1>", "<ipAddress2>"]
}

# Retaining dependent resource state for the above update operation
# TF destroys when resource instance not in .tf file

data "oneview_id_pools_ipv4_subnets" "ipv4_subnets_data" {
  network_id = "<network_id>"
}

resource "oneview_ipv4_range" "ipv4range" {
  name = "IpRange"
  subnet_uri = data.oneview_id_pools_ipv4_subnets.ipv4_subnets_data.uri
  start_stop_fragments {
    start_address = "<startAddress>"
    end_address = "<endAddress>"
  }
}

resource "oneview_ethernet_network" "ethernetnetwork" {
  name    = "SubnetEthernet"
  type    = "ethernet-networkV4"
  vlan_id = 157
  subnet_uri = data.oneview_id_pools_ipv4_subnets.ipv4_subnets_data.uri
}
