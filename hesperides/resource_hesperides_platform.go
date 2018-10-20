package hesperides

import (
	"bytes"
	"encoding/json"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
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
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"production": {
				Type:     schema.TypeBool,
				Required: true,
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

	platformCreate(*provider, application, bytes.NewBuffer(platformJson))

	d.SetId(buildTwoPartID(&application, &name))

	return resourceHesperidesPlatformRead(d, meta)
}

func resourceHesperidesPlatformRead(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(*Config)

	application := d.Get("application").(string)
	name := d.Get("name").(string)

	log.Printf("[DEBUG] Reading Hesperides Platform: %s", name)

	platformRead(*provider, application, name)

	return nil
}

func resourceHesperidesPlatformUpdate(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(*Config)

	applicationName, platformName := parseTwoPartID(d.Id())
	version := d.Get("version").(string)
	production := d.Get("production").(bool)

	platform := hesperidesPlatform{ApplicationName: applicationName, PlatformName: platformName, ApplicationVersion: version, Production: production, VersionId: 1, Modules: []string{}}
	platformJson, _ := json.Marshal(platform)

	log.Printf("[INFO] Updating Hesperides Platform: %s", platformJson)

	platformUpdate(*provider, applicationName, platformName, bytes.NewBuffer(platformJson))

	return resourceHesperidesPlatformRead(d, meta)
}

func resourceHesperidesPlatformDelete(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(*Config)

	applicationName, platformName := parseTwoPartID(d.Id())

	log.Printf("[INFO] Deleting Hesperides Platform: %s", d.Id())

	platformDelete(*provider, applicationName, platformName)

	return nil
}
