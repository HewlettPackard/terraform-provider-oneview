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

// Package ov -
package ov

import (
	"github.com/HewlettPackard/oneview-golang/utils"
)

// LogicalDrive logical drive options
type LogicalDrive struct {
	Bootable  bool   `json:"bootable"`            // "bootable": true,
	RaidLevel string `json:"raidLevel,omitempty"` // "raidLevel": "RAID1"
}

// LocalStorageOptions -
type LocalStorageOptions struct { // "localStorage": {
	Controllers        []LocalStorageEmbeddedController `json:"controllers,omitempty"`        //  The list of embedded local storage controllers.
	Initialize         bool                             `json:"initialize,omitempty"`         // 				"initialize": true
	LogicalDrives      []LogicalDrive                   `json:"logicalDrives,omitempty"`      // "logicalDrives": [],
	ManageLocalStorage bool                             `json:"manageLocalStorage,omitempty"` // "manageLocalStorage": true,
	ReapplyState       string                           `json:"reapplyState,omitempty"`       //Current reapply state of SAN storage component.
	SasLogicalJBODs    []LogicalJbod                    `json:"sasLogicalJBODs,omitempty"`    // "sasLogicalJBODs": [],
}

// StoragePath storage path host-to-target paths
//Use GET /rest/storage-systems/{arrayid}/managedPorts?query="expectedNetworkUri EQ '/rest/fc-networks/{netowrk-id}'"
//to retrieve the storage targets for the associated network.
type StoragePath struct {
	ConnectionID   int           `json:"connectionId,omitempty"` // connectionId (required), The ID of the connection associated with this storage path. Use GET /rest/server-profiles/available-networks to retrieve the list of available networks.
	IsEnabled      bool          `json:"isEnabled"`              // isEnabled (required), Identifies whether the storage path is enabled.
	NetworkUri     utils.Nstring `json:"networkUri"`             //The URI of the network associated with this storage path
	Status         string        `json:"status,omitempty"`       // status (read only), The overall health status of the storage path.
	Targets        []Target      `json:"targets,omitempty"`      // only set when storageTargetType is TargetPorts
	TargetSelector string        `json:"targetSelector,omitempty"`
}

type Target struct {
	IpAddress string `json:"IpAddress,omitempty"`
	Name      string `json:"name,omitempty"`
	TcpPort   int    `json:"tcpPort,omitempty"`
}

type PropertiesSP struct {
	DataProtectionLevel           string        `json:"dataProtectionLevel,omitempty"`           //The data protection level of the volume.
	DataTransferLimit             int           `json:"dataTransferLimit,omitempty"`             //The data transfer limit of the volume in mebibytes.
	Description                   string        `json:"description,omitempty"`                   //The description of the volume.
	Folder                        string        `json:"folder,omitempty"`                        //The folder of the volume.
	IsAdaptiveOptimizationEnabled bool          `json:"isAdaptiveOptimizationEnabled,omitempty"` //Indicates whether or not the volume will use adaptive optimization.
	IsCompressed                  bool          `json:"isCompressed,omitempty"`                  //Indicates whether or not compression is enabled on the volume.
	IsDataReductionEnabled        bool          `json:"isDataReductionEnabled,omitempty"`        //Indicates whether or not data reduction, such as deduplication and compression, is enabled on the volume
	IsDeduplicated                bool          `json:"isDeduplicated,omitempty"`                //Indicates whether or not deduplication is enabled on the volume.
	IsEncrypted                   bool          `json:"isEncrypted,omitempty"`                   //Indicates whether or not encryption is enabled on the volume.
	IopsLimit                     int           `json:"iopsLimit,omitempty"`                     //The IOPS (input/output operations per second) limit of the volume.
	IsPinned                      bool          `json:"isPinned,omitempty"`                      //Indicates whether or not the volume is cache pinned.
	IsShareable                   *bool         `json:"isShareable,omitempty"`                   //Indicates whether the volume is shareable or private
	Name                          string        `json:"name,omitempty"`                          //The name of the volume.
	PerformancePolicy             string        `json:"performancePolicy,omitempty"`             //The performance policy of the volume.
	ProvisioningType              string        `json:"provisioningType,omitempty"`              //The provisioning type of the volume.
	Size                          int           `json:"size,omitempty"`                          //The size of the volume.
	SnapshotPool                  utils.Nstring `json:"snapshotPool,omitempty"`                  //A URI reference to the common provisioning group used to create snapshots
	StoragePool                   utils.Nstring `json:"storagePool,omitempty"`                   //The storage pool of the volume.
	TemplateVersion               string        `json:"templateVersion,omitempty"`               //The template version of the volume.
	VolumeSet                     utils.Nstring `json:"volumeSet,omitempty"`                     //URI of a volume set to use for the volum

}

