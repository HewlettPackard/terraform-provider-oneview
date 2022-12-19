// (C) Copyright 2021 Hewlett Packard Enterprise Development LP
//
// Licensed under the Apache License, Version 2.0 (the "License");
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed
// under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package oneview

import (
	"github.com/HewlettPackard/terraform-provider-oneview/internal/mutexkv"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	ovMutexKV          = mutexkv.NewMutexKV()
	serverHardwareURIs = make(map[string]bool)
)

// Provider returns a schema.Provider for oneview
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"ov_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ONEVIEW_OV_DOMAIN", ""),
			},
			"ov_username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ONEVIEW_OV_USER", ""),
			},
			"ov_password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ONEVIEW_OV_PASSWORD", nil),
			},
			"ov_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ONEVIEW_OV_ENDPOINT", nil),
			},
			"ov_sslverify": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ONEVIEW_OV_SSLVERIFY", true),
			},
			"ov_apiversion": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ONEVIEW_OV_API_VERSION", 0),
			},
			"ov_ifmatch": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ONEVIEW_OV_IF_MATCH", "*"),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"oneview_appliance_snmp_v1_trap_destinations": dataSourceSNMPv1TrapDestination(),
			"oneview_appliance_snmpv3_trap_destinations":  dataSourceSNMPv3TrapDestination(),
			"oneview_appliance_snmpv3_user":               dataSourceSNMPv3User(),
			"oneview_appliance_ssh_access":                dataSourceSSHAccess(),
			"oneview_appliance_time_and_locale":           dataSourceTimeAndLocale(),
			"oneview_connection_templates":                dataSourceConnectionTemplates(),
			"oneview_enclosure":                           dataSourceEnclosure(),
			"oneview_enclosure_group":                     dataSourceEnclosureGroup(),
			"oneview_ethernet_network":                    dataSourceEthernetNetwork(),
			"oneview_fc_network":                          dataSourceFCNetwork(),
			"oneview_fcoe_network":                        dataSourceFCoENetwork(),
			"oneview_firmware_drivers":                    dataSourceFirmwareDrivers(),
			"oneview_hypervisor_cluster_profile":          dataSourceHypervisorClusterProfile(),
			"oneview_hypervisor_manager":                  dataSourceHypervisorManager(),
			"oneview_id_pools_ipv4_subnets":               dataSourceIPv4Subnets(),
			"oneview_ipv4_range":                          dataSourceIPv4Ranges(),
			"oneview_interconnect_type":                   dataSourceInterconnectType(),
			"oneview_interconnect":                        dataSourceInterconnects(),
			"oneview_label":                               dataSourceLabel(),
			"oneview_logical_enclosure":                   dataSourceLogicalEnclosure(),
			"oneview_logical_interconnect":                dataSourceLogicalInterconnect(),
			"oneview_logical_interconnect_group":          dataSourceLogicalInterconnectGroup(),
			"oneview_network_set":                         dataSourceNetworkSet(),
			"oneview_rack_manager":                        dataSourceRackManager(),
			"oneview_scope":                               dataSourceScope(),
			"oneview_server_certificate":                  dataSourceServerCertificate(),
			"oneview_server_hardware":                     dataSourceServerHardware(),
			"oneview_server_hardware_type":                dataSourceServerHardwareType(),
			"oneview_storage_attachment":                  dataSourceStorageAttachment(),
			"oneview_server_profile":                      dataSourceServerProfile(),
			"oneview_server_profile_template":             dataSourceServerProfileTemplate(),
			"oneview_storage_pool":                        dataSourceStoragePool(),
			"oneview_storage_system":                      dataSourceStorageSystem(),
			"oneview_storage_volume_template":             dataSourceStorageVolumeTemplate(),
			"oneview_task":                                dataSourceTask(),
			"oneview_uplink_set":                          dataSourceUplinkSet(),
			"oneview_volume":                              dataSourceVolume(),
			"oneview_version":                             dataSourceVersion(),
			"oneview_relative_value":                      dataSourceRelativeValue(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"oneview_appliance_snmp_v1_trap_destinations": resourceSNMPv1TrapDestination(),
			"oneview_appliance_snmpv3_trap_destinations":  resourceSNMPv3TrapDestination(),
			"oneview_appliance_snmpv3_user":               resourceSNMPv3User(),
			"oneview_appliance_ssh_access":                resourceSSHAccess(),
			"oneview_appliance_time_and_locale":           resourceTimeAndLocale(),
			"oneview_connection_templates":                resourceConnectionTemplates(),
			"oneview_enclosure":                           resourceEnclosure(),
			"oneview_enclosure_group":                     resourceEnclosureGroup(),
			"oneview_ethernet_network":                    resourceEthernetNetwork(),
			"oneview_fcoe_network":                        resourceFCoENetwork(),
			"oneview_fc_network":                          resourceFCNetwork(),
			"oneview_firmware_drivers":                    resourceFirmwareDrivers(),
			"oneview_hypervisor_cluster_profile":          resourceHypervisorClusterProfile(),
			"oneview_hypervisor_manager":                  resourceHypervisorManager(),
			"oneview_id_pools_ipv4_subnets":               resourceIPv4Subnets(),
			"oneview_ipv4_range":                          resourceIPv4Ranges(),
			"oneview_label":                               resourceLabel(),
			"oneview_logical_enclosure":                   resourceLogicalEnclosure(),
			"oneview_logical_interconnect_group":          resourceLogicalInterconnectGroup(),
			"oneview_logical_interconnect":                resourceLogicalInterconnect(),
			"oneview_logical_switch_group":                resourceLogicalSwitchGroup(),
			"oneview_network_set":                         resourceNetworkSet(),
			"oneview_rack_manager":                        resourceRackManager(),
			"oneview_scope":                               resourceScope(),
			"oneview_server_certificate":                  resourceServerCertificate(),
			"oneview_server_hardware":                     resourceServerHardware(),
			"oneview_server_profile":                      resourceServerProfile(),
			"oneview_server_profile_template":             resourceServerProfileTemplate(),
			"oneview_storage_system":                      resourceStorageSystem(),
			"oneview_storage_pool":                        resourceStoragePool(),
			"oneview_storage_volume_template":             resourceStorageVolumeTemplate(),
			"oneview_task":                                resourceTask(),
			"oneview_uplink_set":                          resourceUplinkSet(),
			"oneview_volume":                              resourceVolume(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		OVDomain:     d.Get("ov_domain").(string),
		OVUsername:   d.Get("ov_username").(string),
		OVPassword:   d.Get("ov_password").(string),
		OVEndpoint:   d.Get("ov_endpoint").(string),
		OVSSLVerify:  d.Get("ov_sslverify").(bool),
		OVAPIVersion: d.Get("ov_apiversion").(int),
		OVIfMatch:    d.Get("ov_ifmatch").(string),
	}

	if err := config.loadAndValidate(); err != nil {
		return nil, err
	}
	return &config, nil
}
