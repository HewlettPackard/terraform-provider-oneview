provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = var.api_version
  ov_ifmatch    = "*"
}

# Gets The Server Certificates 
# Any one of the these fields can used to get data source

data "oneview_server_certificate" "sc" {
  #        alias_name = "hm_cert"
  remote_ip = "<Server_IP>"
}

output "oneview_server_certificate_value" {
  value =  element(tolist(data.oneview_server_certificate.sc.certificate_details[*].base64_data), 0)
}

# Adds the Certificates to Oneview Appliance
resource "oneview_server_certificate" "ServerCertificate" {
  certificate_details {
    //base64_data = data.oneview_server_certificate.sc.certificate_details[0].base64_data
    base64_data = element(tolist(data.oneview_server_certificate.sc.certificate_details[*].base64_data), 0)

    type        = "CertificateDetailV2"
    alias_name  = "TestServerCertificate"
  }
}

