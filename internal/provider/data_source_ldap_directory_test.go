package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func Test_dataSourceJumpCloudLdapDirectory(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() {},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: dataLdapDirectoryConfig,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

const dataLdapDirectoryConfig = `
data "jumpcloud_ldap_directory" "test" {
}
`
