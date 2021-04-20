provider "oneview" {
  ov_username =   "${var.username}"
  ov_password =   "${var.password}"
  ov_endpoint =   "${var.endpoint}"
  ov_sslverify =  "${var.ssl_enabled}"
  ov_apiversion = 2800
  ov_ifmatch = "*"
}

data "oneview_scope" "scope" {
        name = "Auto-Scope"
}

# Updating Server profile
resource "oneview_server_profile" "SP" {
  name = "TestSP_Renamed"
  hardware_name = "0000A66102, bay 5"
  type = "ServerProfileV12"
  enclosure_group = "Auto-EG Renamed"
  server_hardware_type = "SY 480 Gen9 1"
  initial_scope_uris = ["${data.oneview_scope.scope.uri}"]
  update_type = "put"
}

/*
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
*/
