// (C) Copyright 2016 Hewlett Packard Enterprise Development LP
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
	"errors"
	"fmt"
	"github.com/HewlettPackard/oneview-golang/i3s"
	"github.com/HewlettPackard/oneview-golang/ov"
)

type Config struct {
	OVDomain     string
	OVUsername   string
	OVPassword   string
	OVEndpoint   string
	OVSSLVerify  bool
	OVAPIVersion int
	OVIfMatch    string

	I3SEndpoint string

	ovClient  *ov.OVClient
	i3sClient *i3s.I3SClient
}

var ErrConfigNotInitialized = errors.New("config not initialized!")

func (c *Config) loadAndValidate() error {
	if c == nil {
		return ErrConfigNotInitialized
	}

	client := (&ov.OVClient{}).NewOVClient(c.OVUsername, c.OVPassword, c.OVDomain, c.OVEndpoint, c.OVSSLVerify, c.OVAPIVersion, c.OVIfMatch)
	c.ovClient = client
	apiver, err := c.ovClient.GetAPIVersion()

	//If no api version is provided use the current version to create client
	if c.OVAPIVersion == 0 {
		if err != nil {
			return fmt.Errorf("Could not fetch the appliance %s api version", c.OVEndpoint)
		}
		c.OVAPIVersion = apiver.CurrentVersion
	}
	//Throw error if provided api version is not supported
	if c.OVAPIVersion < apiver.MinimumVersion {
		return fmt.Errorf("The minimum api version supported is %d", apiver.MinimumVersion)
	}
	client = (&ov.OVClient{}).NewOVClient(c.OVUsername, c.OVPassword, c.OVDomain, c.OVEndpoint, c.OVSSLVerify, c.OVAPIVersion, c.OVIfMatch)
	session, err := c.ovClient.SessionLogin()
	if err != nil {
		return err
	}

	c.ovClient.APIKey = session.ID

	return nil
}

func (c *Config) loadAndValidateI3S() error {
	if c == nil {
		return ErrConfigNotInitialized
	}

	client := (&i3s.I3SClient{}).NewI3SClient(c.I3SEndpoint, c.OVSSLVerify, c.OVAPIVersion, c.ovClient.APIKey)

	c.i3sClient = client

	return nil
}
