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
	"path"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIPv4Ranges() *schema.Resource {
	return &schema.Resource{
		Read:   resourceIPv4RangesRead,
		Create: resourceIPv4RangesCreate,
		Update: resourceIPv4RangesUpdate,
		Delete: resourceIPv4RangesDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"allocator_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
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
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"association_type": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"resource_category": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"resource_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"resource_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"collector_id_list": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"collector_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": {
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
				Optional: true,
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
			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_address": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"fragment_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"start_address": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"subnet_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"total_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Range",
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIPv4RangesCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	ipv4Range := ov.CreateIpv4Range{
		Type:      d.Get("type").(string),
		Name:      d.Get("name").(string),
		SubnetUri: utils.NewNstring(d.Get("subnet_uri").(string)),
	}
	if val, ok := d.GetOk("start_stop_fragments"); ok {
		rawfragments := val.(*schema.Set).List()
		fragments := make([]ov.StartStopFragments, 0)
		for _, rawfrag := range rawfragments {
			rawitem := rawfrag.(map[string]interface{})
			startstopfragment := ov.StartStopFragments{
				StartAddress: utils.NewNstring(rawitem["start_address"].(string)),
				EndAddress:   utils.NewNstring(rawitem["end_address"].(string)),
			}
			fragments = append(fragments, startstopfragment)
		}
		ipv4Range.StartStopFragments = fragments
	}

	data, err := config.ovClient.CreateIPv4Range(ipv4Range)
	if err != nil {
		d.SetId("")
		return err
	}

	id := path.Base(data.URI.String())
	d.SetId(id)
	return resourceIPv4RangesRead(d, meta)
}

func resourceIPv4RangesRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	ipv4range, err := config.ovClient.GetIPv4RangebyId("", d.Id())
	if err != nil || ipv4range.URI.IsNil() {
		d.SetId("")
		return nil
	}
	idList := make([]map[string]interface{}, 0)
	d.Set("allocator_count", 0)
	d.Set("collector_id_list", idList)
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
	d.SetId(d.Id())
	return nil
}

func resourceIPv4RangesUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	if d.HasChange("allocator_count") {
		allocator := ov.UpdateAllocatorList{
			Count: d.Get("allocator_count").(int),
		}
		_, err := config.ovClient.AllocateId(allocator, d.Id())
		if err != nil {
			d.SetId("")
			return err
		}

	} else if d.HasChange("collector_id_list") {
		ids := d.Get("collector_id_list").(*schema.Set).List()

		idsList := make([]utils.Nstring, len(ids))
		for i, raw := range ids {
			idsList[i] = utils.Nstring(raw.(string))
		}

		collect := ov.UpdateCollectorList{
			IdList: idsList,
		}
		_, err := config.ovClient.CollectId(collect, d.Id())
		if err != nil {
			d.SetId("")
			return err
		}

	} else {

		ipv4Range := ov.Ipv4Range{
			Type: d.Get("type").(string),
			Name: d.Get("name").(string),
		}
		if val, ok := d.GetOk("start_stop_fragments"); ok {
			rawfragments := val.(*schema.Set).List()
			fragments := make([]ov.StartStopFragments, 0)
			for _, rawfrag := range rawfragments {
				rawitem := rawfrag.(map[string]interface{})
				startstopfragment := ov.StartStopFragments{
					StartAddress: utils.NewNstring(rawitem["start_address"].(string)),
					EndAddress:   utils.NewNstring(rawitem["end_address"].(string)),
				}
				fragments = append(fragments, startstopfragment)
			}
			ipv4Range.StartStopFragments = fragments
		}
		if _, ok := d.GetOk("enabled"); ok {
			ipv4Range.Enabled = d.Get("enabled").(bool)
		}

		_, err := config.ovClient.UpdateIpv4Range(d.Id(), ipv4Range)
		if err != nil {
			d.SetId("")
			return err
		}
	}
	d.SetId(d.Id())

	return resourceIPv4RangesRead(d, meta)
}

func resourceIPv4RangesDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	err := config.ovClient.DeleteIpv4Range(d.Id())
	if err != nil {
		return err
	}
	return nil
}
