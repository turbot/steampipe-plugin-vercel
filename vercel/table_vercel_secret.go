package vercel

import (
	"context"

	"github.com/chronark/vercel-go/endpoints/secret"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableVercelSecret(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "vercel_secret",
		Description: "Secrets in the Vercel account.",
		List: &plugin.ListConfig{
			Hydrate: listSecret,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getSecret,
			KeyColumns: plugin.AnyColumn([]string{"name", "uid"}),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the secret."},
			{Name: "uid", Type: proto.ColumnType_STRING, Transform: transform.FromField("Uid"), Description: "Unique identifier of the secret."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Created"), Description: "Time when the secret was created."},
			{Name: "team_id", Type: proto.ColumnType_STRING, Hydrate: getSecret, Description: "Unique identifier of the team the secret was created for."},
			{Name: "user_id", Type: proto.ColumnType_STRING, Hydrate: getSecret, Description: "Unique identifier of the user who created the secret."},
			{Name: "project_id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the project the secret belongs to."},
			{Name: "value", Type: proto.ColumnType_STRING, Description: "Value of the secret."},
			{Name: "decryptable", Type: proto.ColumnType_BOOL, Description: "True if the secret value can be decrypted after it is created."},
		},
	}
}

func listSecret(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("vercel_secret.listSecret", "connection_error", err)
		return nil, err
	}

	opts := secret.ListSecretsRequest{Limit: 100}
	for {
		plugin.Logger(ctx).Debug("vercel_secret.listSecret", "opts", opts)
		res, err := conn.Secret.ListSecrets(opts)
		if err != nil {
			plugin.Logger(ctx).Error("vercel_secret.listSecret", "query_error", err)
			return nil, err
		}
		for _, i := range res.Secrets {
			d.StreamListItem(ctx, i)
		}
		plugin.Logger(ctx).Debug("vercel_secret.listSecret", "pagination", res.Pagination)
		if res.Pagination.Next == 0 {
			break
		}
		opts.Until = int(res.Pagination.Next)
	}

	return nil, nil
}

func getSecret(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("vercel_secret.getSecret", "connection_error", err)
		return nil, err
	}

	var nameOrId string
	if h.Item != nil {
		s := h.Item.(secret.Secret)
		nameOrId = s.Uid
	} else if d.KeyColumnQuals["uid"] != nil {
		nameOrId = d.KeyColumnQuals["uid"].GetStringValue()
	} else if d.KeyColumnQuals["name"] != nil {
		nameOrId = d.KeyColumnQuals["name"].GetStringValue()
	} else {
		return nil, nil
	}

	res, err := conn.Secret.GetSecret(nameOrId)
	if err != nil {
		plugin.Logger(ctx).Error("vercel_team.getTeam", "query_error", err, "nameOrId", nameOrId)
	}

	return res, err
}