type Volume struct {
	InitialScopeUris utils.Nstring `json:"initialScopeUris,omitempty"` //Initial scopes for the volume.
	IsPermanent      *bool         `json:"isPermanent,omitempty"`      //If true, indicates that the volume will persist when the profile using this volume is deleted.
	Properties       *PropertiesSP `json:"properties"`                 //The properties specific to a storage system family required for the creation of a storage volume.
	TemplateUri      utils.Nstring `json:"templateUri,omitempty"`      //URI of the storage volume template from which the volume will be created.

}

// VolumeAttachment volume attachment
type VolumeAttachment struct {
	AssociatedTemplateAttachmentId string        `json:"associatedTemplateAttachmentId,omitempty"` //A "key" value uniquely identifying the definition of a volume attachment in a template
	BootVolumePriority             string        `json:"bootVolumePriority,omitempty"`             //Identifies whether the volume will be used as a boot volume and with what priority
	ID                             int           `json:"id,omitempty"`                             // id, The ID of the attached storage volume.
	LUN                            string        `json:"lun,omitempty"`                            // lun, The logical unit number.
	LUNType                        string        `json:"lunType,omitempty"`                        // lunType(required), The logical unit number type: Auto or Manual.
	Permanent                      *bool         `json:"permanent,omitempty"`                      // permanent, If true, indicates that the volume will persist when the profile is deleted. If false, then the volume will be deleted when the profile is deleted.
	State                          string        `json:"state,omitempty"`                          // state(read only), current state of the attachment
	Status                         string        `json:"status,omitempty"`                         // status(read only), The current status of the attachment.
	StoragePaths                   []StoragePath `json:"storagePaths,omitempty"`                   // A list of host-to-target path associations.
	Volume                         *Volume       `json:"volume,omitempty"`                         //Contains properties describing a volume to be create
	VolumeStorageSystemURI         utils.Nstring `json:"volumeStorageSystemUri,omitempty"`         // The URI of the storage system associated with this volume attachment. Use GET /rest/server-profiles/available-storage-systems to retrieve the URI of the storage system associated with a volume.
	VolumeURI                      utils.Nstring `json:"volumeUri,omitempty"`                      // The URI of the storage volume associated with this volume attachment. Use GET /rest/server-profiles/available-storage-systems to retrieve the URIs of available storage volumes.
}

// Clone clone volume attachment for submits
type SanSystemCredential struct {
	ChapLevel        string        `json:"chapLevel"`        //The Challenge Handshake Authentication Protocol (CHAP) authentication method
	ChapName         string        `json:"chapName"`         //The Challenge Handshake Authentication Protocol (CHAP) name.
	ChapSecret       string        `json:"chapSecret"`       //The CHAP secret or password. This secret is automatically generated.
	ChapSource       string        `json:"chapSource"`       //The CHAP source is AutoGenerated or UserDefine
	MutualChapName   string        `json:"mutualChapName"`   //The Mutual Challenge Handshake Authentication Protocol
	MutualChapSecret string        `json:"mutualChapSecret"` //The Mutual-CHAP secret.
	StorageSystemUri utils.Nstring `json:"storageSystemUri"` //The iSCSI storage system uri for which CHAP authentication is used
}

