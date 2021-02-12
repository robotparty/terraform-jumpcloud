package provider

import (
	"context"
	jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceJumpCloudLDAPDirectory() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get information about the JumpCloud LDAP directory. Each account has a LDAP directory by default and there can only be one LDAP directory.",
		ReadContext: dataSourceJumpCloudLDAPDirectoryRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Description: "Name of the LDAP directory. Example: `JumpCloud LDAP`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description: "The directory type. This will always be `ldap_server`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceJumpCloudLDAPDirectoryRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	filterFunction := func(dir jcapiv2.Directory) bool {
		return dir.Type_ == "ldap_server"
	}

	directory, err := filterJumpCloudDirectories(meta, filterFunction)
	if err != nil {
		return diag.Errorf("could not find directory with type 'ldap_server'. Previous error message: %v", err)
	}

	d.SetId(directory.Id)
	_ = d.Set("name", directory.Name)
	_ = d.Set("type", directory.Type_)

	// indicates that everything went well
	return nil
}
