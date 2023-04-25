provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}
# Create SNMPv3 User
resource "oneview_appliance_snmpv3_user" "snmpvuser" {
  user_name                 = "<user_name>"
  security_level            = "<security_level>"
  authentication_protocol   = "<authentication_protocol>"
  authentication_passphrase = "<authentication_passphrase>"
  privacy_protocol          = "AES-128"
  privacy_passphrase        = "1234567812345678"
}