// SanStorageOptions specify san storage
// No San
// 		"sanStorage": {
// 				"volumeAttachments": [],
// 				"manageSanStorage": false
// 		},
type SanStorageOptions struct { // sanStorage
	HostOSType                 string                `json:"hostOSType,omitempty"`                 // hostOSType(required),  The operating system type of the host. To retrieve the list of supported host OS types, issue a REST Get request using the /rest/storage-systems/host-types API.
	ManageSanStorage           bool                  `json:"manageSanStorage"`                     // manageSanStorage(required),  Identifies whether SAN storage is managed in this profile.
	ReapplyState               string                `json:"reapplyState,omitempty"`               //Current reapply state of SAN storage component.
	SanSystemCredentials       []SanSystemCredential `json:"sanSystemCredentials,omitempty"`       //array of SanSystemCredentialsV2
	SecretsGenerated           string                `json:"secretsGenerated,omitempty"`           // The date and time the SAN Storage CHAP secrets were generated
	VolumeAttachments          []VolumeAttachment    `json:"volumeAttachments,omitempty"`          // volumeAttachments, The list of storage volume attachments. array of Volume Attachment
	ScopeUri                   utils.Nstring         `json:"scopeUri,omitempty"`                   //The URI for the resource scope assignments.
	SerialNumber               string                `json:"serialNumber,omitempty"`               // serialNumber (searchable) A 10-byte value that is exposed to the Operating System as the server hardware's Serial Number. The value can be a virtual serial number, user defined serial number or physical serial number read from the server's ROM. It cannot be modified after the profile is created.
	SerialNumberType           string                `json:"serialNumberType,omitempty"`           // serialNumberType (searchable) Specifies the type of Serial Number and UUID to be programmed into the server ROM. The value can be 'Virtual', 'UserDefined', or 'Physical'. The serialNumberType defaults to 'Virtual' when serialNumber or uuid are not specified. It cannot be modified after the profile is created.
	ServerHardwareReapplyState string                `json:"serverHardwareReapplyState,omitempty"` //Current reapply state of the server that is associated with this server profile
	ServerHardwareTypeURI      utils.Nstring         `json:"serverHardwareTypeUri,omitempty"`      // serverHardwareTypeUri Identifies the server hardware type for which the Server Profile was designed. The serverHardwareTypeUri is determined when the profile is created and cannot be modified. Use GET /rest/server-hardware-types to retrieve the list of server hardware types.
	ServerHardwareURI          utils.Nstring         `json:"serverHardwareUri,omitempty"`          // serverHardwareUri Identifies the server hardware to which the server profile is currently assigned, if applicable. Use GET /rest/server-profiles/available-targets to retrieve the list of available servers.
	ServerProfileTemplateUri   utils.Nstring         `json:"serverProfileTemplateUri,omitempty"`   //Identifies the Server profile template the Server Profile is based on
	ServiceManager             string                `json:"serviceManager,omitempty"`             //Name of a service manager that is designated owner of the profile
	State                      string                `json:"state,omitempty"`                      // state (searchable, readonly) Current State of this Server Profile
	Status                     string                `json:"status,omitempty"`                     // status (searchable, readonly) Overall health status of this Server Profile
	TaskURI                    utils.Nstring         `json:"taskUri,omitempty"`                    // taskUri (read only) URI of the task currently executing or most recently executed on this server profile.
	TemplateCompliance         string                `json:"templateCompliance,omitempty"`         //The compliance state of the server profile with the server profile template.
	Type                       string                `json:"type,omitempty"`                       // type (read only) Identifies the resource type. This field should always be set to 'ServerProfileV4'.
	URI                        utils.Nstring         `json:"uri,omitempty"`                        // uri (searchable, readonly) URI of this Server Profile. The URI is automatically generated when the server profile is created and cannot be modified.
	UUID                       string                `json:"uuid,omitempty"`                       // uuid (searchable) A 36-byte value that is exposed to the Operating System as the server hardware's UUID. The value can be a virtual uuid, user defined uuid or physical uuid read from the server's ROM. It cannot be modified after the profile is created.
	WWNType                    string                `json:"wwnType,omitempty"`                    // wwnType (searchable) Specifies the type of WWN address to be programmed into the IO devices. The value can be 'Virtual' or 'Physical'. It cannot be modified after the profile is created.
}
