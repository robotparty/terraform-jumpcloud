package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

// TODO attributes needs to be fixed
func Test_resourceUserGroup(t *testing.T) {
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)

	resource.ParallelTest(t, resource.TestCase{
		//PreCheck: func() {
		//	preCheck(t)
		//},
		PreCheck: func() {
		},
		ProviderFactories: providerFactories,
		//CheckDestroy:
		Steps: []resource.TestStep{
			// Create step
			{
				Config: fmt.Sprintf(`resource "jumpcloud_user_group" "test_%s" {
name = "test_group_%s"
}`, rName, rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("jumpcloud_user_group.test_%s", rName), "name", fmt.Sprintf("test_group_%s", rName)),
				),
			},
			userImportStep(fmt.Sprintf("jumpcloud_user_group.test_%s", rName)),
		},
	})
}
