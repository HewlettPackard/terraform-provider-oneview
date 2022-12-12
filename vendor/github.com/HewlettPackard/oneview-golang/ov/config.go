package ov

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/HewlettPackard/oneview-golang/utils"
)

type Configuration struct {
	OVCred                      *OVCred
	IdPoolsIpv4Subnet           *IdPoolsIpv4Subnet           `json:"id_pools_ipv4_subnet,omitempty"`
	IdPoolsIpv4SubnetRange      *IdPoolsIpv4SubnetRange      `json:"id_pools_ipv4_range,omitempty"`
	ServerProfileConfig         *ServerProfileConfig         `json:"server_profile,omitempty"`
	ServerProfileTemplateConfig *ServerProfileTemplateConfig `json:"server_profile_template,omitempty"`
	HypervisorManagerConfig     *HypervisorManagerConfig     `json:"hypervisor_manager,omitempty"`
	StorageSystemConfig         *StorageSystemConfig         `json:"storage_system,omitempty"`
	LigName                     string                       `json:"ligName"`
	EgName                      string                       `json:"egName"`
	MgmtNetworkName             string                       `json:"mgmtNetworkName"`
	IscsiNetworkName            string                       `json:"iscsiNetworkName"`
	FcNetworkName               string                       `json:"fcNetworkName"`
	ServerCertificateIp         string                       `json:"serverCertificateIp"`
}
type OVCred struct {
	UserName   string `json:"username"`
	Password   string `json:"password"`
	Endpoint   string `json:"endpoint"`
	Domain     string `json:"domain"`
	ApiVersion int    `json:"apiversion,string"`
	SslVerify  bool   `json:"sslverify"`
	IfMatch    string `json:"ifmatch"`
}

type IdPoolsIpv4Subnet struct {
	Count        string        `json:"count"`         //: "3",
	Domain       string        `json:"domain"`        //: "example.com",
	NetworkId    string        `json:"network_id"`    //: "10.1.0.0",
	EndAddress   utils.Nstring `json:"end_address"`   //: "10.1.19.61",
	NewDomain    string        `json:"newDomain"`     //: "awesome.com",
	StartAddress utils.Nstring `json:"start_address"` //: "10.1.19.56",
	SubnetName   string        `json:"subnet_name"`   //: "Test IPv4 Subnet",
	SubnetMask   string        `json:"subnetmask"`    //: "255.255.192.0",
	Gateway      string        `json:"gateway"`       //: "10.1.0.1",
	Type         string        `json:"type"`          //: "Subnet",
	DnsServers   []string      `json:"dnsServers"`    //: []
}

type IdPoolsIpv4SubnetRange struct {
	Count        string          `json:"count"`         //: "3",
	Domain       string          `json:"domain"`        //: "example.com",
	NetworkId    string          `json:"networkId"`     //: "10.1.0.0",
	EndAddress1  utils.Nstring   `json:"end_address1"`  //: "10.1.19.61",
	EndAddress2  utils.Nstring   `json:"end_address2"`  //: "10.1.19.61",
	NewDomain    string          `json:"newDomain"`     //: "awesome.com",
	StartAddress utils.Nstring   `json:"start_address"` //: "10.1.19.56",
	SubnetName   string          `json:"subnet_name"`   //: "Test IPv4 Subnet",
	SubnetMask   string          `json:"subnetmask"`    //: "255.255.192.0",
	Gateway      string          `json:"gateway"`       //: "10.1.0.1",
	Type         string          `json:"type"`          //: "Subnet",
	DnsServers   []string        `json:"dnsServers"`    //: []
	IdList       []utils.Nstring `json:"idList"`
}
type ServerProfileTemplateConfig struct {
	ConnectionNetworkName      string `json:"connection_network_name"`
	ServerPrpofileTemplateName string `json:"server_profile_template_name"`
	ServerHardwareTypeName     string `json:"server_hardware_type_name"`
	EnclosureGroupName         string `json:"enclosure_group_name"`
}
type HypervisorManagerConfig struct {
	Username  string `json:"hypervisor_manager_username"`
	Password  string `json:"hypervisor_manager_password"`
	IpAddress string `json:"hypervisor_manager_ip"`
}
type StorageSystemConfig struct {
	Username   string `json:"storage_username"`
	Password   string `json:"storage_password"`
	IpAddress  string `json:"storage_IP"`
	IpAddress2 string `json:"storage_IP2"`
	Family     string `json:"storage_family"`
}

type ServerProfileConfig struct {
	InventoryHostName        string `json:"connection_network_name"`
	ServerProfileDescription string `json:"server_profile_template_name"`
	ServerHardwareName       string `json:"server_hardware_name"`
	EnclosureGroupName       string `json:"enclosure_group_name"`
	ServerProfileName        string `json:"server_profile_name"`
	NetworkNamestring        string `json:"network_name"`
	OvTemplatestring         string `json:"ov_template"`
}

func LoadConfigFile(configFile string) (Configuration, error) {
	_, filename, _, _ := runtime.Caller(1)
	configFilePath := filepath.Join(filepath.Dir(filename), configFile)
	configF, err := os.Open(configFilePath)
	var config Configuration
	defer configF.Close()
	if err != nil {
		fmt.Println(err)
		fmt.Println("error opening json file")
		return config, err
	}
	jsonParser := json.NewDecoder(configF)

	err_unmarshal := jsonParser.Decode(&config)
	if err_unmarshal != nil {
		fmt.Println("error unmarshaling json file")
		fmt.Println(err_unmarshal)

	}

	return config, nil
}
