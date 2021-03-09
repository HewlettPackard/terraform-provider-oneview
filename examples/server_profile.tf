provider "oneview" {
  ov_username =   "${var.username}"
  ov_password =   "${var.password}"
  ov_endpoint =   "${var.endpoint}"
  ov_sslverify =  "${var.ssl_enabled}"
  ov_apiversion = 2400
  ov_ifmatch = "*"
}

data "oneview_scope" "scope" {
        name = "testing"
}

# Creates a server profile from server profile tempalte

resource "oneview_server_profile" "SP2" {
  name = "TestSP"
  hardware_name = "Synergy-Encl-2, bay 8"
  type = "ServerProfileV12"
  enclosure_group = "SYN03_EC"
  initial_scope_uris = ["${data.oneview_scope.scope.uri}"]
  template = "TestServerProfileTemplate"
}

# Creation of Server Profile without template

resource "oneview_server_profile" "SP2" {
  name = "TestSP2"
  hardware_name = "0000A66101, bay 5"
  type = "ServerProfileV12"
  enclosure_group = "EG"
  boot_order = ["HardDisk"]
  boot_mode = {
	manage_mode = true
	mode = "UEFIOptimized"
	pxe_boot_policy = "Auto"
  }
  connection_settings = {
    connections = [{
	  id = 1
  	  name = "Deployment Network A"
	  function_type = "Ethernet"
	  network_uri = "/rest/ethernet-networks/728af64d-c9aa-4287-a331-ba4c6654dd15"
	  port_id = "Mezz 3:1-a"
	  boot = {
		priority = "Primary"
		ethernet_boot_type = "PXE"
	  }
    }]
  }
}

# Updating Server profile
resource "oneview_server_profile" "SP2" {
  name = "TestSP_Renamed_withCon2"
  hardware_name = "0000A66101, bay 5"
  type = "ServerProfileV12"
  enclosure_group = "EG"
  server_hardware_type = "SY 480 Gen9 1"
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

# To import an existing server profile to terraform, use the below code and run the following command:

# terraform import <resource>.<instance_name> <resource_name>
# Eg: terraform import oneview_server_profile.serverProfile Test

resource "oneview_server_profile" "serverProfile" {
}
