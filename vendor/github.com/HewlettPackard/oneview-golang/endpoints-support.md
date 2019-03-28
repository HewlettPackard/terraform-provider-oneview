***Legend***

| Item | Meaning |
| ------------------ | --------------------------------------------- |
|  :white_check_mark: | Endpoint implemented in the GO SDK and tested for this API version :tada: |
|  :heavy_multiplication_x:  | Endpoint considered as 'out-of-scope' for the GO SDK.              |
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
|     **Enclosure Groups**                                                                                                                          |
|<sub>/rest/enclosure-groups</sub>                                                        | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/enclosure-groups</sub>                                                        | POST     | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/enclosure-groups/{id}</sub>                                                   | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/enclosure-groups/{id}</sub>                                                   | PUT      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/enclosure-groups/{id}</sub>                                                   | DELETE   | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/enclosure-groups/{id}/script</sub>                                            | GET      | :heavy_minus_sign:   | :heavy_minus_sign:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/enclosure-groups/{id}/script</sub>                                            | PUT      | :heavy_minus_sign:   | :heavy_minus_sign:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|     **Ethernet Networks**                                                                                                                         |
|<sub>/rest/ethernet-networks</sub>                                                       | GET      | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/ethernet-networks</sub>                                                       | POST     | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/ethernet-networks/bulk</sub>                                                  | POST     | :heavy_minus_sign:   | :heavy_minus_sign:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/ethernet-networks/{id}</sub>                                                  | GET      | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/ethernet-networks/{id}</sub>                                                  | PUT      | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/ethernet-networks/{id}</sub>                                                  | PATCH    | :heavy_minus_sign:   | :heavy_minus_sign:   | :heavy_minus_sign:         | :heavy_minus_sign:   | :heavy_minus_sign:   |
|<sub>/rest/ethernet-networks/{id}</sub>                                                  | DELETE   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/ethernet-networks/{id}/associatedProfiles</sub>                               | GET      | :heavy_minus_sign:   | :heavy_minus_sign:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/ethernet-networks/{id}/associatedUplinkGroups</sub>                           | GET      | :heavy_minus_sign:   | :heavy_minus_sign:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   |
|     **FC Networks**                                                                                                                               |
|<sub>/rest/fc-networks</sub>                                                             | GET      | :white_check_mark:   | :white_check_mark:   | :white_check_mark:         | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/fc-networks</sub>                                                             | POST     | :white_check_mark:   | :white_check_mark:   | :white_check_mark:         | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/fc-networks/{id}</sub>                                                        | GET      | :white_check_mark:   | :white_check_mark:   | :white_check_mark:         | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/fc-networks/{id}</sub>                                                        | PATCH    | :heavy_minus_sign:   | :heavy_minus_sign:   | :heavy_multiplication_x:   | :heavy_minus_sign:   | :heavy_minus_sign:   |
|<sub>/rest/fc-networks/{id}</sub>                                                        | PUT      | :white_check_mark:   | :white_check_mark:   | :white_check_mark:         | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/fc-networks/{id}</sub>                                                        | DELETE   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:         | :white_check_mark:   | :white_check_mark:   |
|     **FCoE Networks**                                                                                                                             |
|<sub>/rest/fcoe-networks</sub>                                                           | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/fcoe-networks</sub>                                                           | POST     | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/fcoe-networks/{id}</sub>                                                      | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/fcoe-networks/{id}</sub>                                                      | PATCH    | :heavy_minus_sign:   | :heavy_minus_sign:   | :heavy_multiplication_x:   |
|<sub>/rest/fcoe-networks/{id}</sub>                                                      | PUT      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/fcoe-networks/{id}</sub>                                                      | DELETE   | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|     **Interconnect Types**                                                                                                                        |
|<sub>/rest/interconnect-types</sub>                                                      | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/interconnect-types/{id}</sub>                                                 | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|     **Logical Interconnect Groups**                                                                                                               |
|<sub>/rest/logical-interconnect-groups</sub>                                             | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/logical-interconnect-groups</sub>                                             | POST     | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/logical-interconnect-groups/{id}</sub>                                        | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/logical-interconnect-groups/{id}</sub>                                        | PUT      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/logical-interconnect-groups/{id}</sub>                                        | PATCH    | :heavy_minus_sign:   | :heavy_minus_sign:   | :heavy_multiplication_x:   |
|<sub>/rest/logical-interconnect-groups/{id}</sub>                                        | DELETE   | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|     **Logical Switch Groups**                                                                                                                     |
|<sub>/rest/logical-switch-groups</sub>                                                   | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/logical-switch-groups</sub>                                                   |POST      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/logical-switch-groups/{id}</sub>                                              |GET       | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/logical-switch-groups/{id}</sub>                                              |PUT       | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/logical-switch-groups/{id}</sub>                                              |DELETE    | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|     **Logical Switches**                                                                                                                         |
|<sub>/rest/logical-switches</sub>                                                        | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/logical-switches</sub>                                                        |POST      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/logical-switches/{id}</sub>                                                   |GET       | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/logical-switches/{id}</sub>                                                   |PUT       | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/logical-switches/{id}</sub>                                                   |DELETE    | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/logical-switches/{id}/refresh</sub>                                           |PUT       | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|     **Network Sets**                                                                                                                              |
|<sub>/rest/network-sets</sub>                                                            | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/network-sets</sub>                                                            | POST     | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/network-sets/withoutEthernet</sub>                                            | GET      | :heavy_minus_sign:   | :heavy_minus_sign:   | :heavy_multiplication_x:   |
|<sub>/rest/network-sets/{id}</sub>                                                       | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/network-sets/{id}</sub>                                                       | PUT      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/network-sets/{id}</sub>                                                       | DELETE   | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/network-sets/{id}</sub>                                                       | PATCH    | :heavy_minus_sign:   | :heavy_minus_sign:   | :heavy_multiplication_x:   |
|     **OS Deployment Plans**                                                                                                                      |
|<sub>/rest/os-deployment-plans/</sub>                                                    | GET      | :heavy_minus_sign:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/os-deployment-plans/{id}</sub>                                                | GET      | :heavy_minus_sign:   | :white_check_mark:   | :heavy_multiplication_x:   |
|     **Power Devices**                                                                                                                             |
|<sub>/rest/power-devices</sub>                                                           | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/power-devices/{id}</sub>                                                      | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|     **SAN Managers**                                                                                                                             |
|<sub>/rest/fc-sans/device-managers</sub>                                                 | GET      | :white_check_mark:   | :heavy_minus_sign:   | :heavy_multiplication_x:   |
|<sub>/rest/fc-sans/device-managers/{id}</sub>                                            | GET      | :white_check_mark:   | :heavy_minus_sign:   | :heavy_multiplication_x:   |
|<sub>/rest/fc-sans/device-managers/{id}</sub>                                            | PUT      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/fc-sans/device-managers/{id}</sub>                                            | DELETE   | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/fc-sans/providers/{id}/device-managers</sub>                                  | POST     | :white_check_mark:   | :heavy_minus_sign:   | :heavy_multiplication_x:   |
|     **Server Hardware**                                                                                                                          |
|<sub>/rest/server-hardware</sub>                                                         | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/server-hardware/{id}</sub>                                                    | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|     **Server Hardware Types**                                                                                                                     |
|<sub>/rest/server-hardware-types</sub>                                                   | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/server-hardware-types/{id}</sub>                                              | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/server-hardware-types/{id}</sub>                                              | PUT      | :heavy_minus_sign:   | :heavy_minus_sign:   | :heavy_minus_sign:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   |
|<sub>/rest/server-hardware-types/{id}</sub>                                              | DELETE   | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   | :heavy_multiplication_x:   |
|     **Server Profile Templates**                                                                                                                  |
|<sub>/rest/server-profile-templates</sub>                                                | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/server-profile-templates</sub>                                                | POST     | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/server-profile-templates/{id}</sub>                                           | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/server-profile-templates/{id}</sub>                                           | PUT      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/server-profile-templates/{id}</sub>                                           | DELETE   | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|     **Server Profiles**                                                                                                                           |
|<sub>/rest/server-profiles</sub>                                                         | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/server-profiles</sub>                                                         | POST     | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/server-profiles</sub>                                                         | DELETE   | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
<sub>/rest/server-profiles</sub>                                                          | PUT      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|     **Storage Pools**                                                                                                                                   |
|<sub>/rest/storage-pools</sub>                                                           | GET      | :heavy_multiplication_x: | :heavy_multiplication_x: | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | 
|<sub>/rest/storage-pools/{id}</sub>                                                      | GET      | :heavy_multiplication_x: | :heavy_multiplication_x: | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/storage-pools/{id}</sub>                                                      | PUT      | :heavy_multiplication_x: | :heavy_multiplication_x: | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   |
|     **Switch Types**                                                                                                                              |
|<sub>/rest/switch-types</sub>                                                            | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|<sub>/rest/switch-types/{id}</sub>                                                       | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|     **Storage Volume Attachments**                                                                                                                                   |
|<sub>/rest/storage-volume-attachments</sub>                                              | GET      | :heavy_multiplication_x:  | :heavy_multiplication_x: | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | 
|<sub>/rest/storage-volume-attachments/{id}</sub>                                         | GET      | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   |
|     **Tasks**                                                                                                                                     |
|<sub>/rest/tasks</sub>                                                                   | GET      | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   |
|     **Version**                                                                                                                                   |
|<sub>/rest/version</sub>                                                                 | GET      | :white_check_mark:   | :white_check_mark:   | :heavy_multiplication_x:   |
|     **Volumes**                                                                                                                                   |
|<sub>/rest/storage-volumes</sub>                                                         | GET      | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | 
|<sub>/rest/storage-volumes/{id}</sub>                                                    | GET      | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/storage-volumes</sub>                                                         | POST     | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | 
|<sub>/rest/storage-volumes/{id}</sub>                                                    | PUT      | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   |
|<sub>/rest/storage-volumes/{id}</sub>                                                    | DELETE   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   | :white_check_mark:   |
