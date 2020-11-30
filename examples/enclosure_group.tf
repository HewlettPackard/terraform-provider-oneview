provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2200
  ov_ifmatch    = "*"
}

variable "LIG_name" {
  type        = string
  description = "Logical Interconnect Name"
  default     = "LIG-Name"
}

variable "enc_grp" {
  type        = string
  description = "Name of the Enclosure Group"
  default     = "EG-Name"
}

variable "scope_name_1" {
  type        = string
  description = "Name of the Scope"
  default     = "Scope-Name"
}

variable "scope_name_2" {
  type        = string
  description = "Name of the Scope"
  default     = "Scope-Name"
}

# Fetching Logical Interconnect Group
data "oneview_logical_interconnect_group" "logical_interconnect_group" {
  name = var.LIG_name
}

# Creating Enclosure Group
resource "oneview_enclosure_group" "eg_inst" {
  name               = var.enc_grp
  description        = "Testing creation of Enclosure Group"
  ip_addressing_mode = "External"
  enclosure_count    = 3
  initial_scope_uris = [var.scope_name_1, var.scope_name_2]
  interconnect_bay_mappings {
    interconnect_bay               = 3
    logical_interconnect_group_uri = data.oneview_logical_interconnect_group.logical_interconnect_group.uri
  }
  interconnect_bay_mappings {
    interconnect_bay               = 6
    logical_interconnect_group_uri = data.oneview_logical_interconnect_group.logical_interconnect_group.uri
  }
}

# Fetching Enclosure Group
data "oneview_enclosure_group" "enclosure_group" {
  name = var.enc_grp
}

output "oneview_enclosure_group_value" {
  value = data.oneview_enclosure_group.enclosure_group.uri
}

# Updates the resource created above
# To update uncomment the below lines and add the values to the attributes mentioned
resource "oneview_enclosure_group" "eg_inst" {
  name                     = "${var.enc_grp} Renamed"
  enclosure_count          = 3
  ip_addressing_mode       = "External"
  power_mode               = "RedundantPowerFeed"
  enclosure_type_uri       = data.oneview_enclosure_group.enclosure_group.enclosure_type_uri
  ambient_temperature_mode = data.oneview_enclosure_group.enclosure_group.ambient_temperature_mode
  interconnect_bay_mappings {
    interconnect_bay               = 3
    logical_interconnect_group_uri = data.oneview_logical_interconnect_group.logical_interconnect_group.uri
  }
  interconnect_bay_mappings {
    interconnect_bay               = 6
    logical_interconnect_group_uri = data.oneview_logical_interconnect_group.logical_interconnect_group.uri
  }
  interconnect_bay_mapping_count = 2
  type                           = "EnclosureGroupV8"
  stacking_mode                  = "Enclosure"
}

