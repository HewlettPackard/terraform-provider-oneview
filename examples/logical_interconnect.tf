provider "oneview" {
  ov_username =   var.username
  ov_password =   var.password
  ov_endpoint =   var.endpoint
  ov_sslverify =  var.ssl_enabled
  ov_apiversion = 2200
  ov_ifmatch = "*"
}

# Get Logical Interconnect to terraform
data "oneview_logical_interconnect" "logical_interconnect" {
  name = "05c45bdf-a2eb-4461-95b9-3e971b9e105e"
}

output "oneview_logical_interconnect_value" {
  value = data.oneview_logical_interconnect.logical_interconnect.uri
  depends_on = [ oneview_logical_interconnect.logical_interconnect]
}

# To import an existing logical interconnect to terraform for update, use the below code and run the following command:
# terraform import <resource>.<instance_name> <resource_id>
# Eg: terraform import oneview_logical_interconnect.logical_interconnect dc411e73-c106-49dd-a523-112a0e169f12
/*
resource "oneview_logical_interconnect" "logical_interconnect" {
}
*/
# Returns logical interconnects to a consistent state. The current logical interconnect state is compared to the associated logical interconnect group.
resource "oneview_logical_interconnect" "logical_interconnect" {
  update_type = "updateComplianceById"
}

