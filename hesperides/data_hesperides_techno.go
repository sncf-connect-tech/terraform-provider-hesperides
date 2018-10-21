package hesperides

import (
	"encoding/json"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func dataHesperidesTechno() *schema.Resource {
	return &schema.Resource{
		Read: dataHesperidesTechnoRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: false,
				Required: true,
				Computed: false,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: false,
				Required: true,
				Computed: false,
			},
			"working_copy": {
				Type:     schema.TypeBool,
				Optional: false,
				Required: true,
				Computed: false,
			},
		},
	}
}

func dataHesperidesTechnoRead(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(*Config)

	name := d.Get("name").(string)
	version := d.Get("version").(string)
	workingCopy := d.Get("working_copy").(bool)

	techno := hesperidesTechno{Name: name, Version: version, WorkingCopy: workingCopy}
	technoJson, _ := json.Marshal(techno)

	log.Printf("[DEBUG] Reading Hesperides Techno: %s", technoJson)

	if workingCopy {
		technoReadTemplates(*provider, name, version, WorkingCopy)
	} else {
		technoReadTemplates(*provider, name, version, Release)
	}

	return resourceHesperidesTechnoRead(d, meta)
}
