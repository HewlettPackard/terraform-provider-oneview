provider "oneview" {
  ov_username = "<ov_username>"
  ov_password = "<ov_password>"
  ov_endpoint = "<ov_endpoint>"
  ov_sslverify = false
  ov_apiversion = <ov_apiversion>
  ov_ifmatch = "*"
}

data "oneview_scope" "scope_obj" {
        name = "test"
}

data "oneview_storage_pool" "st_pool" {
        name = "CPG_FC-AO"
}
data "oneview_storage_volume_template" "st_vt" {
        name = "tr_vt"
}

resource "oneview_volume" "volume" {
  properties = {
    "name" = "testvol"
    "storage_pool"= "${data.oneview_storage_pool.st_pool.uri}",
    "size"= 268435456,
    "provisioning_type"= "Thin",
    "is_deduplicated"= false
  }
  template_uri= "${data.oneview_storage_volume_template.st_vt.uri}",
  is_permanent= true,
  name = "testvol"
  description = "Test Volume"
  initial_scope_uris = ["${data.oneview_scope.scope_obj.uri}"]
  provisioned_capacity = "268435456"
}

/* Datasource

data "oneview_volume" "volume" {
  name = "testvol2"
}

output "oneview_volume_value" {
  value = "${data.oneview_volume.volume.uri}"
}
*/
