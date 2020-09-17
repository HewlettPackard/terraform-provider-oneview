provider "oneview" {
        ov_username = "<ov_username>"
        ov_password = "<ov_password"
        ov_endpoint = "<ov_endpoint>"
        ov_sslverify = false
        ov_apiversion = <ov_apiversion>
        ov_ifmatch = "*"
}

# Get Logical Interconnect to terraform
data "oneview_logical_interconnect" "logical_interconnect" {
	name = "5fe6ef52-42d6-4895-8500-dfb633478225"
}

output "oneview_logical_interconnect_value" {
	value = "${data.oneview_logical_interconnect.logical_interconnect.uri}"
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
