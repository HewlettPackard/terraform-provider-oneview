# terraform-provider-oneview

[![Build Status](https://travis-ci.org/HewlettPackard/terraform-provider-oneview.svg?branch=master)](https://travis-ci.org/HewlettPackard/terraform-provider-oneview)

A Terraform provider for oneview

## Installation

* Install Go 1.8. For previous versions, you may have to set your `$GOPATH` manually, if you haven't done it yet.
* Install Terraform 0.9.x or above [from here](https://www.terraform.io/downloads.html) and save it into `/usr/local/bin/terraform` folder (create it if it doesn't exists)
* Download the code by issuing a `go get` command.

```bash
# Download the source code for terraform-provider-oneview
# and build the needed binary, by saving it inside $GOPATH/bin
go get -u github.com/HewlettPackard/terraform-provider-oneview

# Copy the binary to have it along the terraform binary
mv $GOPATH/bin/terraform-provider-oneview /usr/local/bin/terraform
```

## Example terraform file to provision a server with an operating system
```
provider "oneview" {
  ov_username   = "Administrator"
  ov_password   = "thisisapassword"
  ov_endpoint   = "https://oneview_instance.com"
}

resource "oneview_server_profile" "default" {
  name              = "test"
  template          = "Web Server Template"

  //specify this value so icsp can find the public ip address.
  public_connection = "pub_conn_1"
}

resource "icsp_server" "default" {
  ilo_ip = "15.x.x.x"
  user_name = "ilo_user"
  password = "ilo_password"
  serial_number = "${oneview_server_profile.default.serial_number}"
  build_plans = ["/rest/os-deployment-build-plans/1570001"]

  //this attribute gets you the public ip address
  public_mac = "${oneview_server_profile.default.public_mac}"
}
```

More information about how to configure the provider can be found [here](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/index.html.markdown)

## Resources
Any resource that OneView can manage is on the roadmap for Terraform to also manage. Below is the current list of resources that Terraform can manage. Open an issue if there is a resource that needs to be developed as soon as possible.

#### [Server Profile](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/server_profile.html.markdown)

```js
resource "oneview_server_profile" "default" {
  name = "test-server-profile"
  template = "${oneview_server_profile_template.test.name}"
}
```

#### [ICSP Server](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/icsp_server.html.markdown)
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

#### [Image Streamer Deployment Plan](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/i3s_plan.html.markdown)
This block takes an already provisioned server and through Image Streamer lays down an Operating System.

```js
resource "oneview_i3s_plan" "default" {
  server_name = "${oneview_server_profile.default.name}"
  os_deployment_plan = "Ubuntu 16.04"
  deploy_net_name = "I3S-Deploy-v301"
}
```

#### [Server Profile Template](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/server_profile_template.html.markdown)

```js
resource "oneview_server_profile_template" "default" {
  name = "test-server-profile-template"
  enclosure_group = "my_enclosure_group"
  server_hardware_type = "BL460c Gen9 1"
}
```

#### [Ethernet Network](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/ethernet_network.html.markdown)

```js
resource "oneview_ethernet_network" "default" {
  name = "test-ethernet-network"
  vlanId = 71
}
```

#### [Fibre Channel Network](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/fc_network.html.markdown)

```js
resource "oneview_fc_network" "default" {
  name = "test-fc-network"
}
```

#### [Fibre Channel over Ethernet Network](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/fcoe_network.html.markdown)

```js
resource "oneview_fcoe_network" "default" {
  name = "test-fcoe-network"
  vlanId = 71
}
```

#### [Network Set](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/network_set.html.markdown)

```js
resource "oneview_network_set" "default" {
  name = "test-network-set"
  network_uris = ["${oneview_ethernet_network.default.*.uri}"]
  native_network_uri = "${oneview_ethernet_network.default.1.uri}"
}
```

#### [Logical Interconnect Group](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/logical_interconnect_group.html.markdown)


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

#### [Enclosure Group](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/enclosure_group.html.markdown)

```js
resource "oneview_enclosure_group" "default" {
  name = "default-enclosure-group"
  logical_interconnect_groups = ["${oneview_logical_interconnect_group.primary.name}",
                                 "${oneview_logical_interconnect_group.secondary.name}"]
}
```

#### [Logical Switch](https://github.com/HewlettPackard/terraform-provider-oneview/blob/master/docs/r/logical_switch.html.markdown)

```js
resource "oneview_logical_switch" "default" {
  name = "test-logical-switch"
  switch_type_name = "Cisco Nexus 6xxx"
  switch_count = 1
}
```

### License

This project is licensed under the Apache 2.0 license.
