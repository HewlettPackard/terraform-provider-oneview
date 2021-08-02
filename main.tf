provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 3000
  ov_ifmatch    = "*"
}

# Creating Server Profile Template with Management Settings
resource "oneview_server_profile_template" "SPTwithMgmtSettings" {
  name                 = "Test1234"
  server_hardware_type = "DL560 Gen10 1"
  boot_mode {
    manage_mode     = true
    mode            = "UEFIOptimized"
    pxe_boot_policy = "Auto"
  }
}


