package oneview

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ethernetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"purpose": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "General",
			},
			"private_network": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"smart_link": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ethernet_network_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Tagged",
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ethernet-networkV3",
			},
			"connection_template_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fabric_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scopesuri": {
				Optional: true,
				Type:     schema.TypeString,
				Computed: true,
			},
			"initial_scope_uris": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
		},
	}
}
