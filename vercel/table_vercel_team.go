package vercel

import (
	"context"

	"github.com/chronark/vercel-go/endpoints/team"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableVercelTeam(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "vercel_team",
		Description: "Teams in the Vercel account.",
		List: &plugin.ListConfig{
			Hydrate: listTeam,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getTeam,
			KeyColumns: plugin.AnyColumn([]string{"id", "slug"}),
			// Not found for teams returns unauthorized
			ShouldIgnoreError: isNotFoundOrUnauthorizedError,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Unique identifier of the team."},
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "Slug of the team."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the team."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: ""},
			// Other columns
			{Name: "allow_project_transfers", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "avatar", Type: proto.ColumnType_STRING, Description: "Avatar for the team."},
			{Name: "billing", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the team was created."},
			{Name: "creator_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Creatorid"), Description: "ID of the user who created the team."},
			{Name: "invite_code", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "membership", Type: proto.ColumnType_JSON, Description: "Membership of the team."},
			{Name: "platform_version", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "preview_deployment_suffix", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "profiles", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "resource_config", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "soft_block", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "staging_prefix", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("UpdatedAt").Transform(transform.UnixMsToTimestamp), Description: "Time when the team was last updated."},
		},
	}
}

func listTeam(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("vercel_team.listTeam", "connection_error", err)
		return nil, err
	}

	opts := team.ListTeamsRequest{Limit: 100}
	for {
		plugin.Logger(ctx).Debug("vercel_team.listTeam", "opts", opts)
		res, err := conn.Team.ListTeams(opts)
		if err != nil {
			plugin.Logger(ctx).Error("vercel_team.listTeam", "query_error", err)
			return nil, err
		}
		for _, i := range res.Teams {
			d.StreamListItem(ctx, i)
		}
		plugin.Logger(ctx).Debug("vercel_team.listTeam", "pagination", res.Pagination)
		if res.Pagination.Next == 0 {
			break
		}
		opts.Until = int(res.Pagination.Next)
	}

	return nil, nil
}

func getTeam(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("vercel_team.getTeam", "connection_error", err)
		return nil, err
	}

	var req team.GetTeamRequest
	if h.Item != nil {
		s := h.Item.(team.Team)
		req.ID = s.ID
	} else if d.EqualsQuals["id"] != nil {
		req.ID = d.EqualsQuals["id"].GetStringValue()
	} else if d.EqualsQuals["slug"] != nil {
		req.Slug = d.EqualsQuals["slug"].GetStringValue()
	} else {
		return nil, nil
	}

	plugin.Logger(ctx).Debug("vercel_team.getTeam", "req", req)

	res, err := conn.Team.Get(req)
	if err != nil {
		plugin.Logger(ctx).Error("vercel_team.getTeam", "query_error", err, "req", req)
	}

	return res, err
}
