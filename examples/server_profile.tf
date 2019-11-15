provider "oneview" {
        ov_username = "<ov_username>"
        ov_password = "<ov_password"
        ov_endpoint = "<ov_endpoint>"
        ov_sslverify = false
        ov_apiversion = <ov_apiversion>
        ov_ifmatch = "*"
}

# Creates a server profile or Updates if already existing

resource "oneview_server_profile" "SP" {
  name = "TestSP"
  hardware_name = "0000A66102, bay 3"
  type = "ServerProfileV9"
}

# Creation of Server Profile without template
resource "oneview_server_profile" "SP" {
  name = "TestSP"
  hardware_name = "SYN03_Frame3, bay 1"
  type = "ServerProfileV10"
  enclosure_group = "SYN03_EC"
  initial_scope_uris = ["${data.oneview_scope.scope.uri}"]
}

# Updating Server profile
resource "oneview_server_profile" "SP" {
  name = "TestSP_Renamed"
  hardware_name = "SYN03_Frame3, bay 1"
  type = "ServerProfileV10"
  enclosure_group = "SYN03_EC"
  server_hardware_type = "SY 480 Gen9 3"
  initial_scope_uris = ["${data.oneview_scope.scope.uri}"]
  update_type = "put"
}

# Patch request - Server profile Refresh
resource "oneview_server_profile" "SP" {
        update_type = "patch"
        options = [
        {
          op = "replace"
          path = "/refreshState"
          value = "RefreshPending"
        }
        ]
        name = "TestSP_Renamed"
        type = "ServerProfileV10"
        server_hardware_type = "SY 480 Gen9 3"
        enclosure_group = "SYN03_EC"
        hardware_name = "SYN03_Frame3, bay 1"
}

#Data source for server profile

data "oneview_server_profile" "sp" {
        name = "TestAll"
}

output "oneview_server_profile_value" {
        value = "${data.oneview_server_profile.sp.uri}"
}

# To import an existing server profile to terraform, use the below code and run the following command:

# terraform import <resource>.<instance_name> <resource_name>
# Eg: terraform import oneview_server_profile.serverProfile Test

resource "oneview_server_profile" "serverProfile" {
}
