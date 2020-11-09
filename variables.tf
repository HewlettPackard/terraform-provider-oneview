variable "username" {
 type = "string"
 description = "Oneview Appliance Username"
}

variable "password" {
 type = "string"
description = "Oneview Appliance Password"
}

variable "endpoint" {
 type = "string"
 description = "Oneview Appliance Ip Address"
}

variable "ssl_enabled" {
 default = false
 description = "SSL Enabled"
}
