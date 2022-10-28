package vercel

import (
	"context"

	"github.com/chronark/vercel-go/endpoints/deployment"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableVercelDeployment(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "vercel_deployment",
		Description: "Deployments in the Vercel account.",
		List: &plugin.ListConfig{
			Hydrate: listDeployment,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the deployment."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL of the deployment."},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "One of: BUILDING, ERROR, INITIALIZING, QUEUED, READY, CANCELED."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Created").Transform(transform.UnixMsToTimestamp), Description: "Time when the deployment was created."},
			{Name: "creator", Type: proto.ColumnType_JSON, Description: "Creator of the deployment."},
			{Name: "building_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("BuildingAt").Transform(transform.UnixMsToTimestamp), Description: "Time when deployment started to build."},
			{Name: "ready", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Ready").Transform(transform.UnixMsToTimestamp), Description: "Time when deployment is ready to view."},
			{Name: "meta", Type: proto.ColumnType_JSON, Description: "GitHub metadata associated with the deployment."},
		},
	}
}

func listDeployment(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("vercel_deployment.listDeployment", "connection_error", err)
		return nil, err
	}

	req := deployment.ListDeploymentsRequest{Limit: 100}
	limit := d.QueryContext.GetLimit()
	plugin.Logger(ctx).Warn("vercel_deployment.listDeployment", "queryContext", "limit", limit)
	if limit == -1 {
		limit = int64(req.Limit)
	}

	total := 0
	for {
		res, err := conn.Deployment.List(req)
		plugin.Logger(ctx).Debug("vercel_dns.listDeployment", "res", res)
		if err != nil {
			plugin.Logger(ctx).Error("vercel_domain.listDeployment", "query_error", err)
			return nil, err
		}
		for _, i := range res.Deployments {
			d.StreamListItem(ctx, i)
			total += 1
			if int64(total) == limit {
				break
			}
		}
		req.Until = int(res.Pagination.Next)
		if int64(total) == limit {
			break
		}
	}

	return nil, nil
}
