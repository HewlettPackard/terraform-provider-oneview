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

data  "oneview_enclosure_group" "eg" {
  name = "EG_Range"
}

# Allocates 2 Ip addresses from range to associated EG
resource "oneview_ipv4_range" "ipv4range" {
  allocator_count = 2
  name = "IpRange-renamed"
  subnet_uri = data.oneview_id_pools_ipv4_subnets.ipv4_subnets_data.uri
  start_stop_fragments {
    start_address = "<start_address>"
    end_address = "<end_address>"
  }
  associated_resources {
    resource_category = data.oneview_enclosure_group.eg.category
    resource_name     = data.oneview_enclosure_group.eg.name
    resource_uri      = data.oneview_enclosure_group.eg.uri
  }

}


# Retains existing EG for allocation
resource "oneview_enclosure_group" "eg_inst" {
  name               = "EG_Range"
  enclosure_count    = 3
  ip_addressing_mode = data.oneview_enclosure_group.eg.ip_addressing_mode
  ip_range_uris = data.oneview_enclosure_group.eg.ip_range_uris
  enclosure_type_uri       = data.oneview_enclosure_group.eg.enclosure_type_uri
  ambient_temperature_mode = data.oneview_enclosure_group.eg.ambient_temperature_mode
  type           = "EnclosureGroupV8"
  stacking_mode  = "Enclosure"
  power_mode  = "RedundantPowerFeed"
  status = "OK"

}

