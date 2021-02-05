package provider

import (
	jcapiv1 "github.com/TheJumpCloud/jcapi-go/v1"
	jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"log"
)

func resourceApplication() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource for adding an Amazon Web Services (AWS) account application. **Note:** This resource is due to change in future versions to be more generic and allow for adding various applications supported by JumpCloud.",
		Create:      resourceApplicationCreate,
		Read:        resourceApplicationRead,
		Update:      resourceApplicationUpdate,
		Delete:      resourceApplicationDelete,

		Schema: map[string]*schema.Schema{
			"display_label": {
				Description: "Name of the application to display",
				Type:        schema.TypeString,
				Required:    true,
			},
			"sso_url": {
				Description: "The SSO URL suffix to use",
				Type:        schema.TypeString,
				Required:    true,
			},
			"saml_role_attribute": {
				Description: "Value of the `https://aws.amazon.com/SAML/Attributes/Role` attribute.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"aws_session_duration": {
				Description: "Value of the `https://aws.amazon.com/SAML/Attributes/SessionDuration` attribute.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"metadata_xml": {
				Description: "The JumpCloud metadata XML file.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

// We receive a v2config from the TF base code but need a v1config to continue. So, we take the only
// preloaded element (the x-api-key) and populate the v1config with it.
//func convertV2toV1Config(v2config *jcapiv2.Configuration) *jcapiv1.Configuration {
//	configv1 := jcapiv1.NewConfiguration()
//	configv1.AddDefaultHeader("x-api-key", v2config.DefaultHeader["x-api-key"])
//	configv1.AddDefaultHeader("x-org-id", v2config.DefaultHeader["x-org-id"])
//	return configv1
//}

func resourceApplicationCreate(d *schema.ResourceData, m interface{}) error {
	configv1 := convertV2toV1Config(m.(*jcapiv2.Configuration))
	client := jcapiv1.NewAPIClient(configv1)

	payload := jcapiv1.Application{
		// TODO clearify if previous Active: true is translated to Beta: false
		Beta:         false,
		Name:         "aws",
		DisplayLabel: d.Get("display_label").(string),
		SsoUrl:       d.Get("sso_url").(string),
		Config: &jcapiv1.ApplicationConfig{
			ConstantAttributes: &jcapiv1.ApplicationConfigConstantAttributes{
				Value: []jcapiv1.ApplicationConfigConstantAttributesValue{
					{
						Name:  "https://aws.amazon.com/SAML/Attributes/SessionDuration",
						Value: d.Get("aws_session_duration").(string),
					},
					{
						Name:  "https://aws.amazon.com/SAML/Attributes/Role",
						Value: d.Get("saml_role_attribute").(string),
					},
				},
			},
		},
	}

	request := map[string]interface{}{
		"body": payload,
	}

	returnstruc, _, err := client.ApplicationsApi.ApplicationsPost(context.TODO(), request)
	if err != nil {
		return err
	}

	d.SetId(returnstruc.Id)
	return resourceApplicationRead(d, m)
}

func resourceApplicationRead(d *schema.ResourceData, m interface{}) error {
	configv1 := convertV2toV1Config(m.(*jcapiv2.Configuration))
	client := jcapiv1.NewAPIClient(configv1)

	response, _, err := client.ApplicationsApi.ApplicationsGet(context.TODO(), d.Id(), nil)

	// If the object does not exist, unset the ID
	if err != nil {
		if err.Error() == "EOF" {
			d.SetId("")
			return nil
		}
		return err
	}

	d.SetId(response.Id)

	if err := d.Set("display_label", response.DisplayLabel); err != nil {
		return err
	}
	if err := d.Set("sso_url", response.SsoUrl); err != nil {
		return err
	}

	constantAttributeValues := response.Config.ConstantAttributes.Value
	for _, el := range constantAttributeValues {
		if el.Name == "https://aws.amazon.com/SAML/Attributes/SessionDuration" {
			if err := d.Set("aws_session_duration", el.Value); err != nil {
				return err
			}
		}

		if el.Name == "https://aws.amazon.com/SAML/Attributes/Role" {
			if err := d.Set("saml_role_attribute", el.Value); err != nil {
				return err
			}
		}
	}

	if response.Id != "" {
		log.Println("[INFO] response ID is ", response.Id)
		orgId := configv1.DefaultHeader["x-org-id"]
		apiKey := configv1.DefaultHeader["x-api-key"]

		metadataXml, err := GetApplicationMetadataXml(orgId, response.Id, apiKey)
		if err != nil {
			return err
		}

		if err := d.Set("metadata_xml", metadataXml); err != nil {
			return err
		}
	} else {
		log.Println("[INFO] no ID in response, skipping metadata XML retrieval")
	}

	return nil
}

func resourceApplicationUpdate(d *schema.ResourceData, m interface{}) error {
	configv1 := convertV2toV1Config(m.(*jcapiv2.Configuration))
	client := jcapiv1.NewAPIClient(configv1)

	payload := jcapiv1.Application{
		// TODO clearify if previous Active: true is translated to Beta: false
		Beta:         false,
		Name:         "aws",
		DisplayLabel: d.Get("display_label").(string),
		SsoUrl:       d.Get("sso_url").(string),
		Config: &jcapiv1.ApplicationConfig{
			ConstantAttributes: &jcapiv1.ApplicationConfigConstantAttributes{
				Value: []jcapiv1.ApplicationConfigConstantAttributesValue{
					{
						Name:  "https://aws.amazon.com/SAML/Attributes/SessionDuration",
						Value: d.Get("aws_session_duration").(string),
					},
					{
						Name:  "https://aws.amazon.com/SAML/Attributes/Role",
						Value: d.Get("saml_role_attribute").(string),
					},
				},
			},
		},
	}

	request := map[string]interface{}{
		"body": payload,
	}

	_, _, err := client.ApplicationsApi.ApplicationsPut(context.TODO(), d.Id(), request)
	if err != nil {
		return err
	}
	return resourceApplicationRead(d, m)
}

func resourceApplicationDelete(d *schema.ResourceData, m interface{}) error {
	configv1 := convertV2toV1Config(m.(*jcapiv2.Configuration))
	client := jcapiv1.NewAPIClient(configv1)

	_, _, err := client.ApplicationsApi.ApplicationsDelete(context.TODO(), d.Id(), nil)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
