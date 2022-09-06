package vercel

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
)

type vercelConfig struct {
	APIToken *string `cty:"api_token"`
	Team     *string `cty:"team"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_token": {
		Type: schema.TypeString,
	},
	"team": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &vercelConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) vercelConfig {
	if connection == nil || connection.Config == nil {
		return vercelConfig{}
	}
	config, _ := connection.Config.(vercelConfig)
	return config
}
