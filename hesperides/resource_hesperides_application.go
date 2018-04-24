package hesperides

import (
	"bytes"
	"encoding/json"
	"github.com/hashicorp/terraform/helper/schema"
	"net/http"
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
			},
		},
	}
}

func resourceHesperidesApplicationCreate(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(*Config)
	n := d.Get("name").(string)
	application := HesperidesApplication{Name: n}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(application)
	req, _ := http.NewRequest(http.MethodPost, provider.Endpoint+"/rest/api/applications", body)
	client := &http.Client{}
	client.Do(req)
	return resourceHesperidesApplicationRead(d, meta)
}

func resourceHesperidesApplicationRead(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(*Config)
	n := d.Get("name").(string)
	application := HesperidesApplication{Name: n}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(application)
	req, _ := http.NewRequest(http.MethodGet, provider.Endpoint+"/rest/api/applications", body)
	client := &http.Client{}
	client.Do(req)
	return nil

}

func resourceHesperidesApplicationUpdate(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(*Config)
	n := d.Get("name").(string)
	application := HesperidesApplication{Name: n}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(application)
	req, _ := http.NewRequest(http.MethodPut, provider.Endpoint+"/rest/api/applications", body)
	client := &http.Client{}
	client.Do(req)
	return resourceHesperidesApplicationRead(d, meta)
}

func resourceHesperidesApplicationDelete(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(*Config)
	n := d.Get("name").(string)
	application := HesperidesApplication{Name: n}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(application)
	req, _ := http.NewRequest(http.MethodDelete, provider.Endpoint+"/rest/api/applications", body)
	client := &http.Client{}
	client.Do(req)
	return nil
}

type HesperidesApplication struct {
	Name string
}
