package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccUserGroupMembership(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:          func() {},
		ProviderFactories: providerFactories,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
		resource "jumpcloud_user" "test_user" {
		username = "test_user"
		email = "test@sagewave.io"
		firstname = "sage"
		lastname = "wave"
		enable_mfa = false
	}

		resource "jumpcloud_user_group" "test_group" {
			name = "testgroup"
		}

		resource "jumpcloud_user_group_membership" "test_membership" {
  			user_id = "${jumpcloud_user.test_user.id}"
			group_id = "${jumpcloud_user_group.test_group.id}"
  		}
	`),
				Check: resource.TestCheckResourceAttrSet("jumpcloud_user_group_membership.test_membership",
					"user_id"),
			},
		},
	})
}
