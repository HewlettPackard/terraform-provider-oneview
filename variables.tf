/* 
              Set the the below environment varaiables to set the variable values.
        ------------------------------------------------------------------------------------
        |  No | Variable Name                               | Value                        |
        ------------------------------------------------------------------------------------
        |  1  | TF_VAR_username                             | <oneview_user_name>          |
        ------------------------------------------------------------------------------------
        |  2  | TF_VAR_password                             | <oneview_password            |
        ------------------------------------------------------------------------------------
        |  3  | TF_VAR_endpoint                             | <oneview_ip>                 |
        ------------------------------------------------------------------------------------
        |  4  | TF_VAR_ssl_enabled                          | <false>                      |
        ------------------------------------------------------------------------------------
        |  5  | TF_VAR_api_version                          | <api_version>                |
        ------------------------------------------------------------------------------------
*/

variable "username" {
 type = string
 description = "Oneview Appliance Username"
}

variable "password" {
 type = string
description = "Oneview Appliance Password"
}

variable "endpoint" {
 type = string
 description = "Oneview Appliance IP Address"
}

variable "api_version"{
 type = number
 description = "Oneview Appliance REST API Version"
}

variable "ssl_enabled" {
 default = false
 description = "SSL Enabled"
}

