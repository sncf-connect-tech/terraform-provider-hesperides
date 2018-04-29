package hesperides

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
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
				ForceNew: true,
			},
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
	provider := meta.(*Config)

	application := d.Get("application").(string)
	name := d.Get("name").(string)
	version := d.Get("version").(string)
	production := d.Get("production").(bool)

	platform := hesperidesPlatform{ApplicationName: application, PlatformName: name, ApplicationVersion: version, Production: production, VersionId: 0, Modules: []string{}}
	platformJson, _ := json.Marshal(platform)

	log.Printf("[INFO] Creating Hesperides Platform: %s", platformJson)

	req, _ := http.NewRequest(http.MethodPost, provider.Endpoint+"/rest/applications/"+application+"/platforms", bytes.NewBuffer(platformJson))
	req.Header.Add("Authorization", "Basic "+provider.Token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	d.SetId(application + "-" + name)

	return nil
}

func resourceHesperidesPlatformRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceHesperidesPlatformUpdate(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(*Config)

	application := d.Get("application").(string)
	name := d.Get("name").(string)
	version := d.Get("version").(string)
	production := d.Get("production").(bool)

	platform := hesperidesPlatform{ApplicationName: application, PlatformName: name, ApplicationVersion: version, Production: production, VersionId: 0, Modules: []string{}}
	platformJson, _ := json.Marshal(platform)

	log.Printf("[INFO] Updating Hesperides Platform: %s", platformJson)

	req, _ := http.NewRequest(http.MethodPut, provider.Endpoint+"/rest/applications/"+application+"/platforms", bytes.NewBuffer(platformJson))
	req.Header.Add("Authorization", "Basic "+provider.Token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return nil
}

func resourceHesperidesPlatformDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

type hesperidesPlatform struct {
	ApplicationName    string   `json:"application_name"`
	PlatformName       string   `json:"platform_name"`
	ApplicationVersion string   `json:"application_version"`
	Production         bool     `json:"production"`
	VersionId          int      `json:"version_id"`
	Modules            []string `json:"modules"`
}
