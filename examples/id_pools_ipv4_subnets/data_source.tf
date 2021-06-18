provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Testing data source with subnet_id or network_id
# subnet_id = "<subnet_id>"
# network_id = "<network_id>"

data "oneview_id_pools_ipv4_subnets" "ipv4_subnets" {
  network_id = "<network_id>"
}

output "oneview_id_pools_ipv4_subnets_value" {
  value = data.oneview_id_pools_ipv4_subnets.ipv4_subnets.uri
}
