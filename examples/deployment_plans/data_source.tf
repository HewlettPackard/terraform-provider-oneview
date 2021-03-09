provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  i3s_endpoint  = var.i3s_endpoint
  ov_apiversion = 2010
  ov_ifmatch    = "*"
}

# Test for data source

data "oneview_deployment_plan" "deployment_plan" {
  name = "TestDP"
}

output "oneview_deployment_plan_value" {
  value = data.oneview_deployment_plan.deployment_plan.oe_build_plan_uri
}

# Importing an existing resource from teh appliance
/*
resource "oneview_deployment_plan" "dp_inst"{
}
*/
