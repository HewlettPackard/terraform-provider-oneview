provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

variable "enc_grp" {
  type        = string
  description = "Name of the Enclosure Group"
  default     = "Auto-EG"
}

# Fetching Enclosure Group
data "oneview_enclosure_group" "enclosure_group" {
  name = var.enc_grp
}

output "oneview_enclosure_group_value" {
  value = data.oneview_enclosure_group.enclosure_group.uri
}

