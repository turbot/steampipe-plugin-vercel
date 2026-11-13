package vercel

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type vercelConfig struct {
	APIToken *string `hcl:"api_token"`
	Team     *string `hcl:"team"`
}

func ConfigInstance() interface{} {
	return &vercelConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) vercelConfig {
	if connection == nil || connection.GetConfig() == nil {
		return vercelConfig{}
	}
	config, _ := connection.GetConfig().(vercelConfig)
	return config
}
