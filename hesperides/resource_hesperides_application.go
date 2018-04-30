package hesperides

import (
	"log"

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
				Required: true,
			},
		},
	}
}

func resourceHesperidesApplicationCreate(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)

	log.Printf("[INFO] Creating Hesperides Application: %s", name)

	d.SetId(name)

	return nil
}

func resourceHesperidesApplicationRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceHesperidesApplicationUpdate(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)

	log.Printf("[INFO] Updating Hesperides Application: %s", name)

	return nil
}

func resourceHesperidesApplicationDelete(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)

	log.Printf("[INFO] Deleting Hesperides Application: %s", name)

	return nil
}
