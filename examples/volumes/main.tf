provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 3600
  ov_ifmatch    = "*"
}

# Extracting resources for volume creation
data "oneview_scope" "scope_obj" {
  name = "Auto-Scope"
}

data "oneview_storage_pool" "st_pool" {
  name = "CPG-SSD-AO"
}

data "oneview_storage_volume_template" "st_vt" {
  name = "RenameDemoStorageTemplate"
}

# Creates volume resource
resource "oneview_volume" "volume" {
  properties {
    name              = "testvol"
     /* uncomment the below two line to provide the size and storage pool uri or it will be taken from the volume template*/
    //storage_pool      = data.oneview_storage_pool.st_pool.uri
    //size              = 1221225472/
    provisioning_type = "Thin" 
    is_deduplicated   = false
  }
  template_name        = "RenameDemoStorageTemplate"
  template_uri         = data.oneview_storage_volume_template.st_vt.uri
  name                 = "testvol"
  description          = "Test Volume"
  initial_scope_uris   = [data.oneview_scope.scope_obj.uri]
  provisioned_capacity = "268435456"
}

