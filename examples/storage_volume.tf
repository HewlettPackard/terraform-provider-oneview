provider "oneview" {
  ov_username = "<ov_username>"
  ov_password = "<ov_password>"
  ov_endpoint = "<ov_endpoint>"
  ov_sslverify = false
  ov_apiversion = <ov_apiversion>
  ov_ifmatch = "*"
}

resource "oneview_volume" "volume" {
  properties = {
    "name" = "testvol"
    "storage_pool"= "/rest/storage-pools/4EF64D4-FB48-4209-8790-AB200070738C",
    "size"= 268435456,
    "provisioning_type"= "Thin",
    "is_deduplicated"= false
  }
  template_uri= "/rest/storage-volume-templates/01953309-b02e-47d2-921b-aaaf0099d392",
  is_permanent= true,
  name = "testvol"
  description = "Test Volume"
  initial_scope_uris = ["/rest/scopes/4dff2b83-edb0-4629-a002-4a6a18951115"]
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
