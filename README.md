# terraform-provider-oneview
A Terraform provider for oneview

## Installation 

* Install golang 1.7 or better
* Install Terraform 0.8.0 or better
* Create a file structure similar to ~/workspace/go/src/github.com/HewlettPackard/

```
  export GOPATH = ~/worspace/go
  cd ~/workspace/go/src/github.com/HewlettPackard/
  git clone https://github.com/HewlettPackard/terraform-provider-oneview.git
  cd terraform-provider-oneview
  go get
  go build -o terraform-provider-oneview
  mv terraform-provider-oneview /usr/local/bin/terraform
```

## Example terraform file to provision a server
```
provider "oneview" {
  ov_username   = "Administrator"
  ov_password   = "thisisapassword"
  ov_endpoint   = "https://oneview_instance.com"
}

resource "oneview_server_profile" "test" {
  name             = "test"
  template         = "Web Server Template"
}
```
More information about how to configure the provider can be found [here](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/index.html.markdown)

## Resources
Any resource that OneView can manage is on the roadmap for Terraform to also manage. Below is the current list of resources that Terraform can manage. Open an issue if there is a resource that needs to be developed as soon as possible. 

####[Server Profile](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/server_profile.html.markdown)

```js
resource "oneview_server_profile" "default" {
  name = "test-server-profile"
  template = "${oneview_server_profile_template.test.name}"
}
```

####[ICSP Server](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/icsp_server.html.markdown)
This block takes an already provsioned server and through ICSP lays down an operating system or 
whatever is specified in the build plans.

```js
resource "icsp_server" "default" {
  ilo_ip = "15.x.x.x"
  user_name = "ilo_user"
  password = "ilo_password"
  serial_number = "${oneview_server_profile.default.serial_number}"
  build_plans = ["/rest/os-deployment-build-plans/1570001"]
}
```

####[Image Streamer Deployment Plan](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/i3s_plan.html.markdown)
This block takes an already provisioned server and through Image Streamer lays down an Operating System. 

```js
resource "oneview_i3s_plan" "default" {
  server_name = "${oneview_server_profile.default.name}"
  os_deployment_plan = "Ubuntu 16.04"
}
```

####[Server Profile Template](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/server_profile_template.html.markdown)

```js
resource "oneview_server_profile_template" "default" {
  name = "test-server-profile-template"
  enclosure_group = "my_enclosure_group"
  server_hardware_type = "BL460c Gen9 1"
}
```

####[Ethernet Network](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/ethernet_network.html.markdown)

```js
resource "oneview_ethernet_network" "default" {
  name = "test-ethernet-network"
  vlanId = 71
}
```

####[Fibre Channel Network](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/fc_network.html.markdown)

```js
resource "oneview_fc_network" "default" {
  name = "test-fc-network"
}
```

####[Fibre Channel over Ethernet Network](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/fcoe_network.html.markdown)

```js
resource "oneview_fcoe_network" "default" {
  name = "test-fcoe-network"
  vlanId = 71
}
```

####[Network Set](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/network_set.html.markdown)

```js
resource "oneview_network_set" "default" {
  name = "test-network-set"
  network_uris = ["${oneview_ethernet_network.default.*.uri}"]
  native_network_uri = "${oneview_ethernet_network.default.1.uri}"
}
```

####[Logical Interconnect Group](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/logical_interconnect_group.html.markdown)


```js
resource "oneview_logical_interconnect_group" "default" {
  name = "test-logical-interconnect-group"
  
  internal_network_uris = ["${oneview_ethernet_network.default.0.uri}"]
  
  interconnect_map_entry_template {
    interconnect_type_name = "HP VC FlexFabric-20/40 F8 Module"
    bay_number = 1
  }
  
  uplink_set {
    name = "uplink-default"
    network_uris = ["${oneview_ethernet_network.test.1.uri}"]
    logical_port_config {
      bay_num = 4
      port_num = [20,21]
    }
  }
}
```

####[Enclosure Group](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/enclosure_group.html.markdown)

```js
resource "oneview_enclosure_group" "default" {
  name = "default-enclosure-group"
  logical_interconnect_groups = ["${oneview_logical_interconnect_group.primary.name}", 
                                 "${oneview_logical_interconnect_group.secondary.name}"]
}
```

####[Logical Switch](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/logical_switch.html.markdown)

```js
resource "oneview_logical_switch" "default" {
  name = "test-logical-switch"
  switch_type_name = "Cisco Nexus 6xxx"
  switch_count = 1
}
```

### License

This project is licensed under the Apache 2.0 license.
