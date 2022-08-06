package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var providerFactories = map[string]func() (*schema.Provider, error){
	"jumpcloud": func() (*schema.Provider, error) {
		return New("test")(), nil
	},
}

func TestProvider(t *testing.T) {
	if err := New("dev")().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

//func testAccPreCheck(t *testing.T) {
//	// You can add code here to run prior to any test case execution, for example assertions
//	// about the appropriate environment variables being set are common to see in a pre-check
//	// function.
//}

//func preCheck(t *testing.T) {
//	// Validate that required variables are provided in env
//	variables := []string{
//		"api_key",
//		"org_id",
//	}
//
//	for _, required_variable := range variables {
//		val := os.Getenv(required_variable)
//		if val == "" {
//			t.Fatalf("Unable to test missing environment varialbe: %s", required_variable)
//		}
//	}
//}

func importStep(name string, ignore ...string) resource.TestStep {
	step := resource.TestStep{
		ResourceName:      name,
		ImportState:       true,
		ImportStateVerify: true,
	}

	if len(ignore) > 0 {
		step.ImportStateVerifyIgnore = ignore
	}

	return step
}
