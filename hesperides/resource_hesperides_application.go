package hesperides

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceHesperidesApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceHesperidesApplicationCreate,
		Read:   resourceHesperidesApplicationRead,
		Update: resourceHesperidesApplicationUpdate,
		Delete: resourceHesperidesApplicationDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: false,
				Required: true,
				Computed: false,
			},
		},
	}
}

func resourceHesperidesApplicationCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceHesperidesApplicationRead(d *schema.ResourceData, meta interface{}) error {
	return nil

}

func resourceHesperidesApplicationUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceHesperidesApplicationDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

type HesperidesApplication struct {
	Name string
}
