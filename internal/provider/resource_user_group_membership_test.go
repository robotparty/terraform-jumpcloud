package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccUserGroupMembership(t *testing.T) {
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() {},
		ProviderFactories: providerFactories,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
		resource "jumpcloud_user" "test_user_%s" {
		username = "test_user_%s"
		email = "test@sagewave.io"
		firstname = "sage"
		lastname = "wave"
		enable_mfa = false
	}

		resource "jumpcloud_user_group" "test_group_%s" {
			name = "testgroup_%s"
		}

		resource "jumpcloud_user_group_membership" "test_membership_%s" {
  			user_id = "${jumpcloud_user.test_user_%s.id}"
			group_id = "${jumpcloud_user_group.test_group_%s.id}"
  		}
	`, rName, rName, rName, rName, rName, rName, rName),
				Check: resource.TestCheckResourceAttrSet(fmt.Sprintf("jumpcloud_user_group_membership.test_membership_%s", rName),
					"user_id"),
			},
		},
	})
}
