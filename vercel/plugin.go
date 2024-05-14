package vercel

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-vercel",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		ConnectionKeyColumns: []plugin.ConnectionKeyColumn{
			{
				Name:    "user_uid",
				Hydrate: getUserUid,
			},
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError,
		},
		TableMap: map[string]*plugin.Table{
			"vercel_deployment": tableVercelDeployment(ctx),
			"vercel_dns_record": tableVercelDnsRecord(ctx),
			"vercel_domain":     tableVercelDomain(ctx),
			"vercel_project":    tableVercelProject(ctx),
			"vercel_secret":     tableVercelSecret(ctx),
			"vercel_team":       tableVercelTeam(ctx),
			"vercel_user":       tableVercelUser(ctx),
		},
	}
	return p
}
