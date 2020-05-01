provider "oneview" {
    ov_username = "<username>"
	ov_password = "<password>"
	ov_endpoint = "<endpoint>"
	ov_sslverify = false
	ov_apiversion = <ov-apiversion>
	ov_ifmatch = "*"
}

data "oneview_scope" "scope_obj" {
        name = "test_scope"
}

resource "oneview_server_certificate" "ServerCertificate" {
    certificate_details = [{
                        base64_data="-----BEGIN CERTIFICATE-----\nMIIEKTCCAxGgAwIBAgIJAJNHxltN7DJuMA0GCSqGSIb3DQEBCwUAMFoxCzAJBgNV\nBAYTAklOMQswCQYDVQQIDAJLQTELMAkGA1UEBwwCQkExDDAKBgNVBAoMA0hQRTEM\nMAoGA1UECwwDRU1MMRUwEwYDVQQDDAwxNzIuMTguMTMuMTEwIBcNMjAwNDAzMDgw\nMzM4WhgPMjEyMDAzMTAwODAzMzhaMFoxCzAJBgNVBAYTAklOMQswCQYDVQQIDAJL\nQTELMAkGA1UEBwwCQkExDDAKBgNVBAoMA0hQRTEMMAoGA1UECwwDRU1MMRUwEwYD\nVQQDDAwxNzIuMTguMTMuMTEwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIB\nAQDfZ2u+s0R3CVpPAAGGbUDkuw4zueh0IM10Fma7iw5zJrxZqchDKzPLRxPMlIUq\nHcHWtKOANNd23r5kTcSJply3oo+P6qrhc9B2PtzJMKlUpgAaga8s5Ii/czK6P2X9\nlEksLkqaluyXQmec1AjvhVpkzbgp8FdSihbjENRFzEJDLPmmPDzS5IhgmWNCNSDf\nQaDFHmp938lP1CTHt4l57Xz71+3cx+xHEL6yqvH1CKux3RGVVaSZocJwLRVCUPbr\n+SjrX1tJ3Mkd6o0WarBstL5n+3UwKOphEPaha2J/DLpasoRXbkVi24jtttHUChV0\nUBnNcV2SSWtgWC9jidHOIQwpAgMBAAGjge8wgewwHQYDVR0OBBYEFLUnePCMBEVu\nMDMdRp2g8US0SctqMIGMBgNVHSMEgYQwgYGAFLUnePCMBEVuMDMdRp2g8US0Sctq\noV6kXDBaMQswCQYDVQQGEwJJTjELMAkGA1UECAwCS0ExCzAJBgNVBAcMAkJBMQww\nCgYDVQQKDANIUEUxDDAKBgNVBAsMA0VNTDEVMBMGA1UEAwwMMTcyLjE4LjEzLjEx\nggkAk0fGW03sMm4wCwYDVR0PBAQDAgSwMAkGA1UdEwQCMAAwDwYDVR0RBAgwBocE\nrBINCzATBgNVHSUEDDAKBggrBgEFBQcDATANBgkqhkiG9w0BAQsFAAOCAQEAVBFX\nDc4sbHbSSJYeCxWQFo5DVUKqaWUewmMZPcGfia1BzltC1JDvpUqdE5NOGbr8xS+8\nnX+pNcO/qfMzpQrwtE8kp4KlB69rcPAUNyZ5o/tJ+Qew45j6hqPsoIZ16EmqZscq\n6AAqbXZeyixbNCTspJa38BOJFdGVhdjTYaxWQOZBz8eWK9DRNFpfkOWiIXOiyDkd\nRXMh17YfhADjvbt/npBw1BCCXXERl3hBz4g6GbZJhAPtfd99sZO01id6RBABgzng\n3fLtFd48SUnrCnxOZURzsU9PHyzhBqyLm6bFZ210ZDL21ugszCewgz2BcZXHzZxi\nWHkHstA1BIXrYjCeYQ==\n-----END CERTIFICATE-----"
                        type="CertificateDetailV2"
                        alias_name = "TestServerCertificate"
                        }]

}

