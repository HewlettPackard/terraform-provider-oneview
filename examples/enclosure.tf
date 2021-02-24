provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2400
  ov_ifmatch    = "*"
}

/*
# Add Enclosure to Oneview Appliance

variable "enc_hostname" {
  type        = string
  description = "Enclosure IP Address"
  default     = "https://x.x.x.x"
}

variable "enc_username" {
  type        = string
  description = "Enclosure Username"
  default     = "username"
}

variable "enc_password" {
  type        = string
  description = "Enclosure Password"
  default     = "password"
}

variable "enc_name" {
  type        = string
  description = "Enclosure name"
  default     = "Enc-1"
}

data "oneview_enclosure_group" "enclosure_group" {
  name = "TestEnclosureGroup_Renamed"
}

data "oneview_scope" "scope" {
  name = "test"
}

resource "oneview_enclosure" "enclosure_inst" {
  enclosure_group_uri = data.oneview_enclosure_group.enclosure_group.uri
  host_name           = var.en_hostname
  user_name           = var.en_username
  password            = var.enc_password
  licensing_intent    = "Oneview"
  initial_scope_uris  = ["scope_name1", "scope_name2"]
  name                = var.enc_name
}
*/

#Importing Existing resource to update
#run the below command to import the resource:
#terraform import oneview_enclosure.enclosure_inst <your_enclosure_name>

/*
resource "oneview_enclosure" "import_enc" {
}
*/

/*
# Testing data source
data "oneview_enclosure" "enclosure" {
  name = "Synergy-Encl-1"
}

output "oneview_enclosure_value" {
  value = data.oneview_enclosure.enclosure.uuid
}

# Update Enclosure
resource "oneview_enclosure" "import_enc" {
  op                  = "replace"
  path                = "/name"
  value               = "Synergy-Encl 1"
  name                = "Synergy-Encl 1"
  enclosure_group_uri = data.oneview_enclosure.enclosure.enclosure_group_uri
}
*/
