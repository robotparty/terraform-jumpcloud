package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TODO attributes needs to be fixed
func Test_resourceUserGroup(t *testing.T) {
	randSuffix := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)

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
				Config: fmt.Sprintf(`
resource "jumpcloud_user_group" "test" {
	name = "test_group_%s"
}
`, randSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("jumpcloud_user_group.test", "name", fmt.Sprintf("test_group_%s", randSuffix)),
				),
			},
			userImportStep("jumpcloud_user_group.test"),
		},
	})
}