/*
resource "oneview_server_certificate" "ServerCertificate" {
    certificate_details = [{
                        base64_data="-----BEGIN CERTIFICATE-----\nMIIEKtCCAxGgAwIBAgIJAJNHxltN7DJuMA0GCSqGSIb3DQEBCwUAMFoxCzAJBgNV\nBAYTAklOMQswCQYDVQQIDAJLQTELMAkGA1UEBwwCQkExDDAKBgNVBAoMA0hQRTEM\nMAoGA1UECwwDRU1MMRUwEwYDVQQDDAwxNzIuMTguMTMuMTEwIBcNMjAwNDAzMDgw\nMzM4WhgPMjEyMDAzMTAwODAzMzhaMFoxCzAJBgNVBAYTAklOMQswCQYDVQQIDAJL\nQTELMAkGA1UEBwwCQkExDDAKBgNVBAoMA0hQRTEMMAoGA1UECwwDRU1MMRUwEwYD\nVQQDDAwxNzIuMTguMTMuMTEwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIB\nAQDfZ2u+s0R3CVpPAAGGbUDkuw4zueh0IM10Fma7iw5zJrxZqchDKzPLRxPMlIUq\nHcHWtKOANNd23r5kTcSJply3oo+P6qrhc9B2PtzJMKlUpgAaga8s5Ii/czK6P2X9\nlEksLkqaluyXQmec1AjvhVpkzbgp8FdSihbjENRFzEJDLPmmPDzS5IhgmWNCNSDf\nQaDFHmp938lP1CTHt4l57Xz71+3cx+xHEL6yqvH1CKux3RGVVaSZocJwLRVCUPbr\n+SjrX1tJ3Mkd6o0WarBstL5n+3UwKOphEPaha2J/DLpasoRXbkVi24jtttHUChV0\nUBnNcV2SSWtgWC9jidHOIQwpAgMBAAGjge8wgewwHQYDVR0OBBYEFLUnePCMBEVu\nMDMdRp2g8US0SctqMIGMBgNVHSMEgYQwgYGAFLUnePCMBEVuMDMdRp2g8US0Sctq\noV6kXDBaMQswCQYDVQQGEwJJTjELMAkGA1UECAwCS0ExCzAJBgNVBAcMAkJBMQww\nCgYDVQQKDANIUEUxDDAKBgNVBAsMA0VNTDEVMBMGA1UEAwwMMTcyLjE4LjEzLjEx\nggkAk0fGW03sMm4wCwYDVR0PBAQDAgSwMAkGA1UdEwQCMAAwDwYDVR0RBAgwBocE\nrBINCzATBgNVHSUEDDAKBggrBgEFBQcDATANBgkqhkiG9w0BAQsFAAOCAQEAVBFX\nDc4sbHbSSJYeCxWQFo5DVUKqaWUewmMZPcGfia1BzltC1JDvpUqdE5NOGbr8xS+8\nnX+pNcO/qfMzpQrwtE8kp4KlB69rcPAUNyZ5o/tJ+Qew45j6hqPsoIZ16EmqZscq\n6AAqbXZeyixbNCTspJa38BOJFdGVhdjTYaxWQOZBz8eWK9DRNFpfkOWiIXOiyDkd\nRXMh17YfhADjvbt/npBw1BCCXXERl3hBz4g6GbZJhAPtfd99sZO01id6RBABgzng\n3fLtFd48SUnrCnxOZURzsU9PHyzhBqyLm6bFZ210ZDL21ugszCewgz2BcZXHzZxi\nWHkHstA1BIXrYjCeYQ==\n-----END CERTIFICATE-----"
                        type="CertificateDetailV2"
                        alias_name = "TestServerCertificate"
                        }]
}
*/
/* Testing data source

data "oneview_server_certificate" "sc" {
/* Any one of the these fields can used to get data source
         alias_name = "hm_cert"
         /*remote_ip = "172.18.13.11"*/
}
output "oneview_server_certificate_value" {
        value = "${data.oneview_server_certificate.sc.type}"
}
*/
