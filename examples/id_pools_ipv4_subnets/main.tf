provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Creates ipv4 subnet with the networkId, subnetMask and gateway
resource "oneview_id_pools_ipv4_subnets" "ipv4_subnets" {
  name          = "<subnet_name>"
  network_id    = "<network_id>"
  subnet_mask   = "<subnetmask>"
  gateway       = "<gateway>"
  domain        = "Terraform.com"
}

