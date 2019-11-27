***Legend***

| Item | Meaning |
| ------------------ | --------------------------------------------- |
|  :white_check_mark: | Endpoint implemented in the Terraform SDK and tested for this API version :tada: |
|  :heavy_multiplication_x:  | Endpoint considered as 'out-of-scope' for the Terraform SDK.              |
|  :heavy_minus_sign: | Endpoint not available for this API Version |

<br />

***Notes***

* If an endpoint is marked as implemented in a previous version of the API, it will likely already be working for newer API versions, however in these cases it is important to:
1. Specify the 'type' of the resource when using an untested API, as it will not get set by default
2. If an example is not working, verify the [HPE OneView REST API Documentation](http://h17007.www1.hpe.com/docs/enterprise/servers/oneview3.1/cic-api/en/api-docs/current/index.html)  for the API version being used, since the expected attributes for that resource might have changed.

<br />

## HPE OneView

| Endpoints                                                                       | Verb     | V200 | V300 | V500 |V600 |V800
| --------------------------------------------------------------------------------------- | -------- | :------------------: | :------------------: | :------------------: | :------------------: | :------------------: |
|     **Enclosure**                                                                                                                         |
|<sub>/rest/enclosure/{id}</sub>                                                          | GET      | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   |
|<sub>/rest/enclosure</sub>                                                               | POST     | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   |
|<sub>/rest/enclosure/{id}</sub>                                                          | PUT      | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   |
|<sub>/rest/enclosure/{id}</sub>                                                          | DELETE   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   |
|     **Enclosure Group**                                                                                                                  |
|<sub>/rest/enclosure-group</sub>                                                         | POST     | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   |
|<sub>/rest/enclosure-group/{id}</sub>                                                    | PUT      | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   |
|<sub>/rest/enclosure-group/{id}</sub>                                                    | DELETE   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   |
|     **Ethernet Networks**                                                                                                                         |
|<sub>/rest/ethernet-networks</sub>                                                       | POST     | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   |
|<sub>/rest/ethernet-networks/{id}</sub>                                                  | PUT      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   |
|<sub>/rest/ethernet-networks/{id}</sub>                                                  | DELETE   | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   |
|     **FC Networks**                                                                                                                               |
|<sub>/rest/fc-networks/{id}</sub>                                                        | GET      | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:         | :heavy_multiplication_x:   | :white_check_mark:   |
|<sub>/rest/fc-networks</sub>                                                             | POST     | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:         | :heavy_multiplication_x:   | :white_check_mark:   |
|<sub>/rest/fc-networks/{id}</sub>                                                        | PUT      | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:         | :heavy_multiplication_x:   | :white_check_mark:   |
|<sub>/rest/fc-networks/{id}</sub>                                                        | DELETE   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:         | :heavy_multiplication_x:   | :white_check_mark:   |
|     **FCoE Networks**                                                                                                                             |
|<sub>/rest/fcoe-networks</sub>                                                           | POST     | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/fcoe-networks/{id}</sub>                                                      | PUT      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/fcoe-networks/{id}</sub>                                                      | DELETE   | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|     **Interconnect Types**                                                                                                                           |
|<sub>/rest/Interconnect-types/{id}</sub>                                                 | GET      | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   |
|     **Interconnects**                                                                                                                          |
|<sub>/rest/interconnects/{id}</sub>                                                      | GET      | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|     **Logical Enclosure**                                                                                                                         |
|<sub>/rest/logical-enclosures/</sub>                                                     | POST     |  :heavy_multiplication_x:  | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/logical-enclosures/{id}</sub>                                                 | GET      | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/logical-enclosures/{id}</sub>                                                 | PUT      | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/logical-enclosures/{id}</sub>                                                 | DELETE   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:  | :white_check_mark:   | :white_check_mark:   |
|     **Logical Interconnects**                                                                                                               |
|<sub>/rest/logical-interconnects/{id}</sub>                                              | GET      | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:         | :heavy_multiplication_x:   | :white_check_mark:   |
|<sub>/rest/logical-interconnects/{id}/compliance</sub>                                   | PUT      | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:         | :heavy_multiplication_x:   | :white_check_mark:   |
|     **Logical Interconnect Groups**                                                                                                             |
|<sub>/rest/logical-interconnect-groups</sub>                                             | POST     | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/logical-interconnect-groups/{id}</sub>                                        | PUT      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/logical-interconnect-groups/{id}</sub>                                        | DELETE   | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|     **Logical Switch Groups**                                                                                                                     |
|<sub>/rest/logical-switch-groups</sub>                                                   |POST      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/logical-switch-groups/{id}</sub>                                              |PUT       | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/logical-switch-groups/{id}</sub>                                              |DELETE    | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|     **Network Sets**                                                                                                                              |
|<sub>/rest/network-sets/{id}</sub>                                                        | GET      | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:         | :heavy_multiplication_x:   | :white_check_mark:   |
|<sub>/rest/network-sets</sub>                                                            | POST     | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/network-sets/{id}</sub>                                                       | PUT      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/network-sets/{id}</sub>                                                       | DELETE   | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|     **Scope**                                                                                                                  |
|<sub>/rest/scopes</sub>                                                                  | POST     | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:         | :heavy_multiplication_x:   | :white_check_mark:   |
|<sub>/rest/scopes/{id}</sub>                                                             | PUT      | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:         | :heavy_multiplication_x:   | :white_check_mark:   |
|<sub>/rest/scopes/{id}</sub>                                                             | DELETE   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:         | :heavy_multiplication_x:   | :white_check_mark:   |
|     **Server Hardware**                                                                                                                           |
|<sub>/rest/server-hardware/{id}</sub>                                                    | GET      | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|     **Server Hardware Type**                                                                                                                  |
|<sub>/rest/server-hardware-type/{id}</sub>                                               | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|     **Server Profile Templates**                                                                                                                  |
|<sub>/rest/server-profile-templates/{id}</sub>                                           | GET      | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/server-profile-templates</sub>                                                | POST     | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/server-profile-templates/{id}</sub>                                           | PUT      | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/server-profile-templates/{id}</sub>                                           | DELETE   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   |
|     **Server Profiles**    
|<sub>/rest/server-profiles/{id}</sub>                                                    | GET      | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/server-profiles</sub>                                                         | POST     | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/server-profiles</sub>                                                         | DELETE   | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/server-profiles</sub>                                                         | PUT      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|     **Storage System**    
|<sub>/rest/storage-systems/{id}</sub>                                                    | GET      | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   |
|<sub>/rest/storage-systems</sub>                                                         | POST     | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   |
|<sub>/rest/storage-systems</sub>                                                         | DELETE   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   |
|<sub>/rest/storage-systems</sub>                                                         | PUT      | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   |
|     **Uplink Sets**                                                                                                                           |
|<sub>/rest/uplink-sets</sub>                                                             | GET      | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/uplink-sets</sub>                                                    	  | POST     | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/uplink-sets/{id}</sub>                                                        | DELETE   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/uplink-sets/{id}</sub>                                                        | PUT      | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |

## HPE Synergy Image Streamer

| Endpoints                                                                       | Verb     | V200 | V300 | V500 |V600 |V800
| --------------------------------------------------------------------------------------- | -------- | :------------------: | :------------------: | :------------------: | :------------------: | :------------------: |
|     **OS Deployment Plans**                                                                                                                           |
|<sub>/rest/os-deployment-plans</sub>                                                         | POST     | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   |
|<sub>/rest/os-deployment-plans</sub>                                                         | DELETE   | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   |
<sub>/rest/os-deployment-plans</sub>                                                          | PUT      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   |
