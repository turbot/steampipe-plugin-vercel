package vercel

import (
	"context"

	"github.com/chronark/vercel-go/endpoints/project"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableVercelProject(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "vercel_project",
		Description: "Projects in the Vercel account.",
		List: &plugin.ListConfig{
			Hydrate: listProject,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the project."},
			// Other columns
			{Name: "accountid", Type: proto.ColumnType_STRING, Description: "Account ID for the project."},
			{Name: "alias", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "analytics", Type: proto.ColumnType_JSON, Description: "Analytics information, if enabled for the project."},
			{Name: "auto_expose_system_envs", Type: proto.ColumnType_BOOL, Description: "If true then system environment variables are exposed for use."},
			{Name: "build_command", Type: proto.ColumnType_STRING, Description: "The build command for this project."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CreatedAt").Transform(transform.UnixMsToTimestamp), Description: "Time when the project was created."},
			{Name: "dev_command", Type: proto.ColumnType_STRING, Description: "The dev command for this project."},
			{Name: "directory_listing", Type: proto.ColumnType_BOOL, Description: "If true then the project is listed in the Vercel directory."},
			{Name: "env", Type: proto.ColumnType_JSON, Description: "Environment variables for the project."},
			{Name: "framework", Type: proto.ColumnType_STRING, Description: "Framework used in the project, e.g. nextjs."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Id"), Description: "ID of the project."},
			{Name: "install_command", Type: proto.ColumnType_STRING, Description: "The install command for this project."},
			{Name: "latest_deployments", Type: proto.ColumnType_JSON, Description: "Information about the latest deployments of the project."},
			{Name: "link", Type: proto.ColumnType_JSON, Description: "Details of the link from this project to a source code repository."},
			{Name: "live", Type: proto.ColumnType_BOOL, Description: "If true, the project is live."},
			{Name: "node_version", Type: proto.ColumnType_STRING, Description: "Node version used by the project, e.g. 16.x."},
			{Name: "output_directory", Type: proto.ColumnType_STRING, Description: "Directory where output of the build will go."},
			{Name: "password_protection", Type: proto.ColumnType_JSON, Description: "Password protection information, if enabled."},
			{Name: "permissions", Type: proto.ColumnType_JSON, Description: "Permissions settings."},
			{Name: "public_source", Type: proto.ColumnType_BOOL, Description: "If true, the project is linked to a public source."},
			{Name: "root_directory", Type: proto.ColumnType_STRING, Description: "Root directory for the build process."},
			{Name: "serverless_function_region", Type: proto.ColumnType_STRING, Description: "Region where serverless functions will be deployed."},
			{Name: "source_files_outside_root_directory", Type: proto.ColumnType_BOOL, Description: "If true then source files are outside the root directory."},
			{Name: "sso_protection", Type: proto.ColumnType_JSON, Description: "SSO protection information, if enabled."},
			{Name: "targets", Type: proto.ColumnType_JSON, Description: "Targets of the build."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("UpdatedAt").Transform(transform.UnixMsToTimestamp), Description: "Time when the project was last updated."},
		},
	}
}

func listProject(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("vercel_project.listProject", "connection_error", err)
		return nil, err
	}

	opts := project.ListProjectsRequest{Limit: 100}
	for {
		plugin.Logger(ctx).Debug("vercel_project.listProject", "opts", opts)
		res, err := conn.Project.List(opts)
		if err != nil {
			plugin.Logger(ctx).Error("vercel_project.listProject", "query_error", err)
			return nil, err
		}
		for _, i := range res.Projects {
			d.StreamListItem(ctx, i)
		}
		plugin.Logger(ctx).Debug("vercel_project.listProject", "pagination", res.Pagination)
		if res.Pagination.Next == 0 {
			break
		}
		opts.Until = res.Pagination.Next
	}

	return nil, nil
}
