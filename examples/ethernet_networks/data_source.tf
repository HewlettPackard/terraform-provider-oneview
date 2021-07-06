provider "oneview" {
        ov_username =   "${var.username}"
        ov_password =   "${var.password}"
        ov_endpoint =   "${var.endpoint}"
        ov_sslverify =  "${var.ssl_enabled}"
        ov_apiversion = 3000
        ov_ifmatch = "*"

}

# Testing data source
data "oneview_ethernet_network" "ethernetnetworks" {
  name = "TestEthNetwork_terraform_Rename"
}

#Importing Existing resource
#resource "oneview_ethernet_network" "import_eth"{
#}
