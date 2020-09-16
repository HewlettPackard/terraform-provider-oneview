package ov

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Configuration struct {
	UserName   string `json:"username"`
	Password   string `json:"password"`
	Endpoint   string `json:"endpoint"`
	Domain     string `json:"domain"`
	ApiVersion int    `json:"apiversion"`
	SslVerify  bool   `json:"sslverify"`
	IfMatch    string `json:"ifmatch"`
}

func LoadConfigFile(configFile string) (Configuration, error) {
	_, filename, _, _ := runtime.Caller(1)
	configFilePath := filepath.Join(filepath.Dir(filename), configFile)
	configF, err := os.Open(configFilePath)
	var config Configuration
	defer configF.Close()
	if err != nil {
		fmt.Println(err)
		return config, err
	}
	jsonParser := json.NewDecoder(configF)
	jsonParser.Decode(&config)

	return config, nil
}
