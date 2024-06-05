package viettelidc

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

//https://spacelift.io/blog/terraform-data-sources-how-they-are-utilised

func resourceKubernetesV1() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceKubernetesCreate,
		ReadContext:   resourceKubernetesRead,
		UpdateContext: resourceKubernetesUpdate,
		DeleteContext: resourceKubernetesDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"account": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: false,
			},
		},
	}
}

func resourceKubernetesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	//config := meta.(*Config)
	// if config == nil {
	// 	fmt.Println("Config is nil")
	// } else {
	// 	userName := config.Username
	// 	region := config.Region
	// 	fmt.Println("UserName: ", userName)
	// 	fmt.Println("Region: ", region)
	// }

	name := d.Get("name").(string)
	account := d.Get("account").(string)

	d.SetId(account)
	fmt.Println("Name: ", name)

	resp, err := http.Get("https://www.uuidtools.com/api/generate/v1/count/10")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	return resourceKubernetesRead(ctx, d, meta)
}

func resourceKubernetesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceKubernetesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceKubernetesRead(ctx, d, meta)
}

func resourceKubernetesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
