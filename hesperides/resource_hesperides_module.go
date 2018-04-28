package hesperides

import (
	"bytes"
	"encoding/json"
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

	req, _ := http.NewRequest(http.MethodPost, provider.Endpoint+"/rest/modules", bytes.NewBuffer(moduleJson))
	req.Header.Add("Authorization", "Basic "+provider.Token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return resourceHesperidesApplicationRead(d, meta)
}

func resourceHesperidesModuleRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceHesperidesModuleUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
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
