provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Extracting required resource 
data "oneview_scope" "scope_obj" {
  name = "<scope>"
}

data "oneview_storage_pool" "st_pool" {
  name = "<storage_pool>"
}

data "oneview_storage_volume_template" "st_vt" {
  name = "<storage_volume_template>"
}

# Updates the volume resource created through main.tf
resource "oneview_volume" "volume" {
  properties {
    name              = "<re_name>"
     /* uncomment the below two line to provide the size and storage pool uri or it will be taken from the volume template*/
    //storage_pool      = data.oneview_storage_pool.st_pool.uri
    //size              = 1221225472/   
    provisioning_type = "Thin"
    is_deduplicated   = false
  }
  template_name        = "<storage_volume_template>"
  template_uri         = data.oneview_storage_volume_template.st_vt.uri
  name                 = "<re_name>"
  description          = "Test Volume Updated"
  initial_scope_uris   = [data.oneview_scope.scope_obj.uri]
  provisioned_capacity = "1221225472"
}
