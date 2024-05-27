package viettelidc

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/http"
)

//https://spacelift.io/blog/terraform-data-sources-how-they-are-utilised

func resourceServer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceServerCreate,
		ReadContext:   resourceServerRead,
		UpdateContext: resourceServerUpdate,
		DeleteContext: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"uuid_count": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	uuid_count := d.Get("uuid_count").(string)

	d.SetId(uuid_count)

	// https://www.uuidtools.com/api/generate/v1/count/uuid_count
	//body := strings.NewReader(uuid_count)
	resp, err := http.Get("https://www.uuidtools.com/api/generate/v1/count/10")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return resourceServerRead(ctx, d, meta)
}

func resourceServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceServerRead(ctx, d, meta)
}

func resourceServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
