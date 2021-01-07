provider "oneview" {
  ov_username =   var.username
  ov_password =   var.password
  ov_endpoint =   var.endpoint
  ov_sslverify =  var.ssl_enabled
  ov_apiversion = 2400
  ov_ifmatch = "*"
}
/*
data "oneview_scope" "scope" {
  name = "test"
}

data "oneview_ethernet_network" "eth_network" {
  name = "Prod_1103"
}

resource "oneview_scope" "scope_inst" {
  name                = "TestScope"
  description         = "Testing creation of scope"
  type                = "ScopeV3"
  initial_scope_uris  = [data.oneview_scope.scope.uri]
  added_resource_uris = [data.oneview_ethernet_network.eth_network.uri]
}
*/

# Updates the resource created above 
# To update uncomment the below and ad the attributes  to be updated
resource "oneview_scope" "scope_inst" {
  name        = "TestScope_Renamed"
  type        = "ScopeV3"
  description = "Rename the Scope"
}


#Importing Existing resource
#resource "oneview_scope" "import_scope" {
#}
