provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Test for data source  
data "oneview_logical_interconnect_group" "logical_interconnect_group" {
  name = "TestLIG5"
}

output "lig_value" {
  value = data.oneview_logical_interconnect_group.logical_interconnect_group.redundancy_type
}

# Importing an existing resource from appliance
resource "oneview_logical_interconnect_group" "import_lig" {
}

