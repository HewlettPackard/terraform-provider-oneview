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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIPv4Ranges() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIPv4RangesRead,

		Schema: map[string]*schema.Schema{
			"allocated_fragment_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allocated_id_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"allocator_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_resources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"association_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_category": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"collector_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_range": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"end_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"free_fragment_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"free_id_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"prefix": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"range_category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reserved_id_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"start_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"start_stop_fragments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fragment_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"start_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"subnet_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIPv4RangesRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	id := d.Get("id").(string)
	ipv4range, err := config.ovClient.GetIPv4RangebyId("", id)
	if err != nil {
		d.SetId("")
		return err
	} else if ipv4range.URI.IsNil() {
		d.SetId("")
		return nil
	}
	d.Set("name", ipv4range.Name)
	d.Set("type", ipv4range.Type)
	d.Set("created", ipv4range.Created)
	d.Set("modified", ipv4range.Modified)
	d.Set("uri", ipv4range.URI.String())
	d.Set("category", ipv4range.Category)
	d.Set("etag", ipv4range.ETAG)
	d.Set("subnet_uri", ipv4range.SubnetUri)
	d.Set("total_count", ipv4range.TotalCount)
	d.Set("reserved_id_count", ipv4range.ReservedIdCount)
	d.Set("start_address", ipv4range.StartAddress)
	d.Set("range_category", ipv4range.RangeCategory)
	d.Set("prefix", ipv4range.Prefix)
	d.Set("free_fragment_uri", ipv4range.FreeFragmentUri)
	d.Set("free_id_count", ipv4range.FreeIdCount)
	d.Set("end_address", ipv4range.EndAddress)
	d.Set("enabled", ipv4range.Enabled)
	d.Set("default_range", ipv4range.DefaultRange)
	d.Set("allocated_fragment_uri", ipv4range.AllocatedFragmentUri)
	d.Set("collector_uri", ipv4range.CollectorUri)
	d.Set("allocated_id_count", ipv4range.AllocatedIdCount)
	d.Set("allocator_uri", ipv4range.AllocatorUri)

	startStopFragments := make([]map[string]interface{}, 0, len(ipv4range.StartStopFragments))

	for _, startStopFragment := range ipv4range.StartStopFragments {
		startStopFragments = append(startStopFragments, map[string]interface{}{
			"end_address":   startStopFragment.EndAddress,
			"fragment_type": startStopFragment.FragmentType,
			"start_address": startStopFragment.StartAddress,
		})
	}

	d.Set("start_stop_fragments", startStopFragments)

	AssociatedResources := make([]map[string]interface{}, 0, len(ipv4range.AssociatedResources))

	for _, AssociatedResource := range ipv4range.AssociatedResources {
		AssociatedResources = append(AssociatedResources, map[string]interface{}{
			"association_type":  AssociatedResource.AssociationType,
			"resource_category": AssociatedResource.ResourceCategory,
			"resource_name":     AssociatedResource.ResourceName,
			"resource_uri":      AssociatedResource.ResourceUri,
		})
	}

	d.Set("associated_resources", AssociatedResources)

	d.SetId(id)
	return nil
}
