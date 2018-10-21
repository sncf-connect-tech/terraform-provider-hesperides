package hesperides

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func dataHesperidesPlatform() *schema.Resource {
	return &schema.Resource{
		Read: dataHesperidesPlatformRead,

		Schema: map[string]*schema.Schema{
			"application": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"version": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataHesperidesPlatformRead(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(*Config)

	application := d.Get("application").(string)
	name := d.Get("name").(string)

	log.Printf("[DEBUG] Reading Hesperides Platform: %s", name)

	platformRead(*provider, application, name)

	return nil
}
