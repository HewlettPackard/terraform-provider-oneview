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
task_id = "5376a1fc-7967-4e39-a769-4011a33ef428"
}

output "oneview_task_by_value" {
  value = data.oneview_task.tasks
}

