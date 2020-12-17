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
  name = "49c7fb2c-8173-4dbd-870a-617bfb0a6552"
}

output "oneview_logical_interconnect_value" {
  value = data.oneview_logical_interconnect.logical_interconnect.uri
}

# To import an existing logical interconnect to terraform for update, use the below code and run the following command:
# terraform import <resource>.<instance_name> <resource_id>
# Eg: terraform import oneview_logical_interconnect.logical_interconnect dc411e73-c106-49dd-a523-112a0e169f12

resource "oneview_logical_interconnect" "logical_interconnect" {
}

# Returns logical interconnects to a consistent state. The current logical interconnect state is compared to the associated logical interconnect group.
resource "oneview_logical_interconnect" "logical_interconnect" {
  update_type = "updateComplianceById"
}

