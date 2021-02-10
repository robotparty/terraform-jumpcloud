package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccUserGroupAssociation(t *testing.T) {

	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)
	resource.Test(t, resource.TestCase{
		PreCheck:          func() {},
		ProviderFactories: providerFactories,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
		resource "jumpcloud_application" "test_application_%s" {
						display_label = "test_aws_account"
						sso_url = "https://sso.jumpcloud.com/saml2/example-application"
						saml_role_attribute = "arn:aws:iam::AWS_ACCOUNT_ID:role/MY_ROLE,arn:aws:iam::AWS_ACCOUNT_ID:saml-provider/MY_SAML_PROVIDER"
						aws_session_duration = 432000
		}

		resource "jumpcloud_user_group" "test_group_%s" {
			name = "testgroup_%s"
		}

		resource "jumpcloud_user_group_association" "test_association_%s" {
 			object_id = "${jumpcloud_application.test_application_%s.id}"
			group_id = "${jumpcloud_user_group.test_group_%s.id}"
			type = "application"
 		}
	`, rName, rName, rName, rName, rName, rName),
				Check: resource.TestCheckResourceAttrSet(fmt.Sprintf(`jumpcloud_user_group_association.test_association_%s`, rName),
					"group_id"),
			},
		},
	})
}
