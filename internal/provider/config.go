package provider

import jcapiv2 "github.com/TheJumpCloud/jcapi-go/v2"

const (
	headerAccept = "application/json"
)

// Config holds the JC configuration
type Config struct {
	APIKey string // User specific auth token
	OrgId  string // Organization ID
}

// Client instantiates a jcapiv2.Configuration struct that is passed
// to every Resource operation
func (c *Config) Client() interface{} {
	config := jcapiv2.NewConfiguration()
	config.AddDefaultHeader("x-api-key", c.APIKey)
	config.AddDefaultHeader("x-org-id", c.OrgId)

	// Instantiate the API client
	return config
}
