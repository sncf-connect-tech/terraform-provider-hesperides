package hesperides

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceHesperidesTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceHesperidesTemplateCreate,
		Read:   resourceHesperidesTemplateRead,
		Update: resourceHesperidesTemplateUpdate,
		Delete: resourceHesperidesTemplateDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: false,
				Required: true,
				Computed: false,
			},
			"filename": {
				Type:     schema.TypeString,
				Optional: false,
				Required: true,
				Computed: false,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: false,
				Required: true,
				Computed: false,
			},
			"version_id": {
				Type:     schema.TypeInt,
				Optional: false,
				Required: true,
				Computed: false,
			},
		},
	}
}

func resourceHesperidesTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceHesperidesTemplateRead(d, meta)
}

func resourceHesperidesTemplateRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceHesperidesTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceHesperidesTemplateRead(d, meta)
}

func resourceHesperidesTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
