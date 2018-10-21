package hesperides

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func dataHesperidesModule() *schema.Resource {
	return &schema.Resource{
		Read: dataHesperidesModuleRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: false,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: false,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"working_copy": {
				Type:     schema.TypeBool,
				Optional: false,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func dataHesperidesModuleRead(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(*Config)

	name := d.Get("name").(string)
	version := d.Get("version").(string)
	workingCopy := d.Get("working_copy").(bool)

	log.Printf("[DEBUG] Reading Hesperides Module: %s", name)

	if workingCopy {
		moduleRead(*provider, name, version, WorkingCopy)
	} else {
		moduleRead(*provider, name, version, Release)
	}

	return nil
}
