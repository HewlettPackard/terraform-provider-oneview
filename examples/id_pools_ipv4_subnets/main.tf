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
  name          = "Test IPv4 Subnet"
  network_id    = "10.1.0.0"
  subnet_mask   = "255.255.192.0"
  gateway       = "10.1.0.1"
  domain        = "Terraform.com"
}


