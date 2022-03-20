provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

/*

   Import the existing the resource from appliance into the terraform instance.
   Create a empty resource and execute the following command.

        terraform import oneview_logical_interconnect.<instance> <resource_id>
        instance - instance name declared in the empty resource declared.
        resource_id - id of the logical interconnect as per the appliance.
	terraform import <resource>.<instance_name> <resource_id>

        Example: terraform import oneview_logical_interconnect.li d4468f89-4442-4324-9c01-624c7382db2d
*/

/*
        Use the the particular update string for a particular update functionality.

        ------------------------------------------------------------------------------------
        |  No |   Update Funtion                            |   update_string              |
        ------------------------------------------------------------------------------------
        |  1  | UpdateLogicalInterconnectConsistentById     | updateComplianceById         |
        ------------------------------------------------------------------------------------
*/

resource "oneview_logical_interconnect" "logical_interconnect" {
}

