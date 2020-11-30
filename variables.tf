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
        |  5  | TF_VAR_i3s_endpoint                         | <image_streamer_ip>          |
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

variable "ssl_enabled" {
 type = bool
 default = false
 description = "SSL Enabled"
}

variable "i3s_endpoint" {
 type = string
 default = false
 description = "Image Streamer IP Address"
}
