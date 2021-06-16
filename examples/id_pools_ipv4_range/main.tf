provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

resource "oneview_ipv4_range" "ipv4range" {
  name = "IpRange"
  subnet_uri = "/rest/id-pools/ipv4/subnets/40f76df9-1e39-4e5a-81fc-14614efea5e8"
  start_stop_fragments {
    start_address = "<ip_address>"
    end_address = "<ip_address>"
  }
  start_stop_fragments {
    start_address = "<ip_address>"
    end_address = "<ip_address>"
  }
}
