package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func Test_resourceUser(t *testing.T) {
	randomSuffix := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			// Create step
			{
				Config: fmt.Sprintf(`resource "jumpcloud_user" "test" {
						username   = "test_user_%s"
						email      = "test_%s@sagewave.io"
						firstname  = "sage"
						lastname   = "wave"
						enable_mfa = false
					}`, randomSuffix, randomSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("jumpcloud_user.test", "username", fmt.Sprintf("test_user_%s", randomSuffix)),
					resource.TestCheckResourceAttr("jumpcloud_user.test", "email", fmt.Sprintf("test_%s@sagewave.io", randomSuffix)),
					resource.TestCheckResourceAttr("jumpcloud_user.test", "firstname", "sage"),
					resource.TestCheckResourceAttr("jumpcloud_user.test", "lastname", "wave"),
					resource.TestCheckResourceAttr("jumpcloud_user.test", "enable_mfa", "false"),
				),
			},
			userImportStep("jumpcloud_user.test"),

			// Update Step
			{
				Config: fmt.Sprintf(`resource "jumpcloud_user" "test" {
						username = "test_user_%s"
						email = "test_%s@sagewave.io"
						firstname = "updatedSage"
						lastname = "wave"
						enable_mfa = false
					}`, randomSuffix, randomSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("jumpcloud_user.test", "username", fmt.Sprintf("test_user_%s", randomSuffix)),
					resource.TestCheckResourceAttr("jumpcloud_user.test", "email", fmt.Sprintf("test_%s@sagewave.io", randomSuffix)),
					resource.TestCheckResourceAttr("jumpcloud_user.test", "firstname", "updatedSage"),
					resource.TestCheckResourceAttr("jumpcloud_user.test", "lastname", "wave"),
					resource.TestCheckResourceAttr("jumpcloud_user.test", "enable_mfa", "false"),
				),
			},
			userImportStep("jumpcloud_user.test"),
		},
	})
}

func userImportStep(name string) resource.TestStep {
	return importStep(name, "allow_existing", "skip_forget_on_destroy")
}
