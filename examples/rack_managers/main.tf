provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

data "oneview_scope" "scope_obj" {
  name = "testing"
}
# Add the Rack Manager 
resource "oneview_rack_manager" "RM" {
  hostname               = "<ipaddress>"
  username               = "<user>"
  password               = "<password>"
  initial_scope_uris = [data.oneview_scope.scope_obj.uri]
  
}

// Uncomment below  to import an existing rack manager
// Run 'terraform improt oneview_rack_manager.import_rm <rack-manager-name>'

# resource "oneview_rack_manager" "import_rm"{
# }

