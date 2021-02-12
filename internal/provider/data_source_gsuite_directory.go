package provider

import (
	"context"
	jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceJumpCloudGSuiteDirectory() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get information about a JumpCloud G Suite directory.",
		ReadContext: dataSourceJumpCloudGSuiteDirectoryRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Description: "The user defined name, e.g. `My G Suite directory`.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"type": {
				Description: "The directory type. This will always be `g_suite`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceJumpCloudGSuiteDirectoryRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	filterFunction := func(dir jcapiv2.Directory) bool {
		return dir.Type_ == "g_suite" && dir.Name == d.Get("name")
	}

	directory, err := filterJumpCloudDirectories(meta, filterFunction)
	if err != nil {
		return diag.Errorf("could not find directory with type 'g_suite'. Previous error message: %v", err)
	}

	d.SetId(directory.Id)
	_ = d.Set("name", directory.Name)
	_ = d.Set("type", directory.Type_)

	// indicates that everything went well
	return nil
}
