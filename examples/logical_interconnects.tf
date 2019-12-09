provider "oneview" {
	ov_username = "<ov_username>"
        ov_password = "<ov_password"
        ov_endpoint = "<ov_endpoint>"
        ov_sslverify = false
        ov_apiversion = <ov_apiversion>
	ov_ifmatch = "*"
}

/*

   Import the existing the resource from appliance into the teriraform instance.
   Create a empty resource and execute the following command.
	
	terraform import oneview_logical_interconnect.<instance> <resource_name>
	instance - instance name declared in the empty resource declared.
	resource_name - name of the logical interconnect as per the appliance.
*/
/*resource "oneview_logical_interconnect" "li" {
}*/

/*
	Use the the particular update string for a particular update functionality.
	
	------------------------------------------------------------------------------------
	|  No |   Update Funtion                            |   update_string              |
	------------------------------------------------------------------------------------
	|  1  | UpdateLogicalInterconnectConsistentById     | updateComplianceById         |
	------------------------------------------------------------------------------------
*/
resource "oneview_logical_interconnect" "li" {
	uri = "/rest/logical-interconnects/d4468f89-4442-4324-9c01-624c7382db2d"
	update_type = "updateComplianceById"
}
