provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# IMPORTANT: Only one operation i.e., either of Enable Range or Edit Range
# can be performed in a single call of this API.
# Example for Edit Range
resource "oneview_ipv4_range" "ipv4range" {
  name = "IpRange-renamed"
  start_stop_fragments {
    start_address = "<ip_address>"
    end_address = "<ip_address>"
  }
}

# Example for Enable Range
resource "oneview_ipv4_range" "ipv4range" {
  enableb = false    
}