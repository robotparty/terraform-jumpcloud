package provider

import (
	"context"
	jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceJumpCloudLDAPDirectory() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get information about the JumpCloud LDAP directory.",
		ReadContext: dataSourceJumpCloudLDAPDirectoryRead,
		Schema: map[string]*schema.Schema{
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

func dataSourceJumpCloudLDAPDirectoryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*jcapiv2.Configuration)
	client := jcapiv2.NewAPIClient(config)

	directories, _, err := client.DirectoriesApi.DirectoriesList(
		context.TODO(), "", "", nil)
	if err != nil {
		return diag.FromErr(err)
	}

	// there can only be a single GSuite directory per JumpCloud account
	for _, dir := range directories {
		if dir.Type_ == "ldap_server" {
			d.SetId(dir.Id)
			d.Set("name", dir.Name)
			d.Set("type", dir.Type_)
		}
		return nil
	}

	return diag.Errorf("couldn't find a directory with type 'ldap_server'")
}
