package provider

import (
	"context"
	"fmt"
	jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUserGroupAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserGroupAssociationCreate,
		Read:   resourceUserGroupAssociationRead,
		Update: nil,
		Delete: resourceUserGroupAssociationDelete,
		Schema: map[string]*schema.Schema{
			"group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"object_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errors []error) {
					allowedValues := []string{
						"active_directory",
						"application",
						"command",
						"g_suite",
						"ldap_server",
						"office_365",
						"policy",
						"radius_server",
						"system",
						"system_group",
					}

					v := val.(string)
					if !stringInSlice(v, allowedValues) {
						errors = append(errors, fmt.Errorf("%q must be one of %q", key, allowedValues))
					}
					return
				},
			},
		},
	}
}

func modifyUserGroupAssociation(client *jcapiv2.APIClient,
	d *schema.ResourceData, action string) error {

	payload := jcapiv2.UserGroupGraphManagementReq{
		Op:    action,
		Type_: d.Get("type").(string),
		Id:    d.Get("object_id").(string),
	}

	req := map[string]interface{}{
		"body": payload,
	}

	_, err := client.UserGroupAssociationsApi.GraphUserGroupAssociationsPost(
		context.TODO(), d.Get("group_id").(string), "", "", req)

	return err
}

func resourceUserGroupAssociationCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(*jcapiv2.Configuration)
	client := jcapiv2.NewAPIClient(config)

	err := modifyUserGroupAssociation(client, d, "add")
	if err != nil {
		return err
	}
	return resourceUserGroupAssociationRead(d, m)
}

func resourceUserGroupAssociationRead(d *schema.ResourceData, m interface{}) error {
	config := m.(*jcapiv2.Configuration)
	client := jcapiv2.NewAPIClient(config)

	optionals := map[string]interface{}{
		"groupId": d.Get("group_id").(string),
		"limit":   int32(100),
	}

	graphconnect, _, err := client.UserGroupAssociationsApi.GraphUserGroupAssociationsList(
		context.TODO(), d.Get("group_id").(string), "", "", []string{d.Get("type").(string)}, optionals)
	if err != nil {
		return err
	}

	// the ID of the specified object is buried in a complex construct
	for _, v := range graphconnect {
		if v.To.Id == d.Get("object_id") {
			resourceId := d.Get("group_id").(string) + "/" + d.Get("object_id").(string)
			d.SetId(resourceId)
			return nil
		}
	}

	// element does not exist; unset ID
	d.SetId("")
	return nil
}

func resourceUserGroupAssociationDelete(d *schema.ResourceData, m interface{}) error {
	config := m.(*jcapiv2.Configuration)
	client := jcapiv2.NewAPIClient(config)
	return modifyUserGroupAssociation(client, d, "remove")
}
