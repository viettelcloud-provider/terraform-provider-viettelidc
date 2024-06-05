package viettelidc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKubernetesV1() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKubernetesV1Read,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
			},

			"host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceKubernetesV1Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	dataSourceKubernetesV1Attributes(d)

	return nil
}

func dataSourceKubernetesV1Attributes(d *schema.ResourceData) {
	d.SetId("server1")
	d.Set("name", "server1")
	d.Set("status", "PEDNING")
	d.Set("host", "localhost")

}
