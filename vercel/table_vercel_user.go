package vercel

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableVercelUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "vercel_user",
		Description: "The currently authenticated user making the request.",
		List: &plugin.ListConfig{
			Hydrate: listUser,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the user."},
			{Name: "uid", Type: proto.ColumnType_STRING, Transform: transform.FromField("Uid"), Description: "Unique identifier of the user."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Email address of the user."},
			{Name: "username", Type: proto.ColumnType_STRING, Description: "Username of the user."},
			{Name: "staging_prefix", Type: proto.ColumnType_STRING, Description: "Username of the user."},
			{Name: "platform_version", Type: proto.ColumnType_INT, Description: ""},
			{Name: "bio", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "website", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "billing", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "profiles", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "date", Type: proto.ColumnType_TIMESTAMP, Description: "Date when the user was created."},
			{Name: "soft_block", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "remote_caching", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "resource_config", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "has_trial_available", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "import_flow_git_namespace", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "import_flow_git_namespace_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "import_flow_git_provider", Type: proto.ColumnType_STRING, Description: ""},
		},
	}
}

func listUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("vercel_user.listUser", "connection_error", err)
		return nil, err
	}

	res, err := conn.User.Get()
	if err != nil {
		plugin.Logger(ctx).Error("vercel_user.listUser", "query_error", err)
		return nil, err
	}

	plugin.Logger(ctx).Debug("vercel_user.listUser", "res", res)

	d.StreamListItem(ctx, res.User)

	return nil, nil
}
