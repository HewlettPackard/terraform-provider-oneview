provider "oneview" {
  ov_username =   "${var.username}"
  ov_password =   "${var.password}"
  ov_endpoint =   "${var.endpoint}"
  ov_sslverify =  "${var.ssl_enabled}"
  ov_apiversion = 2200
  ov_ifmatch = "*"
}

data "oneview_scope" "scope" {
        name = "testing"
}

# Creates a server profile or Updates if already existing

resource "oneview_server_profile" "SP" {
  name = "TestSP"
  hardware_name = "Synergy-Encl-2, bay 8"
  type = "ServerProfileV12"
}

# Creation of Server Profile without template

resource "oneview_server_profile" "SP2" {
  name = "TestSP2"
  hardware_name = "SYN03_Frame3, bay 1"
  type = "ServerProfileV12"
  enclosure_group = "SYN03_EC"
  initial_scope_uris = ["${data.oneview_scope.scope.uri}"]
}



# Updating Server profile
resource "oneview_server_profile" "SP" {
  name = "TestSP_Renamed"
  hardware_name = "Synergy-Encl-2, bay 8"
  type = "ServerProfileV12"
  enclosure_group = "EG-Synergy-Local"
  server_hardware_type = "SY 480 Gen9 1"
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
        type = "ServerProfileV12"
        server_hardware_type = "SY 480 Gen9 1"
        enclosure_group = "EG-Synergy-Local"
        hardware_name = "Synergy-Encl-2, bay 8"
}


#Data source for server profile

data "oneview_server_profile" "sp" {
        name = "TestSP_Renamed"
}

output "oneview_server_profile_value" {
        value = "${data.oneview_server_profile.sp.uri}"
}

/*
# To import an existing server profile to terraform, use the below code and run the following command:

# terraform import <resource>.<instance_name> <resource_name>
# Eg: terraform import oneview_server_profile.serverProfile Test

resource "oneview_server_profile" "serverProfile" {
}*/
