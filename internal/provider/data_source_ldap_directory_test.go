package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func Test_dataSourceJumpCloudLdapDirectory(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() {},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`data "jumpcloud_ldap_directory" "test" {
				}`),

				Check: resource.ComposeTestCheckFunc(),
			},
		},
	})
}
