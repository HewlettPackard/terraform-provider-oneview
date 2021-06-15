provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2800
  ov_ifmatch    = "*"
}

# Updates the resource created from main.tf
resource "oneview_id_pools_ipv4_subnets" "ipv4_subnets" {
  name="RenamedSubnet"
  network_id="192.169.1.0"
  subnet_mask="255.255.255.0"
  gateway="192.169.1.1"
}
