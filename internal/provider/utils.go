package provider

import (
	jcapiv1 "github.com/TheJumpCloud/jcapi-go/v1"
	jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"
	"github.com/go-resty/resty/v2"
	"log"
)

// Gets an application's metadata XML for SAML authentication
// this direct API call is a needed workaround since JumpCloud does not offer this endpoint through its SDK
func GetApplicationMetadataXml(orgId string, applicationId string, apiKey string) (string, error) {
	url := "https://console.jumpcloud.com/api/organizations/" + orgId + "/applications/" + applicationId + "/metadata.xml"

	// debug is always set to true, but output will only be shown if TF_LOG=DEBUG is set
	client := resty.New().SetDebug(true)

	resp, err := client.R().
		SetHeader("x-api-key", apiKey).
		Get(url)

	if err != nil {
		return "", err
	}

	log.Println("Status Code:", resp.StatusCode())
	log.Println("Status     :", resp.Status())
	log.Println("Time       :", resp.Time())
	log.Println("Received At:", resp.ReceivedAt())
	log.Println("Body       :\n", resp)

	return string(resp.Body()), nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// We receive a v2config from the TF base code but need a v1config to continue. So, we take the only
// preloaded element (the x-api-key) and populate the v1config with it.
func convertV2toV1Config(v2config *jcapiv2.Configuration) *jcapiv1.Configuration {
	const apiKeyHeader = "x-api-key"
	const orgIdHeader = "x-org-id"

	configv1 := jcapiv1.NewConfiguration()

	configv1.AddDefaultHeader(apiKeyHeader, v2config.DefaultHeader[apiKeyHeader])
	if v2config.DefaultHeader[orgIdHeader] != "" {
		configv1.AddDefaultHeader(orgIdHeader, v2config.DefaultHeader[orgIdHeader])
	}
	return configv1
}
