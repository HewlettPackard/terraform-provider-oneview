/*
(c) Copyright [2015] Hewlett Packard Enterprise Development LP

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ov

type BootOptionV3 struct {
	BootTargetLun        string `json:"bootTargetLun,omitempty"`        // "bootTargetLun": "0",
	BootTargetName       string `json:"bootTargetName,omitempty"`       // "bootTargetName": "iqn.2015-02.com.hpe:iscsi.target",
	BootVlanId           int    `json:"bootVlanId,omitempty"`           //The virtual LAN ID of the boot connection.
	BootVolumeSource     string `json:"bootVolumeSource,omitempty"`     // "bootVolumeSource": "",
	ChapLevel            string `json:"chapLevel,omitempty"`            // "chapLevel": "None",
	ChapName             string `json:"chapName,omitempty"`             // "chapName": "chap name",
	ChapSecret           string `json:"chapSecret,omitempty"`           // "chapSecret": "super secret chap secret",
	FirstBootTargetIp    string `json:"firstBootTargetIp,omitempty"`    // "firtBootTargetIp": "10.0.0.50",
	FirstBootTargetPort  string `json:"firstBootTargetPort,omitempty"`  // "firstBootTargetPort": "8080",
	InitiatorGateway     string `json:"initiatorGateway,omitempty"`     // "initiatorGateway": "3260",
	InitiatorIp          string `json:"initiatorIp,omitempty"`          // "initiatorIp": "192.168.6.21",
	InitiatorName        string `json:"initiatorName,omitempty"`        // "initiatorName": "iqn.2015-02.com.hpe:oneview-vcgs02t012",
	InitiatorNameSource  string `json:"initiatorNameSource,omitempty"`  // "initiatorNameSource": "UserDefined"
	InitiatorSubnetMask  string `json:"initiatorSubnetMask,omitempty"`  // "initiatorSubnetMask": "255.255.240.0",
	InitiatorVlanId      int    `json:"initiatorVlanId,omitempty"`      // "initiatorVlanId": 77,
	MutualChapName       string `json:"mutualChapName,omitempty"`       // "mutualChapName": "name of mutual chap",
	MutualChapSecret     string `json:"mutualChapSecret,omitempty"`     // "mutualChapSecret": "secret of mutual chap",
	SecondBootTargetIp   string `json:"secondBootTargetIp,omitempty"`   // "secondBootTargetIp": "10.0.0.51",
	SecondBootTargetPort string `json:"secondBootTargetPort,omitempty"` // "secondBootTargetPort": "78"
}
