package hesperides

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"

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
	provider := meta.(*Config)

	name := d.Get("name").(string)
	version := d.Get("version").(string)
	workingCopy := d.Get("working_copy").(bool)

	techno := hesperidesTechno{Name: name, Version: version, WorkingCopy: workingCopy}
	technoJson, _ := json.Marshal(techno)

	log.Printf("[INFO] Creating Hesperides Techno: %s", technoJson)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, _ := http.NewRequest(http.MethodPost, provider.Endpoint+"/rest/templates/packages/"+name+"/"+version+"/workingcopy/templates", bytes.NewBuffer(technoJson))
	req.Header.Add("Authorization", "Basic "+provider.Token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var workingCopyStr string
	if workingCopy {
		workingCopyStr = WorkingCopy
	} else {
		workingCopyStr = Release
	}

	d.SetId(name + "-" + version + "-" + workingCopyStr)

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
	Name        string `json:"name"`
	Version     string `json:"version"`
	WorkingCopy bool   `json:"working_copy"`
}
