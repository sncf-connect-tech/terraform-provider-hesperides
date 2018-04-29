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
	return nil
}

func resourceHesperidesTemplateRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceHesperidesTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceHesperidesTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

type hesperidesTemplate struct {
	Name      string           `json:"name"`
	Filename  string           `json:"filename"`
	Location  string           `json:"location"`
	Rights    hesperidesRights `json:"rights"`
	VersionId int              `json:"version_id"`
}

type hesperidesRights struct {
	User  hesperidesFileRights `json:"user"`
	Group hesperidesFileRights `json:"group"`
	Other hesperidesFileRights `json:"other"`
}

type hesperidesFileRights struct {
	Read    bool `json:"read"`
	Write   bool `json:"write"`
	Execute bool `json:"execute"`
}
