package hesperides

import (
	"bytes"
	"encoding/json"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
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

	module := hesperidesModule{Name: name, Version: version, WorkingCopy: workingCopy, Technos: []hesperidesTechno{}, VersionId: versionId}
	moduleJson, _ := json.Marshal(module)

	log.Printf("[INFO] Creating Hesperides Module: %s", moduleJson)

	var workingCopyStr string
	if workingCopy {
		workingCopyStr = WorkingCopy
	} else {
		workingCopyStr = Release
	}

	moduleCreate(*provider, bytes.NewBuffer(moduleJson))

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

	module := hesperidesModule{Name: name, Version: version, WorkingCopy: workingCopy, Technos: []hesperidesTechno{}, VersionId: versionId}
	moduleJson, _ := json.Marshal(module)

	log.Printf("[INFO] Updating Hesperides Module: %s", moduleJson)

	moduleUpdate(*provider, bytes.NewBuffer(moduleJson))

	return resourceHesperidesApplicationRead(d, meta)
}

func resourceHesperidesModuleDelete(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(*Config)

	name := d.Get("name").(string)
	version := d.Get("version").(string)
	workingCopy := d.Get("working_copy").(bool)
	versionId := d.Get("version_id").(int)

	module := hesperidesModule{Name: name, Version: version, WorkingCopy: workingCopy, Technos: []hesperidesTechno{}, VersionId: versionId}
	moduleJson, _ := json.Marshal(module)

	log.Printf("[INFO] Deleting Hesperides Module: %s", moduleJson)

	if workingCopy {
		moduleDelete(*provider, name, version, WorkingCopy)
	} else {
		moduleDelete(*provider, name, version, Release)
	}

	return resourceHesperidesApplicationRead(d, meta)
}
