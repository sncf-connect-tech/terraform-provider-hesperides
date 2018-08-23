package hesperides

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceHesperidesTechno() *schema.Resource {
	return &schema.Resource{
		Create: resourceHesperidesTechnoCreate,
		Read:   resourceHesperidesTechnoRead,
		Update: resourceHesperidesTechnoUpdate,
		Delete: resourceHesperidesTechnoDelete,
		Importer: &schema.ResourceImporter{
			State: resourceHesperidesTechnoImportState,
		},

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
			"templates": {
				Type:     schema.TypeList,
				Optional: false,
				Required: true,
				Computed: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: false,
							Required: true,
							Computed: false,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: false,
							Required: true,
							Computed: false,
						},
						"filename": {
							Type:     schema.TypeString,
							Optional: false,
							Required: true,
							Computed: false,
						},
						"location": {
							Type:     schema.TypeString,
							Optional: false,
							Required: true,
							Computed: false,
						},
						"content": {
							Type:     schema.TypeString,
							Optional: false,
							Required: true,
							Computed: false,
						},
						"rights": {
							Type:     schema.TypeList,
							Optional: true,
							Required: false,
							Computed: false,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"user": {
										Type:     schema.TypeList,
										Optional: false,
										Required: true,
										Computed: false,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"read": {
													Type:     schema.TypeBool,
													Optional: false,
													Required: true,
													Computed: false,
												},
												"write": {
													Type:     schema.TypeBool,
													Optional: false,
													Required: true,
													Computed: false,
												},
												"execute": {
													Type:     schema.TypeBool,
													Optional: false,
													Required: true,
													Computed: false,
												},
											},
										},
									},
									"group": {
										Type:     schema.TypeList,
										Optional: false,
										Required: true,
										Computed: false,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"read": {
													Type:     schema.TypeBool,
													Optional: false,
													Required: true,
													Computed: false,
												},
												"write": {
													Type:     schema.TypeBool,
													Optional: false,
													Required: true,
													Computed: false,
												},
												"execute": {
													Type:     schema.TypeBool,
													Optional: false,
													Required: true,
													Computed: false,
												},
											},
										},
									},
									"other": {
										Type:     schema.TypeList,
										Optional: false,
										Required: true,
										Computed: false,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"read": {
													Type:     schema.TypeBool,
													Optional: false,
													Required: true,
													Computed: false,
												},
												"write": {
													Type:     schema.TypeBool,
													Optional: false,
													Required: true,
													Computed: false,
												},
												"execute": {
													Type:     schema.TypeBool,
													Optional: false,
													Required: true,
													Computed: false,
												},
											},
										},
									},
								},
							},
						},
						"version_id": {
							Type:     schema.TypeInt,
							Optional: false,
							Required: true,
							Computed: false,
						},
					},
				},
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

	var workingCopyStr string
	if workingCopy {
		workingCopyStr = WorkingCopy
	} else {
		workingCopyStr = Release
	}

	if len(d.Get("templates").([]interface{})) == 0 {
		return fmt.Errorf("no template found")
	}

	if _, ok := d.GetOk("templates"); ok {
		for index, raw := range d.Get("templates").([]interface{}) {
			templateRaw := raw.(map[string]interface{})
			templateName := templateRaw["name"].(string)
			templateNamespace := templateRaw["namespace"].(string)
			templateFilename := templateRaw["filename"].(string)
			templateLocation := templateRaw["location"].(string)
			templateContent := templateRaw["content"].(string)
			templateVersionId := templateRaw["version_id"].(int)

			var rights hesperidesTemplateRights
			if _, ok := d.GetOk("templates." + strconv.Itoa(index) + ".rights"); ok {
				if _, ok := d.GetOk("templates." + strconv.Itoa(index) + ".rights.0.user"); ok {
					var user hesperidesTemplateFileRights
					if v, ok := d.GetOk("templates." + strconv.Itoa(index) + ".rights.0.user.0.read"); ok {
						user.Read = v.(bool)
					}
					if v, ok := d.GetOk("templates." + strconv.Itoa(index) + ".rights.0.user.0.write"); ok {
						user.Write = v.(bool)
					}
					if v, ok := d.GetOk("templates." + strconv.Itoa(index) + ".rights.0.user.0.execute"); ok {
						user.Execute = v.(bool)
					}
					rights.User = user
				}
				if _, ok := d.GetOk("templates." + strconv.Itoa(index) + ".rights.0.group"); ok {
					var group hesperidesTemplateFileRights
					if v, ok := d.GetOk("templates." + strconv.Itoa(index) + ".rights.0.group.0.read"); ok {
						group.Read = v.(bool)
					}
					if v, ok := d.GetOk("templates." + strconv.Itoa(index) + ".rights.0.group.0.write"); ok {
						group.Write = v.(bool)
					}
					if v, ok := d.GetOk("templates." + strconv.Itoa(index) + ".rights.0.group.0.execute"); ok {
						group.Execute = v.(bool)
					}
					rights.Group = group
				}
				if _, ok := d.GetOk("templates." + strconv.Itoa(index) + ".rights.0.other"); ok {
					var other hesperidesTemplateFileRights
					if v, ok := d.GetOk("templates." + strconv.Itoa(index) + ".rights.0.other.0.read"); ok {
						other.Read = v.(bool)
					}
					if v, ok := d.GetOk("templates." + strconv.Itoa(index) + ".rights.0.other.0.write"); ok {
						other.Write = v.(bool)
					}
					if v, ok := d.GetOk("templates." + strconv.Itoa(index) + ".rights.0.other.0.execute"); ok {
						other.Execute = v.(bool)
					}
					rights.Other = other
				}
			}

			template := hesperidesTemplate{Name: templateName, Namespace: templateNamespace, Filename: templateFilename, Location: templateLocation, Content: templateContent, Rights: rights, VersionId: templateVersionId}
			templateJson, _ := json.Marshal(template)

			log.Printf("[INFO] Adding Hesperides Template: %s", templateJson)

			http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

			var method string
			if index == 0 {
				method = http.MethodPost
			} else {
				method = http.MethodPut
			}

			req, _ := http.NewRequest(method, provider.Endpoint+"/rest/templates/packages/"+name+"/"+version+"/workingcopy/templates", bytes.NewBuffer(templateJson))
			req.Header.Add("Authorization", "Basic "+provider.Token)
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			_, err := client.Do(req)
			if err != nil {
				panic(err)
			}
		}
	}

	// Release the techno if it was created directly as it
	if !workingCopy {
		log.Printf("[INFO] Releasing Hesperides Techno: %s", technoJson)

		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

		req, _ := http.NewRequest(http.MethodPost, provider.Endpoint+"/rest/templates/packages/create_release?techno_name="+name+"&techno_version="+version, nil)
		req.Header.Add("Authorization", "Basic "+provider.Token)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		_, err := client.Do(req)
		if err != nil {
			panic(err)
		}
	}

	d.SetId(name + "-" + version + "-" + workingCopyStr)

	return nil
}

func resourceHesperidesTechnoRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceHesperidesTechnoUpdate(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(*Config)

	name := d.Get("name").(string)
	version := d.Get("version").(string)
	workingCopy := d.Get("working_copy").(bool)

	techno := hesperidesTechno{Name: name, Version: version, WorkingCopy: workingCopy}
	technoJson, _ := json.Marshal(techno)

	// Illegal change: use the same resource for a different techno
	if d.HasChange("name") {
		return fmt.Errorf("illegal change: \"name\" can not be changed, consider creating a new techno")
	}

	// Illegal change: the techno pass from a release to a working copy using the same version
	if !d.HasChange("version") && d.HasChange("working_copy") && workingCopy {
		return fmt.Errorf("illegal change: could not pass from a release to a working copy using the same version")
	}

	// Hesperides consider a change in the version as a new techno
	if d.HasChange("version") {
		return resourceHesperidesTechnoCreate(d, meta)
	}

	// Release the techno
	if d.HasChange("working_copy") && !workingCopy {
		log.Printf("[INFO] Releasing Hesperides Techno: %s", technoJson)

		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

		req, _ := http.NewRequest(http.MethodPost, provider.Endpoint+"/rest/templates/packages/create_release?techno_name="+name+"&techno_version="+version, nil)
		req.Header.Add("Authorization", "Basic "+provider.Token)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		_, err := client.Do(req)
		if err != nil {
			panic(err)
		}
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

func resourceHesperidesTechnoDelete(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(*Config)

	name := d.Get("name").(string)
	version := d.Get("version").(string)
	workingCopy := d.Get("working_copy").(bool)

	techno := hesperidesTechno{Name: name, Version: version, WorkingCopy: workingCopy}
	technoJson, _ := json.Marshal(techno)

	log.Printf("[INFO] Deleting Hesperides Techno: %s", technoJson)

	var workingCopyStr string
	if workingCopy {
		workingCopyStr = WorkingCopy
	} else {
		workingCopyStr = Release
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, _ := http.NewRequest(http.MethodDelete, provider.Endpoint+"/rest/templates/packages/"+name+"/"+version+"/"+workingCopyStr+"/templates", nil)
	req.Header.Add("Authorization", "Basic "+provider.Token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return resourceHesperidesApplicationRead(d, meta)
}

func resourceHesperidesTechnoImportState(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	provider := meta.(*Config)

	name, version, workingCopyStr := parseThreePartID(d.Id())

	var workingCopy = workingCopyStr == WorkingCopy

	techno := hesperidesTechno{Name: name, Version: version, WorkingCopy: workingCopy}
	technoJson, _ := json.Marshal(techno)

	log.Printf("[INFO] Importing Hesperides Techno: %s", technoJson)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, _ := http.NewRequest(http.MethodGet, provider.Endpoint+"/rest/templates/packages/"+name+"/"+version+"/"+workingCopyStr, nil)
	req.Header.Add("Authorization", "Basic "+provider.Token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	d.SetId(name + "-" + version + "-" + workingCopyStr)

	return []*schema.ResourceData{d}, nil
}
