package vercel

import (
	"context"

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

	plugin.Logger(ctx).Debug("vercel_deployment.listDeployment")
	res, err := conn.Deployment.List()
	if err != nil {
		plugin.Logger(ctx).Error("vercel_domain.listDomain", "query_error", err)
		return nil, err
	}
	for _, i := range res.Deployments {
		d.StreamListItem(ctx, i)
	}

	return nil, nil
}
