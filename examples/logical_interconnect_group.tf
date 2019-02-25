provider "oneview" {
  ov_username   = "administrator"
  ov_password   = "madhav123"
  ov_endpoint   = "https://10.170.16.44"
  ov_sslverify  = false
  ov_apiversion = 800
  ov_ifmatch    = "*"
}

resource "oneview_logical_interconnect_group" "LIG" {
  name = "TestLIG"
  type = "logical-interconnect-groupV5"

  interconnect_map_entry_template = [
    {
      bay_number             = 1
      interconnect_type_name = "HP VC Flex-10 Enet Module"
    },
  ]
}

# Updates the resource created above by changing the bay_number
# To update, uncomment the below and add the attributes to be updated.
/*
resource "oneview_logical_interconnect_group" "LIG" {
        name = "TestLIG"
        type = "logical-interconnect-groupV5"
        interconnect_map_entry_template = [
                {
                        bay_number = 2
                        interconnect_type_name = "HP VC Flex-10 Enet Module"
                }
        ]
}
*/

