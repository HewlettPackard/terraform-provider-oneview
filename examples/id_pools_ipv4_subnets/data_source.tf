provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2800
  ov_ifmatch    = "*"
}

# Testing data source with subnet_id
data "oneview_id_pools_ipv4_subnets" "ipv4_subnets" {
  subnet_id = "1ca48743-6c42-4a14-b7bf-a32deac20b80"
}
