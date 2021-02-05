package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			DataSourcesMap: map[string]*schema.Resource{
				"jumpcloud_gsuite_directory":    dataSourceJumpCloudGSuiteDirectory(),
				"jumpcloud_ldap_directory":      dataSourceJumpCloudLDAPDirectory(),
				"jumpcloud_office365_directory": dataSourceJumpCloudOffice365Directory(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"jumpcloud_application":            resourceApplication(),
				"jumpcloud_user":                   resourceUser(),
				"jumpcloud_user_group":             resourceUserGroup(),
				"jumpcloud_user_group_membership":  resourceUserGroupMembership(),
				"jumpcloud_user_group_association": resourceUserGroupAssociation(),
			},
			Schema: map[string]*schema.Schema{
				"api_key": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("JUMPCLOUD_API_KEY", nil),
					Description: "The admin API key to access JumpCloud resources",
				},
				"org_id": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("JUMPCLOUD_ORG_ID", nil),
					Description: "The JumpCloud organization ID",
				},
			},
		}

		p.UserAgent("terraform-provider-jumpcloud", version)
		p.ConfigureContextFunc = configure

		return p
	}
}

func configure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := Config{
		APIKey: d.Get("api_key").(string),
		OrgId:  d.Get("org_id").(string),
	}

	return config.Client(), nil
}
