package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccUserGroupAssociation(t *testing.T) {
	randSuffix := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() {},
		ProviderFactories: providerFactories,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testUserGroupAssocConfig(randSuffix),
				Check: resource.TestCheckResourceAttrSet("jumpcloud_user_group_association.test_association",
					"group_id"),
			},
		},
	})
}

func testUserGroupAssocConfig(randSuffix string) string {
	return fmt.Sprintf(`
resource "jumpcloud_application" "test_application" {
	display_label        = "test_aws_account_%s"
	sso_url              = "https://sso.jumpcloud.com/saml2/example-application-%s"
	saml_role_attribute  = "arn:aws:iam::AWS_ACCOUNT_ID:role/MY_ROLE,arn:aws:iam::AWS_ACCOUNT_ID:saml-provider/MY_SAML_PROVIDER"
	aws_session_duration = 432000
}

resource "jumpcloud_user_group" "test_group" {
	name = "testgroup_%s"
}

resource "jumpcloud_user_group_association" "test_association" {
	object_id = jumpcloud_application.test_application.id
	group_id  = jumpcloud_user_group.test_group.id
	type      = "application"
}
`, randSuffix, randSuffix, randSuffix)
}
