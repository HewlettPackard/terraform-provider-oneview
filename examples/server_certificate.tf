provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 2400
  ov_ifmatch    = "*"
}

data "oneview_scope" "scope_obj" {
  name = "testing"
}

# Testing data source 
data "oneview_server_certificate" "sc" {
  # Any one of the these fields can used to get data source
  #         alias_name = "hm_cert"
  remote_ip = "172.18.13.11"
}

output "oneview_server_certificate_value" {
  value = data.oneview_server_certificate.sc.certificate_details[0].base64_data
}


resource "oneview_server_certificate" "ServerCertificate" {
  certificate_details {
    base64_data = data.oneview_server_certificate.sc.certificate_details[0].base64_data
    type        = "CertificateDetailV2"
    alias_name  = "TestServerCertificate"
  }
}

/*
resource "oneview_server_certificate" "ServerCertificate" {
  certificate_details {
    base64_data = " -----BEGIN CERTIFICATE-----\nMIIEKTCCAxGgAwIBAgIJAMZHZwLUTbU9MA0GCSqGSIb3DQEBCwUAMFoxCzAJBgNV\nBAYTAklOMQswCQYDVQQIDAJLQTELMAkGA1UEBwwCQkExDDAKBgNVBAoMA0hQRTEM\nMAoGA1UECwwDRU1MMRUwEwYDVQQDDAwxNzIuMTguMTMuMTEwIBcNMjEwMTA0MTEx\nNjAxWhgPMjEyMDEyMTExMTE2MDFaMFoxCzAJBgNVBAYTAklOMQswCQYDVQQIDAJL\nQTELMAkGA1UEBwwCQkExDDAKBgNVBAoMA0hQRTEMMAoGA1UECwwDRU1MMRUwEwYD\nVQQDDAwxNzIuMTguMTMuMTEwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIB\nAQCd+XftwvAJRe1yLlocD7Fa0Sd8sRCkxjd3NBLNgpREik4i1e9vhPQuh033oSOx\nIgZVlpZMHRfoEOHCUMTqTQFuBVtRkVRiUTUbhN2X25Pu6XWTCsPvoTqE31Lb2Z3d\n6e6FoZHFCsvjP7f3/oBEydP+/LpZxs8NK85CB/mO1HKjeKZeQhOJ0vvkPWuEi1JO\nuR+B+CR99geWLTKSfRiS4ST9AiIr2WS7ev2HSLS5Xggoo9i71x4YJ26lyiKv64C/\ngnAqjoygeXMK9Sa8OcPL/xWTuN/gykkocr8sxegLrmo2iln77RK5zYYjDVLCc1i5\nLsD9S/VKEcivAzShUfLdyHJnAgMBAAGjge8wgewwHQYDVR0OBBYEFJa+Mep5B5C8\nTSnBcaQoaaaHKLKWMIGMBgNVHSMEgYQwgYGAFJa+Mep5B5C8TSnBcaQoaaaHKLKW\noV6kXDBaMQswCQYDVQQGEwJJTjELMAkGA1UECAwCS0ExCzAJBgNVBAcMAkJBMQww\nCgYDVQQKDANIUEUxDDAKBgNVBAsMA0VNTDEVMBMGA1UEAwwMMTcyLjE4LjEzLjEx\nggkAxkdnAtRNtT0wCwYDVR0PBAQDAgSwMAkGA1UdEwQCMAAwDwYDVR0RBAgwBocE\nrBINCzATBgNVHSUEDDAKBggrBgEFBQcDATANBgkqhkiG9w0BAQsFAAOCAQEAF36W\n7TnObJJPYEQdjg/X2oa+c5rmmZOcjw0iWcKiJYkTc978q4NgP+0GKac2N12bAFCl\nVhLtAVBc1FKPQLczOw/7jgtE+gyd4aWZzVmcRVmQWH4iQDDN3lzrXoKufhYx6xcv\nsSfcgkFRht03e3qoQ4JtHafD+tZanaYY3wm5TOZiTwGmhD3RCy+uEZav5vaFV7d4\nSLqiqYDM8dbbddKyEjw1LBJ0NgHzprgqr04QKm0J2/XX71SXO68ecxNeVveHotIh\nAQBPrvC3oKaFPq/zF/WxKIhF615c2Em3fWwFH3v+fraxGc9l1SNo5x3fQvqZMd6X\nvoVtu7MnHIXwWBpO8Q==\n-----END CERTIFICATE-----"
           
    type        = "CertificateDetailV2"
    alias_name  = "TestServerCertificate"
  }
}
*/

# Import existing certificate
 /*resource "oneview_server_certificate" "ServerCertificate" {
 }
 */
