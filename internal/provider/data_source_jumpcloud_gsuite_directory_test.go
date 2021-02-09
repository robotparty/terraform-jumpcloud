package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func dataSourceJumpCloudGSuiteDirectoryTest(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() {},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`resource "g_suite" "test" {
	name = test
	type = g_suite
}
`),
			},
		},
	})
}
