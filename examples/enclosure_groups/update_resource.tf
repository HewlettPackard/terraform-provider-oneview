provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
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
  default     = "Auto-EG"
}

variable "scope_name_1" {
  type        = string
  description = "Name of the Scope"
  default     = "Auto-Scope"
}

data "oneview_enclosure_group" "enclosure_group" {
  name = var.enc_grp
}

# Updates the resource created through main.tf
resource "oneview_enclosure_group" "eg_inst" {
  name                     = "${var.enc_grp} Renamed"
  enclosure_count          = 3
  ip_addressing_mode       = "External"
  power_mode               = "RedundantPowerFeed"
  enclosure_type_uri       = data.oneview_enclosure_group.enclosure_group.enclosure_type_uri
  ambient_temperature_mode = data.oneview_enclosure_group.enclosure_group.ambient_temperature_mode
  interconnect_bay_mappings {
    interconnect_bay               = 3
    logical_interconnect_group_name = "${var.LIG_name}"
  }
  interconnect_bay_mappings {
    interconnect_bay               = 6
    logical_interconnect_group_name = "${var.LIG_name}"
  }
  interconnect_bay_mapping_count = 2
  type                           = "EnclosureGroupV8"
  stacking_mode                  = "Enclosure"
}

