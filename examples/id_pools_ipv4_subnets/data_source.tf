provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}


data "oneview_id_pools_ipv4_subnets" "ipv4_subnets" {
  network_id = "10.1.0.0"
}

output "oneview_id_pools_ipv4_subnets_value" {
  value = data.oneview_id_pools_ipv4_subnets.ipv4_subnets.uri
}
