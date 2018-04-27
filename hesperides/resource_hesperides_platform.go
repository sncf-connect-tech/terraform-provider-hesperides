package hesperides

import (
	"github.com/hashicorp/terraform/helper/schema"
	"bytes"
	"encoding/json"
	"net/http"
	"crypto/tls"
	"fmt"
)

func resourceHesperidesPlatform() *schema.Resource {
	return &schema.Resource{
		Create: resourceHesperidesPlatformCreate,
		Read:   resourceHesperidesPlatformRead,
		Update: resourceHesperidesPlatformUpdate,
		Delete: resourceHesperidesPlatformDelete,

		Schema: map[string]*schema.Schema{
			"application": {
				Type:     schema.TypeString,
				Optional: false,
				Required: true,
				Computed: false,
			},
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
			"production": {
				Type:     schema.TypeBool,
				Optional: false,
				Required: true,
				Computed: false,
			},
		},
	}
}

func resourceHesperidesPlatformCreate(d *schema.ResourceData, meta interface{}) error {
	fmt.Println("hello")
	provider := meta.(*Config)
	application := d.Get("application").(string)
	name := d.Get("name").(string)
	version := d.Get("version").(string)
	production := d.Get("production").(bool)
	var modules []string
	platform := HesperidesPlatform{application_name: application, platform_name: name, application_version: version, production: production, version_id: 0, modules: modules}
	fmt.Println(platform)
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(platform)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/rest/applications/"+application+"/platforms", body)
	req.Header.Add("Authorization", "Basic "+provider.Token)
	client := &http.Client{Transport: tr}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
	return nil
}

func resourceHesperidesPlatformRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceHesperidesPlatformUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceHesperidesPlatformDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

type HesperidesPlatform struct {
	application_name    string
	platform_name       string
	application_version string
	production          bool
	version_id          int
	modules             []string
}
