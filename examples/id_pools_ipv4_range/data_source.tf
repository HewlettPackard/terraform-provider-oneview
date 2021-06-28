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

# Fetches range details using range_id
# When multiple Ranges there, gets last one details

data "oneview_ipv4_range" "ipranges" {
  id = basename(join("",data.oneview_id_pools_ipv4_subnets.ipv4_subnets_data.range_uris))
}

output "ipranges_value" {
  value =  data.oneview_ipv4_range.ipranges
}
