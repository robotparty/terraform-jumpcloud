package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func Test_resourceApplication(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			// Create step
			{
				Config: fmt.Sprintf(`resource "jumpcloud_application" "test_application" {
						display_label = "test_aws_account"
						sso_url = "https://sso.jumpcloud.com/saml2/example-application"
						saml_role_attribute = "arn:aws:iam::AWS_ACCOUNT_ID:role/MY_ROLE,arn:aws:iam::AWS_ACCOUNT_ID:saml-provider/MY_SAML_PROVIDER"
						aws_session_duration = 432000
					}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("jumpcloud_application.test_application", "display_label", "test_aws_account"),
				),
			},
			userImportStep("jumpcloud_application.test_application"),
			// Update Step
			{
				Config: fmt.Sprintf(`resource "jumpcloud_application" "test_application" {
						display_label = "test_aws_account2"
						sso_url = "https://sso.jumpcloud.com/saml2/example-application"
						saml_role_attribute = "arn:aws:iam::AWS_ACCOUNT_ID:role/MY_ROLE,arn:aws:iam::AWS_ACCOUNT_ID:saml-provider/MY_SAML_PROVIDER"
						aws_session_duration = 432000
					}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("jumpcloud_application.test_application", "display_label", "test_aws_account2"),
				),
			},
			userImportStep("jumpcloud_application.test_application"),
		},
	})
}
