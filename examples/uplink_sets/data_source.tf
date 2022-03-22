provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Example for data source
data "oneview_uplink_set" "uplink_set" {
  name = "Auto-UplinkSet-Updated"
}

output "oneview_uplink_set_value" {
  value = data.oneview_uplink_set.uplink_set.uri
}

#Importing Existing resource
#resource "oneview_uplink_set" "import_us"{
#}
