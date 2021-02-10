package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func Test_resourceUser(t *testing.T) {

	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		//PreCheck: func() {
		//	preCheck(t)
		//},
		PreCheck: func() {
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			// Create step
			{
				Config: fmt.Sprintf(`resource "jumpcloud_user" "test_%s" {
						username = "test_user_%s"
						email = "test_%s@sagewave.io"
						firstname = "sage"
						lastname = "wave"
						enable_mfa = false
					}`, rName, rName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("jumpcloud_user.test_%s", rName), "username", fmt.Sprintf("test_user_%s", rName)),
					resource.TestCheckResourceAttr(fmt.Sprintf("jumpcloud_user.test_%s", rName), "email", fmt.Sprintf("test_%s@sagewave.io", rName)),
					resource.TestCheckResourceAttr(fmt.Sprintf("jumpcloud_user.test_%s", rName), "firstname", "sage"),
					resource.TestCheckResourceAttr(fmt.Sprintf("jumpcloud_user.test_%s", rName), "lastname", "wave"),
					resource.TestCheckResourceAttr(fmt.Sprintf("jumpcloud_user.test_%s", rName), "enable_mfa", "false"),
				),
			},
			userImportStep(fmt.Sprintf("jumpcloud_user.test_%s", rName)),

			// Update Step
			{
				Config: fmt.Sprintf(`resource "jumpcloud_user" "test_%s" {
						username = "test_user_%s"
						email = "test_%s@sagewave.io"
						firstname = "updatedSage"
						lastname = "wave"
						enable_mfa = false
					}`, rName, rName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("jumpcloud_user.test_%s", rName), "username", fmt.Sprintf("test_user_%s", rName)),
					resource.TestCheckResourceAttr(fmt.Sprintf("jumpcloud_user.test_%s", rName), "email", fmt.Sprintf("test_%s@sagewave.io", rName)),
					resource.TestCheckResourceAttr(fmt.Sprintf("jumpcloud_user.test_%s", rName), "firstname", "updatedSage"),
					resource.TestCheckResourceAttr(fmt.Sprintf("jumpcloud_user.test_%s", rName), "lastname", "wave"),
					resource.TestCheckResourceAttr(fmt.Sprintf("jumpcloud_user.test_%s", rName), "enable_mfa", "false"),
				),
			},
			userImportStep(fmt.Sprintf("jumpcloud_user.test_%s", rName)),
		},
	})
}

func userImportStep(name string) resource.TestStep {
	return importStep(name, "allow_existing", "skip_forget_on_destroy")
}
