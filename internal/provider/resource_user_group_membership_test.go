package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccUserGroupMembership(t *testing.T) {
	randomSuffix := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() {},
		ProviderFactories: providerFactories,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testUserGroupMembershipConfig(randomSuffix),
				Check: resource.TestCheckResourceAttrSet("jumpcloud_user_group_membership.test_membership",
					"user_id"),
			},
		},
	})
}

func testUserGroupMembershipConfig(randSuffix string) string {
	return fmt.Sprintf(`
resource "jumpcloud_user" "test_user" {
	username   = "test_user_%s"
	email      = "test_%s@sagewave.io"
	firstname  = "sage"
	lastname   = "wave"
	enable_mfa = false
}

resource "jumpcloud_user_group" "test_group" {
	name = "testgroup_%s"
}

resource "jumpcloud_user_group_membership" "test_membership" {
	user_id  = jumpcloud_user.test_user.id
	group_id = jumpcloud_user_group.test_group.id
}
`, randSuffix, randSuffix, randSuffix)
}
