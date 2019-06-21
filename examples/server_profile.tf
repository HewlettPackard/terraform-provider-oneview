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
  template = "testspt"
}

