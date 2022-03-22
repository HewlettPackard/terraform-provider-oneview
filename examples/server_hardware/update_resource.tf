provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

// Updates oneTimeBoot of the server hardware
// it can be Usb, Network, Hdd, Cdrom, Normal
// values are case sensive for Idempotency
resource "oneview_server_hardware" "sh" {
  one_time_boot = "Usb"
}

//Enables or Disables maintenance mode
/*resource "oneview_server_hardware" "sh" {
  maintenance_mode = "true"
}

// Turn On or Off Uid Light
resource "oneview_server_hardware" "sh" {
  uid_state = "On"
}

//Updates Power State of the server hardware
resource "oneview_server_hardware" "sh" {
  server_power_state {
        power_state = "Off"
        power_control = "MomentaryPress"
  }
}*/
