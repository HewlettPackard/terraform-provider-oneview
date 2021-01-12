/*
        Use the the below environment varaiables to set the variable values.
        ------------------------------------------------------------------------------------
        |  No |   Variable Name                             |   value                      |
        ------------------------------------------------------------------------------------
        |  1  | TF_VAR_username                             | <oneview_user_name>          |
        ------------------------------------------------------------------------------------
        |  2  | TF_VAR_password                             | <oneview_password            |
        ------------------------------------------------------------------------------------
        |  3  | TF_VAR_endpoint                             | <oneview_ip>                 |
        ------------------------------------------------------------------------------------
        |  4  | TF_VAR_ssl_enabled                          | <false>                      |
        ------------------------------------------------------------------------------------
*/

variable "username" {
  type        = string
  description = "Oneview Appliance Username"
}

variable "password" {
  type        = string
  description = "Oneview Appliance Password"
}

variable "endpoint" {
  type        = string
  description = "Oneview Appliance Ip Address"
}

variable "ssl_enabled" {
  default     = false
  description = "SSL Enabled"
}

