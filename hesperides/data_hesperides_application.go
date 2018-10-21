package hesperides

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func dataHesperidesApplication() *schema.Resource {
	return &schema.Resource{
		Read: dataHesperidesApplicationRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataHesperidesApplicationRead(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(*Config)

	name := d.Get("name").(string)

	log.Printf("[DEBUG] Reading Hesperides Application: %s", name)

	applicationRead(*provider, name)

	return nil
}
