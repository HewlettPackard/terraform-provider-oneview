provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Data source for server profile

data "oneview_server_profile" "sp" {
  name = "TestSP_Renamed"
}

output "oneview_server_profile_value" {
  value = data.oneview_server_profile.sp.uri
}

/*
# To import an existing server profile to terraform, use the below code and run the following command:

# terraform import <resource>.<instance_name> <resource_name>
# Eg: terraform import oneview_server_profile.serverProfile Test

resource "oneview_server_profile" "serverProfile" {
}
*/
