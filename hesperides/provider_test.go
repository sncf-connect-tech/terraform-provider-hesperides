package hesperides

import (
	"log"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"hesperides": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	endpoint := testAccGetEndpoint()
	log.Printf("[INFO] Test: Using %s as test endpoint", endpoint)
	os.Setenv("HESPERIDES_ENDPOINT", endpoint)

	token := testAccGetToken()
	log.Printf("[INFO] Test: Using %s as test token", token)
	os.Setenv("HESPERIDES_TOKEN", token)

	err := testAccProvider.Configure(terraform.NewResourceConfig(nil))
	if err != nil {
		t.Fatal(err)
	}
}

func testAccGetEndpoint() string {
	v := os.Getenv("HESPERIDES_ENDPOINT")
	if v == "" {
		return "localhost:8080"
	}
	return v
}

func testAccGetToken() string {
	v := os.Getenv("HESPERIDES_TOKEN")
	if v == "" {
		return "localhost:8080"
	}
	return v
}
