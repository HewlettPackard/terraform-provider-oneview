package ov

import (
	"encoding/json"
	"fmt"
	"os"
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
	path, errPath := os.Getwd()
	if errPath != nil {
		panic(errPath)
	}
	configFilePath := path + "/examples/" + configFile
	configF, err1 := os.Open(configFilePath)
	var config Configuration
	defer configF.Close()
	if err1 != nil {
		fmt.Println(err1)
		return config, err1
	}

	jsonParser := json.NewDecoder(configF)
	jsonParser.Decode(&config)
	fmt.Println(config.Password)

	return config, nil
}
