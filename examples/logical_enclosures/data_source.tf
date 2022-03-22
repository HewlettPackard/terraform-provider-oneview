provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

variable "LE_name" {
  type        = string
  description = "Logical Enclosure Name"
  default     = "Auto-LE"
}

variable "EG_name" {
  type        = string
  description = "Enclosure Group Name"
  default     = "Auto-EG"
}

# Fetching Enclosure Group
data "oneview_enclosure_group" "enclosure_group" {
  name = var.EG_name
}

# Datasource
data "oneview_logical_enclosure" "logical_enclosure" {
  name                = var.LE_name
  enclosure_group_uri = data.oneview_enclosure_group.enclosure_group.uri
}

output "oneview_logical_enclosure_value" {
  value = data.oneview_logical_enclosure.logical_enclosure.name
}

# Import an Exisiting Logical Enclosure
# terraform import oneview_logical_enclosure.LogicalEnclosure <LE_Name>
#resource "oneview_logical_enclosure" "LogicalEnclosure" {
#}
