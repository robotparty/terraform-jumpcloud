package provider

import (
	"context"
	"encoding/json"
	"net/http"

	jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUserGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides a JumpCloud user group resource. Refer to the [JumpCloud API model](https://docs.jumpcloud.com/2.0/models/usergroup) for further details.",
		CreateContext: resourceUserGroupCreate,
		ReadContext:   resourceUserGroupRead,
		UpdateContext: resourceUserGroupUpdate,
		DeleteContext: resourceUserGroupDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the group. Example: `My Group`.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"attributes": {
				// TODO add description
				Description: "TODO",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"posix_groups": {
							Type: schema.TypeString,
							// PosixGroups cannot be edited after group creation.
							ForceNew: true,
							Optional: true,
						},
					},
					// enable_samba has a more complicated lifecycle,
					// Commenting out for now as it is ignored in CRU by the JCAPI
					// From Jumpcloud UI:
					// Samba Authentication must be configured in the
					// JumpCloud LDAP Directory and LDAP sync must be enabled
					// on this group before Samba Authentication can be enabled.
					// "enable_samba": {
					// 	Type:     schema.TypeBool,
					// 	Optional: true,
					// },
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceUserGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*jcapiv2.Configuration)
	client := jcapiv2.NewAPIClient(config)

	body := jcapiv2.UserGroupPost{Name: d.Get("name").(string)}

	// For Attributes.PosixGroups, only the first member of the slice
	// is considered by the JCAPI
	if attr, ok := expandAttributes(d.Get("attributes")); ok {
		body.Attributes = attr
	}

	req := map[string]interface{}{
		"body": body,
	}
	group, res, err := client.UserGroupsApi.GroupsUserPost(context.TODO(),
		"", headerAccept, req)
	if err != nil {
		// TODO: sort out error essentials
		return diag.Errorf("error creating user group %s: %s - response = %+v",
			(req["body"].(jcapiv2.UserGroupPost)).Name, err, res)
	}

	d.SetId(group.Id)
	return resourceUserGroupRead(ctx, d, meta)
}

// resourceUserGroupRead uses a helper function that consumes the
// JC's HTTP API directly; the groups' attributes need to be kept in state
// as they are required for resourceUserGroupUpdate and the current
// implementation of the JC SDK doesn't support their retrieval
func resourceUserGroupRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*jcapiv2.Configuration)

	group, ok, err := userGroupReadHelper(config, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if !ok {
		// not found
		d.SetId("")
		return nil
	}

	d.SetId(group.ID)
	if err := d.Set("name", group.Name); err != nil {
		return diag.FromErr(err)
	}
	// TODO bring back after attributes are set correctly
	//if err := d.Set("attributes", flattenAttributes(&group.Attributes)); err != nil {
	//	return diag.FromErr(err)
	//}

	return nil
}

func userGroupReadHelper(config *jcapiv2.Configuration, id string) (ug *UserGroup,
	ok bool, err error) {

	req, err := http.NewRequest(http.MethodGet,
		config.BasePath+"/usergroups/"+id, nil)
	if err != nil {
		return
	}

	req.Header.Add("x-api-key", config.DefaultHeader["x-api-key"])
	if config.DefaultHeader["x-org-id"] != "" {
		req.Header.Add("x-org-id", config.DefaultHeader["x-org-id"])
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return
	}

	ok = true
	err = json.NewDecoder(res.Body).Decode(&ug)
	return
}

func resourceUserGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*jcapiv2.Configuration)
	client := jcapiv2.NewAPIClient(config)

	body := jcapiv2.UserGroupPost{Name: d.Get("name").(string)}
	if attr, ok := expandAttributes(d.Get("attributes")); ok {
		body.Attributes = attr
	} else {
		return diag.Errorf("Unable to update, attributes not expandable")
	}

	req := map[string]interface{}{
		"body": body,
	}
	// behaves like PUT, will fail if
	// attributes.posixGroups isn't sent, see GODOC
	_, res, err := client.UserGroupsApi.GroupsUserPatch(context.TODO(),
		d.Id(), "", headerAccept, req)
	if err != nil {
		// TODO: sort out error essentials
		return diag.Errorf("error deleting user group:%s; response = %+v", err, res)
	}

	return resourceUserGroupRead(ctx, d, meta)
}

func resourceUserGroupDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*jcapiv2.Configuration)
	client := jcapiv2.NewAPIClient(config)

	res, err := client.UserGroupsApi.GroupsUserDelete(context.TODO(),
		d.Id(), "", headerAccept, nil)
	if err != nil {
		// TODO: sort out error essentials
		return diag.Errorf("error deleting user group:%s; response = %+v", err, res)
	}

	d.SetId("")
	return nil
}
