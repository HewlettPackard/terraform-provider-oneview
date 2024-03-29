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
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"oneview": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	v := os.Getenv("ONEVIEW_OV_ENDPOINT")
	if v == "" {
		t.Fatal("ONEVIEW_OV_ENDPOINT must be set for acceptance tests")
	}

	v = os.Getenv("ONEVIEW_OV_USER")
	if v == "" {
		t.Fatal("ONEVIEW_OV_USER must be set for acceptance test")
	}

	v = os.Getenv("ONEVIEW_OV_PASSWORD")
	if v == "" {
		t.Fatal("ONEVIEW_OV_PASSWORD must be set for acceptance test")
	}

	v = os.Getenv("ONEVIEW_SSLVERIFY")
	if v == "" {
		t.Fatal("ONEVIEW_OV_SSLVERIFY must be set for acceptance test")
	}

}

func testProviderConfig() (*Config, error) {
	config := testAccProvider.Meta().(*Config)
	if config == nil {
		return nil, fmt.Errorf("Unable to obtain provider config\n")
	}
	return config, nil
}
