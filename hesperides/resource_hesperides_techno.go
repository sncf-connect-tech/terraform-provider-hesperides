package hesperides

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceHesperidesTechno() *schema.Resource {
	return &schema.Resource{
		Create: resourceHesperidesTechnoCreate,
		Read:   resourceHesperidesTechnoRead,
		Update: resourceHesperidesTechnoUpdate,
		Delete: resourceHesperidesTechnoDelete,

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

func resourceHesperidesTechnoCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceHesperidesTechnoRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceHesperidesTechnoUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceHesperidesTechnoDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

type hesperidesTechno struct {
	Name        string             `json:"name"`
	Version     string             `json:"version"`
	WorkingCopy bool               `json:"is_working_copy"`
	Template    hesperidesTemplate `json:"template"`
}
