provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Testing data source
data "oneview_ethernet_network" "ethernetnetworks1" {
  name = "TestNetwork_1"
}

data "oneview_ethernet_network" "ethernetnetworks2" {
  name = "iscsi_nw"
}

data "oneview_scope" "scope" {
  name = "Auto-Scope"
}

#Create a Server Profile using a Server profile template and get task uri

data "oneview_server_hardware" "sh" {
  name = "172.18.31.4"
}

data "oneview_server_profile_template" "spt" {
  name = "Test_spt"
}


resource "oneview_server_profile_async_task_uri" "sp_from_spt" {
  name                 = "sp_from_spt"
  hardware_name        = "0000A66101, bay 4"
  server_hardware_type = "SY 660 Gen9 2"
  template             = data.oneview_server_profile_template.spt.name
  scopes_uri           = data.oneview_scope.scope.uri
  force_flags          = ["ignoreServerHealth"]
}
