package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestUserAcc_resourceUser(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		//PreCheck: func() {
		//	preCheck(t)
		//},
		PreCheck: func() {
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`resource "jumpcloud_user" "test" {
	username = "test_user"
	email = "test@sagewave.io"
	firstname = "sage"
	lastname = "wave"
	enable_mfa = false
}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("jumpcloud_user.test", "username", "test_user"),
				//resource.TestCheckResourceAttr("jumpcloud_user.test", "email", "test@sagewave.io"),
				//resource.TestCheckResourceAttr("jumpcloud_user.test", "firstname", "sage"),
				//resource.TestCheckResourceAttr("jumpcloud_user.test", "lastname", "wave"),
				//resource.TestCheckResourceAttr("jumpcloud_user.test", "enable_mfa", "false"),
				),
			},
			{
				ResourceName:      "jumpcloud_user",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
