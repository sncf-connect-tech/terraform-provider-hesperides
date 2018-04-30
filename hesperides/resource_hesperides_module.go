package hesperides

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceHesperidesModule() *schema.Resource {
	return &schema.Resource{
		Create: resourceHesperidesModuleCreate,
		Read:   resourceHesperidesModuleRead,
		Update: resourceHesperidesModuleUpdate,
		Delete: resourceHesperidesModuleDelete,

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
			"version_id": {
				Type:     schema.TypeInt,
				Optional: false,
				Required: true,
				Computed: false,
			},
		},
	}
}

func resourceHesperidesModuleCreate(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(*Config)

	name := d.Get("name").(string)
	version := d.Get("version").(string)
	workingCopy := d.Get("working_copy").(bool)
	versionId := d.Get("version_id").(int)

	module := hesperidesModule{Name: name, Version: version, WorkingCopy: workingCopy, Technos: []string{}, VersionId: versionId}
	moduleJson, _ := json.Marshal(module)

	log.Printf("[INFO] Creating Hesperides Module: %s", moduleJson)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, _ := http.NewRequest(http.MethodPost, provider.Endpoint+"/rest/modules", bytes.NewBuffer(moduleJson))
	req.Header.Add("Authorization", "Basic "+provider.Token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var workingCopyStr string
	if workingCopy {
		workingCopyStr = "workingcopy"
	} else {
		workingCopyStr = "release"
	}

	d.SetId(name + "-" + version + "-" + workingCopyStr)

	return nil
}

func resourceHesperidesModuleRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceHesperidesModuleUpdate(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(*Config)

	name := d.Get("name").(string)
	version := d.Get("version").(string)
	workingCopy := d.Get("working_copy").(bool)
	versionId := d.Get("version_id").(int)

	module := hesperidesModule{Name: name, Version: version, WorkingCopy: workingCopy, Technos: []string{}, VersionId: versionId}
	moduleJson, _ := json.Marshal(module)

	log.Printf("[INFO] Updating Hesperides Module: %s", moduleJson)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, _ := http.NewRequest(http.MethodPut, provider.Endpoint+"/rest/modules", bytes.NewBuffer(moduleJson))
	req.Header.Add("Authorization", "Basic "+provider.Token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return resourceHesperidesApplicationRead(d, meta)
}

func resourceHesperidesModuleDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

type hesperidesModule struct {
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	WorkingCopy bool     `json:"working_copy"`
	Technos     []string `json:"technos"`
	VersionId   int      `json:"version_id"`
}
