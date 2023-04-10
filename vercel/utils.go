package vercel

import (
	"context"
	"os"
	"strings"

	"github.com/chronark/vercel-go"
	"github.com/pkg/errors"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*vercel.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "vercel"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*vercel.Client), nil
	}

	// Default to the env var settings
	apiToken := os.Getenv("VERCEL_API_TOKEN")
	team := os.Getenv("VERCEL_TEAM")

	// Prefer config settings
	vercelConfig := GetConfig(d.Connection)
	if vercelConfig.APIToken != nil {
		apiToken = *vercelConfig.APIToken
	}
	if vercelConfig.Team != nil {
		team = *vercelConfig.Team
	}

	// Error if the minimum config is not set
	if apiToken == "" {
		return nil, errors.New("api_token must be configured")
	}

	config := vercel.NewClientConfig{Token: apiToken}
	if team != "" {
		config.Teamid = team
	}
	plugin.Logger(ctx).Debug("vercel.connect", "config", config)
	conn := vercel.New(config)

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}

func isNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "status code 404")
}

func isNotFoundOrUnauthorizedError(err error) bool {
	return strings.Contains(err.Error(), "Unable to fetch")
}
