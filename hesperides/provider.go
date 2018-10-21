package hesperides

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["endpoint"],
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["token"],
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"hesperides_application": dataHesperidesApplication(),
			"hesperides_module":      dataHesperidesModule(),
			"hesperides_platform":    dataHesperidesPlatform(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"hesperides_application": resourceHesperidesApplication(),
			"hesperides_module":      resourceHesperidesModule(),
			"hesperides_platform":    resourceHesperidesPlatform(),
			"hesperides_techno":      resourceHesperidesTechno(),
			"hesperides_template":    resourceHesperidesTemplate(),
		},

		ConfigureFunc: configureProvider,
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"endpoint": "The endpoint where Hesperides operations will take place",

		"token": "session token. A session token is only required if you are\n" +
			"using temporary security credentials.",
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Endpoint: d.Get("endpoint").(string),
		Token:    d.Get("token").(string),
	}

	return &config, nil
}
