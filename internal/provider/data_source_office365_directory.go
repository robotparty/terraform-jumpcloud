package provider

import (
	"context"
	jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceJumpCloudOffice365Directory() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get information about a JumpCloud Office 365 directory.",
		ReadContext: dataSourceJumpCloudOffice365DirectoryRead,
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
				Description: "The directory type. This will always be `office_365`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceJumpCloudOffice365DirectoryRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	filterFunction := func(dir jcapiv2.Directory) bool {
		return dir.Type_ == "office_365" && dir.Name == d.Get("name")
	}

	directory, err := filterJumpCloudDirectories(meta, filterFunction)
	if err != nil {
		return diag.Errorf("could not find directory with type 'office_365'. Previous error message: %v", err)
	}

	d.SetId(directory.Id)
	_ = d.Set("name", directory.Name)
	_ = d.Set("type", directory.Type_)

	// indicates that everything went well
	return nil
}
