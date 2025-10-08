provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Test configuration to verify data source with task URI
data "oneview_server_profile_async_task" "test_by_uri" {
  task_uri_input = "/rest/tasks/ce16faf1-865b-44f2-8bc5-3c397c2b8410"
}

output "test_task_details" {
  value = {
    uri    = data.oneview_server_profile_async_task.test_by_uri.task_uri
    state  = data.oneview_server_profile_async_task.test_by_uri.task_state
    status = data.oneview_server_profile_async_task.test_by_uri.task_status
    progress = data.oneview_server_profile_async_task.test_by_uri.percent_complete
    percent_complete = data.oneview_server_profile_async_task.test_by_uri.percent_complete
  }
}