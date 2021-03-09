provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2600
  ov_ifmatch    = "*"
}

variable "LIG_name" {
  type        = string
  description = "Logical Interconnect Name"
  default     = "Auto-LIG"
}

variable "enc_grp" {
  type        = string
  description = "Name of the Enclosure Group"
  default     = "EG"
}

variable "scope_name_1" {
  type        = string
  description = "Name of the Scope"
  default     = "Auto-Scope"
}

# Fetching Logical Interconnect Group
#data "oneview_logical_interconnect_group" "logical_interconnect_group" {
#        name = "${var.LIG_name}"
#}
data "oneview_logical_interconnect_group" "logical_interconnect_group" {
  name = "Auto-LIG"
}

# Creating Enclosure Group
resource "oneview_enclosure_group" "eg_inst" {
  name               = var.enc_grp
  description        = "Testing creation of Enclosure Group"
  ip_addressing_mode = "External"
  enclosure_count    = 3
  initial_scope_uris = [var.scope_name_1]
  interconnect_bay_mappings {
    interconnect_bay               = 3
    logical_interconnect_group_uri = "/rest/logical-interconnect-groups/08cc9fdd-c831-4437-9414-ca3cfa12d292" #"${data.oneview_logical_interconnect_group.logical_interconnect_group.uri}"
  }
  interconnect_bay_mappings {
    interconnect_bay               = 6
    logical_interconnect_group_uri = "/rest/logical-interconnect-groups/08cc9fdd-c831-4437-9414-ca3cfa12d292" #"${data.oneview_logical_interconnect_group.logical_interconnect_group.uri}"
  }
}

