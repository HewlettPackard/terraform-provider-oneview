provider "oneview" {
  ov_username =   var.username
  ov_password =   var.password
  ov_endpoint =   var.endpoint
  ov_sslverify =  var.ssl_enabled
  ov_apiversion = 2200
  ov_ifmatch = "*"
}
/*
data "oneview_scope" "scope" {
  name = "Auto-Scope"
}

data "oneview_fc_network" "fc_network" {
  name = "UCFCNET"
}

resource "oneview_scope" "scope_inst" {
  name                = "TestScope"
  description         = "Testing creation of scope"
  type                = "ScopeV3"
  initial_scope_uris  = [data.oneview_scope.scope.uri]
  added_resource_uris = [data.oneview_fc_network.fc_network.uri]
}

# Updates the resource created above 
# To update uncomment the below and ad the attributes  to be updated
resource "oneview_scope" "scope_inst" {
  name        = "TestScope_Renamed"
  type        = "ScopeV3"
  description = "Rename the Scope"
}*/
#Importing Existing resource
#resource "oneview_scope" "import_scope" {
#}
