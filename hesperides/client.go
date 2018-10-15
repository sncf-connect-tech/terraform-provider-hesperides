package hesperides

import (
	"crypto/tls"
	"io"
	"net/http"
)

func hesperidesClient(config Config, method string, url string, body io.Reader) (*http.Response) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, _ := http.NewRequest(method, config.Endpoint+url, body)
	req.Header.Add("Authorization", "Basic "+config.Token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return response
}

func technoAddTemplates(config Config, name string, version string, releaseType string, body io.Reader) {
	url := "/rest/templates/packages/" + name + "/" + version + "/" + releaseType + "/templates"
	hesperidesClient(config, http.MethodPost, url, body)
}

func technoDeleteTemplates(config Config, name string, version string, releaseType string) {
	url := "/rest/templates/packages/" + name + "/" + version + "/" + releaseType + "/templates"
	hesperidesClient(config, http.MethodDelete, url, nil)
}

func technoReadTemplates(config Config, name string, version string, releaseType string) (*http.Response) {
	url := "/rest/templates/packages/" + name + "/" + version + "/" + releaseType + "/templates"
	return hesperidesClient(config, http.MethodGet, url, nil)
}

func technoRelease(config Config, name string, version string) {
	url := "/rest/templates/packages/create_release?techno_name=" + name + "&techno_version=" + version
	hesperidesClient(config, http.MethodPost, url, nil)
}

func technoUpdateTemplates(config Config, name string, version string, releaseType string, body io.Reader) {
	url := "/rest/templates/packages/" + name + "/" + version + "/" + releaseType + "/templates"
	hesperidesClient(config, http.MethodPut, url, body)
}
