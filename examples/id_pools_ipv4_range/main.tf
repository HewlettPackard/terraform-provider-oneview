provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

data "oneview_id_pools_ipv4_subnets" "ipv4_subnets_data" {
  network_id = "<network_id>"
}

# Creates a subnet range
resource "oneview_ipv4_range" "ipv4range" {
  name = "IpRange"
  subnet_uri = data.oneview_id_pools_ipv4_subnets.ipv4_subnets_data.uri
  start_stop_fragments {
    start_address = "<start_address>"
    end_address = "<end_address>"
  }
}
