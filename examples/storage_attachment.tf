provider "oneview" {
        ov_username = "<ov_username>"
        ov_password = "<ov_password"
        ov_endpoint = "<ov_endpoint>"
        ov_sslverify = false
        ov_apiversion = <ov_apiversion>
        ov_ifmatch = "*"
}

data "oneview_storage_attachment" "storage_attach" {
        name = "<name>"
}

output "oneview_storage_attachment" {
        value = "${data.oneview_storage_attachment.storage_attach.uri}"
}
