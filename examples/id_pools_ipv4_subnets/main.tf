provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2800
  ov_ifmatch    = "*"
}

# Creates ipv4 subnet with the networkId, subnetMask and gateway
resource "oneview_id_pools_ipv4_subnets" "ipv4_subnets" {
  name="SubnetTF"
  network_id="192.169.1.0"
  subnet_mask="255.255.255.0"
  gateway="192.169.1.1"
  domain= "Terraform.com"
}

