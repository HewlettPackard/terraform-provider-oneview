provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Testing data source

data "oneview_task" "tasks" {
	filter="taskState='Running'"
}

resource "oneview_task" "update_patch" {
	uri = data.oneview_task.tasks.uri
}
