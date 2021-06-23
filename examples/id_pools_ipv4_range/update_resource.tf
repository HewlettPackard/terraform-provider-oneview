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

# IMPORTANT: Only one operation i.e., either of Enable Range or Edit Range
# can be performed in a single call of this API.
# Example for Edit Range
resource "oneview_ipv4_range" "ipv4range" {
  name = "IpRange-renamed"
  subnet_uri = data.oneview_id_pools_ipv4_subnets.ipv4_subnets_data.uri
  start_stop_fragments {
    start_address = "<start_address>"
    end_address = "<end_address>"
  }
}


# Creating Enclosure Group and associate the subnet range with it
resource "oneview_enclosure_group" "eg_inst" {
  name               = "EG_Range"
  ip_addressing_mode = "ipPool"
  ip_range_uris = data.oneview_id_pools_ipv4_subnets.ipv4_subnets_data.range_uris
  enclosure_count    = 3
}
